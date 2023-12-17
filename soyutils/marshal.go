package soyutils

import (
	"fmt"
	"os"
	"reflect"

	"github.com/pkg/errors"
)

type (
	marshalFunc   func(interface{}) ([]byte, error)
	unmarshalFunc func([]byte, interface{}) error
)

func ReadAndUnmarshalFilePointer[T any](filename string, f unmarshalFunc) (*T, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file to unmarshal")
	}

	var t T
	if err := f(b, &t); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal file %s content", filename)
	}

	return &t, nil
}

func ReadAndUnmarshalFile[T any](filename string, f unmarshalFunc) (T, error) {
	p, err := ReadAndUnmarshalFilePointer[T](filename, f)
	if err != nil {
		var t T
		return t, err
	}
	if p == nil {
		var t T
		return t, fmt.Errorf("got nil pointer after unmarshaling file %s", filename)
	}

	return *p, nil
}

func MarshalAndWriteToFile[T any](t T, f marshalFunc, filename string) error {
	b, err := f(t)
	if err != nil {
		return errors.Wrapf(err, "error marshaling type %s", reflect.TypeOf(t).String())
	}

	fp, err := os.Create(filename)
	defer fp.Close() //nolint:staticcheck

	if err != nil {
		return errors.Wrapf(err, "error creating file %s", filename)
	}

	_, err = fp.Write(b)
	if err != nil {
		return errors.Wrapf(err, "error writing to file %s", filename)
	}

	return nil
}
