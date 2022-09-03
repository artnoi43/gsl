package wgraph

import (
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/mgl/data/list"
)

type graphWeight interface {
	constraints.Ordered
}

type WeightedGraph[T graphWeight, S ~string] interface {
	SetDirection(bool)
	HasDirection() bool
	AddNode(node WeightedNode[T, S])
	AddEdge(n1, n2 WeightedNode[T, S], weight T)
	GetNodes() []WeightedNode[T, S]
	GetEdges() map[WeightedNode[T, S]][]WeightedEdge[T, S]
	GetNodeEdges(node WeightedNode[T, S]) []WeightedEdge[T, S]
}

type WeightedNode[T graphWeight, S ~string] interface {
	list.ItemPQ[T] // For priority queue
	SetCost(T)     // Set accumulated cost to the node

	GetKey() S
	GetThrough() WeightedNode[T, S]
	SetThrough(WeightedNode[T, S])
}

type WeightedEdge[T graphWeight, S ~string] interface {
	GetNode() WeightedNode[T, S]
	GetWeight() T
}
