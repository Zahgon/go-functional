package it

import "iter"

// Chain yields values from multiple iterators in sequence.
func Chain[V any](iterators ...func(func(V) bool)) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// Chain2 yields values from multiple iterators in sequence.
func Chain2[V, W any](iterators ...func(func(V, W) bool)) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}
