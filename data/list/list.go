package list

type BasicList[T any] interface {
	Pop() *T
	Push(T)
	Len() int
	IsEmpty() bool
}
