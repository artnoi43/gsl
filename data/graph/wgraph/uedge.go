package wgraph

type UndirectedEdgeImpl[T graphWeight, S ~string] struct {
	node   UndirectedNode[T, S]
	weight T
}

func (self *UndirectedEdgeImpl[T, S]) GetNode() UndirectedNode[T, S] {
	return self.node
}

func (self *UndirectedEdgeImpl[T, S]) GetWeight() T {
	return self.weight
}
