package wgraph

type EdgeWeightedImpl[T graphWeight, S ~string] struct {
	node   WeightedNode[T, S]
	weight T
}

func (self *EdgeWeightedImpl[T, S]) GetNode() WeightedNode[T, S] {
	return self.node
}

func (self *EdgeWeightedImpl[T, S]) GetWeight() T {
	return self.weight
}
