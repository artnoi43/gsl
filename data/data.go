package data

type SetValuer[T any] interface {
	SetValue(T)
}

type GetValuer[T any] interface {
	GetValue() T
}

type Valuer[T any] interface {
	SetValuer[T]
	GetValuer[T]
}

type Set[T comparable] interface {
	HasDuplicate(x T) bool
}
