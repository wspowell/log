package log

import (
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// DefaultLogger returns a new default logger that logs at info level.
func DefaultLogger() *Log {
	return NewLog(NewConfig())
}

// DebugLogger returns a new default logger that logs at debug level.
func DebugLogger() *Log {
	return NewLog(NewConfig().WithLevel(LevelDebug))
}

// ErrorLogger returns a new default logger that logs at error level.
func ErrorLogger() *Log {
	return NewLog(NewConfig().WithLevel(LevelDebug))
}

type Logger interface {
	Tag(name string, value any)
	Tags() map[string]any
	Printf(format string, values ...any)
	Trace(format string, values ...any)
	Debug(format string, values ...any)
	Info(format string, values ...any)
	Warn(format string, values ...any)
	Error(format string, values ...any)
	Fatal(format string, values ...any)
}

type LoggerConfig interface {
	Level() Level
	Tags() map[string]any
	Output() io.Writer
}

var _ Logger = (*Log)(nil)

type Log struct {
	logger zerolog.Logger
	tags   map[string]any
	// Perf: Copying the level into the struct performed better.
	level Level
}

func NewLog(config LoggerConfig) *Log {
	return &Log{
		logger: log.Output(config.Output()).Level(zerolog.Level(config.Level())).With().Fields(config.Tags()).Logger(),
		tags:   config.Tags(),
		level:  config.Level(),
	}
}

func (self *Log) Tag(name string, value any) {
	self.tags[name] = value
	self.logger = self.logger.With().Fields(map[string]any{
		name: value,
	}).Logger()
}

// Tags cloned value.
// Returns tags for the Logger only, not the Configer.
func (self *Log) Tags() map[string]any {
	clone := make(map[string]any, len(self.tags))
	for key, value := range self.tags {
		clone[key] = value
	}

	return clone
}

func (self *Log) Printf(format string, values ...any) {
	// Log at INFO to match logrus.
	if LevelInfo >= self.level {
		self.printf(format, values...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) Trace(format string, values ...any) {
	if LevelTrace >= self.level {
		self.trace(format, values...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) Debug(format string, values ...any) {
	if LevelDebug >= self.level {
		self.debug(format, values...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) Info(format string, values ...any) {
	if LevelInfo >= self.level {
		self.info(format, values...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) Warn(format string, values ...any) {
	if LevelWarn >= self.level {
		self.warn(format, values...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) Error(format string, values ...any) {
	if LevelError >= self.level {
		self.error(format, values...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) Fatal(format string, values ...any) {
	if LevelFatal >= self.level {
		self.fatal(format, values...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) Panic(format string, values ...any) {
	if LevelPanic >= self.level {
		self.panic(format, values...)
	}
}

// Localize Log to the next Context. Called in context.Localize().
// This will copy all log tags on to the localized Log value.
// If a fresh Logger is desired, use WithContext() to override.
func (self *Log) Localize() any {
	return &Log{
		logger: self.logger,
		tags:   self.Tags(),
		level:  self.level,
	}
}
