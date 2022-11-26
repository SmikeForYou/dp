package dp

type Cache[K comparable, V any] map[K]V

func NewCache[K comparable, V any]() Cache[K, V] {
	return make(map[K]V)
}

func NewCacheFromArr[K comparable, V any](data []V, visitor func(elem V) (K, V)) Cache[K, V] {
	cache := NewCache[K, V]()
	for _, v := range data {
		cache.Push(visitor(v))
	}
	return cache
}

func (c Cache[K, V]) Push(key K, val V) {
	c[key] = val
}

func (c Cache[K, V]) Exists(key K) bool {
	_, ok := c[key]
	return ok
}

func (c Cache[K, V]) Get(key K) (V, bool) {
	v, ok := c[key]
	return v, ok
}
func (c Cache[K, V]) GetWithCallback(key K, onFail func(key K) (V, bool)) (V, bool) {
	v, ok := c[key]
	if ok {
		return v, ok
	}
	return onFail(key)
}

func (c Cache[K, V]) Remove(key K) {
	delete(c, key)
}
