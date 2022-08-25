package data

import "testing"

func TestStack(t *testing.T) {

	values := []interface{}{0, 1, 2, "last"}
	stack := NewStack()

	// Test Push
	for _, expected := range values {
		stack.Push(expected)
		v := stack.Pop()
		if v != expected {
			t.Fatalf("Stack.Pop failed - expected %v, got %v", expected, v)
		}
	}

	// Test Len
	for i, v := range values {
		stack.Push(v)

		if newLen := stack.Len(); newLen != i+1 {
			t.Fatalf("Stack.Len failed - expected %d, got %d", newLen, i+1)
		}
	}

	// Test Pop
	for i, qLen := 0, stack.Len(); i < qLen; i++ {
		popped := stack.Pop()
		expected := values[qLen-i-1]
		if popped != expected {
			t.Fatalf("Stack.Pop failed - expected %v, got %v", expected, popped)
		}
	}
}
