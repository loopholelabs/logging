// SPDX-License-Identifier: Apache-2.0

package slog

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/loopholelabs/logging/types"
)

var _ types.Logger = (*Logger)(nil)

var (
	ReplaceAttr = func(_ []string, a slog.Attr) slog.Attr {
		switch a.Key {
		case slog.TimeKey:
			a.Key = types.TimestampKey
		}
		return a
	}
)

type Logger struct {
	logger *slog.Logger
	level  *slog.LevelVar
	output io.Writer
	source string
}

func New(source string, level types.Level, output io.Writer) *Logger {
	slogLevel := new(slog.LevelVar)
	s := newSlog(source, slogLevel, output)
	s.SetLevel(level)
	return s
}

func newSlog(source string, slogLevel *slog.LevelVar, output io.Writer) *Logger {
	sl := slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{
		Level:       slogLevel,
		ReplaceAttr: ReplaceAttr,
	}).WithAttrs([]slog.Attr{
		{Key: types.SourceKey, Value: slog.StringValue(source)},
	}))
	s := &Logger{
		logger: sl,
		output: output,
		level:  slogLevel,
		source: source,
	}
	return s
}

func (s *Logger) SetLevel(level types.Level) {
	var slogLevel slog.Level
	switch level {
	case types.FatalLevel:
		slogLevel = slog.LevelError + 1
	case types.ErrorLevel:
		slogLevel = slog.LevelError
	case types.WarnLevel:
		slogLevel = slog.LevelWarn
	case types.InfoLevel:
		slogLevel = slog.LevelInfo
	case types.DebugLevel:
		slogLevel = slog.LevelDebug
	case types.TraceLevel:
		slogLevel = slog.LevelDebug - 1
	}
	s.level.Set(slogLevel)
}

func (s *Logger) SubLogger(source string) types.Logger {
	sloglevel := new(slog.LevelVar)
	sloglevel.Set(s.level.Level())
	return newSlog(fmt.Sprintf("%s:%s", s.source, source), sloglevel, s.output)
}

func (s *Logger) Fatal() types.Event {
	return &Event{
		level:  slog.LevelError + 1,
		logger: s.logger,
	}
}

func (s *Logger) Error() types.Event {
	return &Event{
		level:  slog.LevelError,
		logger: s.logger,
	}
}

func (s *Logger) Warn() types.Event {
	return &Event{
		level:  slog.LevelWarn,
		logger: s.logger,
	}
}

func (s *Logger) Info() types.Event {
	return &Event{
		level:  slog.LevelInfo,
		logger: s.logger,
	}
}

func (s *Logger) Debug() types.Event {
	return &Event{
		level:  slog.LevelDebug,
		logger: s.logger,
	}
}

func (s *Logger) Trace() types.Event {
	return &Event{
		level:  slog.LevelDebug - 1,
		logger: s.logger,
	}
}