package graph

import (
	"testing"
)

type dijkTestUtil[T dijkstraWeight, S ~string] struct {
	inititalValue      T
	expectedFinalValue T
	expectedPathHops   int
	expectedPathway    []*UndirectedWeightedNodeImpl[T, S]

	edges []*struct {
		to     *UndirectedWeightedNodeImpl[T, S]
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

func constructDijkstraTestGraph[T dijkstraWeight, S ~string](nameStart, nameFinish S) map[UndirectedWeightedNode[T, S]]*dijkTestUtil[T, S] {
	// TODO: infinity is way too low, because dijkstraWeight also has uint8
	infinity := T(100)
	nodeStart := &UndirectedWeightedNodeImpl[T, S]{
		Name: nameStart,
		Cost: T(0),
	}
	nodeA := &UndirectedWeightedNodeImpl[T, S]{
		Name: "a",
		Cost: infinity,
	}
	nodeB := &UndirectedWeightedNodeImpl[T, S]{
		Name: "b",
		Cost: infinity,
	}
	nodeC := &UndirectedWeightedNodeImpl[T, S]{
		Name: "c",
		Cost: infinity,
	}
	nodeD := &UndirectedWeightedNodeImpl[T, S]{
		Name: "d",
		Cost: infinity,
	}
	nodeFinish := &UndirectedWeightedNodeImpl[T, S]{
		Name: nameFinish,
		Cost: infinity,
	}
	m := map[UndirectedWeightedNode[T, S]]*dijkTestUtil[T, S]{
		nodeStart: {
			inititalValue:      T(0),
			expectedFinalValue: T(0),
			expectedPathHops:   0,
			expectedPathway:    []*UndirectedWeightedNodeImpl[T, S]{},
			edges: []*struct {
				to     *UndirectedWeightedNodeImpl[T, S]
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
			expectedPathway:    []*UndirectedWeightedNodeImpl[T, S]{nodeD, nodeFinish},
			edges: []*struct {
				to     *UndirectedWeightedNodeImpl[T, S]
				weight T
			}{},
		},
		nodeA: {
			inititalValue:      infinity,
			expectedFinalValue: T(2),
			expectedPathHops:   1,
			expectedPathway:    []*UndirectedWeightedNodeImpl[T, S]{nodeA},
			edges: []*struct {
				to     *UndirectedWeightedNodeImpl[T, S]
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
			expectedPathway:    []*UndirectedWeightedNodeImpl[T, S]{nodeA, nodeB},
			edges: []*struct {
				to     *UndirectedWeightedNodeImpl[T, S]
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
			expectedPathway:    []*UndirectedWeightedNodeImpl[T, S]{nodeA, nodeC},
			edges: []*struct {
				to     *UndirectedWeightedNodeImpl[T, S]
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
			expectedPathway:    []*UndirectedWeightedNodeImpl[T, S]{nodeD},
			edges: []*struct {
				to     *UndirectedWeightedNodeImpl[T, S]
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
func testDijkstra[T dijkstraWeight, S ~string](t *testing.T, nameStart, nameFinish S) {
	nodesMap := constructDijkstraTestGraph[T](nameStart, nameFinish)

	// Prepare graph
	djikGraph := NewDijkstraGraph[T, S]()
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

	var startNode UndirectedWeightedNode[T, S]
	for node := range nodesMap {
		if node.GetKey() == nameStart {
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
		if node.GetValue() != util.expectedFinalValue {
			t.Fatalf(fatalMsgCost, nameStart, node.GetKey(), node.GetValue(), util.expectedFinalValue)
		}
		// Test paths
		shortestPath := shortestPaths[node]
		if hops := len(shortestPath); hops != util.expectedPathHops {
			t.Fatalf(fatalMsgPathLen, nameStart, node.GetKey(), hops, util.expectedPathHops)
		}
		for i, actualPathway := range shortestPath {
			if expectedPathway := util.expectedPathway[i]; expectedPathway != actualPathway {
				t.Fatalf(fatalMsgPathVia, nameStart, node.GetKey(), i, actualPathway.GetKey(), expectedPathway.Name)
			}
		}
	}
}
