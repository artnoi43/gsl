package data

type Stack []interface{}

func (self *Stack) Push(x interface{}) {
	*self = append(*self, x)
}

func (self *Stack) Pop() interface{} {
	h := *self
	l := len(h)
	if l == 0 {
		return nil
	}

	var elem interface{}
	elem, *self = h[l-1], h[0:l-1]

	return elem
}

func (self *Stack) Len() int {
	h := *self
	return len(h)
}

func (self *Stack) IsEmpty() bool {
	h := *self
	return len(h) == 0
}

func NewStack() *Stack {
	return &Stack{}
}
