// errors.go: Implement the "errors" API
// This file is hereby released into the Public Domain.

package errors

import "errors"

func New(text string) error {
	return errors.New(text)
}
