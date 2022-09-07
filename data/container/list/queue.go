package list

type QueueImpl[T any] []T

func NewSafeQueue[T any]() SafeList[T, *QueueImpl[T]] {
	return WrapSafeList[T](new(QueueImpl[T]))
}

func (self *QueueImpl[T]) Push(x T) {
	*self = append(*self, x)
}

func (self *QueueImpl[T]) PushSlice(slice []T) {
	for _, elem := range slice {
		self.Push(elem)
	}
}

// Pop pops and returns the left-most element of self,
// returning nil if self is empty
func (self *QueueImpl[T]) Pop() *T {
	state := *self
	l := len(state)
	if l == 0 {
		return nil
	}

	var elem T
	elem, *self = state[0], state[1:l]

	return &elem
}

func (self *QueueImpl[T]) Len() int {
	state := *self
	return len(state)
}

func (self *QueueImpl[T]) IsEmpty() bool {
	state := *self
	return len(state) == 0
}

func NewQueue[T any]() SafeList[T, *QueueImpl[T]] {
	return WrapSafeList[T](new(QueueImpl[T]))
}

func NewQueueUnsafe[T any]() *QueueImpl[T] {
	return new(QueueImpl[T])
}
