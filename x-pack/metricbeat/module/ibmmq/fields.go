// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package ibmmq

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "ibmmq", asset.ModuleFieldsPri, AssetIbmmq); err != nil {
		panic(err)
	}
}

// AssetIbmmq returns asset data.
// This is the base64 encoded gzipped contents of module/ibmmq.
func AssetIbmmq() string {
	return "eJxUzjEOgkAQheF+T/GHhooLTGFhZ0FhbSxARty4C+vuUHB7o5ioU07el/ca7roKvo/x4cC8BRXqw76lPdYOsgbtigq9Wudg0HLJPpmfJ2HnALYwcR6WoA6KmvlpLMKpupml6uzg6jUMRd6gYeqifltfZ2tSYczzkj6fX7Gp/y3PAAAA//8gwjWY"
}
