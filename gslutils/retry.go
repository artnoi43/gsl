package gslutils

import (
	"strings"
	"time"

	"github.com/pkg/errors"
)

type retryConfig struct {
	attempts    int
	delay       time.Duration
	lastErrOnly bool
}

// RetryOption is a function that takes in (and modifies) *Config
type RetryOption func(*retryConfig)

// Attempts set retry attempts
func Attempts(attempts int) RetryOption {
	return func(conf *retryConfig) {
		conf.attempts = attempts
	}
}

// Delay sets retry delay
func Delay(delay time.Duration) RetryOption {
	return func(conf *retryConfig) {
		conf.delay = delay
	}
}

// LastErrorOnly sets retry option to only return the last error
func LastErrorOnly(lastErrOnly bool) RetryOption {
	return func(conf *retryConfig) {
		conf.lastErrOnly = lastErrOnly
	}
}

func retry(
	f func() error,
	opts ...RetryOption,
) error {
	var err error
	var retryErrors []error

	conf := new(retryConfig)
	for _, applyOption := range opts {
		applyOption(conf)
	}

	var lastIndex int // Index of last error
	for i := 0; i < conf.attempts; i++ {
		// Overwrite err with last error
		err = f()

		if err != nil {
			retryErrors = append(retryErrors, err)
			lastIndex = i

			time.Sleep(conf.delay)
			continue
		}

		break
	}

	// Return nil if last error is nil
	if err == nil {
		return nil
	}

	// Return only last error
	if conf.lastErrOnly {
		return retryErrors[lastIndex]
	}

	// Return all errors, concatenated.
	errorStrings := make([]string, lastIndex)
	for _, retryError := range retryErrors {
		errorStrings = append(errorStrings, retryError.Error())
	}
	err = errors.New(strings.Join(errorStrings, ", "))

	return err
}

// Retry wraps retry with action string.
func Retry(action string, f func() error, opts ...RetryOption) error {
	if err := retry(f, opts...); err != nil {
		return errors.Wrapf(err, "error when retrying %s", action)
	}

	return nil
}

// RetryWithReturn wraps |f| in a `func() error`
// and captures the T value in that function,
// and returns T returned by |f|.
func RetryWithReturn[T any](
	action string,
	f func() (T, error),
	opts ...RetryOption,
) (
	T,
	error,
) {
	var t T
	var err error

	err = retry(func() error {
		t, err = f()
		if err != nil {
			return err
		}

		return nil
	},
		opts...,
	)

	if err != nil {
		return t, errors.Wrapf(err, "error when retrying %s", action)
	}

	return t, nil
}
