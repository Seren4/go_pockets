package cache

import (
	"sync"
	"time"
)

// Cache is key-value storage.
type Cache[K comparable, V any] struct {
	// add a mutex next to the resource we want to protect (the data map).
	mu   sync.Mutex
	// rmu  sync.RWMutex 
	data map[K]entryWithTimeout[V]
	ttl  time.Duration
}

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time // After that time, the value is useless.
}

// New creates a usable Cache (it initializes the map for us).
func New[K comparable, V any](ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:  ttl,
		data: make(map[K]entryWithTimeout[V]),
	}
}

// This method accepts a key of the adequate type, and returns the value - also of the adequate type
func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var noValue V

	v, found := c.data[key]
	if !found {
		return noValue, false
	}
	if v.expires.Before(time.Now()) {
		delete(c.data, key)
		return noValue, false
	}
	return v.value, true
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = entryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	// Do not return an error for the moment,
	// but it can happen in the near future.
	return nil
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
