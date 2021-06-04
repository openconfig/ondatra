// Package usererr wraps server errors to distinguish them from infrastructure failures.
package usererr

import (
	"github.com/pkg/errors"
)

type usererr struct {
	error
}

// New constructs a new user error with a format specifier.
func New(format string, args ...interface{}) error {
	return &usererr{errors.Errorf(format, args...)}
}

// Wrap wraps the specified err as a user error.
func Wrap(err error) error {
	return &usererr{err}
}

// Wrapf wraps the specified err as a user error with a format specifier.
func Wrapf(err error, format string, args ...interface{}) error {
	return &usererr{errors.Wrapf(err, format, args...)}
}

// In returns whether any error in the specified errors chain is a user error.
func In(err error) bool {
	for err != nil {
		if _, ok := err.(*usererr); ok {
			return true
		}
		err = unwrap(err)
	}
	return false
}

// unwrap can unwrap errors produced by errors.Wrap or by the %w verb.
func unwrap(err error) error {
	if c, ok := err.(interface {
		Cause() error
	}); ok {
		return c.Cause()
	}
	if u, ok := err.(interface {
		Unwrap() error
	}); ok {
		return u.Unwrap()
	}
	return nil
}
