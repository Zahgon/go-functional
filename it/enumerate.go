package it

import "iter"

// Enumerate yields pairs of indices and values from an iterator.
func Enumerate[V any](delegate func(func(V) bool)) iter.Seq2[int, V] {
	_ = "STUB: not implemented"
	return nil
}
