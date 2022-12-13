package wgraph

import (
	"container/heap"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/gsl/data/container/list"
)

var ErrDijkstraNegativeWeightEdge = errors.New("Dijkstra edge must not be negative")

// WeightDjikstra represents the allowed types for edge weights and node costs.
// TODO: integrate more comparable types, like big.Int and big.Float
type WeightDjikstra interface {
	Weight
	constraints.Integer | constraints.Float
}

type NodeDijkstra[T WeightDjikstra] interface {
	NodeWeighted[T]

	GetPrevious() NodeDijkstra[T]     // When using with Dijkstra code, gets the previous (prior node) from a Dijkstra walk.
	SetPrevious(prev NodeDijkstra[T]) // In Dijkstra code, set a node's previous node value
}

// GraphDijkstra[T] wraps GraphWeightedImpl[T], where T is generic type numeric types and S is ~string.
// It uses HashMapGraphWeighted as the underlying data structure.
type GraphDijkstra[T WeightDjikstra] struct {
	graph GraphWeighted[NodeDijkstra[T], EdgeWeighted[T, NodeDijkstra[T]], T]
}

// This type is the Dijkstra shortest path answer. It has 2 fields, (1) `From` the 'from' node, and (2) `Paths`.
// DijkstraShortestPath.Paths is a hash map where the key is a node, and the value is the previous node with the lowest cost to that key node.
// Because each instance holds all best route to every reachable node from From node, you can reconstruct the shortest path from any nodes in
// that Paths map with ReconstructPathTo
type DijstraShortestPath[T WeightDjikstra] struct {
	From  NodeDijkstra[T]
	Paths map[NodeDijkstra[T]]NodeDijkstra[T]
}

// NewDikstraGraph calls NewGraphWeightedUnsafe[T], and return the wrapped graph.
// Alternatively, you can create your own implementation of GraphWeighted[T].
func NewDijkstraGraphUnsafe[T WeightDjikstra](directed bool) *GraphDijkstra[T] {
	return &GraphDijkstra[T]{
		graph: &HashMapGraphWeightedImpl[T, NodeDijkstra[T]]{},
	}
}

// NewDikstraGraph calls NewGraphWeighted[T], and return the wrapped graph.
// Alternatively, you can create your own implementation of GraphWeighted[T].
func NewDijkstraGraph[T WeightDjikstra](directed bool) *GraphDijkstra[T] {
	return &GraphDijkstra[T]{
		graph: NewGraphWeighted[T, NodeDijkstra[T]](directed),
	}
}

func (g *GraphDijkstra[T]) SetDirection(directed bool) {
	g.graph.SetDirection(directed)
}

func (g *GraphDijkstra[T]) IsDirected() bool {
	return g.graph.IsDirected()
}

func (g *GraphDijkstra[T]) AddNode(node NodeDijkstra[T]) {
	g.graph.AddNode(node)
}

// AddDijkstraEdge validates if weight is valid, and then calls GraphWeighted.AddEdge
func (g *GraphDijkstra[T]) AddEdge(n1, n2 NodeDijkstra[T], weight T) error {
	var zeroValue T
	if weight < zeroValue {
		return errors.Wrapf(ErrDijkstraNegativeWeightEdge, "negative edge weight %v", weight)
	}

	g.graph.AddEdge(n1, n2, weight)
	return nil
}

func (g *GraphDijkstra[T]) GetNodes() []NodeDijkstra[T] {
	nodes := any(g.graph.GetNodes())
	return nodes.([]NodeDijkstra[T])
}

func (g *GraphDijkstra[T]) GetEdges() []EdgeWeighted[T, NodeDijkstra[T]] {
	return g.graph.GetEdges()
}

func (g *GraphDijkstra[T]) GetNodeNeighbors(node NodeDijkstra[T]) []NodeDijkstra[T] {
	return g.graph.GetNodeNeighbors(node)
}

func (g *GraphDijkstra[T]) GetNodeEdges(node NodeDijkstra[T]) []EdgeWeighted[T, NodeDijkstra[T]] {
	return g.graph.GetNodeEdges(node)
}

// DjisktraFrom takes a *NodeImpl[T] startNode, and finds the shortest path from startNode to all other nodes.
// This implementation uses PriorityQueue[T], so the nodes' values must satisfy constraints.Ordered.
func (g *GraphDijkstra[T]) DijkstraShortestPathFrom(startNode NodeDijkstra[T]) *DijstraShortestPath[T] {
	var zeroValue T
	startNode.SetValueOrCost(zeroValue)
	startNode.SetPrevious(nil)

	visited := make(map[NodeDijkstra[T]]bool)
	parents := make(map[NodeDijkstra[T]]NodeDijkstra[T])

	pq := list.NewPriorityQueue[T](list.MinHeap)
	heap.Push(pq, startNode)

	for !pq.IsEmpty() {
		// Pop the top of pq and mark it as visited
		popped := heap.Pop(pq)
		if popped == nil {
			panic("popped nil - should not happen")
		}
		current, ok := popped.(NodeDijkstra[T])
		if !ok {
			typeOfCurrent := reflect.TypeOf(current)
			panic(fmt.Sprintf("current is %s, not *Node[T]", typeOfCurrent))
		}

		visited[current] = true
		edges := g.GetNodeEdges(current)

		for _, edge := range edges {
			edgeNode := edge.ToNode()

			// Skip visited
			if visited[edgeNode] {
				continue
			}

			heap.Push(pq, edgeNode)
			// If getting to edge from current is cheaper that the edge current cost state,
			// update it to pass via current instead
			if newCost := current.GetValue() + edge.GetWeight(); newCost < edgeNode.GetValue() {
				edgeNode.SetValueOrCost(newCost)
				edgeNode.SetPrevious(current)
				// Save (best) path answer to parents
				parents[edgeNode] = current
			}
		}
	}

	return &DijstraShortestPath[T]{
		From:  startNode,
		Paths: parents,
	}
}

// ReconstructPathTo reconstructs shortest path as slice of nodes, from `d.from` to `to`
// using DijkstraShortestPathReconstruct
func (d *DijstraShortestPath[T]) ReconstructPathTo(to NodeDijkstra[T]) []NodeDijkstra[T] {
	return DijkstraShortestPathReconstruct(d.Paths, d.From, to)
}

// DijkstraShortestPathReconstruct reconstructs a path as an array of nodes
// from dst back until it found nil, that is, the first node after the start node.
// For example, if you have a shortestPaths map lile this:
/*
	dubai: nil
	helsinki: dubai
	budapest: helsinki
*/
// Then, the returned slice will be [budapest, helsinki, dubai],
// and the returned length will be 3 (inclusive). The path reconstruct from this function
// starts from the destination and goes all the way back to the source.
func DijkstraShortestPathReconstruct[T WeightDjikstra](
	shortestPaths map[NodeDijkstra[T]]NodeDijkstra[T],
	src NodeDijkstra[T],
	dst NodeDijkstra[T],
) []NodeDijkstra[T] {
	prevNodes := []NodeDijkstra[T]{dst}
	prev, found := shortestPaths[dst]
	if !found {
		return prevNodes
	}
	prevNodes = append(prevNodes, prev)

	for prev.GetPrevious() != nil {
		prevPrev, found := shortestPaths[prev]
		if !found {
			continue
		}

		prevNodes = append(prevNodes, prevPrev)
		prev = prevPrev

		// This allows us to have partial path
		if prev == src {
			break
		}
	}
	return prevNodes
}
