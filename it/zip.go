package it

import "iter"

// Zip yields pairs of values from two iterators.
func Zip[V, W any](left func(func(V) bool), right func(func(W) bool)) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// Left is a convenience function that unzips an iterator and returns the left
// iterator, closing the right iterator.
func Left[V, W any](delegate func(func(V, W) bool)) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// Right is a convenience function that unzips an iterator and returns the
// right iterator, closing the left iterator.
func Right[V, W any](delegate func(func(V, W) bool)) iter.Seq[W] {
	_ = "STUB: not implemented"
	return nil
}
