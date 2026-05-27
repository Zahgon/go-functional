package it

import "iter"

// Exhausted is an iterator that yields no values.
func Exhausted[V any]() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// Exhausted2 is an iterator that yields no values.
func Exhausted2[V, W any]() iter.Seq2[V, W] { _ = "STUB: not implemented"; return nil }
