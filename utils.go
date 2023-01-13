package gsl

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type GoNumber interface {
	constraints.Integer | constraints.Float
}

// GroupConsecutive *sorts* slice |s| in-place,
// and calls GroupConsecutiveSorted on |s|.
func GroupConsecutive[N GoNumber](s []N) [][2]N {
	sort.Slice(s, func(i, j int) bool {
		return i < j
	})

	return GroupConsecutiveSorted(s)
}

// GroupConsecutiveSorted groups input numbers in slice []N |s|
// that are consecutive to each other (i.e. difference of 1) into a slice of [2]N.
// The result type [2]N represents the start and end of a consecutive range,
// that is, result[0] is "from" while result[1] is "to" of such ranges.
//
// If a breakpoint (non-consecutive element) is found, GroupConsecutive creates a new result.
// e.g. [1, 2, 3, 5, 6, 8, 9, 10] will be mapped to [{1, 3}, {5, 6}, {8, 10}].
//
// If |s| has length of 0, it returns [][2]N{ {0, 0} }.
// If |s| has length of 1 and is []N{n}, it returns [][2]N{ {n, n} }.
func GroupConsecutiveSorted[N GoNumber](s []N) [][2]N {
	l := len(s)

	switch l {
	case 0:
		return [][2]N{{0, 0}}
	case 1:
		n := s[0]
		return [][2]N{{n, n}}
	}

	var consecs [][2]N

	for i := 0; i < l; i++ {
		var curr N
		var prev N

		curr = s[i]

		// First element - init new default consec range [curr, curr].
		if i == 0 {
			consecs = append(consecs, [2]N{curr, curr})
			continue
		}

		prev = s[i-1]

		// Skip duplicate member
		if prev == curr {
			continue
		}

		// If breakpoint found, add new default consec range [curr, curr].
		if prev != curr-1 {
			consecs = append(consecs, [2]N{curr, curr})
			continue
		}

		consecs[len(consecs)-1][1] = curr
	}

	return consecs
}
