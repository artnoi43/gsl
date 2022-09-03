package wgraph

// TODO: more relaxed generic types

import (
	"sync"
)

type WeightedGraphImpl[T graphWeight, S ~string] struct {
	Direction bool
	Nodes     []WeightedNode[T, S]
	Edges     map[WeightedNode[T, S]][]WeightedEdge[T, S] // Edges is a map of a WeightedNodeImpl's edges

	// TODO: should we make mut public/exported?
	mut *sync.RWMutex
}

func NewWeightedGraph[T graphWeight, S ~string](hasDirection bool) WeightedGraph[T, S] {
	return &WeightedGraphImpl[T, S]{
		Direction: hasDirection,
		Nodes:     []WeightedNode[T, S]{},
		Edges:     make(map[WeightedNode[T, S]][]WeightedEdge[T, S]),
		mut:       &sync.RWMutex{},
	}
}

func (self *WeightedGraphImpl[T, S]) SetDirection(hasDirection bool) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.Direction = hasDirection
}

func (self *WeightedGraphImpl[T, S]) HasDirection() bool {
	self.mut.Lock()
	defer self.mut.Unlock()

	return self.Direction
}

func (self *WeightedGraphImpl[T, S]) GetNodes() []WeightedNode[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Nodes
}

func (self *WeightedGraphImpl[T, S]) GetEdges() map[WeightedNode[T, S]][]WeightedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Edges
}

func (self *WeightedGraphImpl[T, S]) GetNodeEdges(node WeightedNode[T, S]) []WeightedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Edges[node]
}

func (self *WeightedGraphImpl[T, S]) AddNode(node WeightedNode[T, S]) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.Nodes = append(self.Nodes, node)
}

func (self *WeightedGraphImpl[T, S]) AddEdge(n1, n2 WeightedNode[T, S], weight T) {
	self.mut.Lock()
	defer self.mut.Unlock()

	// Add and edge from n1 leading to n2
	self.Edges[n1] = append(self.Edges[n1], &WeightedEdgeImpl[T, S]{node: n2, weight: weight})

	if self.Direction {
		return
	}
	// If it's not directed, then both nodes have links from and to each other
	self.Edges[n2] = append(self.Edges[n2], &WeightedEdgeImpl[T, S]{node: n1, weight: weight})
}
