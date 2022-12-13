package wgraph

import "github.com/artnoi43/gsl/data/graph"

// NodeWeighted should be able to put in a priority queue, just in case topological sort is needed.
type NodeWeighted[T graphWeight, S ~string] interface {
	// Inherit some from unweighted graphs
	graph.Node[T]

	GetKey() S // Get the node's key, names, IDs
}

type NodeWeightedImpl[T graphWeight, S ~string] struct {
	Name        S
	ValueOrCost T
	// Previous is for Dijkstra shortest path algorithm, or other sort programs.
}

type NodeDijkstraImpl[T WeightDjikstra, S ~string] struct {
	NodeWeightedImpl[T, S]
	Previous NodeDijkstra[T, S]
}

// Implements data.Valuer[T]
func (n *NodeWeightedImpl[T, S]) GetValue() T {
	return n.ValueOrCost
}

func (n *NodeWeightedImpl[T, S]) GetKey() S {
	return n.Name
}

func (n *NodeWeightedImpl[T, S]) SetValueOrCost(value T) {
	n.ValueOrCost = value
}

func (n *NodeDijkstraImpl[T, S]) GetPrevious() NodeDijkstra[T, S] {
	return n.Previous
}

func (n *NodeDijkstraImpl[T, S]) SetPrevious(prev NodeDijkstra[T, S]) {
	n.Previous = prev
}
