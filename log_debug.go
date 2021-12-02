//go:build !release
// +build !release

package log

import "github.com/sirupsen/logrus"

func (self Log) printf(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Printf(format, v...)
}

// Perf: Each log function has its own function based on benchmark feedback. Not sure why though.
// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) trace(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Tracef(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) debug(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Debugf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) info(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Infof(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) warn(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Warnf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) error(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Errorf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) fatal(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Fatalf(format, v...)
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func (self Log) panic(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).WithFields(logrus.Fields{"function": getCallerFunctionName()}).Panicf(format, v...)
}
