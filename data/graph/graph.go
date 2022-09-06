package graph

// GraphImpl is a basic implementation of Graph, not safe for concurrent code.
type GraphImpl[T any] struct {
	direction bool
	Nodes     []Node[T]
	Edges     map[Node[T]][]Node[T]
}

// NewGraphUnsafe[T] returns the default implementation of unweighted graph (*GraphImpl[T])
// without the mutex field. If your code is not concurrent, use this type, otherwise,
// consider calling NewGraph[T] instead.
func NewGraphUnsafe[T any](hasDirection bool) Graph[T] {
	return &GraphImpl[T]{
		direction: hasDirection,
		Edges:     make(map[Node[T]][]Node[T]),
	}
}

// NewGraph[T] returns the default implementation of unweighted graph (*GraphImpl[T])
// wrapped inside a of SafeGraph[N any, E any, M any, W any]
func NewGraph[T any](hasDirection bool) Graph[T] {
	return WrapSafeGraph[
		Node[T],
		Node[T],
		map[Node[T]][]Node[T],
		any,
	](
		NewGraphUnsafe[T](hasDirection),
	)
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
