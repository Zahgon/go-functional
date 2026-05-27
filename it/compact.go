package it

import "iter"

// Compact yields all values from a delegate iterator that are not zero values.
func Compact[V comparable](delegate func(func(V) bool)) iter.Seq[V] {
	_ = "STUB: not implemented"
	return nil
}
