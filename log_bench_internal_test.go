package log

import (
	"fmt"
	"io"
	"testing"

	"github.com/rs/zerolog/log"
)

func Benchmark_debug_at_error_level(b *testing.B) {
	log := NewLog(newTestConfig())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("hello: %s", "world")
	}
}

func Benchmark_debug_at_error_level_with_complex_parameters(b *testing.B) {
	log := NewLog(newTestConfig())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug(fmt.Sprintf("hello: %s", "world"))
	}
}

func Benchmark_error_at_error_level(b *testing.B) {
	log := NewLog(newTestConfig())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Error("hello: %s", "world")
	}
}

func Benchmark_debug_at_error_level_1000x(b *testing.B) {
	log := NewLog(newTestConfig())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 1000; k++ {
			log.Debug("hello: %s", "world")
		}
	}
}

func Benchmark_error_at_error_level_1000x(b *testing.B) {
	log := NewLog(newTestConfig())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 1000; k++ {
			log.Error("hello: %s", "world")
		}
	}
}

func Benchmark_error_at_error_level_1000x_zerolog(b *testing.B) {
	logger := log.Output(io.Discard).With().Fields(map[string]interface{}{
		"global": "global",
		"test1":  "value1",
		"test2":  "value2",
	}).Logger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 1000; k++ {
			logger.Error().Msgf("hello: %s", "world")
		}
	}
}

func Benchmark_NewLog(b *testing.B) {
	config := NewConfig().WithLevel(LevelError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewLog(config)
	}
}

func Benchmark_ConfigCopy_NewLog(b *testing.B) {
	config := NewConfig().WithLevel(LevelError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cfg := config
		NewLog(cfg)
	}
}

func Benchmark_Tag(b *testing.B) {
	log := NewLog(newTestConfig())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Tag("tag", "value")
	}
}
