package str

import "strings"

func ToUpper[T ~string](s T) T {
	return T(strings.ToUpper(string(s)))
}

func ToLower[T ~string](s T) T {
	return T(strings.ToLower(string(s)))
}
