package data

type Queue []interface{}

func (self *Queue) Push(x interface{}) {
	*self = append(*self, x)
}

func (self *Queue) Pop() interface{} {
	h := *self
	l := len(h)
	if l == 0 {
		return nil
	}

	var elem interface{}
	elem, *self = h[0], h[1:l]

	return elem
}

func (self *Queue) Len() int {
	h := *self
	return len(h)
}

func (self *Queue) IsEmpty() bool {
	h := *self
	return len(h) == 0
}

func NewQueue() *Queue {
	return &Queue{}
}
