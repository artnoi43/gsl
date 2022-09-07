package data

type ArraySortDirection uint8

const (
	Ascending ArraySortDirection = iota
	Descending
)

type Valuer[T any] interface {
	GetValue() T
}

type Set[T comparable] interface {
	HasDuplicate(x T) bool
}
