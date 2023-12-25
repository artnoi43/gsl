package data

import (
	"math/big"
	"testing"
)

type testCaseQuickSort struct {
	inputs   []int64
	expected []int64
	ordering SortOrder
}

func defaultTestCases() []testCaseQuickSort {
	return []testCaseQuickSort{
		{
			inputs:   []int64{},
			expected: []int64{},
			ordering: Ascending,
		},
		{
			inputs:   []int64{-1},
			expected: []int64{-1},
			ordering: Ascending,
		},
		{
			inputs:   []int64{2, 2, 2},
			expected: []int64{2, 2, 2},
			ordering: Ascending,
		},
		{
			inputs:   []int64{2, 3, 60, 1, 70, 234, -1},
			expected: []int64{-1, 1, 2, 3, 60, 70, 234},
			ordering: Ascending,
		},
		{
			inputs:   []int64{2, 3, 60, 1, 70, 234, -1},
			expected: []int64{234, 70, 60, 3, 2, 1, -1},
			ordering: Descending,
		},
	}
}

func TestQuickSort(t *testing.T) {
	tests := defaultTestCases()
	for i := range tests {
		testCase := tests[i]
		testQuickSort(t, testCase)
	}

	// See if it'll overflow on 10M ints
	var s []int = make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		s[i] = i
	}

	QuickSort(s, Descending)
}

func testQuickSort(t *testing.T, testCase testCaseQuickSort) {
	result := QuickSort(testCase.inputs, testCase.ordering)

	for i := range testCase.expected {
		expected := testCase.expected[i]
		actual := result[i]

		if expected != actual {
			t.Fatalf("got unexpected value at position %d: expecting %d, got %d", i, expected, actual)
		}
	}
}

type testCaseQuickSortCmp[T CmpOrdered[T]] struct {
	inputs   []T
	expected []T
	ordering SortOrder
}

func TestQuickSortBigInts(t *testing.T) {
	testCasesInt64 := defaultTestCases()

	tests := make([]testCaseQuickSortCmp[*big.Int], len(testCasesInt64))
	for i := range testCasesInt64 {
		testCaseInt64 := &testCasesInt64[i]
		inputs := make([]*big.Int, len(testCaseInt64.inputs))
		expected := make([]*big.Int, len(testCaseInt64.expected))

		for j := range inputs {
			inputs[j] = big.NewInt(testCaseInt64.inputs[j])
		}
		for j := range expected {
			expected[j] = big.NewInt(testCaseInt64.expected[j])
		}

		tests[i] = testCaseQuickSortCmp[*big.Int]{
			inputs:   inputs,
			expected: expected,
			ordering: testCaseInt64.ordering,
		}
	}

	for i := range tests {
		testQuickSortCmp(t, tests[i])
	}
}

type foo struct {
	value int64
}

func (f *foo) Cmp(other *foo) int {
	switch {
	case f.value == other.value:
		return 0
	case f.value < other.value:
		return 1
	}

	return -1
}

func TestQuickSortCustomCmp(t *testing.T) {
	testCasesInt64 := defaultTestCases()
	tests := make([]testCaseQuickSortCmp[*foo], len(testCasesInt64))

	for i := range testCasesInt64 {
		testCaseInt64 := testCasesInt64[i]
		inputs := make([]*foo, len(testCaseInt64.inputs))
		expected := make([]*foo, len(testCaseInt64.expected))

		for j := range testCaseInt64.inputs {
			inputs[j] = &foo{
				value: testCaseInt64.inputs[j],
			}
		}
		for j := range testCaseInt64.expected {
			inputs[j] = &foo{
				value: testCaseInt64.expected[j],
			}
		}

		tests[i] = testCaseQuickSortCmp[*foo]{
			inputs:   inputs,
			expected: expected,
			ordering: testCaseInt64.ordering,
		}
	}

	for i := range tests {
		testQuickSortCmp(t, tests[i])
	}
}

func testQuickSortCmp[T CmpOrdered[T]](t *testing.T, testCase testCaseQuickSortCmp[T]) {
	result := QuickSortCmp[T](testCase.inputs, testCase.ordering)

	for i := range testCase.expected {
		expected := testCase.expected[i]
		actual := result[i]

		if cmp := expected.Cmp(actual); cmp != 0 {
			t.Logf("unexpected value at position %d - expecting %v, got %v", i, expected, actual)
			t.Fatalf("unexpected Cmp result - expecting 0, got %d", cmp)
		}
	}
}
