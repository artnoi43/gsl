package wgraph

// TODO: more relaxed generic types

import (
	"github.com/artnoi43/gsl/data/graph"
)

// HashMapGraphWeightedImpl[T] is the default implementation of GraphWeighted[T].
// To use this type concurrently, consider wrapping it with graph.WrapSafeGraph.
type HashMapGraphWeightedImpl[T graphWeight, S ~string, N NodeWeighted[T, S]] struct {
	Directed bool
	Nodes    []N
	Edges    map[NodeWeighted[T, S]][]EdgeWeighted[T, S] // Edges is a map of a NodeWeightedImpl's edges
}

// NewGraphWeightedUnsafe[T, S] returns the default implementation of GraphWeighted[T, S] without the concurrency wrapper.
// If your code is concurrent, use NewGraphWeighted[T, S] instead
func NewGraphWeightedUnsafe[T graphWeight, S ~string, N NodeWeighted[T, S]](directed bool) GraphWeighted[T, S, N] {
	return &HashMapGraphWeightedImpl[T, S, N]{
		Directed: directed,
		Nodes:    make([]N, 0),
		Edges:    make(map[NodeWeighted[T, S]][]EdgeWeighted[T, S]),
	}
}

// WrapSafeGraphWeighted wraps any graph g that implements GraphWeighted[T, S] with graph.SafeGraph[N, E, M, W].
func WrapSafeGraphWeighted[T graphWeight, S ~string, N NodeWeighted[T, S]](
	g GraphWeighted[T, S, N],
) GraphWeighted[T, S, N] {
	// The type parameters mirror how GraphWeighted[T, S] implements BasicGraph[N, E, M, W]
	return graph.WrapSafeGenericGraph[
		N,
		EdgeWeighted[T, S],
		T,
	](
		g,
	)
}

// NewGraphWeighted[T, S] returns the default implementation of GraphWeighted[T, S] with the concurrency wrapper.
// If your code is not concurrent, use NewGraphWeightedUnsafe[T, S] instead
func NewGraphWeighted[T graphWeight, S ~string, N NodeWeighted[T, S]](directed bool) GraphWeighted[T, S, N] {
	g := &HashMapGraphWeightedImpl[T, S, N]{
		Directed: directed,
		Nodes:    make([]N, 0),
		Edges:    make(map[NodeWeighted[T, S]][]EdgeWeighted[T, S]),
	}

	return WrapSafeGraphWeighted[T, S, N](g)
}

func (s *HashMapGraphWeightedImpl[T, S, N]) SetDirection(directed bool) { s.Directed = directed }

func (s *HashMapGraphWeightedImpl[T, S, N]) IsDirected() bool { return s.Directed }

func (s *HashMapGraphWeightedImpl[T, S, N]) AddNode(node N) {
	s.Nodes = append(s.Nodes, node)
}

// AddEdge adds edge from n1 to n2. This particular method does not return error in any case.
func (s *HashMapGraphWeightedImpl[T, S, N]) AddEdge(n1, n2 N, weight T) error {
	// Add and edge from n1 leading to n2
	s.Edges[n1] = append(s.Edges[n1], &EdgeWeightedImpl[T, S]{toNode: n2, weight: weight})

	if s.Directed {
		return nil
	}

	// If it's not directed, then both nodes have links from and to each other
	s.Edges[n2] = append(s.Edges[n2], &EdgeWeightedImpl[T, S]{toNode: n1, weight: weight})
	return nil
}

func (s *HashMapGraphWeightedImpl[T, S, N]) GetNodes() []N { return s.Nodes }

func (s *HashMapGraphWeightedImpl[T, S, N]) GetEdges() []EdgeWeighted[T, S] {
	var edges []EdgeWeighted[T, S]
	for _, nodeEdges := range s.Edges {
		edges = append(edges, nodeEdges...)
	}

	return edges
}

func (s *HashMapGraphWeightedImpl[T, S, N]) GetNodeNeighbors(node N) []N {
	edges := s.Edges[node]
	neighbors := make([]N, len(edges))

	for i, edge := range edges {
		neighbors[i] = edge.ToNode().(N)
	}

	return neighbors
}

func (s *HashMapGraphWeightedImpl[T, S, N]) GetNodeEdges(node N) []EdgeWeighted[T, S] {
	return s.Edges[node]
}
