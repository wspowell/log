// +build release

package log

func (self Log) printf(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Printf(format, v...)
}

// Perf: Each log function has its own function based on benchmark feedback. Not sure why though.
func (self Log) trace(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Tracef(format, v...)
}

func (self Log) debug(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Debugf(format, v...)
}

func (self Log) info(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Infof(format, v...)
}

func (self Log) warn(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Warnf(format, v...)
}

func (self Log) error(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Errorf(format, v...)
}

func (self Log) fatal(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Fatalf(format, v...)
}

func (self Log) panic(format string, v ...interface{}) {
	self.logger.WithFields(self.tags).WithFields(self.cfg.Tags()).Panicf(format, v...)
}
