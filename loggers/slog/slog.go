// SPDX-License-Identifier: Apache-2.0

package slog

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/loopholelabs/logging"
)

var _ logging.Logger = (*Logger)(nil)

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		a.Key = logging.TimestampKey
	}
	return a
}

type Logger struct {
	logger *slog.Logger
	level  *slog.LevelVar
	output io.Writer
	source string
}

func New(source string, level logging.Level, output io.Writer) *Logger {
	slogLevel := new(slog.LevelVar)
	s := newSlog(source, slogLevel, output)
	s.SetLevel(level)
	return s
}

func newSlog(source string, slogLevel *slog.LevelVar, output io.Writer) *Logger {
	sl := slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{
		Level:       slogLevel,
		ReplaceAttr: replaceAttr,
	}).WithAttrs([]slog.Attr{
		{Key: logging.SourceKey, Value: slog.StringValue(source)},
	}))
	s := &Logger{
		logger: sl,
		output: output,
		level:  slogLevel,
		source: source,
	}
	return s
}

func (s *Logger) SetLevel(level logging.Level) {
	var slogLevel slog.Level
	switch level {
	case logging.FatalLevel:
		slogLevel = slog.LevelError + 1
	case logging.ErrorLevel:
		slogLevel = slog.LevelError
	case logging.WarnLevel:
		slogLevel = slog.LevelWarn
	case logging.InfoLevel:
		slogLevel = slog.LevelInfo
	case logging.DebugLevel:
		slogLevel = slog.LevelDebug
	case logging.TraceLevel:
		slogLevel = slog.LevelDebug - 1
	}
	s.level.Set(slogLevel)
}

func (s *Logger) SubLogger(source string) logging.Logger {
	sloglevel := new(slog.LevelVar)
	sloglevel.Set(s.level.Level())
	return newSlog(fmt.Sprintf("%s:%s", s.source, source), sloglevel, s.output)
}

func (s *Logger) Fatal() logging.Event {
	return &Event{
		level:  slog.LevelError + 1,
		logger: s.logger,
	}
}

func (s *Logger) Error() logging.Event {
	return &Event{
		level:  slog.LevelError,
		logger: s.logger,
	}
}

func (s *Logger) Warn() logging.Event {
	return &Event{
		level:  slog.LevelWarn,
		logger: s.logger,
	}
}

func (s *Logger) Info() logging.Event {
	return &Event{
		level:  slog.LevelInfo,
		logger: s.logger,
	}
}

func (s *Logger) Debug() logging.Event {
	return &Event{
		level:  slog.LevelDebug,
		logger: s.logger,
	}
}

func (s *Logger) Trace() logging.Event {
	return &Event{
		level:  slog.LevelDebug - 1,
		logger: s.logger,
	}
}
