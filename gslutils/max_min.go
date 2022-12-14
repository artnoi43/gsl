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

func Sum[T constraints.Integer | constraints.Float](items ...T) T {
	var sum T
	for _, item := range items {
		sum += item
	}

	return sum
}

func Avg[T constraints.Integer | constraints.Float](items ...T) T {
	return Sum(items...) / T(len(items))
}
