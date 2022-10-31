// Copyright 2019 Google LLC
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

package telemetry_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/openconfig/gnmi/errdiff"
	cfgdev "github.com/openconfig/ondatra/config/device"
	"github.com/openconfig/ondatra/internal/fakegnmi"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/ondatra/telemetry/device"
	"github.com/openconfig/ondatra/telemetry/interfaces"
	"github.com/openconfig/testt"
	"github.com/openconfig/ygot/testutil"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"google.golang.org/protobuf/testing/protocmp"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var fakeGNMI = func() *fakegnmi.FakeGNMI {
	fg, err := fakegnmi.Start(0)
	if err != nil {
		log.Fatalf("could not start fake gNMI server: %v", err)
	}
	return fg
}()

func telemRoot() *device.DevicePath {
	root := device.DeviceRoot("fakeDUT")
	root.PutCustomData(genutil.DefaultClientKey, func(ctx context.Context) (gpb.GNMIClient, error) {
		return fakeGNMI.Dial(ctx)
	})
	return root
}

func configRoot() *cfgdev.DevicePath {
	root := cfgdev.DeviceRoot("fakeDUT")
	root.PutCustomData(genutil.DefaultClientKey, func(ctx context.Context) (gpb.GNMIClient, error) {
		return fakeGNMI.Dial(ctx)
	})
	return root
}

func gnmiPath(t *testing.T, s string, args ...any) *gpb.Path {
	t.Helper()
	p, err := ygot.StringToStructuredPath(fmt.Sprintf(s, args...))
	if err != nil {
		t.Fatal(err)
	}
	if len(p.GetElem()) == 0 || p.GetElem()[0].GetName() != "meta" {
		p.Origin = "openconfig"
	}
	return p
}

// verifySubscriptionPathsSent verifies the paths of the sent subscription requests is the same as wantPaths.
func verifySubscriptionPathsSent(t *testing.T, wantPaths ...*gpb.Path) {
	t.Helper()
	requests := fakeGNMI.Requests()
	if len(requests) != 1 {
		t.Errorf("Number of subscription requests sent is not 1: %v", requests)
		return
	}

	var gotPaths []*gpb.Path
	req := requests[0].GetSubscribe()
	for _, sub := range req.GetSubscription() {
		got, err := util.JoinPaths(req.GetPrefix(), sub.GetPath())
		if err != nil {
			t.Fatal(err)
		}
		got.Target = ""
		gotPaths = append(gotPaths, got)
	}
	if diff := cmp.Diff(wantPaths, gotPaths, protocmp.Transform(), cmpopts.SortSlices(testutil.PathLess)); diff != "" {
		t.Errorf("subscription paths (-want, +got):\n%s", diff)
	}
}

// verifyTelemetryErrSubstr verifies the error components of a given
// ComplianceErrors value given the consitutent error substrings.
func verifyTelemetryErrSubstr(t *testing.T, complianceErrs *genutil.ComplianceErrors, wantPathErrSubstr, wantTypeErrSubstr string) {
	t.Helper()
	var pathErrs, typeErrs []*genutil.TelemetryError
	var validateErrs []error
	if complianceErrs != nil {
		pathErrs = complianceErrs.PathErrors
		typeErrs = complianceErrs.TypeErrors
		validateErrs = complianceErrs.ValidateErrors
	}
	if len(pathErrs) > 1 {
		t.Fatalf("unmarshal: got more than one path unmarshal error: %v", pathErrs)
	}
	if len(typeErrs) > 1 {
		t.Fatalf("unmarshal: got more than one type unmarshal error: %v", typeErrs)
	}
	if len(validateErrs) > 1 {
		t.Fatalf("unmarshal: got more than one validate error: %v", validateErrs)
	}

	// Populate errors for validation.
	var pathErr, typeErr error
	if len(pathErrs) == 1 {
		pathErr = pathErrs[0].Err
	}
	if len(typeErrs) == 1 {
		typeErr = typeErrs[0].Err
	}

	// Validate expected errors
	if diff := errdiff.Substring(pathErr, wantPathErrSubstr); diff != "" {
		t.Fatalf("unmarshal: did not get expected path Err substring:\n%s", diff)
	}
	if diff := errdiff.Substring(typeErr, wantTypeErrSubstr); diff != "" {
		t.Fatalf("unmarshal: did not get expected type Err substring:\n%s", diff)
	}
}

