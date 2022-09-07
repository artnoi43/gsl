package list

type StackImpl[T any] []T

func NewStack[T any]() SafeList[T, *StackImpl[T]] {
	return WrapSafeList[T](new(StackImpl[T]))
}

func NewStackUnsafe[T any]() *StackImpl[T] {
	return new(StackImpl[T])
}

func (self *StackImpl[T]) Push(x T) {
	*self = append(*self, x)
}

func (self *StackImpl[T]) PushSlice(slice []T) {
	for _, elem := range slice {
		self.Push(elem)
	}
}

// Pop pops and returns the right-most element of self,
// returning nil if self is empty
func (self *StackImpl[T]) Pop() *T {
	state := *self
	l := len(state)
	if l == 0 {
		return nil
	}

	var elem T
	elem, *self = state[l-1], state[0:l-1]

	return &elem
}

func (self *StackImpl[T]) Len() int {
	state := *self
	return len(state)
}

func (self *StackImpl[T]) IsEmpty() bool {
	state := *self
	return len(state) == 0
}
