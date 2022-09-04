package main

import (
	"fmt"
	"math"
	"reflect"

	"github.com/artnoi43/mgl/data/graph"
	"github.com/artnoi43/mgl/data/graph/wgraph"
	"github.com/artnoi43/mgl/mglutils"
)

// This code purpose is to show an example of how to use the graph and wgraph packages.
// The code provides an example of solving the best flight path between major cities.
// It uses a weight graph to do Dijkstra sort, and an unweighted graph for BFS.

type cityName string

// *city implements wgraph.UndirectedNode
type city struct {
	name    cityName
	cost    float64
	through *city
}

func main() {
	infinity := math.MaxFloat64

	tokyo := &city{name: cityName("Tokyo"), cost: 0, through: nil}
	bangkok := &city{name: cityName("Bangkok"), cost: infinity, through: nil}
	hongkong := &city{name: cityName("Hongkong"), cost: infinity, through: nil}
	dubai := &city{name: cityName("Dubai"), cost: infinity, through: nil}
	helsinki := &city{name: cityName("Helsinki"), cost: infinity, through: nil}
	budapest := &city{name: cityName("Budapest"), cost: infinity, through: nil}

	// See file flight_graph.png
	graphEdges := map[wgraph.WeightedNode[float64, cityName]][]struct {
		to       wgraph.WeightedNode[float64, cityName]
		flighDur float64
	}{
		tokyo: {
			{
				to:       bangkok,
				flighDur: 4,
			},
			{
				to:       hongkong,
				flighDur: 1.5,
			},
			{
				to:       dubai,
				flighDur: 7,
			},
		},
		hongkong: {
			{
				to:       helsinki,
				flighDur: 11.5,
			},
		},
		bangkok: {
			{
				to:       dubai,
				flighDur: 6,
			},
			{
				to:       helsinki,
				flighDur: 9,
			},
		},
		dubai: {
			{
				to:       helsinki,
				flighDur: 3,
			},
			{
				to:       budapest,
				flighDur: 5,
			},
		},
		helsinki: {
			{
				to:       budapest,
				flighDur: 1.5,
			},
		},
		budapest: {},
	}

	hasDirection := false
	dijkGraph := wgraph.NewDijkstraGraph[float64, cityName](hasDirection)
	unweightedGraph := graph.NewGraph[float64](hasDirection)

	// Add edges and nodes to graphs
	for node, nodeEdges := range graphEdges {
		dijkGraph.AddNode(node)
		unweightedGraph.AddNode(node)
		for _, nodeEdge := range nodeEdges {
			// fmt.Println(node.GetKey()+":", "adding edge", nodeEdge.to.GetKey(), "weight", nodeEdge.flighDur)
			if err := dijkGraph.AddEdge(node, nodeEdge.to, nodeEdge.flighDur); err != nil {
				panic("failed to add dijkstra-compatible graph edge: " + err.Error())
			}
			unweightedGraph.AddEdge(node, nodeEdge.to, nil)
		}
	}

	fromNode := tokyo
	shortestPathsFromTokyo := dijkGraph.DijkstraShortestPathFrom(fromNode)

	fmt.Println("Dijkstra result")
	for _, dst := range dijkGraph.GetNodes() {
		if dst == fromNode {
			continue
		}
		pathToNode := wgraph.DijkstraShortestPathReconstruct(shortestPathsFromTokyo.Paths, shortestPathsFromTokyo.From, dst)
		mglutils.ReverseInPlace(pathToNode)

		fmt.Println("> from", fromNode.GetKey(), "to", dst.GetKey(), "min cost", dst.GetValue())
		for _, via := range pathToNode {
			fmt.Printf("%s (%v) ", via.GetKey(), via.GetValue())
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("BFS result")
	for _, dst := range unweightedGraph.GetNodes() {
		shortestHopsFromTokyo, hops, found := graph.BFSShortestPath[float64](unweightedGraph, fromNode, dst)
		fmt.Println("path to", dst.(*city).GetKey(), "found", found, "shortestHops", hops)
		mglutils.ReverseInPlace(shortestHopsFromTokyo)
		for i, hop := range shortestHopsFromTokyo {
			fmt.Println("hop", i, hop.(*city).GetKey())
		}
	}

	takeUnweightedGraph(unweightedGraph, fromNode)
	takeWeightedGraph(dijkGraph, fromNode)
	// Compile error: wgraph.WeightedGraph are not compatible with graph.Graph.
	// But the node types are pretty much interchangable.
	// takeUnweightedGraph(dijkGraph, fromNode)
	// takeWeightedGraph(unweightedGraph, fromNode)

	// Compile error: graph.GenericGraph is too primitive to be used easily. Use graph.Graph or wgraph.GraphWeighted instead
	// takeGenericGraph(unweightedGraph, fromNode)
	// takeGenericGraph(dijkGraph, fromNode)

}

// graph.GenericGraph is impractical
func takeGenericGraph(g graph.GenericGraph[*city, *city, float64, map[*city]*city], from graph.Node[float64]) {

}

func takeUnweightedGraph(g graph.Graph[float64], from graph.Node[float64]) {

}

func takeWeightedGraph(g wgraph.WeightedGraph[float64, cityName], from wgraph.WeightedNode[float64, cityName]) {

}

func (self *city) GetValue() float64 {
	return self.cost
}

func (self *city) GetKey() cityName {
	return self.name
}

func (self *city) GetThrough() wgraph.WeightedNode[float64, cityName] {
	if self.through == nil {
		return nil
	}
	return self.through
}

func (self *city) SetCost(newCost float64) {
	self.cost = newCost
}

func (self *city) SetThrough(node wgraph.WeightedNode[float64, cityName]) {
	if node == nil {
		self.through = nil
		return
	}

	via, ok := node.(*city)
	if !ok {
		typeOfNode := reflect.TypeOf(node)
		panic(fmt.Sprintf("node not *city but %s", typeOfNode))
	}
	self.through = via
}
