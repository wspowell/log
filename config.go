package log

import (
	"io"
	"os"
)

type Level int8

const (
	LevelTrace Level = iota - 1 // Match zerolog levels.
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

type Config struct {
	level Level

	// Tags are added to each Logger created.
	// Therefore, these tags are global and must not be altered.
	globalTags map[string]any

	output io.Writer
}

// NewConfig with defaults.
func NewConfig() *Config {
	return &Config{
		level:      LevelInfo,
		globalTags: map[string]any{},
		output:     os.Stdout,
	}
}

func (self *Config) WithLevel(level Level) *Config {
	self.level = level
	return self
}

func (self *Config) WithTags(tags map[string]any) *Config {
	self.globalTags = tags
	return self
}

func (self *Config) WithOutput(output io.Writer) *Config {
	self.output = output
	return self
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

func (self *Config) Output() io.Writer {
	return os.Stdout
}
