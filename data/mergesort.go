package data

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](arr []T, ordering SortOrder) []T {
	length := len(arr)

	if length < 2 {
		return arr
	}

	if length == 2 {
		a, b := arr[0], arr[1]

		lessFunc := LessFuncOrdered[T](ordering)

		if lessFunc(a, b) {
			return arr
		}

		return []T{b, a}
	}

	mid := length / 2

	left := MergeSort[T](arr[:mid], ordering)
	right := MergeSort[T](arr[mid:], ordering)

	sorted := make([]T, length)
	var sp, lp, rp int

	lessFunc := LessFuncOrdered[T](ordering)
	for lp < len(left) && rp < len(right) {
		l, r := left[lp], right[rp]

		if lessFunc(l, r) {
			sorted[sp] = l
			lp++
			sp++

			continue
		}

		sorted[sp] = r
		rp++
		sp++
	}

	for lp < len(left) {
		sorted[sp] = left[lp]
		lp++
		sp++
	}

	for rp < len(right) {
		sorted[sp] = right[rp]
		rp++
		sp++
	}

	return sorted
}
