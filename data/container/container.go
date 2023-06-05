package container

import "github.com/soyart/gsl/data"

type BasicContainer[T any] interface {
	Push(x T)
	Pop() T // Pop returns the right-most value and removing it from the Set
	Len() int

	GetAt(idx int) T  // Get returns a value at index idx
	RemoveAt(idx int) // RemoveAt removes value at index idx if it exists from the slice
}

type Set[T comparable] interface {
	BasicContainer[T]
	data.Set[T]
}
