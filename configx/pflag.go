// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package configx

import (
	"github.com/knadh/koanf/providers/posflag"
	"github.com/ory/x/jsonschemax"
)

type PFlagProvider struct {
	p     *posflag.Posflag
	paths []jsonschemax.Path
}
