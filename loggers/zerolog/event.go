// SPDX-License-Identifier: Apache-2.0

package zerolog

import (
	"net"

	"github.com/rs/zerolog"

	"github.com/loopholelabs/logging/types"
)

var _ types.Event = (*Event)(nil)

type Event zerolog.Event

func (e *Event) Str(key string, val string) types.Event {
	return (*Event)((*zerolog.Event)(e).Str(key, val))
}

func (e *Event) Bool(key string, val bool) types.Event {
	return (*Event)((*zerolog.Event)(e).Bool(key, val))
}

func (e *Event) Int(key string, val int) types.Event {
	return (*Event)((*zerolog.Event)(e).Int(key, val))
}

func (e *Event) Int8(key string, val int8) types.Event {
	return (*Event)((*zerolog.Event)(e).Int8(key, val))
}

func (e *Event) Int16(key string, val int16) types.Event {
	return (*Event)((*zerolog.Event)(e).Int16(key, val))
}

func (e *Event) Int32(key string, val int32) types.Event {
	return (*Event)((*zerolog.Event)(e).Int32(key, val))
}

func (e *Event) Int64(key string, val int64) types.Event {
	return (*Event)((*zerolog.Event)(e).Int64(key, val))
}

func (e *Event) Uint(key string, val uint) types.Event {
	return (*Event)((*zerolog.Event)(e).Uint(key, val))
}

func (e *Event) Uint8(key string, val uint8) types.Event {
	return (*Event)((*zerolog.Event)(e).Uint8(key, val))
}

func (e *Event) Uint16(key string, val uint16) types.Event {
	return (*Event)((*zerolog.Event)(e).Uint16(key, val))
}

func (e *Event) Uint32(key string, val uint32) types.Event {
	return (*Event)((*zerolog.Event)(e).Uint32(key, val))
}

func (e *Event) Uint64(key string, val uint64) types.Event {
	return (*Event)((*zerolog.Event)(e).Uint64(key, val))
}

func (e *Event) Float32(key string, val float32) types.Event {
	return (*Event)((*zerolog.Event)(e).Float32(key, val))
}

func (e *Event) Float64(key string, val float64) types.Event {
	return (*Event)((*zerolog.Event)(e).Float64(key, val))
}

func (e *Event) IPAddr(key string, ipAddr net.IP) types.Event {
	return (*Event)((*zerolog.Event)(e).IPAddr(key, ipAddr))
}

func (e *Event) MACAddr(key string, macAddr net.HardwareAddr) types.Event {
	return (*Event)((*zerolog.Event)(e).MACAddr(key, macAddr))
}

func (e *Event) Err(err error) types.Event {
	return (*Event)((*zerolog.Event)(e).Err(err))
}

func (e *Event) Msg(msg string) {
	(*zerolog.Event)(e).Msg(msg)
}

func (e *Event) Msgf(format string, args ...interface{}) {
	(*zerolog.Event)(e).Msgf(format, args...)
}
