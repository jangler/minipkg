package writers

import (
	"io"
	"testing"
)

func TestDiscarderImplementsWriter(t *testing.T) {
	var d Discarder
	var _ io.Writer = d
}

func TestDiscarderWrite(t *testing.T) {
	var d Discarder
	n, err := d.Write([]byte{})
	if n != 0 || err != nil {
		t.Errorf("d.Write([]byte{}) == (%#v, %#v); want (%#v, %#v)",
			n, err, 0, nil)
	}
	n, err = d.Write([]byte{0})
	if n != 1 || err != nil {
		t.Errorf("d.Write([]byte{0}) == (%#v, %#v); want (%#v, %#v)",
			n, err, 1, nil)
	}
}
