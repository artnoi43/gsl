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

	left := MergeSort(arr[:mid], ordering)
	right := MergeSort(arr[mid:], ordering)

	sorted := make([]T, length)

	// p points to insertion point in sorted,
	// lp and rp are current read heads for left and right
	var p, lp, rp int

	// Merge left and right into sort until either one is used up,
	// then blindly fill sorted with the remainers
	lessFunc := LessFuncOrdered[T](ordering)
	for lp < len(left) && rp < len(right) {
		l, r := left[lp], right[rp]

		if lessFunc(l, r) {
			sorted[p] = l
			lp++
			p++

			continue
		}

		sorted[p] = r
		rp++
		p++
	}

	for lp < len(left) {
		sorted[p] = left[lp]
		lp++
		p++
	}

	for rp < len(right) {
		sorted[p] = right[rp]
		rp++
		p++
	}

	return sorted
}
