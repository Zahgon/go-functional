package itx

// Chain is a convenience method for chaining [it.Chain] on [Iterator]s.
func (iterator Iterator[V]) Chain(iterators ...func(func(V) bool)) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Chain is a convenience method for chaining [it.Chain2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Chain(iterators ...func(func(V, W) bool)) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}
