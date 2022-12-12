package data

import (
	"testing"
)

func TestArraySortDirection(t *testing.T) {
	if !Ascending.IsValid() {
		t.Error("direction Ascending is invalid")
	}
	if !Descending.IsValid() {
		t.Error("direction Descending is invalid")
	}

	d := ArraySortDirection(69)
	if d.IsValid() {
		t.Errorf("invalid direction %d is valid", d)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{2, 3, 60, 1, 70, 234, -1}
	out := QuickSort(arr, Ascending)

	if out[0] != -1 {
		t.Fatal("unexpected result")
	}
	if out[len(out)-1] != 234 {
		t.Fatal("unexpected result")
	}
}
