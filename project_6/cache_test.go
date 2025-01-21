package cache_test

import (
	"testing"
	cache "learngo-pockets/genericcache"
  "github.com/stretchr/testify/assert"
)

// You can start writing a unit test that writes, reads, 
// checks the returned type, checks the returned value for an absent key,
// writes another value for the same key, etc.

func TestCache(t *testing.T) {
	c := cache.New[string, string]()
	
	c.Upsert("serena", "06778899")
	value, found := c.Read("serena")
	assert.True(t, found)
	assert.Equal(t, "06778899", value)

}