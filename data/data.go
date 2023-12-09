package data

import (
	"golang.org/x/exp/constraints"
)

type Setter[T any] interface {
	SetValue(T)
}

type Getter[T any] interface {
	GetValue() T
}

type Valuer[T any] interface {
	Setter[T]
	Getter[T]
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

func NewValuer[T any](t T) Setter[T] {
	return &wrapper[T]{value: t}
}

func NewGetValuer[T any](t T) Getter[T] {
	return &wrapper[T]{value: t}
}

func ToValues[T any](getters []Getter[T]) []T {
	values := make([]T, len(getters))
	for i := range getters {
		values[i] = getters[i].GetValue()
	}

	return values
}

func FromValues[T any](values []T) []Getter[T] {
	getters := make([]Getter[T], len(values))
	for i := range getters {
		values[i] = getters[i].GetValue()
	}

	return getters
}

func MaxValuer[T constraints.Ordered](values []Getter[T]) T {
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

func MinValuer[T constraints.Ordered](values []Getter[T]) T {
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
