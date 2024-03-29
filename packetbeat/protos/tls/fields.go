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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package tls

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "tls", asset.ModuleFieldsPri, AssetTls); err != nil {
		panic(err)
	}
}

// AssetTls returns asset data.
// This is the base64 encoded gzipped contents of protos/tls.
func AssetTls() string {
	return "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"
}
