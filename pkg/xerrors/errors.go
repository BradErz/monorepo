package xerrors

import (
	"fmt"
)

type Error struct {
	err     error
	msg     string
	code    Code
	details map[string]string
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
		code:    code,
		err:     err,
		details: details,
		msg:     fmt.Sprintf(format, a...),
	}
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v %v", e.msg, e.err, e.details)
	}
	return e.msg
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() map[string]string {
	return e.details
}
