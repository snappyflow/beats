// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.elastic.co/apm"

	"github.com/snappyflow/beats/v7/libbeat/beat"
	"github.com/snappyflow/beats/v7/libbeat/beat/events"
	"github.com/snappyflow/beats/v7/libbeat/common"
	"github.com/snappyflow/beats/v7/libbeat/esleg/eslegclient"
	"github.com/snappyflow/beats/v7/libbeat/logp"
	"github.com/snappyflow/beats/v7/libbeat/outputs"
	"github.com/snappyflow/beats/v7/libbeat/outputs/outil"
	"github.com/snappyflow/beats/v7/libbeat/publisher"
	"github.com/snappyflow/beats/v7/libbeat/testing"
)

// Client is an elasticsearch client.
type Client struct {
	conn eslegclient.Connection

	index    outputs.IndexSelector
	pipeline *outil.Selector

	observer outputs.Observer

	log *logp.Logger
}

// ClientSettings contains the settings for a client.
type ClientSettings struct {
	eslegclient.ConnectionSettings
	Index    outputs.IndexSelector
	Pipeline *outil.Selector
	Observer outputs.Observer
}

type bulkResultStats struct {
	acked        int // number of events ACKed by Elasticsearch
	duplicates   int // number of events failed with `create` due to ID already being indexed
	fails        int // number of failed events (can be retried)
	nonIndexable int // number of failed events (not indexable -> must be dropped)
	tooMany      int // number of events receiving HTTP 429 Too Many Requests
}

const (
	defaultEventType = "doc"
)

