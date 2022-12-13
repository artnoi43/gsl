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

// TODO: integrate more comparable types, like big.Int and big.Float
// WeightDjikstra represents the allowed types for edge weights and node costs.
type WeightDjikstra interface {
	constraints.Integer | constraints.Float
}

type NodeDijkstra[T WeightDjikstra, S ~string] interface {
	NodeWeighted[T, S]

	SetValueOrCost(value T)              // Save cost or value to the node
	GetPrevious() NodeDijkstra[T, S]     // When using with Dijkstra code, gets the previous (prior node) from a Dijkstra walk.
	SetPrevious(prev NodeDijkstra[T, S]) // In Dijkstra code, set a node's previous node value
}

// GraphDijkstra[T] wraps GraphWeightedImpl[T], where T is generic type numeric types and S is ~string.
// It uses HashMapGraphWeighted as the underlying data structure.
type GraphDijkstra[T WeightDjikstra, S ~string] struct {
	graph GraphWeighted[T, S, NodeDijkstra[T, S]]
}

// This type is the Dijkstra shortest path answer. It has 2 fields, (1) `From` the 'from' node, and (2) `Paths`.
// DijkstraShortestPath.Paths is a hash map where the key is a node, and the value is the previous node with the lowest cost to that key node.
// Because each instance holds all best route to every reachable node from From node, you can reconstruct the shortest path from any nodes in
// that Paths map with ReconstructPathTo
type DijstraShortestPath[T WeightDjikstra, S ~string] struct {
	From  NodeDijkstra[T, S]
	Paths map[NodeDijkstra[T, S]]NodeDijkstra[T, S]
}

// NewDikstraGraph calls NewGraphWeightedUnsafe[T, S], and return the wrapped graph.
// Alternatively, you can create your own implementation of GraphWeighted[T, S].
func NewDijkstraGraphUnsafe[T WeightDjikstra, S ~string](directed bool) *GraphDijkstra[T, S] {
	return &GraphDijkstra[T, S]{
		graph: &HashMapGraphWeightedImpl[T, S, NodeDijkstra[T, S]]{},
	}
}

// NewDikstraGraph calls NewGraphWeighted[T, S], and return the wrapped graph.
// Alternatively, you can create your own implementation of GraphWeighted[T, S].
func NewDijkstraGraph[T WeightDjikstra, S ~string](directed bool) *GraphDijkstra[T, S] {
	return &GraphDijkstra[T, S]{
		graph: NewGraphWeighted[T, S, NodeDijkstra[T, S]](directed),
	}
}

func (g *GraphDijkstra[T, S]) SetDirection(directed bool) {
	g.graph.SetDirection(directed)
}

func (g *GraphDijkstra[T, S]) IsDirected() bool {
	return g.graph.IsDirected()
}

func (g *GraphDijkstra[T, S]) AddNode(node NodeDijkstra[T, S]) {
	g.graph.AddNode(node)
}

// AddDijkstraEdge validates if weight is valid, and then calls GraphWeighted.AddEdge
func (g *GraphDijkstra[T, S]) AddEdge(n1, n2 NodeDijkstra[T, S], weight T) error {
	var zeroValue T
	if weight < zeroValue {
		return errors.Wrapf(ErrDijkstraNegativeWeightEdge, "negative edge weight %v", weight)
	}

	g.graph.AddEdge(n1, n2, weight)
	return nil
}

func (g *GraphDijkstra[T, S]) GetNodes() []NodeDijkstra[T, S] {
	nodes := any(g.graph.GetNodes())
	return nodes.([]NodeDijkstra[T, S])
}

func (g *GraphDijkstra[T, S]) GetEdges() []EdgeWeighted[T, S] {
	return g.graph.GetEdges()
}

func (g *GraphDijkstra[T, S]) GetNodeNeighbors(node NodeDijkstra[T, S]) []NodeDijkstra[T, S] {
	return g.graph.GetNodeNeighbors(node)
}

func (g *GraphDijkstra[T, S]) GetNodeEdges(node NodeDijkstra[T, S]) []EdgeWeighted[T, S] {
	return g.graph.GetNodeEdges(node)
}

// DjisktraFrom takes a *NodeImpl[T] startNode, and finds the shortest path from startNode to all other nodes.
// This implementation uses PriorityQueue[T], so the nodes' values must satisfy constraints.Ordered.
func (g *GraphDijkstra[T, S]) DijkstraShortestPathFrom(startNode NodeDijkstra[T, S]) *DijstraShortestPath[T, S] {
	var zeroValue T
	startNode.SetValueOrCost(zeroValue)
	startNode.SetPrevious(nil)

	visited := make(map[NodeDijkstra[T, S]]bool)
	parents := make(map[NodeDijkstra[T, S]]NodeDijkstra[T, S])

	pq := list.NewPriorityQueue[T](list.MinHeap)
	heap.Push(pq, startNode)

	for !pq.IsEmpty() {
		// Pop the top of pq and mark it as visited
		popped := heap.Pop(pq)
		if popped == nil {
			panic("popped nil - should not happen")
		}
		current, ok := popped.(NodeDijkstra[T, S])
		if !ok {
			typeOfCurrent := reflect.TypeOf(current)
			panic(fmt.Sprintf("current is %s, not *Node[T]", typeOfCurrent))
		}

		visited[current] = true
		edges := g.GetNodeEdges(current)

		for _, edge := range edges {
			edgeNode := edge.ToNode().(NodeDijkstra[T, S])

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

	return &DijstraShortestPath[T, S]{
		From:  startNode,
		Paths: parents,
	}
}

// ReconstructPathTo reconstructs shortest path as slice of nodes, from `d.from` to `to`
// using DijkstraShortestPathReconstruct
func (d *DijstraShortestPath[T, S]) ReconstructPathTo(to NodeDijkstra[T, S]) []NodeDijkstra[T, S] {
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
func DijkstraShortestPathReconstruct[T WeightDjikstra, S ~string](
	shortestPaths map[NodeDijkstra[T, S]]NodeDijkstra[T, S],
	src NodeDijkstra[T, S],
	dst NodeDijkstra[T, S],
) []NodeDijkstra[T, S] {
	prevNodes := []NodeDijkstra[T, S]{dst}
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
