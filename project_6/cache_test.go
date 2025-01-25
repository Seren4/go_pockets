package cache_test

import (
	"fmt"
	cache "learngo-pockets/genericcache"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := cache.New[string, string](3, time.Millisecond * 100)

	c.Upsert("serena", "06778899")
	value, found := c.Read("serena")
	assert.True(t, found)
	assert.Equal(t, "06778899", value)
	assert.Equal(t, reflect.TypeOf(value).String(), "string")

	time.Sleep(time.Millisecond * 200)
	v, f := c.Read("serena")
	assert.False(t, f)
	assert.Equal(t, "", v)

	value2, found2 := c.Read("alice")
	assert.False(t, found2)
	assert.Equal(t, "", value2)

	c.Upsert("serena", "06778800")
	value3, found3 := c.Read("serena")
	assert.True(t, found3)
	assert.Equal(t, "06778800", value3)

}

func TestCache_Parallel_goroutines(t *testing.T) {

	c := cache.New[int, string](10, 30)

	const parallelTasks = 10

	wg := sync.WaitGroup{}

	wg.Add(parallelTasks)

	for i := 0; i < parallelTasks; i++ {
		go func(j int) {
			defer wg.Done()
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}
	wg.Wait()
}

// Alternatively, we can make use of the testing package to execute parallel tests.
func TestCache_Parallel(t *testing.T) {

	c := cache.New[int, string](2, 30)

	// This goroutine can be executed along with another.
	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	// This goroutine can be executed along with another.
	t.Run("write kuus", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}

func TestCache_TTL(t *testing.T) {
	t.Parallel()
	c := cache.New[string, string](2, time.Millisecond * 100)
	c.Upsert("Norwegian", "Blue")
	// Check the item is there.
	got, found := c.Read("Norwegian")
	assert.True(t, found)
	assert.Equal(t, "Blue", got)
	time.Sleep(time.Millisecond * 200)
	got, found = c.Read("Norwegian")
	assert.False(t, found)
	assert.Equal(t, "", got)
}

// TestCache_MaxSize tests the maximum capacity feature of a cache.
// It checks that update items are properly requeued as "new" items,
// and that we make room by removing the most ancient item for the new ones.
func TestCache_MaxSize(t *testing.T) {
	t.Parallel()                                                  
	// Give it a TTL long enough to survive this test
	c := cache.New[int, int](3, time.Minute) 
	c.Upsert(1, 1)
	c.Upsert(2, 2)
	c.Upsert(3, 3)

	got, found := c.Read(1)
	assert.True(t, found)
	assert.Equal(t, 1, got)

	// Update 1, which will no longer make it the oldest
	c.Upsert(1, 10)
	// Adding a fourth element will discard the oldest - 2 in this case.
	c.Upsert(4, 4)
	// Trying to retrieve an element that should've been discarded by now.
	got, found = c.Read(2)
	assert.False(t, found)
	assert.Equal(t, 0, got)
}