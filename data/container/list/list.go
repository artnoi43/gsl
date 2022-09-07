package list

type BasicList[T any] interface {
	Push(T)
	Pop() *T
	Len() int
	IsEmpty() bool
}

type Set[T comparable] interface {
	BasicList[T]
	HasDuplicate(x T) bool
}
