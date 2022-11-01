package soyutils

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// ReadFileYAML reads YAML file |filename| and parses it as T
// and returns the result
func ReadFileYAML[T any](filename string) (T, error) {
	var t T
	b, err := os.ReadFile(filename)
	if err != nil {
		return t, errors.Wrap(err, "failed to read config")
	}

	if err := yaml.Unmarshal(b, &t); err != nil {
		return t, errors.Wrap(err, "failed to parse config")
	}

	return t, nil
}

// ReadFileYAML reads YAML file |filename| and parses it as T
// and returns the pointer to the result
func ReadFileYAMLPointer[T any](filename string) (*T, error) {
	value, err := ReadFileYAML[T](filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed get value")
	}

	return &value, nil
}
