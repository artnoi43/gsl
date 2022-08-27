package data

// TODO: more relaxed generic types

import (
	"container/heap"
	"fmt"
	"reflect"
	"sync"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Unsigned] struct {
	Name    string
	Cost    T
	Through *Node[T]
}

// Implements data.ItemPQ[T]
func (self *Node[T]) Value() T {
	return self.Cost
}

type Edge[T constraints.Unsigned] struct {
	Node   *Node[T]
	weight T
}

type WeightedGraph[T constraints.Unsigned] struct {
	Nodes []*Node[T]
	Edges map[*Node[T]][]*Edge[T] // Edges is a map of a Node's edges

	// TODO: should we make mut public/exported?
	mut sync.RWMutex
}

func NewWeightedGraph[T constraints.Unsigned]() *WeightedGraph[T] {
	return &WeightedGraph[T]{
		Edges: make(map[*Node[T]][]*Edge[T]),
	}
}

func (self *WeightedGraph[T]) GetNodeByName(name string) (Node *Node[T]) {
	self.mut.RLock()
	defer self.mut.RUnlock()
	for _, n := range self.Nodes {
		if n.Name == name {
			Node = n
		}
	}
	return
}

func (self *WeightedGraph[T]) AddNode(n *Node[T]) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Nodes = append(self.Nodes, n)
}

func AddNodes[T constraints.Unsigned](graph *WeightedGraph[T], names ...string) (nodes map[string]*Node[T]) {
	nodes = make(map[string]*Node[T])
	for _, name := range names {
		n := &Node[T]{
			Name:    name,
			Cost:    ^T(0),
			Through: nil,
		}
		graph.AddNode(n)
		nodes[name] = n
	}
	return
}

func (self *WeightedGraph[T]) AddEdge(n1, n2 *Node[T], weight T) {
	self.mut.Lock()
	defer self.mut.Unlock()
	self.Edges[n1] = append(self.Edges[n1], &Edge[T]{Node: n2, weight: weight})
	self.Edges[n2] = append(self.Edges[n2], &Edge[T]{Node: n1, weight: weight})
}

func (self *WeightedGraph[T]) DjikstraFrom(startNode *Node[T]) (shortestPaths map[*Node[T]][]*Node[T]) {
	startNode.Cost = 0
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

		Edges := self.Edges[current]
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
	shortestPaths = make(map[*Node[T]][]*Node[T], len(self.Nodes))
	for _, node := range self.Nodes {
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
