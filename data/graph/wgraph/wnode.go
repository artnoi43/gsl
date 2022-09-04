package wgraph

type NodeWeightedImpl[T graphWeight, S ~string] struct {
	Name    S
	Cost    T
	Through WeightedNode[T, S]
}

// Implements data.ItemPQ[T]
func (self *NodeWeightedImpl[T, S]) GetValue() T {
	return self.Cost
}
func (self *NodeWeightedImpl[T, S]) GetKey() S {
	return self.Name
}
func (self *NodeWeightedImpl[T, S]) GetThrough() WeightedNode[T, S] {
	return self.Through
}
func (self *NodeWeightedImpl[T, S]) SetCost(value T) {
	self.Cost = value
}

func (self *NodeWeightedImpl[T, S]) SetThrough(node WeightedNode[T, S]) {
	self.Through = node
}
