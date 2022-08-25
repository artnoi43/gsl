package data

type Queue []interface{}

func (self *Queue) Push(x interface{}) {
	*self = append(*self, x)
}

func (self *Queue) Pop() interface{} {
	h := *self
	var elem interface{}
	l := len(h)
	elem, *self = h[0], h[1:l]

	return elem
}

func (self *Queue) Len() int {
	h := *self
	return len(h)
}

func NewQueue() *Queue {
	return &Queue{}
}
