package itx

// Transform is a convenience method for chaining [it.Map] on [Iterator]s where
// the provided functions argument type is the same as its return type.
//
// This is a limited version of [it.Map] due to a limitation on Go's type
// system whereby new generic type parameters cannot be defined on methods.
func (iterator Iterator[V]) Transform(f func(V) V) Iterator[V] {
	_ = "STUB: not implemented"
	return nil
}

// Transform is a convenience method for chaining [it.Map2] on [Iterator2]s
// where the provided functions argument type is the same as its return type.
//
// This is a limited version of [it.Map2] due to a limitation on Go's type
// system whereby new generic type parameters cannot be defined on methods.
func (iterator Iterator2[V, W]) Transform(f func(V, W) (V, W)) Iterator2[V, W] {
	_ = "STUB: not implemented"
	return nil
}

// TransformError is a convenience method for chaining [it.MapError] on
// [Iterator]s where the provided functions argument type is the same as its
// return type.
//
// This is a limited version of [it.MapError] due to a limitation on Go's type
// system whereby new generic type parameters cannot be defined on methods.
func (iterator Iterator[V]) TransformError(f func(V) (V, error)) Iterator2[V, error] {
	_ = "STUB: not implemented"
	return nil
}
