package data

type Valuer[T any] interface {
	GetValue() T
}
