// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package logrusx

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/sirupsen/logrus"
)

type (
	options struct {
		l         *logrus.Logger
		level     *logrus.Level
		formatter logrus.Formatter

		exitFunc func(int)
		hooks    []logrus.Hook
	}
)

//go:embed config.schema.json
var ConfigSchema string

const ConfigSchemaID = "cowk8s://logging-config"

// AddConfigSchema adds the logging schema to the compiler.
// The interface is specified instead of `jsonschema.Compiler` to allow the use of any jsonschema library fork or version.
func AddConfigSchema(c interface {
	AddResource(url string, r io.Reader) error
}) error {
	return c.AddResource(ConfigSchemaID, bytes.NewBufferString(ConfigSchema))
}

func newLogger(parent *logrus.Logger, o *options) *logrus.Logger {
	l := parent
	if l == nil {
		l = logrus.New()
	}

	if o.exitFunc != nil {
		l.ExitFunc = o.exitFunc
	}

	for _, hook := range o.hooks {
		l.AddHook(hook)
	}

	return l
}
