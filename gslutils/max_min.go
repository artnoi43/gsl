package gslutils

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](items ...T) T {
	if len(items) == 0 {
		var zeroValue T
		return zeroValue
	}

	max := items[0]
	for _, num := range items {
		if num > max {
			max = num
		}
	}

	return max
}

func Min[T constraints.Ordered](items ...T) T {
	if len(items) == 0 {
		var zeroValue T
		return zeroValue
	}

	min := items[0]
	for _, num := range items {
		if num < min {
			min = num
		}
	}

	return min
}
