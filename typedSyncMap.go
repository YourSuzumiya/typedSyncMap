package typedSyncMap

import (
	"sync"
)

type TypedSyncMap[K comparable, V any] struct {
	m sync.Map
}

func (t *TypedSyncMap[K, V]) Load(key K) (V, bool) {
	value, ok := t.m.Load(key)
	if !ok {
		var zeroValue V
		return zeroValue, false
	}
	return value.(V), true
}

func (t *TypedSyncMap[K, V]) Store(key K, value V) {
	t.m.Store(key, value)
}

func (t *TypedSyncMap[K, V]) Delete(key K) {
	t.m.Delete(key)
}

func (t *TypedSyncMap[K, V]) Range(f func(key K, value V) bool) {
	t.m.Range(func(k, v interface{}) bool {
		return f(k.(K), v.(V))
	})
}
