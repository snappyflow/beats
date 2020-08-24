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

package beat

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "beat", asset.ModuleFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded gzipped contents of module/beat.
func AssetBeat() string {
	return "eJzcl0+L2zwQxu/5FIPP7+YD+PAWSlvYSwtloYdSimxPvGJljSuNdvG3L5LXiWPJdtJmE6hOiRU/z0/zR1Lu4Am7HAoUvAFgyQpzyN6j4GwDUKEtjWxZks7h/w0AgJ+ChiqncANgUKGwmEMtNgAWmaWubQ7fM2tV9h9kj8xt9sPPPZLhnyXpnaxz2All/fs7iaqyeVC+Ay0a3LP4wV3rtQ259vVJguhYZawkq/2jQesJuxcy4+dJxX6Etd5/2EbCXusC0v69WNyy4Fh9HIRTtKcqx5kaxjRyY45GaFFjg5q3qEWhsDr62QBWECkUejK3gOfHvYUSNRuhRjbwavMujRNqbluS05wEkZqxRnMeyGfXFGiAdq/6FlJrHRjIcet4678kEeISOAVBNOgBenFwFisoupDFJMQvhw7fiCFoQ4FS12mQcZHaixSp/asidS3LBreNTcZCka7PC0TA6kWTfsbpYFiTIcdS46WMD4V4kPZu2qdC6vlyULIopnNz+TiB41MINpTUNKSBCYRSwXy60FRWIOqVaHoJ7QQ8P770jTItnjUwmN/AY7x0F50B6cdD1446exEIn1EnVwOrETsD6KN3gbCFoplzW4rfGFmUT9GREFMnOuEPwOGoQ/pgeQBNLwqreoHjQMvyeS7nb4vbWy+neHT9EVw+RtvKdUgD4irBQFoZatvblsAawh7VtUqWgm8ZV3ugWAfeCRlfuK4Ku0Kw30mJGqG7m5Aykb8+dqe2FhMLdUXSB+8HehLZxXPAoJiL+cVOga8oqtnTE844AYruVg0VnJeCNTpZjaHZs+5tKT3fkv+A+GIkL19HLpD1b97kn0i7Dxfj9B9nzHnDzIeMTgGWPvwOAAD//2j1+Zk="
}
