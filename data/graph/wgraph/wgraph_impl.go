package wgraph

// TODO: more relaxed generic types

import (
	"github.com/artnoi43/gsl/data/graph"
)

// GraphWeightedImpl[T] is the default implementation of GraphWeighted[T].
// To use this type concurrently, consider wrapping it with graph.WrapSafeGraph.
type GraphWeightedImpl[T graphWeight, S ~string] struct {
	Directed bool
	Nodes    []NodeWeighted[T, S]
	Edges    map[NodeWeighted[T, S]][]EdgeWeighted[T, S] // Edges is a map of a NodeWeightedImpl's edges
}

// NewGraphWeightedUnsafe[T, S] returns the default implementation of GraphWeighted[T, S] without the concurrency wrapper.
// If your code is concurrent, use NewGraphWeighted[T, S] instead
func NewGraphWeightedUnsafe[T graphWeight, S ~string](directed bool) HashMapGraphWeighted[T, S] {
	return &GraphWeightedImpl[T, S]{
		Directed: directed,
		Nodes:    []NodeWeighted[T, S]{},
		Edges:    make(map[NodeWeighted[T, S]][]EdgeWeighted[T, S]),
	}
}

// WrapSafeGraphWeighted wraps any graph g that implements GraphWeighted[T, S] with graph.SafeGraph[N, E, M, W].
func WrapSafeGraphWeighted[T graphWeight, S ~string](g HashMapGraphWeighted[T, S]) HashMapGraphWeighted[T, S] {
	// The type parameters mirror how GraphWeighted[T, S] implements BasicGraph[N, E, M, W]
	return graph.WrapSafeGenericGraph[
		NodeWeighted[T, S],
		EdgeWeighted[T, S],
		map[NodeWeighted[T, S]][]EdgeWeighted[T, S],
		T,
	](
		g,
	)
}

// NewGraphWeighted[T, S] returns the default implementation of GraphWeighted[T, S] with the concurrency wrapper.
// If your code is not concurrent, use NewGraphWeightedUnsafe[T, S] instead
func NewGraphWeighted[T graphWeight, S ~string](directed bool) HashMapGraphWeighted[T, S] {
	return WrapSafeGraphWeighted(
		NewGraphWeightedUnsafe[T, S](directed),
	)
}

func (self *GraphWeightedImpl[T, S]) SetDirection(directed bool) { self.Directed = directed }

func (self *GraphWeightedImpl[T, S]) HasDirection() bool { return self.Directed }

func (self *GraphWeightedImpl[T, S]) GetNodes() []NodeWeighted[T, S] { return self.Nodes }

func (self *GraphWeightedImpl[T, S]) GetEdges() map[NodeWeighted[T, S]][]EdgeWeighted[T, S] {
	return self.Edges
}

func (self *GraphWeightedImpl[T, S]) GetNodeEdges(node NodeWeighted[T, S]) []EdgeWeighted[T, S] {
	return self.Edges[node]
}

func (self *GraphWeightedImpl[T, S]) AddNode(node NodeWeighted[T, S]) {
	self.Nodes = append(self.Nodes, node)
}

// AddEdge adds edge from n1 to n2. This particular method does not return error in any case.
func (self *GraphWeightedImpl[T, S]) AddEdge(n1, n2 NodeWeighted[T, S], weight T) error {
	// Add and edge from n1 leading to n2
	self.Edges[n1] = append(self.Edges[n1], &EdgeWeightedImpl[T, S]{toNode: n2, weight: weight})

	if self.Directed {
		return nil
	}

	// If it's not directed, then both nodes have links from and to each other
	self.Edges[n2] = append(self.Edges[n2], &EdgeWeightedImpl[T, S]{toNode: n1, weight: weight})
	return nil
}
