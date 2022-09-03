package wgraph

type WeightedNodeImpl[T graphWeight, S ~string] struct {
	Name    S
	Cost    T
	Through WeightedNode[T, S]
}

// Implements data.ItemPQ[T]
func (self *WeightedNodeImpl[T, S]) GetValue() T {
	return self.Cost
}
func (self *WeightedNodeImpl[T, S]) GetKey() S {
	return self.Name
}
func (self *WeightedNodeImpl[T, S]) GetThrough() WeightedNode[T, S] {
	return self.Through
}
func (self *WeightedNodeImpl[T, S]) SetCost(value T) {
	self.Cost = value
}

func (self *WeightedNodeImpl[T, S]) SetThrough(node WeightedNode[T, S]) {
	self.Through = node
}
