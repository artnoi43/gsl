package data

import (
	"golang.org/x/exp/constraints"
)

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

func MaxValuer[T constraints.Ordered](values []GetValuer[T]) T {
	var t T
	if len(values) == 0 {
		return t
	}

	max := values[0].GetValue()
	for i := range values {
		v := values[i].GetValue()
		if v > max {
			max = v
		}
	}

	return max
}

func MinValuer[T constraints.Ordered](values []GetValuer[T]) T {
	var t T
	if len(values) == 0 {
		return t
	}
	min := values[0].GetValue()
	for i := range values {
		v := values[i].GetValue()
		if v < min {
			min = v
		}
	}

	return min
}
