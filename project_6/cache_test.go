package cache_test

import (
	"testing"
	cache "learngo-pockets/genericcache"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestCache(t *testing.T) {
	c := cache.New[string, string]()
	
	c.Upsert("serena", "06778899")
	value, found := c.Read("serena")
	assert.True(t, found)
	assert.Equal(t, "06778899", value)
	assert.Equal(t, reflect.TypeOf(value).String(), "string")

	value2, found2 := c.Read("alice")
	assert.False(t, found2)
	assert.Equal(t, "", value2)

	c.Upsert("serena", "06778800")
	value3, found3 := c.Read("serena")
	assert.True(t, found3)
	assert.Equal(t, "06778800", value3)

}