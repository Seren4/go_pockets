package cache

// Cache is key-value storage.
type Cache[K comparable, V any] struct {
	data  map[K]V
}

// New creates a usable Cache (it initializes the map for us).
func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
				 data: make(map[K]V),
} }

//This method accepts a key of the adequate type, and returns the value - also of the adequate type
func (c *Cache[K, V]) Read(key K) (V, bool){
	v, found := c.data[key]
	return v, found
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.data[key] = value
	// Do not return an error for the moment,
  // but it can happen in the near future.
  return nil
}