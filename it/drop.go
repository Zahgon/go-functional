package it

import "iter"

// Drop yields all values from a delegate iterator except the first `count`
// values.
func Drop[V any](delegate func(func(V) bool), count uint) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// Drop2 yields all pairs of values from a delegate iterator except the first
// `count` pairs.
func Drop2[V, W any](delegate func(func(V, W) bool), count uint) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// DropWhile yields all values from a delegate iterator after the predicate
// returns false.
func DropWhile[V any](delegate func(func(V) bool), predicate func(V) bool) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// DropWhile2 yields all pairs of values from a delegate iterator after the
// predicate returns false.
func DropWhile2[V, W any](delegate func(func(V, W) bool), predicate func(V, W) bool) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}
