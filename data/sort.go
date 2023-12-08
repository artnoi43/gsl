package data

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type SortOrder uint8

const (
	Ascending SortOrder = iota
	Descending
)

func badOrder(ordering SortOrder) string {
	return fmt.Sprintf("bad SortOrder %d", ordering)
}

func (d SortOrder) IsValid() bool {
	switch d {
	case Ascending, Descending:
		return true
	}

	return false
}

type LessFunc[T any] func(list []T, i, j int) bool

// LessFuncOrdered selects the appropriate comparison function to check if the elements are ordered.
// If the returned function returns true, then the elements are sorted according to its |ordering|
func LessFuncOrdered[T constraints.Ordered](ordering SortOrder) func(T, T) bool {
	switch ordering {
	case Ascending:
		return func(v1, v2 T) bool {
			return v1 <= v2
		}

	case Descending:
		return func(v1, v2 T) bool {
			return v1 >= v2
		}
	}

	panic(badOrder(ordering))
}
