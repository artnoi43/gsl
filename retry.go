package gsl

import (
	"strings"
	"time"

	"github.com/pkg/errors"
)

type retryConfig struct {
	attempts    int
	delay       time.Duration
	lastErrOnly bool
	stopOnErr   error
}

// RetryOption is a function that takes in (and modifies) *Config
type RetryOption func(*retryConfig)

var ErrRetry = errors.New("error when retrying")

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

// StopOnError sets a specific error, which, if seen when retrying, breaks the retry loop.
// If |err| is found during retry, the loop breaks, and it is treated just like any other error,
// i.e. the retry still fails.
func StopOnError(err error) RetryOption {
	return func(conf *retryConfig) {
		conf.stopOnErr = err
	}
}

// retry does not wrap any error, but it does collect multiple errors.
func retry(
	f func() error,
	opts ...RetryOption,
) error {
	conf := new(retryConfig)
	for _, applyOption := range opts {
		applyOption(conf)
	}

	var retryErrors []error // Attempt errors
	var err error           // Current error
	var lastIndex int       // Index of last error

	for i := 0; i < conf.attempts; i++ {
		// Overwrite err with last error
		err = f()

		if err != nil {
			retryErrors = append(retryErrors, err)
			lastIndex = i

			if errors.Is(conf.stopOnErr, err) {
				break
			}

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
	errorStrings := make([]string, lastIndex+1)
	for i, retryErr := range retryErrors {
		errorStrings[i] = retryErr.Error()
	}

	return errors.New(strings.Join(errorStrings, ", "))
}

// Retry wraps retry with action string.
func Retry(action string, f func() error, opts ...RetryOption) error {
	if err := retry(f, opts...); err != nil {
		return wrapErrRetry(action, err)
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
		return t, wrapErrRetry(action, err)
	}

	return t, nil
}

func wrapErrRetry(action string, err error) error {
	retryErr := errors.Wrapf(errors.New(action), ErrRetry.Error())

	return errors.Wrap(err, retryErr.Error())
}
