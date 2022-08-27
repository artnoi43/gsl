package data

import (
	"testing"

	"golang.org/x/exp/constraints"
)

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

	graph := NewWeightedGraph[T]()
	infinity := ^T(0)

	nodeStart := &Node[T]{
		Name: nameStart,
		Cost: 0,
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
	for _, Node := range nodes {
		graph.AddNode(Node)
	}

	graph.AddEdge(nodeStart, nodeA, T(2))
	graph.AddEdge(nodeStart, nodeB, T(4))
	graph.AddEdge(nodeStart, nodeD, T(4))
	graph.AddEdge(nodeA, nodeB, T(1))
	graph.AddEdge(nodeA, nodeC, T(2))
	graph.AddEdge(nodeC, nodeD, T(2))
	graph.AddEdge(nodeB, nodeFinish, T(5))
	graph.AddEdge(nodeD, nodeFinish, T(3))

	startNode := graph.GetNodeByName(nameStart)
	shortestPaths := graph.DjikstraFrom(startNode)
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
			expectedCost = 0
			expectedHops = 0
			expectedVias = []*Node[T]{}
		case nodeFinish:
			expectedCost = 7
			expectedHops = 2 // -> d -> finish
			expectedVias = []*Node[T]{nodeD, nodeFinish}
		case nodeA:
			expectedCost = 2
			expectedHops = 1 // -> a
			expectedVias = []*Node[T]{nodeA}
		case nodeB:
			expectedCost = 3
			expectedHops = 2 // -> a -> b
			expectedVias = []*Node[T]{nodeA, nodeB}
		case nodeC:
			expectedCost = 4
			expectedHops = 2 // -> a -> c
			expectedVias = []*Node[T]{nodeA, nodeC}
		case nodeD:
			expectedCost = 4
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
