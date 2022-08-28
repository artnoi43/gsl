package wgraph

// TODO: more relaxed generic types

import (
	"sync"
)

type UndirectedGraphImpl[T graphWeight, S ~string] struct {
	Nodes []UndirectedNode[T, S]
	Edges map[UndirectedNode[T, S]][]UndirectedEdge[T, S] // Edges is a map of a UndirectedNodeImpl's edges

	// TODO: should we make mut public/exported?
	mut *sync.RWMutex
}

func NewWeightedGraph[T graphWeight, S ~string]() UndirectedGraph[T, S] {
	return &UndirectedGraphImpl[T, S]{
		Nodes: []UndirectedNode[T, S]{},
		Edges: make(map[UndirectedNode[T, S]][]UndirectedEdge[T, S]),
		mut:   &sync.RWMutex{},
	}
}

func (self *UndirectedGraphImpl[T, S]) GetNodes() []UndirectedNode[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Nodes
}

func (self *UndirectedGraphImpl[T, S]) GetEdges() map[UndirectedNode[T, S]][]UndirectedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Edges
}

func (self *UndirectedGraphImpl[T, S]) GetNodeEdges(node UndirectedNode[T, S]) []UndirectedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()
	return self.Edges[node]
}

func (self *UndirectedGraphImpl[T, S]) AddNode(node UndirectedNode[T, S]) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Nodes = append(self.Nodes, node)
}

func (self *UndirectedGraphImpl[T, S]) AddEdge(n1, n2 UndirectedNode[T, S], weight T) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Edges[n1] = append(self.Edges[n1], &UndirectedEdgeImpl[T, S]{node: n2, weight: weight})
	self.Edges[n2] = append(self.Edges[n2], &UndirectedEdgeImpl[T, S]{node: n1, weight: weight})
}
