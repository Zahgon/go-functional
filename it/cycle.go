package it

import "iter"

// Cycle yields values from an iterator repeatedly.
//
// Note: this is an infinite iterator.
//
// Note: memory usage will grow until all values from the underlying iterator
// are stored in memory.
func Cycle[V any](delegate func(func(V) bool)) iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// Cycle2 yields pairs of values from an iterator repeatedly.
//
// Note: this is an infinite iterator.
//
// Note: memory usage will grow until all values from the underlying iterator
// are stored in memory.
func Cycle2[V, W any](delegate func(func(V, W) bool)) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}
