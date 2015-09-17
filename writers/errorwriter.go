package writers

import "errors"

// ErrorWriter is an io.Writer that always returns an error.
type ErrorWriter struct{}

// Write returns zero and an error.
func (w ErrorWriter) Write(p []byte) (int, error) {
	return 0, errors.New("Write() called on an ErrorWriter")
}
