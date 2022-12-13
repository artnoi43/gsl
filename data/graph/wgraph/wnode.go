package wgraph

import "github.com/artnoi43/gsl/data/graph"

// NodeWeighted should be able to put in a priority queue, just in case topological sort is needed.
type NodeWeighted[T graphWeight, S ~string] interface {
	// Inherit some from unweighted graphs
	graph.Node[T]

	// Other node with weighted edge methods

	SetValueOrCost(value T)              // Save cost or value to the node
	GetKey() S                           // Get the node's key, names, IDs
	GetPrevious() NodeWeighted[T, S]     // When using with Dijkstra code, gets the previous (prior node) from a Dijkstra walk.
	SetPrevious(prev NodeWeighted[T, S]) // In Dijkstra code, set a node's previous node value
}

type NodeWeightedImpl[T graphWeight, S ~string] struct {
	Name        S
	ValueOrCost T
	// Previous is for Dijkstra shortest path algorithm, or other sort programs.
	Previous NodeWeighted[T, S]
}

// Implements data.Valuer[T]
func (n *NodeWeightedImpl[T, S]) GetValue() T {
	return n.ValueOrCost
}
func (n *NodeWeightedImpl[T, S]) GetKey() S {
	return n.Name
}
func (n *NodeWeightedImpl[T, S]) GetPrevious() NodeWeighted[T, S] {
	return n.Previous
}
func (n *NodeWeightedImpl[T, S]) SetValueOrCost(value T) {
	n.ValueOrCost = value
}

func (n *NodeWeightedImpl[T, S]) SetPrevious(prev NodeWeighted[T, S]) {
	n.Previous = prev
}
