package wgraph

import (
	"testing"

	"github.com/artnoi43/mgl/mglutils"
)

type dijkstraTestUtils[T dijkstraWeight, S ~string] struct {
	inititalValue      T
	expectedFinalValue T
	expectedPathHops   int
	expectedPathway    []*NodeWeightedImpl[T, S]

	edges []*struct {
		to     *NodeWeightedImpl[T, S]
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

func constructDijkstraTestGraph[T dijkstraWeight, S ~string](nameStart, nameFinish S) map[WeightedNode[T, S]]*dijkstraTestUtils[T, S] {
	// TODO: infinity is way too low, because dijkstraWeight also has uint8
	infinity := T(100)
	nodeStart := &NodeWeightedImpl[T, S]{
		Name: nameStart,
		Cost: T(0),
	}
	nodeA := &NodeWeightedImpl[T, S]{
		Name: "A",
		Cost: infinity,
	}
	nodeB := &NodeWeightedImpl[T, S]{
		Name: "B",
		Cost: infinity,
	}
	nodeC := &NodeWeightedImpl[T, S]{
		Name: "C",
		Cost: infinity,
	}
	nodeD := &NodeWeightedImpl[T, S]{
		Name: "D",
		Cost: infinity,
	}
	nodeFinish := &NodeWeightedImpl[T, S]{
		Name: nameFinish,
		Cost: infinity,
	}
	m := map[WeightedNode[T, S]]*dijkstraTestUtils[T, S]{
		nodeStart: {
			inititalValue:      T(0),
			expectedFinalValue: T(0),
			expectedPathHops:   1,
			expectedPathway:    []*NodeWeightedImpl[T, S]{},
			edges: []*struct {
				to     *NodeWeightedImpl[T, S]
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
			expectedPathHops:   3,
			expectedPathway:    []*NodeWeightedImpl[T, S]{nodeStart, nodeD, nodeFinish},
			edges: []*struct {
				to     *NodeWeightedImpl[T, S]
				weight T
			}{},
		},
		nodeA: {
			inititalValue:      infinity,
			expectedFinalValue: T(2),
			expectedPathHops:   2,
			expectedPathway:    []*NodeWeightedImpl[T, S]{nodeStart, nodeA},
			edges: []*struct {
				to     *NodeWeightedImpl[T, S]
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
			expectedPathHops:   3,
			expectedPathway:    []*NodeWeightedImpl[T, S]{nodeStart, nodeA, nodeB},
			edges: []*struct {
				to     *NodeWeightedImpl[T, S]
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
			expectedPathHops:   3,
			expectedPathway:    []*NodeWeightedImpl[T, S]{nodeStart, nodeA, nodeC},
			edges: []*struct {
				to     *NodeWeightedImpl[T, S]
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
			expectedPathHops:   2,
			expectedPathway:    []*NodeWeightedImpl[T, S]{nodeStart, nodeD},
			edges: []*struct {
				to     *NodeWeightedImpl[T, S]
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
	hasDirection := true
	djikGraph := NewDijkstraGraph[T, S](hasDirection)
	for node, util := range nodesMap {
		// Add node
		djikGraph.AddNode(node)
		// Add edges
		nodeEdges := util.edges
		for _, edge := range nodeEdges {
			if err := djikGraph.AddEdge(node, edge.to, edge.weight); err != nil {
				t.Error(err.Error())
			}
		}
	}

	var startNode WeightedNode[T, S]
	for node := range nodesMap {
		if node.GetKey() == nameStart {
			startNode = node
		}
	}

	dijkShortestPaths := djikGraph.DijkstraShortestPathFrom(startNode)
	fatalMsgCost := "invalid answer for (%s->%s): %d, expecting %d"
	fatalMsgPathLen := "invalid returned path length (%s->%s): %d, expecting %d"
	fatalMsgPathVia := "invalid via path (%s->%s)[%d]: %s, expecting %s"

	// The check is hard-coded
	for node, util := range nodesMap {
		// Test costs
		if node.GetValue() != util.expectedFinalValue {
			t.Fatalf(fatalMsgCost, nameStart, node.GetKey(), node.GetValue(), util.expectedFinalValue)
		}

		if node == startNode {
			continue
		}
		nodeKey := node.GetKey()
		// t.Logf("dst node: %v\n", nodeKey)
		// Test paths
		pathToNode := DijkstraShortestPathReconstruct(dijkShortestPaths.Paths, startNode, node)
		mglutils.ReverseInPlace(pathToNode)
		if hops := len(pathToNode); hops != util.expectedPathHops {
			t.Fatalf(fatalMsgPathLen, nameStart, nodeKey, hops, util.expectedPathHops)
		}
		for i, actual := range pathToNode {
			expected := util.expectedPathway[i]
			// t.Logf("prev[%d]: %v, expected: %v\n", i, actual.GetKey(), expected.GetKey())
			if expected != actual {
				t.Fatalf(fatalMsgPathVia, nameStart, nodeKey, i, actual.GetKey(), expected.GetKey())
			}
		}
	}
}
