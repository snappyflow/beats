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

package jolokia

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "jolokia", asset.ModuleFieldsPri, AssetJolokia); err != nil {
		panic(err)
	}
}

// AssetJolokia returns asset data.
// This is the base64 encoded gzipped contents of module/jolokia.
func AssetJolokia() string {
	return "eJx8kFFOwzAQRP99ipG/6QX8wQG4AkLISjbptrbX8m5RensUkoARqPM5uzN6mhOudA+4SJIrRwcYW6IA/7I53gEj6dC4GksJeHYAsF+RZbwlcoCepdn7IGXiOWCKSVe3UaKoFDCv1UpmXGYNePWqyT/Bn82qf3PAxJRGDV/lJ5SYqYdaZfe6FjW51d35h2vTHsQgxSIXRSZrPChoqaI04oPj91Ocqdge7ik2km2dvHTtf0EewvyMdcnLQULWffxe6dDB8hkAAP//TGJ6CQ=="
}