func TestMetadata(t *testing.T) {
	connPath := gnmiPath(t, "meta/connected")
	syncPath := gnmiPath(t, "meta/sync")
	leavesAddedPath := gnmiPath(t, "meta/targetLeavesAdded")

	testsPass := []struct {
		desc        string
		stubFn      func(s *fakegnmi.Stubber)
		lookupFn    func(t testing.TB) genutil.QualifiedValue
		wantSubPath *gpb.Path
		wantValue   genutil.QualifiedValue
	}{{
		desc: "check connected",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: connPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: func(t testing.TB) genutil.QualifiedValue {
			return telemRoot().Meta().Connected().Lookup(t)
		},
		wantSubPath: connPath,
		wantValue: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      connPath,
			}}).SetVal(true),
	}, {
		desc: "check connected -- unconnected",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: connPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: false}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: func(t testing.TB) genutil.QualifiedValue {
			return telemRoot().Meta().Connected().Lookup(t)
		},
		wantSubPath: connPath,
		wantValue: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      connPath,
			}}).SetVal(false),
	}, {
		desc: "check sync",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: syncPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: func(t testing.TB) genutil.QualifiedValue {
			return telemRoot().Meta().Sync().Lookup(t)
		},
		wantSubPath: syncPath,
		wantValue: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      syncPath,
			}}).SetVal(true),
	}, {
		desc: "check sync -- unsynced",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: syncPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: false}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: func(t testing.TB) genutil.QualifiedValue {
			return telemRoot().Meta().Sync().Lookup(t)
		},
		wantSubPath: syncPath,
		wantValue: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      syncPath,
			}}).SetVal(false),
	}, {
		desc: "check leaves added",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: leavesAddedPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
					}},
					Timestamp: 200,
				}).
				Sync()
		},
		lookupFn: func(t testing.TB) genutil.QualifiedValue {
			return telemRoot().Meta().TargetLeavesAdded().Lookup(t)
		},
		wantSubPath: leavesAddedPath,
		wantValue: (&telemetry.QualifiedInt64{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 200),
				Path:      leavesAddedPath,
			}}).SetVal(42),
	}, {
		desc: "no values for leaves added",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(&gpb.Notification{}).Sync()
		},
		lookupFn: func(t testing.TB) genutil.QualifiedValue {
			ret := telemRoot().Meta().TargetLeavesAdded().Lookup(t)
			if ret != nil { // avoid typed-nil
				return ret
			}
			return nil
		},
		wantSubPath: leavesAddedPath,
		wantValue:   nil,
	}, {
		desc: "no values for leaves added without connected return value",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		lookupFn: func(t testing.TB) genutil.QualifiedValue {
			ret := telemRoot().Meta().TargetLeavesAdded().Lookup(t)
			if ret != nil { // avoid typed-nil
				return ret
			}
			return nil
		},
		wantSubPath: leavesAddedPath,
		wantValue:   nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := tt.lookupFn(t)
			verifySubscriptionPathsSent(t, tt.wantSubPath)
			if got != nil {
				checkJustReceived(t, got.GetRecvTimestamp())
			}
			if diff := cmp.Diff(tt.wantValue, got, cmpopts.IgnoreFields(telemetry.QualifiedBool{}, "RecvTimestamp"), cmpopts.IgnoreFields(telemetry.QualifiedInt64{}, "RecvTimestamp"), cmp.AllowUnexported(telemetry.QualifiedBool{}, telemetry.QualifiedInt64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	intfName := "Et1/2/3"
	intfPath := gnmiPath(t, "interfaces/interface[name=%s]/state/oper-status", intfName)

	testsPass := []struct {
		desc         string
		stubFn       func(s *fakegnmi.Stubber)
		wantValue    *telemetry.QualifiedE_Interface_OperStatus
		wantValFatal string
	}{{
		desc: "one notification",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "no sync response",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				})
		},
		wantValue: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "one notification with interpolation",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "one notification with a prefix",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intfPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}},
					Prefix:    &gpb.Path{Origin: intfPath.GetOrigin(), Elem: intfPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(telemetry.Interface_OperStatus_DOWN),
	}, {
		desc: "one notification with a prefix with interpolation",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intfPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}},
					Prefix:    &gpb.Path{Origin: intfPath.GetOrigin(), Elem: intfPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(telemetry.Interface_OperStatus_DOWN),
	}, {
		desc: "multiple notifications with the first one having no update value",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(&gpb.Notification{Update: []*gpb.Update{}}).
				Notification(&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 101,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 101),
				Path:      intfPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		wantValue:    nil,
		wantValFatal: "No value present\n",
	}, {
		desc: "with connection",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 102,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 102),
				Path:      intfPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := telemRoot().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(intfName).OperStatus().Lookup(t)
			verifySubscriptionPathsSent(t, intfPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedE_Interface_OperStatus{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
			fatalMsg := testt.CaptureFatal(t, func(t testing.TB) {
				if diff := cmp.Diff(tt.wantValue.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			})
			if fatalMsg != nil && *fatalMsg != tt.wantValFatal {
				t.Errorf("Got fatal msg different from expected: got %q, want %q", *fatalMsg, tt.wantValFatal)
			}
		})
	}

	testsFail := []struct {
		desc         string
		stubFn       func(s *fakegnmi.Stubber)
		wantFatalMsg string
	}{{
		desc: "multiple values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}}}).
				Notification(
					&gpb.Notification{
						Update: []*gpb.Update{{
							Path: intfPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
						}}}).
				Sync()
		},
		wantFatalMsg: "got multiple",
	}, {
		desc: "returned path uses deprecated Element field",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: func() *gpb.Path {
							pathStr := fmt.Sprintf("interfaces/interface[name=%s]/state/oper-status", intfName)
							p, err := ygot.StringToStringSlicePath(pathStr)
							if err != nil {
								t.Fatal(err)
							}
							p.Origin = "openconfig"
							return p
						}(),
						Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
				}).
				Sync()
		},
		wantFatalMsg: "deprecated and unsupported Element field",
	}, {
		desc: "last element is different than the query path",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: gnmiPath(t, "interfaces/interface[name=%s]/state/description", intfName),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
				}).
				Sync()
		},
		wantFatalMsg: "does not match the query path",
	}, {
		desc: "returned prefix + path is incompatible",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intfPath.GetElem()[len(intfPath.GetElem())-1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					// missing first element.
					Prefix: &gpb.Path{Elem: intfPath.GetElem()[1 : len(intfPath.GetElem())-1]},
				}).
				Sync()
		},
		wantFatalMsg: "does not match the query path",
	}, {
		desc: "returned path is incompatible",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: append(intfPath.GetElem(), &gpb.PathElem{Name: "does-not-exist"})},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
				}).
				Sync()
		},
		// The path cannot be unmarshalled because it doesn't exist in the schema.
		wantFatalMsg: "does-not-exist",
	}, {
		desc: "one invalid notification where Val of Update is nil",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  nil,
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantFatalMsg: "invalid nil Val",
	}, {
		desc: "schema noncompliance: one notification that doesn't unmarshal because type doesn't match",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantFatalMsg: "failed to unmarshal",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := testt.ExpectFatal(t, func(t testing.TB) {
				telemRoot().Interface(intfName).OperStatus().Lookup(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Lookup failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestGetDefault(t *testing.T) {
	intfName := "Ethernet3/1/1"
	statePath := gnmiPath(t, "interfaces/interface[name=%s]/state/enabled", intfName)
	configPath := gnmiPath(t, "interfaces/interface[name=%s]/config/enabled", intfName)

	testsPass := []struct {
		desc      string
		stubFn    func(s *fakegnmi.Stubber)
		lookupFn  func(t testing.TB) *telemetry.QualifiedBool
		wantPath  *gpb.Path
		wantValue *telemetry.QualifiedBool
	}{{
		desc: "one notification",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statePath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: false}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: telemRoot().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(intfName).Enabled().Lookup,
		wantPath: statePath,
		wantValue: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      statePath,
			}}).SetVal(false),
	}, {
		desc: "telemetry path API: zero notifications on a leaf that has a default value",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		lookupFn: telemRoot().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(intfName).Enabled().Lookup,
		wantPath: statePath,
		wantValue: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Path: statePath,
			}}).SetVal(true),
	}, {
		desc: "config path API: zero notifications on a leaf that has a default value",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		lookupFn: configRoot().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(intfName).Enabled().Lookup,
		wantPath: configPath,
		wantValue: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Config: true,
				Path:   configPath,
			}}).SetVal(true),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := tt.lookupFn(t)
			verifySubscriptionPathsSent(t, tt.wantPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedBool{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantValue.IsPresent() {
				if diff := cmp.Diff(tt.wantValue.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	intfName := "Et1/2/3"
	intfPath := gnmiPath(t, "interfaces/interface[name=%s]/config/description", intfName)

	testsPass := []struct {
		desc      string
		stubFn    func(s *fakegnmi.Stubber)
		wantValue *telemetry.QualifiedString
	}{{
		desc: "one notification",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal("foo"),
	}, {
		desc: "one notification with interpolation",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal("foo"),
	}, {
		desc: "one notification with a prefix",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intfPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bar"}},
					}},
					Prefix:    &gpb.Path{Origin: intfPath.GetOrigin(), Elem: intfPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal("bar"),
	}, {
		desc: "one notification with a prefix with interpolation",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intfPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bar"}},
					}},
					Prefix:    &gpb.Path{Origin: intfPath.GetOrigin(), Elem: intfPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal("bar"),
	}, {
		desc: "multiple notifications with the first one having no update value",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(&gpb.Notification{}).
				Notification(&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					Timestamp: 101,
				}).
				Sync()
		},
		wantValue: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 101),
				Path:      intfPath,
			}}).SetVal("foo"),
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		wantValue: nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := configRoot().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(intfName).Description().Lookup(t)
			verifySubscriptionPathsSent(t, intfPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedString{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got desc Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantValue.IsPresent() {
				if diff := cmp.Diff(tt.wantValue.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got desc val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}

	testsFail := []struct {
		desc         string
		stubFn       func(s *fakegnmi.Stubber)
		wantFatalMsg string
	}{{
		desc: "multiple values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
				}).
				Notification(
					&gpb.Notification{
						Update: []*gpb.Update{{
							Path: intfPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bar"}},
						}},
					}).
				Sync()
		},
		wantFatalMsg: "got multiple",
	}, {
		desc: "one invalid notification where Val of Update is nil",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  nil,
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantFatalMsg: "invalid nil Val",
	}, {
		desc: "schema noncompliance: one notification that doesn't unmarshal because type doesn't match",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantFatalMsg: "failed to unmarshal",
	}, {
		desc: "unmarshal fails because returned prefix and path don't match",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intfPath.GetElem()[len(intfPath.GetElem())-1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					// missing first element.
					Prefix: &gpb.Path{Elem: intfPath.GetElem()[1 : len(intfPath.GetElem())-1]},
				}).
				Sync()
		},
		wantFatalMsg: "does not match the query path",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := testt.ExpectFatal(t, func(t testing.TB) {
				configRoot().Interface(intfName).Description().Lookup(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Lookup failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}

	t.Run("use get for config", func(t *testing.T) {
		root := configRoot()
		genutil.PutUseGetForConfig(root, true)
		fakeGNMI.Stub().GetResponse(&gpb.GetResponse{
			Notification: []*gpb.Notification{{
				Timestamp: 100,
				Update: []*gpb.Update{{
					Path: intfPath,
					Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{JsonIetfVal: []byte(`"foo"`)}},
				}},
			}},
		})

		got := root.Interface(intfName).Description().Lookup(t)
		checkJustReceived(t, got.RecvTimestamp)
		wantValue := (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:        true,
				Timestamp:     time.Unix(0, 100),
				Path:          intfPath,
				RecvTimestamp: got.RecvTimestamp,
			}}).SetVal("foo")
		wantGetRequest := &gpb.GetRequest{
			Path:     []*gpb.Path{intfPath},
			Type:     gpb.GetRequest_CONFIG,
			Encoding: gpb.Encoding_JSON_IETF,
		}
		if diff := cmp.Diff(wantValue, got, cmp.AllowUnexported(telemetry.QualifiedString{}), protocmp.Transform()); diff != "" {
			t.Errorf("Got desc Qualified type different from expected (-want,+got):\n %s", diff)
		}
		if diff := cmp.Diff(wantGetRequest, fakeGNMI.GetRequests()[len(fakeGNMI.GetRequests())-1], protocmp.IgnoreFields(&gpb.GetRequest{}, "prefix"), protocmp.Transform()); diff != "" {
			t.Errorf("Got GetRequest different from expected (-want,+got):\n %s", diff)
		}
		if diff := cmp.Diff(wantValue.Val(t), got.Val(t)); diff != "" {
			t.Errorf("Got desc val different from expected (-want,+got):\n %s", diff)
		}
	})
}

