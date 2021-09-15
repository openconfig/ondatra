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

package telemgo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/openconfig/ygot/ygot"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// DevicePath represents the /device YANG schema element.
type DevicePath struct {
	*ygot.DeviceRootBase
}

// WithReplica adds the replica number to the context metadata of the gNMI
// server query.
func (n *DevicePath) WithReplica(replica int) *DevicePath {
	PutReplica(n, replica)
	return n
}

// DeviceRoot returns a new path object from which YANG paths can be constructed.
func DeviceRoot(id string) *DevicePath {
	return &DevicePath{ygot.NewDeviceRootBase(id)}
}

// InterfacePath represents the /openconfig-interfaces/interfaces/interface YANG schema element.
type InterfacePath struct {
	*ygot.NodePath
}

// InterfaceDescriptionPath represents the /openconfig-interfaces/interfaces/interface/config/description YANG schema element.
type InterfaceDescriptionPath struct {
	*ygot.NodePath
}

// Interface returns from DevicePath the path struct for its child "interface".
// Name: string
func (n *DevicePath) Interface(Name string) *InterfacePath {
	return &InterfacePath{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Description returns from InterfacePath the path struct for its child "description".
func (n *InterfacePath) Description() *InterfaceDescriptionPath {
	return &InterfaceDescriptionPath{
		NodePath: ygot.NewNodePath(
			[]string{"config", "description"},
			map[string]interface{}{},
			n,
		),
	}
}

func mustPath(t *testing.T, s string) *gpb.Path {
	t.Helper()
	p, err := ygot.StringToStructuredPath(s)
	if err != nil {
		t.Fatalf("cannot converting string %s to path, got err: %v", s, err)
	}
	return p
}

type batchSetInput struct {
	path ygot.PathStruct
	val  interface{}
	op   setOperation
}

func TestBatchSetAndReset(t *testing.T) {
	tests := []struct {
		desc    string
		inOps   []batchSetInput
		wantReq *gpb.SetRequest
	}{{
		desc: "replace",
		inOps: []batchSetInput{{
			DeviceRoot("42").Interface("eth1").Description(),
			ygot.String("foo"),
			replacePath,
		}},
		wantReq: &gpb.SetRequest{
			Replace: []*gpb.Update{{
				Path: mustPath(t, "/interfaces/interface[name=eth1]/config/description"),
				Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{[]byte(`"foo"`)}},
			}},
		},
	}, {
		desc: "replace-update-delete-replace",
		inOps: []batchSetInput{{
			DeviceRoot("42").Interface("eth1").Description(),
			ygot.String("foo"),
			replacePath,
		}, {
			DeviceRoot("42").Interface("eth2").Description(),
			ygot.String("foo"),
			updatePath,
		}, {
			DeviceRoot("42").Interface("eth3"),
			nil,
			deletePath,
		}, {
			DeviceRoot("42").Interface("eth4").Description(),
			ygot.String("foo"),
			replacePath,
		}},
		wantReq: &gpb.SetRequest{
			Delete: []*gpb.Path{
				mustPath(t, "/interfaces/interface[name=eth3]"),
			},
			Replace: []*gpb.Update{{
				Path: mustPath(t, "/interfaces/interface[name=eth1]/config/description"),
				Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{[]byte(`"foo"`)}},
			}, {
				Path: mustPath(t, "/interfaces/interface[name=eth4]/config/description"),
				Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{[]byte(`"foo"`)}},
			}},
			Update: []*gpb.Update{{
				Path: mustPath(t, "/interfaces/interface[name=eth2]/config/description"),
				Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{[]byte(`"foo"`)}},
			}},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			batch := NewSetRequestBatch(DeviceRoot("42"))
			test := func() {
				for _, op := range tt.inOps {
					err := batch.batchSet(op.path, op.val, op.op)
					if err != nil {
						t.Fatal(err)
					}
				}
				if diff := cmp.Diff(batch.req, tt.wantReq, protocmp.Transform()); diff != "" {
					t.Errorf("(-got, +want):\n%s", diff)
				}
			}
			test()
			batch.Reset()
			if diff := cmp.Diff(batch.req, &gpb.SetRequest{}, protocmp.Transform()); diff != "" {
				t.Errorf("after Reset (-got, +want):\n%s", diff)
			}
			test()
		})
	}
}

func TestBatchReplaceAndReset_BadPath(t *testing.T) {
	tests := []struct {
		desc string
		inOp batchSetInput
	}{{
		desc: "device ID doesn't match up",
		inOp: batchSetInput{
			DeviceRoot("43").Interface("eth1").Description(),
			ygot.String("foo"),
			replacePath,
		},
	}, {
		desc: "has custom data",
		inOp: batchSetInput{
			DeviceRoot("42").WithReplica(5).Interface("eth1").Description(),
			ygot.String("foo"),
			updatePath,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			batch := NewSetRequestBatch(DeviceRoot("42"))
			makeReplace := func() {
				if err := batch.batchSet(tt.inOp.path, tt.inOp.val, tt.inOp.op); err == nil {
					t.Fatalf("got nil error")
				}
			}
			makeReplace()
			batch.Reset()
			if diff := cmp.Diff(batch.req, &gpb.SetRequest{}, protocmp.Transform()); diff != "" {
				t.Errorf("after Reset (-got, +want):\n%s", diff)
			}
			makeReplace()
		})
	}
}
