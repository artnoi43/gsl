package wgraph

// TODO: more relaxed generic types

import (
	"sync"
)

type GraphWeightedImpl[T graphWeight, S ~string] struct {
	Direction bool
	Nodes     []WeightedNode[T, S]
	Edges     map[WeightedNode[T, S]][]WeightedEdge[T, S] // Edges is a map of a WeightedNodeImpl's edges
}

type GraphWeightedImplSafe[T graphWeight, S ~string] struct {
	graph *GraphWeightedImpl[T, S]

	// TODO: should we make mut public/exported?
	mut *sync.RWMutex
}

func NewWeightedGraph[T graphWeight, S ~string](hasDirection bool) WeightedGraph[T, S] {
	return &GraphWeightedImplSafe[T, S]{
		graph: &GraphWeightedImpl[T, S]{
			Direction: hasDirection,
			Nodes:     []WeightedNode[T, S]{},
			Edges:     make(map[WeightedNode[T, S]][]WeightedEdge[T, S]),
		},
		mut: &sync.RWMutex{},
	}
}

func (self *GraphWeightedImpl[T, S]) SetDirection(hasDirection bool) {
	self.Direction = hasDirection
}

func (self *GraphWeightedImpl[T, S]) HasDirection() bool {
	return self.Direction
}

func (self *GraphWeightedImpl[T, S]) GetNodes() []WeightedNode[T, S] {
	return self.Nodes
}

func (self *GraphWeightedImpl[T, S]) GetEdges() map[WeightedNode[T, S]][]WeightedEdge[T, S] {
	return self.Edges
}

func (self *GraphWeightedImpl[T, S]) GetNodeEdges(node WeightedNode[T, S]) []WeightedEdge[T, S] {
	return self.Edges[node]
}

func (self *GraphWeightedImpl[T, S]) AddNode(node WeightedNode[T, S]) {
	self.Nodes = append(self.Nodes, node)
}

func (self *GraphWeightedImplSafe[T, S]) AddNode(node WeightedNode[T, S]) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.graph.AddNode(node)
}

// AddEdge adds edge from n1 to n2. This particular method does not return error in any case.
func (self *GraphWeightedImpl[T, S]) AddEdge(n1, n2 WeightedNode[T, S], weight T) error {
	// Add and edge from n1 leading to n2
	self.Edges[n1] = append(self.Edges[n1], &EdgeWeightedImpl[T, S]{node: n2, weight: weight})

	if self.Direction {
		return nil
	}

	// If it's not directed, then both nodes have links from and to each other
	self.Edges[n2] = append(self.Edges[n2], &EdgeWeightedImpl[T, S]{node: n1, weight: weight})
	return nil
}

func (self *GraphWeightedImplSafe[T, S]) SetDirection(hasDirection bool) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.graph.SetDirection(hasDirection)
}

func (self *GraphWeightedImplSafe[T, S]) HasDirection() bool {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.graph.HasDirection()
}

func (self *GraphWeightedImplSafe[T, S]) GetNodes() []WeightedNode[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.graph.GetNodes()
}

func (self *GraphWeightedImplSafe[T, S]) GetEdges() map[WeightedNode[T, S]][]WeightedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.graph.GetEdges()
}

func (self *GraphWeightedImplSafe[T, S]) GetNodeEdges(node WeightedNode[T, S]) []WeightedEdge[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.graph.GetNodeEdges(node)
}

// AddEdge adds edge from n1 to n2. This particular method does not return error in any case.
func (self *GraphWeightedImplSafe[T, S]) AddEdge(n1, n2 WeightedNode[T, S], weight T) error {
	self.mut.Lock()
	defer self.mut.Unlock()

	return self.graph.AddEdge(n1, n2, weight)
}
