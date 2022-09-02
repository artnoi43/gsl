package wgraph

import (
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/mgl/data/list"
)

type graphWeight interface {
	constraints.Ordered
}

type UndirectedGraph[T graphWeight, S ~string] interface {
	AddNode(node UndirectedNode[T, S])
	AddEdge(n1, n2 UndirectedNode[T, S], weight T)
	GetNodes() []UndirectedNode[T, S]
	GetEdges() map[UndirectedNode[T, S]][]UndirectedEdge[T, S]
	GetNodeEdges(node UndirectedNode[T, S]) []UndirectedEdge[T, S]
}

type UndirectedNode[T graphWeight, S ~string] interface {
	list.ItemPQ[T]
	GetKey() S
	GetThrough() UndirectedNode[T, S]
	SetCost(T)
	SetThrough(UndirectedNode[T, S])
}

type UndirectedEdge[T graphWeight, S ~string] interface {
	GetNode() UndirectedNode[T, S]
	GetWeight() T
}
