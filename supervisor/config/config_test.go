package config

import (
	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestEntry_Load(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		teatPath := gdebug.CallerFilePath()
		config := NewConfig(teatPath + "/test.conf")
		err := config.Load()
		t.Assert(err, nil)
		entrys := config.GetPrograms()
		t.Assert(len(entrys), 6)
	})
}
