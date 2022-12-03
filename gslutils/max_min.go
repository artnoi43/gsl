package gslutils

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](items ...T) T {
	var max T
	for _, num := range items {
		if num > max {
			max = num
		}
	}
	return max
}

func Min[N constraints.Ordered](items ...N) N {
	var zeroValue N
	if len(items) == 0 {
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
