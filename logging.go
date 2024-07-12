// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"io"
	"testing"

	testingAdapter "github.com/loopholelabs/logging/adapters/testing"
	"github.com/loopholelabs/logging/loggers/noop"
	"github.com/loopholelabs/logging/loggers/slog"
	"github.com/loopholelabs/logging/loggers/zerolog"
)

type Kind int

const (
	Noop Kind = iota
	Zerolog
	Slog
)

type Level int

const (
	FatalLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

const (
	TimestampKey = "time"
	ErrorKey     = "error"
	SourceKey    = "source"
)

type Logger interface {
	SetLevel(level Level)
	SubLogger(source string) Logger

	Fatal() Event
	Error() Event
	Warn() Event
	Info() Event
	Debug() Event
	Trace() Event
}

// New creates a new logger based on the given kind, source, and output
// and a default level of Info.
func New(kind Kind, source string, output io.Writer) Logger {
	switch kind {
	case Noop:
		return noop.New()
	case Zerolog:
		return zerolog.New(source, InfoLevel, output)
	case Slog:
		return slog.New(source, InfoLevel, output)
	default:
		return nil
	}
}

func NewTest(t testing.TB, kind Kind, source string) {
	switch kind {
	case Noop:
		noop.New()
	case Zerolog:
		zerolog.New(source, InfoLevel, testingAdapter.New(t))
	case Slog:
		slog.New(source, InfoLevel, testingAdapter.New(t))
	}
}
