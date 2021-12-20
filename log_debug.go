//go:build !release
// +build !release

package log

func (self Log) printf(format string, v ...any) {
	self.logger.Debug().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// Perf: Each log function has its own function based on benchmark feedback. Not sure why though.
// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) trace(format string, v ...any) {
	self.logger.Trace().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) debug(format string, v ...any) {
	self.logger.Debug().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) info(format string, v ...any) {
	self.logger.Info().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) warn(format string, v ...any) {
	self.logger.Warn().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) error(format string, v ...any) {
	self.logger.Error().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) fatal(format string, v ...any) {
	self.logger.Fatal().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) panic(format string, v ...any) {
	self.logger.Panic().Fields(self.tags).Fields(self.cfg.Tags()).Str("function", getCallerFunctionName()).Msgf(format, v...)
}
