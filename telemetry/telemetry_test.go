// Copyright 2021 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package telemetry

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/gnmi/errdiff"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

func TestAddToBatch(t *testing.T) {
	pathWithCustomData := ygot.NewDeviceRootBase("test")
	pathWithCustomData.PutCustomData("fake", "data")

	tests := []struct {
		desc    string
		want    []*gpb.Path
		wantErr string
		path    ygot.PathStruct
	}{{
		desc:    "path with custom data",
		path:    pathWithCustomData,
		wantErr: "batch cannot accept a path that has custom request options",
	}, {
		desc:    "path with invalid target",
		path:    ygot.NewDeviceRootBase("bad id"),
		wantErr: "path target \"bad id\" doesn't match batch target \"test\"",
	}, {
		desc: "success adding path",
		path: ygot.NewDeviceRootBase("test"),
		want: []*gpb.Path{{
			Target: "test",
			Origin: "openconfig",
		}},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			b := &Batch{
				root: ygot.NewDeviceRootBase("test"),
			}
			err := addToBatch(b, tt.path)
			if diff := errdiff.Substring(err, tt.wantErr); diff != "" {
				t.Errorf("AddPath(%v) returned unexpected error diff:\n%s", tt.path, diff)
			}
			if diff := cmp.Diff(tt.want, b.paths, protocmp.Transform()); diff != "" {
				t.Errorf("AddPath(%v) unexpected path diff (-want +got):\n%s", tt.path, diff)
			}
		})
	}
}
