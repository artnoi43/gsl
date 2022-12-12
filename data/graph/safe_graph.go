package graph

import (
	"sync"
)

// SafeGraph[N, E, R, W] wraps any BasicGraph[N, E, R, W] graphs.
type SafeGraph[N, E, R, W any] struct {
	mut   *sync.RWMutex
	Graph GenericGraph[N, E, R, W]
}

// WrapSafeGenericGraph[N, E, R, W] wraps BasicGraph[N, E, R, W]
// with *SafeGraph[N, E, R, W] to use mutex to avoid data races
func WrapSafeGenericGraph[N, E, R, W any](g GenericGraph[N, E, R, W]) *SafeGraph[N, E, R, W] {
	return &SafeGraph[N, E, R, W]{
		Graph: g,
		mut:   &sync.RWMutex{},
	}
}

func (g *SafeGraph[N, E, R, W]) SetDirection(value bool) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.SetDirection(value)
}

func (g *SafeGraph[N, E, R, W]) HasDirection() bool {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.HasDirection()
}

func (g *SafeGraph[N, E, R, W]) AddNode(node N) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.AddNode(node)
}

func (g *SafeGraph[N, E, R, W]) GetNodes() []N {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodes()
}

func (g *SafeGraph[N, E, R, W]) AddEdge(n1, n2 N, weight W) error {
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.Graph.AddEdge(n1, n2, weight)
}

func (g *SafeGraph[N, E, R, W]) GetEdges() R {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetEdges()
}

func (g *SafeGraph[N, E, R, W]) GetNodeEdges(node N) []E {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodeEdges(node)
}
