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

package redis

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "redis", asset.ModuleFieldsPri, AssetRedis); err != nil {
		panic(err)
	}
}

// AssetRedis returns asset data.
// This is the base64 encoded gzipped contents of module/redis.
func AssetRedis() string {
	return "eJysk81u2zoQhfd+ioOs7gUSda9FgaJB0C6KAmkfQCw5pgYiOSpJWdXbF5Ls1HboNk47K5sjnu/MD+/Q0VQjkuG0ATJnRzVuHuf/NxvAUNKR+8wSarzdAMCSwycxg6MNsGVyJtVL6g5BefolN0eeeqphowz9/qSgeSpzLOXEPp2VxC4KrrGadWKxZUfpKHcOPPEvS2nHsZI7mkaJ5iz3G/4cX1taFCFb5Jb2njikrIKmCu9VwDeChOWLxquUKTa3aJJTO5p/6JadafDfViIe7x/evPv8gDFy5mCx5P6/fUaViCZRyBzINdVTulhvz+clreUqxyqdZXqV2xp9FE0pVc9verZRra3IcaAiz9GO3JVEJ7Yq3XsJz1NKypZHeplYvnWJd2AlJ+NfLu0XJ+Ncb4KKhEg5Mu3IYBvF77dnxwoKgfIosYOWEEjPUtULN1z78sRfveBavFfBgH6QHjKZqog1w9q7ajhv+Yp3Eux17A9Lq4IFZ2SRDlkOHpbHdvDFAZ51lERagkllexeewfWu5o58vD88+O8DxalM7Gj6t3PoaIIEjC3r9qQBo0p/GI6KtjyVV5tR0Q6eQk4YObcXTGnl3GzpZwAAAP//AoCyqA=="
}
