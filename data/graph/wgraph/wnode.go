package wgraph

type NodeWeightedImpl[T graphWeight, S ~string] struct {
	Name        S
	ValueOrCost T
	// Previous is for Dijkstra shortest path algorithm, or other sort programs.
	Previous NodeWeighted[T, S]
}

// Implements data.Valuer[T]
func (n *NodeWeightedImpl[T, S]) GetValue() T {
	return n.ValueOrCost
}
func (n *NodeWeightedImpl[T, S]) GetKey() S {
	return n.Name
}
func (n *NodeWeightedImpl[T, S]) GetPrevious() NodeWeighted[T, S] {
	return n.Previous
}
func (n *NodeWeightedImpl[T, S]) SetValueOrCost(value T) {
	n.ValueOrCost = value
}

func (n *NodeWeightedImpl[T, S]) SetPrevious(prev NodeWeighted[T, S]) {
	n.Previous = prev
}
