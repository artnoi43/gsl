package wgraph

import (
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/gsl/data/graph"
)

type graphWeight interface {
	constraints.Ordered
}

// HashMapGraphWeighted is a graph.GenericGraph, using a hash map to represent node connections.
type HashMapGraphWeighted[
	T graphWeight,
	S ~string,
] graph.GenericGraph[
	// The weighted graph's node interface
	NodeWeighted[T, S],
	// The weighted graph's edge interface
	EdgeWeighted[T, S],
	// The edge weight can be of any types that implements graphWeight
	T,
]

// NodeWeighted should be able to put in a priority queue, just in case topological sort is needed.
type NodeWeighted[T graphWeight, S ~string] interface {
	// Inherit some from unweighted graphs
	graph.Node[T]

	// Other node with weighted edge methods

	SetValueOrCost(value T)              // Save cost or value to the node
	GetKey() S                           // Get the node's key, names, IDs
	GetPrevious() NodeWeighted[T, S]     // When using with Dijkstra code, gets the previous (prior node) from a Dijkstra walk.
	SetPrevious(prev NodeWeighted[T, S]) // In Dijkstra code, set a node's previous node value
}

// EdgeWeighted represents what a weighted edge should be able to do.
type EdgeWeighted[T graphWeight, S ~string] interface {
	ToNode() (to NodeWeighted[T, S])
	GetWeight() T
}
