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

package memcache

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "memcache", asset.ModuleFieldsPri, AssetMemcache); err != nil {
		panic(err)
	}
}

// AssetMemcache returns asset data.
// This is the base64 encoded gzipped contents of protos/memcache.
func AssetMemcache() string {
	return "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"
}
