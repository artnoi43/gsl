package mglutils

import "testing"

func TestMaxMin(t *testing.T) {
	floats := []float64{1, 2, 3, 4, 5}
	if Max(floats) != 5 {
		t.Fatal("unexpected Max result")
	}
	if r := Min(floats); r != 1 {
		t.Fatalf("unexpected Min result %v", r)
	}
}
