package wgraph

type NodeWeightedImpl[T graphWeight, S ~string] struct {
	Name        S
	ValueOrCost T
	// Previous is for Dijkstra shortest path algorithm, or other sort programs.
	Previous NodeWeighted[T, S]
}

// Implements data.Valuer[T]
func (self *NodeWeightedImpl[T, S]) GetValue() T {
	return self.ValueOrCost
}
func (self *NodeWeightedImpl[T, S]) GetKey() S {
	return self.Name
}
func (self *NodeWeightedImpl[T, S]) GetPrevious() NodeWeighted[T, S] {
	return self.Previous
}
func (self *NodeWeightedImpl[T, S]) SetValueOrCost(value T) {
	self.ValueOrCost = value
}

func (self *NodeWeightedImpl[T, S]) SetPrevious(prev NodeWeighted[T, S]) {
	self.Previous = prev
}
