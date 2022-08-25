package data

// Queue can be used in many scenarios, including when implementing a graph.
// See queueForGraph below for examples.

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	values0 := []uint8{0, 1, 200, 20}
	q0 := NewQueue[uint8]()
	testQueue(t, values0, q0)

	values1 := []string{"foo", "bar", "baz"}
	q1 := NewQueue[string]()
	testQueue(t, values1, q1)

	// Composite type queue - any comparable types should be ok in tests
	valuesComposite := []interface{}{1, 2, "last"}
	qComposite := NewQueue[interface{}]()
	// Test Push for composite queue
	qComposite.PushSlice(valuesComposite...)
	// Test Pop for composite queue
	for i, qLen := 0, qComposite.Len(); i < qLen; i++ {
		expected := valuesComposite[i]
		p := qComposite.Pop()
		if p == nil {
			t.Fatalf("Queue.Pop failed - expected %v, got nil", expected)
		}
		value := *p
		if value != expected {
			t.Fatalf("Queue.Pop failed - expected %v, got %v", expected, value)
		}
	}
	// Test Queue.IsEmpty for composite queue
	if !qComposite.IsEmpty() {
		t.Fatalf("Queue.IsEmpty failed - expected to be emptied")
	}
}

func testQueue[T comparable](t *testing.T, values []T, q *Queue[T]) {
	// Test Push
	for _, expected := range values {
		q.Push(expected)
		v := q.Pop()
		if *v != expected {
			t.Fatalf("Queue.Pop failed - expected %v, got %v", expected, v)
		}
	}

	// Test Len
	for i, v := range values {
		q.Push(v)

		if newLen := q.Len(); newLen != i+1 {
			t.Fatalf("Queue.Len failed - expected %d, got %d", newLen, i+1)
		}
	}

	// Test Pop
	for i, qLen := 0, q.Len(); i < qLen; i++ {
		popped := q.Pop()
		expected := values[i]
		if *popped != expected {
			t.Fatalf("Queue.Pop failed - expected %v, got %v", expected, popped)
		}
	}

	// Test IsEmpty
	if !q.IsEmpty() {
		t.Fatal("Stack.IsEmpty failed - expected true")
	}

	// Test Pop after emptied
	v := q.Pop()
	t.Logf("value of Pop() after emptied: %v\n", v)
}

func TestQueueForGraph(t *testing.T) {
	const (
		art  = "art"
		bob  = "bob"
		ron  = "ron"
		may  = "may"
		liz  = "liz"
		abe  = "abe"
		chad = "chad"
	)

	// Implement a graph with Go hash map
	graph := make(map[string][]string)
	graph[art] = []string{bob, liz}
	graph[bob] = []string{art, liz}
	graph[ron] = []string{may}
	graph[may] = []string{liz, ron}
	graph[liz] = []string{art, bob, may}
	graph[abe] = []string{chad}
	graph[chad] = []string{abe}

	if hops, found := isConnectedNotBestPath(graph, art, may); !found {
		t.Fatalf("queueForGraph (hops: %d): expected true", hops)
	} else {
		t.Logf("queueForGraph (hops: %d)", hops)
	}
	if hops, found := isConnectedNotBestPath(graph, ron, liz); !found {
		t.Fatalf("queueForGraph (hops: %d): expected true", hops)
	} else {
		t.Logf("queueForGraph (hops: %d)", hops)
	}
	if hops, found := isConnectedNotBestPath(graph, art, chad); found {
		t.Fatalf("queueForGraph (hops: %d): expected false", hops)
	} else {
		t.Logf("queueForGraph (hops: %d)", hops)
	}
}

// isConnectedNotBestPath returns the number of hops needed to look from node src to dst and whether dst is found.
// The returned returned number of hops is not neccessarily shortest path (which sucks).
func isConnectedNotBestPath[T comparable](inGraph map[T][]T, src, dst T) (int, bool) {
	var hops int
	searched := make(map[T]bool)

	if src == dst {
		return hops, true
	}
	searched[src] = true

	fmt.Println("hops++ 1")
	hops++
	// Copy graph
	graph := make(map[T][]T, len(inGraph))
	for k, v := range inGraph {
		graph[k] = v
	}
	// Get first neigbors and delete from map
	firstNeighbors := graph[src]
	delete(graph, src)

	if firstNeighbors == nil {
		fmt.Println("no neighbors for src", src)
		return hops, false
	}

	q := NewQueue[T]()
	q.PushSlice(firstNeighbors...)

	for {
		if q.IsEmpty() {
			fmt.Println("empty!")
			break
		}
		popped := q.Pop()
		if popped == nil {
			continue
		}

		thisNeighbor := *popped
		if searched[thisNeighbor] {
			continue
		}
		if thisNeighbor == dst {
			return hops, true
		}

		// Push nth-degree connections to queue
		q.PushSlice(graph[thisNeighbor]...)
		searched[thisNeighbor] = true

		hops++
		fmt.Println("hops++ n", thisNeighbor)
	}
	return hops, false
}
