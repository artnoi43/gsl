package data

import (
	"math/big"
	"testing"
)

func TestArraySortDirection(t *testing.T) {
	if !Ascending.IsValid() {
		t.Error("direction Ascending is invalid")
	}
	if !Descending.IsValid() {
		t.Error("direction Descending is invalid")
	}

	d := SortOrder(69)
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

	// See if it'll overflow
	var s []int = make([]int, 1000000000)
	for i := 0; i < 1000000000; i++ {
		s[i] = i
	}

	QuickSort(arr, Descending)
}

func TestQuickSortCmp(t *testing.T) {
	ints := []int64{2, 3, 60, 1, 70, 234, -1}
	arr := make([]*big.Int, len(ints))

	for i := range ints {
		arr[i] = big.NewInt(ints[i])
	}

	out := QuickSortCmp(arr, Ascending)

	if out[0].Int64() != -1 {
		t.Fatalf("unexpected result - expecting -1, got %d", out[0].Int64())
	}

	if out[len(out)-1].Int64() != 234 {
		t.Fatalf("unexpected result - expecting 234, got %d", out[len(out)-1].Int64())
	}
}
