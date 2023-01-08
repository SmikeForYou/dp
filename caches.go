package dp

import "time"

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

type tcacheval[V any] struct {
	exp   time.Time
	value V
}

type TempCache[K comparable, V any] struct {
	cache Cache[K, tcacheval[V]]
}

func NewTempCache[K comparable, V any]() TempCache[K, V] {
	cache := make(map[K]tcacheval[V])
	return TempCache[K, V]{
		cache: cache,
	}
}

func (tc TempCache[K, V]) Push(key K, val V, ttl time.Duration) {
	tc.cache[key] = tcacheval[V]{exp: time.Now().Add(ttl), value: val}
}

func (tc TempCache[K, V]) Get(key K) (V, bool) {
	tcv, ok := tc.cache[key]
	if tcv.exp.Before(time.Now()) {
		delete(tc.cache, key)
		var t V
		return t, false
	}
	return tcv.value, ok
}

func (tc TempCache[K, V]) Remove(key K) {
	delete(tc.cache, key)
}
