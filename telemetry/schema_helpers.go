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

// Package telemetry contains generated code to perform gNMI Get operations on devices.
package telemetry

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// BinarySliceToFloat32 although ideally should be part of the ygot library,
// lives here due to "[]Binary", a type defined in the ygen-generated code,
// not being able to be assigned to [][]byte, preventing such a function from being used.
func BinarySliceToFloat32(in []Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// Lookup uses gNMI Get to fill the input GoStruct with values at the input path.
// This func is used by generated code and doesn't need to be call directly.
func Lookup(t testing.TB, n ygot.PathStruct, goStructName string, gs ygot.GoStruct, isLeaf, preferShadowPath bool) (*genutil.Metadata, bool) {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	md, ok, err := genutil.Unmarshal(datapoints, GetSchema(), goStructName, gs, queryPath, isLeaf, preferShadowPath)
	if md.ComplianceErrors != nil {
		if isLeaf {
			t.Fatalf("noncompliant data encountered while unmarshalling leaf: %v", md.ComplianceErrors)
		}
		log.V(0).Infof("noncompliant data encountered while unmarshalling: %v", md.ComplianceErrors)
	}
	if err != nil {
		t.Fatal(err)
	}
	return md, ok
}

// GetSchema return the generated ytypes schema used for unmarshaling datapoints.
// This func is used by generated code and doesn't need to be call directly.
func GetSchema() *ytypes.Schema {
	return &ytypes.Schema{
		Root:       &Device{},
		SchemaTree: SchemaTree,
		Unmarshal:  Unmarshal,
	}
}

// BatchWatcher observes a stream of Device samples.
type BatchWatcher struct {
	W       *genutil.Watcher
	lastVal *QualifiedDevice
}

// Await waits until predicate becomes true or the watch times out.
// The final merged device and the status of the predicate are returned.
func (u *BatchWatcher) Await(t testing.TB) (*QualifiedDevice, bool) {
	t.Helper()
	return u.lastVal, u.W.Await(t)
}

// BatchCollection is a telemetry Collection whose Await method returns a slice of Device samples.
type BatchCollection struct {
	W    *genutil.Watcher
	vals []*QualifiedDevice
}

// Await blocks for the telemetry collection to be complete, and then returns the slice of samples received.
func (u *BatchCollection) Await(t testing.TB) []*QualifiedDevice {
	t.Helper()
	u.W.Await(t)
	return u.vals
}

// NewBatch creates a new batch object.
// This doesn't need to be called directly. Use dut.Telemetry().NewBatch() instead.
func NewBatch(root genutil.FakeRootPathStruct) *Batch {
	return &Batch{
		root: root,
	}
}

// Batch allows gNMI operations on multiple paths in the at the same time.
type Batch struct {
	paths []*gpb.Path
	root  genutil.FakeRootPathStruct
}

// MustAddToBatch calls Add and fails the calling test fatally on error.
func MustAddToBatch(t testing.TB, b *Batch, p ygot.PathStruct) {
	t.Helper()
	if err := addToBatch(b, p); err != nil {
		t.Fatal(err)
	}
}

// addToBatch adds the path struct to the batch object
func addToBatch(b *Batch, p ygot.PathStruct) error {
	path, customData, err := genutil.ResolvePath(p)
	if err != nil {
		return err
	}
	if _, ok := customData[genutil.DefaultClientKey]; ok && len(customData) > 1 || !ok && len(customData) > 0 {
		return fmt.Errorf("batch cannot accept a path that has custom request options; please set request options solely via the batch object")
	}
	if path.GetTarget() != b.root.Id() {
		return fmt.Errorf("path target %q doesn't match batch target %q", path.GetTarget(), b.root.Id())
	}
	b.paths = append(b.paths, path)
	return nil
}

// Reset clears all paths from the batch object.
func (b *Batch) Reset() {
	b.paths = nil
}

var fakeRootName = reflect.TypeOf(Device{}).Name()

// Lookup gets a value sample for all paths added to the batch.
func (b *Batch) Lookup(t testing.TB) *QualifiedDevice {
	t.Helper()
	datapoints, _ := genutil.MustGet(t, b.root, b.paths...)
	gs := &Device{}
	path, _, err := genutil.ResolvePath(b.root)
	if err != nil {
		t.Fatalf("error resolving root path: %v", err)
	}

	md, ok, err := genutil.Unmarshal(datapoints, GetSchema(), fakeRootName, gs, path, false, true)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		return nil
	}
	return (&QualifiedDevice{
		Metadata: md,
	}).SetVal(gs)
}

// Watch subscribes to all paths in the batch and evaluates the predicate on each received notification.
// The predicate is evaluated on result of merging the data in the notification with the existing values.
// Note: The Device object is reused between each evaluation of the predicate, so make a copy using ygot.DeepCopy
// if you need to do anything beyond reading the struct within the timeframe of the predicate's evaluation.
func (b *Batch) Watch(t testing.TB, duration time.Duration, predicate func(*QualifiedDevice) bool) *BatchWatcher {
	t.Helper()
	dev := &Device{}
	w := &BatchWatcher{}
	w.W = genutil.MustWatch(t, b.root, b.paths, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, GetSchema(), fakeRootName, dev, queryPath, false, false)
		qd := (&QualifiedDevice{
			Metadata: md,
		}).SetVal(dev)
		return qd, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedDevice)
		w.lastVal = val
		return ok && predicate(val)
	})
	return w
}

// Collect retrieves a Collection sample for the paths in the batch object.
// Note: this func deep copies the device object on every notification,
// consider using Watch instead.
func (b *Batch) Collect(t testing.TB, duration time.Duration) *BatchCollection {
	t.Helper()
	dev := &Device{}
	bc := &BatchCollection{}
	bc.W = genutil.MustWatch(t, b.root, b.paths, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		copy, err := ygot.DeepCopy(dev)
		if err != nil {
			t.Fatal(err)
		}
		md, ok, err := genutil.Unmarshal(upd, GetSchema(), fakeRootName, copy, queryPath, false, false)
		if err != nil || !ok {
			return nil, err
		}
		dev = copy.(*Device)
		qd := (&QualifiedDevice{
			Metadata: md,
		}).SetVal(dev)
		return qd, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		bc.vals = append(bc.vals, qualVal.(*QualifiedDevice))
		return false
	})
	return bc
}
