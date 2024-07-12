// SPDX-License-Identifier: Apache-2.0

package slog

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/loopholelabs/logging"
)

var _ logging.Event = (*Event)(nil)

var (
	defaultContext = context.Background()
)

type Event struct {
	level  slog.Level
	logger *slog.Logger
	attr   []slog.Attr
}

func (e *Event) Str(key string, val string) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.StringValue(val),
	})
	return e
}

func (e *Event) Bool(key string, val bool) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.BoolValue(val),
	})
	return e
}

func (e *Event) Int(key string, val int) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.IntValue(val),
	})
	return e
}

func (e *Event) Int8(key string, val int8) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.IntValue(int(val)),
	})
	return e
}

func (e *Event) Int16(key string, val int16) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.IntValue(int(val)),
	})
	return e
}

func (e *Event) Int32(key string, val int32) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.IntValue(int(val)),
	})
	return e
}

func (e *Event) Int64(key string, val int64) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Int64Value(val),
	})
	return e
}

func (e *Event) Uint(key string, val uint) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return e
}

func (e *Event) Uint8(key string, val uint8) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return e

}

func (e *Event) Uint16(key string, val uint16) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return e
}

func (e *Event) Uint32(key string, val uint32) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return e
}

func (e *Event) Uint64(key string, val uint64) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(val),
	})
	return e
}

func (e *Event) Float32(key string, val float32) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Float64Value(float64(val)),
	})
	return e
}

func (e *Event) Float64(key string, val float64) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.Float64Value(val),
	})
	return e
}

func (e *Event) IPAddr(key string, ipAddr net.IP) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.StringValue(ipAddr.String()),
	})
	return e
}

func (e *Event) MACAddr(key string, macAddr net.HardwareAddr) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   key,
		Value: slog.StringValue(macAddr.String()),
	})
	return e
}

func (e *Event) Err(err error) logging.Event {
	e.attr = append(e.attr, slog.Attr{
		Key:   logging.ErrorKey,
		Value: slog.StringValue(err.Error()),
	})
	return e
}

func (e *Event) Msg(msg string) {
	e.logger.LogAttrs(defaultContext, e.level, msg, e.attr...)
}

func (e *Event) Msgf(format string, args ...interface{}) {
	e.logger.LogAttrs(defaultContext, e.level, fmt.Sprintf(format, args...), e.attr...)
}
