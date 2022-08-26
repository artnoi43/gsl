package data

type queueType any

type Queue[T queueType] []T

func (self *Queue[T]) Push(x T) {
	*self = append(*self, x)
}

func (self *Queue[T]) PushSlice(slice ...T) {
	for _, elem := range slice {
		self.Push(elem)
	}
}

// Pop pops and returns the left-most element of self,
// returning nil if self is empty
func (self *Queue[T]) Pop() *T {
	state := *self
	l := len(state)
	if l == 0 {
		return nil
	}

	var elem T
	elem, *self = state[0], state[1:l]

	return &elem
}

func (self *Queue[T]) Len() int {
	state := *self
	return len(state)
}

func (self *Queue[T]) IsEmpty() bool {
	state := *self
	return len(state) == 0
}

func NewQueue[T queueType]() *Queue[T] {
	return &Queue[T]{}
}
