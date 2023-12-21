package gsl

import (
	"os"
	"reflect"
	"testing"
)

var envsCi [2]string = [2]string{"CI", "GITHUB_ACTIONS"}

func TestGroupConsecutive(t *testing.T) {
	ci := false
	for i := range envsCi {
		s, found := os.LookupEnv(envsCi[i])
		if found && s != "" {
			ci = true
			break
		}
	}

	if !ci {
		t.Skip("Skipping tests due to test not being run on CI platform")
		return
	}

	testGroupConsecutive(t)
}

func testGroupConsecutive(t *testing.T) {
	type testCase[T GoNumber] struct {
		in  []T
		out [][2]T
	}

	testCases := []testCase[uint64]{
		{
			in:  []uint64{1, 2, 3, 5, 6, 7, 9},
			out: [][2]uint64{{1, 3}, {5, 7}, {9, 9}},
		},
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
		{
			in:  []uint64{1, 1, 1, 1, 1, 1, 1, 7},
			out: [][2]uint64{{1, 1}, {7, 7}},
		},
		{
			in:  []uint64{1, 2, 2, 2, 3, 4, 5, 10, 11, 12, 13, 13, 14, 14, 15, 17, 19, 20},
			out: [][2]uint64{{1, 5}, {10, 15}, {17, 17}, {19, 20}},
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

	// See if it'll overflow
	n := 100000
	var s []int = make([]int, n)
	for i := int(0); i < n; i++ {
		s[i] = 7
	}

	ranges := GroupConsecutive(s)
	if l := len(ranges); l != 1 {
		t.Errorf("expecting length 1, got %d", l)
	}
}
