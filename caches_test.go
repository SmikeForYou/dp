package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTempCache(t *testing.T) {
	tcache := NewTempCache[string, string]()
	tcache.Push("key1", "value1", time.Second*2)
	time.Sleep(time.Second)
	v, ok := tcache.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", v)
	tcache.Push("key2", "value2", time.Second)
	time.Sleep(time.Second * 2)
	_, ok = tcache.Get("key2")
	assert.False(t, ok)
}
