package gslutils

import "testing"

func TestMaxMin(t *testing.T) {
	min := float64(1)
	max := float64(100)
	floats := []float64{69, min, 3, 4, max}
	if Max(floats...) != max {
		t.Fatal("unexpected Max result")
	}
	if r := Min(floats...); r != min {
		t.Fatalf("unexpected Min result %v", r)
	}
}
