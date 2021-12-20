//go:build !release
// +build !release

package frametest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"

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

	type logStructure struct {
		Level    string    `json:"level"`
		Global   string    `json:"global"`
		Test1    string    `json:"test1"`
		Test2    string    `json:"test2"`
		Function string    `json:"function"`
		Message  string    `json:"message"`
		Time     time.Time `json:"time"`
	}

	expectedLogStructure := &logStructure{
		Level:    "error",
		Global:   "global",
		Test1:    "value1",
		Test2:    "value2",
		Function: "github.com/wspowell/log/frametest_test.Test_Log_Error_NonContext:58",
		Message:  "error: test",
		Time:     time.Time{},
	}

	actualLogStructure := &logStructure{}
	assert.Nil(t, json.Unmarshal(cfg.logCapture.Bytes(), actualLogStructure))

	assert.NotEqual(t, time.Time{}, actualLogStructure.Time)
	actualLogStructure.Time = time.Time{}

	fmt.Println(cfg.logCapture.String())
	assert.Equal(t, expectedLogStructure, actualLogStructure)
}

func Test_Log_Error_Context(t *testing.T) {
	t.Parallel()

	cfg := newTestConfig(log.LevelError)
	ctx := context.Background()
	ctx = log.WithContext(ctx, cfg)

	log.Tag(ctx, "test1", "value1")
	log.Tag(ctx, "test2", "value2")

	log.Error(ctx, "error: %s", "test")

	type logStructure struct {
		Level    string    `json:"level"`
		Global   string    `json:"global"`
		Test1    string    `json:"test1"`
		Test2    string    `json:"test2"`
		Function string    `json:"function"`
		Message  string    `json:"message"`
		Time     time.Time `json:"time"`
	}

	expectedLogStructure := &logStructure{
		Level:    "error",
		Global:   "global",
		Test1:    "value1",
		Test2:    "value2",
		Function: "github.com/wspowell/log/frametest_test.Test_Log_Error_Context:100",
		Message:  "error: test",
		Time:     time.Time{},
	}

	actualLogStructure := &logStructure{}
	assert.Nil(t, json.Unmarshal(cfg.logCapture.Bytes(), actualLogStructure))

	assert.NotEqual(t, time.Time{}, actualLogStructure.Time)
	actualLogStructure.Time = time.Time{}

	fmt.Println(cfg.logCapture.String())
	assert.Equal(t, expectedLogStructure, actualLogStructure)
}
