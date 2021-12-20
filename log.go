package log

import (
	"github.com/rs/zerolog"
)

type Logger interface {
	Tag(name string, value any)
	Tags() map[string]any
	Printf(format string, v ...any)
	Trace(format string, v ...any)
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	Fatal(format string, v ...any)
}

type baseLog struct {
	logger zerolog.Logger
	cfg    Configuration
	tags   map[string]any
}

var _ Logger = (*Log)(nil)

// Perf: Log value receivers performed better.
type Log struct {
	// Perf: Store least used values in one struct because copying one reference is faster than copying many values.
	*baseLog

	// Perf: Copying the level into the struct performed better.
	level Level
}

func NewLog(config Configuration) Log {
	if logConfig, ok := config.(*Config); ok {
		return Log{
			baseLog: &baseLog{
				logger: logConfig.logger,
				cfg:    config,
				tags:   map[string]any{},
			},
			level: logConfig.level,
		}
	}

	return Log{
		baseLog: &baseLog{
			logger: newZerologLogger(config),
			cfg:    config,
			tags:   map[string]any{},
		},
		level: config.Level(),
	}
}

func (self Log) Tag(name string, value any) {
	self.tags[name] = value
}

// Tags cloned value.
// Returns tags for the Logger only, not the Configer.
func (self Log) Tags() map[string]any {
	clone := make(map[string]any, len(self.tags))
	for key, value := range self.tags {
		clone[key] = value
	}

	return clone
}

func (self Log) Printf(format string, v ...any) {
	// Log at INFO to match logrus.
	if LevelInfo >= self.level {
		self.printf(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Trace(format string, v ...any) {
	if LevelTrace >= self.level {
		self.trace(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Debug(format string, v ...any) {
	if LevelDebug >= self.level {
		self.debug(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Info(format string, v ...any) {
	if LevelInfo >= self.level {
		self.info(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Warn(format string, v ...any) {
	if LevelWarn >= self.level {
		self.warn(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Error(format string, v ...any) {
	if LevelError >= self.level {
		self.error(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Fatal(format string, v ...any) {
	if LevelFatal >= self.level {
		self.fatal(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Panic(format string, v ...any) {
	if LevelPanic >= self.level {
		self.panic(format, v...)
	}
}

// Localize Log to the next Context. Called in context.Localize().
// This will copy all log tags on to the localized Log value.
// If a fresh Logger is desired, use WithContext() to override.
func (self Log) Localize() any {
	clone := self.baseLog.cfg.Logger()
	for key, value := range self.tags {
		clone.Tag(key, value)
	}

	return clone
}
