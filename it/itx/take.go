package itx

// Take is a convenience method for chaining [it.Take] on [Iterator]s.
func (iterator Iterator[V]) Take(limit uint) Iterator[V] { _ = "STUB: not implemented"; return nil }

// Take is a convenience method for chaining [it.Take2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Take(limit uint) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// TakeWhile is a convenience method for chaining [it.TakeWhile] on
// [Iterator]s.
func (iterator Iterator[V]) TakeWhile(predicate func(V) bool) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// TakeWhile is a convenience method for chaining [it.TakeWhile2] on
// [Iterator2]s.
func (iterator Iterator2[V, W]) TakeWhile(predicate func(V, W) bool) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}
