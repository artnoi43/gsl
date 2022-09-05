package data

import (
	"testing"
)

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
