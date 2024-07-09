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

var _ Logger = (*NopLogger)(nil)

// Noop is a no-op logger
type Noop struct {}

func (n *Noop) Fatal(string) {}

func (n *Noop) Fatalf(string, ...interface{}) {}

func (n *Noop) Error(string) {}

func (n *Noop) Errorf(string, ...interface{}) {}

func (n *Noop) Warn(string) {}

func (n *Noop) Warnf(string, ...interface{}) {}

func (n *Noop) Info(msg string) {}

func (n *Noop) Infof(string, ...interface{}) {}

func (n *Noop) Debug(string) {}

func (n *Noop) Debugf(string, ...interface{}) {}
