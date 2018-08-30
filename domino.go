// domino.go: Methods of the Domino type
// This file is hereby released into the public domain.

package errors

func (d Domino) Cause() error {
	return d.c
}

func (d Domino) Error() string {
	// could not read config: [read failed: unable to open file or directory] caused: unable to unmarshal
	d.s + ": " + Quotes[0] + d.c.Error() + Quotes[1] + " caused: " + d.e.Error()
}

func (d Domino) StackTrace() StackTrace {
	return d.st
}
