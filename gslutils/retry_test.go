package gslutils

import (
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestRetry(t *testing.T) {
	fooErr := errors.New("foo")
	lastErr := errors.New("lastErr")

	var i int
	attempts := 3

	f := func() error {
		i++

		if i < attempts {
			return fooErr
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

	// Expecting non-nil err, but the retry should break at the fist run
	var j int
	delay := time.Second * 3
	start := time.Now()
	err = retry(
		func() error {
			j++
			if j < 2 {
				return fooErr
			}

			return nil
		},
		Delay(delay), Attempts(attempts), StopOnError(fooErr), LastErrorOnly(true),
	)
	if err == nil {
		t.Fatal("expecting non-nil error")
	}
	if !errors.Is(fooErr, err) {
		t.Fatal("expecing fooErr")
	}
	if elapsed := time.Since(start); elapsed.Round(time.Second) > delay {
		t.Fatal("this retry should break right away")
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
		err := Retry("testRetryDuration", *f, Delay(conf.delay), Attempts(conf.attempts))
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
		t.Log(conf, rounded, err)
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
		_, err := RetryWithReturn("testRetryWithReturnDuration", *f, Delay(conf.delay), Attempts(conf.attempts))
		elapsed := time.Since(start)

		var expectedDuration time.Duration
		if conf.noError {
			if err != nil {
				t.Fatal("expecting nil error")
			}

			t.Log("err", err)
			expectedDuration = conf.sleepDuration
		} else {
			if err == nil {
				t.Fatal("expecting non-nil error")
			}

			t.Log("err", err)
			expectedDuration = time.Duration(conf.attempts * (int(conf.sleepDuration) + int(conf.delay)))
		}

		rounded := elapsed.Round(time.Second)
		t.Log(conf, rounded)
		if rounded > expectedDuration {
			t.Fatalf("took to long for f - expecting %v, got %v\n", expectedDuration, rounded)
		}
	}
}
