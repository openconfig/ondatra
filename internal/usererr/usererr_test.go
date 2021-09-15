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

package usererr

import (
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

func TestIn(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		wantIn bool
	}{{
		name:   "new",
		err:    New("foo"),
		wantIn: true,
	}, {
		name:   "wrap",
		err:    Wrap(errors.New("foo")),
		wantIn: true,
	}, {
		name:   "wrapped with errors.Wrap",
		err:    errors.Wrap(New("foo"), "bar"),
		wantIn: true,
	}, {
		name:   "wrapped with xerrors %w",
		err:    xerrors.Errorf("foo: %w", New("bar")),
		wantIn: true,
	}, {
		name:   "wrapped multiple times",
		err:    errors.Wrap(xerrors.Errorf("foo: %w", New("bar")), "baz"),
		wantIn: true,
	}, {
		name:   "not user err",
		err:    errors.New("foo"),
		wantIn: false,
	}, {
		name:   "not user error - wrapped",
		err:    errors.Wrap(xerrors.Errorf("foo: %w", errors.New("bar")), "baz"),
		wantIn: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.err); got != tt.wantIn {
				t.Errorf("In(%v) got %v, want %v", tt.err, got, tt.wantIn)
			}
		})
	}
}
