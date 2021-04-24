package log

import (
	"fmt"
	"testing"

	"github.com/wspowell/context"
)

func Benchmark_Context_debug_at_error_level(b *testing.B) {
	ctx := context.Local()
	ctx = WithContext(ctx, newTestConfig(LevelError))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debug(ctx, "hello: %s", "world")
	}
}

func Benchmark_Context_debug_at_error_level_with_complex_parameters(b *testing.B) {
	ctx := context.Local()
	ctx = WithContext(ctx, newTestConfig(LevelError))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debug(ctx, fmt.Sprintf("hello: %s", "world"))
	}
}

func Benchmark_Context_error_at_error_level(b *testing.B) {
	ctx := context.Local()
	ctx = WithContext(ctx, newTestConfig(LevelError))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Error(ctx, "hello: %s", "world")
	}
}

func Benchmark_Context_debug_at_error_level_1000x(b *testing.B) {
	ctx := context.Local()
	ctx = WithContext(ctx, newTestConfig(LevelError))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 1000; k++ {
			Debug(ctx, "hello: %s", "world")
		}
	}
}

func Benchmark_Context_error_at_error_level_1000x(b *testing.B) {
	ctx := context.Local()
	ctx = WithContext(ctx, newTestConfig(LevelError))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 1000; k++ {
			Error(ctx, "hello: %s", "world")
		}
	}
}

func Benchmark_Context_WithContext(b *testing.B) {
	config := NewConfig(LevelError)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WithContext(context.Local(), config)
	}
}
