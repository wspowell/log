//go:build !release
// +build !release

package log

func (self Log) printf(format string, v ...any) {
	self.logger.Debug().Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// Perf: Each log function has its own function based on benchmark feedback. Not sure why though.
// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) trace(format string, v ...any) {
	self.logger.Trace().Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) debug(format string, v ...any) {
	self.logger.Debug().Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) info(format string, v ...any) {
	self.logger.Info().Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) warn(format string, v ...any) {
	self.logger.Warn().Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) error(format string, v ...any) {
	self.logger.Error().Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) fatal(format string, v ...any) {
	self.logger.Fatal().Str("function", getCallerFunctionName()).Msgf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) panic(format string, v ...any) {
	self.logger.Panic().Str("function", getCallerFunctionName()).Msgf(format, v...)
}
