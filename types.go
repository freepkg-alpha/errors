// types.go: Public types used by this package
// This file is hereby released into the public domain.

package errors

type Cause struct {
	s  string
	st StackTrace
}

type Causer interface {
	Cause() error
}

type Domino struct {
	cause  error
	effect error
	s      string
	st     StackTrace
}

type Frame uintptr

type StackTrace []Frame

type StackTracer interface {
	StackTrace() StackTrace
}
