package itx

// Repeat yields the same value indefinitely.
func Repeat[V any](value V) Iterator[V] { _ = "STUB: not implemented"; return nil }

// Repeat2 yields the same two values indefinitely.
func Repeat2[V, W any](value1 V, value2 W) Iterator2[V, W] { _ = "STUB: not implemented"; return nil }
