package list

import "sync"

// SafeList[T] wraps BasicList[T] and uses sync.RWMutex to avoid data races.
type SafeList[T any] struct {
	basicList BasicList[T]
	mut       *sync.RWMutex
}

// WrapSafeList[T] wraps a BasicList[T] into SafeList[T].
func WrapSafeList[T any](basicList BasicList[T]) *SafeList[T] {
	return &SafeList[T]{
		basicList: basicList,
		mut:       &sync.RWMutex{},
	}
}

func (self *SafeList[T]) Push(x T) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.basicList.Push(x)
}

func (self *SafeList[T]) Pop() *T {
	self.mut.Lock()
	defer self.mut.Unlock()

	return self.basicList.Pop()
}

func (self *SafeList[T]) Len() int {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.basicList.Len()
}

func (self *SafeList[T]) IsEmpty() bool {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.basicList.IsEmpty()
}
