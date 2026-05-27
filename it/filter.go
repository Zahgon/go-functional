package it

import "iter"

// Filter yields values from an iterator that satisfy a predicate.
func Filter[V any](delegate func(func(V) bool), predicate func(V) bool) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// Exclude yields values from an iterator that do not satisfy a predicate.
func Exclude[V any](delegate func(func(V) bool), predicate func(V) bool) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}

// Filter2 yields values from an iterator that satisfy a predicate.
func Filter2[V, W any](delegate func(func(V, W) bool), predicate func(V, W) bool) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// Exclude2 yields values from an iterator that do not satisfy a predicate.
func Exclude2[V, W any](delegate func(func(V, W) bool), predicate func(V, W) bool) iter.Seq2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// FilterError yields values from an iterator that satisfy a predicate where
// the predicate can return an error.
func FilterError[V any](delegate func(func(V) bool), predicate func(V) (bool, error)) iter.Seq2[V, error] {
	_ = "STUB: not implemented"
	return nil
}

// ExcludeError yields values from an iterator that do not satisfy a predicate
// where the predicate can return an error.
func ExcludeError[V any](delegate func(func(V) bool), predicate func(V) (bool, error)) iter.Seq2[V, error] {
	_ = "STUB: not implemented"
	return nil
}

func not[V any](predicate func(V) bool) func(V) bool { _ = "STUB: not implemented"; return nil }

func not2[V, W any](predicate func(V, W) bool) func(V, W) bool {
	_ = "STUB: not implemented"
	return nil
}

func notError[V any](predicate func(V) (bool, error)) func(V) (bool, error) {
	_ = "STUB: not implemented"
	return nil
}
