package graph

// This package provides a very simple interface for graph.
// A concurrency-ready example implementation is available as GraphImpl.

import (
	"errors"

	"github.com/artnoi43/mgl/data"
)

var ErrEdgeWeightNotNull = errors.New("found edge weight in unweighted graph")

type nodeValue any

// The Graph interface was modeled after wgraph.WeightedGraph, to try to make both interfaces
// very similar, in case I'll be good at Go enough to compose a single interface.
type Graph[T nodeValue] interface {
	SetDirection(bool)
	HasDirection() bool
	AddNode(node Node[T])
	// AddEdge for unweighted graph DOES NOT do anything with weight when adding an edge.
	// In the real implementation provided in this package, AddEdge returns an error if weight is not nil.
	AddEdge(n1, n2 Node[T], weight any) error
	GetNodes() []Node[T]
	GetEdges() map[Node[T]][]Node[T]
	GetNodeEdges(node Node[T]) []Node[T]
}

type Node[T nodeValue] data.Valuer[T]
