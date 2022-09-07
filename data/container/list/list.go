package list

type BasicList[T any] interface {
	Push(x T)
	PushSlice(values []T)
	Pop() *T
	Len() int
	IsEmpty() bool
}

// Set[T] extends BasicList[T] with HasDuplicate(x T).
// Default implementation is SetImpl[T]
type Set[T comparable] interface {
	BasicList[T]
	HasDuplicate(x T) bool
}
