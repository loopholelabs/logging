// SPDX-License-Identifier: Apache-2.0

package noop

import "github.com/loopholelabs/logging"

var _ logging.Logger = (*Logger)(nil)

type Logger struct{}

func New() *Logger {
	return new(Logger)
}

func (s *Logger) SetLevel(logging.Level) {}

func (s *Logger) SubLogger(string) logging.Logger { return s }

func (s *Logger) Fatal() logging.Event {
	return new(Event)
}

func (s *Logger) Error() logging.Event {
	return new(Event)
}

func (s *Logger) Warn() logging.Event {
	return new(Event)
}

func (s *Logger) Info() logging.Event {
	return new(Event)
}

func (s *Logger) Debug() logging.Event {
	return new(Event)
}

func (s *Logger) Trace() logging.Event {
	return new(Event)
}
