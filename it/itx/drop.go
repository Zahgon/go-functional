package itx

// Drop is a convenience method for chaining [it.Drop] on [Iterator]s.
func (iterator Iterator[V]) Drop(count uint) Iterator[V] { _ = "STUB: not implemented"; return nil }

// Drop is a convenience method for chaining [it.Drop2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Drop(count uint) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// DropWhile is a convenience method for chaining [it.DropWhile] on
// [Iterator]s.
func (iterator Iterator[V]) DropWhile(predicate func(V) bool) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// DropWhile is a convenience method for chaining [it.DropWhile2] on
// [Iterator2]s.
func (iterator Iterator2[V, W]) DropWhile(predicate func(V, W) bool) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}