func TestGetNonleaf(t *testing.T) {
	intfName := "Et1/2/3"
	intfPath := gnmiPath(t, "interfaces/interface[name=%s]", intfName)
	statusPath := gnmiPath(t, "interfaces/interface[name=%s]/state/oper-status", intfName)
	namePath := gnmiPath(t, "interfaces/interface[name=%s]/state/name", intfName)
	nameConfigPath := gnmiPath(t, "interfaces/interface[name=%s]/config/name", intfName)
	enabledPath := gnmiPath(t, "interfaces/interface[name=%s]/state/enabled", intfName)

	testsPass := []struct {
		desc        string
		stubFn      func(s *fakegnmi.Stubber)
		lookupFn    func(t testing.TB) *telemetry.QualifiedInterface
		wantValue   *telemetry.QualifiedInterface
		wantPathErr string
	}{{
		desc: "one notification",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: telemRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_UP,
		}),
	}, {
		desc: "one notification with shadow path value, which is ignored",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: namePath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: intfName}},
					}, {
						Path: nameConfigPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bogus"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: telemRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			Name: ygot.String(intfName),
		}),
	}, {
		desc: "one notification using config path API",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: nameConfigPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: intfName}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: configRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			Name: ygot.String(intfName),
		}),
	}, {
		desc: "one notification using config path API, with non-shadow and shadow path value, the latter of which is ignored",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: nameConfigPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: intfName}},
					}, {
						Path: namePath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bogus"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: configRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			Name: ygot.String(intfName),
		}),
	}, {
		desc: "one notification using config path API, with shadow path value, which is ignored",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: namePath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bogus"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: configRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{}),
	}, {
		desc: "one notification with interpolation and two data points",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}, {
						Path: enabledPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 200,
				}).
				Sync()
		},
		lookupFn: telemRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 200),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_UP,
			Enabled:    ygot.Bool(true),
		}),
	}, {
		desc: "one notification with a prefix",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: statusPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}},
					Prefix:    &gpb.Path{Origin: "openconfig", Elem: statusPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn: telemRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_DOWN,
		}),
	}, {
		desc: "multiple notifications",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}, {
						Path: gnmiPath(t, "interfaces/interface[name=%s]/ethernet/state/mac-address", intfName),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "00:de:ad:be:ef:22"}},
					}},
					Timestamp: 300,
				}).
				Notification(
					&gpb.Notification{
						Update: []*gpb.Update{{
							Path: statusPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
						}},
						Timestamp: 400,
					}).
				Sync()
		},
		lookupFn: telemRoot().Interface(intfName).Lookup,
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 400),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_UP,
			Ethernet: &telemetry.Interface_Ethernet{
				MacAddress: ygot.String("00:de:ad:be:ef:22"),
			},
		}),
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		lookupFn:  telemRoot().Interface(intfName).Lookup,
		wantValue: nil,
	}, {
		desc: "unmarshal fails because returned prefix and path don't match",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: statusPath.GetElem()[2:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}},
					Prefix:    &gpb.Path{Elem: statusPath.GetElem()[1:2]},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn:    telemRoot().Interface(intfName).Lookup,
		wantValue:   nil,
		wantPathErr: "does not match the query path",
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := tt.lookupFn(t)
			verifySubscriptionPathsSent(t, intfPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s\n", diff)
				if got != nil {
					t.Logf("ComplianceErrors:\n%v", got.GetComplianceErrors())
				}
			}
			// Test Val(t) API.
			if tt.wantValue.IsPresent() {
				if diff := cmp.Diff(tt.wantValue.Val(t), got.Val(t), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{})); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
			if got != nil {
				verifyTelemetryErrSubstr(t, got.ComplianceErrors, tt.wantPathErr, "")
			}
		})
	}

	rootPath := gnmiPath(t, "/")
	testsPassRoot := []struct {
		desc        string
		stubFn      func(s *fakegnmi.Stubber)
		lookupFn    func(t testing.TB) *telemetry.QualifiedDevice
		wantSubPath *gpb.Path
		wantValue   *telemetry.QualifiedDevice
	}{{
		desc: "one notification on root",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		lookupFn:    telemRoot().Lookup,
		wantSubPath: rootPath,
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      rootPath,
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intfName: {
					Name:       &intfName,
					OperStatus: telemetry.Interface_OperStatus_UP,
				},
			},
		}),
	}}

	for _, tt := range testsPassRoot {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := tt.lookupFn(t)
			verifySubscriptionPathsSent(t, tt.wantSubPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantValue.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s\nComplianceErrors:\n%v", diff, got.ComplianceErrors)
			}
			// Test Val(t) API.
			if tt.wantValue.IsPresent() {
				if diff := cmp.Diff(tt.wantValue.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}

	testsFail := []struct {
		desc        string
		stubFn      func(s *fakegnmi.Stubber)
		intfPathFn  func() *interfaces.InterfacePath
		wantValue   *telemetry.QualifiedInterface
		wantPathErr string
		wantTypeErr string
	}{{
		desc: "schema noncompliance: one notification that doesn't unmarshal because path doesn't match",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: gnmiPath(t, "interfaces/interface[name=%s]/bogus", intfName),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		intfPathFn: func() *interfaces.InterfacePath {
			return telemRoot().Interface(intfName)
		},
		wantValue:   nil,
		wantPathErr: "no match found",
	}, {
		desc: "schema noncompliance: one notification with interpolation and two data points, where the type doesn't match for one of them",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}, {
						Path: enabledPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 200,
				}).
				Sync()
		},
		intfPathFn: func() *interfaces.InterfacePath {
			return telemRoot().Interface(intfName)
		},
		wantValue: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 200),
				Path:      intfPath,
			}}).SetVal(&telemetry.Interface{
			Enabled: ygot.Bool(true),
		}),
		wantTypeErr: "failed to unmarshal",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := tt.intfPathFn().Lookup(t)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantValue.IsPresent() {
				if diff := cmp.Diff(tt.wantValue.Val(t), got.Val(t), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{})); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
			if got != nil {
				verifyTelemetryErrSubstr(t, got.ComplianceErrors, tt.wantPathErr, tt.wantTypeErr)
			}
		})
	}
}

