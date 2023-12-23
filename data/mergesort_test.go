package data

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := [][]int{
		{1},
		{1, 2, 3, -1},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{5, 3, 1, 20, -4, 8},
	}

	for i := range tests {
		out := MergeSort[int](tests[i], Ascending)

		prev := out[0]
		for j := range out {
			elem := out[j]
			if elem < prev {
				t.Fatalf("unexpected value")
			}

			prev = elem
		}
	}
}
