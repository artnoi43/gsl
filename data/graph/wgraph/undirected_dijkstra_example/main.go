package main

import (
	"fmt"
	"math"
	"reflect"

	"github.com/artnoi43/mgl/data/graph/wgraph"
	"github.com/artnoi43/mgl/mglutils"
)

// *city implements wgraph.UndirectedNode
type city struct {
	name    string
	cost    float64
	through *city
}

func main() {
	infinity := math.MaxFloat64

	tokyo := &city{name: "Tokyo", cost: 0, through: nil}
	bangkok := &city{name: "Bangkok", cost: infinity, through: nil}
	hongkong := &city{name: "Hongkong", cost: infinity, through: nil}
	dubai := &city{name: "Dubai", cost: infinity, through: nil}
	helsinki := &city{name: "Helsinki", cost: infinity, through: nil}
	budapest := &city{name: "Budapest", cost: infinity, through: nil}

	// See file flight_graph.png
	graphEdges := map[wgraph.UndirectedNode[float64, string]][]struct {
		to       wgraph.UndirectedNode[float64, string]
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

	graph := wgraph.NewDijkstraGraph[float64, string]()
	// Add nodes
	for node, nodeEdges := range graphEdges {
		graph.AddNode(node)
		for _, nodeEdge := range nodeEdges {
			// fmt.Println(node.GetKey()+":", "adding edge", nodeEdge.to.GetKey(), "weight", nodeEdge.flighDur)
			if err := graph.AddEdge(node, nodeEdge.to, nodeEdge.flighDur); err != nil {
				panic("failed to add dijkstra-compatible graph edge: " + err.Error())
			}
		}
	}

	fromNode := tokyo
	shortestPathsFromTokyo := graph.DijkstraShortestPathFrom(fromNode)

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
}

func (self *city) GetValue() float64 {
	return self.cost
}

func (self *city) GetKey() string {
	return self.name
}

func (self *city) GetThrough() wgraph.UndirectedNode[float64, string] {
	if self.through == nil {
		return nil
	}
	return self.through
}

func (self *city) SetCost(newCost float64) {
	self.cost = newCost
}

func (self *city) SetThrough(node wgraph.UndirectedNode[float64, string]) {
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
