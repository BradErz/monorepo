package xerrors

import (
	"errors"
	"fmt"
)

type Error struct {
	Err     error
	Message string
	Code    Code
	Details map[string]string
}

// Wrapf returns a wrapped error.
func Wrapf(code Code, err error, format string, a ...interface{}) error {
	return wrapf(code, err, nil, format, a...)
}

// Newf instantiates a new error.
func Newf(code Code, format string, a ...interface{}) error {
	return Wrapf(code, nil, format, a...)
}

func Detailsf(code Code, err error, details map[string]string, format string, a ...interface{}) error {
	return wrapf(code, err, details, format, a...)
}

func wrapf(code Code, err error, details map[string]string, format string, a ...interface{}) error {
	return &Error{
		Code:    code,
		Err:     err,
		Details: details,
		Message: fmt.Sprintf(format, a...),
	}
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v %v", e.Message, e.Err, e.Details)
	}
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

func NotValidObjectID(id string) error {
	return Newf(CodeInvalidArgument, "%s is not a valid object id", id)
}

func IsNotFound(err error) bool {
	myError := &Error{}
	if !errors.As(err, &myError) {
		return false
	}
	return myError.Code == CodeNotFound
}
