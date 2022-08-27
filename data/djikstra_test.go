package data

import (
	"testing"

	"golang.org/x/exp/constraints"
)

// TODO: Add tests for other types

func TestDjikstra(t *testing.T) {
	testDjikstra[uint](t)
	testDjikstra[uint8](t)
	testDjikstra[uint32](t)
	testDjikstra[uint64](t)
}

func testDjikstra[T constraints.Unsigned](t *testing.T) {
	const (
		nameStart  = "start"
		nameFinish = "finish"
	)

	djikGraph := NewDjikstraGraph[T]()
	infinity := ^T(0)

	nodeStart := &Node[T]{
		Name: nameStart,
		Cost: T(0),
	}
	nodeA := &Node[T]{
		Name: "a",
		Cost: infinity,
	}
	nodeB := &Node[T]{
		Name: "b",
		Cost: infinity,
	}
	nodeC := &Node[T]{
		Name: "c",
		Cost: infinity,
	}
	nodeD := &Node[T]{
		Name: "d",
		Cost: infinity,
	}
	nodeFinish := &Node[T]{
		Name: nameFinish,
		Cost: infinity,
	}
	nodes := []*Node[T]{nodeStart, nodeA, nodeB, nodeC, nodeD, nodeFinish}

	for _, node := range nodes {
		djikGraph.graph.AddNode(node)
	}
	djikGraph.AddDjikstraEdge(nodeStart, nodeA, T(2))
	djikGraph.AddDjikstraEdge(nodeStart, nodeB, T(4))
	djikGraph.AddDjikstraEdge(nodeStart, nodeD, T(4))
	djikGraph.AddDjikstraEdge(nodeA, nodeB, T(1))
	djikGraph.AddDjikstraEdge(nodeA, nodeC, T(2))
	djikGraph.AddDjikstraEdge(nodeC, nodeD, T(2))
	djikGraph.AddDjikstraEdge(nodeB, nodeFinish, T(5))
	djikGraph.AddDjikstraEdge(nodeD, nodeFinish, T(3))

	shortestPaths := djikGraph.DjikstraFrom(nodeStart)
	fatalMsgCost := "invalid answer for (%s->%s): %d, expecting %d"
	fatalMsgPathLen := "invalid returned path (%s->%s): %d, expecting %d"
	fatalMsgPathVia := "invalid via path (%s->%s)[%d]: %s, expecting %s"

	// The check is hard-coded
	for _, node := range nodes {
		var expectedCost T
		var expectedHops int
		var expectedVias []*Node[T]
		switch node {
		case nodeStart:
			expectedCost = T(0)
			expectedHops = 0
			expectedVias = []*Node[T]{}
		case nodeFinish:
			expectedCost = T(7)
			expectedHops = 2 // -> d -> finish
			expectedVias = []*Node[T]{nodeD, nodeFinish}
		case nodeA:
			expectedCost = T(2)
			expectedHops = 1 // -> a
			expectedVias = []*Node[T]{nodeA}
		case nodeB:
			expectedCost = T(3)
			expectedHops = 2 // -> a -> b
			expectedVias = []*Node[T]{nodeA, nodeB}
		case nodeC:
			expectedCost = T(4)
			expectedHops = 2 // -> a -> c
			expectedVias = []*Node[T]{nodeA, nodeC}
		case nodeD:
			expectedCost = T(4)
			expectedHops = 1 // -> d
			expectedVias = []*Node[T]{nodeD}
		}

		// Test costs
		if node.Cost != expectedCost {
			t.Fatalf(fatalMsgCost, nameStart, node.Name, node.Cost, expectedCost)
		}

		// Test paths
		shortestPath := shortestPaths[node]
		if hops := len(shortestPath); hops != expectedHops {
			t.Fatalf(fatalMsgPathLen, nameStart, node.Name, hops, expectedHops)
		}
		for i, actualVia := range shortestPath {
			if expectedVia := expectedVias[i]; expectedVia != actualVia {
				t.Fatalf(fatalMsgPathVia, nameStart, node.Name, i, actualVia.Name, expectedVia.Name)
			}
		}
	}
}
