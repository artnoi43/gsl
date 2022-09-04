package graph

// This package provides a very simple interface for graph.
// A concurrency-ready example implementation is available as GraphImpl.

import (
	"errors"

	"github.com/artnoi43/mgl/data"
)

var ErrEdgeWeightNotNull = errors.New("found edge weight in unweighted graph")

// GenericGraph represents what an mgl graph should look like.
// It is not intended to be used by production code, but more like an internal building block for mgl graphs.
// It's not minimal, and was designed around flexibility and coverage.
// This interface can be used for both unweighted and weighted graph (see wgraph package).
type GenericGraph[
	N any, // Type for graph node
	E any, // Type for graph edge
	M any, // Type representing the edge implementation of the graph, typically map[N]E.
	W any, // Type for graph weight
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

// Graph is a GenericGraph, but with more node constraints.
type Graph[T nodeValue] GenericGraph[
	// The graph node is Node[T]
	Node[T],
	// Since there's no edge weight, this graph will use the connected nodes to represent a node's edges
	Node[T],
	// The graph's implements edges as a map of node to other nodes
	map[Node[T]][]Node[T],
	// The weight can be of any types, but only nil is valid if using the default implementation of unweighted graph
	any,
]

type Node[T nodeValue] interface {
	data.Valuer[T]
}
