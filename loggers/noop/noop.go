// SPDX-License-Identifier: Apache-2.0

package noop

import "github.com/loopholelabs/logging/types"

var _ types.Logger = (*Logger)(nil)

type Logger struct {
	level types.Level
}

func New(level types.Level) *Logger {
	return &Logger{level: level}
}

func (s *Logger) SetLevel(level types.Level) {
	s.level = level
}

func (s *Logger) Level() types.Level {
	return s.level
}

func (s *Logger) SubLogger(string) types.SubLogger { return s }

func (s *Logger) Fatal() types.Event {
	return new(Event)
}

func (s *Logger) Error() types.Event {
	return new(Event)
}

func (s *Logger) Warn() types.Event {
	return new(Event)
}

func (s *Logger) Info() types.Event {
	return new(Event)
}

func (s *Logger) Debug() types.Event {
	return new(Event)
}

func (s *Logger) Trace() types.Event {
	return new(Event)
}
