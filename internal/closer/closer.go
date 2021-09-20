// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
