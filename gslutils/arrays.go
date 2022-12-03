package gslutils

func CopySlice[T any](arr []T) []T {
	ret := make([]T, len(arr))
	copy(ret, arr)

	return ret
}

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

// Contains iterates over |arr| and returns true if an element in |arr| is equal to |item|.
func Contains[T comparable](arr []T, item T) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}

	return false
}

// FilterSlice iterates over |arr| and calls |filterFunc| on each of the element |elem|.
// If filterFunc returns true, |elem| is added to the return slice.
func FilterSlice[T any](arr []T, filterFunc func(elem T) bool) []T {
	if arr == nil {
		return nil
	}

	var filtered []T
	for _, elem := range arr {
		if filterFunc(elem) {
			filtered = append(filtered, elem)
		}
	}

	return filtered
}

// CollectPointers iterates over |arr| and returns a new slice containing all references
// to the elements of |arr|.
func CollectPointers[T any](arr []T) []*T {
	if arr == nil {
		return nil
	}

	out := make([]*T, len(arr))
	for i := range arr {
		out[i] = &arr[i]
	}

	return out
}

// CollectPointersIf iterates over |arr| and calling filterFunc on the element |elem|.
// If filterFunc returns true, |elem| is added to the return slice.
func CollectPointersIf[T any](arr []T, filterFunc func(elem T) bool) []*T {
	if arr == nil {
		return nil
	}
	var filtered []*T
	for i := range arr {
		value := arr[i]
		if filterFunc(value) {
			filtered = append(filtered, &value)
		}
	}

	return filtered
}

// DerefValues iterates over |arr| and return a new slice containing dereferenced values
// of the elements in |arr|. If an element in |arr| is nil at index `i`, the value of the
// return slice at index `i` is a default value for type T.
func DerefValues[T any](arr []*T) []T {
	if arr == nil {
		return nil
	}

	var t T // Default value

	values := make([]T, len(arr))
	for i := range arr {
		if arr[i] == nil {
			values[i] = t
			continue
		}

		values[i] = *arr[i]
	}

	return values
}

// DerefValuesIf iterates over |arr| and calls |filterFunc| on the element |elem|.
// If filterFunc returns true, |elem|'s dereferenced value is added to the return slice.
// Unlike DerefValues but similar to CollectPointersIf, if |elem| is nil,
// its value will not be added to the return slice
func DerefValuesIf[T any](arr []*T, filterFunc func(elem T) bool) []T {
	if arr == nil {
		return nil
	}

	var filtered []T
	for _, elem := range arr {
		if elem == nil {
			continue
		}

		value := *elem
		if filterFunc(value) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}
