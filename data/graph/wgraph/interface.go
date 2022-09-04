package wgraph

import (
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/mgl/data/graph"
	"github.com/artnoi43/mgl/data/list"
)

type graphWeight interface {
	constraints.Ordered
}

type WeightedGraph[T graphWeight, S ~string] interface {
	graph.GenericGraph[
		WeightedNode[T, S],
		WeightedEdge[T, S],
		T,
		map[WeightedNode[T, S]][]WeightedEdge[T, S],
	]
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
