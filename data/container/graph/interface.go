package graph

// This package provides a very simple interface for graph.
// A concurrency-ready example implementation is available as GraphImpl.

import (
	"errors"

	"github.com/artnoi43/gsl/data"
)

var ErrEdgeWeightNotNull = errors.New("found edge weight in unweighted graph")

type nodeValue any

// Graph is a GenericGraph, but with more node constraints. Type `GraphImpl` implements this interface.
type Graph[T nodeValue] GenericGraph[
	// The graph node is Node[T]
	Node[T],
	// Since there's no edge weight, this graph will use the connected nodes to represent a node's edges
	Node[T],
	// The graph's implements edges as a map of node to other nodes
	map[Node[T]][]Node[T],
	// The weight can be of any types, BUT ONLY NIL IS VALID if using the default implementation of unweighted graph
	any,
]

type Node[T nodeValue] data.Valuer[T]
