package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wspowell/local"
)

func Test_WithContext(t *testing.T) {
	ctx := local.NewLocalized()

	WithContext(ctx, NewConfig(LevelDebug))

	assert.NotNil(t, fromContext(ctx))
	assert.IsType(t, Log{}, fromContext(ctx))
}

func Test_WithContext_child(t *testing.T) {
	ctx := local.NewLocalized()

	WithContext(ctx, NewConfig(LevelDebug))

	assert.NotNil(t, fromContext(ctx))
	assert.IsType(t, Log{}, fromContext(ctx))

	childCtx := local.FromContext(ctx)

	assert.NotNil(t, fromContext(childCtx))
	assert.IsType(t, Log{}, fromContext(childCtx))
}

func Test_Context_Tags(t *testing.T) {
	ctx := local.NewLocalized()

	config := newTestConfig(LevelError)
	config.Tags()["global"] = "global"

	WithContext(ctx, config)

	Tag(ctx, "parent", "parent")

	childCtx := local.FromContext(ctx)

	Tag(childCtx, "test1", "value1")
	Tag(childCtx, "test2", "value2")

	cfg, ok := childCtx.Value(configContextKey{}).(Configer)
	assert.True(t, ok)

	if value, ok := cfg.Tags()["global"]; !ok || value != "global" {
		t.Error("missing global tag")
	}

	if value, ok := fromContext(ctx).(Log).tags["parent"]; !ok || value != "parent" {
		t.Error("missing parent tag")
	}

	if value, ok := fromContext(childCtx).(Log).tags["parent"]; ok || value == "parent" {
		t.Error("parent tag should not exist")
	}

	if value, ok := fromContext(childCtx).(Log).tags["test1"]; !ok || value != "value1" {
		t.Error("missing test1 tag")
	}

	if value, ok := fromContext(childCtx).(Log).tags["test2"]; !ok || value != "value2" {
		t.Error("missing test2 tag")
	}
}
