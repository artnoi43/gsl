package data

// Queue can be used in many scenarios, including when implementing a graph.
// See queueForGraph below for examples.

import (
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
	for _, value := range valuesComposite {
		qComposite.Push(value)
	}
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

	if hops, found := isConnected(graph, art, may, 0); !found {
		t.Fatalf("queueForGraph (hops: %d): expected true", hops)
	}
	if hops, found := isConnected(graph, ron, liz, 0); !found {
		t.Fatalf("queueForGraph (hops: %d): expected true", hops)
	}
	if hops, found := isConnected(graph, art, chad, 0); found {
		t.Fatalf("queueForGraph (hops: %d): expected false", hops)
	}
}

// isConnected recursively traveses through a copy of inGraph to find dst.
// inGraph elements can point to themselves for this function.
func isConnected[T comparable](inGraph map[T][]T, src, dst T, r int) (int, bool) {
	if src == dst {
		return r, true
	}

	// Copy graph
	graph := make(map[T][]T, len(inGraph))
	for k, v := range inGraph {
		graph[k] = v
	}

	directNeighbors := graph[src]
	// Delete key src to avoid redundant works
	delete(graph, src)

	// If node src points to nothing
	if directNeighbors == nil {
		return r, false
	}

	// Check 1st-degree connection (src's direct directNeighbors)
	// (1) If we found dst, that's it
	// (2) Otherwise, push the neighbor into q
	r++
	q := NewQueue[T]()
	for _, directNeighbor := range directNeighbors {
		if directNeighbor == src {
			continue
		}
		if directNeighbor == dst {
			// We found dst within the first neighbor
			return r, true
		}
		q.Push(directNeighbor)
	}

	// Check 2nd-degree connection, or directNeighbor's directNeighbors.
	// We do this by first popping q to get pointer to the directNeighbor
	// And then call this func recursively again on the neighbor.
	for i, qLen := 0, q.Len(); i < qLen; i++ {
		pDirectNeighbor := q.Pop()
		if pDirectNeighbor == nil {
			continue
		}
		if r, found := isConnected(graph, *pDirectNeighbor, dst, r); found {
			return r, true
		}
	}
	return r, false
}
