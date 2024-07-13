// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"io"
	"testing"

	"github.com/loopholelabs/logging/loggers/noop"
	"github.com/loopholelabs/logging/loggers/slog"
	"github.com/loopholelabs/logging/loggers/zerolog"

	"github.com/loopholelabs/logging/types"

	testingAdapter "github.com/loopholelabs/logging/adapters/testing"
)

type Kind int

const (
	Noop Kind = iota
	Zerolog
	Slog
)

// New creates a new logger based on the given kind, source, and output
// and a default level of Info.
func New(kind Kind, source string, output io.Writer) types.Logger {
	switch kind {
	case Noop:
		return noop.New(types.InfoLevel)
	case Zerolog:
		return zerolog.New(source, types.InfoLevel, output)
	case Slog:
		return slog.New(source, types.InfoLevel, output)
	default:
		return nil
	}
}

func Test(t testing.TB, kind Kind, source string) types.Logger {
	switch kind {
	case Noop:
		return noop.New(types.InfoLevel)
	case Zerolog:
		return zerolog.New(source, types.InfoLevel, testingAdapter.New(t))
	case Slog:
		return slog.New(source, types.InfoLevel, testingAdapter.New(t))
	default:
		return nil
	}
}
