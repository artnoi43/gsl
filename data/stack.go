package data

type stackType any

type Stack[T stackType] []T

func (self *Stack[T]) Push(x T) {
	*self = append(*self, x)
}

// Pop pops and returns the left-most element of self,
// returning nil if self is empty
func (self *Stack[T]) Pop() *T {
	h := *self
	l := len(h)
	if l == 0 {
		return nil
	}

	var elem T
	elem, *self = h[l-1], h[0:l-1]

	return &elem
}

func (self *Stack[T]) Len() int {
	h := *self
	return len(h)
}

func (self *Stack[T]) IsEmpty() bool {
	h := *self
	return len(h) == 0
}

func NewStack[T stackType]() *Stack[T] {
	return &Stack[T]{}
}
