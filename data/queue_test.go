package data

import "testing"

func TestQueue(t *testing.T) {

	values := []interface{}{0, 1, 2, "last"}
	q := NewQueue()

	// Test Push
	for _, expected := range values {
		q.Push(expected)
		v := q.Pop()
		if v != expected {
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
		if popped != expected {
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
