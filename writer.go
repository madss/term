package term

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

// A Writer is a implements the io.Writer interface.
//
// The embedded Context can be used for formatting the output.
type Writer struct {
	Context
	w io.Writer
}

// New returns a new Writer.
//
// This is similar to NewWriter(os.Stdout), except that formatting will only be
// applied if os.Stdout is a real tty.
func New() *Writer {
	w := NewWriter(os.Stdout)
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		w.bypass = true
	}
	return w
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// With creates a new child Writer.
//
// The child Writer will inherit the parameters of the parent, only overriding
// the parameters provided. This is useful for temporarily changing the
// formatting.
func (w Writer) With(opts ...Option) *Writer {
	return &Writer{
		Context: *w.Context.With(opts...),
		w:       w.w,
	}
}

// Write implements the io.Writer interface.
func (w Writer) Write(b []byte) (n int, err error) {
	return w.Context.Write(w.w, b)
}
