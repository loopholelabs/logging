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

import "testing"

var _ Logger = (*TestLogger)(nil)

// TestLogger is a logger for use in tests
type TestLogger struct {
	t testing.TB
}

// NewTestLogger creates a new instance of the *TestLogger
func NewTestLogger(t testing.TB) *TestLogger {
	return &TestLogger{t: t}
}

func (t *TestLogger) Fatal(msg string) {
	t.t.Log("[FATAL] " + msg)
}

func (t *TestLogger) Fatalf(msg string, args ...interface{}) {
	t.t.Logf("[FATAL] "+msg, args...)
}

func (t *TestLogger) Error(msg string) {
	t.t.Log("[ERROR] " + msg)
}

func (t *TestLogger) Errorf(msg string, args ...interface{}) {
	t.t.Logf("[ERROR] "+msg, args...)
}

func (t *TestLogger) Warn(msg string) {
	t.t.Log("[WARN] " + msg)
}

func (t *TestLogger) Warnf(msg string, args ...interface{}) {
	t.t.Logf("[WARN] "+msg, args...)
}

func (t *TestLogger) Info(msg string) {
	t.t.Log("[INFO] " + msg)
}

func (t *TestLogger) Infof(msg string, args ...interface{}) {
	t.t.Logf("[INFO] "+msg, args...)
}

func (t *TestLogger) Debug(msg string) {
	t.t.Log("[DEBUG] " + msg)
}

func (t *TestLogger) Debugf(msg string, args ...interface{}) {
	t.t.Logf("[DEBUG] "+msg, args...)
}
