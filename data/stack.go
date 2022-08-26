package data

type stackType any

type Stack[T stackType] []T

func (self *Stack[T]) Push(x T) {
	*self = append(*self, x)
}

// Pop pops and returns the left-most element of self,
// returning nil if self is empty
func (self *Stack[T]) Pop() *T {
	state := *self
	l := len(state)
	if l == 0 {
		return nil
	}

	var elem T
	elem, *self = state[l-1], state[0:l-1]

	return &elem
}

func (self *Stack[T]) Len() int {
	state := *self
	return len(state)
}

func (self *Stack[T]) IsEmpty() bool {
	state := *self
	return len(state) == 0
}

func NewStack[T stackType]() *Stack[T] {
	return &Stack[T]{}
}
