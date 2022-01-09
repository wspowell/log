package log

import (
	"github.com/wspowell/context"
)

type configContextKey struct{}
type levelContextKey struct{}
type logContextKey struct{}

// WithConfig adds a LoggerConfig to the Context which enables new Loggers to be created.
// This creates a new Logger from the LoggerConfig and adds it as a local Context value.
func WithContext(ctx context.Context, config LoggerConfig) context.Context {
	ctx = context.WithValue(ctx, configContextKey{}, config)
	ctx = context.WithValue(ctx, levelContextKey{}, config.Level())
	log := NewLog(config)
	withLogger(ctx, log)

	return ctx
}

func withLogger(ctx context.Context, logger Logger) {
	// Logger is localized to the context so it guarantees a data race cannot occur.
	context.WithLocalValue(ctx, logContextKey{}, logger)
}

func fromContext(ctx context.Context, logLevel Level) (Logger, bool) {
	if configLevel, ok := ctx.Value(levelContextKey{}).(Level); ok {
		if logLevel < configLevel {
			return nil, false
		}

		return getLogger(ctx)
	}

	return nil, false
}

func getLogger(ctx context.Context) (Logger, bool) {
	if log, ok := ctx.Value(logContextKey{}).(Logger); ok {
		return log, true
	}

	return newLogger(ctx)
}

func newLogger(ctx context.Context) (Logger, bool) {
	// Create new Logger from the LoggerConfig, if present.
	if config, ok := ctx.Value(configContextKey{}).(LoggerConfig); ok {
		log := NewLog(config)
		withLogger(ctx, log)

		return log, true
	}

	return nil, false
}

func Tag(ctx context.Context, name string, value any) {
	if log, ok := fromContext(ctx, LevelPanic); ok {
		log.Tag(name, value)
	}
}

func Printf(ctx context.Context, format string, values ...any) {
	// Log at INFO to match logrus.
	if log, ok := fromContext(ctx, LevelInfo); ok {
		// perf: copy the values so that the compiler does not allocate on the heap. This prevents an allocation happening even
		//       when the log level is too low and the log function is never called.
		//       See: https://stackoverflow.com/questions/27788813/variadic-functions-causing-unnecessary-heap-allocations-in-go
		valuesCopy := make([]any, len(values))
		copy(valuesCopy, values)

		log.Printf(format, valuesCopy...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func Trace(ctx context.Context, format string, values ...any) {
	if log, ok := fromContext(ctx, LevelTrace); ok {
		// perf: copy the values so that the compiler does not allocate on the heap. This prevents an allocation happening even
		//       when the log level is too low and the log function is never called.
		//       See: https://stackoverflow.com/questions/27788813/variadic-functions-causing-unnecessary-heap-allocations-in-go
		valuesCopy := make([]any, len(values))
		copy(valuesCopy, values)

		log.Trace(format, valuesCopy...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func Debug(ctx context.Context, format string, values ...any) {
	if log, ok := fromContext(ctx, LevelDebug); ok {
		// perf: copy the values so that the compiler does not allocate on the heap. This prevents an allocation happening even
		//       when the log level is too low and the log function is never called.
		//       See: https://stackoverflow.com/questions/27788813/variadic-functions-causing-unnecessary-heap-allocations-in-go
		valuesCopy := make([]any, len(values))
		copy(valuesCopy, values)

		log.Debug(format, valuesCopy...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func Info(ctx context.Context, format string, values ...any) {
	if log, ok := fromContext(ctx, LevelInfo); ok {
		// perf: copy the values so that the compiler does not allocate on the heap. This prevents an allocation happening even
		//       when the log level is too low and the log function is never called.
		//       See: https://stackoverflow.com/questions/27788813/variadic-functions-causing-unnecessary-heap-allocations-in-go
		valuesCopy := make([]any, len(values))
		copy(valuesCopy, values)

		log.Info(format, valuesCopy...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func Warn(ctx context.Context, format string, values ...any) {
	if log, ok := fromContext(ctx, LevelWarn); ok {
		// perf: copy the values so that the compiler does not allocate on the heap. This prevents an allocation happening even
		//       when the log level is too low and the log function is never called.
		//       See: https://stackoverflow.com/questions/27788813/variadic-functions-causing-unnecessary-heap-allocations-in-go
		valuesCopy := make([]any, len(values))
		copy(valuesCopy, values)

		log.Warn(format, valuesCopy...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func Error(ctx context.Context, format string, values ...any) {
	if log, ok := fromContext(ctx, LevelError); ok {
		// perf: copy the values so that the compiler does not allocate on the heap. This prevents an allocation happening even
		//       when the log level is too low and the log function is never called.
		//       See: https://stackoverflow.com/questions/27788813/variadic-functions-causing-unnecessary-heap-allocations-in-go
		valuesCopy := make([]any, len(values))
		copy(valuesCopy, values)

		log.Error(format, valuesCopy...)
	}
}

// nolint:goprintffuncname // reason: keep in line with logger function naming
func Fatal(ctx context.Context, format string, values ...any) {
	if log, ok := fromContext(ctx, LevelFatal); ok {
		// perf: copy the values so that the compiler does not allocate on the heap. This prevents an allocation happening even
		//       when the log level is too low and the log function is never called.
		//       See: https://stackoverflow.com/questions/27788813/variadic-functions-causing-unnecessary-heap-allocations-in-go
		valuesCopy := make([]any, len(values))
		copy(valuesCopy, values)

		log.Fatal(format, valuesCopy...)
	}
}
