// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"github.com/snappyflow/beats/v7/libbeat/cmd"

	// register central management
	"github.com/snappyflow/beats/v7/x-pack/libbeat/licenser"
	_ "github.com/snappyflow/beats/v7/x-pack/libbeat/management"

	// Register fleet
	_ "github.com/snappyflow/beats/v7/x-pack/libbeat/management/fleet"
	// register processors
	_ "github.com/snappyflow/beats/v7/x-pack/libbeat/processors/add_cloudfoundry_metadata"

	// register autodiscover providers
	_ "github.com/snappyflow/beats/v7/x-pack/libbeat/autodiscover/providers/aws/ec2"
	_ "github.com/snappyflow/beats/v7/x-pack/libbeat/autodiscover/providers/aws/elb"
)

// AddXPack extends the given root folder with XPack features
func AddXPack(root *cmd.BeatsRootCmd, name string) {
	licenser.Enforce(name, licenser.BasicAndAboveOrTrial)
	root.AddCommand(genEnrollCmd(name, ""))
}
