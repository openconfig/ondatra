// Package closer provides a utility for closing resources.
package closer

import (
	log "github.com/golang/glog"
	"github.com/pkg/errors"
)

// dummy internal error to indicate the close should only log.
var errLogOnly = errors.New("logOnly")

// Close runs the given close function, assigning the close error, if any,
// to rerr if rerr is non-nil. If rerr is nil, the error is logged.
// Must be called in a defer on a return value of exactly "&rerr".
func Close(rerr *error, fn func() error, msg string) {
	if cerr := fn(); cerr != nil {
		cerr = errors.Wrap(cerr, msg)
		if *rerr == nil {
			*rerr = cerr
		} else {
			log.Error(cerr)
		}
	}
}

// CloseOnErr runs the given close function if rerr is not-nil.
// The close error, if any, is logged.
// Must be called in a defer on a return value of exactly "&rerr".
func CloseOnErr(rerr *error, fn func() error, msg string) {
	if *rerr != nil {
		CloseAndLog(fn, msg)
	}
}

// CloseAndLog runs the given close function and logs the close error, if any.
// Must be called in a defer.
func CloseAndLog(fn func() error, msg string) {
	Close(&errLogOnly, fn, msg)
}

// CloseVoidOnErr runs the given close function if rerr is not-nil.
// Must be called in a defer on a return value of exactly "&rerr".
func CloseVoidOnErr(rerr *error, fn func()) {
	if *rerr != nil {
		fn()
	}
}
