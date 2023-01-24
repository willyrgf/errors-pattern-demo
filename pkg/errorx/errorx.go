package errorx

import "fmt"

type StandardError struct {
	cause    error
	msg      string
	exitCode uint
}

type SpecificError struct {
	cause    error
	msg      string
	exitCode uint
}

var ErrStandard *StandardError
var ErrSpecific *SpecificError

func (e *StandardError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

func (e *StandardError) Unwrap() error {
	return e.cause
}

func (e *StandardError) Is(target error) bool {
	return target == ErrStandard
}

func (e *StandardError) ExitCode() uint {
	return e.exitCode
}

func (e *SpecificError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

func (e *SpecificError) Unwrap() error {
	return e.cause
}

func (e *SpecificError) Is(target error) bool {
	return target == ErrSpecific
}

func (e *SpecificError) ExitCode() uint {
	return e.exitCode
}

func NewStandardError(cause error) error {
	return &StandardError{
		cause:    cause,
		msg:      "a standard error with exit code",
		exitCode: 1,
	}
}

func NewSpecificError(cause error) error {
	return &SpecificError{
		cause:    cause,
		msg:      "a specific error with exit code",
		exitCode: 127,
	}
}
