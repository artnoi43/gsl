package data

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := testCasesDefault()
	for i := range tests {
		testCase := tests[i]
		testSortFunc(t, testCase, QuickSort)
	}

	// See if it'll overflow on 10M ints
	var s []int = make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		s[i] = i
	}

	QuickSort(s, Descending)
}

func TestQuickSortBigInts(t *testing.T) {
	tests := testCasesBigInt()
	for i := range tests {
		testSortFuncCmp(t, tests[i], QuickSortCmp)
	}
}

func TestQuickSortCustomCmp(t *testing.T) {
	tests := testCasesFoo()
	for i := range tests {
		testSortFuncCmp[*foo](t, tests[i], QuickSortCmp)
	}
}
