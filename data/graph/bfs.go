package graph

import (
	"github.com/artnoi43/mgl/data/list"
)

// BFSShortestPath calls BFSSearch and uses its output to call BFSShortestPathReconstruct.
// It then returns the shortest path (slice of nodes), the number of hops it takes from `src` to `dst`,
// and the boolean value whethere there's a path from `src` to `dst`
func BFSShortestPath(g Graph, src Node, dst Node) ([]Node, int, bool) {
	rawPath, found := BFSSearch(g, src, dst)
	if !found {
		return nil, -1, false
	}

	shortestPath, hops := BFSShortestPathReconstruct(rawPath, src, dst)
	return shortestPath, hops, found
}

func BFSSearch(g Graph, src Node, dst Node) (map[Node]Node, bool) {
	q := list.NewQueue[Node]()
	q.Push(src)

	visited := make(map[Node]bool)
	prev := make(map[Node]Node)
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

func BFSShortestPathReconstruct(backwardPath map[Node]Node, src, dst Node) ([]Node, int) {
	current := dst
	shortestPath := []Node{current}
	var hops int
	if current == src {
		return shortestPath, hops
	}

	for {
		if current == src {
			break
		}

		next, found := backwardPath[current]
		if !found {
			break
		}

		shortestPath = append(shortestPath, next)
		current = next
		hops++
	}

	return shortestPath, hops
}
