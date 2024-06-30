// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package configx

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func MergeAllTypes(src, dst map[string]interface{}) error {
	rawSrc, err := json.Marshal(src)
	if err != nil {
		return errors.WithStack(err)
	}

	dstSrc, err := json.Marshal(dst)
}
