package itx

import (
	"io"
)

// Lines yields lines from an io.Reader.
//
// Note: lines longer than 65536 will cauese an error.
func Lines(r io.Reader) Iterator2[[]byte, error] { _ = "STUB: not implemented"; return nil }

// LinesString yields lines from an io.Reader as strings.
//
// Note: lines longer than 65536 will cauese an error.
func LinesString(r io.Reader) Iterator2[string, error] { _ = "STUB: not implemented"; return nil }
