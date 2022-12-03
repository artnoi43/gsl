package gslutils

import (
	"reflect"
	"testing"
)

func TestCopySlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 20, 69}
	out := CopySlice(arr)

	if !reflect.DeepEqual(arr, out) {
		t.Log("original", arr)
		t.Log("copy", out)
		t.Fatal("unexpected copy result")
	}
}

func TestReverseSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expecteds := []int{4, 3, 2, 1}

	reversed := Reverse(arr)
	for i, item := range reversed {
		if item == arr[i] {
			t.Fatal("not copied but in-place")
		}
		if expected := expecteds[i]; item != expected {
			t.Fatalf("unexpected value -- expecting %v, got %v", expected, item)
		}
	}

	ReverseInPlace(arr)
	for i, item := range arr {
		if expected := expecteds[i]; item != expected {
			t.Fatalf("unexpected value -- expecting %v, got %v", expected, item)
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
		t.Fatalf("unexpected Contains result -- expecting true")
	}
	if Contains(arr1, i) {
		t.Log("item", i)
		t.Log("arr", arr1)
		t.Fatalf("unexpected Contains result -- expecting false")
	}
}

func TestFilterSlice(t *testing.T) {
	arr := []int{-1, -2, 9, 69, 0}
	actual := FilterSlice(arr, func(i int) bool {
		return i <= 0
	})
	expecteds := []int{-1, -2, 0}

	for i, elem := range actual {
		if expected := expecteds[i]; elem != expected {
			t.Fatalf("unexpected value -- expecting %d, got %d", expected, elem)
		}
	}
}

func TestCollectPointers(t *testing.T) {
	arr := []int{1, 2, 6}
	out := CollectPointers(arr)

	for i := range arr {
		if &arr[i] != out[i] {
			t.Fatalf("wrong pointer collected -- expected %p found %p\n", &arr[i], out[i])
		}
	}

	out = CollectPointersIf(arr, func(i int) bool {
		return i > 5 // Only 6 should be filtered
	})

	if l := len(out); l != 1 {
		t.Fatalf("wrong pointer filtered -- expecting 1 result, got %d\n", l)
	}

	if arr[len(arr)-1] != *out[0] {
		t.Fatalf("wrong pointer filtered -- invalid pointer")
	}
}

func TestDerefValues(t *testing.T) {
	type foo string

	f1 := foo("henlo")
	f2 := foo("lol")
	f3 := foo("kuyhee")

	arr := []*foo{&f1, &f2, nil, &f3}

	// Test DerefValues
	expecteds := []foo{f1, f2, foo(""), f3}
	values := DerefValues(arr)
	for i, actual := range values {
		expected := expecteds[i]
		if actual != expected {
			t.Fatalf("unexpected value -- expecting %s, got %s", expected, actual)
		}
	}

	// Test DerefValuesIf
	expecteds = []foo{f1, f3}
	values = DerefValuesIf(arr, func(elem foo) bool {
		return len(elem) > 3
	})
	for i, actual := range values {
		expected := expecteds[i]
		if actual != expected {
			t.Fatalf("unexpected value -- expecting %s, got %s", expected, actual)
		}
	}
}
