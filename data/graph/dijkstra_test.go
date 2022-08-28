package graph

import (
	"testing"
)

type dijkTestUtil[T dijkstraWeight] struct {
	inititalValue      T
	expectedFinalValue T
	expectedPathHops   int
	expectedPathway    []*Node[T]

	edges []*struct {
		to     *Node[T]
		weight T
	}
}

// TODO: Add tests for other types

func TestDijkstra(t *testing.T) {
	const (
		nameStart  = "start"
		nameFinish = "finish"
	)
	testDijkstra[uint](t, nameStart, nameFinish)
	testDijkstra[uint8](t, nameStart, nameFinish)
	testDijkstra[int32](t, nameStart, nameFinish)
	testDijkstra[float64](t, nameStart, nameFinish)
}

func constructDijkstraTestGraph[T dijkstraWeight](nameStart, nameFinish string) map[*Node[T]]*dijkTestUtil[T] {
	// TODO: infinity is way too low, because dijkstraWeight also has uint8
	infinity := T(100)
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
	m := map[*Node[T]]*dijkTestUtil[T]{
		nodeStart: {
			inititalValue:      T(0),
			expectedFinalValue: T(0),
			expectedPathHops:   0,
			expectedPathway:    []*Node[T]{},
			edges: []*struct {
				to     *Node[T]
				weight T
			}{
				{
					to:     nodeA,
					weight: T(2),
				}, {
					to:     nodeB,
					weight: T(4),
				}, {
					to:     nodeD,
					weight: T(4),
				},
			},
		},
		nodeFinish: {
			inititalValue:      infinity,
			expectedFinalValue: T(7),
			expectedPathHops:   2,
			expectedPathway:    []*Node[T]{nodeD, nodeFinish},
			edges: []*struct {
				to     *Node[T]
				weight T
			}{},
		},
		nodeA: {
			inititalValue:      infinity,
			expectedFinalValue: T(2),
			expectedPathHops:   1,
			expectedPathway:    []*Node[T]{nodeA},
			edges: []*struct {
				to     *Node[T]
				weight T
			}{
				{
					to:     nodeB,
					weight: T(1),
				},
				{
					to:     nodeC,
					weight: T(2),
				},
			},
		},
		nodeB: {
			inititalValue:      infinity,
			expectedFinalValue: T(3),
			expectedPathHops:   2,
			expectedPathway:    []*Node[T]{nodeA, nodeB},
			edges: []*struct {
				to     *Node[T]
				weight T
			}{
				{
					to:     nodeFinish,
					weight: T(5),
				},
			},
		},
		nodeC: {
			inititalValue:      infinity,
			expectedFinalValue: T(4),
			expectedPathHops:   2,
			expectedPathway:    []*Node[T]{nodeA, nodeC},
			edges: []*struct {
				to     *Node[T]
				weight T
			}{
				{
					to:     nodeD,
					weight: T(2),
				},
			},
		},
		nodeD: {
			inititalValue:      infinity,
			expectedFinalValue: T(4),
			expectedPathHops:   1,
			expectedPathway:    []*Node[T]{nodeD},
			edges: []*struct {
				to     *Node[T]
				weight T
			}{
				{
					to:     nodeFinish,
					weight: T(3),
				},
			},
		},
	}
	return m
}

// The weighted graph used in this test can be viewed at assets/dijkstra_test_graph.png
func testDijkstra[T dijkstraWeight](t *testing.T, nameStart, nameFinish string) {
	nodesMap := constructDijkstraTestGraph[T](nameStart, nameFinish)

	// Prepare graph
	djikGraph := NewDijkstraGraph[T]()
	for node, util := range nodesMap {
		// Add node
		djikGraph.AddDijkstraNode(node)
		// Add edges
		nodeEdges := util.edges
		for _, edge := range nodeEdges {
			if err := djikGraph.AddDijkstraEdge(node, edge.to, edge.weight); err != nil {
				t.Error(err.Error())
			}
		}
	}

	var startNode *Node[T]
	for node := range nodesMap {
		if node.Name == nameStart {
			startNode = node
		}
	}

	shortestPaths := djikGraph.DijkstraFrom(startNode)
	fatalMsgCost := "invalid answer for (%s->%s): %d, expecting %d"
	fatalMsgPathLen := "invalid returned path (%s->%s): %d, expecting %d"
	fatalMsgPathVia := "invalid via path (%s->%s)[%d]: %s, expecting %s"

	// The check is hard-coded
	for node, util := range nodesMap {
		// Test costs
		if node.Cost != util.expectedFinalValue {
			t.Fatalf(fatalMsgCost, nameStart, node.Name, node.Cost, util.expectedFinalValue)
		}
		// Test paths
		shortestPath := shortestPaths[node]
		if hops := len(shortestPath); hops != util.expectedPathHops {
			t.Fatalf(fatalMsgPathLen, nameStart, node.Name, hops, util.expectedPathHops)
		}
		for i, actualPathway := range shortestPath {
			if expectedPathway := util.expectedPathway[i]; expectedPathway != actualPathway {
				t.Fatalf(fatalMsgPathVia, nameStart, node.Name, i, actualPathway.Name, expectedPathway.Name)
			}
		}
	}
}
