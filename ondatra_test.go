// Copyright 2021 Google LLC
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

package ondatra

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/flags"
)

func TestReserveOnRun(t *testing.T) {
	flagParseFn = func() (*flags.Values, error) {
		return &flags.Values{}, nil
	}
	origRunTests := runTestsFn
	defer func() {
		reserveFn = reserve
		releaseFn = release
		runTestsFn = origRunTests
	}()
	tests := []struct {
		desc                  string
		sig                   os.Signal
		reservErr, releaseErr error
		wantErr               string
	}{{
		desc:      "error on reserve",
		reservErr: errors.New("error reserving testbed"),
		wantErr:   "error reserving testbed",
	}, {
		desc:       "error on release",
		releaseErr: errors.New("error releasing testbed"),
		wantErr:    "error releasing testbed",
	}, {
		desc: "release on signal",
		sig:  os.Interrupt,
	}, {
		desc: "release on test completion",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			reserveFn = func(*flags.Values) error {
				return test.reservErr
			}
			var releaseMu sync.Mutex
			var released bool
			releaseCh := make(chan struct{}, 1)
			releaseErr := test.releaseErr
			defer close(releaseCh)
			releaseFn = func() error {
				releaseMu.Lock()
				defer releaseMu.Unlock()
				if !released {
					released = true
					releaseCh <- struct{}{}
				}
				return releaseErr
			}
			runTestsFn = func(*fixture, *testing.M, time.Duration) {
				if test.sig != nil {
					sigc <- test.sig
				}
			}
			var initBindCalled bool
			initBindFn = func(b binding.Binding) {
				initBindCalled = true
			}
			fakeBinder := func() (binding.Binding, error) { return nil, nil }
			if gotErr := doRun(nil, fakeBinder); (gotErr != nil) != (test.wantErr != "") {
				t.Fatalf("doRun: got err %v, wanted err? %t", gotErr, test.wantErr != "")
			}
			if !initBindCalled {
				t.Errorf("doRun did not initialize the binding")
			}
			wantReleased := test.reservErr == nil
			if wantReleased {
				<-releaseCh
			}
			releaseMu.Lock()
			defer releaseMu.Unlock()
			if released != wantReleased {
				t.Errorf("doRun: testbed released? %t, wanted testbed released? %t", released, wantReleased)
			}
		})
	}
}
