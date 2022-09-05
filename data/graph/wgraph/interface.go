package wgraph

import (
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/gsl/data"
	"github.com/artnoi43/gsl/data/graph"
)

type graphWeight interface {
	constraints.Ordered
}

// WeightedGraph is a graph.GenericGraph
type WeightedGraph[
	T graphWeight,
	S ~string,
] graph.GenericGraph[
	// The graph's node type
	WeightedNode[T, S],
	// The graph's edge type
	WeightedEdge[T, S],
	// The graph represents the edges as map of node to edges
	map[WeightedNode[T, S]][]WeightedEdge[T, S],
	// The edge weight can be of any types that implements graphWeight
	T,
]

// WeightedNode should be able to put in a priority queue, just in case topological sort is needed.
type WeightedNode[T graphWeight, S ~string] interface {
	data.Valuer[T] // For priority queue
	SetCost(T)     // Set accumulated cost to the node

	GetKey() S
	GetThrough() WeightedNode[T, S]
	SetThrough(WeightedNode[T, S])
}

// WeightedEdge represents what a weighted edge should be able to do.
type WeightedEdge[T graphWeight, S ~string] interface {
	GetNode() WeightedNode[T, S]
	GetWeight() T
}
