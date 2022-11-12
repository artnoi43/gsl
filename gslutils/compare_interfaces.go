package gslutils

import (
	"reflect"

	"github.com/pkg/errors"
)

var ErrNotConvertible = errors.New("types are not convertible")

// CompareInterfaceValues will compares |a| and |b| as type T.
// It returns true if |a| and |b| are of type T and are equal in value,
// or if |a| and |b| can be converted to T and are equal in value.
// If |a| and |b| are not equal but can be converted to T, then it returns false.
// CompareInterfaceValues *only* returns non-nil error if any of the values cannot be converted into T.
func CompareInterfaceValues[T comparable](a, b interface{}) (bool, error) {
	var t T
	typeTarget := reflect.TypeOf(t)

	assertedA, ok := a.(T)
	if !ok {
		typeA := reflect.TypeOf(a)
		if typeA.ConvertibleTo(typeTarget) {
			// Convert to T
			innerVal := reflect.ValueOf(a).Convert(typeTarget)
			assertedA = innerVal.Interface().(T)
		} else {
			return false, errors.Wrapf(ErrNotConvertible, "cannot convert %s to %s", typeA.String(), typeTarget.String())
		}
	}

	assertedB, ok := b.(T)
	if !ok {
		typeB := reflect.TypeOf(b)
		if typeB.ConvertibleTo(typeTarget) {
			// Convert to T
			innerVal := reflect.ValueOf(b).Convert(typeTarget)
			assertedB = innerVal.Interface().(T)
		} else {
			return false, errors.Wrapf(ErrNotConvertible, "cannot convert %s to %s", typeB.String(), typeTarget.String())
		}
	}

	return assertedA == assertedB, nil
}
