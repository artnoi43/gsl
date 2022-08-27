package data

import (
	"container/heap"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

var ErrNegativeWeightEdge = errors.New("Djikstra edge must not be negative")

type djiskstraWeight interface{ constraints.Ordered }

// DjikstraGraph[T] wraps WeightedGraphImpl[T], where T is generic type numeric types.
// Only constraints.Unsigned T is being tested.
// To prevent bugs, so if the weight data source is of non-T type (i.e. a float), users will need to perform
// multiplications on the data, e.g. 0.1 -> 1000, 0.01 -> 100.
type DjikstraGraph[T djiskstraWeight] struct {
	graph WeightedGraph[T]
}

func NewDjikstraGraph[T djiskstraWeight]() *DjikstraGraph[T] {
	return &DjikstraGraph[T]{
		graph: NewWeightedGraph[T](),
	}
}

func (self *DjikstraGraph[T]) AddDjikstraEdge(n1, n2 *Node[T], weight T) error {
	var zeroValue T
	if weight < zeroValue {
		return errors.Wrapf(ErrNegativeWeightEdge, "negative edge weight %v", weight)
	}
	err := self.graph.AddEdge(n1, n2, weight)
	return err
}

func (self *DjikstraGraph[T]) DjikstraFrom(startNode *Node[T]) (shortestPaths map[*Node[T]][]*Node[T]) {
	var zeroValue T
	startNode.Cost = zeroValue
	startNode.Through = nil

	pq := NewPriorityQueue[T](MinHeap)
	heap.Push(pq, startNode)

	visited := make(map[*Node[T]]bool)
	for !pq.IsEmpty() {
		// Pop the top of pq and mark it as visited
		currentHeap := heap.Pop(pq)
		if currentHeap == nil {
			panic("popped nil - should not happen")
		}
		current, ok := currentHeap.(*Node[T])
		if !ok {
			typeOfCurrent := reflect.TypeOf(current)
			panic(fmt.Sprintf("current is %s, not *Node[T]", typeOfCurrent))
		}
		visited[current] = true

		Edges := self.graph.GetNodeEdges(current)
		for _, Edge := range Edges {
			// Skip visited
			if visited[Edge.Node] {
				continue
			}

			heap.Push(pq, Edge.Node)
			// If getting to Edge from current is cheaper that the Edge current cost state,
			// update it to pass via current instead
			if newCost := current.Cost + Edge.weight; newCost < Edge.Node.Cost {
				Edge.Node.Cost = newCost
				Edge.Node.Through = current
			}
		}
	}

	// Reconstruct path
	shortestPaths = make(map[*Node[T]][]*Node[T], len(self.graph.GetNodes()))
	for _, node := range self.graph.GetNodes() {
		var path []*Node[T]
		// Keep going back until start (nil Through)
		for via := node; via.Through != nil; via = via.Through {
			path = append(path, via)
		}
		lenPath := len(path)
		if lenPath == 0 {
			continue
		}
		// Reverse path to forwardPath
		for i, j := 0, lenPath-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		shortestPaths[node] = path
	}
	return shortestPaths
}
