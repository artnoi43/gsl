package wgraph

import (
	"container/heap"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/gsl/data/list"
)

var ErrDijkstraNegativeWeightEdge = errors.New("Dijkstra edge must not be negative")

// TODO: integrate more comparable types, like big.Int and big.Float
// dijkstraWeight represents the allowed types for edge weights and node costs.
type dijkstraWeight interface {
	constraints.Integer | constraints.Float
}

// This type is the Dijkstra shortest path answer. It has 2 fields, (1) `From` the 'from' node, and (2) `Paths`.
// DijkstraShortestPath.Paths is a hash map where the key is a node, and the value is the previous node with the lowest cost to that key node.
type DijstraShortestPath[T dijkstraWeight, S ~string] struct {
	From  NodeWeighted[T, S]
	Paths map[NodeWeighted[T, S]]NodeWeighted[T, S]
}

// DijkstraGraph[T] wraps GraphWeightedImpl[T], where T is generic type numeric types and S is ~string.
// Only constraints.Unsigned T is being tested.
// To prevent bugs, so if the weight data source is of non-T type (i.e. a float), users will need to perform
// multiplications on the data, e.g. 0.1 -> 1000, 0.01 -> 100.
type DijkstraGraph[T dijkstraWeight, S ~string] struct {
	graph GraphWeighted[T, S]
}

// NewDikstraGraph calls NewGraphWeighted[T, S], and return the graph.
// Alternatively, you can create your own implementation of GraphWeighted
func NewDijkstraGraph[T dijkstraWeight, S ~string](hasDirection bool) *DijkstraGraph[T, S] {
	return &DijkstraGraph[T, S]{
		graph: NewGraphWeighted[T, S](hasDirection),
	}
}

func (self *DijkstraGraph[T, S]) SetDirection(hasDirection bool) {
	self.graph.SetDirection(hasDirection)
}

func (self *DijkstraGraph[T, S]) HasDirection() bool {
	return self.graph.HasDirection()
}

func (self *DijkstraGraph[T, S]) AddNode(node NodeWeighted[T, S]) {
	self.graph.AddNode(node)
}

// AddDijkstraEdge validates if weight is valid, and then calls GraphWeighted.AddEdge
func (self *DijkstraGraph[T, S]) AddEdge(n1, n2 NodeWeighted[T, S], weight T) error {
	var zeroValue T
	if weight < zeroValue {
		return errors.Wrapf(ErrDijkstraNegativeWeightEdge, "negative edge weight %v", weight)
	}

	self.graph.AddEdge(n1, n2, weight)
	return nil
}

func (self *DijkstraGraph[T, S]) GetNodes() []NodeWeighted[T, S] {
	return self.graph.GetNodes()
}

func (self *DijkstraGraph[T, S]) GetEdges() map[NodeWeighted[T, S]][]EdgeWeighted[T, S] {
	return self.graph.GetEdges()
}

func (self *DijkstraGraph[T, S]) GetNodeEdges(node NodeWeighted[T, S]) []EdgeWeighted[T, S] {
	return self.graph.GetNodeEdges(node)
}

// DjisktraFrom takes a *NodeImpl[T] startNode, and finds the shortest path from startNode to all other nodes.
// This implementation uses PriorityQueue[T], so the nodes' values must satisfy constraints.Ordered.
func (self *DijkstraGraph[T, S]) DijkstraShortestPathFrom(startNode NodeWeighted[T, S]) *DijstraShortestPath[T, S] {
	var zeroValue T
	startNode.SetValueOrCost(zeroValue)
	startNode.SetThrough(nil)

	visited := make(map[NodeWeighted[T, S]]bool)
	prev := make(map[NodeWeighted[T, S]]NodeWeighted[T, S])

	pq := list.NewPriorityQueue[T](list.MinHeap)
	heap.Push(pq, startNode)

	for !pq.IsEmpty() {
		// Pop the top of pq and mark it as visited
		popped := heap.Pop(pq)
		if popped == nil {
			panic("popped nil - should not happen")
		}
		current, ok := popped.(NodeWeighted[T, S])
		if !ok {
			typeOfCurrent := reflect.TypeOf(current)
			panic(fmt.Sprintf("current is %s, not *Node[T]", typeOfCurrent))
		}

		visited[current] = true
		edges := self.GetNodeEdges(current)

		for _, edge := range edges {
			edgeNode := edge.GetNode()

			// Skip visited
			if visited[edgeNode] {
				continue
			}

			heap.Push(pq, edgeNode)
			// If getting to edge from current is cheaper that the edge current cost state,
			// update it to pass via current instead
			if newCost := current.GetValue() + edge.GetWeight(); newCost < edgeNode.GetValue() {
				edgeNode.SetValueOrCost(newCost)
				edgeNode.SetThrough(current)
				// Save path answer to prev
				prev[edgeNode] = current
			}
		}
	}

	return &DijstraShortestPath[T, S]{
		From:  startNode,
		Paths: prev,
	}
}

// ReconstructPathTo reconstructs shortest path as slice of nodes, from `d.from` to `to`
// using DijkstraShortestPathReconstruct
func (d *DijstraShortestPath[T, S]) ReconstructPathTo(to NodeWeighted[T, S]) []NodeWeighted[T, S] {
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
func DijkstraShortestPathReconstruct[T dijkstraWeight, S ~string](
	shortestPaths map[NodeWeighted[T, S]]NodeWeighted[T, S],
	src NodeWeighted[T, S],
	dst NodeWeighted[T, S],
) []NodeWeighted[T, S] {
	prevNodes := []NodeWeighted[T, S]{dst}
	prev, found := shortestPaths[dst]
	if !found {
		return prevNodes
	}
	prevNodes = append(prevNodes, prev)

	for prev.GetThrough() != nil {
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
