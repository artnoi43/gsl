package graph

import "sync"

type GraphImpl[T any] struct {
	Direction bool
	Nodes     []Node[T]
	Edges     map[Node[T]][]Node[T]

	mut sync.RWMutex
}

func NewGraph[T any](hasDirection bool) Graph[T] {
	return &GraphImpl[T]{
		Direction: hasDirection,
		Edges:     make(map[Node[T]][]Node[T]),
	}
}

func (g *GraphImpl[T]) SetDirection(hasDirection bool) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Direction = hasDirection
}

func (g *GraphImpl[T]) HasDirection() bool {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Direction
}

func (g *GraphImpl[T]) GetNodes() []Node[T] {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Nodes
}

func (g *GraphImpl[T]) GetEdges() map[Node[T]][]Node[T] {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Edges
}

func (g *GraphImpl[T]) GetNodeEdges(node Node[T]) []Node[T] {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Edges[node]
}

func (g *GraphImpl[T]) AddNode(node Node[T]) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Nodes = append(g.Nodes, node)
}

func (g *GraphImpl[T]) AddEdge(n1, n2 Node[T], weight any) error {
	if weight != nil {
		return ErrEdgeWeightNotNull
	}
	g.mut.Lock()
	defer g.mut.Unlock()

	// Add and edge from n1 leading to n2
	g.Edges[n1] = append(g.Edges[n1], n2)

	if g.Direction {
		return nil
	}
	// If it's not directed, then both nodes have links from and to each other
	g.Edges[n2] = append(g.Edges[n2], n1)
	return nil
}
