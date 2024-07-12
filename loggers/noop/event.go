// SPDX-License-Identifier: Apache-2.0

package noop

import (
	"net"

	"github.com/loopholelabs/logging"
)

var _ logging.Event = (*Event)(nil)

type Event struct{}

func (e *Event) Str(string, string) logging.Event { return e }

func (e *Event) Bool(string, bool) logging.Event { return e }

func (e *Event) Int(string, int) logging.Event { return e }

func (e *Event) Int8(string, int8) logging.Event { return e }

func (e *Event) Int16(string, int16) logging.Event { return e }

func (e *Event) Int32(string, int32) logging.Event { return e }

func (e *Event) Int64(string, int64) logging.Event { return e }

func (e *Event) Uint(string, uint) logging.Event { return e }

func (e *Event) Uint8(string, uint8) logging.Event { return e }

func (e *Event) Uint16(string, uint16) logging.Event { return e }

func (e *Event) Uint32(string, uint32) logging.Event { return e }

func (e *Event) Uint64(string, uint64) logging.Event { return e }

func (e *Event) Float32(string, float32) logging.Event { return e }

func (e *Event) Float64(string, float64) logging.Event { return e }

func (e *Event) IPAddr(string, net.IP) logging.Event { return e }

func (e *Event) MACAddr(string, net.HardwareAddr) logging.Event { return e }

func (e *Event) Err(error) logging.Event { return e }

func (e *Event) Msg(string) {}

func (e *Event) Msgf(string, ...interface{}) {}
