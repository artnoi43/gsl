package mathutils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Max[N Number](slice []N) N {
	var max N
	for _, num := range slice {
		if num > max {
			max = num
		}
	}
	return max
}

func Min[N Number](slice []N) N {
	if len(slice) == 0 {
		return 0
	}
	min := slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
	}
	return min
}