// NewClient instantiates a new client.
func NewClient(
	s ClientSettings,
	onConnect *callbacksRegistry,
) (*Client, error) {
	pipeline := s.Pipeline
	if pipeline != nil && pipeline.IsEmpty() {
		pipeline = nil
	}

	conn, err := eslegclient.NewConnection(eslegclient.ConnectionSettings{
		URL:              s.URL,
		Username:         s.Username,
		Password:         s.Password,
		APIKey:           s.APIKey,
		Headers:          s.Headers,
		TLS:              s.TLS,
		Kerberos:         s.Kerberos,
		Proxy:            s.Proxy,
		ProxyDisable:     s.ProxyDisable,
		Parameters:       s.Parameters,
		CompressionLevel: s.CompressionLevel,
		EscapeHTML:       s.EscapeHTML,
		Timeout:          s.Timeout,
	})
	if err != nil {
		return nil, err
	}

	conn.OnConnectCallback = func() error {
		globalCallbackRegistry.mutex.Lock()
		defer globalCallbackRegistry.mutex.Unlock()

		for _, callback := range globalCallbackRegistry.callbacks {
			err := callback(conn)
			if err != nil {
				return err
			}
		}

		if onConnect != nil {
			onConnect.mutex.Lock()
			defer onConnect.mutex.Unlock()

			for _, callback := range onConnect.callbacks {
				err := callback(conn)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	client := &Client{
		conn:     *conn,
		index:    s.Index,
		pipeline: pipeline,

		observer: s.Observer,

		log: logp.NewLogger("elasticsearch"),
	}

	return client, nil
}

// Clone clones a client.
func (client *Client) Clone() *Client {
	// when cloning the connection callback and params are not copied. A
	// client's close is for example generated for topology-map support. With params
	// most likely containing the ingest node pipeline and default callback trying to
	// create install a template, we don't want these to be included in the clone.

	c, _ := NewClient(
		ClientSettings{
			ConnectionSettings: eslegclient.ConnectionSettings{
				URL:   client.conn.URL,
				Proxy: client.conn.Proxy,
				// Without the following nil check on proxyURL, a nil Proxy field will try
				// reloading proxy settings from the environment instead of leaving them
				// empty.
				ProxyDisable:      client.conn.Proxy == nil,
				TLS:               client.conn.TLS,
				Kerberos:          client.conn.Kerberos,
				Username:          client.conn.Username,
				Password:          client.conn.Password,
				APIKey:            client.conn.APIKey,
				Parameters:        nil, // XXX: do not pass params?
				Headers:           client.conn.Headers,
				Timeout:           client.conn.Timeout,
				CompressionLevel:  client.conn.CompressionLevel,
				OnConnectCallback: nil,
				Observer:          nil,
				EscapeHTML:        false,
			},
			Index:    client.index,
			Pipeline: client.pipeline,
		},
		nil, // XXX: do not pass connection callback?
	)
	return c
}

func (client *Client) Publish(ctx context.Context, batch publisher.Batch) error {
	events := batch.Events()
	rest, err := client.publishEvents(ctx, events)
	if len(rest) == 0 {
		batch.ACK()
	} else {
		batch.RetryEvents(rest)
	}
	return err
}

// PublishEvents sends all events to elasticsearch. On error a slice with all
// events not published or confirmed to be processed by elasticsearch will be
// returned. The input slice backing memory will be reused by return the value.
func (client *Client) publishEvents(ctx context.Context, data []publisher.Event) ([]publisher.Event, error) {
	span, ctx := apm.StartSpan(ctx, "publishEvents", "output")
	defer span.End()
	begin := time.Now()
	st := client.observer

	if st != nil {
		st.NewBatch(len(data))
	}

	if len(data) == 0 {
		return nil, nil
	}

	// encode events into bulk request buffer, dropping failed elements from
	// events slice
	origCount := len(data)
	span.Context.SetLabel("events_original", origCount)
	data, bulkItems := bulkEncodePublishRequest(client.log, client.conn.GetVersion(), client.index, client.pipeline, data)
	newCount := len(data)
	span.Context.SetLabel("events_encoded", newCount)
	if st != nil && origCount > newCount {
		st.Dropped(origCount - newCount)
	}
	if newCount == 0 {
		return nil, nil
	}

	status, result, sendErr := client.conn.Bulk(ctx, "", "", nil, bulkItems)
	if sendErr != nil {
		err := apm.CaptureError(ctx, fmt.Errorf("failed to perform any bulk index operations: %w", sendErr))
		err.Send()
		client.log.Error(err)
		return data, sendErr
	}
	pubCount := len(data)
	span.Context.SetLabel("events_published", pubCount)

	client.log.Debugf("PublishEvents: %d events have been published to elasticsearch in %v.",
		pubCount,
		time.Now().Sub(begin))

	// check response for transient errors
	var failedEvents []publisher.Event
	var stats bulkResultStats
	if status != 200 {
		failedEvents = data
		stats.fails = len(failedEvents)
	} else {
		failedEvents, stats = bulkCollectPublishFails(client.log, result, data)
	}

	failed := len(failedEvents)
	span.Context.SetLabel("events_failed", failed)
	if st := client.observer; st != nil {
		dropped := stats.nonIndexable
		duplicates := stats.duplicates
		acked := len(data) - failed - dropped - duplicates

		st.Acked(acked)
		st.Failed(failed)
		st.Dropped(dropped)
		st.Duplicate(duplicates)
		st.ErrTooMany(stats.tooMany)
	}

	if failed > 0 {
		if sendErr == nil {
			sendErr = eslegclient.ErrTempBulkFailure
		}
		return failedEvents, sendErr
	}
	return nil, nil
}

func interceptDocument(log *logp.Logger, index outputs.IndexSelector, event beat.Event) (bool, beat.Event) {
	var valueData map[string]interface{}
	processRequired := false
	newEvent := beat.Event{}

	// event.SetID()
	// beat, err := event.Private
	// log.Infof("Event meta  %+v", )
	indexName, err := index.Select(&event)
	newIndex := ""
	if err != nil {
		log.Debug("Index not found: Parse error")
	}
	log.Debugf("Index name %+v", indexName)

	msg, err := getEventMessage(log, indexName, &event)
	if err != nil {
		log.Infof("Dropping event: %+v", err)
	}
	json.Unmarshal(msg, &valueData)
	// log.Infof("Document data %+v", valueData)
	if labels, ok := valueData["labels"]; ok {

		profileId := labels.(map[string]interface{})["_tag_profileId"].(string)
		projectName, projFound := labels.(map[string]interface{})["_tag_projectName"].(string)
		redactBody := labels.(map[string]interface{})["_tag_redact_body"]

		// Dropping record if project name is not found
		if !projFound {
			return processRequired, newEvent
		}
		// Delete Cookies from headers
		event.Fields.Delete("http.request.headers.Cookies")
		event.Fields.Delete("http.response.headers.Cookies")

		var httpBodyString interface{} = nil
		httpRequestBodyFound := false

		if http, httpFound := valueData["http"]; httpFound {
			if request, requestFound := http.(map[string]interface{})["request"]; requestFound {
				httpBodyString = request.(map[string]interface{})["body"]
				if httpBodyString != nil {
					httpRequestBodyFound = true
				}
			}
		}

		// Only remove if redactBody Tag is present in labels
		if redactBody != nil && httpRequestBodyFound && strings.Contains(strings.ToLower(indexName), "trace") {
			indexType := "log"
			docType := "user-input"
			tagIndexType := labels.(map[string]interface{})["_tag_IndexType"]
			tagDocType := labels.(map[string]interface{})["_tag_documentType"]
			if tagIndexType != nil {
				if strings.Contains(strings.ToLower(tagIndexType.(string)), "metric") {
					indexType = "metric"
				}
			}
			if tagDocType != nil {
				docType = tagDocType.(string)
			}
			newIndex = indexType + "-" + profileId + "-" + getMetricName(projectName) + "-$_write"
			// newEvent := event
			newEvent.Meta = event.Meta.Clone()
			newEvent.Fields = event.Fields.Clone()
			newEvent.Private = event.Private
			newEvent.Timestamp = event.Timestamp
			// newEvent.Fields.Put()

			for key, _ := range valueData {
				newEvent.Fields.Delete(key)
			}
			newEvent.Meta.Put("snappyflow_index", newIndex)

			traceBody := make(map[string]interface{})
			if httpBodyString != nil {
				httpBodyString = httpBodyString.(map[string]interface{})["original"]

				log.Debugf("trace httpd body found : %+v ", httpBodyString)
				httpBody := map[string]interface{}{}
				switch v := httpBodyString.(type) {
				case string:
					if err := json.Unmarshal([]byte(httpBodyString.(string)), &httpBody); err == nil {
						newEvent.Fields.Put("body_json", httpBody)
						// traceBody["request_body"] = httpBody
					}
				case map[string]interface{}:
					// traceBody["request_body"] = httpBodyString
					newEvent.Fields.Put("body_json", httpBodyString)
				default:
					log.Debugf("Found unexpected type %+v", v)
				}

			} else {
				// If httpd body is not found do not send log data
				log.Debug("Trace http body not found")
			}
			traceBody["url"] = valueData["url"]
			traceBody["service"] = valueData["service"]
			traceBody["labels"] = valueData["labels"]
			delete(traceBody["labels"].(map[string]interface{}), "_tag_documentType")
			delete(traceBody["labels"].(map[string]interface{}), "_tag_IndexType")
			delete(traceBody["labels"].(map[string]interface{}), "_tag_redact_body")
			traceBody["processor"] = valueData["processor"]
			traceBody["source"] = valueData["source"]
			traceBody["agent"] = valueData["agent"]
			traceBody["status"], _ = event.Fields.GetValue("http.response.status_code")
			traceBody["user_agent"], _ = event.Fields.GetValue("user_agent")

			if valueData["user"] == nil {
				traceBody["user"] = make(map[string]interface{})
			} else {
				traceBody["user"] = valueData["user"]
			}
			// log.Debugf("new Index %+v", newIndex)
			// log.Debugf("Inside trace http body :")
			newEvent.Fields.Put("details_json", traceBody)
			newEvent.Fields.Put("time", int(valueData["timestamp"].(map[string]interface{})["us"].(float64))/1000)

			newEvent.Fields.Put("_plugin", "trace_body")
			newEvent.Fields.Put("_documentType", docType)

			newEvent.Fields.Put("_tag_projectName", labels.(map[string]interface{})["_tag_projectName"].(string))
			newEvent.Fields.Put("_tag_appName", labels.(map[string]interface{})["_tag_appName"].(string))

			newEvent.Fields.Put("message", "Trace request body")
			newEvent.Fields.Put("trace_id", valueData["trace"].(map[string]interface{})["id"])
			newEvent.Fields.Put("transaction_id", valueData["transaction"].(map[string]interface{})["id"])

			log.Debug("Deleteing http body from trace data")
			event.Fields.Delete("http.request.body")
			processRequired = true
		}
	}
	return processRequired, newEvent
}

// bulkEncodePublishRequest encodes all bulk requests and returns slice of events
// successfully added to the list of bulk items and the list of bulk items.
func bulkEncodePublishRequest(
	log *logp.Logger,
	version common.Version,
	index outputs.IndexSelector,
	pipeline *outil.Selector,
	data []publisher.Event,
) ([]publisher.Event, []interface{}) {
	var newData []publisher.Event
	for i := range data {
		event := data[i].Content
		// Drop event if profile id is not present
		_, err := event.Fields.GetValue("labels._tag_profileId")
		if err != nil {
			continue
		}
		// Drop metric event
		procEvent, err := event.Fields.GetValue("processor.event")
		if err == nil {
			if procEvent.(string) == "metric" {
				continue
			}
		}
		newData = append(newData, data[i])
		processRequired, newEvent := interceptDocument(log, index, event)
		if processRequired {
			pEvent := publisher.Event{}
			pEvent.Cache = data[i].Cache
			pEvent.Flags = data[i].Flags
			pEvent.Content = newEvent
			newData = append(newData, pEvent)
		}
	}
	okEvents := newData[:0]
	bulkItems := []interface{}{}

	// pEvent
	for i := range newData {
		event := &newData[i].Content
		meta, err := createEventBulkMeta(log, version, index, pipeline, event)
		if err != nil {
			log.Errorf("Failed to encode event meta data: %+v", err)
			continue
		}
		if opType := events.GetOpType(*event); opType == events.OpTypeDelete {
			// We don't include the event source in a bulk DELETE
			bulkItems = append(bulkItems, meta)
		} else {
			bulkItems = append(bulkItems, meta, event)
		}
		okEvents = append(okEvents, newData[i])
	}
	return okEvents, bulkItems
}

func createEventBulkMeta(
	log *logp.Logger,
	version common.Version,
	indexSel outputs.IndexSelector,
	pipelineSel *outil.Selector,
	event *beat.Event,
) (interface{}, error) {
	eventType := ""
	if version.Major == 6 {
		eventType = defaultEventType
	}

	pipeline, err := getPipeline(event, pipelineSel)
	if err != nil {
		err := fmt.Errorf("failed to select pipeline: %v", err)
		return nil, err
	}

	index, err := indexSel.Select(event)
	if index == "" {
		metricIndex, notFoundErr := event.Meta.GetValue("snappyflow_index")
		if notFoundErr == nil {
			index = metricIndex.(string)
		}
		if err != nil && notFoundErr != nil {
			err := fmt.Errorf("failed to select event index: %v", err)
			return nil, err
		}
	}
	// For rum agent - redirection to rum index based on project name
	agent, err := event.Fields.GetValue("labels._tag_agent")
	if err == nil {
		if agent == "rum" {
			profileId, _ := event.Fields.GetValue("labels._tag_profileId")
			projectName, _ := event.Fields.GetValue("labels._tag_projectName")
			index = fmt.Sprintf("rum-%s-%s-$_write", profileId.(string), getMetricName(projectName.(string)))
		}
	}

	// Only process for Snappyflow indexes
	canProcess := false
	indexArr := strings.Split(index, "-")
	validSFIndexes := map[string]string{
		"log":    "valid",
		"trace":  "valid",
		"metric": "valid",
	}
	if len(indexArr) > 0 {
		if _, ok := validSFIndexes[indexArr[0]]; ok {
			canProcess = true
		}
	}

	if !canProcess {
		err := fmt.Errorf("not a valid index: %v", index)
		return nil, err
	}

	id, _ := events.GetMetaStringValue(*event, events.FieldMetaID)
	opType := events.GetOpType(*event)

	meta := eslegclient.BulkMeta{
		Index:    index,
		DocType:  eventType,
		Pipeline: pipeline,
		ID:       id,
	}

	if opType == events.OpTypeDelete {
		if id != "" {
			return eslegclient.BulkDeleteAction{Delete: meta}, nil
		} else {
			return nil, fmt.Errorf("%s %s requires _id", events.FieldMetaOpType, events.OpTypeDelete)
		}
	}
	if id != "" || version.Major > 7 || (version.Major == 7 && version.Minor >= 5) {
		if opType == events.OpTypeIndex {
			return eslegclient.BulkIndexAction{Index: meta}, nil
		}
		return eslegclient.BulkCreateAction{Create: meta}, nil
	}
	return eslegclient.BulkIndexAction{Index: meta}, nil
}

func getPipeline(event *beat.Event, pipelineSel *outil.Selector) (string, error) {
	if event.Meta != nil {
		pipeline, err := events.GetMetaStringValue(*event, events.FieldMetaPipeline)
		if err == common.ErrKeyNotFound {
			return "", nil
		}
		if err != nil {
			return "", errors.New("pipeline metadata is no string")
		}

		return strings.ToLower(pipeline), nil
	}

	if pipelineSel != nil {
		return pipelineSel.Select(event)
	}
	return "", nil
}

// bulkCollectPublishFails checks per item errors returning all events
// to be tried again due to error code returned for that items. If indexing an
// event failed due to some error in the event itself (e.g. does not respect mapping),
// the event will be dropped.
func bulkCollectPublishFails(
	log *logp.Logger,
	result eslegclient.BulkResult,
	data []publisher.Event,
) ([]publisher.Event, bulkResultStats) {
	reader := newJSONReader(result)
	if err := bulkReadToItems(reader); err != nil {
		log.Errorf("failed to parse bulk response: %v", err.Error())
		return nil, bulkResultStats{}
	}

	count := len(data)
	failed := data[:0]
	stats := bulkResultStats{}
	for i := 0; i < count; i++ {
		status, msg, err := bulkReadItemStatus(log, reader)
		if err != nil {
			log.Error(err)
			return nil, bulkResultStats{}
		}

		if status < 300 {
			stats.acked++
			continue // ok value
		}

		if status == 409 {
			// 409 is used to indicate an event with same ID already exists if
			// `create` op_type is used.
			stats.duplicates++
			continue // ok
		}

		if status < 500 {
			if status == http.StatusTooManyRequests {
				stats.tooMany++
			} else {
				// hard failure, don't collect
				log.Warnf("Cannot index event %#v (status=%v): %s", data[i], status, msg)
				stats.nonIndexable++
				continue
			}
		}

		log.Debugf("Bulk item insert failed (i=%v, status=%v): %s", i, status, msg)
		stats.fails++
		failed = append(failed, data[i])
	}

	return failed, stats
}

func (client *Client) Connect() error {
	return client.conn.Connect()
}

func (client *Client) Close() error {
	return client.conn.Close()
}

func (client *Client) String() string {
	return "elasticsearch(" + client.conn.URL + ")"
}

func (client *Client) Test(d testing.Driver) {
	client.conn.Test(d)
}
