package it

import "iter"

// Map yields values from an iterator that have had the provided function
// applied to each value.
func Map[V, W any](delegate func(func(V) bool), f func(V) W) iter.Seq[W] {
	_ = "STUB: not implemented"
	return nil
}

// Map2 yields pairs of values from an iterator that have had the provided
// function applied to each pair.
func Map2[V, W, X, Y any](delegate func(func(V, W) bool), f func(V, W) (X, Y)) iter.Seq2[X, Y] {
	_ = "STUB: not implemented"
	return nil
}

// MapError yields values from an iterator that have had the provided function
// applied to each value where the function can return an error.
func MapError[V, W any](delegate func(func(V) bool), f func(V) (W, error)) iter.Seq2[W, error] {
	_ = "STUB: not implemented"
	return nil
}
