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

package token

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/snappyflow/beats/v7/libbeat/common"
	"github.com/snappyflow/beats/v7/libbeat/common/bus"
)

func TestTokenAppender(t *testing.T) {
	tests := []struct {
		eventConfig string
		event       bus.Event
		result      common.MapStr
		config      string
	}{
		// Appender without a condition should apply the config regardless
		// Empty event config should return a config with only the headers
		{
			event: bus.Event{},
			result: common.MapStr{
				"headers": map[string]interface{}{
					"Authorization": "Bearer foo bar",
				},
			},
			eventConfig: "",
			config: `
token_path: "test"
`,
		},
		// Metricbeat module config should return a config that has headers section
		{
			event: bus.Event{},
			result: common.MapStr{
				"module": "prometheus",
				"hosts":  []interface{}{"1.2.3.4:8080"},
				"headers": map[string]interface{}{
					"Authorization": "Bearer foo bar",
				},
			},
			eventConfig: `
module: prometheus
hosts: ["1.2.3.4:8080"]
`,
			config: `
token_path: "test"
`,
		},
	}

	for _, test := range tests {
		config, err := common.NewConfigWithYAML([]byte(test.config), "")
		if err != nil {
			t.Fatal(err)
		}

		eConfig, err := common.NewConfigWithYAML([]byte(test.eventConfig), "")
		if err != nil {
			t.Fatal(err)
		}

		test.event["config"] = []*common.Config{eConfig}
		writeFile("test", "foo bar")

		appender, err := NewTokenAppender(config)
		assert.NoError(t, err)
		assert.NotNil(t, appender)

		appender.Append(test.event)
		cfgs, _ := test.event["config"].([]*common.Config)
		assert.Equal(t, len(cfgs), 1)

		out := common.MapStr{}
		cfgs[0].Unpack(&out)

		assert.Equal(t, out, test.result)
		deleteFile("test")
	}
}

func writeFile(name, message string) {
	ioutil.WriteFile(name, []byte(message), os.ModePerm)
}

func deleteFile(name string) {
	os.Remove(name)
}
