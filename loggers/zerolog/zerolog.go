// SPDX-License-Identifier: Apache-2.0

package zerolog

import (
	"fmt"
	"io"

	"github.com/rs/zerolog"

	"github.com/loopholelabs/logging/types"
)

var _ types.Logger = (*Logger)(nil)

type Logger struct {
	logger zerolog.Logger
	source string
}

func init() {
	zerolog.TimestampFieldName = types.TimestampKey
	zerolog.ErrorFieldName = types.ErrorKey
}

func New(source string, level types.Level, output io.Writer) *Logger {
	zl := zerolog.New(output)
	z := &Logger{
		logger: zl,
		source: source,
	}
	z.SetLevel(level)
	return z
}

func (z *Logger) SetLevel(level types.Level) {
	var zerologLevel zerolog.Level
	switch level {
	case types.FatalLevel:
		zerologLevel = zerolog.FatalLevel
	case types.ErrorLevel:
		zerologLevel = zerolog.ErrorLevel
	case types.WarnLevel:
		zerologLevel = zerolog.WarnLevel
	case types.InfoLevel:
		zerologLevel = zerolog.InfoLevel
	case types.DebugLevel:
		zerologLevel = zerolog.DebugLevel
	case types.TraceLevel:
		zerologLevel = zerolog.TraceLevel
	}
	z.logger.Level(zerologLevel)
}

func (z *Logger) SubLogger(source string) types.Logger {
	return &Logger{
		logger: z.logger,
		source: fmt.Sprintf("%s:%s", z.source, source),
	}
}

func (z *Logger) Fatal() types.Event {
	return (*Event)(z.logger.Fatal().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Error() types.Event {
	return (*Event)(z.logger.Error().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Warn() types.Event {
	return (*Event)(z.logger.Warn().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Info() types.Event {
	return (*Event)(z.logger.Info().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Debug() types.Event {
	return (*Event)(z.logger.Debug().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Trace() types.Event {
	return (*Event)(z.logger.Trace().Timestamp().Str(types.SourceKey, z.source))
}
