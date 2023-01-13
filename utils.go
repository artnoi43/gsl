package gsl

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// GroupConsecutive collect sorts and groups block numbers in slice []I
// that are consecutive (difference of 1) to each other into a slice [][2]I,
// e.g. [1, 2, 3, 5, 6, 8, 9, 10] will be mapped to [{1, 3}, {5, 6}, {8, 10}].
// Result[0] is "from", i.e. the start of the consecutive range, while Result[1] is "to", the end of such range.
//
// If |s| has length of 0, it returns [][2]I{ {0, 0} }
// If |s| has length of 1 and is []I{n}, it returns [][2]I{ {n, n} }
func GroupConsecutive[I interface {
	constraints.Integer | constraints.Float
}](s []I) [][2]I {
	var zero I

	l := len(s)
	switch l {
	case 0:
		return [][2]I{{zero, zero}}
	case 1:
		n := s[0]
		return [][2]I{{n, n}}
	}

	sort.Slice(s, func(i, j int) bool {
		return i < j
	})

	var ranges [][2]I

	for i := 0; i < l; i++ {
		var curr I
		var prev I

		curr = s[i]

		if i != 0 {
			prev = s[i-1]
		} else if i == 0 {
			ranges = append(ranges, [2]I{curr, curr})
		}

		// Skip duplicate member
		if prev == curr {
			continue
		}

		currRange := &ranges[len(ranges)-1]

		// Not yet last element
		if i != l-1 {
			// Next element to |curr|
			next := s[i+1]

			if curr+1 == next {
				// Check if previous element a break point
				if prev != curr-1 {
					// Found break point, add new ranges member
					ranges = append(ranges, [2]I{curr, curr})
					continue
				}
			}

			// Update "to"
			currRange[1] = curr

		} else if i == l-1 {
			if curr-1 == prev {
				// Update "to"
				currRange[1] = curr
			} else {
				// Found break point, add new ranges member
				ranges = append(ranges, [2]I{curr, curr})
			}
		}
	}

	return ranges
}
