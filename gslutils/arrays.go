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
