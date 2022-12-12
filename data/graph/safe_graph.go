package graph

import (
	"sync"
)

// SafeGraph[N, E, R, W] wraps any BasicGraph[N, E, R, W] graphs.
type SafeGraph[N, E, W any] struct {
	mut   *sync.RWMutex
	Graph GenericGraph[N, E, W]
}

// WrapSafeGenericGraph[N, E, R, W] wraps BasicGraph[N, E, R, W]
// with *SafeGraph[N, E, R, W] to use mutex to avoid data races
func WrapSafeGenericGraph[N, E, W any](g GenericGraph[N, E, W]) *SafeGraph[N, E, W] {
	return &SafeGraph[N, E, W]{
		Graph: g,
		mut:   &sync.RWMutex{},
	}
}

func (g *SafeGraph[N, E, W]) SetDirection(value bool) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.SetDirection(value)
}

func (g *SafeGraph[N, E, W]) IsDirected() bool {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.IsDirected()
}

func (g *SafeGraph[N, E, W]) AddNode(node N) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.AddNode(node)
}

func (g *SafeGraph[N, E, W]) GetNodes() []N {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodes()
}

func (g *SafeGraph[N, E, W]) AddEdge(n1, n2 N, weight W) error {
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.Graph.AddEdge(n1, n2, weight)
}

func (g *SafeGraph[N, E, W]) GetNodeEdges(node N) []E {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodeEdges(node)
}
