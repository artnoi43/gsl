package list

import "sync"

// SafeList[T] wraps BasicList[T] and uses sync.RWMutex to avoid data races.
// L was added to make sure that the underlying list type is always accessible from the instance type,
// for example, SafeList[float64, *Queue[float64]]
// if L was not the type parameter, then a safe uint8 stack, safe uint8 queue, etc,
// will be indifferentiable with the same type `SafeList[uint8]`.
// SafeList[T, BasicList[T]] also implements BasicList[T]
type SafeList[T any, L BasicList[T]] struct {
	mut       *sync.RWMutex
	basicList L
}

// WrapSafeList[T] wraps a BasicList[T] into SafeList[T],
// where T is the underlying entity (item) type and L is the underlying BasicList[T] type.
// If you're wrapping a variable `foo“ of type `*Stack[uint8]`, then call this function with:
// WrapSafeList[uint8, *Stack[uint8]](foo)
func WrapSafeList[T any, L BasicList[T]](basicList L) *SafeList[T, L] {
	return &SafeList[T, L]{
		basicList: basicList,
		mut:       &sync.RWMutex{},
	}
}

func (self *SafeList[T, L]) Push(x T) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.basicList.Push(x)
}

func (self *SafeList[T, L]) Pop() *T {
	self.mut.Lock()
	defer self.mut.Unlock()

	return self.basicList.Pop()
}

func (self *SafeList[T, L]) Len() int {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.basicList.Len()
}

func (self *SafeList[T, L]) IsEmpty() bool {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.basicList.IsEmpty()
}
