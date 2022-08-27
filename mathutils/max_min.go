package mathutils

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](slice []T) T {
	var max T
	for _, num := range slice {
		if num > max {
			max = num
		}
	}
	return max
}

func Min[N constraints.Ordered](slice []N) N {
	var zeroValue N
	if len(slice) == 0 {
		return zeroValue
	}
	min := slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
	}
	return min
}
