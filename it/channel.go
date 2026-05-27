package it

import "iter"

// FromChannel yields values from a channel.
//
// In order to avoid a deadlock, the channel must be closed before attempting
// to called `stop` on a pull-style iterator.
func FromChannel[V any](channel <-chan V) iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// ToChannel sends yielded values to a channel.
//
// The channel is closed when the iterator is exhausted. Beware of leaked go
// routines when using this function with an infinite iterator.
func ToChannel[V any](seq func(func(V) bool)) <-chan V { _ = "STUB: not implemented"; return nil }
