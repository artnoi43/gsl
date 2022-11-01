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

func TestCollectPointers(t *testing.T) {
	arr := []int{1, 2, 6}
	out := CollectPointers(&arr)

	for i := range arr {
		if &arr[i] != out[i] {
			t.Fatalf("wrong pointer collected - expected %p found %p\n", &arr[i], out[i])
		}
	}

	out = CollectPointersIf(&arr, func(i int) bool {
		return i > 5 // Only 6 should be filtered
	})

	if l := len(out); l != 1 {
		t.Fatalf("wrong pointer filtered - expecting 1 result, got %d\n", l)
	}

	if &arr[len(arr)-1] != out[0] {
		t.Fatalf("wrong pointer filtered - invalid pointer")
	}
}
