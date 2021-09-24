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

package genutil

import (
	"fmt"
	"testing"

	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// FakeRootPathStruct is an interface that is implemented by the fake root path
// struct type.
type FakeRootPathStruct interface {
	ygot.PathStruct
	Id() string
	CustomData() map[string]interface{}
	PutCustomData(key string, val interface{})
}

// SetRequestBatch implements batch operations for doing a gNMI SetRequest
// based on a fake-root PathStruct object.
type SetRequestBatch struct {
	deviceRoot FakeRootPathStruct
	req        *gpb.SetRequest
}

// NewSetRequestBatch creates a new SetRequestBatch object based on a
// deviceRoot PathStruct object to which to send SetRequests.
func NewSetRequestBatch(deviceRoot FakeRootPathStruct) *SetRequestBatch {
	return &SetRequestBatch{
		deviceRoot: deviceRoot,
		req:        &gpb.SetRequest{},
	}
}

// WithReplica adds the replica number to the context metadata of the gNMI
// server query.
func (b *SetRequestBatch) WithReplica(replica int) *SetRequestBatch {
	PutReplica(b.deviceRoot, replica)
	return b
}

// WithSubscriptionMode specifies the subscription mode in the underlying gNMI
// subscribe.
func (b *SetRequestBatch) WithSubscriptionMode(mode gpb.SubscriptionMode) *SetRequestBatch {
	PutSubscriptionMode(b.deviceRoot, mode)
	return b
}

// Set creates and makes a single gNMI SetRequest call for the batched set requests.
func (b *SetRequestBatch) Set(t testing.TB) *gpb.SetResponse {
	t.Helper()
	resp, err := batchSet("openconfig", b.deviceRoot.Id(), b.deviceRoot.CustomData(), b.req)
	if err != nil {
		t.Fatalf("SetRequestBatch.Set: %v", err)
	}
	return resp
}

// Reset clears any buffered requests.
func (b *SetRequestBatch) Reset() {
	b.req = &gpb.SetRequest{}
}

// BatchDelete buffers a delete operation in the SetRequestBatch.
func (b *SetRequestBatch) BatchDelete(t testing.TB, n ygot.PathStruct) {
	t.Helper()
	if err := b.batchSet(n, nil, updatePath); err != nil {
		t.Fatal(err)
	}
}

// BatchReplace buffers an replace operation in the SetRequestBatch.
func (b *SetRequestBatch) BatchReplace(t testing.TB, n ygot.PathStruct, val interface{}) {
	t.Helper()
	if err := b.batchSet(n, val, replacePath); err != nil {
		t.Fatal(err)
	}
}

// BatchUpdate buffers an update operation in the SetRequestBatch.
func (b *SetRequestBatch) BatchUpdate(t testing.TB, n ygot.PathStruct, val interface{}) {
	t.Helper()
	if err := b.batchSet(n, val, updatePath); err != nil {
		t.Fatal(err)
	}
}

// batchSet buffers a replace Update object in the SetRequestBatch.
func (b *SetRequestBatch) batchSet(n ygot.PathStruct, val interface{}, op setOperation) error {
	path, customData, errs := ygot.ResolvePath(n)
	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	if path.GetTarget() != b.deviceRoot.Id() {
		return fmt.Errorf("path target doesn't equal the device target for the batch: got %q, want %q", path.GetTarget(), b.deviceRoot.Id())
	}
	if len(customData) != 0 {
		return fmt.Errorf("batching cannot accept a path that has its custom request options; please set request options solely via the batch object")
	}

	if err := populateSetRequest(b.req, path, val, op); err != nil {
		return err
	}
	return nil
}

// setOperation is an enum representing the different kinds of SetRequest
// operations available.
type setOperation int

const (
	// deletePath represents a SetRequest delete.
	deletePath setOperation = iota
	// replacePath represents a SetRequest replace.
	replacePath
	// updatePath represents a SetRequest update.
	updatePath
)

func populateSetRequest(req *gpb.SetRequest, path *gpb.Path, val interface{}, op setOperation) error {
	if req == nil {
		return fmt.Errorf("cannot populate a nil SetRequest")
	}
	if len(path.Elem) == 0 {
		return fmt.Errorf("got empty path for replace operation")
	}

	// Keep only the path elements.
	path = &gpb.Path{Elem: path.Elem}
	switch op {
	case deletePath:
		req.Delete = append(req.Delete, path)
	case replacePath, updatePath:
		// Since the GoStructs are generated using preferOperationalState, we
		// need to turn on preferShadowPath to prefer marshalling config paths.
		js, err := ygot.Marshal7951(val, ygot.JSONIndent("  "), &ygot.RFC7951JSONConfig{AppendModuleName: true, PreferShadowPath: true})
		if err != nil {
			return fmt.Errorf("could not encode value into JSON format: %w", err)
		}
		update := &gpb.Update{
			Path: path,
			Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{js}},
		}
		switch op {
		case replacePath:
			req.Replace = append(req.Replace, update)
		case updatePath:
			req.Update = append(req.Update, update)
		}
	}

	return nil
}
