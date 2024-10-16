// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"bytes"
	"fmt"
	slogLogger "log/slog"
	"testing"
	"time"

	"github.com/rs/zerolog"

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

func fillZerologSubloggerTestFields(t *testing.T, format string, depth int) string {
	args := make([]interface{}, 0, depth+2)
	args = append(args, zeroTime.Format(zerolog.TimeFieldFormat), t.Name())
	for i := 0; i < depth; i++ {
		args = append(args, t.Name())
	}
	return fmt.Sprintf(format, args...)
}

func fillSlogTestFields(t *testing.T, format string) string {
	return fmt.Sprintf(format, t.Name())
}

func fillSlogSubloggerTestFields(t *testing.T, format string, depth int) string {
	args := make([]interface{}, 0, depth+1)
	args = append(args, t.Name())
	for i := 0; i < depth; i++ {
		args = append(args, t.Name())
	}
	return fmt.Sprintf(format, args...)
}

func TestInfo(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		t.Run("noop", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Noop, t.Name(), out)
			log.Info().Msg("")
			log.Debug().Msg("") // Should not be logged.
			assert.Equal(t, "", out.String())
		})

		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			log.Info().Msg("")
			log.Debug().Msg("") // Should not be logged.
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s\"}\n"), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)
			log.Info().Msg("")
			log.Debug().Msg("") // Should not be logged.
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s\n"), out.String())
		})
	})

	t.Run("one-field", func(t *testing.T) {
		t.Run("noop", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Noop, t.Name(), out)
			log.Info().Str("foo", "bar").Msg("")
			log.Debug().Str("foo", "bar").Msg("") // Should not be logged.
			assert.Equal(t, "", out.String())
		})

		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			log.Info().Str("foo", "bar").Msg("")
			log.Debug().Str("foo", "bar").Msg("") // Should not be logged.
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s\",\"foo\":\"bar\"}\n"), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)
			log.Info().Str("foo", "bar").Msg("")
			log.Debug().Str("foo", "bar").Msg("") // Should not be logged.
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s foo=bar\n"), out.String())
		})
	})

	t.Run("two-field", func(t *testing.T) {
		t.Run("noop", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Noop, t.Name(), out)
			log.Info().
				Str("foo", "bar").
				Int("n", 123).
				Msg("")
			log.Debug(). // Should not be logged.
					Str("foo", "bar").
					Int("n", 123).
					Msg("")
			assert.Equal(t, "", out.String())
		})

		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			log.Info().
				Str("foo", "bar").
				Int("n", 123).
				Msg("")
			log.Debug(). // Should not be logged.
					Str("foo", "bar").
					Int("n", 123).
					Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s\",\"foo\":\"bar\",\"n\":123}\n"), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)
			log.Info().
				Str("foo", "bar").
				Int("n", 123).
				Msg("")
			log.Debug(). // Should not be logged.
					Str("foo", "bar").
					Int("n", 123).
					Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s foo=bar n=123\n"), out.String())
		})
	})

	t.Run("with", func(t *testing.T) {
		t.Run("noop", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Noop, t.Name(), out)
			logger := log.With().
				Str("foo", "bar").
				Logger()
			logger.Info().Msg("")
			assert.Equal(t, "", out.String())
		})

		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)

			logger := log.With().
				Str("foo", "bar").
				Logger()
			logger.Info().Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"time\":\"%s\",\"source\":\"%s\"}\n"), out.String())

			// Log with per-message attribute.
			out.Reset()
			logger.Info().Int("n", 123).Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"time\":\"%s\",\"source\":\"%s\",\"n\":123}\n"), out.String())

			// Retain attributes on sublogger.
			out.Reset()
			logger2 := logger.With().
				Str("foo2", "bar2").
				Logger()
			logger2.Info().Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"foo2\":\"bar2\",\"time\":\"%s\",\"source\":\"%s\"}\n"), out.String())

			// Ensure original loggers were not modified.
			out.Reset()
			log.Info().Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s\"}\n"), out.String())

			out.Reset()
			logger.Info().Msg("")
			assert.Equal(t, fillZerologTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"time\":\"%s\",\"source\":\"%s\"}\n"), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)

			logger := log.With().
				Str("foo", "bar").
				Logger()
			logger.Info().Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s foo=bar\n"), out.String())

			// Log with per-message attribute.
			out.Reset()
			logger.Info().Int("n", 123).Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s foo=bar n=123\n"), out.String())

			// Retain attributes on sublogger.
			out.Reset()
			logger2 := logger.With().
				Str("foo2", "bar2").
				Logger()
			logger2.Info().Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s foo=bar foo2=bar2\n"), out.String())

			// Ensure original loggers were not modified.
			out.Reset()
			log.Info().Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s\n"), out.String())

			out.Reset()
			logger.Info().Msg("")
			assert.Equal(t, fillSlogTestFields(t, "level=INFO msg=\"\" source=%s foo=bar\n"), out.String())
		})
	})
}

