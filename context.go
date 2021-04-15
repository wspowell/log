package logging

import (
	"github.com/wspowell/local"
)

type logLevelContextKey struct{}
type configContextKey struct{}
type logContextKey struct{}

func WithContext(ctx local.Context, config Configer) {
	local.WithValue(ctx, configContextKey{}, config)
	local.WithValue(ctx, logLevelContextKey{}, config.Level())
}

func shouldLog(ctx local.Context, logLevel Level) bool {
	if configLevel, ok := ctx.Value(logLevelContextKey{}).(Level); ok {
		return logLevel >= configLevel
	}
	return false
}

func fromContext(ctx local.Context) Logger {
	if log, ok := ctx.Value(logContextKey{}).(Logger); ok {
		return log
	}

	if config, ok := ctx.Value(configContextKey{}).(Configer); ok {
		// Logger is localized to the context so it guarantees a data race cannot occur.
		log := config.Logger()
		ctx.Localize(logContextKey{}, log)
		return log
	}

	return nil
}

func Tag(ctx local.Context, name string, value interface{}) {
	if log := fromContext(ctx); log != nil {
		log.Tag(name, value)
	}
}

func Printf(ctx local.Context, format string, v ...interface{}) {
	// Log at INFO to match logrus.
	if shouldLog(ctx, LevelInfo) {
		if log := fromContext(ctx); log != nil {
			log.Printf(format, v)
		}
	}

}

func Trace(ctx local.Context, format string, v ...interface{}) {
	if shouldLog(ctx, LevelTrace) {
		if log := fromContext(ctx); log != nil {
			log.Trace(format, v)
		}
	}
}

func Debug(ctx local.Context, format string, v ...interface{}) {
	if shouldLog(ctx, LevelDebug) {
		if log := fromContext(ctx); log != nil {
			log.Debug(format, v)
		}
	}
}

func Info(ctx local.Context, format string, v ...interface{}) {
	if shouldLog(ctx, LevelInfo) {
		if log := fromContext(ctx); log != nil {
			log.Info(format, v)
		}
	}
}

func Warn(ctx local.Context, format string, v ...interface{}) {
	if shouldLog(ctx, LevelWarn) {
		if log := fromContext(ctx); log != nil {
			log.Warn(format, v)
		}
	}
}

func Error(ctx local.Context, format string, v ...interface{}) {
	if shouldLog(ctx, LevelError) {
		if log := fromContext(ctx); log != nil {
			log.Error(format, v)
		}
	}
}

func Fatal(ctx local.Context, format string, v ...interface{}) {
	if shouldLog(ctx, LevelFatal) {
		if log := fromContext(ctx); log != nil {
			log.Fatal(format, v)
		}
	}
}
