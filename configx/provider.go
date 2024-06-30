// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package configx

import (
	"context"
	"sync"
)

type tuple struct {
	Key   string
	Value interface{}
}

type Provider struct {
	l sync.RWMutex

	schema []byte
}

const (
	FlagConfig = "config"
	Delimiter  = "."
)

func New(ctx context.Context, schema []byte) (*Provider, error) {
	p := &Provider{
		schema: schema,
	}

	return p, nil
}
