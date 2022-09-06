package graph

import "sync"

type SafeGraph[N any, E any, M any, W any] struct {
	Graph GenericGraph[N, E, M, W]
	mut   *sync.RWMutex
}

// WrapSafeGenericGraph[N, E, M, W] wraps BasicGraph[N, E, M, W]
// with *SafeGraph[N, E, M, W] to use mutex to avoid data races
func WrapSafeGenericGraph[
	N any,
	E any,
	M any,
	W any,
](g GenericGraph[N, E, M, W]) *SafeGraph[N, E, M, W] {
	return &SafeGraph[N, E, M, W]{
		Graph: g,
		mut:   &sync.RWMutex{},
	}
}

func (g *SafeGraph[N, E, M, W]) SetDirection(value bool) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.SetDirection(value)
}

func (g *SafeGraph[N, E, M, W]) HasDirection() bool {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.HasDirection()
}

func (g *SafeGraph[N, E, M, W]) AddNode(node N) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.Graph.AddNode(node)
}

func (g *SafeGraph[N, E, M, W]) GetNodes() []N {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodes()
}

func (g *SafeGraph[N, E, M, W]) AddEdge(n1, n2 N, weight W) error {
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.Graph.AddEdge(n1, n2, weight)
}

func (g *SafeGraph[N, E, M, W]) GetEdges() M {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetEdges()
}

func (g *SafeGraph[N, E, M, W]) GetNodeEdges(node N) []E {
	g.mut.RLock()
	defer g.mut.RUnlock()

	return g.Graph.GetNodeEdges(node)
}
