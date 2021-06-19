package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorCode string

const (
	ErrRouteNotFound = ErrorCode("ERR-00000404")
	ErrorBadRequest  = ErrorCode("ERR-00000400")

	ErrUnknown       = ErrorCode("ERR-10000000")
	ErrNotSupport    = ErrorCode("ERR-10000001")
	ErrNeedImplement = ErrorCode("ERR-10000002")
)

func (e ErrorCode) New(msg string) error {
	return customError{
		code:          e,
		originalError: errors.New(msg),
	}
}

func (e ErrorCode) Newf(msg string, args ...interface{}) error {
	return customError{
		code:          e,
		originalError: fmt.Errorf(msg, args...),
	}
}

func (e ErrorCode) Wrap(err error, msg string) error {
	return e.Wrapf(err, msg)
}

func (e ErrorCode) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)
	return customError{
		code:          e,
		originalError: newErr,
		contextInfo:   nil,
	}
}
