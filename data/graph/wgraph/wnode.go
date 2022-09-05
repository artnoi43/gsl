package wgraph

type NodeWeightedImpl[T graphWeight, S ~string] struct {
	Name        S
	ValueOrCost T
	// Through is for Dijkstra shortest path algorithm.
	Through NodeWeighted[T, S]
}

// Implements data.Valuer[T]
func (self *NodeWeightedImpl[T, S]) GetValue() T {
	return self.ValueOrCost
}
func (self *NodeWeightedImpl[T, S]) GetKey() S {
	return self.Name
}
func (self *NodeWeightedImpl[T, S]) GetThrough() NodeWeighted[T, S] {
	return self.Through
}
func (self *NodeWeightedImpl[T, S]) SetValueOrCost(value T) {
	self.ValueOrCost = value
}

func (self *NodeWeightedImpl[T, S]) SetThrough(node NodeWeighted[T, S]) {
	self.Through = node
}
