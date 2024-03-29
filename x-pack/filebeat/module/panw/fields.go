// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package panw

import (
	"github.com/snappyflow/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "panw", asset.ModuleFieldsPri, AssetPanw); err != nil {
		panic(err)
	}
}

// AssetPanw returns asset data.
// This is the base64 encoded gzipped contents of module/panw.
func AssetPanw() string {
	return "eJzMmE1v4zYQhu/5FXNzC8S59ZJDgbSLAAHS1NjdoMcFTY4s1hRHOxzF1f76grQsa2060jrZILpZCvk+8/HyI3NYY3sNtfKbCwCx4rD/ZTBotrVY8tfw+wUAwF9kGodQEMNCOYIbJwQPKBvidYBfFjcP878//XoBUFh0JlynQXPwqtpPGx9pa7yGFVNTd28yYvG5TfNAwVSBlJjmgCpRXHV/NJSCAz0K+9c51eekv9MnTvKZoLcxg6NVuBqOPcIakHHjMKB8L9XhrbHdEJuDb88xAsCDqhCoSIxxcpBSCVRKdIkGpLQBAoZgyV9lgQI1rDHLc5SucZouaUKA/wl6k7CE6rnDJ3SdGNDyX9RydTA6l7Yh6Tfyh5xjuZtAHJ9PW6wo0NX7VNaGPNYLcqGOkve6UL3KD5B5ddhg8GxRJxItKMj84ebzrozKGMYQLsEWu1fxqw1QIxfEFZpjxtN1HmQ2B7gL4MTHCfzHEdwtcoD9KkKcy+MOxJFfvR5KFNvDZK1qMIj1Ks77Rn4dKL47034YsL0v5w7J3qN9h1Udenj4/mVGHrXyiJkneuhETHlXT/D1qLNfxjXB4ugNowonDH7WAeFzPBikOUHteg8EuYpYw8Jmgfz2sPNG602nduZaU2tVf7E5076CqRdKr1FAq1oaRrj7kAytQEpGdaqs8EJXT3GYpqpqvJU2H/qU8CemID5/7tRSBhxt5qUKZX9Kji3/21yaen9IP9FYhXVvdeqMUme2VAzuvH75w3rFLfTZ2XXKliagl8i7RFBeufYb5uuybFMs/1hnbi3HcfxkNebWt3yRs7lv2L1R6ht2Z2ZeK8EVcftz3Hyb+jXV4/Hjfdz+ZBYS++PH+147v43EsbuCXKYxT8jGaoG4sJa4rbDyBmzIToBWSmSYVcpZbakJs0uYrVi1G8U4uwRimC3R25WfjZnI0ebY9i/YLO7iccUrB76pkK0Ga9CLLSxy6mJUujw+wOQvlvi1Qa/xi2+qJXKWMbPRjgDe0wrQC7dDsnTltQGs14wVekHTyYtVzmXq+Ojt1wb3ITlaJaSRmLrFnvGZW/NZeY9dR7ztnCg1ZXM5gHrNNsj8s+OgEZLNf4Au/vipfH3WcmRDIKVPXqCmsMyPYG7ShCBqjb4n6D3yfwAAAP//PRbQnA=="
}
