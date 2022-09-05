package wgraph

type EdgeWeightedImpl[T graphWeight, S ~string] struct {
	node   NodeWeighted[T, S]
	weight T
}

// If E is an edge from nodes A to B, then E.GetNode() returns B.
func (self *EdgeWeightedImpl[T, S]) GetNode() NodeWeighted[T, S] {
	return self.node
}

func (self *EdgeWeightedImpl[T, S]) GetWeight() T {
	return self.weight
}
