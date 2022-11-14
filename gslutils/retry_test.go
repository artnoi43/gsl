package gslutils

import (
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestRetry(t *testing.T) {
	var lastErr = errors.New("lastErr")
	var i int
	attempts := 3

	f := func() error {
		i++

		if i < attempts {
			return errors.New("foo")
		}
		if i == attempts {
			return lastErr
		}

		return nil
	}

	err := retry(f, Attempts(attempts), LastErrorOnly(true))
	if !errors.Is(err, lastErr) {
		t.Fatal("expecing lastErr")
	}

	err = retry(f, Attempts(attempts), LastErrorOnly(false))
	if errors.Is(err, lastErr) {
		t.Fatal("expecting combined errors")
	}
}

func TestRetryDuration(t *testing.T) {
	sleepDuration := 2 * time.Second
	delay := 2 * time.Second
	attempts := 3

	f := func(err error) error {
		time.Sleep(sleepDuration)
		return err
	}

	// No error
	f0 := func() error {
		return f(nil)
	}

	// Always error
	f1 := func() error {
		return f(errors.New("foo"))
	}

	type retryConf struct {
		sleepDuration time.Duration
		delay         time.Duration
		attempts      int
		noError       bool
	}

	tests := make(map[*func() error]retryConf)
	tests[&f0] = retryConf{
		sleepDuration: sleepDuration,
		delay:         delay,
		attempts:      attempts,
		noError:       true,
	}
	tests[&f1] = retryConf{
		sleepDuration: sleepDuration,
		delay:         delay,
		attempts:      attempts,
	}

	for f, conf := range tests {
		start := time.Now()
		err := retry(*f, Delay(conf.delay), Attempts(conf.attempts))
		elapsed := time.Since(start)

		var expectedDuration time.Duration
		if conf.noError {
			if err != nil {
				t.Fatal("expecting nil error")
			}
			expectedDuration = conf.sleepDuration
		} else {
			if err == nil {
				t.Fatal("expecting non-nil error")
			}
			expectedDuration = time.Duration(conf.attempts * (int(conf.sleepDuration) + int(conf.delay)))
		}

		rounded := elapsed.Round(time.Second)
		t.Log(conf, rounded)
		if rounded > expectedDuration {
			t.Fatalf("took to long for f - expecting %v, got %v\n", expectedDuration, rounded)
		}
	}
}

func TestRetryWithReturnDuration(t *testing.T) {
	sleepDuration := 2 * time.Second
	delay := 2 * time.Second
	attempts := 3

	f := func(i int) (int, error) {
		time.Sleep(sleepDuration)

		if i < 0 {
			return 0, errors.New("foo")
		}

		return i, nil
	}

	// No error
	f0 := func() (int, error) {
		return f(5)
	}
	// Always error
	f1 := func() (int, error) {
		return f(-1)
	}

	type retryConf struct {
		sleepDuration time.Duration
		delay         time.Duration
		attempts      int
		noError       bool
	}

	tests := make(map[*func() (int, error)]retryConf)
	tests[&f0] = retryConf{
		sleepDuration: sleepDuration,
		delay:         delay,
		attempts:      attempts,
		noError:       true,
	}
	tests[&f1] = retryConf{
		sleepDuration: sleepDuration,
		delay:         delay,
		attempts:      attempts,
	}

	for f, conf := range tests {
		start := time.Now()
		_, err := RetryWithReturn("test", *f, Delay(conf.delay), Attempts(conf.attempts))
		elapsed := time.Since(start)

		var expectedDuration time.Duration
		if conf.noError {
			if err != nil {
				t.Fatal("expecting nil error")
			}
			expectedDuration = conf.sleepDuration
		} else {
			if err == nil {
				t.Fatal("expecting non-nil error")
			}
			expectedDuration = time.Duration(conf.attempts * (int(conf.sleepDuration) + int(conf.delay)))
		}

		rounded := elapsed.Round(time.Second)
		t.Log(conf, rounded)
		if rounded > expectedDuration {
			t.Fatalf("took to long for f - expecting %v, got %v\n", expectedDuration, rounded)
		}
	}
}