func TestSubLoggers(t *testing.T) {
	t.Run("depth=1", func(t *testing.T) {
		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			sublogger := log.SubLogger(t.Name())
			sublogger.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillZerologSubloggerTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s:%s\",\"foo\":\"bar\"}\n", 1), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)
			sublogger := log.SubLogger(t.Name())
			sublogger.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillSlogSubloggerTestFields(t, "level=INFO msg=\"\" source=\"%s:%s\" foo=bar\n", 1), out.String())
		})
	})

	t.Run("depth=2", func(t *testing.T) {
		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			sublogger0 := log.SubLogger(t.Name())
			sublogger1 := sublogger0.SubLogger(t.Name())
			sublogger1.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillZerologSubloggerTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s:%s:%s\",\"foo\":\"bar\"}\n", 2), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)
			sublogger0 := log.SubLogger(t.Name())
			sublogger1 := sublogger0.SubLogger(t.Name())
			sublogger1.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillSlogSubloggerTestFields(t, "level=INFO msg=\"\" source=\"%s:%s:%s\" foo=bar\n", 2), out.String())
		})
	})

	t.Run("depth=3", func(t *testing.T) {
		t.Run("zerolog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Zerolog, t.Name(), out)
			sublogger0 := log.SubLogger(t.Name())
			sublogger1 := sublogger0.SubLogger(t.Name())
			sublogger2 := sublogger1.SubLogger(t.Name())
			sublogger2.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillZerologSubloggerTestFields(t, "{\"level\":\"info\",\"time\":\"%s\",\"source\":\"%s:%s:%s:%s\",\"foo\":\"bar\"}\n", 3), out.String())
		})

		t.Run("slog", func(t *testing.T) {
			out := &bytes.Buffer{}
			log := New(Slog, t.Name(), out)
			sublogger0 := log.SubLogger(t.Name())
			sublogger1 := sublogger0.SubLogger(t.Name())
			sublogger2 := sublogger1.SubLogger(t.Name())
			sublogger2.Info().Str("foo", "bar").Msg("")
			assert.Equal(t, fillSlogSubloggerTestFields(t, "level=INFO msg=\"\" source=\"%s:%s:%s:%s\" foo=bar\n", 3), out.String())
		})
	})

	t.Run("with", func(t *testing.T) {

		t.Run("before-depth=1", func(t *testing.T) {
			t.Run("zerolog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Zerolog, t.Name(), out).With().Str("foo", "bar").Logger()
				sublogger := log.SubLogger(t.Name())
				sublogger.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillZerologSubloggerTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"time\":\"%s\",\"source\":\"%s:%s\",\"foo1\":\"bar1\"}\n", 1), out.String())
			})

			t.Run("slog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Slog, t.Name(), out).With().Str("foo", "bar").Logger()
				sublogger := log.SubLogger(t.Name())
				sublogger.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillSlogSubloggerTestFields(t, "level=INFO msg=\"\" source=\"%s:%s\" foo=bar foo1=bar1\n", 1), out.String())
			})
		})

		t.Run("before-depth=2", func(t *testing.T) {
			t.Run("zerolog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Zerolog, t.Name(), out).With().Str("foo", "bar").Logger()
				sublogger0 := log.SubLogger(t.Name())
				sublogger1 := sublogger0.SubLogger(t.Name())
				sublogger1.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillZerologSubloggerTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"time\":\"%s\",\"source\":\"%s:%s:%s\",\"foo1\":\"bar1\"}\n", 2), out.String())
			})

			t.Run("slog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Slog, t.Name(), out).With().Str("foo", "bar").Logger()
				sublogger0 := log.SubLogger(t.Name())
				sublogger1 := sublogger0.SubLogger(t.Name())
				sublogger1.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillSlogSubloggerTestFields(t, "level=INFO msg=\"\" source=\"%s:%s:%s\" foo=bar foo1=bar1\n", 2), out.String())
			})
		})

		t.Run("after-depth=1", func(t *testing.T) {
			t.Run("zerolog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Zerolog, t.Name(), out)
				sublogger := log.SubLogger(t.Name()).With().Str("foo", "bar").Logger()
				sublogger.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillZerologSubloggerTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"time\":\"%s\",\"source\":\"%s:%s\",\"foo1\":\"bar1\"}\n", 1), out.String())
			})

			t.Run("slog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Slog, t.Name(), out)
				sublogger := log.SubLogger(t.Name()).With().Str("foo", "bar").Logger()
				sublogger.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillSlogSubloggerTestFields(t, "level=INFO msg=\"\" source=\"%s:%s\" foo=bar foo1=bar1\n", 1), out.String())
			})
		})

		t.Run("after-depth=2", func(t *testing.T) {
			t.Run("zerolog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Zerolog, t.Name(), out)
				sublogger0 := log.SubLogger(t.Name())
				sublogger1 := sublogger0.SubLogger(t.Name()).With().Str("foo", "bar").Logger()
				sublogger1.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillZerologSubloggerTestFields(t, "{\"level\":\"info\",\"foo\":\"bar\",\"time\":\"%s\",\"source\":\"%s:%s:%s\",\"foo1\":\"bar1\"}\n", 2), out.String())
			})

			t.Run("slog", func(t *testing.T) {
				out := &bytes.Buffer{}
				log := New(Slog, t.Name(), out)
				sublogger0 := log.SubLogger(t.Name())
				sublogger1 := sublogger0.SubLogger(t.Name()).With().Str("foo", "bar").Logger()
				sublogger1.Info().Str("foo1", "bar1").Msg("")
				assert.Equal(t, fillSlogSubloggerTestFields(t, "level=INFO msg=\"\" source=\"%s:%s:%s\" foo=bar foo1=bar1\n", 2), out.String())
			})
		})
	})
}
