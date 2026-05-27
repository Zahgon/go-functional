package itx

import (
	"iter"
)

type (
	// Iterator is a wrapper around [iter.Seq] that allows for method chaining of
	// most iterators found in the `it` package.
	Iterator[V any] func(func(V) bool)

	// Iterator2 is a wrapper around [iter.Seq2] that allows for method chaining
	// of most iterators found in the `it` package.
	Iterator2[V, W any] func(func(V, W) bool)
)

// From converts an iterator in an [Iterator] to support method chaining.
func From[V any](iterator func(func(V) bool)) Iterator[V] { _ = "STUB: not implemented"; return nil }

// From2 converts an iterator in an [Iterator2] to support method chaining.
func From2[V, W any](iterator func(func(V, W) bool)) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// FromSlice converts a slice to an [Iterator].
func FromSlice[V any](slice []V) Iterator[V] { _ = "STUB: not implemented"; return nil }

// FromMap converts a map to an [Iterator2].
func FromMap[V comparable, W any](m map[V]W) Iterator2[V, W] { _ = "STUB: not implemented"; return nil }

// Seq converts an [Iterator] to an [iter.Seq].
func (iterator Iterator[V]) Seq() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// Seq converts an [Iterator2] to an [iter.Seq2].
func (iterator Iterator2[V, W]) Seq() iter.Seq2[V, W] { _ = "STUB: not implemented"; return nil }

// Collect is a convenience method for chaining [slices.Collect] on
// [Iterator]s.
func (iterator Iterator[V]) Collect() []V { _ = "STUB: not implemented"; return nil }

// ForEach is a convenience method for chaining [it.ForEach] on [Iterator]s.
func (iterator Iterator[V]) ForEach(fn func(V)) { _ = "STUB: not implemented"; return }

// ForEach is a convenience method for chaining [it.ForEach2] on [Iterator2]s.
func (iterator Iterator2[V, W]) ForEach(fn func(V, W)) { _ = "STUB: not implemented"; return }

// Find is a convenience method for chaining [it.Find] on [Iterator]s.
func (iterator Iterator[V]) Find(predicate func(V) bool) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Find is a convenience method for chaining [it.Find2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Find(predicate func(V, W) bool) (V, W, bool) {
	_ = "STUB: not implemented"
	return *new(V), *new(W), false
}

// Collect2 consumes an [iter.Seq2] and returns two slices of values.
func (iterator Iterator2[V, W]) Collect() ([]V, []W) {
	_ = "STUB: not implemented"
	return nil,

		// Len is a convenience method for chaining [it.Len] on [Iterator]s.
		nil
}

func (iterator Iterator[V]) Len() int { _ = "STUB: not implemented"; return 0 }

// Len is a convenience method for chaining [it.Len2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Len() int { _ = "STUB: not implemented"; return 0 }

// Drain is a convenience method for chaining [it.Drain] on [Iterator]s.
func (iterator Iterator[V]) Drain() {
	_ = "STUB: not implemented"

	// Drain2 is a convenience method for chaining [it.Drain2] on [Iterator2]s.
	return
}

func (iterator Iterator2[V, W]) Drain() { _ = "STUB: not implemented"; return }
