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
		// These are 'the majority', a closed, well-connected group of people
		art = "art"
		bob = "bob"
		liz = "liz"
		mac = "mac"
		ron = "ron"
		may = "may"
		// weired0 helps connect the majority to aloof via may,
		// and weird1 bridges weird0 to aloof
		weird0 = "weird0"
		weird1 = "weird1"
		aloof  = "aloof"
		// abe and chad just have no friends
		abe  = "abe"
		chad = "chad"
	)

	graph := make(map[string][]string)
	graph[art] = []string{bob, liz}
	graph[bob] = []string{art, liz, mac}
	graph[liz] = []string{art, bob, mac, may}
	graph[mac] = []string{bob, liz}
	graph[ron] = []string{may}
	graph[may] = []string{liz, ron, weird0}
	// weird0 links may to weird1
	graph[weird0] = []string{may, weird1}
	graph[weird1] = []string{weird0, aloof}
	graph[aloof] = []string{weird1}
	// abe and chad is not connected to the majority
	graph[abe] = []string{chad}
	graph[chad] = []string{abe}

	t.Logf("graph %+v\n", graph)
	pront(t, graph, art, liz, true)
	pront(t, graph, art, bob, true)
	pront(t, graph, art, mac, true)
	pront(t, graph, art, may, true)
	pront(t, graph, art, abe, false) // expects not found
}

// searchBreadthFirst returns the number of hops needed to look from node src to dst and whether dst is found.
func searchBreadthFirst[T comparable](mapGraph map[T][]T, src, dst T) bool {
	// Copy graph
	graph := make(map[T][]T, len(mapGraph))
	for k, v := range mapGraph {
		graph[k] = v
	}
	// Get first neigbors and delete from map
	firstNeighbors := graph[src]
	delete(graph, src)

	var found bool
	var searched = make(map[T]bool)

	if src == dst {
		return true
	}
	searched[src] = true

	if firstNeighbors == nil {
		return false
	}

	// Prepare for 1st hop
	q := NewQueue[T]()
	q.PushSlice(firstNeighbors...)

	for {
		if q.IsEmpty() {
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
		searched[thisNeighbor] = true

		if thisNeighbor == dst {
			found = true
			continue
		}
		// Push nth-degree connections to queue
		q.PushSlice(graph[thisNeighbor]...)
	}

	return found
}

func pront[T comparable](t *testing.T, g map[T][]T, src, dst T, expectToFind bool) {
	t.Log(src, "->", dst)
	found := searchBreadthFirst(g, src, dst)
	if found != expectToFind {
		t.Fatalf("unexpected result 'found' - expected %v, got %v\n", expectToFind, found)
	}
	t.Logf("shortestPath %v->%v found: %v\n", src, dst, found)
}
