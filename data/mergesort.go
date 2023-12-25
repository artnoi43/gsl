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

	return MergeSortedArrays(left, right, LessFuncOrdered[T](ordering))
}

func MergeSortedArrays[T constraints.Ordered](
	a []T,
	b []T,
	lessFunc func(a T, b T) bool,
) []T {
	var p, lp, rp int
	sorted := make([]T, len(a)+len(b))

	for lp < len(a) && rp < len(b) {
		l, r := a[lp], b[rp]

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

	for lp < len(a) {
		sorted[p] = a[lp]
		lp++
		p++
	}

	for rp < len(b) {
		sorted[p] = b[rp]
		rp++
		p++
	}

	return sorted
}
