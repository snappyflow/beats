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

// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"github.com/snappyflow/beats/v7/libbeat/common"
)

// Config stores the configuration of Journalbeat
type Config struct {
	Inputs       []*common.Config `config:"inputs"`
	RegistryFile string           `config:"registry_file"`
}

var (
	// DefaultConfig are the defaults of a Journalbeat instance
	DefaultConfig = Config{
		RegistryFile: "registry",
	}
)
