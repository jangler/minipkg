package writers

import (
	"io"
	"testing"
)

func TestErrorWriterImplementsWriter(t *testing.T) {
	var w ErrorWriter
	var _ io.Writer = w
}

func TestErrorWriterWrite(t *testing.T) {
	var w ErrorWriter
	n, err := w.Write([]byte{})
	if n != 0 || err == nil {
		t.Errorf("w.Write([]byte{}) == (%#v, %#v); want (0, error)", n, err)
	}
}
