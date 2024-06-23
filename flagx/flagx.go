// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package flagx

import (
	"time"

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