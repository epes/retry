package retry

import (
	"fmt"
	"time"
)

// TesterFn is a function which tests the current status
// of a retry and returns whether to break out of the
// retry loop
type TesterFn func() (bool, error)

// Do retries the TesterFn until it returns true or
// it hits the max amount of retries
func Do(
	fn TesterFn,
	intervalMs int,
	maxRetries int,
	retryFn IntervalFn,
) error {
	for i := 0; i < maxRetries; i++ {
		result, err := fn()
		if err != nil {
			return err
		}

		if result {
			return nil
		}

		time.Sleep(time.Millisecond * time.Duration(retryFn(intervalMs, i)))
	}

	return fmt.Errorf("fn() false after %d retries", maxRetries)
}

// DoUntil retries the TesterFn until it returns true or
// the blocker TesterFn returns true first
func DoUntil(
	fn TesterFn,
	intervalMs int,
	blocker TesterFn,
	retryFn IntervalFn,
) error {
	iteration := 0

	for {
		block, err := blocker()
		if err != nil {
			return err
		}

		if block {
			return fmt.Errorf("DoUntil blocker is true")
		}

		result, err := fn()
		if err != nil {
			return err
		}

		if result {
			return nil
		}
		iteration++
		time.Sleep(time.Millisecond * time.Duration(retryFn(intervalMs, iteration)))
	}
}
