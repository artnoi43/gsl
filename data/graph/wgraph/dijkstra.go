package wgraph

import (
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

var ErrDijkstraNegativeWeightEdge = errors.New("Dijkstra edge must not be negative")

// WeightDijkstra represents the allowed types for edge weights and node costs.
// TODO: integrate more comparable types, like big.Int and big.Float
type WeightDijkstra interface {
	Weight
	constraints.Integer | constraints.Float
}

type NodeDijkstra[T WeightDijkstra] interface {
	NodeWeighted[T]

	GetPrevious() NodeDijkstra[T]     // When using with Dijkstra code, gets the previous (prior node) from a Dijkstra walk.
	SetPrevious(prev NodeDijkstra[T]) // In Dijkstra code, set a node's previous node value
}

type GraphDijkstra[T WeightDijkstra] interface {
	GraphWeighted[NodeDijkstra[T], EdgeWeighted[T, NodeDijkstra[T]], T]
	DijkstraShortestPathFrom(startNode NodeDijkstra[T]) *DijstraShortestPath[T]
}

// This type is the Dijkstra shortest path answer. It has 2 fields, (1) `From` the 'from' node, and (2) `Paths`.
// DijkstraShortestPath.Paths is a hash map where the key is a node, and the value is the previous node with the lowest cost to that key node.
// Because each instance holds all best route to every reachable node from From node, you can reconstruct the shortest path from any nodes in
// that Paths map with ReconstructPathTo
type DijstraShortestPath[T WeightDijkstra] struct {
	From  NodeDijkstra[T]
	Paths map[NodeDijkstra[T]]NodeDijkstra[T]
}

// NewDikstraGraph calls NewGraphWeightedUnsafe[T], and return the wrapped graph.
// Alternatively, you can create your own implementation of GraphWeighted[T].
func NewDijkstraGraphUnsafe[T WeightDijkstra](directed bool) GraphDijkstra[T] {
	return &GraphDijkstraImpl[T]{
		graph: &HashMapGraphWeightedImpl[T, NodeDijkstra[T]]{},
	}
}

// NewDikstraGraph calls NewGraphWeighted[T], and return the wrapped graph.
// Alternatively, you can create your own implementation of GraphWeighted[T].
func NewDijkstraGraph[T WeightDijkstra](directed bool) GraphDijkstra[T] {
	return &GraphDijkstraImpl[T]{
		graph: NewGraphWeighted[T, NodeDijkstra[T]](directed),
	}
}

// DijkstraShortestPathReconstruct reconstructs a path as an array of nodes
// from dst back until it found nil, that is, the first node after the start node.
// For example, if you have a shortestPaths map lile this:
/*
	dubai: nil
	helsinki: dubai
	budapest: helsinki
*/
// Then, the returned slice will be [budapest, helsinki, dubai],
// and the returned length will be 3 (inclusive). The path reconstruct from this function
// starts from the destination and goes all the way back to the source.
func DijkstraShortestPathReconstruct[T WeightDijkstra](
	shortestPaths map[NodeDijkstra[T]]NodeDijkstra[T],
	src NodeDijkstra[T],
	dst NodeDijkstra[T],
) []NodeDijkstra[T] {
	prevNodes := []NodeDijkstra[T]{dst}
	prev, found := shortestPaths[dst]
	if !found {
		return prevNodes
	}
	prevNodes = append(prevNodes, prev)

	for prev.GetPrevious() != nil {
		prevPrev, found := shortestPaths[prev]
		if !found {
			continue
		}

		prevNodes = append(prevNodes, prevPrev)
		prev = prevPrev

		// This allows us to have partial path
		if prev == src {
			break
		}
	}
	return prevNodes
}
