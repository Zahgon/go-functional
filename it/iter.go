package it

import (
	"cmp"
)

// ForEach consumes an iterator and applies a function to each value yielded.
func ForEach[V any](iterator func(func(V) bool), fn func(V)) { _ = "STUB: not implemented"; return }

// ForEach2 consumes an iterator and applies a function to each pair of values.
func ForEach2[V, W any](iterator func(func(V, W) bool), fn func(V, W)) {
	_ = "STUB: not implemented"
	return
}

// Fold will fold every element into an accumulator by applying a function and
// passing an initial value.
func Fold[V, R any](iterator func(func(V) bool), fn func(R, V) R, initial R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// Fold2 will fold every element into an accumulator by applying a function and
// passing an initial value.
func Fold2[V, W, R any](iterator func(func(V, W) bool), fn func(R, V, W) R, initial R) R {
	_ = "STUB: not implemented"
	return *new(R)
}

// Max consumes an iterator and returns the maximum value yielded and true if
// there was at least one value, or the zero value and false if the iterator
// was empty.
func Max[V cmp.Ordered](iterator func(func(V) bool)) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Min consumes an iterator and returns the minimum value yielded and true if
// there was at least one value, or the zero value and false if the iterator
// was empty.
func Min[V cmp.Ordered](iterator func(func(V) bool)) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Find consumes an iterator until a value is found that satisfies a predicate.
// It returns the value and true if one was found, or the zero value and false
// if the iterator was exhausted.
func Find[V any](iterator func(func(V) bool), pred func(V) bool) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Find2 consumes an iterator until a pair of values is found that satisfies a
// predicate. It returns the pair and true if one was found, or the zero values
// and false if the iterator was exhausted.
func Find2[V, W any](iterator func(func(V, W) bool), pred func(V, W) bool) (V, W, bool) {
	_ = "STUB: not implemented"
	return *new(V), *new(W), false
}

// TryCollect consumes an [iter.Seq2] where the right side yields errors and
// returns a slice of values and the first error encountered. Iteration stops
// at the first error.
func TryCollect[V any](iterator func(func(V, error) bool)) ([]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustCollect consumes an [iter.Seq2] where the right side yields errors and
// returns a slice of values. If an error is encountered this function will
// panic.
func MustCollect[V any](iterator func(func(V, error) bool)) []V {
	_ = "STUB: not implemented"
	return nil
}

// Collect2 consumes an [iter.Seq2] and returns two slices of values.
func Collect2[V, W any](iterator func(func(V, W) bool)) ([]V, []W) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Len consumes an [iter.Seq] and returns the number of values yielded.
func Len[V any](iterator func(func(V) bool)) int { _ = "STUB: not implemented"; return 0 }

// Len2 consumes an [iter.Seq2] and returns the number of pairs of values
// yielded.
func Len2[V, W any](iterator func(func(V, W) bool)) int { _ = "STUB: not implemented"; return 0 }

// Contains consumes an [iter.Seq] until the provided value is found and
// returns true. If the value is not found, it returns false when the iterator
// is exhausted.
func Contains[V comparable](iterator func(func(V) bool), v V) bool {
	_ = "STUB: not implemented"
	return false
}

// Drain consumes an [iter.Seq] completely, dropping all values.
//
// You may wish to use this to execute side effects without needing to collect
// values.
func Drain[V any](iterator func(func(V) bool)) { _ = "STUB: not implemented"; return }

// Drain2 consumes an [iter.Seq2] completely, dropping all values.
//
// You may wish to use this to execute side effects without needing to collect
// values.
func Drain2[V, W any](iterator func(func(V, W) bool)) { _ = "STUB: not implemented"; return }

// All consumes an [iter.Seq] of `bool`s and returns true if all values are
// true, or false otherwise. Iteration will terminate early if a `false` is
// encountered.
func All(iterator func(func(bool) bool)) bool { _ = "STUB: not implemented"; return false }

// All consumes an [iter.Seq] of `bool`s and returns true if any of the values
// are true, or false otherwise. Iteration will terminate early if a `true` is
// encountered.
func Any(iterator func(func(bool) bool)) bool { _ = "STUB: not implemented"; return false }
