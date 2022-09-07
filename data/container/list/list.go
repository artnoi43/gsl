package list

type BasicList[T any] interface {
	Push(T)
	Pop() *T
	Len() int
	IsEmpty() bool
}
