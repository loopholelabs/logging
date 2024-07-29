package zerolog

import (
	"net"

	"github.com/loopholelabs/logging/types"
	"github.com/rs/zerolog"
)

var _ types.Context = (*Context)(nil)

type Context struct {
	l       *Logger
	zeroCtx zerolog.Context
}

func (c *Context) Logger() types.SubLogger {
	return &Logger{
		logger: c.zeroCtx.Logger(),
		source: c.l.source,
		level:  c.l.level,
	}
}

func (c *Context) Str(key string, val string) types.Context {
	c.zeroCtx = c.zeroCtx.Str(key, val)
	return c
}

func (c *Context) Bool(key string, val bool) types.Context {
	c.zeroCtx = c.zeroCtx.Bool(key, val)
	return c
}

func (c *Context) Int(key string, val int) types.Context {
	c.zeroCtx = c.zeroCtx.Int(key, val)
	return c
}

func (c *Context) Int8(key string, val int8) types.Context {
	c.zeroCtx = c.zeroCtx.Int8(key, val)
	return c
}

func (c *Context) Int16(key string, val int16) types.Context {
	c.zeroCtx = c.zeroCtx.Int16(key, val)
	return c
}

func (c *Context) Int32(key string, val int32) types.Context {
	c.zeroCtx = c.zeroCtx.Int32(key, val)
	return c
}

func (c *Context) Int64(key string, val int64) types.Context {
	c.zeroCtx = c.zeroCtx.Int64(key, val)
	return c
}

func (c *Context) Uint(key string, val uint) types.Context {
	c.zeroCtx = c.zeroCtx.Uint(key, val)
	return c
}

func (c *Context) Uint8(key string, val uint8) types.Context {
	c.zeroCtx = c.zeroCtx.Uint8(key, val)
	return c
}

func (c *Context) Uint16(key string, val uint16) types.Context {
	c.zeroCtx = c.zeroCtx.Uint16(key, val)
	return c
}

func (c *Context) Uint32(key string, val uint32) types.Context {
	c.zeroCtx = c.zeroCtx.Uint32(key, val)
	return c
}

func (c *Context) Uint64(key string, val uint64) types.Context {
	c.zeroCtx = c.zeroCtx.Uint64(key, val)
	return c
}

func (c *Context) Float32(key string, val float32) types.Context {
	c.zeroCtx = c.zeroCtx.Float32(key, val)
	return c
}

func (c *Context) Float64(key string, val float64) types.Context {
	c.zeroCtx = c.zeroCtx.Float64(key, val)
	return c
}

func (c *Context) IPAddr(key string, val net.IP) types.Context {
	c.zeroCtx = c.zeroCtx.IPAddr(key, val)
	return c
}

func (c *Context) MACAddr(key string, val net.HardwareAddr) types.Context {
	c.zeroCtx = c.zeroCtx.MACAddr(key, val)
	return c
}
