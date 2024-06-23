// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package flagx

import (
	"time"

	"github.com/spf13/pflag"

	"github.com/spf13/cobra"

	"github.com/cowk8s/x/cmdx"
)

func NewFlagSet(name string) *pflag.FlagSet {
	return pflag.NewFlagSet(name, pflag.ContinueOnError)
}

// MustGetBool returns a bool flag or fatals if an error occurs.
func MustGetBool(cmd *cobra.Command, name string) bool {
	ok, err := cmd.Flags().GetBool(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return ok
}

// MustGetString returns a string flag or fatals if an error occurs.
func MustGetString(cmd *cobra.Command, name string) string {
	s, err := cmd.Flags().GetString(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return s
}

// MustGetDuration returns a time.Duration flag or fatals if an error occurs.
func MustGetDuration(cmd *cobra.Command, name string) time.Duration {
	d, err := cmd.Flags().GetDuration(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return d
}
