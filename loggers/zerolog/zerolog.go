// SPDX-License-Identifier: Apache-2.0

package zerolog

import (
	"fmt"
	"io"

	"github.com/rs/zerolog"

	"github.com/loopholelabs/logging"
)

var _ logging.Logger = (*Logger)(nil)

type Logger struct {
	logger zerolog.Logger
	source string
}

func init() {
	zerolog.TimestampFieldName = logging.TimestampKey
	zerolog.ErrorFieldName = logging.ErrorKey
}

func New(source string, level logging.Level, output io.Writer) *Logger {
	zl := zerolog.New(output)
	z := &Logger{
		logger: zl,
		source: source,
	}
	z.SetLevel(level)
	return z
}

func (z *Logger) SetLevel(level logging.Level) {
	var zerologLevel zerolog.Level
	switch level {
	case logging.FatalLevel:
		zerologLevel = zerolog.FatalLevel
	case logging.ErrorLevel:
		zerologLevel = zerolog.ErrorLevel
	case logging.WarnLevel:
		zerologLevel = zerolog.WarnLevel
	case logging.InfoLevel:
		zerologLevel = zerolog.InfoLevel
	case logging.DebugLevel:
		zerologLevel = zerolog.DebugLevel
	case logging.TraceLevel:
		zerologLevel = zerolog.TraceLevel
	}
	z.logger.Level(zerologLevel)
}

func (z *Logger) SubLogger(source string) logging.Logger {
	return &Logger{
		logger: z.logger,
		source: fmt.Sprintf("%s:%s", z.source, source),
	}
}

func (z *Logger) Fatal() logging.Event {
	return (*Event)(z.logger.Fatal().Timestamp().Str(logging.SourceKey, z.source))
}

func (z *Logger) Error() logging.Event {
	return (*Event)(z.logger.Error().Timestamp().Str(logging.SourceKey, z.source))
}

func (z *Logger) Warn() logging.Event {
	return (*Event)(z.logger.Warn().Timestamp().Str(logging.SourceKey, z.source))
}

func (z *Logger) Info() logging.Event {
	return (*Event)(z.logger.Info().Timestamp().Str(logging.SourceKey, z.source))
}

func (z *Logger) Debug() logging.Event {
	return (*Event)(z.logger.Debug().Timestamp().Str(logging.SourceKey, z.source))
}

func (z *Logger) Trace() logging.Event {
	return (*Event)(z.logger.Trace().Timestamp().Str(logging.SourceKey, z.source))
}
