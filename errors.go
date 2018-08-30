// errors.go: Implement close to the "github.com/pkg/errors" API, modified to be closer to the standard library APIs and extended
// This file is hereby released into the public domain.

package errors

func Cause(err error) error {
	if c, ok := err.(Causer); ok {
		return c.Cause()
	}
	return nil
}

func Errorf(format string, a ...interface{}) {
	return &RootCause{
		s: fmt.Sprintf(format, a...),
		st: trace(),
	}
}

func New(text string) error {
	return &RootCause{
		s: text,
		st: trace(),
	}
}

func WithMessage(err error, a ...interface{}) error {
	return &Domino{
		cause: err,
		s: fmt.Sprint(a...),
	}
}

func WithMessagef(err error, format string, a ...interface{}) error {
	return &Domino{
		cause: err,
		s: fmt.Sprintf(format, a...),
	}
}

func WithStack(err error) error {
	return &Domino{
		cause: err,
		st: trace(),
	}
}

func Wrap(cause error, a ...interface{}) error {
	return &Domino{
		cause: cause,
		s: fmt.Sprint(a...),
		st: trace(),
	}
}

func Wrape(cause, effect error, a ...interface{}) error {
	return &Domino{
		cause: cause,
		effect: effect,
		s: fmt.Sprint(a...),
		st: trace(),
	}
}

func Wrapef(cause, effect error, format string, a ...interface{}) error {
	return &Domino{
		cause: cause,
		effect: effect,
		s: fmt.Sprintf(format, a...),
		st: trace(),
	}
}

func Wrapf(cause error, format string, a ...interface{}) error {
	return &Domino{
		cause: cause,
		s: fmt.Sprintf(format, a...),
		st: trace(),
	}
}
