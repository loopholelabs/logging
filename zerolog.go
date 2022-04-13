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

import "github.com/rs/zerolog"

var _ Logger = (*ZerologAdapter)(nil)

// ZerologAdapter is an adapter to convert a zerolog.Logger into a struct that satisfies the Logger interface
type ZerologAdapter struct {
	l *zerolog.Logger
}

// ConvertZerolog converts a *zerolog.Logger to a *ZerologAdapter
func ConvertZerolog(l *zerolog.Logger) *ZerologAdapter {
	return &ZerologAdapter{l: l}
}

func (z *ZerologAdapter) Fatal(msg string) {
	z.l.Fatal().Msg(msg)
}

func (z *ZerologAdapter) Fatalf(msg string, args ...interface{}) {
	z.l.Fatal().Msgf(msg, args...)
}

func (z *ZerologAdapter) Error(msg string) {
	z.l.Error().Msg(msg)
}

func (z *ZerologAdapter) Errorf(msg string, args ...interface{}) {
	z.l.Error().Msgf(msg, args...)
}

func (z *ZerologAdapter) Warn(msg string) {
	z.l.Warn().Msg(msg)
}

func (z *ZerologAdapter) Warnf(msg string, args ...interface{}) {
	z.l.Warn().Msgf(msg, args...)
}

func (z *ZerologAdapter) Info(msg string) {
	z.l.Info().Msg(msg)
}

func (z *ZerologAdapter) Infof(msg string, args ...interface{}) {
	z.l.Info().Msgf(msg, args...)
}

func (z *ZerologAdapter) Debug(msg string) {
	z.l.Debug().Msg(msg)
}

func (z *ZerologAdapter) Debugf(msg string, args ...interface{}) {
	z.l.Debug().Msgf(msg, args...)
}
