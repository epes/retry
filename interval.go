package retry

import (
	"math"
)

// IntervalFn is a function that takes an interval and the
// iteration count and returns an int of milliseconds to
// wait before retrying
type IntervalFn func(interval int, iteration int) int

// Regular is a IntervalFn with no fallback
func Regular(interval int, iteration int) int {
	return interval
}

// Exponential is a IntervalFn with exponential fallback
func Exponential(interval int, iteration int) int {
	return interval * int(math.Pow(2.0, float64(iteration)))
}
