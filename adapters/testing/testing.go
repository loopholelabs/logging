// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"io"
	"testing"
)

var _ io.Writer = (*Testing)(nil)

// Testing is an adapter for testing.TB
type Testing struct {
	t testing.TB
}

// New creates a new instance of *Testing
func New(t testing.TB) *Testing {
	return &Testing{t: t}
}

func (t *Testing) Write(p []byte) (n int, err error) {
	t.t.Log(string(p))
	return len(p), nil
}
