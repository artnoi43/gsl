package graph

import "sync"

// GraphImpl is a basic implementation of Graph, not safe for concurrent code.
type GraphImpl[T any] struct {
	direction bool
	Nodes     []Node[T]
	Edges     map[Node[T]][]Node[T]
}

// GraphImplSafe[T] wraps GraphImpl[T] methods with locking and unlocking mutex.
type GraphImplSafe[T any] struct {
	Graph Graph[T]
	mut   *sync.RWMutex
}

// NewGraph[T] returns the default implementation of Graph[T], which is GraphImplSafe[T].
// If your code is not concurrent, you can use GraphImpl instead.
func NewGraph[T any](hasDirection bool) Graph[T] {
	return &GraphImplSafe[T]{
		Graph: &GraphImpl[T]{
			direction: hasDirection,
			Edges:     make(map[Node[T]][]Node[T]),
		},
		mut: &sync.RWMutex{},
	}
}

func (g *GraphImpl[T]) SetDirection(hasDirection bool) { g.direction = hasDirection }

func (g *GraphImpl[T]) HasDirection() bool { return g.direction }

func (g *GraphImpl[T]) GetNodes() []Node[T] { return g.Nodes }

func (g *GraphImpl[T]) GetEdges() map[Node[T]][]Node[T] { return g.Edges }

func (g *GraphImpl[T]) GetNodeEdges(node Node[T]) []Node[T] { return g.Edges[node] }

func (g *GraphImpl[T]) AddNode(node Node[T]) {
	g.Nodes = append(g.Nodes, node)
}

func (g *GraphImpl[T]) AddEdge(n1, n2 Node[T], weight any) error {
	if weight != nil {
		return ErrEdgeWeightNotNull
	}

	// Add and edge from n1 leading to n2
	g.Edges[n1] = append(g.Edges[n1], n2)

	if g.direction {
		return nil
	}
	// If it's not directed, then both nodes have links from and to each other
	g.Edges[n2] = append(g.Edges[n2], n1)
	return nil
}

func (g *GraphImplSafe[T]) SetDirection(value bool) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.SetDirection(value)
}

func (g *GraphImplSafe[T]) HasDirection() bool {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.HasDirection()
}

func (g *GraphImplSafe[T]) GetNodes() []Node[T] {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodes()
}

func (g *GraphImplSafe[T]) GetEdges() map[Node[T]][]Node[T] {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetEdges()
}

func (g *GraphImplSafe[T]) GetNodeEdges(node Node[T]) []Node[T] {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodeEdges(node)
}

func (g *GraphImplSafe[T]) AddNode(node Node[T]) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.AddNode(node)
}

func (g *GraphImplSafe[T]) AddEdge(n1, n2 Node[T], weight any) error {
	if weight != nil {
		return ErrEdgeWeightNotNull
	}
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.Graph.AddEdge(n1, n2, weight)
}
