package graph

import "sync"

type GraphImpl struct {
	Direction bool
	Nodes     []Node
	Edges     map[Node][]Node

	mut sync.RWMutex
}

func NewGraph(hasDirection bool) Graph {
	return &GraphImpl{
		Direction: hasDirection,
		Edges:     make(map[Node][]Node),
	}
}

func (g *GraphImpl) SetDirection(hasDirection bool) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Direction = hasDirection
}

func (g *GraphImpl) HasDirection() bool {
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.Direction
}

func (g *GraphImpl) GetNodes() []Node {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Nodes
}

func (g *GraphImpl) GetEdges() map[Node][]Node {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Edges
}

func (g *GraphImpl) GetNodeEdges(node Node) []Node {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Edges[node]
}

func (g *GraphImpl) AddNode(node Node) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Nodes = append(g.Nodes, node)
}

func (g *GraphImpl) AddEdge(n1, n2 Node) {
	g.mut.Lock()
	defer g.mut.Unlock()

	// Add and edge from n1 leading to n2
	g.Edges[n1] = append(g.Edges[n1], n2)

	if g.Direction {
		return
	}
	// If it's not directed, then both nodes have links from and to each other
	g.Edges[n2] = append(g.Edges[n2], n1)
}
