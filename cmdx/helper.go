// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package cmdx

import (
	"fmt"
	"os"
)

// Fatalf prints to os.Stderr and exists with code 1.
func Fatalf(message string, args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, message)
	}
	os.Exit(1)
}
