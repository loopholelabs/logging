package noop

import (
	"net"

	"github.com/loopholelabs/logging/types"
)

var _ types.Context = (*Context)(nil)

type Context struct {
	l *Logger
}

func (c *Context) Logger() types.SubLogger {
	return c.l
}

func (c *Context) Str(key string, val string) types.Context {
	return c
}

func (c *Context) Bool(key string, val bool) types.Context {
	return c
}

func (c *Context) Int(key string, val int) types.Context {
	return c
}

func (c *Context) Int8(key string, val int8) types.Context {
	return c
}

func (c *Context) Int16(key string, val int16) types.Context {
	return c
}

func (c *Context) Int32(key string, val int32) types.Context {
	return c
}

func (c *Context) Int64(key string, val int64) types.Context {
	return c
}

func (c *Context) Uint(key string, val uint) types.Context {
	return c
}

func (c *Context) Uint8(key string, val uint8) types.Context {
	return c
}

func (c *Context) Uint16(key string, val uint16) types.Context {
	return c
}

func (c *Context) Uint32(key string, val uint32) types.Context {
	return c
}

func (c *Context) Uint64(key string, val uint64) types.Context {
	return c
}

func (c *Context) Float32(key string, val float32) types.Context {
	return c
}

func (c *Context) Float64(key string, val float64) types.Context {
	return c
}

func (c *Context) IPAddr(key string, ipAddr net.IP) types.Context {
	return c
}

func (c *Context) MACAddr(key string, macAddr net.HardwareAddr) types.Context {
	return c
}
