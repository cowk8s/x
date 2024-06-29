// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package otelx

import (
	"go.opentelemetry.io/otel/trace"

	"github.com/cowk8s/x/logrusx"
	"github.com/cowk8s/x/stringsx"
)

type Tracer struct {
	tracer trace.Tracer
}

// Creates a new tracer. If name is empty, a default tracer name is used
// instead. See: https://godocs.io/go.opentelemetry.io/otel/sdk/trace#TracerProvider.Tracer
func New(name string, l *logrusx.Logger, c *Config) (*Tracer, error) {
	t := &Tracer{}

	if err := t.setup(name, l, c); err != nil {
		return nil, err
	}

	return t, nil
}

// Creates a new no-op tracer.
func NewNoop(_ *logrusx.Logger, c *Config) *Tracer {
	tp := trace.NewNoopTracerProvider()
	t := &Tracer{tracer: tp.Tracer("")}
	return t
}

// setup constructs the tracer based on the given configuration.
func (t *Tracer) setup(name string, l *logrusx.Logger, c *Config) error {
	switch f := stringsx.SwitchExact(c.Provider); {
	case f.AddCase("jaeger"):
		tracer, err := SetupJaeger(t, name, c)
		if err != nil {
			return err
		}

		t.tracer = tracer
		l.Infof("Jaeger tracer configured! Sending spans to %s", c.Providers.Jaeger.LocalAgentAddress)
	case f.AddCase(""):
		l.Infof("No tracer configured - skipping tracing setup")
		t.tracer = trace.NewNoopTracerProvider().Tracer(name)
	default:
		return f.ToUnknownCaseErr()
	}

	return nil
}

// IsLoaded returns true if the tracer has been loaded.
func (t *Tracer) IsLoaded() bool {
	if t == nil || t.tracer == nil {
		return false
	}
	return true
}

// Tracer returns the underlying OpenTelemetry tracer.
func (t *Tracer) Tracer() trace.Tracer {
	return t.tracer
}
