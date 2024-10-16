// SPDX-License-Identifier: Apache-2.0

package zerolog

import (
	"fmt"
	"io"
	"sync"

	"github.com/rs/zerolog"

	"github.com/loopholelabs/logging/types"
)

var _ types.RootLogger = (*Logger)(nil)

type Logger struct {
	level      types.Level
	logger     *zerolog.Logger
	loggerLock sync.RWMutex
	source     string
}

func init() {
	zerolog.TimestampFieldName = types.TimestampKey
	zerolog.ErrorFieldName = types.ErrorKey
}

func New(source string, level types.Level, output io.Writer) *Logger {
	return newLogger(source, level, zerolog.New(output))
}

func newLogger(source string, level types.Level, zl zerolog.Logger) *Logger {
	z := &Logger{
		logger: &zl,
		source: source,
	}
	z.SetLevel(level)
	return z
}

func (z *Logger) Level() types.Level {
	return z.level
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

	newLogger := z.logger.Level(zerologLevel)

	z.loggerLock.Lock()
	defer z.loggerLock.Unlock()

	z.logger = &newLogger
	z.level = level
}

func (z *Logger) SubLogger(source string) types.Logger {
	return newLogger(fmt.Sprintf("%s:%s", z.source, source), z.level, *z.logger)
}

func (z *Logger) With() types.Context {
	z.loggerLock.RLock()
	defer z.loggerLock.RUnlock()

	return &Context{
		l:       z,
		zeroCtx: z.logger.With(),
	}
}

func (z *Logger) Fatal() types.Event {
	z.loggerLock.RLock()
	defer z.loggerLock.RUnlock()

	return (*Event)(z.logger.Fatal().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Error() types.Event {
	z.loggerLock.RLock()
	defer z.loggerLock.RUnlock()

	return (*Event)(z.logger.Error().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Warn() types.Event {
	z.loggerLock.RLock()
	defer z.loggerLock.RUnlock()

	return (*Event)(z.logger.Warn().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Info() types.Event {
	z.loggerLock.RLock()
	defer z.loggerLock.RUnlock()

	return (*Event)(z.logger.Info().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Debug() types.Event {
	z.loggerLock.RLock()
	defer z.loggerLock.RUnlock()

	return (*Event)(z.logger.Debug().Timestamp().Str(types.SourceKey, z.source))
}

func (z *Logger) Trace() types.Event {
	z.loggerLock.RLock()
	defer z.loggerLock.RUnlock()

	return (*Event)(z.logger.Trace().Timestamp().Str(types.SourceKey, z.source))
}
