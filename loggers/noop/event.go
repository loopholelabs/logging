// SPDX-License-Identifier: Apache-2.0

package noop

import (
	"net"

	"github.com/loopholelabs/logging/types"
)

var _ types.Event = (*Event)(nil)

type Event struct{}

func (e *Event) Str(string, string) types.Event { return e }

func (e *Event) Bool(string, bool) types.Event { return e }

func (e *Event) Int(string, int) types.Event { return e }

func (e *Event) Int8(string, int8) types.Event { return e }

func (e *Event) Int16(string, int16) types.Event { return e }

func (e *Event) Int32(string, int32) types.Event { return e }

func (e *Event) Int64(string, int64) types.Event { return e }

func (e *Event) Uint(string, uint) types.Event { return e }

func (e *Event) Uint8(string, uint8) types.Event { return e }

func (e *Event) Uint16(string, uint16) types.Event { return e }

func (e *Event) Uint32(string, uint32) types.Event { return e }

func (e *Event) Uint64(string, uint64) types.Event { return e }

func (e *Event) Float32(string, float32) types.Event { return e }

func (e *Event) Float64(string, float64) types.Event { return e }

func (e *Event) IPAddr(string, net.IP) types.Event { return e }

func (e *Event) MACAddr(string, net.HardwareAddr) types.Event { return e }

func (e *Event) Err(error) types.Event { return e }

func (e *Event) Msg(string) {}

func (e *Event) Msgf(string, ...interface{}) {}
