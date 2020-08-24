// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/snappyflow/beats/v7/libbeat/logp"
	"github.com/snappyflow/beats/v7/x-pack/functionbeat/function/provider"
)

// Errors generated by the cliHandler.
var (
	errNoFunctionGiven = errors.New("no function given")
)

// cliHandler takes a provider.CLIManager and acts a bridge between user enterred content from the CLI
// and the type managing the function on the provider. It allow to specify multiple functions at
// the command line but will do a single invocation on the CLIManager and will do general validation
// and normalization of the values. It also communicate the status of the operations to the user.
//
// NOTES: Each execution call of the CLIManager are independant, this mean that a fail call will not
// stop other calls to succeed.
//
// TODO(ph) functions could be merged into a single call , but I thought it was premature to do
// it.
type cliHandler struct {
	clis      map[string]provider.CLIManager
	log       *logp.Logger
	errOutput io.Writer
	output    io.Writer

	functionsByProvider map[string]string
}

func newCLIHandler(clis map[string]provider.CLIManager, functionsByProvider map[string]string, errOutput io.Writer, output io.Writer) *cliHandler {
	return &cliHandler{
		clis:                clis,
		errOutput:           errOutput,
		functionsByProvider: functionsByProvider,
		output:              output,
		log:                 logp.NewLogger("cli-handler"),
	}
}

func (c *cliHandler) Deploy(names []string) error {
	c.log.Debugf("Starting deploy for: %s", strings.Join(names, ", "))
	defer c.log.Debug("Deploy execution ended")

	if len(names) == 0 {
		return errNoFunctionGiven
	}

	errCount := c.iterateCLIFuncVerbose(names, "deploy", provider.CLIManager.Deploy)

	if errCount > 0 {
		return fmt.Errorf("Fail to deploy %d function(s)", errCount)
	}
	return nil
}

func (c *cliHandler) Update(names []string) error {
	c.log.Debugf("Starting update for: %s", strings.Join(names, ", "))
	defer c.log.Debug("Update execution ended")

	if len(names) == 0 {
		return errNoFunctionGiven
	}

	errCount := c.iterateCLIFuncVerbose(names, "update", provider.CLIManager.Update)

	if errCount > 0 {
		return fmt.Errorf("fail to update %d function(s)", errCount)
	}
	return nil
}

func (c *cliHandler) Remove(names []string) error {
	c.log.Debugf("Starting remove for: %s", strings.Join(names, ", "))
	defer c.log.Debug("Remove execution ended")

	if len(names) == 0 {
		return errNoFunctionGiven
	}

	errCount := c.iterateCLIFuncVerbose(names, "remove", provider.CLIManager.Remove)

	if errCount > 0 {
		return fmt.Errorf("fail to remove %d function(s)", errCount)
	}
	return nil
}

func (c *cliHandler) Export(names []string) error {
	c.log.Debugf("Starting export for: %s", strings.Join(names, ", "))
	defer c.log.Debug("Export execution ended")

	if len(names) == 0 {
		return errNoFunctionGiven
	}

	errCount := c.iterateCLIFuncQuiet(names, "export", provider.CLIManager.Export)

	if errCount > 0 {
		return fmt.Errorf("fail to export %d function(s)", errCount)
	}
	return nil
}

func (c *cliHandler) Package(outputPattern string) error {
	for _, cli := range c.clis {
		err := cli.Package(outputPattern)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *cliHandler) iterateCLIFuncVerbose(names []string, operation string, f func(provider.CLIManager, string) error) int {
	return c.iterateCLIFunc(names, operation, f, true)
}

func (c *cliHandler) iterateCLIFuncQuiet(names []string, operation string, f func(provider.CLIManager, string) error) int {
	return c.iterateCLIFunc(names, operation, f, false)
}

func (c *cliHandler) iterateCLIFunc(names []string, operation string, f func(provider.CLIManager, string) error, verbose bool) int {
	errCount := 0
	for _, name := range names {
		providerName, ok := c.functionsByProvider[name]
		if !ok {
			fmt.Fprintf(c.errOutput, "Function: %s, could not be %sed. Enable it.\n", name, operation)
			errCount++
			continue
		}

		cli, ok := c.clis[providerName]
		if !ok {
			fmt.Fprintf(c.errOutput, "Function: %s, could not be %s. Selected provider '%s' cannot be found", name, operation, providerName)
			errCount++
			continue
		}

		err := f(cli, name)
		if err != nil {
			fmt.Fprintf(c.errOutput, "Function: %s, could not %s, error: %s\n", name, operation, err)
			errCount++
			continue
		}

		if verbose {
			fmt.Fprintf(c.output, "Function: %s, %s successful\n", name, operation)
		}
	}
	return errCount
}
