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

package php_fpm

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "php_fpm", asset.ModuleFieldsPri, AssetPhpFpm); err != nil {
		panic(err)
	}
}

// AssetPhpFpm returns asset data.
// This is the base64 encoded gzipped contents of module/php_fpm.
func AssetPhpFpm() string {
	return "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"
}
