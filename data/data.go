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

type wrapper[T any] struct {
	value T
}

func (w *wrapper[T]) SetValue(value T) {
	w.value = value
}

func (w *wrapper[T]) GetValue() T {
	return w.value
}

func NewValuer[T any](t T) SetValuer[T] {
	return &wrapper[T]{value: t}
}

func NewGetValuer[T any](t T) GetValuer[T] {
	return &wrapper[T]{value: t}
}
