package main

import (
	"fmt"
	"math"
	"reflect"

	graphlib "github.com/artnoi43/mgl/data/graph"
	"github.com/artnoi43/mgl/data/graph/wgraph"
	"github.com/artnoi43/mgl/mglutils"
)

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
	graph := wgraph.NewDijkstraGraph[float64, cityName](hasDirection)
	dumbGraph := graphlib.NewGraph(false)

	// Add nodes
	for node, nodeEdges := range graphEdges {
		graph.AddNode(node)
		dumbGraph.AddNode(node)
		for _, nodeEdge := range nodeEdges {
			// fmt.Println(node.GetKey()+":", "adding edge", nodeEdge.to.GetKey(), "weight", nodeEdge.flighDur)
			if err := graph.AddEdge(node, nodeEdge.to, nodeEdge.flighDur); err != nil {
				panic("failed to add dijkstra-compatible graph edge: " + err.Error())
			}
			dumbGraph.AddEdge(node, nodeEdge.to)
		}
	}

	fromNode := tokyo
	shortestPathsFromTokyo := graph.DijkstraShortestPathFrom(fromNode)

	fmt.Println("Dijkstra result")
	for _, dst := range graph.GetNodes() {
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

	fmt.Println("BFS result")
	shortestHopsFromTokyo, hops, found := graphlib.BFSShortestPath(dumbGraph, fromNode, budapest)
	fmt.Println("found", found, "shortestHops", hops)
	for i, hop := range shortestHopsFromTokyo {
		hopNode := hop.(*city)
		fmt.Println("hop", i, hopNode.GetKey())
	}
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
