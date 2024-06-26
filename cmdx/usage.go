// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package cmdx

import (
	"bytes"
	"text/template"

	"github.com/spf13/cobra"
)

var usageTemplateFuncs = template.FuncMap{}

func TemplateCommandField(cmd *cobra.Command, field string) (string, error) {
	t := template.New("")
	t.Funcs(usageTemplateFuncs)
	t, err := t.Parse(field)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	if err := t.Execute(&out, cmd); err != nil {
		return "", err
	}
	return out.String(), nil
}
