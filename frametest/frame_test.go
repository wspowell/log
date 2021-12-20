//go:build !release
// +build !release

package frametest_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wspowell/context"

	"github.com/wspowell/log"
)

type testConfig struct {
	*log.Config

	logCapture *bytes.Buffer
}

func newTestConfig(level log.Level) *testConfig {
	config := log.NewConfig(level)

	return &testConfig{
		Config:     config,
		logCapture: bytes.NewBuffer([]byte{}),
	}
}

func (self *testConfig) Tags() map[string]any {
	return map[string]any{
		"global": "global",
	}
}

func (self *testConfig) Out() io.Writer {
	return self.logCapture
}

func (self *testConfig) Logger() log.Logger {
	return log.NewLog(self)
}

func Test_Log_Error_NonContext(t *testing.T) {
	t.Parallel()

	cfg := newTestConfig(log.LevelError)
	logger := log.NewLog(cfg)

	logger.Tag("test1", "value1")
	logger.Tag("test2", "value2")

	logger.Error("error: %s", "test")

	// Log minus the timestamp.
	expectedLog := `level=error msg="error: test" function=github.com/wspowell/log/frametest_test.Test_Log_Error_NonContext global=global test1=value1 test2=value2` + "\n"

	assert.True(t, strings.HasSuffix(cfg.logCapture.String(), expectedLog), "actual: %s", cfg.logCapture.String())
}

func Test_Log_Error_Context(t *testing.T) {
	t.Parallel()

	cfg := newTestConfig(log.LevelError)
	ctx := context.Local()
	ctx = log.WithContext(ctx, cfg)

	log.Tag(ctx, "test1", "value1")
	log.Tag(ctx, "test2", "value2")

	log.Error(ctx, "error: %s", "test")

	// Log minus the timestamp.
	expectedLog := `level=error msg="error: test" function=github.com/wspowell/log/frametest_test.Test_Log_Error_Context global=global test1=value1 test2=value2` + "\n"

	assert.True(t, strings.HasSuffix(cfg.logCapture.String(), expectedLog), "actual: %s", cfg.logCapture.String())
}
