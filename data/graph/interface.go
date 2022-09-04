package graph

import "errors"

var ErrEdgeWeightNotNull = errors.New("found edge weight in unweighted graph")

// This package provides a ery simple interface for graph. Currently, there's no limitation on what a Node can be.
// A concurrency-ready example implementation is available as GraphImpl.
type Graph interface {
	SetDirection(bool)
	HasDirection() bool
	AddNode(node Node)
	// AddEdge for unweighted graph DOES NOT do anything with weight when adding an edge.
	// In the real implementation provided in this package, AddEdge returns an error if weight is not nil.
	AddEdge(n1, n2 Node, weight any) error
	GetNodes() []Node
	GetEdges() map[Node][]Node
	GetNodeEdges(node Node) []Node
}

type Node any
