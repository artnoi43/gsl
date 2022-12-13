package wgraph

// EdgeWeighted represents what a weighted edge should be able to do.
type EdgeWeighted[T graphWeight, S ~string] interface {
	ToNode() (to NodeWeighted[T, S])
	GetWeight() T
}

type EdgeWeightedImpl[T graphWeight, S ~string] struct {
	toNode NodeWeighted[T, S]
	weight T
}

// If E is an edge from nodes A to B, then E.GetToNode() returns B.
func (s *EdgeWeightedImpl[T, S]) ToNode() NodeWeighted[T, S] {
	return s.toNode
}

func (s *EdgeWeightedImpl[T, S]) GetWeight() T {
	return s.weight
}
