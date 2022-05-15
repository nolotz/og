package boo

import "github.com/nolotz/og/internal/runtimeutil"

// WithStack returns a StackError wrapping any given error
// together with the corresponding program counter.
func WithStack(err error) error {
	return &StackError{
		PC:    runtimeutil.FailingPC(),
		Cause: err,
	}
}
