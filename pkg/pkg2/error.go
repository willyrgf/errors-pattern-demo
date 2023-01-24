package pkg2

import (
	"fmt"
)

type pkg2SpecificError struct {
	cause error
	msg   string
}

var ErrPkg2 *pkg2SpecificError

func (e *pkg2SpecificError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

func (e *pkg2SpecificError) Unwrap() error {
	return e.cause
}

func (e *pkg2SpecificError) Is(target error) bool {
	return target == ErrPkg2
}

func newPkg2Error(cause error) error {
	return &pkg2SpecificError{
		cause: cause,
		msg:   "something goes wrong in pkg2",
	}
}
