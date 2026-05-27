package op

// Add returns the sum of `a` and `b`.
func Add[V ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~string | ~float32 | ~float64](a, b V) V {
	_ = "STUB: not implemented"

	// Ref returns a reference to a copy of the provided value.
	//
	// This may be useful when interacting with packages that use pointers as
	// proxies for optional values.
	return *new(V)
}

func Ref[V any](v V) *V {
	_ = "STUB: not implemented"

	// Deref returns the value pointed to by the provided pointer.
	return nil
}

func Deref[V any](v *V) V { _ = "STUB: not implemented"; return *new(V) }