func TestWildcardGet(t *testing.T) {
	octetsPath := func(t *testing.T, intfName string) *gpb.Path {
		return gnmiPath(t, "interfaces/interface[name=%s]/state/counters/in-octets", intfName)
	}
	intf1Path := octetsPath(t, "Ethernet3/1/1")
	intf2Path := octetsPath(t, "Ethernet3/2/2")
	wantSubPath := octetsPath(t, "*")
	startTime := time.Now()

	testsPass := []struct {
		desc       string
		stubFn     func(s *fakegnmi.Stubber)
		wantValues []*telemetry.QualifiedUint64
	}{{
		desc: "one value",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Sync()
		},
		wantValues: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: intf1Path,
				}}).SetVal(123),
		},
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		wantValues: nil,
	}, {
		desc: "multiple values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}}}).
				Sync()
		},
		wantValues: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      intf1Path,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      intf2Path,
				}}).SetVal(456),
		},
	}, {
		desc: "multiple values, but one of them has a path which does not match the query",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: &gpb.Path{Elem: intf2Path.Elem[1:]},
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}}}).
				Sync()
		},
		wantValues: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: intf1Path,
				}}).SetVal(123),
		},
	}, {
		desc: "schema noncompliance: one notification that doesn't unmarshal because type doesn't match",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}}}).
				Sync()
		},
		wantValues: nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := telemRoot().InterfaceAny().Counters().InOctets().Lookup(t)
			verifySubscriptionPathsSent(t, wantSubPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				if i < len(tt.wantValues) {
					tt.wantValues[i].RecvTimestamp = q.RecvTimestamp
				}
			}
			if diff := cmp.Diff(tt.wantValues, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got in octets different from expected, diff(-want,+got):\n %s", diff)
			}

			var wantVals, gotVals []uint64
			for _, wantValue := range tt.wantValues {
				wantVals = append(wantVals, wantValue.Val(t))
			}
			for _, inOctet := range got {
				gotVals = append(gotVals, inOctet.Val(t))
			}
			if diff := cmp.Diff(wantVals, gotVals); diff != "" {
				t.Errorf("Lookup(wildcard leaf): Got values different from expected (-want,+got):\n %s", diff)
			}
			inOctetVals := telemRoot().InterfaceAny().Counters().InOctets().Get(t)
			if diff := cmp.Diff(wantVals, inOctetVals); diff != "" {
				t.Errorf("Get(wildcard leaf): Got values different from expected (-want,+got):\n %s", diff)
			}
		})
	}

	t.Run("one invalid value -- Update's Val is empty", func(t *testing.T) {
		fakeGNMI.Stub().Notification(
			&gpb.Notification{
				Timestamp: startTime.UnixNano(),
				Update: []*gpb.Update{{
					Path: intf1Path,
					Val:  nil,
				}}}).
			Sync()
		got := testt.ExpectFatal(t, func(t testing.TB) {
			telemRoot().InterfaceAny().Counters().InOctets().Lookup(t)
		})
		if wantFatal := "invalid nil Val"; !strings.Contains(got, wantFatal) {
			t.Errorf("Wildcard Get failed with message %q, want %q", got, wantFatal)
		}
	})
}

