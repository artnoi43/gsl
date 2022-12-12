package soyutils

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	unmarshalErrorFmtJSON = "error unmarshaling JSON file %s"
	marshalErrorFmtJSON   = "error marshaling JSON file %s"
)

// ReadFileJSON reads JSON file |filename| and parses it as T
// and returns the pointer to the result structure T value
func ReadFileJSONPointer[T any](filename string) (*T, error) {
	p, err := ReadAndUnmarshalFilePointer[T](
		filename,
		json.Unmarshal,
	)
	if err != nil {
		return nil, errors.Wrapf(err, unmarshalErrorFmtJSON, filename)
	}
	if p == nil {
		return nil, errors.Wrapf(errors.New("nil pointer unmarshaled"), unmarshalErrorFmtJSON, filename)
	}

	return p, nil
}

// ReadFileJSON reads JSON file |filename| and parses it as T
// and returns the result structure T
func ReadFileJSON[T any](filename string) (T, error) {
	t, err := ReadAndUnmarshalFile[T](
		filename,
		json.Unmarshal,
	)
	if err != nil {
		return t, errors.Wrapf(err, unmarshalErrorFmtJSON, filename)
	}

	return t, nil
}

// WriteFileJSON marshals |t| to JSON-encoded bytes and write the JSON string to file |filename|
func MarshalAndWriteFileJSON[T any](t T, filename string) error {
	if err := MarshalAndWriteToFile(t, json.Marshal, filename); err != nil {
		return errors.Wrapf(err, marshalErrorFmtJSON, filename)
	}

	return nil
}
