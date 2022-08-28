package main

import (
	"fmt"
	"math"
	"reflect"

	"github.com/artnoi43/mgl/data/graph/wgraph"
)

// *city implements wgraph.UndirectedNode
type city struct {
	name    string
	value   float64
	through *city
}

func main() {
	infinity := math.MaxFloat64

	tokyo := &city{name: "tokyo", value: 0, through: nil}
	bangkok := &city{name: "bangkok", value: infinity, through: nil}
	hongkong := &city{name: "hongkong", value: infinity, through: nil}
	dubai := &city{name: "dubai", value: infinity, through: nil}
	helsinki := &city{name: "helsinki", value: infinity, through: nil}
	budapest := &city{name: "budapest", value: infinity, through: nil}

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
				flighDur: 7,
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

	shortestPaths := graph.DijkstraFrom(tokyo)
	for dst, vias := range shortestPaths {
		fmt.Println(">", "from", tokyo.GetKey(), "to", dst.GetKey(), "dur", dst.GetValue())
		fmt.Print("via")
		for _, via := range vias {
			fmt.Print(" " + via.GetKey() + " ")
		}
		fmt.Println()
	}
}

func (self *city) GetValue() float64 {
	return self.value
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

func (self *city) SetValue(value float64) {
	self.value = value
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
