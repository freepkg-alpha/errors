// cause.go: Methods of the Cause type
// This file is hereby released into the public domain.

package errors

func (c Cause) Error() string {
	return c.s
}

func (c Cause) StackTrace() StackTrace {
	return c.st
}
