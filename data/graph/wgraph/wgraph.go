package wgraph

// TODO: more relaxed generic types

import (
	"sync"
)

type GraphWeightedImpl[T graphWeight, S ~string] struct {
	direction bool
	Nodes     []NodeWeighted[T, S]
	Edges     map[NodeWeighted[T, S]][]EdgeWeighted[T, S] // Edges is a map of a NodeWeightedImpl's edges
}

// GraphWeightedImplSafe[T] wraps GraphWeightedImpl[T] methods with locking and unlocking mutex.
type GraphWeightedImplSafe[T graphWeight, S ~string] struct {
	Graph GraphWeighted[T, S]

	// TODO: should we make mut public/exported?
	mut *sync.RWMutex
}

func NewGraphWeighted[T graphWeight, S ~string](hasDirection bool) GraphWeighted[T, S] {
	return &GraphWeightedImplSafe[T, S]{
		Graph: &GraphWeightedImpl[T, S]{
			direction: hasDirection,
			Nodes:     []NodeWeighted[T, S]{},
			Edges:     make(map[NodeWeighted[T, S]][]EdgeWeighted[T, S]),
		},
		mut: &sync.RWMutex{},
	}
}

func (self *GraphWeightedImpl[T, S]) SetDirection(hasDirection bool) {
	self.direction = hasDirection
}

func (self *GraphWeightedImpl[T, S]) HasDirection() bool {
	return self.direction
}

func (self *GraphWeightedImpl[T, S]) GetNodes() []NodeWeighted[T, S] {
	return self.Nodes
}

func (self *GraphWeightedImpl[T, S]) GetEdges() map[NodeWeighted[T, S]][]EdgeWeighted[T, S] {
	return self.Edges
}

func (self *GraphWeightedImpl[T, S]) GetNodeEdges(node NodeWeighted[T, S]) []EdgeWeighted[T, S] {
	return self.Edges[node]
}

func (self *GraphWeightedImpl[T, S]) AddNode(node NodeWeighted[T, S]) {
	self.Nodes = append(self.Nodes, node)
}

func (self *GraphWeightedImplSafe[T, S]) AddNode(node NodeWeighted[T, S]) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.Graph.AddNode(node)
}

// AddEdge adds edge from n1 to n2. This particular method does not return error in any case.
func (self *GraphWeightedImpl[T, S]) AddEdge(n1, n2 NodeWeighted[T, S], weight T) error {
	// Add and edge from n1 leading to n2
	self.Edges[n1] = append(self.Edges[n1], &EdgeWeightedImpl[T, S]{node: n2, weight: weight})

	if self.direction {
		return nil
	}

	// If it's not directed, then both nodes have links from and to each other
	self.Edges[n2] = append(self.Edges[n2], &EdgeWeightedImpl[T, S]{node: n1, weight: weight})
	return nil
}

func (self *GraphWeightedImplSafe[T, S]) SetDirection(hasDirection bool) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.Graph.SetDirection(hasDirection)
}

func (self *GraphWeightedImplSafe[T, S]) HasDirection() bool {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Graph.HasDirection()
}

func (self *GraphWeightedImplSafe[T, S]) GetNodes() []NodeWeighted[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Graph.GetNodes()
}

func (self *GraphWeightedImplSafe[T, S]) GetEdges() map[NodeWeighted[T, S]][]EdgeWeighted[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Graph.GetEdges()
}

func (self *GraphWeightedImplSafe[T, S]) GetNodeEdges(node NodeWeighted[T, S]) []EdgeWeighted[T, S] {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Graph.GetNodeEdges(node)
}

// AddEdge adds edge from n1 to n2. This particular method does not return error in any case.
func (self *GraphWeightedImplSafe[T, S]) AddEdge(n1, n2 NodeWeighted[T, S], weight T) error {
	self.mut.Lock()
	defer self.mut.Unlock()

	return self.Graph.AddEdge(n1, n2, weight)
}
