package itx

// Filter is a convenience method for chaining [it.Filter] on [Iterator]s.
func (iterator Iterator[V]) Filter(predicate func(V) bool) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Exclude is a convenience method for chaining [it.Exclude] on [Iterator]s.
func (iterator Iterator[V]) Exclude(predicate func(V) bool) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Filter is a convenience method for chaining [it.Filter2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Filter(predicate func(V, W) bool) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// Exclude is a convenience method for chaining [it.Exclude2] on [Iterator2]s.
func (iterator Iterator2[V, W]) Exclude(predicate func(V, W) bool) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// FilterError is a convenience method for chaining [it.FilterError] on
// [Iterator]s.
func (iterator Iterator[V]) FilterError(predicate func(V) (bool, error)) Iterator2[V, error] {
	_ = "STUB: not implemented"
	return nil
}

// ExcludeError is a convenience method for chaining [it.ExcludeError] on
// [Iterator]s.
func (iterator Iterator[V]) ExcludeError(predicate func(V) (bool, error)) Iterator2[V, error] {
	_ = "STUB: not implemented"
	return nil
}
