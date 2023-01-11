package gsl

import "strings"

type stringer interface {
	String() string
}

func ToUpper[T ~string](s T) T {
	return T(strings.ToUpper(string(s)))
}

func ToLower[T ~string](s T) T {
	return T(strings.ToLower(string(s)))
}

// StringerToUpperString calls String() on s
// and returns the result of calling `strings.ToUpper` on that string.
func StringerToUpperString(s stringer) string {
	return strings.ToUpper(s.String())
}

// StringerToLowerString calls String() on s
// and returns the result of calling `strings.ToLower` on that string.
func StringerToLowerString(s stringer) string {
	return strings.ToLower(s.String())
}
