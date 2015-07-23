package writers

import (
	"bytes"
	"io"
)

// A CRLF wraps an existing io.Writer, replacing all written \n with \r\n.
type CRLF struct {
	w io.Writer
}

// NewCRLF returns a pointer to a new CRLF wrapping w.
func NewCRLF(w io.Writer) *CRLF {
	return &CRLF{w}
}

// Write writes p to the wrapped io.Writer with each instance of \n in p
// replaced with \r\n. The byte count returned is no larger than len(p).
func (crlf *CRLF) Write(p []byte) (int, error) {
	max := len(p)
	p = bytes.Replace(p, []byte("\n"), []byte("\r\n"), -1)
	n, err := crlf.w.Write(p)
	if n > max {
		n = max
	}
	return n, err
}
