// errors.go: Implement part of the "fmt" API
// This file is hereby released into the Public Domain.

package errors

import "fmt"

func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
