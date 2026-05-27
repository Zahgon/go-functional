package it

import "iter"

// Take yields the first `limit` values from a delegate iterator.
func Take[V any](delegate func(func(V) bool), limit uint) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// Take2 yields the first `limit` pairs of values from a delegate iterator.
func Take2[V, W any](delegate func(func(V, W) bool), limit uint) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// TakeWhile yields values from a delegate iterator until the predicate returns
// false.
func TakeWhile[V any](delegate func(func(V) bool), predicate func(V) bool) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// TakeWhile2 yields pairs of values from a delegate iterator until the
// predicate returns false.
func TakeWhile2[V, W any](delegate func(func(V, W) bool), predicate func(V, W) bool) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}
