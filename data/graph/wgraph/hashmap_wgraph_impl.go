package wgraph

// TODO: more relaxed generic types

import (
	"github.com/artnoi43/gsl/data/graph"
)

// HashMapGraphWeightedImpl[W, N] is the default implementation of GraphWeighted[N, EdgeWeighted[W], W].
type HashMapGraphWeightedImpl[T Weight, N NodeWeighted[T]] struct {
	Directed bool
	Nodes    []N
	Edges    map[NodeWeighted[T]][]EdgeWeighted[T, N] // Edges is a map of a NodeWeightedImpl's edges
}

// NewGraphWeightedUnsafe[T] returns the default implementation of GraphWeighted[T] without the concurrency wrapper.
// If your code is concurrent, use NewGraphWeighted[T] instead
func NewGraphWeightedUnsafe[T Weight, N NodeWeighted[T]](directed bool) GraphWeighted[N, EdgeWeighted[T, N], T] {
	return &HashMapGraphWeightedImpl[T, N]{
		Directed: directed,
		Nodes:    make([]N, 0),
		Edges:    make(map[NodeWeighted[T]][]EdgeWeighted[T, N]),
	}
}

// WrapSafeGraphWeighted wraps any graph g that implements GraphWeighted[T] with graph.SafeGraph[N, E, M, W].
func WrapSafeGraphWeighted[T Weight, N NodeWeighted[T]](
	g GraphWeighted[N, EdgeWeighted[T, N], T],
) GraphWeighted[N, EdgeWeighted[T, N], T] {
	// The type parameters mirror how GraphWeighted[T] implements BasicGraph[N, E, M, W]
	return graph.WrapSafeGenericGraph[
		N,
		EdgeWeighted[T, N],
		T,
	](
		g,
	)
}

// NewGraphWeighted[T] returns the default implementation of GraphWeighted[T] with the concurrency wrapper.
// If your code is not concurrent, use NewGraphWeightedUnsafe[T] instead
func NewGraphWeighted[T Weight, N NodeWeighted[T]](directed bool) GraphWeighted[N, EdgeWeighted[T, N], T] {
	return WrapSafeGraphWeighted(NewGraphWeightedUnsafe[T, N](directed))
}

func (g *HashMapGraphWeightedImpl[T, N]) SetDirection(directed bool) {
	g.Directed = directed
}

func (g *HashMapGraphWeightedImpl[T, N]) IsDirected() bool {
	return g.Directed
}

func (g *HashMapGraphWeightedImpl[T, N]) AddNode(node N) {
	g.Nodes = append(g.Nodes, node)
}

// AddEdge adds edge from n1 to n2. This particular method does not return error in any case.
func (g *HashMapGraphWeightedImpl[T, N]) AddEdge(n1, n2 N, weight T) error {
	// Add and edge from n1 leading to n2
	g.Edges[n1] = append(g.Edges[n1], &EdgeWeightedImpl[T, N]{toNode: n2, weight: weight})

	if g.Directed {
		return nil
	}

	// If it's not directed, then both nodes have links from and to each other
	g.Edges[n2] = append(g.Edges[n2], &EdgeWeightedImpl[T, N]{toNode: n1, weight: weight})
	return nil
}

func (g *HashMapGraphWeightedImpl[T, N]) GetNodes() []N {
	return g.Nodes
}

func (g *HashMapGraphWeightedImpl[T, N]) GetEdges() []EdgeWeighted[T, N] {
	var edges []EdgeWeighted[T, N]
	for _, nodeEdges := range g.Edges {
		edges = append(edges, nodeEdges...)
	}

	return edges
}

func (g *HashMapGraphWeightedImpl[T, N]) GetNodeNeighbors(node N) []N {
	edges := g.Edges[node]
	neighbors := make([]N, len(edges))

	for i, edge := range edges {
		neighbors[i] = edge.ToNode()
	}

	return neighbors
}

func (g *HashMapGraphWeightedImpl[T, N]) GetNodeEdges(node N) []EdgeWeighted[T, N] {
	return g.Edges[node]
}
