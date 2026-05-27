package itx

// Left is a convenience method that unzips an [Iterator2] and returns the left
// iterator, closing the right iterator.
func (iterator Iterator2[V, W]) Left() Iterator[V] { _ = "STUB: not implemented"; return nil }

// Right is a convenience method that unzips an [Iterator2] and returns the
// right iterator, closing the left iterator.
func (iterator Iterator2[V, W]) Right() Iterator[W] { _ = "STUB: not implemented"; return nil }
