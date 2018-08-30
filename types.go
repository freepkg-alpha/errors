// types.go: Public types used by this package
// This file is hereby released into the public domain.

package errors

type Causer interface {
	Cause() error
}

type Domino struct {
	cause  error
	effect error
	s      string
	st     StackTrace
}

type Frame *runtime.Frame

type RootCause struct {
	s  string
	st StackTrace
}

type StackTrace []Frame

type StackTracer interface {
	StackTrace() StackTrace
}
