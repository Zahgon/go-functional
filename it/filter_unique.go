package it

import "iter"

// FilterUnique yields all the unique values from an iterator.
//
// Note: All unique values seen from an iterator are stored in memory.
func FilterUnique[V comparable](iterator func(func(V) bool)) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}
