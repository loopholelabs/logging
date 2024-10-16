// SPDX-License-Identifier: Apache-2.0

package types

import (
	"errors"
	"net"
	"strings"
)

var (
	ErrInvalidLogLevel = errors.New("invalid log level")
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

func (l Level) String() string {
	switch l {
	case FatalLevel:
		return "FATAL"
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	case TraceLevel:
		return "TRACE"
	default:
		return "invalid"
	}
}

func (l Level) Type() string {
	return "log level"
}

func (l *Level) Set(v string) error {
	switch strings.ToLower(v) {
	case "fatal":
		*l = FatalLevel
	case "error":
		*l = ErrorLevel
	case "warn":
		*l = WarnLevel
	case "info":
		*l = InfoLevel
	case "debug":
		*l = DebugLevel
	case "trace":
		*l = TraceLevel
	default:
		return ErrInvalidLogLevel
	}

	return nil
}

const (
	TimestampKey = "time"
	ErrorKey     = "error"
	SourceKey    = "source"
)

type RootLogger interface {
	SetLevel(level Level)
	Logger
}

type Logger interface {
	Level() Level
	SubLogger(source string) Logger
	With() Context

	Fatal() Event
	Error() Event
	Warn() Event
	Info() Event
	Debug() Event
	Trace() Event
}

type Event interface {
	taggable[Event]

	Msg(msg string)
	Msgf(format string, args ...interface{})
}

type Context interface {
	taggable[Context]

	Logger() Logger
}

// taggable represents values that can receive structured fields.
type taggable[T any] interface {
	Str(key string, val string) T
	Bool(key string, val bool) T

	Int(key string, val int) T
	Int8(key string, val int8) T
	Int16(key string, val int16) T
	Int32(key string, val int32) T
	Int64(key string, val int64) T

	Uint(key string, val uint) T
	Uint8(key string, val uint8) T
	Uint16(key string, val uint16) T
	Uint32(key string, val uint32) T
	Uint64(key string, val uint64) T

	Float32(key string, val float32) T
	Float64(key string, val float64) T

	IPAddr(key string, ipAddr net.IP) T
	MACAddr(key string, macAddr net.HardwareAddr) T

	Err(err error) T
}