func TestWildcardNonleafGet(t *testing.T) {
	wildcardPath := gnmiPath(t, "interfaces/interface[name=*]")
	intf1Name := "Et1/2/3"
	intf2Name := "Et4/5/6"
	intf1Path := gnmiPath(t, "interfaces/interface[name=%s]", intf1Name)
	intf2Path := gnmiPath(t, "interfaces/interface[name=%s]", intf2Name)
	status1Path := gnmiPath(t, "interfaces/interface[name=%s]/state/oper-status", intf1Name)
	status2Path := gnmiPath(t, "interfaces/interface[name=%s]/state/oper-status", intf2Name)

	testsPass := []struct {
		desc       string
		stubFn     func(s *fakegnmi.Stubber)
		wantValues []*telemetry.QualifiedInterface
	}{{
		desc: "one notification without interpolation",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: status2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantValues: []*telemetry.QualifiedInterface{
			(&telemetry.QualifiedInterface{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(0, 100),

					Path: intf2Path,
				}}).SetVal(&telemetry.Interface{
				OperStatus: telemetry.Interface_OperStatus_UP,
			}),
		},
	}, {
		desc: "two notifications under different containers, and a completely bogus path (which is ignored since its datapoint group is wholly noncompliant)",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: status1Path.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}, {
						Path: &gpb.Path{Elem: status2Path.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}, {
						Path: &gpb.Path{Elem: gnmiPath(t, "interfaces/interface[name=%s]/state/enabled", intf1Name).GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					// The first element, "interface", is shared.
					Prefix:    &gpb.Path{Origin: status2Path.GetOrigin(), Elem: status2Path.GetElem()[:1]},
					Timestamp: 200,
				}).
				Notification(
					&gpb.Notification{
						Update: []*gpb.Update{{
							Path: gnmiPath(t, "interfaces/interface[name=%s]/ethernet/state/mac-address", intf2Name),
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "00:de:ad:be:ef:22"}},
						}, {
							Path: gnmiPath(t, "bogus/interface[name=%s]/enabled", intf2Name),
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
						}},
						Timestamp: 400,
					},
				).
				Sync()
		},
		wantValues: []*telemetry.QualifiedInterface{
			(&telemetry.QualifiedInterface{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(0, 200),
					Path:      intf1Path,
				}}).SetVal(&telemetry.Interface{
				OperStatus: telemetry.Interface_OperStatus_UP,
				Enabled:    ygot.Bool(true),
			}),
			(&telemetry.QualifiedInterface{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(0, 400),
					Path:      intf2Path,
				}}).SetVal(&telemetry.Interface{
				OperStatus: telemetry.Interface_OperStatus_DOWN,
				Ethernet: &telemetry.Interface_Ethernet{
					MacAddress: ygot.String("00:de:ad:be:ef:22"),
				},
			}),
		},
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		wantValues: nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := telemRoot().InterfaceAny().Lookup(t)
			verifySubscriptionPathsSent(t, wildcardPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				if i < len(tt.wantValues) {
					tt.wantValues[i].RecvTimestamp = q.RecvTimestamp
				}
			}
			if diff := cmp.Diff(tt.wantValues, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Interface GoStructs different from expected, diff(-want,+got):\n %s", diff)
			}

			var wantVals, gotVals []*telemetry.Interface
			for _, wantValue := range tt.wantValues {
				wantVals = append(wantVals, wantValue.Val(t))
			}
			for _, qType := range got {
				gotVals = append(gotVals, qType.Val(t))
			}
			if diff := cmp.Diff(wantVals, gotVals); diff != "" {
				t.Errorf("Lookup(wildcard non-leaf): Got values different from expected (-want,+got):\n %s", diff)
			}

			interfaceVals := telemRoot().InterfaceAny().Get(t)
			if diff := cmp.Diff(wantVals, interfaceVals); diff != "" {
				t.Errorf("Get(wildcard non-leaf): Got values different from expected (-want,+got):\n %s", diff)
			}
		})
	}

	t.Run("schema noncompliance: one notification that doesn't unmarshal because path doesn't match", func(t *testing.T) {
		fakeGNMI.Stub().Notification(
			&gpb.Notification{
				Update: []*gpb.Update{{
					Path: gnmiPath(t, "interfaces/interface[name=%s]/bogus", intf2Name),
					Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
				}, {
					Path: status2Path,
					Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
				}},
				Timestamp: 100,
			}).
			Sync()
		got := telemRoot().InterfaceAny().Lookup(t)
		if len(got) != 1 {
			t.Fatalf("Expected exactly one piece of data for this test, but got %v: %v", len(got), got)
		}
		verifyTelemetryErrSubstr(t, got[0].GetComplianceErrors(), "no match found", "")
	})
}

