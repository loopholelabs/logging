/*
	Copyright 2024 Loophole Labs

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

var _ Logger = (*NoopLogger)(nil)

// NoopLogger is a no-op logger
type NoopLogger struct{}

// NewNoopLogger creates a new instance of the *NoopLogger
func NewNoopLogger() *NoopLogger {
	return new(NoopLogger)
}

func (n *NoopLogger) Fatal(string) {}

func (n *NoopLogger) Fatalf(string, ...interface{}) {}

func (n *NoopLogger) Error(string) {}

func (n *NoopLogger) Errorf(string, ...interface{}) {}

func (n *NoopLogger) Warn(string) {}

func (n *NoopLogger) Warnf(string, ...interface{}) {}

func (n *NoopLogger) Info(string) {}

func (n *NoopLogger) Infof(string, ...interface{}) {}

func (n *NoopLogger) Debug(string) {}

func (n *NoopLogger) Debugf(string, ...interface{}) {}
