// SPDX-License-Identifier: Apache-2.0

package noop

import "github.com/loopholelabs/logging/types"

var _ types.Logger = (*Logger)(nil)

type Logger struct{}

func New() *Logger {
	return new(Logger)
}

func (s *Logger) SetLevel(types.Level) {}

func (s *Logger) SubLogger(string) types.Logger { return s }

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
