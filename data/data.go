package data

type Valuer[T any] interface {
	Value() T
}

type BasicList[T any] interface {
	Pop() *T
	Push(T)
	Len() int
	IsEmpty() bool
}