func TestCollect(t *testing.T) {
	octetsPath := func(t *testing.T, intfName string) *gpb.Path {
		return gnmiPath(t, "interfaces/interface[name=%s]/state/counters/out-octets", intfName)
	}
	intf1Name := "Et1/2/3"
	intf2Name := "Et4/5/6"
	intf1Path := octetsPath(t, intf1Name)
	intf2Path := octetsPath(t, intf2Name)
	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounteres an EOF.
	testsPass := []struct {
		desc        string
		stubFn      func(s *fakegnmi.Stubber)
		intfName    string
		wantSubPath *gpb.Path
		wantValue   []*telemetry.QualifiedUint64
	}{{
		desc: "one value",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf2Name,
		wantSubPath: intf2Path,
		wantValue: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      intf2Path,
				}}).SetVal(123),
		},
	}, {
		desc: "one value with interpolation",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf1Name,
		wantSubPath: intf1Path,
		wantValue: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      intf1Path,
				}}).SetVal(123),
		},
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf2Name,
		wantSubPath: intf2Path,
		wantValue:   nil,
	}, {
		desc: "multiple values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf2Name,
		wantSubPath: intf2Path,
		wantValue: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      intf2Path,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      intf2Path,
				}}).SetVal(456),
		},
	}, {
		desc: "multiple values with the last one being type noncompliant",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: -456}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf2Name,
		wantSubPath: intf2Path,
		wantValue: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      intf2Path,
				}}).SetVal(123),
		},
	}, {
		desc: "single value that is path noncompliant",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intf2Path.Elem[len(intf2Path.Elem)-2:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
					}},
					Prefix: &gpb.Path{Elem: append([]*gpb.PathElem{{Name: "foo"}}, intf2Path.Elem[:len(intf2Path.Elem)-2]...)},
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf2Name,
		wantSubPath: intf2Path,
		wantValue:   nil,
	}, {
		desc: "multiple values with different prefixes",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: intf2Path.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}},
					Prefix: &gpb.Path{Origin: intf2Path.GetOrigin(), Elem: intf2Path.GetElem()[:1]},
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: &gpb.Path{Elem: intf2Path.GetElem()[3:]},
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}},
						Prefix: &gpb.Path{Origin: intf2Path.GetOrigin(), Elem: intf2Path.GetElem()[:3]},
					}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf2Name,
		wantSubPath: intf2Path,
		wantValue: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      intf2Path,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      intf2Path,
				}}).SetVal(456),
		},
	}, {
		desc: "with deletes",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}}}).
				Notification(
					&gpb.Notification{
						Prefix:    &gpb.Path{Elem: intf2Path.GetElem()[0:1]},
						Timestamp: startTime.Add(20 * time.Millisecond).UnixNano(),
						Delete: []*gpb.Path{
							{Origin: intf2Path.GetOrigin(), Elem: intf2Path.GetElem()[1:]},
						}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(30 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 789}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(40 * time.Millisecond).UnixNano(),
						Delete: []*gpb.Path{
							intf2Path,
						}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(50 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 111}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		intfName:    intf2Name,
		wantSubPath: intf2Path,
		wantValue: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      intf2Path,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      intf2Path,
				}}).SetVal(456),
			{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(20 * time.Millisecond),
					Path:      intf2Path,
				},
			},
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(30 * time.Millisecond),
					Path:      intf2Path,
				}}).SetVal(789),
			{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(40 * time.Millisecond),
					Path:      intf2Path,
				},
			},
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(50 * time.Millisecond),
					Path:      intf2Path,
				}}).SetVal(111),
		},
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got := telemRoot().Interface(tt.intfName).Counters().OutOctets().Collect(t, time.Second).Await(t)
			verifySubscriptionPathsSent(t, tt.wantSubPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				if i < len(tt.wantValue) {
					tt.wantValue[i].RecvTimestamp = q.RecvTimestamp
				}
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
		})
	}
}

func TestWatch(t *testing.T) {
	intfName := "Ethernet3/1/1"
	intfPath := gnmiPath(t, "interfaces/interface[name=%s]/state/counters/out-octets", intfName)
	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc       string
		stubFn     func(s *fakegnmi.Stubber)
		predicate  func(*telemetry.QualifiedUint64) bool
		wantValue  *telemetry.QualifiedUint64
		wantStatus bool
	}{{
		desc: "no values at path",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Sync()
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				})
		},
		predicate: func(val *telemetry.QualifiedUint64) bool {
			return !val.IsPresent()
		},
		wantStatus: true,
		wantValue: &telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Path: intfPath,
			},
		},
	}, {
		desc: "values and predicate never true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intfPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		predicate: func(val *telemetry.QualifiedUint64) bool {
			if val.IsPresent() {
				return val.Val(t) > 100
			}
			return false
		},
		wantValue: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      intfPath,
			},
		}).SetVal(80),
	}, {
		desc: "values and predicate becomes true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intfPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intfPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intfPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		predicate: func(val *telemetry.QualifiedUint64) bool {
			if val.IsPresent() {
				return val.Val(t) > 100
			}
			return false
		},
		wantStatus: true,
		wantValue: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      intfPath,
			},
		}).SetVal(120),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			got, predStatus := telemRoot().Interface(intfName).Counters().OutOctets().Watch(t, time.Second, tt.predicate).Await(t)
			verifySubscriptionPathsSent(t, intfPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
			}
		})
	}

	t.Run("multiple awaits", func(t *testing.T) {
		fakeGNMI.Stub().Sync().Notification(&gpb.Notification{Timestamp: startTime.Add(time.Second).UnixNano()})
		watcher := telemRoot().Interface("eth0").Counters().OutOctets().Watch(t, time.Second, func(val *telemetry.QualifiedUint64) bool { return true })
		got, gotStatus := watcher.Await(t)
		want := &telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Path: gnmiPath(t, "interfaces/interface[name=eth0]/state/counters/out-octets"),
			},
		}
		wantStatus := true

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
			t.Errorf("First call to Await() returned unexpected diff(-want,+got):\n %s", diff)
		}
		if gotStatus != wantStatus {
			t.Errorf("First call to Await() returned unexpected status got: %v want: %v", gotStatus, wantStatus)
		}

		got, gotStatus = watcher.Await(t)
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
			t.Errorf("Second call to Await() returned unexpected diff(-want,+got):\n %s", diff)
		}
		if gotStatus != wantStatus {
			t.Errorf("Second call to Await() returned unexpected status got: %v want: %v", gotStatus, wantStatus)
		}

	})
}

func TestWildcardWatch(t *testing.T) {
	octetsPath := func(t *testing.T, intfName string) *gpb.Path {
		return gnmiPath(t, "interfaces/interface[name=%s]/state/counters/out-octets", intfName)
	}
	ethPath := octetsPath(t, "eth0")
	loPath := octetsPath(t, "lo0")
	wildcardPath := octetsPath(t, "*")
	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc       string
		stubFn     func(s *fakegnmi.Stubber)
		wantValue  *telemetry.QualifiedUint64
		wantStatus bool
	}{{
		desc: "values and predicate never true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: ethPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: loPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		wantValue: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      loPath,
			},
		}).SetVal(50),
	}, {
		desc: "multiple values in notfication",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: ethPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
					}, {
						Path: loPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		wantValue: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime,
				Path:      loPath,
			},
		}).SetVal(50),
	}, {
		desc: "values and predicate becomes true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: ethPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: loPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		wantStatus: true,
		wantValue: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      loPath,
			},
		}).SetVal(150),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())

			var ethCond, loCond bool
			got, predStatus := telemRoot().InterfaceAny().Counters().OutOctets().Watch(t, time.Second, func(val *telemetry.QualifiedUint64) bool {
				if !ethCond {
					ethCond = val.IsPresent() && val.Path.String() == ethPath.String() && val.Val(t) < 100
				}
				if !loCond {
					loCond = val.IsPresent() && val.Path.String() == loPath.String() && val.Val(t) > 100
				}
				return ethCond && loCond
			}).Await(t)
			verifySubscriptionPathsSent(t, wildcardPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantValue.RecvTimestamp = got.RecvTimestamp

			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
			}
		})
	}
}

