package data

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](arr []T, ordering SortOrder) []T {
	length := len(arr)
	lessFunc := LessFuncOrdered[T](ordering)

	switch {
	case length < 2:
		return arr

	case length == 2:
		a, b := arr[0], arr[1]

		if lessFunc(a, b) {
			return arr
		}

		return []T{b, a}

	default:
		mid := length / 2

		left := MergeSort[T](arr[:mid], ordering)
		right := MergeSort[T](arr[mid:], ordering)

		lp, rp := len(left)-1, len(right)-1
		sorted := make([]T, length)
		sp := length - 1

		for rp >= 0 && lp >= 0 {
			l, r := left[lp], right[rp]

			if lessFunc(l, r) {
				sorted[sp] = r
				rp--
				sp--

				continue
			}

			sorted[sp] = l
			lp--
			sp--
		}

		for lp >= 0 {
			sorted[sp] = left[lp]
			lp--
			sp--
		}

		for rp >= 0 {
			sorted[sp] = right[rp]
			rp--
			sp--
		}

		return sorted
	}
}
