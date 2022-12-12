package soyutils

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	unmarshalErrorFmtYAML = "error unmarshaling YAML file %s"
	marshalErrorFmtYAML   = "error marshaling YAML file %s"
)

// ReadFileYAML reads YAML file |filename| and parses it as T
// and returns the pointer to the result structure T
func ReadFileYAMLPointer[T any](filename string) (*T, error) {
	p, err := ReadAndUnmarshalFilePointer[T](filename, yaml.Unmarshal)
	if err != nil {
		return nil, errors.Wrapf(err, unmarshalErrorFmtYAML, filename)
	}
	if p == nil {
		return nil, errors.Wrapf(errors.New("nil pointer unmarshaled"), unmarshalErrorFmtYAML, filename)
	}

	return p, nil
}

// ReadFileYAML reads YAML file |filename| and parses it as T
// and returns the result structure T
func ReadFileYAML[T any](filename string) (T, error) {
	t, err := ReadAndUnmarshalFile[T](filename, yaml.Unmarshal)
	if err != nil {
		return t, err
	}

	return t, nil
}

// WriteFileJSON marshals |t| to YAML-encoded bytes and write the YAML string to file |filename|
func MarshalAndWriteFileYAML[T any](t T, filename string) error {
	if err := MarshalAndWriteToFile(t, yaml.Marshal, filename); err != nil {
		return errors.Wrapf(err, marshalErrorFmtYAML, filename)
	}

	return nil
}
