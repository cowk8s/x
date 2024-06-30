// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package configx

import (
	"regexp"

	"github.com/ory/jsonschema/v3"
)

var isNumRegex = regexp.MustCompile("^[0-9]+$")

func NewKoanfEnv(prefix string, rawSchema []byte, schema *jsonschema.Schema) (*Env, error) {
	paths, err := 
}

// Env implements an environment variables provider.
type Env struct {
	prefix string
	paths  []jsonschemax.Path
}