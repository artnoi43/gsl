package data

import (
	"golang.org/x/exp/constraints"
)

func QuickSortValuer[T constraints.Ordered](
	arr []GetValuer[T],
	direction ArraySortDirection,
) []GetValuer[T] {

	if !direction.IsValid() {
		panic("invalid direction")
	}

	l := len(arr)
	// Base case
	if l < 2 {
		return arr
	}

	// Pivot here is going to be middle elem
	mid := l / 2
	pivot := arr[mid]
	var left []GetValuer[T]
	var right []GetValuer[T]

	// sansPivot is the list without the pivot element/member
	sansPivot := append(arr[:mid], arr[mid+1:]...)

	for _, elem := range sansPivot {
		switch direction {
		case Ascending:
			if elem.GetValue() >= pivot.GetValue() {
				right = append(right, elem)
				continue
			}

		case Descending:
			if elem.GetValue() <= pivot.GetValue() {
				right = append(right, elem)
				continue
			}
		}

		left = append(left, elem)
	}

	sorted := append(QuickSortValuer(left, direction), pivot)
	sorted = append(sorted, QuickSortValuer(right, direction)...)

	return sorted
}

func QuickSort[T constraints.Ordered](
	arr []T,
	direction ArraySortDirection,
) []T {

	if !direction.IsValid() {
		panic("invalid connection")
	}

	l := len(arr)
	// Base case
	if l < 2 {
		return arr
	}

	// Pivot here is going to be middle elem
	mid := l / 2
	pivot := arr[mid]

	var left []T
	var right []T

	// sansPivot is the list without the pivot element/member
	sansPivot := append(arr[:mid], arr[mid+1:]...)

	for _, elem := range sansPivot {
		switch direction {
		case Ascending:
			if elem >= pivot {
				right = append(right, elem)
				continue
			}
		case Descending:
			if elem <= pivot {
				right = append(right, elem)
				continue
			}
		}

		left = append(left, elem)
	}

	sorted := append(QuickSort(left, direction), pivot)
	sorted = append(sorted, QuickSort(right, direction)...)

	return sorted
}
