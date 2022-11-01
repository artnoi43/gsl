package gslutils

import (
	"testing"
)

func TestReverseSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expecteds := []int{4, 3, 2, 1}

	reversed := Reverse(arr)
	for i, item := range reversed {
		if item == arr[i] {
			t.Fatal("not copied but in-place")
		}
		if expected := expecteds[i]; item != expected {
			t.Fatalf("unexpected value, expecting %v, got %v", expected, item)
		}
	}

	ReverseInPlace(arr)
	for i, item := range arr {
		if expected := expecteds[i]; item != expected {
			t.Fatalf("unexpected value, expecting %v, got %v", expected, item)
		}
	}
}

func TestContains(t *testing.T) {
	i := 69
	arr0 := []int{68, 69, 70, 71}
	arr1 := []int{68, 70, 71, 72}

	if !Contains(arr0, i) {
		t.Log("item", i)
		t.Log("arr", arr0)
		t.Fatalf("unexpected Contains result -- expected true")
	}
	if Contains(arr1, i) {
		t.Log("item", i)
		t.Log("arr", arr1)
		t.Fatalf("unexpected Contains result -- expected false")
	}
}
