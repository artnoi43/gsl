package wgraph

import (
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/gsl/data/graph"
)

type graphWeight interface {
	constraints.Ordered
}

// GraphWeighted is a graph.GenericGraph, using a hash map to represent node connections.
type GraphWeighted[
	T graphWeight,
	S ~string,
	N NodeWeighted[T, S],
] graph.Graph[
	// The weighted graph's node interface
	N,
	// The weighted graph's edge interface
	EdgeWeighted[T, S],
	// The edge weight can be of any types that implements graphWeight
	T,
]
