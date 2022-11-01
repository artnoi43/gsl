package gslutils

func ReverseInPlace[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Reverse[T any](arr []T) []T {
	reversed := make([]T, len(arr))
	copy(reversed, arr)

	ReverseInPlace(reversed)
	return reversed
}

func Contains[T comparable](arr []T, item T) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}

	return false
}

func CollectPointers[T any](arr *[]T) []*T {
	derefArr := *arr
	out := make([]*T, len(derefArr))

	for i := range derefArr {
		out[i] = &derefArr[i]
	}

	return out
}

func CollectPointersIf[T any](arr *[]T, filterFunc func(T) bool) []*T {
	derefArr := *arr
	var filtered []*T
	for i := range derefArr {
		if filterFunc(derefArr[i]) {
			filtered = append(filtered, &derefArr[i])
		}
	}

	return filtered
}
