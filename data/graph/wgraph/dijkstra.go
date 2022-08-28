package wgraph

import (
	"container/heap"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"

	"github.com/artnoi43/mgl/data/list"
)

var ErrDijkstraNegativeWeightEdge = errors.New("Dijkstra edge must not be negative")

type dijkstraWeight interface {
	constraints.Integer | constraints.Float
}

// DijkstraGraphUndirected[T] wraps WeightedGraphImpl[T], where T is generic type numeric types and S is ~string.
// Only constraints.Unsigned T is being tested.
// To prevent bugs, so if the weight data source is of non-T type (i.e. a float), users will need to perform
// multiplications on the data, e.g. 0.1 -> 1000, 0.01 -> 100.
type DijkstraGraphUndirected[T dijkstraWeight, S ~string] struct {
	graph UndirectedGraph[T, S]
}

func NewDijkstraGraph[T dijkstraWeight, S ~string]() *DijkstraGraphUndirected[T, S] {
	return &DijkstraGraphUndirected[T, S]{
		graph: NewWeightedGraph[T, S](),
	}
}

func (self *DijkstraGraphUndirected[T, S]) AddNode(node UndirectedNode[T, S]) {
	self.graph.AddNode(node)
}

// AddDijkstraEdge validates if weight is valid, and then calls WeightedGraph.AddEdge
func (self *DijkstraGraphUndirected[T, S]) AddEdge(n1, n2 UndirectedNode[T, S], weight T) error {
	var zeroValue T
	if weight < zeroValue {
		return errors.Wrapf(ErrDijkstraNegativeWeightEdge, "negative edge weight %v", weight)
	}

	self.graph.AddEdge(n1, n2, weight)
	return nil
}

func (self *DijkstraGraphUndirected[T, S]) GetNodes() []UndirectedNode[T, S] {
	return self.graph.GetNodes()
}

func (self *DijkstraGraphUndirected[T, S]) GetEdges() map[UndirectedNode[T, S]][]UndirectedEdge[T, S] {
	return self.graph.GetEdges()
}

func (self *DijkstraGraphUndirected[T, S]) GetNodeEdges(node UndirectedNode[T, S]) []UndirectedEdge[T, S] {
	return self.graph.GetNodeEdges(node)
}

// DjisktraFrom takes a *NodeImpl[T] startNode, and finds the shortest path from startNode to all other nodes.
// This implementation uses PriorityQueue[T], so the nodes' values must satisfy constraints.Ordered.
// It returns a hash map, where the key is the destination node, and the values are all other previous nodes
// between the destination (map key) and startNode.
func (self *DijkstraGraphUndirected[T, S]) DijkstraFrom(startNode UndirectedNode[T, S]) (shortestPaths map[UndirectedNode[T, S]][]UndirectedNode[T, S]) {
	var zeroValue T
	startNode.SetValue(zeroValue)
	startNode.SetThrough(nil)

	pq := list.NewPriorityQueue[T](list.MinHeap)
	heap.Push(pq, startNode)

	visited := make(map[UndirectedNode[T, S]]bool)
	for !pq.IsEmpty() {
		// Pop the top of pq and mark it as visited
		popped := heap.Pop(pq)
		if popped == nil {
			panic("popped nil - should not happen")
		}
		current, ok := popped.(UndirectedNode[T, S])
		if !ok {
			typeOfCurrent := reflect.TypeOf(current)
			panic(fmt.Sprintf("current is %s, not *Node[T]", typeOfCurrent))
		}
		visited[current] = true

		edges := self.GetNodeEdges(current)
		for _, edge := range edges {
			// Skip visited
			if visited[edge.GetNode()] {
				continue
			}

			heap.Push(pq, edge.GetNode())
			// If getting to edge from current is cheaper that the edge current cost state,
			// update it to pass via current instead
			if newCost := current.GetValue() + edge.GetWeight(); newCost < edge.GetNode().GetValue() {
				edge.GetNode().SetValue(newCost)
				edge.GetNode().SetThrough(current)
			}
		}
	}

	// Reconstruct path
	shortestPaths = make(map[UndirectedNode[T, S]][]UndirectedNode[T, S], len(self.graph.GetNodes()))
	for _, node := range self.GetNodes() {
		var path []UndirectedNode[T, S]
		// Keep going back until start (nil Through)
		// i.e. backward path
		for via := node; via.GetThrough() != nil; via = via.GetThrough() {
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
