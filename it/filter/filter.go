// This package contains functions intended for use with [iter.Filter].
package filter

import (
	"cmp"
	"regexp"
)

// IsEven returns true when the provided integer is even.
func IsEven[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](integer T) bool {
	_ = "STUB: not implemented"
	return false

	// IsOdd returns true when the provided integer is odd.
}

func IsOdd[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64](integer T) bool {
	_ = "STUB: not implemented"
	return false

	// IsEqual returns a function that returns true when the provided value is equal
	// to some value.
}

func IsEqual[T comparable](value T) func(T) bool { _ = "STUB: not implemented"; return nil }

// NotEqual returns a function that returns true when the provided value is not
// equal to some value.
func NotEqual[T comparable](value T) func(T) bool { _ = "STUB: not implemented"; return nil }

// IsZero returns true when the provided value is the zero value for its type.
func IsZero[T comparable](value T) bool { _ = "STUB: not implemented"; return false }

// GreaterThan returns a function that returns true when the provided value is
// greater than a threshold.
func GreaterThan[T cmp.Ordered](threshold T) func(T) bool { _ = "STUB: not implemented"; return nil }

// LessThan returns a function that returns true when the provided value is less
// than a threshold.
func LessThan[T cmp.Ordered](threshold T) func(T) bool { _ = "STUB: not implemented"; return nil }

// Passthrough returns a function that returns true for any value.
func Passthrough[V any](value V) bool {
	_ = "STUB: not implemented"

	// Passthrough2 returns a function that returns true for any pair of values.
	return false
}

func Passthrough2[V, W any](v V, w W) bool {
	_ = "STUB: not implemented"

	// Not returns a function that inverts the result of the provided function.
	return false
}

func Not[T any](fn func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

// And returns a function that returns true when all provided functions return
// true.
func And[T any](filters ...func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

// Or returns a function that returns true when any of the provided functions
// return true.
func Or[T any](filters ...func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

// Match returns a function that returns true when the provided string or byte
// slice matches a pattern. See [regexp.MatchString].
func Match[T string | []byte](pattern *regexp.Regexp) func(T) bool {
	_ = "STUB: not implemented"
	return nil
}

// Contains returns a function that returns true when the provided string or
// byte slice is found within another string or byte slice. See
// [strings.Contains] and [bytes.Contains].
func Contains[T string | []byte](t T) func(T) bool { _ = "STUB: not implemented"; return nil }
