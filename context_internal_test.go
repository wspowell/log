package log

import (
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wspowell/context"
)

const (
	global = "global"
	parent = "parent"
	value1 = "value1"
	value2 = "value2"
)

type testConfig struct{}

func (self *testConfig) Tags() map[string]any {
	return map[string]any{
		"global": "global",
	}
}

func (self *testConfig) Output() io.Writer {
	return io.Discard
}

func (self *testConfig) Level() Level {
	return LevelError
}

func newTestConfig() *testConfig {
	return &testConfig{}
}

func Test_WithContext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx = WithContext(ctx, NewConfig().WithLevel(LevelDebug))

	log, ok := fromContext(ctx, LevelDebug)
	assert.True(t, ok)
	assert.NotNil(t, log)
	assert.IsType(t, &Log{}, log)
}

func Test_Context_Tags_Localized(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	config := newTestConfig()

	ctx = WithContext(ctx, config)

	Tag(ctx, "parent", "parent")

	logger, _ := getLogger(ctx)
	parentTags := append([]any{}, logger.(*Log).tags...)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(boundaryCtx context.Context) {
		defer wg.Done()

		// Use localized Logger.
		childCtx := context.Localize(boundaryCtx)

		Tag(childCtx, "test1", value1)
		Tag(childCtx, "test2", value2)

		cfg, ok := childCtx.Value(configContextKey{}).(LoggerConfig)
		assert.True(t, ok)
		var value any
		if value, ok = cfg.Tags()[global]; !ok || value != global {
			t.Error("missing global tag")
		}

		logger, _ := getLogger(childCtx)
		childTags := append([]any{}, logger.(*Log).tags...)
		if &childTags == &parentTags {
			t.Error("child tags must not address the same parent tags")
		}

		if value, ok = tagsContainKey(childTags, global); ok || value == global {
			t.Error("child tags must not include global tag")
		}
		if value, ok = tagsContainKey(childTags, parent); !ok || value != parent {
			t.Error("missing parent tag")
		}
		if value, ok = tagsContainKey(childTags, "test1"); !ok || value != value1 {
			t.Error("missing test1 tag")
		}
		if value, ok = tagsContainKey(childTags, "test2"); !ok || value != value2 {
			t.Error("missing test2 tag")
		}

		// Override Log explicitly.
		childCtx = WithContext(childCtx, config)

		Tag(childCtx, "test1", value1)
		Tag(childCtx, "test2", value2)

		cfg, ok = childCtx.Value(configContextKey{}).(LoggerConfig)
		assert.True(t, ok)
		if value, ok = cfg.Tags()[global]; !ok || value != global {
			t.Error("missing global tag")
		}

		logger, _ = getLogger(childCtx)
		ovrrideTags := append([]any{}, logger.(*Log).tags...)
		if &ovrrideTags == &parentTags {
			t.Error("override tags must not address the same parent tags")
		}
		if &ovrrideTags == &childTags {
			t.Error("override tags must not address the same child tags")
		}

		assert.True(t, ok)
		if value, ok := tagsContainKey(childTags, global); ok || value == global {
			t.Error("child tags must not include global tag")
		}
		if value, ok := tagsContainKey(childTags, parent); !ok || value != parent {
			t.Error("parent tag should exist")
		}
		if value, ok := tagsContainKey(childTags, "test1"); !ok || value != value1 {
			t.Error("missing test1 tag")
		}
		if value, ok := tagsContainKey(childTags, "test2"); !ok || value != value2 {
			t.Error("missing test2 tag")
		}
	}(ctx)

	wg.Wait()

	cfg, ok := ctx.Value(configContextKey{}).(LoggerConfig)
	assert.True(t, ok)
	var value any
	if value, ok = cfg.Tags()[global]; !ok || value != global {
		t.Error("missing global tag")
	}

	if value, ok := tagsContainKey(parentTags, global); ok || value == global {
		t.Error("parent tags should not have global tag")
	}
	if value, ok := tagsContainKey(parentTags, parent); !ok || value != parent {
		t.Error("missing parent tag")
	}
	if value, ok := tagsContainKey(parentTags, "test1"); ok || value == value1 {
		t.Error("test1 tag should not exist")
	}
	if value, ok := tagsContainKey(parentTags, "test2"); ok || value == value2 {
		t.Error("test2 tag should not exist")
	}
}

func tagsContainKey(tags []any, key string) (interface{}, bool) {
	for i := 0; i < len(tags); i += 2 {
		if tags[i].(string) == key {
			return tags[i+1], true
		}
	}
	return nil, false
}
