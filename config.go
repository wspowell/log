package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Level int8

const (
	LevelTrace Level = iota - 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (self Level) String() string {
	return [...]string{
		"trace",
		"debug",
		"info",
		"warn",
		"error",
		"fatal",
		"panic",
	}[self]
}

type Configuration interface {
	Level() Level
	// Tags are added to each Logger created.
	// Therefore, these tags are global and must not be altered.
	Tags() map[string]any
	Out() io.Writer
	Logger() Logger
}

type Config struct {
	logger     zerolog.Logger
	level      Level
	globalTags map[string]any
}

func NewConfig(level Level) *Config {
	cfg := &Config{
		level:      level,
		globalTags: map[string]any{},
	}

	cfg.logger = log.Output(os.Stdout).Level(zerolog.DebugLevel)

	return cfg
}

func (self *Config) Level() Level {
	return self.level
}

func (self *Config) Tags() map[string]any {
	clone := make(map[string]any, len(self.globalTags))
	for key, value := range self.globalTags {
		clone[key] = value
	}

	return clone
}

func (self *Config) Out() io.Writer {
	return os.Stdout
}

func (self *Config) Logger() Logger {
	return NewLog(self)
}

func newZerologLogger(cfg Configuration) zerolog.Logger {
	return log.With().Fields(cfg.Tags()).Logger().Output(cfg.Out()).Level(zerolog.Level(cfg.Level()))
}
