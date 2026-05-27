package it

import "iter"

// Repeat yields the same value indefinitely.
func Repeat[V any](value V) iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// Repeat2 yields the same two values indefinitely.
func Repeat2[V, W any](value1 V, value2 W) iter.Seq2[V, W] { _ = "STUB: not implemented"; return nil }
