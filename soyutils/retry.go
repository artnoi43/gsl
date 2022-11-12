package soyutils

import (
	"github.com/avast/retry-go"
	"github.com/pkg/errors"
)

func Retry(action string, f func() error) error {
	err := retry.Do(func() error {
		if err := f(); err != nil {
			return err
		}
		return nil
	},
		retry.Attempts(30),
		retry.Delay(300),
		retry.LastErrorOnly(true),
	)

	if err != nil {
		return errors.Wrapf(err, "error when retrying %s", action)
	}

	return nil
}

func RetryWithReturn[T any](action string, f func() (T, error)) (T, error) {
	var t T
	var err error

	err = retry.Do(func() error {
		t, err = f()
		if err != nil {
			return err
		}

		return nil
	},
		retry.Attempts(30),
		retry.Delay(300),
		retry.LastErrorOnly(true),
	)

	if err != nil {
		return t, errors.Wrapf(err, "error when retrying %s", action)
	}

	return t, nil
}
