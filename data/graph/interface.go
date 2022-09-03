package graph

// This package provides a ery simple interface for graph. Currently, there's no limitation on what a Node can be.
// A concurrency-ready example implementation is available as GraphImpl.
type Graph interface {
	SetDirection(bool)
	HasDirection() bool
	AddNode(node Node)
	AddEdge(n1, n2 Node)
	GetNodes() []Node
	GetEdges() map[Node][]Node
	GetNodeEdges(node Node) []Node
}

type Node any
