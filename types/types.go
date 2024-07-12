// SPDX-License-Identifier: Apache-2.0

package types

import "net"

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

type Event interface {
	Str(key string, val string) Event
	Bool(key string, val bool) Event

	Int(key string, val int) Event
	Int8(key string, val int8) Event
	Int16(key string, val int16) Event
	Int32(key string, val int32) Event
	Int64(key string, val int64) Event

	Uint(key string, val uint) Event
	Uint8(key string, val uint8) Event
	Uint16(key string, val uint16) Event
	Uint32(key string, val uint32) Event
	Uint64(key string, val uint64) Event

	Float32(key string, val float32) Event
	Float64(key string, val float64) Event

	IPAddr(key string, ipAddr net.IP) Event
	MACAddr(key string, macAddr net.HardwareAddr) Event

	Err(err error) Event

	Msg(msg string)
	Msgf(format string, args ...interface{})
}
