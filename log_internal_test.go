package log

import (
	"sync"
	"testing"
)

func Test_Log_GlobalTags_race(t *testing.T) {
	t.Parallel()

	cfg := newTestConfig()
	log := NewLog(cfg)

	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Error("test")
		}()
	}

	wg.Wait()
}

func Test_Log_Tags(t *testing.T) {
	t.Parallel()

	config := newTestConfig()
	config.Tags()["global"] = "global"

	log := NewLog(config)

	log.Tag("test1", "value1")
	log.Tag("test2", "value2")

	if value, ok := tagsContainKey(log.tags, "global"); ok || value == "global" {
		t.Error("logger should not have global tag")
	}

	if value, ok := tagsContainKey(log.tags, "test1"); !ok || value != "value1" {
		t.Error("missing test1 tag")
	}

	if value, ok := tagsContainKey(log.tags, "test2"); !ok || value != "value2" {
		t.Error("missing test2 tag")
	}
}

func Test_Log_Tags_nil(t *testing.T) {
	t.Parallel()

	log := NewLog(NewConfig().WithLevel(LevelError))

	log.Tag("test1", "value1")
	log.Tag("test2", "value2")

	if value, ok := tagsContainKey(log.tags, "test1"); !ok || value != "value1" {
		t.Fail()
	}

	if value, ok := tagsContainKey(log.tags, "test2"); !ok || value != "value2" {
		t.Fail()
	}
}
