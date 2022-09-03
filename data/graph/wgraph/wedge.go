package wgraph

type WeightedEdgeImpl[T graphWeight, S ~string] struct {
	node   WeightedNode[T, S]
	weight T
}

func (self *WeightedEdgeImpl[T, S]) GetNode() WeightedNode[T, S] {
	return self.node
}

func (self *WeightedEdgeImpl[T, S]) GetWeight() T {
	return self.weight
}
