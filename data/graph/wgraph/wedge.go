package wgraph

type EdgeWeightedImpl[T graphWeight, S ~string] struct {
	toNode NodeWeighted[T, S]
	weight T
}

// If E is an edge from nodes A to B, then E.GetToNode() returns B.
func (self *EdgeWeightedImpl[T, S]) ToNode() NodeWeighted[T, S] {
	return self.toNode
}

func (self *EdgeWeightedImpl[T, S]) GetWeight() T {
	return self.weight
}
