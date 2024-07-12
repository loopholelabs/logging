// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog"
	slogLogger "log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/loopholelabs/logging/loggers/slog"
)

var (
	zeroTime = time.Time{}
)

func init() {
	slog.ReplaceAttr = func(_ []string, a slogLogger.Attr) slogLogger.Attr {
		switch a.Key {
		case slogLogger.TimeKey:
			return slogLogger.Attr{}
		}
		return a
	}

	zerolog.TimestampFunc = func() time.Time {
		return zeroTime
	}
}

func fillZerologTestFields(t *testing.T, format string) string {
	return fmt.Sprintf(format, zeroTime.Format(zerolog.TimeFieldFormat), t.Name())
}

func fillSlogTestFields(t *testing.T, format string) string {
	return fmt.Sprintf(format, t.Name())
}

func TestInfo(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Run("noop", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Noop, t.Name(), out)
			log.Info().Msg("")
			assert.Equal(t, "", out.String())
		})

		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			log.Info().Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s\"}\n"), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, "test", out)
			log.Info().Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s\n"), out.String())
		})
	})

	t.Run("one-field", func(t *testing.T) {
		t.Run("noop", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Noop, t.Name(), out)
			log.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, "", out.String())
		})

		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			log.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s\",\"foo\":\"bar\"}\n"), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)
			log.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s foo=bar\n"), out.String())
		})
	})
	//
	//t.Run("two-field", func(t *testing.T) {
	//	out := &bytes.Buffer{}
	//	log := New(out)
	//	log.Info().
	//		Str("foo", "bar").
	//		Int("n", 123).
	//		Msg("")
	//	if got, want := decodeIfBinaryToString(out.Bytes()), `{"level":"info","foo":"bar","n":123}`+"\n"; got != want {
	//		t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
	//	}
	//})
}
