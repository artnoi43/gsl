package gsl

import (
	"reflect"
	"testing"
)

func TestGroupConsecutive(t *testing.T) {
	type testCase struct {
		in  []uint64
		out [][2]uint64
	}
	testCases := []testCase{
		{
			in:  []uint64{1, 2, 3, 4, 6, 7, 8, 9, 69, 70, 72},
			out: [][2]uint64{{1, 4}, {6, 9}, {69, 70}, {72, 72}},
		},
		{
			in:  []uint64{1, 1, 2, 3, 4, 10, 11, 12, 13, 100}, // Duplicate 1s
			out: [][2]uint64{{1, 4}, {10, 13}, {100, 100}},
		},
		{
			in:  []uint64{1, 1, 1, 3, 4, 10, 11, 11, 11, 12, 13, 100}, // Duplicate 1s and 11s
			out: [][2]uint64{{1, 1}, {3, 4}, {10, 13}, {100, 100}},
		},
	}

	for caseNum, tc := range testCases {
		ranges := GroupConsecutive(tc.in)
		if lr, lo := len(ranges), len(tc.out); lr != lo {
			t.Log("expected", tc.out, "actual", ranges)
			t.Errorf("[%d] len output not match, expecting %d, got %d", caseNum, lr, lo)
			continue
		}

		for i := range ranges {
			actual := ranges[i]
			expected := tc.out[i]
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("[%d] unexpected result, expecting %v, got %v", caseNum, expected, actual)
			}
		}
	}
}
