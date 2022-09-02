package mglutils

import "testing"

func TestReverseSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expecteds := []int{4, 3, 2, 1}

	arr = Reverse(arr)
	for i, item := range arr {
		if expected := expecteds[i]; item != expected {
			t.Fatalf("unexpected value, expecting %v, got %v", expected, item)
		}
	}
}
