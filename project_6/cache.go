package cache

import (
	"slices"
	"sync"
	"time"
)

// Cache is key-value storage.
type Cache[K comparable, V any] struct {
	// add a mutex next to the resource we want to protect (the data map).
	mu sync.Mutex
	// rmu  sync.RWMutex
	data              map[K]entryWithTimeout[V]
	ttl               time.Duration
	maxSize           int
	chronologicalKeys []K
}

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time // After that time, the value is useless.
}

// New creates a usable Cache (it initializes the map for us).
func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:               ttl,
		data:              make(map[K]entryWithTimeout[V]),
		maxSize:           maxSize,
		chronologicalKeys: make([]K, 0, maxSize),
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
		c.deleteKeyValue(key)
		return noValue, false
	}
	return v.value, true
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, found := c.data[key]

	switch {
	case found:
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		c.deleteKeyValue(c.chronologicalKeys[0])
	}
	c.addKeyValue(key, value)

	// Do not return an error for the moment,
	// but it can happen in the near future.
	return nil
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.deleteKeyValue(key)
}

// addKeyValue inserts a key and its value into the cache.
func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = entryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

// deleteKeyValue removes a key and its associated value from the cache.
func (c *Cache[K, V]) deleteKeyValue(key K) {
	delete(c.data, key)
	c.chronologicalKeys = slices.DeleteFunc(c.chronologicalKeys, func(k K) bool { return k == key })
}
