package slog

import (
	"log/slog"
	"net"

	"github.com/loopholelabs/logging/types"
)

var _ types.Context = (*Context)(nil)

type Context struct {
	l     *Logger
	attrs []any
}

func (c *Context) Logger() types.SubLogger {
	l := New(c.l.source, c.l.level, c.l.output)
	l.logger = c.l.logger.With(c.attrs...)
	return l
}

func (c *Context) Str(key string, val string) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.StringValue(val),
	})
	return c
}

func (c *Context) Bool(key string, val bool) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.BoolValue(val),
	})
	return c
}

func (c *Context) Int(key string, val int) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.IntValue(val),
	})
	return c
}

func (c *Context) Int8(key string, val int8) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.IntValue(int(val)),
	})
	return c
}

func (c *Context) Int16(key string, val int16) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.IntValue(int(val)),
	})
	return c
}

func (c *Context) Int32(key string, val int32) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.IntValue(int(val)),
	})
	return c
}

func (c *Context) Int64(key string, val int64) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Int64Value(val),
	})
	return c
}

func (c *Context) Uint(key string, val uint) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return c
}

func (c *Context) Uint8(key string, val uint8) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return c
}

func (c *Context) Uint16(key string, val uint16) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return c
}

func (c *Context) Uint32(key string, val uint32) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(uint64(val)),
	})
	return c
}

func (c *Context) Uint64(key string, val uint64) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Uint64Value(val),
	})
	return c
}

func (c *Context) Float32(key string, val float32) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Float64Value(float64(val)),
	})
	return c
}

func (c *Context) Float64(key string, val float64) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.Float64Value(float64(val)),
	})
	return c
}

func (c *Context) IPAddr(key string, ipAddr net.IP) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.StringValue(ipAddr.String()),
	})
	return c
}

func (c *Context) MACAddr(key string, macAddr net.HardwareAddr) types.Context {
	c.attrs = append(c.attrs, slog.Attr{
		Key:   key,
		Value: slog.StringValue(macAddr.String()),
	})
	return c
}