func TestNonleafWildcardWatch(t *testing.T) {
	countersPath := func(t *testing.T, iname, counter string) *gpb.Path {
		return gnmiPath(t, "interfaces/interface[name=%s]/state/counters/%s", iname, counter)
	}
	wildcardPath := gnmiPath(t, "interfaces/interface[name=*]/state/counters")
	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc       string
		stub       func(s *fakegnmi.Stubber)
		wantValue  *telemetry.QualifiedInterface_Counters
		wantStatus bool
	}{{
		desc: "predicate never true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: countersPath(t, "eth0", "in-octets"),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: countersPath(t, "eth0", "out-octets"),
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
						}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		wantValue: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      gnmiPath(t, "interfaces/interface[name=eth0]/state/counters"),
			},
		}).SetVal(&telemetry.Interface_Counters{
			InOctets:  ygot.Uint64(150),
			OutOctets: ygot.Uint64(50),
		}),
	}, {
		desc: "predicate is true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: countersPath(t, "lo0", "in-octets"),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: countersPath(t, "eth0", "out-octets"),
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
						}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		wantStatus: true,
		wantValue: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime,
				Path:      gnmiPath(t, "interfaces/interface[name=lo0]/state/counters"),
			},
		}).SetVal(&telemetry.Interface_Counters{InOctets: ygot.Uint64(150)}),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			var lo0, eth0 bool
			got, predStatus := telemRoot().InterfaceAny().Counters().Watch(t, time.Second, func(val *telemetry.QualifiedInterface_Counters) bool {
				lo0 = lo0 || (val.GetPath().GetElem()[1].GetKey()["name"] == "lo0" && val.Val(t).GetInOctets() == 150)
				eth0 = eth0 || (val.GetPath().GetElem()[1].GetKey()["name"] == "eth0" && val.Val(t).GetOutOctets() == 50)
				return lo0 && eth0
			}).Await(t)
			verifySubscriptionPathsSent(t, wildcardPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}

			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedInterface_Counters{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
			}
		})
	}
}

func TestBatchGet(t *testing.T) {
	octetsPath := func(t *testing.T, intfName string) *gpb.Path {
		return gnmiPath(t, "interfaces/interface[name=%s]/state/counters/out-octets", intfName)
	}
	intf1Name := "Ethernet3/1/1"
	intf2Name := "Ethernet3/1/2"
	intf1Path := octetsPath(t, intf1Name)
	intf2Path := octetsPath(t, intf2Name)
	wildcardPath := octetsPath(t, "*")

	tests := []struct {
		desc         string
		stubFn       func(s *fakegnmi.Stubber)
		addPaths     func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch)
		wantSubPaths []*gpb.Path
		wantValue    *telemetry.QualifiedDevice
	}{{
		desc: "two leaves in single notifcation",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}, {
						Path: gnmiPath(t, "interfaces/interface[name=%s]/config/enabled", intf1Name),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name: ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}, {
		desc: "leaves with prefix",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Prefix: &gpb.Path{
						Elem:   intf1Path.Elem[0:1],
						Origin: "openconfig",
					},
					Update: []*gpb.Update{{
						Path: &gpb.Path{
							Elem: intf1Path.Elem[1:],
						},
						Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: &gpb.Path{
							Elem: intf2Path.Elem[1:],
						},
						Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name: ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}, {
		desc: "non-subscribed paths",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf2Path},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Path: &gpb.Path{
					Origin: "openconfig",
				},
				Timestamp: time.Unix(0, 100),
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name: ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}, {
		desc: "two leaves with wildcard subcription",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.InterfaceAny().Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{wildcardPath},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name: ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			batch := telemRoot().NewBatch()
			tt.addPaths(t, telemRoot(), batch)
			got := batch.Lookup(t)
			verifySubscriptionPathsSent(t, tt.wantSubPaths...)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantValue.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantValue, got, cmpopts.EquateErrors(), cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Lookup() return unexpected device diff(-want,+got):\n %s", diff)
			}
		})
	}
}

func TestBatchWatch(t *testing.T) {
	octetsPath := func(t *testing.T, intfName string) *gpb.Path {
		return gnmiPath(t, "interfaces/interface[name=%s]/state/counters/out-octets", intfName)
	}
	intf1Name := "Ethernet3/1/1"
	intf2Name := "Ethernet3/1/2"
	intf1Path := octetsPath(t, intf1Name)
	intf2Path := octetsPath(t, intf2Name)
	wildcardPath := octetsPath(t, "*")
	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc         string
		stubFn       func(s *fakegnmi.Stubber)
		addPaths     func(testing.TB, *device.DevicePath, *telemetry.Batch)
		wantSubPaths []*gpb.Path
		wantPredArgs []*telemetry.QualifiedDevice
		wantValue    *telemetry.QualifiedDevice
		wantStatus   bool
	}{{
		desc: "multiple notifications before sync and predicate never true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 30}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf1Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
						}, {
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
						}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name: ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(50),
					},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(80),
					},
				},
			},
		}),
		wantPredArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(50),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(80),
						},
					},
				},
			}),
		},
	}, {
		desc: "collection with delete",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Delete:    []*gpb.Path{intf1Path},
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name:     ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(120),
					},
				},
			},
		}),
		wantPredArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
				},
			}),
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name:     ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "predicate becomes true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantStatus:   true,
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name: ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(75),
					},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(120),
					},
				},
			},
		}),
		wantPredArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(80),
						},
					},
				},
			}),
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "wildcard predicate becomes true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: intf2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(35 * time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.InterfaceAny().Counters().OutOctets().Batch(t, b)
		},
		wantStatus:   true,
		wantSubPaths: []*gpb.Path{wildcardPath},
		wantValue: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),

				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				intf1Name: {
					Name: ygot.String(intf1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(75),
					},
				},
				intf2Name: {
					Name: ygot.String(intf2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(120),
					},
				},
			},
		}),
		wantPredArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(80),
						},
					},
				},
			}),
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValue:    nil,
		wantPredArgs: []*telemetry.QualifiedDevice{},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			batch := telemRoot().NewBatch()
			tt.addPaths(t, telemRoot(), batch)
			i := 0

			got, predStatus := batch.Watch(t, 250*time.Millisecond, func(val *telemetry.QualifiedDevice) bool {
				checkJustReceived(t, val.RecvTimestamp)
				tt.wantPredArgs[i].RecvTimestamp = val.RecvTimestamp
				if diff := cmp.Diff(tt.wantPredArgs[i], val, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
					t.Errorf("Predicate received unexpected device diff (-want,+got):\n %s", diff)
				}
				i++
				return val.IsPresent() && val.Val(t).GetInterface(intf1Name).GetCounters().GetOutOctets() > 60 &&
					val.Val(t).GetInterface(intf2Name).GetCounters().GetOutOctets() > 100
			}).Await(t)
			if i != len(tt.wantPredArgs) {
				t.Errorf("Predicate didn't receive all qualified devices: got %d, want %d", i, len(tt.wantPredArgs))
			}
			verifySubscriptionPathsSent(t, tt.wantSubPaths...)
			if tt.wantValue != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Watch() returned unexpected device diff (-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Watch() returned unexpected status: got %v, want %v ", predStatus, tt.wantStatus)
			}
		})
	}
}

