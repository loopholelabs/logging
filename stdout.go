/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package logging

import (
	"github.com/gookit/color"
	"io"
)

// ConsoleLogLevel is the log level for the ConsoleLogger
type ConsoleLogLevel int

const (
	FatalLevel = iota
	ErrorLevel
	WarnLevel
	ImportantLevel
	SuccessLevel
	InfoLevel
	DebugLevel
)

var _ Logger = (*ConsoleLogger)(nil)

// ConsoleLogger logs to the console and satisfies the Logger interface
type ConsoleLogger struct {
	level ConsoleLogLevel
}

// NewConsoleLogger creates a new instance of the *ConsoleLogger with a default logging level of FatalLevel
func NewConsoleLogger() *ConsoleLogger {
	return new(ConsoleLogger)
}

// SetLevel sets the logging level of the ConsoleLogger to level
func (c *ConsoleLogger) SetLevel(level ConsoleLogLevel) {
	c.level = level
}

// SetOutput sets the output of the ConsoleLogger to an io.Writer
func (c *ConsoleLogger) SetOutput(w io.Writer) {
	color.SetOutput(w)
}

// Debug logs a Debug level message in Cyan
func (c *ConsoleLogger) Debug(msg string) {
	if c.level <= DebugLevel {
		color.Debug.Print(msg)
	}
}

// Debugf logs a Debug level message in Cyan
func (c *ConsoleLogger) Debugf(msg string, args ...interface{}) {
	if c.level <= DebugLevel {
		color.Debug.Printf(msg, args...)
	}
}

// Info logs an Info level message in Gray
func (c *ConsoleLogger) Info(msg string) {
	if c.level <= InfoLevel {
		color.Notice.Print(msg)
	}
}

// Infof logs an Info level message in Gray
func (c *ConsoleLogger) Infof(msg string, args ...interface{}) {
	if c.level <= InfoLevel {
		color.Notice.Printf(msg, args...)
	}
}

// Success logs a Success level message in Green
func (c *ConsoleLogger) Success(msg string) {
	if c.level <= SuccessLevel {
		color.Info.Print(msg)
	}
}

// Successf logs a Success level message in Green
func (c *ConsoleLogger) Successf(msg string, args ...interface{}) {
	if c.level <= SuccessLevel {
		color.Info.Printf(msg, args...)
	}
}

// Important logs an Important level message in Blue
func (c *ConsoleLogger) Important(msg string) {
	if c.level <= ImportantLevel {
		color.Primary.Print(msg)
	}
}

// Importantf logs an Important level message in Blue
func (c *ConsoleLogger) Importantf(msg string, args ...interface{}) {
	if c.level <= ImportantLevel {
		color.Primary.Printf(msg, args)
	}
}

// Error logs an Error level message in Red
func (c *ConsoleLogger) Error(msg string) {
	if c.level <= ErrorLevel {
		color.Error.Printf(msg)
	}
}

// Errorf logs an Error level message in Red
func (c *ConsoleLogger) Errorf(msg string, args ...interface{}) {
	if c.level <= ErrorLevel {
		color.Error.Printf(msg, args...)
	}
}

// Warn logs a Warn level message Orange
func (c *ConsoleLogger) Warn(msg string) {
	if c.level <= WarnLevel {
		color.Danger.Printf(msg)
	}
}

// Warnf logs a Warn level message Orange
func (c *ConsoleLogger) Warnf(msg string, args ...interface{}) {
	if c.level <= WarnLevel {
		color.Danger.Printf(msg, args...)
	}
}

// Fatal logs a Fatal level message in White
func (c *ConsoleLogger) Fatal(msg string) {
	if c.level <= FatalLevel {
		color.Light.Printf(msg)
	}
}

// Fatalf logs a Fatal level message in White
func (c *ConsoleLogger) Fatalf(msg string, args ...interface{}) {
	if c.level <= FatalLevel {
		color.Light.Printf(msg, args...)
	}
}
