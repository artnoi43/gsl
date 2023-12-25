package data

import (
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](arr []T, ordering SortOrder) []T {
	l := len(arr)
	// Base case
	if l < 2 {
		return arr
	}

	// Pivot here is going to be middle elem
	mid := l / 2
	pivot := arr[mid]

	// pivoted is the list without the |pivot| element/member
	pivoted := append(arr[:mid], arr[mid+1:]...) //nolint:gocritic
	var left, right []T                          //nolint:prealloc

	{
		// isLess should have lifetime of no more than the for-looop
		// to help minimize stack size in huge recursive calls
		isLess := LessFuncOrdered[T](ordering)
		for _, elem := range pivoted {
			if isLess(elem, pivot) {
				left = append(left, elem)
				continue
			}

			right = append(right, elem)
		}
	}

	sorted := append(QuickSort(left, ordering), pivot)
	sorted = append(sorted, QuickSort(right, ordering)...)

	return sorted
}

func QuickSortCmp[T CmpOrdered[T]](arr []T, ordering SortOrder) []T {
	l := len(arr)
	// Base case
	if l < 2 {
		return arr
	}

	// Pivot here is going to be middle elem
	mid := l / 2
	pivot := arr[mid]

	// pivoted is the list without the pivot element/member
	pivoted := append(arr[:mid], arr[mid+1:]...) //nolint: gocritic
	var left, right []T                          //nolint:prealloc

	cmpLess := -1
	if ordering == Descending {
		cmpLess = 1
	}

	for _, elem := range pivoted {
		if elem.Cmp(pivot) == cmpLess {
			left = append(left, elem)
			continue
		}

		right = append(right, elem)
	}

	sorted := append(QuickSortCmp(left, ordering), pivot)
	sorted = append(sorted, QuickSortCmp(right, ordering)...)

	return sorted
}
