// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package logrusx

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}
