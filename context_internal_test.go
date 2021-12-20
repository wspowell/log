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

type testConfig struct {
	*Config
}

func (self *testConfig) Tags() map[string]any {
	return map[string]any{
		"global": "global",
	}
}

func (self *testConfig) Out() io.Writer {
	return io.Discard
}

func (self *testConfig) Logger() Logger {
	return NewLog(self)
}

func newTestConfig() *testConfig {
	config := NewConfig(LevelError)

	return &testConfig{
		Config: config,
	}
}

func Test_WithContext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx = WithContext(ctx, NewConfig(LevelDebug))

	log, ok := fromContext(ctx, LevelDebug)
	assert.True(t, ok)
	assert.NotNil(t, log)
	assert.IsType(t, Log{}, log)
}

func Test_Context_Tags_Localized(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	config := newTestConfig()

	ctx = WithContext(ctx, config)

	Tag(ctx, "parent", "parent")

	parentTags := Tags(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(boundaryCtx context.Context) {
		defer wg.Done()

		// Use localized Logger.
		childCtx := context.Localize(boundaryCtx)

		Tag(childCtx, "test1", value1)
		Tag(childCtx, "test2", value2)

		cfg, ok := childCtx.Value(configContextKey{}).(Configer)
		assert.True(t, ok)
		var value any
		if value, ok = cfg.Tags()[global]; !ok || value != global {
			t.Error("missing global tag")
		}

		childTags := Tags(childCtx)
		if &childTags == &parentTags {
			t.Error("child tags must not address the same parent tags")
		}

		log, ok := fromContext(childCtx, LevelError)
		assert.True(t, ok)
		if value, ok = log.Tags()[global]; ok || value == global {
			t.Error("global tag should not exist")
		}
		if value, ok = log.Tags()[parent]; !ok || value != parent {
			t.Error("missing parent tag")
		}
		if value, ok = log.Tags()["test1"]; !ok || value != value1 {
			t.Error("missing test1 tag")
		}
		if value, ok = log.Tags()["test2"]; !ok || value != value2 {
			t.Error("missing test2 tag")
		}

		// Override Log explicitly.
		childCtx = WithContext(childCtx, config)

		Tag(childCtx, "test1", value1)
		Tag(childCtx, "test2", value2)

		cfg, ok = childCtx.Value(configContextKey{}).(Configer)
		assert.True(t, ok)
		if value, ok = cfg.Tags()[global]; !ok || value != global {
			t.Error("missing global tag")
		}

		ovrrideTags := Tags(childCtx)
		if &ovrrideTags == &parentTags {
			t.Error("override tags must not address the same parent tags")
		}
		if &ovrrideTags == &childTags {
			t.Error("override tags must not address the same child tags")
		}

		log, ok = fromContext(childCtx, LevelError)
		assert.True(t, ok)
		if value, ok := log.Tags()[global]; ok || value == global {
			t.Error("global tag should not exist")
		}
		if value, ok := log.Tags()[parent]; ok || value == parent {
			t.Error("parent tag should not exist")
		}
		if value, ok := log.Tags()["test1"]; !ok || value != value1 {
			t.Error("missing test1 tag")
		}
		if value, ok := log.Tags()["test2"]; !ok || value != value2 {
			t.Error("missing test2 tag")
		}
	}(ctx)

	wg.Wait()

	cfg, ok := ctx.Value(configContextKey{}).(Configer)
	assert.True(t, ok)
	var value any
	if value, ok = cfg.Tags()[global]; !ok || value != global {
		t.Error("missing global tag")
	}

	log, ok := fromContext(ctx, LevelError)
	assert.True(t, ok)
	if value, ok := log.Tags()[global]; ok || value == global {
		t.Error("global tag should not exist")
	}
	if value, ok := log.Tags()[parent]; !ok || value != parent {
		t.Error("missing parent tag")
	}
	if value, ok := log.Tags()["test1"]; ok || value == value1 {
		t.Error("test1 tag should not exist")
	}
	if value, ok := log.Tags()["test2"]; ok || value == value2 {
		t.Error("test2 tag should not exist")
	}
}
