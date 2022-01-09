//go:build release
// +build release

package log

func (self *Log) printf(format string, values ...any) {
	self.logger.Debug().Fields(self.tags).Msgf(format, values...)
}

// Perf: Each log function has its own function based on benchmark feedback. Not sure why though.
// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) trace(format string, values ...any) {
	self.logger.Trace().Fields(self.tags).Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) debug(format string, values ...any) {
	self.logger.Debug().Fields(self.tags).Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) info(format string, values ...any) {
	self.logger.Info().Fields(self.tags).Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) warn(format string, values ...any) {
	self.logger.Warn().Fields(self.tags).Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) error(format string, values ...any) {
	self.logger.Error().Fields(self.tags).Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) fatal(format string, values ...any) {
	self.logger.Fatal().Fields(self.tags).Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) panic(format string, values ...any) {
	self.logger.Panic().Fields(self.tags).Msgf(format, values...)
}
