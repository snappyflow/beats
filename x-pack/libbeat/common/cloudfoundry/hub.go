// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cloudfoundry

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/pkg/errors"

	"github.com/snappyflow/beats/v7/libbeat/logp"
)

// Client interface exposed by Hub.Client.
type Client interface {
	// GetAppByGuid returns the application from cloudfoundry.
	GetAppByGuid(guid string) (*cfclient.App, error)
	// StartJanitor keeps the cache of applications clean.
	StartJanitor(interval time.Duration)
	// StopJanitor stops the running janitor.
	StopJanitor()
}

// Hub is central place to get all the required clients to communicate with cloudfoundry.
type Hub struct {
	cfg       *Config
	userAgent string
	log       *logp.Logger
}

// NewHub creates a new hub to get the required clients to communicate with cloudfoundry.
func NewHub(cfg *Config, userAgent string, log *logp.Logger) *Hub {
	return &Hub{cfg, userAgent, log}
}

// Client returns the cloudfoundry client.
func (h *Hub) Client() (Client, error) {
	httpClient, insecure, err := h.httpClient()
	if err != nil {
		return nil, err
	}

	h.log.Debugw(
		"creating cloudfoundry ",
		"client_id", h.cfg.ClientID,
		"client_secret_present", h.cfg.ClientSecret != "",
		"api_address", h.cfg.APIAddress)
	cf, err := cfclient.NewClient(&cfclient.Config{
		ClientID:          h.cfg.ClientID,
		ClientSecret:      h.cfg.ClientSecret,
		ApiAddress:        h.cfg.APIAddress,
		HttpClient:        httpClient,
		SkipSslValidation: insecure,
		UserAgent:         h.userAgent,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error creating cloudfoundry client")
	}
	if h.cfg.DopplerAddress != "" {
		cf.Endpoint.DopplerEndpoint = h.cfg.DopplerAddress
	}
	if h.cfg.UaaAddress != "" {
		cf.Endpoint.AuthEndpoint = h.cfg.UaaAddress
	}
	return newClientCacheWrap(cf, h.cfg.CacheDuration, h.cfg.CacheRetryDelay, h.log), nil
}

// RlpListener returns a listener client that calls the passed callback when the provided events are streamed through
// the loggregator to this client.
func (h *Hub) RlpListener(callbacks RlpListenerCallbacks) (*RlpListener, error) {
	client, err := h.Client()
	if err != nil {
		return nil, err
	}
	return h.RlpListenerFromClient(client, callbacks)
}

// RlpListener returns a listener client that calls the passed callback when the provided events are streamed through
// the loggregator to this client.
//
// In the case that the cloudfoundry client was already needed by the code path, call this method
// as not to create a intermediate client that will not be used.
func (h *Hub) RlpListenerFromClient(client Client, callbacks RlpListenerCallbacks) (*RlpListener, error) {
	var rlpAddress string
	if h.cfg.RlpAddress != "" {
		rlpAddress = h.cfg.RlpAddress
	} else {
		rlpAddress = strings.Replace(h.cfg.APIAddress, "api", "log-stream", 1)
	}
	doer, err := h.doerFromClient(client)
	if err != nil {
		return nil, err
	}
	return newRlpListener(rlpAddress, doer, h.cfg.ShardID, callbacks, h.log), nil
}

func (h *Hub) DopplerConsumer(callbacks DopplerCallbacks) (*DopplerConsumer, error) {
	client, err := h.Client()
	if err != nil {
		return nil, err
	}
	return h.DopplerConsumerFromClient(client, callbacks)
}

func (h *Hub) DopplerConsumerFromClient(client Client, callbacks DopplerCallbacks) (*DopplerConsumer, error) {
	dopplerAddress := h.cfg.DopplerAddress
	if dopplerAddress == "" {
		endpoint, err := cfEndpoint(client)
		if err != nil {
			return nil, errors.Wrap(err, "getting endpoints from client")
		}
		dopplerAddress = endpoint.DopplerEndpoint
	}
	httpClient, _, err := h.httpClient()
	if err != nil {
		return nil, errors.Wrap(err, "getting http client")
	}

	// TODO: Refactor Client so it is easier to access the cfclient
	ccw, ok := client.(*clientCacheWrap)
	if !ok {
		return nil, fmt.Errorf("client without cache wrap")
	}
	cfc, ok := ccw.client.(*cfclient.Client)
	if !ok {
		return nil, fmt.Errorf("client is not a cloud foundry client")
	}

	tr := TokenRefresherFromCfClient(cfc)
	return newDopplerConsumer(dopplerAddress, h.cfg.ShardID, h.log, httpClient, tr, callbacks)
}

// doerFromClient returns an auth token doer using uaa.
func (h *Hub) doerFromClient(client Client) (*authTokenDoer, error) {
	httpClient, _, err := h.httpClient()
	if err != nil {
		return nil, err
	}
	url := h.cfg.UaaAddress
	if url == "" {
		endpoint, err := cfEndpoint(client)
		if err != nil {
			return nil, errors.Wrap(err, "getting endpoints from client")
		}
		url = endpoint.AuthEndpoint
	}
	return newAuthTokenDoer(url, h.cfg.ClientID, h.cfg.ClientSecret, httpClient, h.log), nil
}

// httpClient returns an HTTP client configured with the configuration TLS.
func (h *Hub) httpClient() (*http.Client, bool, error) {
	tls, err := h.cfg.TLSConfig()
	if err != nil {
		return nil, true, err
	}
	httpClient := cfclient.DefaultConfig().HttpClient
	tp := defaultTransport()
	tp.TLSClientConfig = tls
	httpClient.Transport = tp
	return httpClient, tls.InsecureSkipVerify, nil
}

func cfEndpoint(client Client) (cfclient.Endpoint, error) {
	ccw, ok := client.(*clientCacheWrap)
	if !ok {
		return cfclient.Endpoint{}, fmt.Errorf("client without cache wrap")
	}
	cfc, ok := ccw.client.(*cfclient.Client)
	if !ok {
		return cfclient.Endpoint{}, fmt.Errorf("client is not a cloud foundry client")
	}
	return cfc.Endpoint, nil
}

// defaultTransport returns a new http.Transport for http.Client
func defaultTransport() *http.Transport {
	defaultTransport := http.DefaultTransport.(*http.Transport)
	return &http.Transport{
		Proxy:                 defaultTransport.Proxy,
		TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
		ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
	}
}
