package it

import "iter"

// Once yields the provided value once.
func Once[V any](value V) iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// Once2 yields the provided value pair once.
func Once2[V, W any](v V, w W) iter.Seq2[V, W] { _ = "STUB: not implemented"; return nil }
