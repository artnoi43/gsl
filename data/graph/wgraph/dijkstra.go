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

// TODO: integrate more comparable types, like big.Int and big.Float
// dijkstraWeight represents the allowed types for edge weights and node costs.
type dijkstraWeight interface {
	constraints.Integer | constraints.Float
}

// This type is the Dijkstra shortest path answer. It has 2 fields, (1) `From` the 'from' node, and (2) `Paths`.
// DijkstraShortestPath.Paths is a hash map where the key is a node, and the value is the previous node with the lowest cost to that key node.
type DijstraShortestPath[T dijkstraWeight, S ~string] struct {
	From  UndirectedNode[T, S]
	Paths map[UndirectedNode[T, S]]UndirectedNode[T, S]
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
func (self *DijkstraGraphUndirected[T, S]) DijkstraShortestPathFrom(startNode UndirectedNode[T, S]) *DijstraShortestPath[T, S] {
	var zeroValue T
	startNode.SetCost(zeroValue)
	startNode.SetThrough(nil)

	visited := make(map[UndirectedNode[T, S]]bool)
	prev := make(map[UndirectedNode[T, S]]UndirectedNode[T, S])

	pq := list.NewPriorityQueue[T](list.MinHeap)
	heap.Push(pq, startNode)

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
			edgeNode := edge.GetNode()

			// Skip visited
			if visited[edgeNode] {
				continue
			}

			heap.Push(pq, edgeNode)
			// If getting to edge from current is cheaper that the edge current cost state,
			// update it to pass via current instead
			if newCost := current.GetValue() + edge.GetWeight(); newCost < edgeNode.GetValue() {
				edgeNode.SetCost(newCost)
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
	shortestPaths map[UndirectedNode[T, S]]UndirectedNode[T, S],
	src UndirectedNode[T, S],
	dst UndirectedNode[T, S],
) []UndirectedNode[T, S] {
	prevNodes := []UndirectedNode[T, S]{dst}
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
