package dp

import "time"

// Cache is a generic map-based cache.
// K is the type of the keys, and V is the type of the values.
type Cache[K comparable, V any] map[K]V

// NewCache creates a new empty Cache.
func NewCache[K comparable, V any]() Cache[K, V] {
	return make(map[K]V)
}

// NewCacheFromArr creates a new Cache from a slice of values.
// The visitor function is used to generate the key-value pairs for the cache.
func NewCacheFromArr[K comparable, V any](data []V, visitor func(elem V) (K, V)) Cache[K, V] {
	cache := NewCache[K, V]()
	for _, v := range data {
		cache.Push(visitor(v))
	}
	return cache
}

// Push adds a key-value pair to the Cache.
func (c Cache[K, V]) Push(key K, val V) {
	c[key] = val
}

// Exists checks if a key exists in the Cache.
func (c Cache[K, V]) Exists(key K) bool {
	_, ok := c[key]
	return ok
}

// Get retrieves a value from the Cache by key.
// It returns the value and a boolean indicating whether the key was found.
func (c Cache[K, V]) Get(key K) (V, bool) {
	v, ok := c[key]
	return v, ok
}

// GetWithCallback retrieves a value from the Cache by key.
// If the key is not found, the onFail callback is called to provide a value.
func (c Cache[K, V]) GetWithCallback(key K, onFail func(key K) (V, bool)) (V, bool) {
	v, ok := c[key]
	if ok {
		return v, ok
	}
	return onFail(key)
}

// Remove deletes a key-value pair from the Cache by key.
func (c Cache[K, V]) Remove(key K) {
	delete(c, key)
}

// tcacheval is a struct that holds a value and its expiration time.
type tcacheval[V any] struct {
	exp   time.Time
	value V
}

// TempCache is a cache with time-based expiration for its values.
// K is the type of the keys, and V is the type of the values.
type TempCache[K comparable, V any] struct {
	cache Cache[K, tcacheval[V]]
}

// NewTempCache creates a new empty TempCache.
func NewTempCache[K comparable, V any]() TempCache[K, V] {
	cache := make(map[K]tcacheval[V])
	return TempCache[K, V]{
		cache: cache,
	}
}

// Push adds a key-value pair to the TempCache with a time-to-live (TTL).
func (tc TempCache[K, V]) Push(key K, val V, ttl time.Duration) {
	tc.cache[key] = tcacheval[V]{exp: time.Now().Add(ttl), value: val}
}

// Get retrieves a value from the TempCache by key.
// If the value has expired, it is removed from the cache and a zero value is returned.
func (tc TempCache[K, V]) Get(key K) (V, bool) {
	tcv, ok := tc.cache[key]
	if tcv.exp.Before(time.Now()) {
		delete(tc.cache, key)
		var t V
		return t, false
	}
	return tcv.value, ok
}

// Remove deletes a key-value pair from the TempCache by key.
func (tc TempCache[K, V]) Remove(key K) {
	delete(tc.cache, key)
}
