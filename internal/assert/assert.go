package assert

import (
	"testing"
)

func True(t *testing.T, value bool) { _ = "STUB: not implemented"; return }

func False(t *testing.T, value bool) { _ = "STUB: not implemented"; return }

func Equal[T comparable](t *testing.T, a, b T) { _ = "STUB: not implemented"; return }

func SliceEqual[T comparable](t *testing.T, a, b []T) { _ = "STUB: not implemented"; return }

func Empty[E any, Slice ~[]E | ~string](t *testing.T, items Slice) {
	_ = "STUB: not implemented"
	return
}