func TestBatchCollect(t *testing.T) {
	octetsPath := func(t *testing.T, intfName string) *gpb.Path {
		return gnmiPath(t, "interfaces/interface[name=%s]/state/counters/out-octets", intfName)
	}
	intf1Name := "Ethernet3/1/1"
	intf2Name := "Ethernet3/1/2"
	intf1Path := octetsPath(t, intf1Name)
	intf2Path := octetsPath(t, intf2Name)
	wildcardPath := octetsPath(t, "*")
	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc         string
		stubFn       func(s *fakegnmi.Stubber)
		addPaths     func(testing.TB, *device.DevicePath, *telemetry.Batch)
		wantSubPaths []*gpb.Path
		wantValues   []*telemetry.QualifiedDevice
	}{{
		desc: "collection with delete",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Delete:    []*gpb.Path{intf1Path},
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValues: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
				},
			}),
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),

					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name:     ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "multiple paths",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValues: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(80),
						},
					},
				},
			}),
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),

					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "multiple notifications before sync",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValues: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(80),
						},
					},
				},
			}),
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "wildcard subscription",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: intf2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.InterfaceAny().Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{wildcardPath},
		wantValues: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(80),
						},
					},
				},
			}),
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),

					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					intf1Name: {
						Name: ygot.String(intf1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					intf2Name: {
						Name: ygot.String(intf2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(intf1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(intf2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubPaths: []*gpb.Path{intf1Path, intf2Path},
		wantValues:   nil,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			batch := telemRoot().NewBatch()
			tt.addPaths(t, telemRoot(), batch)

			got := batch.Collect(t, 250*time.Millisecond).Await(t)
			for i, qt := range got {
				checkJustReceived(t, qt.RecvTimestamp)
				tt.wantValues[i].RecvTimestamp = qt.RecvTimestamp
			}

			verifySubscriptionPathsSent(t, tt.wantSubPaths...)
			if diff := cmp.Diff(tt.wantValues, got, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Collect() returned unexpected device diff (-want,+got):\n %s", diff)
			}
		})
	}
}

func TestNonleafWatch(t *testing.T) {
	intfName := "Ethernet3/1/1"
	counterStrPath := fmt.Sprintf("interfaces/interface[name=%s]/state/counters", intfName)
	nonLeafPath := gnmiPath(t, counterStrPath)
	inOctPath := gnmiPath(t, counterStrPath+"/in-octets")
	outOctPath := gnmiPath(t, counterStrPath+"/out-octets")
	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc         string
		stubFn       func(s *fakegnmi.Stubber)
		wantPredArgs []*telemetry.QualifiedInterface_Counters
		wantValue    *telemetry.QualifiedInterface_Counters
		wantStatus   bool
	}{{
		desc: "multiple notifications before sync and predicate never true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: inOctPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 30}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: inOctPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
					}, {
						Path: outOctPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		wantPredArgs: []*telemetry.QualifiedInterface_Counters{
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      nonLeafPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets:  ygot.Uint64(50),
				OutOctets: ygot.Uint64(80),
			}),
		},
		wantValue: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      nonLeafPath,
			},
		}).SetVal(&telemetry.Interface_Counters{
			InOctets:  ygot.Uint64(50),
			OutOctets: ygot.Uint64(80),
		}),
	}, {
		desc: "with delete",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: inOctPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Delete:    []*gpb.Path{inOctPath},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		wantPredArgs: []*telemetry.QualifiedInterface_Counters{
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      nonLeafPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets: ygot.Uint64(75),
			}),
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      nonLeafPath,
				},
			}).SetVal(&telemetry.Interface_Counters{}),
		},
		wantValue: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      nonLeafPath,
			},
		}).SetVal(&telemetry.Interface_Counters{}),
	}, {
		desc: "predicate becomes true",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: inOctPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: outOctPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: inOctPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				},
			)
		},
		wantStatus: true,
		wantPredArgs: []*telemetry.QualifiedInterface_Counters{
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      nonLeafPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets:  ygot.Uint64(75),
				OutOctets: ygot.Uint64(80),
			}),
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      nonLeafPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets:  ygot.Uint64(120),
				OutOctets: ygot.Uint64(80),
			}),
		},
		wantValue: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      nonLeafPath,
			},
		}).SetVal(&telemetry.Interface_Counters{
			InOctets:  ygot.Uint64(120),
			OutOctets: ygot.Uint64(80),
		}),
	}, {
		desc: "no values",
		stubFn: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					},
				)
		},
		wantValue:    nil,
		wantPredArgs: []*telemetry.QualifiedInterface_Counters{},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stubFn(fakeGNMI.Stub())
			i := 0

			got, predStatus := telemRoot().Interface(intfName).Counters().Watch(t, 250*time.Millisecond, func(val *telemetry.QualifiedInterface_Counters) bool {
				checkJustReceived(t, val.RecvTimestamp)
				tt.wantPredArgs[i].RecvTimestamp = val.RecvTimestamp
				if diff := cmp.Diff(tt.wantPredArgs[i], val, cmp.AllowUnexported(telemetry.QualifiedInterface_Counters{}), protocmp.Transform()); diff != "" {
					t.Errorf("Predicate received unexpected device diff (-want,+got):\n %s", diff)
				}
				i++
				return val.IsPresent() && val.Val(t).GetOutOctets() > 60 && val.Val(t).GetInOctets() > 100
			}).Await(t)
			if i != len(tt.wantPredArgs) {
				t.Errorf("Predicate didn't receive all qualified devices: got %d, want %d", i, len(tt.wantPredArgs))
			}
			verifySubscriptionPathsSent(t, nonLeafPath)
			if tt.wantValue != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantValue.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantValue, got, cmp.AllowUnexported(telemetry.QualifiedInterface_Counters{}), protocmp.Transform()); diff != "" {
				t.Errorf("Watch() returned unexpected device diff (-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Watch() returned unexpected status: got %v, want %v ", predStatus, tt.wantStatus)
			}
		})
	}
}

// checks that the received time is just before now
func checkJustReceived(t *testing.T, recvTime time.Time) {
	if diffSecs := time.Now().Sub(recvTime).Seconds(); diffSecs <= 0 && diffSecs > 1 {
		t.Errorf("received time is too far (%v seconds) away from now", diffSecs)
	}
}
