package data

// TODO: more relaxed generic types

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type graphWeight interface {
	constraints.Ordered
}

// See WeightedGraphImpl
type WeightedGraph[T graphWeight] interface {
	AddNode(node *Node[T])
	AddEdge(n1, n2 *Node[T], weight T)
	GetNodes() []*Node[T]
	GetEdges() map[*Node[T]][]*Edge[T]
	GetNodeEdges(node *Node[T]) []*Edge[T]
}

type Node[T graphWeight] struct {
	Name    string
	Cost    T
	Through *Node[T]
}

// Implements data.ItemPQ[T]
func (self *Node[T]) Value() T {
	return self.Cost
}

type Edge[T graphWeight] struct {
	Node   *Node[T]
	weight T
}

type WeightedGraphImpl[T graphWeight] struct {
	Nodes []*Node[T]
	Edges map[*Node[T]][]*Edge[T] // Edges is a map of a Node's edges

	// TODO: should we make mut public/exported?
	mut *sync.RWMutex
}

func NewWeightedGraph[T graphWeight]() *WeightedGraphImpl[T] {
	return &WeightedGraphImpl[T]{
		Edges: make(map[*Node[T]][]*Edge[T]),
		mut:   &sync.RWMutex{},
	}
}

func (self *WeightedGraphImpl[T]) GetNodes() []*Node[T] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Nodes
}

func (self *WeightedGraphImpl[T]) GetEdges() map[*Node[T]][]*Edge[T] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Edges
}

func (self *WeightedGraphImpl[T]) GetNodeEdges(node *Node[T]) []*Edge[T] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Edges[node]
}

func (self *WeightedGraphImpl[T]) AddNode(node *Node[T]) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Nodes = append(self.Nodes, node)
}

func (self *WeightedGraphImpl[T]) AddEdge(n1, n2 *Node[T], weight T) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Edges[n1] = append(self.Edges[n1], &Edge[T]{Node: n2, weight: weight})
	self.Edges[n2] = append(self.Edges[n2], &Edge[T]{Node: n1, weight: weight})
}
