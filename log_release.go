//go:build release
// +build release

package log

func (self Log) printf(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Printf(format, v...)
}

// Perf: Each log function has its own function based on benchmark feedback. Not sure why though.
// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) trace(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Tracef(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) debug(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Debugf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) info(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Infof(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) warn(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Warnf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) error(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Errorf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) fatal(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Fatalf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) panic(format string, v ...any) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Panicf(format, v...)
}
