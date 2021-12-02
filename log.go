package log

import "github.com/sirupsen/logrus"

type Logger interface {
	Tag(name string, value interface{})
	Tags() map[string]interface{}
	Printf(format string, v ...interface{})
	Trace(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

type baseLog struct {
	logger *logrus.Logger

	cfg  Configer
	tags logrus.Fields
}

var _ Logger = (*Log)(nil)

// Perf: Log value receivers performed better.
type Log struct {
	// Perf: Store least used values in one struct because copying one reference is faster than copying many values.
	*baseLog

	// Perf: Copying the level into the struct performed better.
	level Level
}

func NewLog(cfg Configer) Log {
	if logConfig, ok := cfg.(*Config); ok {
		return Log{
			baseLog: &baseLog{
				logger: logConfig.logger,
				cfg:    cfg,
				tags:   logrus.Fields{},
			},
			level: logConfig.level,
		}
	}

	return Log{
		baseLog: &baseLog{
			logger: newLogrusLogger(cfg),
			cfg:    cfg,
			tags:   logrus.Fields{},
		},
		level: cfg.Level(),
	}
}

func (self Log) Tag(name string, value interface{}) {
	self.tags[name] = value
}

// Tags cloned value.
// Returns tags for the Logger only, not the Configer.
func (self Log) Tags() map[string]interface{} {
	clone := make(map[string]interface{}, len(self.tags))
	for key, value := range self.tags {
		clone[key] = value
	}

	return clone
}

func (self Log) Printf(format string, v ...interface{}) {
	// Log at INFO to match logrus.
	if LevelInfo >= self.level {
		self.printf(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Trace(format string, v ...interface{}) {
	if LevelTrace >= self.level {
		self.trace(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Debug(format string, v ...interface{}) {
	if LevelDebug >= self.level {
		self.debug(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Info(format string, v ...interface{}) {
	if LevelInfo >= self.level {
		self.info(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Warn(format string, v ...interface{}) {
	if LevelWarn >= self.level {
		self.warn(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Error(format string, v ...interface{}) {
	if LevelError >= self.level {
		self.error(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Fatal(format string, v ...interface{}) {
	if LevelFatal >= self.level {
		self.fatal(format, v...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) Panic(format string, v ...interface{}) {
	if LevelPanic >= self.level {
		self.panic(format, v...)
	}
}

// Localize Log to the next Context. Called in context.Localize().
// This will copy all log tags on to the localized Log value.
// If a fresh Logger is desired, use WithContext() to override.
func (self Log) Localize() interface{} {
	clone := self.baseLog.cfg.Logger()
	for key, value := range self.tags {
		clone.Tag(key, value)
	}

	return clone
}
