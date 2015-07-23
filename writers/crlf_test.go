package writers

import (
	"bytes"
	"testing"
)

func TestCRLF(t *testing.T) {
	var buf bytes.Buffer
	crlf := NewCRLF(&buf)

	inputs := [][]byte{[]byte{}, []byte("ab"), []byte("\n"), []byte("a\nb\n")}
	outputs := []string{"", "ab", "\r\n", "a\r\nb\r\n"}
	for i, input := range inputs {
		buf.Reset()
		n, err := crlf.Write(input)
		if err != nil {
			t.Errorf("CRLF.Write(%#v) returned err == %v; want %v",
				input, err, nil)
		} else if n != len(input) {
			t.Errorf("CRLF.Write(%#v) returned n == %d; want %d",
				input, n, len(input))
		} else if buf.String() != outputs[i] {
			t.Errorf("CRLF.Write(%#v) wrote %#v; want %#v",
				input, buf.String(), outputs[i])
		}
	}
}
