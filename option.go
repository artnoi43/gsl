package gsl

import "github.com/pkg/errors"

// Rust style option
type Option[T any] *T

func OptionSome[T any](t T) Option[T] {
	return &t
}

func OptionNone[T any]() Option[T] {
	return nil
}

func OptionIsNone[T any](o Option[T]) bool {
	return o == nil
}

func OptionIsSome[T any](o Option[T]) bool {
	return !OptionIsNone(o)
}

func OptionValue[T any](o Option[T], result *T) error {
	if o == nil {
		return errors.New("option is nil")
	}

	*result = *o
	return nil
}
