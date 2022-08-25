package data

type Stack []interface{}

func (self *Stack) Push(x interface{}) {
	*self = append(*self, x)
}

func (self *Stack) Pop() interface{} {
	h := *self
	var elem interface{}
	l := len(h)
	elem, *self = h[l-1], h[0:l-1]

	return elem
}

func (self *Stack) Len() int {
	h := *self
	return len(h)
}

func NewStack() *Stack {
	return &Stack{}
}
