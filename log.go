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
	output     io.Writer
	globalTags map[string]any

	logger zerolog.Logger
	// Slice of key/value pairs in sequence.
	tags []any
	// Perf: Copying the level into the struct performed better.
	level Level
}

func NewLog(config LoggerConfig) *Log {
	return &Log{
		output:     config.Output(),
		globalTags: config.Tags(),
		logger:     log.Output(config.Output()).With().Fields(config.Tags()).Logger(),
		tags:       []any{},
		level:      config.Level(),
	}
}

func (self *Log) Tag(name string, value any) {
	self.tags = append(self.tags, name, value)
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
	newLog := &Log{
		output:     self.output,
		globalTags: self.globalTags,
		logger:     log.Output(self.output).With().Fields(self.globalTags).Logger(),
		tags:       make([]any, 0, len(self.tags)),
		level:      self.level,
	}

	for i := 0; i < len(self.tags); i += 2 {
		newLog.tags = append(newLog.tags, self.tags[i].(string), self.tags[i+1])
	}

	return newLog
}
