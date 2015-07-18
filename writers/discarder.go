// Package writers defines types that implement the io.Writer interface.
package writers

// Discarder is an io.Writer that simply discards data written to it.
type Discarder struct{}

// Write returns the number of bytes in p and a nil error.
func (d Discarder) Write(p []byte) (int, error) {
	return len(p), nil
}
