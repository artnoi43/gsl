package graph

import "github.com/artnoi43/gsl/data/container/list"

// The functions in this file is rewrites of their counterparts in bfs.go.
// They were modified to take in GenericGraph[Node[T], Node[T], any, any] instead.

func BFSNg[T comparable](
	g GenericGraph[Node[T], Node[T], any],
	src Node[T],
	dst Node[T],
) (
	[]Node[T],
	int,
	bool,
) {
	rawPath, found := BFSSearchNg(g, src, dst)
	if !found {
		return nil, -1, false
	}

	shortestPath, hops := BFSShortestPathReconstruct(rawPath, src, dst)
	return shortestPath, hops, found
}

func BFSSearchNg[T comparable](
	g GenericGraph[Node[T], Node[T], any],
	src Node[T],
	dst Node[T],
) (
	map[Node[T]]Node[T],
	bool,
) {
	q := list.NewSafeQueue[Node[T]]()
	q.Push(src)

	visited := make(map[Node[T]]bool)
	prev := make(map[Node[T]]Node[T])
	var found bool
	for !q.IsEmpty() {
		popped := q.Pop()
		if popped == nil {
			panic("popped nil - should not happen")
		}
		current := *popped

		neighbors := g.GetNodeEdges(current)
		for _, neighbor := range neighbors {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true

			if neighbor == dst {
				found = true
			}

			q.Push(neighbor)
			prev[neighbor] = current
		}
	}

	return prev, found
}
