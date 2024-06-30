// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package configx

import (
	"path/filepath"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/v2"

	"github.com/pkg/errors"
)

// KoanfFile implements a KoanfFile provider.
type KoanfFile struct {
	subKey string
	path   string
	parser koanf.Parser
}

// NewKoanfFile returns a file provider.
func NewKoanfFile(path string) (*KoanfFile, error) {
	return NewKoanfFileSubKey(path, "")
}

func NewKoanfFileSubKey(path, subKey string) (*KoanfFile, error) {
	kf := &KoanfFile{
		path:   filepath.Clean(path),
		subKey: subKey,
	}

	switch e := filepath.Ext(path); e {
	case ".toml":
		kf.parser = toml.Parser()
	case ".json":
		kf.parser = json.Parser()
	case ".yaml", ".yml":
		kf.parser = yaml.Parser()
	default:
		return nil, errors.Errorf("unknown config file extension: %s", e)
	}

	return kf, nil
}
