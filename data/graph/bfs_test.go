package graph

import (
	"reflect"
	"testing"
)

type (
	genericBFS[T nodeValue] func(GenericGraph[Node[T], Node[T], any], Node[T], Node[T]) ([]Node[T], int, bool)
	hashMapBFS[T nodeValue] func(HashMapGraph[T], Node[T], Node[T]) ([]Node[T], int, bool)
	testFunc[T nodeValue]   interface{ genericBFS[T] | hashMapBFS[T] }
)

type bfsTestResult struct {
	found bool
	hops  int
}

func TestBFS(t *testing.T) {
	t.Run("Directed BFS", testBFS)
	t.Run("Undirected BFS", testUBFS)
}

// See visualization in directory assets
func testBFS(t *testing.T) {
	art := person{name: "art", age: 25}
	beagie := person{name: "beagie", age: 3}
	banana := person{name: "banana", age: 2}
	black := person{name: "black", age: 2}
	makam := person{name: "makam", age: 5}
	muji := person{name: "muji", age: 1}

	tests := map[person]map[person]bfsTestResult{
		art: {
			art: {
				found: true,
				hops:  0,
			},
			makam: {
				found: true,
				hops:  2,
			},
			muji: {
				found: true,
				hops:  1,
			},
		},
		muji: {
			art: {
				found: false,
				hops:  -1,
			},
		},
		banana: {
			black: {
				found: true,
				hops:  2,
			},
		},
		black: {
			banana: {
				found: false,
				hops:  -1,
			},
		},
		makam: {
			black: {
				found: false,
				hops:  -1,
			},
		},
	}

	people := []person{art, beagie, black, makam}
	g := NewGraph[int](true)

	for _, friend := range people {
		g.AddNode(friend)
	}

	g.AddEdge(art, beagie, nil)
	g.AddEdge(art, black, nil)
	g.AddEdge(art, banana, nil)
	g.AddEdge(art, muji, nil)
	g.AddEdge(beagie, art, nil)
	g.AddEdge(beagie, banana, nil)
	g.AddEdge(beagie, black, nil)
	g.AddEdge(banana, art, nil)
	g.AddEdge(banana, beagie, nil)
	g.AddEdge(black, makam, nil)

	loopTestBFS[int](t, tests, g)
}

func testUBFS(t *testing.T) {
	art := person{name: "art", age: 25}      // art knows beagie, banana, and black
	beagie := person{name: "beagie", age: 3} // beagie knows art, banana, and black
	banana := person{name: "banana", age: 2} // banana knows art, beagie, and black
	black := person{name: "black", age: 2}   // black knows art, beagie, banana, and makam
	makam := person{name: "makam", age: 5}   // makam only knows black
	muji := person{name: "muji", age: 1}     // muji knows no one

	tests := map[person]map[person]bfsTestResult{
		art: {
			art: {
				found: true,
				hops:  0,
			},
			makam: {
				found: true,
				hops:  2,
			},
			muji: {
				found: false,
				hops:  -1,
			},
		},
		banana: {
			makam: {
				found: true,
				hops:  2,
			},
		},
	}

	people := []person{art, beagie, black, makam}
	g := NewGraph[int](false)

	for _, friend := range people {
		g.AddNode(friend)
	}

	g.AddEdge(art, beagie, nil)
	g.AddEdge(art, black, nil)
	g.AddEdge(art, banana, nil)
	g.AddEdge(beagie, black, nil)
	g.AddEdge(beagie, banana, nil)
	g.AddEdge(black, makam, nil)
	g.AddEdge(black, banana, nil)

	loopTestBFS[int](t, tests, g)
}

func loopTestBFS[T any](
	t *testing.T,
	tests map[person]map[person]bfsTestResult,
	g HashMapGraph[int],
) {
	var gf genericBFS[int] = BFSNg[int]
	var hf hashMapBFS[int] = BFS[int]
	// Using BFSNg[int] and BFS[int] directly in testFuncs will fail type system
	testFuncs := []interface{}{gf, hf}
	for _, tf := range testFuncs {
		for fromNode, m := range tests {
			for toNode, expected := range m {
				var shortestPath []Node[int]
				var hops int
				var found bool

				switch f := tf.(type) {
				case genericBFS[int]:
					shortestPath, hops, found = f(g.(GenericGraph[Node[int], Node[int], any]), fromNode, toNode)
				case hashMapBFS[int]:
					shortestPath, hops, found = f(g, fromNode, toNode)
				default:
					t.Fatal("unexpected function type", reflect.TypeOf(f).String())
				}

				if found != expected.found {
					t.Log(shortestPath)
					t.Fatalf("unexpected found: %v, expected %v\n", found, expected.found)
				}
				if hops != expected.hops {
					t.Log(shortestPath)
					t.Fatalf("unexpected hops: %v, expected %v\n", hops, expected.hops)
				}
				if hops != len(shortestPath)-1 {
					t.Fatal("unexpected relationship for hops and len(shortestPath)")
				}
			}
		}
	}
}

type person struct {
	name string
	age  int
}

func (p person) GetKey() string {
	return p.name
}

func (p person) GetValue() int {
	return p.age
}
