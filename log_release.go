//go:build release
// +build release

package log

func (self *Log) printf(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Debug().Msgf(format, values...)
}

// Perf: Each log function has its own function based on benchmark feedback. Not sure why though.
// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) trace(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Trace().Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) debug(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Debug().Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) info(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Info().Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) warn(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Warn().Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) error(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Error().Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) fatal(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Fatal().Msgf(format, values...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self *Log) panic(format string, values ...any) {
	logger := self.logger.Logger()
	logger.Panic().Msgf(format, values...)
}
