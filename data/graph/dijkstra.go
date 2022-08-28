package graph

import (
	"container/heap"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/mgl/data/list"
)

var ErrDijkstraNegativeWeightEdge = errors.New("Dijkstra edge must not be negative")

type dijkstraWeight interface{ constraints.Ordered }

// DijkstraGraph[T] wraps WeightedGraphImpl[T], where T is generic type numeric types.
// Only constraints.Unsigned T is being tested.
// To prevent bugs, so if the weight data source is of non-T type (i.e. a float), users will need to perform
// multiplications on the data, e.g. 0.1 -> 1000, 0.01 -> 100.
type DijkstraGraph[T dijkstraWeight] struct {
	graph WeightedGraph[T]
}

func NewDijkstraGraph[T dijkstraWeight]() *DijkstraGraph[T] {
	return &DijkstraGraph[T]{
		graph: NewWeightedGraph[T](),
	}
}

func (self *DijkstraGraph[T]) AddDijkstraNode(node *Node[T]) {
	self.graph.AddNode(node)
}

// AddDijkstraEdge validates if weight is valid, and then calls WeightedGraph.AddEdge
func (self *DijkstraGraph[T]) AddDijkstraEdge(n1, n2 *Node[T], weight T) error {
	var zeroValue T
	if weight < zeroValue {
		return errors.Wrapf(ErrDijkstraNegativeWeightEdge, "negative edge weight %v", weight)
	}

	self.graph.AddEdge(n1, n2, weight)
	return nil
}

// DjisktraFrom takes a *Node[T] startNode, and finds the shortest path from startNode to all other nodes.
// This implementation uses PriorityQueue[T], so the nodes' values must satisfy constraints.Ordered.
// It returns a hash map, where the key is the destination node, and the values are all other previous nodes
// between the destination (map key) and startNode.
func (self *DijkstraGraph[T]) DijkstraFrom(startNode *Node[T]) (shortestPaths map[*Node[T]][]*Node[T]) {
	var zeroValue T
	startNode.Cost = zeroValue
	startNode.Through = nil

	pq := list.NewPriorityQueue[T](list.MinHeap)
	heap.Push(pq, startNode)

	visited := make(map[*Node[T]]bool)
	for !pq.IsEmpty() {
		// Pop the top of pq and mark it as visited
		popped := heap.Pop(pq)
		if popped == nil {
			panic("popped nil - should not happen")
		}
		current, ok := popped.(*Node[T])
		if !ok {
			typeOfCurrent := reflect.TypeOf(current)
			panic(fmt.Sprintf("current is %s, not *Node[T]", typeOfCurrent))
		}
		visited[current] = true

		edges := self.graph.GetNodeEdges(current)
		for _, edge := range edges {
			// Skip visited
			if visited[edge.Node] {
				continue
			}

			heap.Push(pq, edge.Node)
			// If getting to edge from current is cheaper that the edge current cost state,
			// update it to pass via current instead
			if newCost := current.Cost + edge.weight; newCost < edge.Node.Cost {
				edge.Node.Cost = newCost
				edge.Node.Through = current
			}
		}
	}

	// Reconstruct path
	shortestPaths = make(map[*Node[T]][]*Node[T], len(self.graph.GetNodes()))
	for _, node := range self.graph.GetNodes() {
		var path []*Node[T]
		// Keep going back until start (nil Through)
		// i.e. backward path
		for via := node; via.Through != nil; via = via.Through {
			path = append(path, via)
		}
		lenPath := len(path)
		if lenPath == 0 {
			continue
		}
		// Reverse path, i.e. to forward path
		for i, j := 0, lenPath-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		shortestPaths[node] = path
	}
	return shortestPaths
}
