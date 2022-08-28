package graph

// TODO: more relaxed generic types

import (
	"sync"

	"github.com/artnoi43/mgl/data/list"
	"golang.org/x/exp/constraints"
)

type graphWeight interface {
	constraints.Ordered
}

// See UndirectedWeightedGraphImpl
type UndirectedWeightedGraph[T graphWeight, S ~string] interface {
	AddNode(node UndirectedWeightedNode[T, S])
	AddEdge(n1, n2 UndirectedWeightedNode[T, S], weight T)
	GetNodes() []UndirectedWeightedNode[T, S]
	GetEdges() map[UndirectedWeightedNode[T, S]][]*UndirectedWeightedEdge[T, S]
	GetNodeEdges(node UndirectedWeightedNode[T, S]) []*UndirectedWeightedEdge[T, S]
}

type UndirectedWeightedNode[T graphWeight, S ~string] interface {
	list.ItemPQ[T]
	GetKey() S
	GetThrough() UndirectedWeightedNode[T, S]
	SetValue(T)
	SetThrough(UndirectedWeightedNode[T, S])
}

type UndirectedWeightedNodeImpl[T graphWeight, S ~string] struct {
	Name    S
	Cost    T
	Through UndirectedWeightedNode[T, S]
}

// Implements data.ItemPQ[T]
func (self *UndirectedWeightedNodeImpl[T, S]) GetValue() T {
	return self.Cost
}
func (self *UndirectedWeightedNodeImpl[T, S]) GetKey() S {
	return self.Name
}
func (self *UndirectedWeightedNodeImpl[T, S]) GetThrough() UndirectedWeightedNode[T, S] {
	return self.Through
}
func (self *UndirectedWeightedNodeImpl[T, S]) SetValue(value T) {
	self.Cost = value
}

func (self *UndirectedWeightedNodeImpl[T, S]) SetThrough(node UndirectedWeightedNode[T, S]) {
	self.Through = node
}

type UndirectedWeightedEdge[T graphWeight, S ~string] struct {
	Node   UndirectedWeightedNode[T, S]
	weight T
}

type UndirectedWeightedGraphImpl[T graphWeight, S ~string] struct {
	Nodes []UndirectedWeightedNode[T, S]
	Edges map[UndirectedWeightedNode[T, S]][]*UndirectedWeightedEdge[T, S] // Edges is a map of a UndirectedWeightedNodeImpl's edges

	// TODO: should we make mut public/exported?
	mut *sync.RWMutex
}

func NewWeightedGraph[T graphWeight, S ~string]() UndirectedWeightedGraph[T, S] {
	return &UndirectedWeightedGraphImpl[T, S]{
		Nodes: []UndirectedWeightedNode[T, S]{},
		Edges: make(map[UndirectedWeightedNode[T, S]][]*UndirectedWeightedEdge[T, S]),
		mut:   &sync.RWMutex{},
	}
}

func (self *UndirectedWeightedGraphImpl[T, S]) GetNodes() []UndirectedWeightedNode[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Nodes
}

func (self *UndirectedWeightedGraphImpl[T, S]) GetEdges() map[UndirectedWeightedNode[T, S]][]*UndirectedWeightedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Edges
}

func (self *UndirectedWeightedGraphImpl[T, S]) GetNodeEdges(node UndirectedWeightedNode[T, S]) []*UndirectedWeightedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Edges[node]
}

func (self *UndirectedWeightedGraphImpl[T, S]) AddNode(node UndirectedWeightedNode[T, S]) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Nodes = append(self.Nodes, node)
}

func (self *UndirectedWeightedGraphImpl[T, S]) AddEdge(n1, n2 UndirectedWeightedNode[T, S], weight T) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Edges[n1] = append(self.Edges[n1], &UndirectedWeightedEdge[T, S]{Node: n2, weight: weight})
	self.Edges[n2] = append(self.Edges[n2], &UndirectedWeightedEdge[T, S]{Node: n1, weight: weight})
}
