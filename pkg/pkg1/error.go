package pkg1

import (
	"fmt"
)

type pkg1SpecificError struct {
	cause error
	msg   string
}

var ErrPkg1 *pkg1SpecificError

func (e *pkg1SpecificError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

func (e *pkg1SpecificError) Unwrap() error {
	return e.cause
}

func (e *pkg1SpecificError) Is(target error) bool {
	return target == ErrPkg1
}

func newPkg1Error(cause error) error {
	return &pkg1SpecificError{
		cause: cause,
		msg:   "something goes wrong in pkg1",
	}
}
