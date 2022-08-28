package wgraph

type UndirectedNodeImpl[T graphWeight, S ~string] struct {
	Name    S
	Cost    T
	Through UndirectedNode[T, S]
}

// Implements data.ItemPQ[T]
func (self *UndirectedNodeImpl[T, S]) GetValue() T {
	return self.Cost
}
func (self *UndirectedNodeImpl[T, S]) GetKey() S {
	return self.Name
}
func (self *UndirectedNodeImpl[T, S]) GetThrough() UndirectedNode[T, S] {
	return self.Through
}
func (self *UndirectedNodeImpl[T, S]) SetValue(value T) {
	self.Cost = value
}

func (self *UndirectedNodeImpl[T, S]) SetThrough(node UndirectedNode[T, S]) {
	self.Through = node
}
