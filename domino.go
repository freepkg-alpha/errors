// domino.go: Methods of the Domino type
// This file is hereby released into the public domain.

package errors

func (d Domino) Cause() error {
	return d.c
}

func (d Domino) Error() string {
	// could not read config: [read failed: unable to open file or directory] caused: [unable to unmarshal]
	s := d.s
	if d.cause != nil {
		if len(s) > 0 {
			s += ": "
		}
		s += Quotes[0] + d.cause.Error() + Quotes[1]
	}
	if d.effect != nil {
		s += " caused: " + Quotes[0] + d.effect.Error() + Quotes[1]
	}
	return s
}

func (d Domino) StackTrace() StackTrace {
	return d.st
}
