// root.go: Methods of the RootCause type
// This file is hereby released into the public domain.

package errors

func (rc RootCause) Error() string {
	return rc.s
}

func (rc RootCause) StackTrace() StackTrace {
	return rc.st
}
