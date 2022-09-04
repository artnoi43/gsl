package graph

// This package provides a very simple interface for graph.
// A concurrency-ready example implementation is available as GraphImpl.

import (
	"errors"

	"github.com/artnoi43/mgl/data"
)

var ErrEdgeWeightNotNull = errors.New("found edge weight in unweighted graph")

type Directioner interface {
}

type GenericGraph[
	N any, // Type for graph node
	E any, // Type for graph edge
	W any, // Type for graph weight
	M any, // Type representing the edge implementation of the graph.
] interface {
	// SetDirection sets the directionality of the graph.
	SetDirection(bool)
	// HasDirection
	HasDirection() bool
	AddNode(node N)
	GetNodes() []N
	// AddEdge adds an edge to the graph. If the graph is directional, then AddEdge will only adds edge from n1 to n2.
	// If the graph is undirectional, then both n1->n2 and n2->n1 edges are added.
	AddEdge(n1, n2 N, weight W) error
	// GetNodeEdge takes in a node, and returns the connections from that node. If the graph is unweighted,
	// the so-called "edges" are just a list of nodes. If the graph is weighted, then the edges returned are
	// slice of the actual edges.
	GetNodeEdges(node N) []E
	// GetEdges returns all the edges in the graph.
	// In the actual implementations provided by mgl, the type M differs between unweighted and weighted graphs.
	// If the graph is weighted, then the returned type M is usually a map of node to its edges,
	// otherwise if the graph has unweighted edges, It returns a slice of nodes accessible by Node.
	GetEdges() M
}

type nodeValue any

// The Graph interface was modeled after wgraph.WeightedGraph, to try to make both interfaces
// very similar, in case I'll be good at Go enough to compose a single interface.
type Graph[T nodeValue] interface {
	GenericGraph[
		Node[T],
		Node[T],
		any,
		map[Node[T]][]Node[T],
	]
}

type Node[T nodeValue] interface {
	data.Valuer[T]
}
