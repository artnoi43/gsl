package data

import (
	"golang.org/x/exp/constraints"
)

func QuickSortGetter[T constraints.Ordered](
	arr []Getter[T],
	ordering SortOrder,
) []Getter[T] {
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
	var left, right []Getter[T]                  //nolint:prealloc

	{
		// isLess should have lifetime of no more than the for-looop
		// to help minimize stack size in huge recursive calls
		isLess := LessFuncOrdered[T](ordering)
		for _, elem := range pivoted {
			if isLess(elem.GetValue(), pivot.GetValue()) {
				left = append(left, elem)
				continue
			}

			right = append(right, elem)
		}
	}

	sorted := append(QuickSortGetter(left, ordering), pivot)
	sorted = append(sorted, QuickSortGetter(right, ordering)...)

	return sorted
}

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
