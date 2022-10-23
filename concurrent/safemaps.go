package concurrent

import "sync"

type SafeMap[K comparable, T any] interface {
	Get(K) T
	Set(K, T)
}

type safeMapImpl[K comparable, T any] struct {
	sync.RWMutex
	wrappedMap map[K]T
}

func NewSafeMap[K comparable, T any](m map[K]T) SafeMap[K, T] {
	return &safeMapImpl[K, T]{
		wrappedMap: m,
	}
}

func (m *safeMapImpl[K, T]) Get(key K) T {
	m.RLock()
	defer m.RUnlock()

	return m.wrappedMap[key]
}

func (m *safeMapImpl[K, T]) Set(key K, value T) {
	m.Lock()
	defer m.Unlock()

	m.wrappedMap[key] = value
}
