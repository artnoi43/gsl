package gslutils

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

func StringerToUpperString(s stringer) string {
	return strings.ToUpper(s.String())
}

func StringerToLowerString(s stringer) string {
	return strings.ToLower(s.String())
}

func StringerToUpper[T stringer](s T) T {
	return T(StringerToUpper(s))
}

func StringerToLower[T stringer](s T) T {
	return T(StringerToLower(s))
}
