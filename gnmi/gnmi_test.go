// Copyright 2022 Google LLC
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

package gnmi

import (
	"fmt"
	"io"
	"testing"

	"github.com/openconfig/testt"
	"github.com/openconfig/ygnmi/ygnmi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakeWatcher[T any] struct {
	err error
}

func (fw *fakeWatcher[T]) Await() (*ygnmi.Value[T], error) {
	return nil, fw.err
}

func TestWatcherAwait(t *testing.T) {
	tests := []struct {
		desc       string
		inErr      error
		wantStatus bool
		wantErr    bool
	}{{
		desc:  "wrapped deadline exceeded",
		inErr: fmt.Errorf("rpc err: %w", status.Error(codes.DeadlineExceeded, "foo")),
	}, {
		desc:  "wrapped canceled",
		inErr: fmt.Errorf("rpc err: %w", status.Error(codes.Canceled, "foo")),
	}, {
		desc:    "wrapped EOF",
		inErr:   fmt.Errorf("rpc err: %w", io.EOF),
		wantErr: true,
	}, {
		desc:    "other error type",
		inErr:   fmt.Errorf("rpc err"),
		wantErr: true,
	}, {
		desc:       "no error",
		wantStatus: true,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			w := &Watcher[int]{
				watcher: &fakeWatcher[int]{
					err: tt.inErr,
				},
			}
			var gotStatus bool
			err := testt.CaptureFatal(t, func(t testing.TB) {
				_, gotStatus = w.Await(t)
			})
			if err != nil && !tt.wantErr {
				t.Fatalf("Await() returned unexpected error: got %v want none", err)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("Await() returned unexpected value: got %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

type fakeCollector[T any] struct {
	err error
}

func (fw *fakeCollector[T]) Await() ([]*ygnmi.Value[T], error) {
	return nil, fw.err
}

func TestCollectorAwait(t *testing.T) {
	tests := []struct {
		desc       string
		inErr      error
		wantStatus bool
		wantErr    bool
	}{{
		desc:  "wrapped deadline exceeded",
		inErr: fmt.Errorf("rpc err: %w", status.Error(codes.DeadlineExceeded, "foo")),
	}, {
		desc:  "wrapped canceled",
		inErr: fmt.Errorf("rpc err: %w", status.Error(codes.Canceled, "foo")),
	}, {
		desc:    "wrapped EOF",
		inErr:   fmt.Errorf("rpc err: %w", io.EOF),
		wantErr: true,
	}, {
		desc:    "other error type",
		inErr:   fmt.Errorf("rpc err"),
		wantErr: true,
	}, {
		desc: "no error",
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			w := &Collector[int]{
				collector: &fakeCollector[int]{
					err: tt.inErr,
				},
			}
			err := testt.CaptureFatal(t, func(t testing.TB) {
				w.Await(t)
			})
			if err != nil && !tt.wantErr {
				t.Fatalf("Await() returned unexpected error: got %v want none", err)
			}
		})
	}
}
