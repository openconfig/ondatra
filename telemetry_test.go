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

package ondatra

import (
	"golang.org/x/net/context"
	"fmt"
	"strings"
	"testing"
	"time"

	log "github.com/golang/glog"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/openconfig/ygot/testutil"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/ondatra/internal/fakegnmi"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/telemetry/device"
	"github.com/openconfig/ondatra/telemetry/interfaces"
	"github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/testt"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var fakeGNMI = func() *fakegnmi.FakeGNMI {
	fg, err := fakegnmi.Start(0)
	if err != nil {
		log.Fatalf("could not start fake gNMI server: %v", err)
	}
	return fg
}()

func initTelemetryFakes() {
	fakebind.Setup().WithReservation(fakeRes)
	for _, dut := range fakeRes.DUTs {
		dut.(*fakebind.DUT).DialGNMIFn = func(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
			return fakeGNMI.Dial(ctx, opts...)
		}
	}
}

func gnmiPath(t *testing.T, s string) *gpb.Path {
	t.Helper()
	p, err := ygot.StringToStructuredPath(s)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.GetElem()) == 0 || p.GetElem()[0].GetName() != "meta" {
		p.Origin = "openconfig"
	}
	return p
}

func gnmiDeprecatedPath(t *testing.T, s string) *gpb.Path {
	t.Helper()
	p, err := ygot.StringToStringSlicePath(s)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.GetElement()) == 0 || p.GetElement()[0] != "meta" {
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
func verifyTelemetryErrSubstr(t *testing.T, complianceErrs *genutil.ComplianceErrors, wantPathErrSubstr, wantTypeErrSubstr, wantValidateErrSubstr string) {
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
	var pathErr, typeErr, validateErr error
	if len(pathErrs) == 1 {
		pathErr = pathErrs[0].Err
	}
	if len(typeErrs) == 1 {
		typeErr = typeErrs[0].Err
	}
	if len(validateErrs) == 1 {
		validateErr = validateErrs[0]
	}

	// Validate expected errors
	if diff := errdiff.Substring(pathErr, wantPathErrSubstr); diff != "" {
		t.Fatalf("unmarshal: did not get expected path Err substring:\n%s", diff)
	}
	if diff := errdiff.Substring(typeErr, wantTypeErrSubstr); diff != "" {
		t.Fatalf("unmarshal: did not get expected type Err substring:\n%s", diff)
	}
	if diff := errdiff.Substring(validateErr, wantValidateErrSubstr); diff != "" {
		t.Fatalf("unmarshal: did not get expected validateErr substring:\n%s", diff)
	}
}

func TestMetadata(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_arista")
	connPath := gnmiPath(t, "meta/connected")
	syncPath := gnmiPath(t, "meta/sync")
	leavesAddedPath := gnmiPath(t, "meta/targetLeavesAdded")

	testsPass := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		getFunc              func(t testing.TB) genutil.QualifiedValue
		wantSubscriptionPath *gpb.Path
		wantQualified        genutil.QualifiedValue
	}{{
		desc: "check connected",
		stub: func(s *fakegnmi.Stubber) {
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
		getFunc: func(t testing.TB) genutil.QualifiedValue {
			return dut.Telemetry().Meta().Connected().Lookup(t)
		},
		wantSubscriptionPath: connPath,
		wantQualified: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      connPath,
			}}).SetVal(true),
	}, {
		desc: "check connected -- unconnected",
		stub: func(s *fakegnmi.Stubber) {
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
		getFunc: func(t testing.TB) genutil.QualifiedValue {
			return dut.Telemetry().Meta().Connected().Lookup(t)
		},
		wantSubscriptionPath: connPath,
		wantQualified: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      connPath,
			}}).SetVal(false),
	}, {
		desc: "check sync",
		stub: func(s *fakegnmi.Stubber) {
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
		getFunc: func(t testing.TB) genutil.QualifiedValue {
			return dut.Telemetry().Meta().Sync().Lookup(t)
		},
		wantSubscriptionPath: syncPath,
		wantQualified: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      syncPath,
			}}).SetVal(true),
	}, {
		desc: "check sync -- unsynced",
		stub: func(s *fakegnmi.Stubber) {
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
		getFunc: func(t testing.TB) genutil.QualifiedValue {
			return dut.Telemetry().Meta().Sync().Lookup(t)
		},
		wantSubscriptionPath: syncPath,
		wantQualified: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      syncPath,
			}}).SetVal(false),
	}, {
		desc: "check leaves added",
		stub: func(s *fakegnmi.Stubber) {
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
		getFunc: func(t testing.TB) genutil.QualifiedValue {
			return dut.Telemetry().Meta().TargetLeavesAdded().Lookup(t)
		},
		wantSubscriptionPath: leavesAddedPath,
		wantQualified: (&telemetry.QualifiedInt64{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 200),
				Path:      leavesAddedPath,
			}}).SetVal(42),
	}, {
		desc: "no values for leaves added",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(&gpb.Notification{}).Sync()
		},
		getFunc: func(t testing.TB) genutil.QualifiedValue {
			ret := dut.Telemetry().Meta().TargetLeavesAdded().Lookup(t)
			if ret != nil { // avoid typed-nil
				return ret
			}
			return nil
		},
		wantSubscriptionPath: leavesAddedPath,
		wantQualified:        nil,
	}, {
		desc: "no values for leaves added without connected return value",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		getFunc: func(t testing.TB) genutil.QualifiedValue {
			ret := dut.Telemetry().Meta().TargetLeavesAdded().Lookup(t)
			if ret != nil { // avoid typed-nil
				return ret
			}
			return nil
		},
		wantSubscriptionPath: leavesAddedPath,
		wantQualified:        nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := tt.getFunc(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if got != nil {
				checkJustReceived(t, got.GetRecvTimestamp())
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmpopts.IgnoreFields(telemetry.QualifiedBool{}, "RecvTimestamp"), cmpopts.IgnoreFields(telemetry.QualifiedInt64{}, "RecvTimestamp"), cmp.AllowUnexported(telemetry.QualifiedBool{}, telemetry.QualifiedInt64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_cisco")
	port := dut.Port(t, "port1")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/oper-status", iname)
	}

	genericPortName := port.Name()
	resolvedPortName := "Et1/2/3"
	resolvedPath := gnmiPath(t, getStrPath(resolvedPortName))

	staticPortName := "Ethernet3/1/1"
	statusPath := gnmiPath(t, getStrPath(staticPortName))

	testsPass := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inPortKey            string
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedE_Interface_OperStatus
		wantValFatal         string
	}{{
		desc: "one notification",
		stub: func(s *fakegnmi.Stubber) {
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
		inPortKey:            staticPortName,
		wantSubscriptionPath: statusPath,
		wantQualified: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      statusPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "no sync response",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: statusPath,
		wantQualified: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      statusPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "one notification with interpolation",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: resolvedPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inPortKey:            genericPortName,
		wantSubscriptionPath: resolvedPath,
		wantQualified: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      resolvedPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "one notification with a prefix",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: statusPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}},
					Prefix:    &gpb.Path{Origin: statusPath.GetOrigin(), Elem: statusPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: statusPath,
		wantQualified: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      statusPath,
			}}).SetVal(telemetry.Interface_OperStatus_DOWN),
	}, {
		desc: "one notification with a prefix with interpolation",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: resolvedPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}},
					Prefix:    &gpb.Path{Origin: resolvedPath.GetOrigin(), Elem: resolvedPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		inPortKey:            genericPortName,
		wantSubscriptionPath: resolvedPath,
		wantQualified: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      resolvedPath,
			}}).SetVal(telemetry.Interface_OperStatus_DOWN),
	}, {
		desc: "multiple notifications with the first one having no update value",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(&gpb.Notification{Update: []*gpb.Update{}}).
				Notification(&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 101,
				}).
				Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: statusPath,
		wantQualified: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 101),
				Path:      statusPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: statusPath,
		wantQualified:        nil,
		wantValFatal:         "No value present\n",
	}, {
		desc: "with connection",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 102,
				}).
				Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: statusPath,
		wantQualified: (&telemetry.QualifiedE_Interface_OperStatus{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 102),
				Path:      statusPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Telemetry().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(tt.inPortKey).OperStatus().Lookup(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedE_Interface_OperStatus{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
			fatalMsg := testt.CaptureFatal(t, func(t testing.TB) {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
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
		stub         func(s *fakegnmi.Stubber)
		wantFatalMsg string
	}{{
		desc: "multiple values",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}}}).
				Notification(
					&gpb.Notification{
						Update: []*gpb.Update{{
							Path: statusPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
						}}}).
				Sync()
		},
		wantFatalMsg: "got multiple",
	}, {
		desc: "returned path uses deprecated Element field",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: gnmiDeprecatedPath(t, getStrPath(staticPortName)),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
				}).
				Sync()
		},
		wantFatalMsg: "deprecated and unsupported Element field",
	}, {
		desc: "last element is different than the query path",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]/state/description", staticPortName)),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
				}).
				Sync()
		},
		wantFatalMsg: "does not match the query path",
	}, {
		desc: "returned prefix + path is incompatible",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: statusPath.GetElem()[len(statusPath.GetElem())-1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					// missing first element.
					Prefix: &gpb.Path{Elem: statusPath.GetElem()[1 : len(statusPath.GetElem())-1]},
				}).
				Sync()
		},
		wantFatalMsg: "does not match the query path",
	}, {
		desc: "returned path is incompatible",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: append(statusPath.GetElem(), &gpb.PathElem{Name: "does-not-exist"})},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
				}).
				Sync()
		},
		// The path cannot be unmarshalled because it doesn't exist in the schema.
		wantFatalMsg: "does-not-exist",
	}, {
		desc: "one invalid notification where Val of Update is nil",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  nil,
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantFatalMsg: "invalid nil Val",
	}, {
		desc: "schema noncompliance: one notification that doesn't unmarshal because type doesn't match",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
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
			tt.stub(fakeGNMI.Stub())
			got := testt.ExpectFatal(t, func(t testing.TB) {
				dut.Telemetry().Interface(staticPortName).OperStatus().Lookup(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Lookup failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestGetDefault(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_arista")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/enabled", iname)
	}
	getStrConfigPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/config/enabled", iname)
	}

	staticPortName := "Ethernet3/1/1"
	enabledPath := gnmiPath(t, getStrPath(staticPortName))
	enabledConfigPath := gnmiPath(t, getStrConfigPath(staticPortName))

	testsPass := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inGetCall            func(t testing.TB) *telemetry.QualifiedBool
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedBool
	}{{
		desc: "one notification",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: enabledPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: false}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inGetCall:            dut.Telemetry().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(staticPortName).Enabled().Lookup,
		wantSubscriptionPath: enabledPath,
		wantQualified: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      enabledPath,
			}}).SetVal(false),
	}, {
		desc: "telemetry path API: zero notifications on a leaf that has a default value",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inGetCall:            dut.Telemetry().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(staticPortName).Enabled().Lookup,
		wantSubscriptionPath: enabledPath,
		wantQualified: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Path: enabledPath,
			}}).SetVal(true),
	}, {
		desc: "config path API: zero notifications on a leaf that has a default value",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inGetCall:            dut.Config().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(staticPortName).Enabled().Lookup,
		wantSubscriptionPath: enabledConfigPath,
		wantQualified: (&telemetry.QualifiedBool{
			Metadata: &genutil.Metadata{
				Config: true,
				Path:   enabledConfigPath,
			}}).SetVal(true),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := tt.inGetCall(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedBool{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantQualified.IsPresent() {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_arista")
	port := dut.Port(t, "port1")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/config/description", iname)
	}

	genericPortName := port.Name()
	resolvedPortName := "Et1/2/3"
	resolvedPath := gnmiPath(t, getStrPath(resolvedPortName))

	staticPortName := "Ethernet3/1/1"
	descPath := gnmiPath(t, getStrPath(staticPortName))

	testsPass := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inPortKey            string
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedString
	}{{
		desc: "one notification",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: descPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: descPath,
		wantQualified: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      descPath,
			}}).SetVal("foo"),
	}, {
		desc: "one notification with interpolation",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: resolvedPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inPortKey:            genericPortName,
		wantSubscriptionPath: resolvedPath,
		wantQualified: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),

				Path: resolvedPath,
			}}).SetVal("foo"),
	}, {
		desc: "one notification with a prefix",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: descPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bar"}},
					}},
					Prefix:    &gpb.Path{Origin: descPath.GetOrigin(), Elem: descPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: descPath,
		wantQualified: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      descPath,
			}}).SetVal("bar"),
	}, {
		desc: "one notification with a prefix with interpolation",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: resolvedPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bar"}},
					}},
					Prefix:    &gpb.Path{Origin: resolvedPath.GetOrigin(), Elem: resolvedPath.GetElem()[:1]},
					Timestamp: 100,
				}).
				Sync()
		},
		inPortKey:            genericPortName,
		wantSubscriptionPath: resolvedPath,
		wantQualified: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      resolvedPath,
			}}).SetVal("bar"),
	}, {
		desc: "multiple notifications with the first one having no update value",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(&gpb.Notification{}).
				Notification(&gpb.Notification{
					Update: []*gpb.Update{{
						Path: descPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					Timestamp: 101,
				}).
				Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: descPath,
		wantQualified: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 101),
				Path:      descPath,
			}}).SetVal("foo"),
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: descPath,
		wantQualified:        nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Config().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(tt.inPortKey).Description().Lookup(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedString{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got desc Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantQualified.IsPresent() {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got desc val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}

	testsFail := []struct {
		desc         string
		stub         func(s *fakegnmi.Stubber)
		wantFatalMsg string
	}{{
		desc: "multiple values",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: descPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
				}).
				Notification(
					&gpb.Notification{
						Update: []*gpb.Update{{
							Path: descPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bar"}},
						}},
					}).
				Sync()
		},
		wantFatalMsg: "got multiple",
	}, {
		desc: "one invalid notification where Val of Update is nil",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: descPath,
						Val:  nil,
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantFatalMsg: "invalid nil Val",
	}, {
		desc: "schema noncompliance: one notification that doesn't unmarshal because type doesn't match",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: descPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantFatalMsg: "failed to unmarshal",
	}, {
		desc: "unmarshal fails because returned prefix and path don't match",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: descPath.GetElem()[len(descPath.GetElem())-1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
					}},
					// missing first element.
					Prefix: &gpb.Path{Elem: descPath.GetElem()[1 : len(descPath.GetElem())-1]},
				}).
				Sync()
		},
		wantFatalMsg: "does not match the query path",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := testt.ExpectFatal(t, func(t testing.TB) {
				dut.Config().Interface(staticPortName).Description().Lookup(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Lookup failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}

	gnmiGetTest := []struct {
		desc           string
		stub           func(s *fakegnmi.Stubber)
		inDutName      string
		wantGetRequest *gpb.GetRequest
		wantQualified  *telemetry.QualifiedString
	}{{
		desc:      "cisco",
		inDutName: "dut_cisco",
		stub: func(s *fakegnmi.Stubber) {
			s.GetResponse(&gpb.GetResponse{
				Notification: []*gpb.Notification{{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: descPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{JsonIetfVal: []byte(`"foo"`)}},
					}},
				}},
			})
		},
		wantGetRequest: &gpb.GetRequest{
			Path:     []*gpb.Path{descPath},
			Type:     gpb.GetRequest_CONFIG,
			Encoding: gpb.Encoding_JSON_IETF,
		},
		wantQualified: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      descPath,
			}}).SetVal("foo"),
	}, {
		desc:      "juniper",
		inDutName: "dut_juniper",
		stub: func(s *fakegnmi.Stubber) {
			s.GetResponse(&gpb.GetResponse{
				Notification: []*gpb.Notification{{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: descPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_JsonIetfVal{JsonIetfVal: []byte(`"foo"`)}},
					}},
				}},
			})
		},
		wantGetRequest: &gpb.GetRequest{
			Path:     []*gpb.Path{descPath},
			Type:     gpb.GetRequest_CONFIG,
			Encoding: gpb.Encoding_JSON_IETF,
		},
		wantQualified: (&telemetry.QualifiedString{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      descPath,
			}}).SetVal("foo"),
	}}
	for _, tt := range gnmiGetTest {
		t.Run(tt.desc, func(t *testing.T) {
			dut := DUT(t, tt.inDutName)
			tt.stub(fakeGNMI.Stub())

			got := dut.Config().Interface(staticPortName).Description().Lookup(t)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp

			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedString{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got desc Qualified type different from expected (-want,+got):\n %s", diff)
			}
			if diff := cmp.Diff(tt.wantGetRequest, fakeGNMI.GetRequests()[len(fakeGNMI.GetRequests())-1], protocmp.IgnoreFields(&gpb.GetRequest{}, "prefix"), protocmp.Transform()); diff != "" {
				t.Errorf("Got GetRequest different from expected (-want,+got):\n %s", diff)
			}
			if tt.wantQualified.IsPresent() {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got desc val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}
}

func TestGetNonleaf(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_arista")
	port := dut.Port(t, "port1")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/oper-status", iname)
	}

	getStrName := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/name", iname)
	}
	getStrNameConfig := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/config/name", iname)
	}

	// Port paths
	genericPortName := port.Name()
	resolvedPortName := "Et1/2/3"
	resolvedPath := gnmiPath(t, getStrPath(resolvedPortName))

	staticPortName := "Ethernet3/1/1"
	statusPath := gnmiPath(t, getStrPath(staticPortName))

	// Port names
	namePath := gnmiPath(t, getStrName(staticPortName))
	nameConfigPath := gnmiPath(t, getStrNameConfig(staticPortName))

	// Interface list path
	staticIntfPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]", staticPortName))
	resolvedIntfPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]", resolvedPortName))

	// Paths of some fields under the interface list
	resolvedEnabledPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]/state/enabled", resolvedPortName))
	macAddrPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]/ethernet/state/mac-address", staticPortName))

	bogusPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]/bogus", staticPortName))

	testsPass := []struct {
		desc                  string
		stub                  func(s *fakegnmi.Stubber)
		inGetCall             func(t testing.TB) *telemetry.QualifiedInterface
		wantSubscriptionPath  *gpb.Path
		wantQualified         *telemetry.QualifiedInterface
		wantPathErrSubstr     string
		wantTypeErrSubstr     string
		wantValidateErrSubstr string
	}{{
		desc: "one notification",
		stub: func(s *fakegnmi.Stubber) {
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),

				Path: staticIntfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_UP,
		}),
	}, {
		desc: "one notification with shadow path value, which is ignored",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: namePath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: staticPortName}},
					}, {
						Path: nameConfigPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bogus"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inGetCall:            dut.Telemetry().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      staticIntfPath,
			}}).SetVal(&telemetry.Interface{
			Name: ygot.String(staticPortName),
		}),
	}, {
		desc: "one notification using config path API",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: nameConfigPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: staticPortName}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inGetCall:            dut.Config().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      staticIntfPath,
			}}).SetVal(&telemetry.Interface{
			Name: ygot.String(staticPortName),
		}),
	}, {
		desc: "one notification using config path API, with non-shadow and shadow path value, the latter of which is ignored",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: nameConfigPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: staticPortName}},
					}, {
						Path: namePath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "bogus"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inGetCall:            dut.Config().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      staticIntfPath,
			}}).SetVal(&telemetry.Interface{
			Name: ygot.String(staticPortName),
		}),
	}, {
		desc: "one notification using config path API, with shadow path value, which is ignored",
		stub: func(s *fakegnmi.Stubber) {
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
		inGetCall:            dut.Config().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Config:    true,
				Timestamp: time.Unix(0, 100),
				Path:      staticIntfPath,
			}}).SetVal(&telemetry.Interface{}),
	}, {
		desc: "one notification with interpolation and two data points",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: resolvedPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}, {
						Path: resolvedEnabledPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 200,
				}).
				Sync()
		},
		inGetCall:            dut.Telemetry().Interface(genericPortName).Lookup,
		wantSubscriptionPath: resolvedIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 200),
				Path:      resolvedIntfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_UP,
			Enabled:    ygot.Bool(true),
		}),
	}, {
		desc: "one notification with a prefix",
		stub: func(s *fakegnmi.Stubber) {
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      staticIntfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_DOWN,
		}),
	}, {
		desc: "multiple notifications",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}, {
						Path: macAddrPath,
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 400),
				Path:      staticIntfPath,
			}}).SetVal(&telemetry.Interface{
			OperStatus: telemetry.Interface_OperStatus_UP,
			Ethernet: &telemetry.Interface_Ethernet{
				MacAddress: ygot.String("00:de:ad:be:ef:22"),
			},
		}),
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inGetCall:            dut.Telemetry().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified:        nil,
	}, {
		desc: "unmarshal fails because returned prefix and path don't match",
		stub: func(s *fakegnmi.Stubber) {
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).Lookup,
		wantSubscriptionPath: staticIntfPath,
		wantQualified:        nil,
		wantPathErrSubstr:    "does not match the query path",
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := tt.inGetCall(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s\n", diff)
				if got != nil {
					t.Logf("ComplianceErrors:\n%v", got.GetComplianceErrors())
				}
			}
			// Test Val(t) API.
			if tt.wantQualified.IsPresent() {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{})); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
			if got != nil {
				verifyTelemetryErrSubstr(t, got.ComplianceErrors, tt.wantPathErrSubstr, tt.wantTypeErrSubstr, tt.wantValidateErrSubstr)
			}
		})
	}

	rootPath := gnmiPath(t, "/")
	testsPassRoot := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inGetCall            func(t testing.TB) *telemetry.QualifiedDevice
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedDevice
	}{{
		desc: "one notification on root",
		stub: func(s *fakegnmi.Stubber) {
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
		inGetCall:            dut.Telemetry().Lookup,
		wantSubscriptionPath: rootPath,
		wantQualified: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path:      rootPath,
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				staticPortName: &telemetry.Interface{
					Name:       &staticPortName,
					OperStatus: telemetry.Interface_OperStatus_UP,
				},
			},
		}),
	}}

	for _, tt := range testsPassRoot {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := tt.inGetCall(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s\nComplianceErrors:\n%v", diff, got.ComplianceErrors)
			}
			// Test Val(t) API.
			if tt.wantQualified.IsPresent() {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}

	testsFail := []struct {
		desc                  string
		stub                  func(s *fakegnmi.Stubber)
		inInterfacePath       func() *interfaces.InterfacePath
		wantQualified         *telemetry.QualifiedInterface
		wantPathErrSubstr     string
		wantTypeErrSubstr     string
		wantValidateErrSubstr string
		wantFatalMsg          string
	}{{
		desc: "schema noncompliance: one notification that doesn't unmarshal because path doesn't match",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: bogusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		inInterfacePath: func() *interfaces.InterfacePath {
			return dut.Telemetry().Interface(staticPortName)
		},
		wantQualified:     nil,
		wantPathErrSubstr: "no match found",
	}, {
		desc: "schema noncompliance: one notification with interpolation and two data points, where the type doesn't match for one of them",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: resolvedPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}, {
						Path: resolvedEnabledPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					Timestamp: 200,
				}).
				Sync()
		},
		inInterfacePath: func() *interfaces.InterfacePath {
			return dut.Telemetry().Interface(genericPortName)
		},
		wantQualified: (&telemetry.QualifiedInterface{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 200),
				Path:      resolvedIntfPath,
			}}).SetVal(&telemetry.Interface{
			Enabled: ygot.Bool(true),
		}),
		wantTypeErrSubstr: "failed to unmarshal",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			if tt.wantFatalMsg != "" {
				got := testt.ExpectFatal(t, func(t testing.TB) {
					tt.inInterfacePath().Lookup(t)
				})
				if !strings.Contains(got, tt.wantFatalMsg) {
					t.Errorf("Lookup failed with message %q, want %q", got, tt.wantFatalMsg)
				}
				return
			}

			got := tt.inInterfacePath().Lookup(t)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantQualified.IsPresent() {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{})); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
			if got != nil {
				verifyTelemetryErrSubstr(t, got.ComplianceErrors, tt.wantPathErrSubstr, tt.wantTypeErrSubstr, tt.wantValidateErrSubstr)
			}
		})
	}
}

func TestWildcardGet(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_juniper")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/in-octets", iname)
	}

	intf1Path := gnmiPath(t, getStrPath("Ethernet3/1/1"))
	intf2Path := gnmiPath(t, getStrPath("Ethernet3/2/2"))
	wantSubscriptionPath := gnmiPath(t, getStrPath("*"))

	startTime := time.Now()

	testsPass := []struct {
		desc          string
		stub          func(s *fakegnmi.Stubber)
		wantQualified []*telemetry.QualifiedUint64
	}{{
		desc: "one value",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Sync()
		},
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: intf1Path,
				}}).SetVal(123),
		},
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		wantQualified: nil,
	}, {
		desc: "multiple values",
		stub: func(s *fakegnmi.Stubber) {
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
		wantQualified: []*telemetry.QualifiedUint64{
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
		stub: func(s *fakegnmi.Stubber) {
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
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: intf1Path,
				}}).SetVal(123),
		},
	}, {
		desc: "schema noncompliance: one notification that doesn't unmarshal because type doesn't match",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}}}).
				Sync()
		},
		wantQualified: nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Telemetry().InterfaceAny().Counters().InOctets().Lookup(t)
			verifySubscriptionPathsSent(t, wantSubscriptionPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				if i < len(tt.wantQualified) {
					tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
				}
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got in octets different from expected, diff(-want,+got):\n %s", diff)
			}

			var wantVals, gotVals []uint64
			for _, wantQualified := range tt.wantQualified {
				wantVals = append(wantVals, wantQualified.Val(t))
			}
			for _, inOctet := range got {
				gotVals = append(gotVals, inOctet.Val(t))
			}
			if diff := cmp.Diff(wantVals, gotVals); diff != "" {
				t.Errorf("Lookup(wildcard leaf): Got values different from expected (-want,+got):\n %s", diff)
			}

			inOctetVals := dut.Telemetry().InterfaceAny().Counters().InOctets().Get(t)
			if diff := cmp.Diff(wantVals, inOctetVals); diff != "" {
				t.Errorf("Get(wildcard leaf): Got values different from expected (-want,+got):\n %s", diff)
			}
		})
	}

	testsFail := []struct {
		desc         string
		stub         func(s *fakegnmi.Stubber)
		wantFatalMsg string
	}{{
		desc: "one invalid value -- Update's Val is empty",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: intf1Path,
						Val:  nil,
					}}}).
				Sync()
		},
		wantFatalMsg: "invalid nil Val",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := testt.ExpectFatal(t, func(t testing.TB) {
				dut.Telemetry().InterfaceAny().Counters().InOctets().Lookup(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Wildcard Get failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestWildcardNonleafGet(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_arista")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/oper-status", iname)
	}

	wantSubscriptionPath := gnmiPath(t, "interfaces/interface[name=*]")

	// Port paths
	resolvedPortName := "Et1/2/3"
	resolvedPath := gnmiPath(t, getStrPath(resolvedPortName))

	staticPortName := "Ethernet3/1/1"
	statusPath := gnmiPath(t, getStrPath(staticPortName))

	// Interface list path
	staticIntfPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]", staticPortName))
	resolvedIntfPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]", resolvedPortName))

	// Paths of some fields under the interface list
	resolvedEnabledPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]/state/enabled", resolvedPortName))
	macAddrPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]/ethernet/state/mac-address", staticPortName))

	bogusFieldPath := gnmiPath(t, fmt.Sprintf("interfaces/interface[name=%s]/bogus", staticPortName))
	bogusContainerPath := gnmiPath(t, fmt.Sprintf("bogus/interface[name=%s]/enabled", staticPortName))

	testsPass := []struct {
		desc          string
		stub          func(s *fakegnmi.Stubber)
		wantQualified []*telemetry.QualifiedInterface
	}{{
		desc: "one notification without interpolation",
		stub: func(s *fakegnmi.Stubber) {
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
		wantQualified: []*telemetry.QualifiedInterface{
			(&telemetry.QualifiedInterface{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(0, 100),

					Path: staticIntfPath,
				}}).SetVal(&telemetry.Interface{
				OperStatus: telemetry.Interface_OperStatus_UP,
			}),
		},
	}, {
		desc: "two notifications under different containers, and a completely bogus path (which is ignored since its datapoint group is wholly noncompliant)",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: resolvedPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}, {
						Path: &gpb.Path{Elem: statusPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "DOWN"}},
					}, {
						Path: &gpb.Path{Elem: resolvedEnabledPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}},
					// The first element, "interface", is shared.
					Prefix:    &gpb.Path{Origin: statusPath.GetOrigin(), Elem: statusPath.GetElem()[:1]},
					Timestamp: 200,
				}).
				Notification(
					&gpb.Notification{
						Update: []*gpb.Update{{
							Path: macAddrPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "00:de:ad:be:ef:22"}},
						}, {
							Path: bogusContainerPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
						}},
						Timestamp: 400,
					},
				).
				Sync()
		},
		wantQualified: []*telemetry.QualifiedInterface{
			(&telemetry.QualifiedInterface{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(0, 200),
					Path:      resolvedIntfPath,
				}}).SetVal(&telemetry.Interface{
				OperStatus: telemetry.Interface_OperStatus_UP,
				Enabled:    ygot.Bool(true),
			}),
			(&telemetry.QualifiedInterface{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(0, 400),
					Path:      staticIntfPath,
				}}).SetVal(&telemetry.Interface{
				OperStatus: telemetry.Interface_OperStatus_DOWN,
				Ethernet: &telemetry.Interface_Ethernet{
					MacAddress: ygot.String("00:de:ad:be:ef:22"),
				},
			}),
		},
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		wantQualified: nil,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Telemetry().InterfaceAny().Lookup(t)
			verifySubscriptionPathsSent(t, wantSubscriptionPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				if i < len(tt.wantQualified) {
					tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
				}
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Interface GoStructs different from expected, diff(-want,+got):\n %s", diff)
			}

			var wantVals, gotVals []*telemetry.Interface
			for _, wantQualified := range tt.wantQualified {
				wantVals = append(wantVals, wantQualified.Val(t))
			}
			for _, qType := range got {
				gotVals = append(gotVals, qType.Val(t))
			}
			if diff := cmp.Diff(wantVals, gotVals); diff != "" {
				t.Errorf("Lookup(wildcard non-leaf): Got values different from expected (-want,+got):\n %s", diff)
			}

			interfaceVals := dut.Telemetry().InterfaceAny().Get(t)
			if diff := cmp.Diff(wantVals, interfaceVals); diff != "" {
				t.Errorf("Get(wildcard non-leaf): Got values different from expected (-want,+got):\n %s", diff)
			}
		})
	}

	testsFail := []struct {
		desc                  string
		stub                  func(s *fakegnmi.Stubber)
		wantPathErrSubstr     string
		wantTypeErrSubstr     string
		wantValidateErrSubstr string
		wantFatalMsg          string
	}{{
		desc: "schema noncompliance: one notification that doesn't unmarshal because path doesn't match",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Update: []*gpb.Update{{
						Path: bogusFieldPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}, {
						Path: statusPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "UP"}},
					}},
					Timestamp: 100,
				}).
				Sync()
		},
		wantPathErrSubstr: "no match found",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			if tt.wantFatalMsg != "" {
				got := testt.ExpectFatal(t, func(t testing.TB) {
					dut.Telemetry().InterfaceAny().Lookup(t)
				})
				if !strings.Contains(got, tt.wantFatalMsg) {
					t.Errorf("Lookup failed with message %q, want %q", got, tt.wantFatalMsg)
				}
				return
			}
			got := dut.Telemetry().InterfaceAny().Lookup(t)
			if len(got) != 1 {
				t.Fatalf("Expected exactly one piece of data for this test, but got %v: %v", len(got), got)
			}

			verifyTelemetryErrSubstr(t, got[0].GetComplianceErrors(), tt.wantPathErrSubstr, tt.wantTypeErrSubstr, tt.wantValidateErrSubstr)
		})
	}
}

func TestCollect(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_cisco")
	port := dut.Port(t, "port1")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/out-octets", iname)
	}

	genericPortName := port.Name()
	resolvedPortName := "Et1/2/3"
	resolvedOctetsPath := gnmiPath(t, getStrPath(resolvedPortName))

	staticPortName := "Ethernet3/1/1"
	staticOctetsPath := gnmiPath(t, getStrPath(staticPortName))

	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounteres an EOF.
	testsPass := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inPortKey            string
		wantSubscriptionPath *gpb.Path
		wantQualified        []*telemetry.QualifiedUint64
	}{{
		desc: "one value",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: staticOctetsPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      staticOctetsPath,
				}}).SetVal(123),
		},
	}, {
		desc: "one value with interpolation",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: resolvedOctetsPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            genericPortName,
		wantSubscriptionPath: resolvedOctetsPath,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      resolvedOctetsPath,
				}}).SetVal(123),
		},
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified:        nil,
	}, {
		desc: "multiple values",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: staticOctetsPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      staticOctetsPath,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      staticOctetsPath,
				}}).SetVal(456),
		},
	}, {
		desc: "multiple values with the last one being type noncompliant",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: staticOctetsPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: -456}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      staticOctetsPath,
				}}).SetVal(123),
		},
	}, {
		desc: "single value that is path noncompliant",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: staticOctetsPath.Elem[len(staticOctetsPath.Elem)-2:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
					}},
					Prefix: &gpb.Path{Elem: append([]*gpb.PathElem{{Name: "foo"}}, staticOctetsPath.Elem[:len(staticOctetsPath.Elem)-2]...)},
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified:        nil,
	}, {
		desc: "multiple values with different prefixes",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: staticOctetsPath.GetElem()[1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}},
					Prefix: &gpb.Path{Origin: staticOctetsPath.GetOrigin(), Elem: staticOctetsPath.GetElem()[:1]},
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: &gpb.Path{Elem: staticOctetsPath.GetElem()[3:]},
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}},
						Prefix: &gpb.Path{Origin: staticOctetsPath.GetOrigin(), Elem: staticOctetsPath.GetElem()[:3]},
					}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      staticOctetsPath,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      staticOctetsPath,
				}}).SetVal(456),
		},
	}, {
		desc: "with deletes",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: staticOctetsPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 456}},
						}}}).
				Notification(
					&gpb.Notification{
						Prefix:    &gpb.Path{Elem: staticOctetsPath.GetElem()[0:1]},
						Timestamp: startTime.Add(20 * time.Millisecond).UnixNano(),
						Delete: []*gpb.Path{
							&gpb.Path{Origin: staticOctetsPath.GetOrigin(), Elem: staticOctetsPath.GetElem()[1:]},
						}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(30 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 789}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(40 * time.Millisecond).UnixNano(),
						Delete: []*gpb.Path{
							staticOctetsPath,
						}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(50 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 111}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      staticOctetsPath,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      staticOctetsPath,
				}}).SetVal(456),
			{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(20 * time.Millisecond),
					Path:      staticOctetsPath,
				},
			},
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(30 * time.Millisecond),
					Path:      staticOctetsPath,
				}}).SetVal(789),
			{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(40 * time.Millisecond),
					Path:      staticOctetsPath,
				},
			},
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(50 * time.Millisecond),
					Path:      staticOctetsPath,
				}}).SetVal(111),
		},
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Telemetry().Interface(tt.inPortKey).Counters().OutOctets().Collect(t, time.Second).Await(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				if i < len(tt.wantQualified) {
					tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
				}
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
		})
	}
}

func TestWatch(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_juniper")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/out-octets", iname)
	}

	staticPortName := "Ethernet3/1/1"
	staticOctetsPath := gnmiPath(t, getStrPath(staticPortName))

	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inPortKey            string
		predicate            func(*telemetry.QualifiedUint64) bool
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedUint64
		wantStatus           bool
	}{{
		desc: "no values at path",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(time.Minute).UnixNano(),
				})
		},
		predicate: func(val *telemetry.QualifiedUint64) bool {
			return !val.IsPresent()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantStatus:           true,
		wantQualified: &telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Path: staticOctetsPath,
			},
		},
	}, {
		desc: "values and predicate never true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: staticOctetsPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
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
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      staticOctetsPath,
			},
		}).SetVal(80),
	}, {
		desc: "values and predicate becomes true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: staticOctetsPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: staticOctetsPath,
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
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantStatus:           true,
		wantQualified: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      staticOctetsPath,
			},
		}).SetVal(120),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got, predStatus := dut.Telemetry().Interface(tt.inPortKey).Counters().OutOctets().Watch(t, time.Second, tt.predicate).Await(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			}

			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
			}
		})
	}

	t.Run("multiple awaits", func(t *testing.T) {
		fakeGNMI.Stub().Sync().Notification(&gpb.Notification{Timestamp: startTime.Add(time.Second).UnixNano()})

		watcher := dut.Telemetry().Interface("eth0").Counters().OutOctets().Watch(t, time.Second, func(val *telemetry.QualifiedUint64) bool { return true })
		got, gotStatus := watcher.Await(t)

		want := &telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Path: gnmiPath(t, getStrPath("eth0")),
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
	initTelemetryFakes()
	dut := DUT(t, "dut_arista")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/out-octets", iname)
	}

	ethPath := gnmiPath(t, getStrPath("eth0"))
	loPath := gnmiPath(t, getStrPath("lo0"))
	wildcardPath := gnmiPath(t, getStrPath("*"))

	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inPortKey            string
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedUint64
		wantStatus           bool
	}{{
		desc: "values and predicate never true",
		stub: func(s *fakegnmi.Stubber) {
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
		wantSubscriptionPath: wildcardPath,
		wantQualified: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      loPath,
			},
		}).SetVal(50),
	}, {
		desc: "multiple values in notfication",
		stub: func(s *fakegnmi.Stubber) {
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
		wantSubscriptionPath: wildcardPath,
		wantQualified: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime,
				Path:      loPath,
			},
		}).SetVal(50),
	}, {
		desc: "values and predicate becomes true",
		stub: func(s *fakegnmi.Stubber) {
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
		wantSubscriptionPath: wildcardPath,
		wantStatus:           true,
		wantQualified: (&telemetry.QualifiedUint64{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      loPath,
			},
		}).SetVal(150),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())

			var ethCond, loCond bool
			got, predStatus := dut.Telemetry().InterfaceAny().Counters().OutOctets().Watch(t, time.Second, func(val *telemetry.QualifiedUint64) bool {
				if !ethCond {
					ethCond = val.IsPresent() && val.Path.String() == ethPath.String() && val.Val(t) < 100
				}
				if !loCond {
					loCond = val.IsPresent() && val.Path.String() == loPath.String() && val.Val(t) > 100
				}
				return ethCond && loCond
			}).Await(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp

			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
			}
		})
	}
}

func TestNonleafWildcardWatch(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_cisco")

	getStrPath := func(iname, counter string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/%s", iname, counter)
	}

	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inPortKey            string
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedInterface_Counters
		wantStatus           bool
	}{{
		desc: "predicate never true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: gnmiPath(t, getStrPath("eth0", "in-octets")),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: gnmiPath(t, getStrPath("eth0", "out-octets")),
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
						}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		wantSubscriptionPath: gnmiPath(t, "interfaces/interface[name=*]/state/counters"),
		wantQualified: (&telemetry.QualifiedInterface_Counters{
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
						Path: gnmiPath(t, getStrPath("lo0", "in-octets")),
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 150}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: gnmiPath(t, getStrPath("eth0", "out-octets")),
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
						}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		wantStatus:           true,
		wantSubscriptionPath: gnmiPath(t, "interfaces/interface[name=*]/state/counters"),
		wantQualified: (&telemetry.QualifiedInterface_Counters{
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
			got, predStatus := dut.Telemetry().InterfaceAny().Counters().Watch(t, time.Second, func(val *telemetry.QualifiedInterface_Counters) bool {
				lo0 = lo0 || (val.GetPath().GetElem()[1].GetKey()["name"] == "lo0" && val.Val(t).GetInOctets() == 150)
				eth0 = eth0 || (val.GetPath().GetElem()[1].GetKey()["name"] == "eth0" && val.Val(t).GetOutOctets() == 50)
				return lo0 && eth0
			}).Await(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if got != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			}

			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedInterface_Counters{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
			}
		})
	}
}

func TestBatchGet(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_juniper")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/out-octets", iname)
	}
	getStrConfigPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/config/enabled", iname)
	}

	port1Name := "Ethernet3/1/1"
	port2Name := "Ethernet3/1/2"
	port1Path := gnmiPath(t, getStrPath(port1Name))
	port2Path := gnmiPath(t, getStrPath(port2Name))
	wildcardPath := gnmiPath(t, getStrPath("*"))
	enabledConfigPath := gnmiPath(t, getStrConfigPath(port1Name))

	tests := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		addPaths             func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch)
		wantSubscriptionPath []*gpb.Path
		wantQualified        *telemetry.QualifiedDevice
	}{{
		desc: "two leaves in single notifcation",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}, {
						Path: enabledConfigPath,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_BoolVal{BoolVal: true}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantQualified: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name: ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}, {
		desc: "leaves with prefix",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Prefix: &gpb.Path{
						Elem:   port1Path.Elem[0:1],
						Origin: "openconfig",
					},
					Update: []*gpb.Update{{
						Path: &gpb.Path{
							Elem: port1Path.Elem[1:],
						},
						Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: &gpb.Path{
							Elem: port2Path.Elem[1:],
						},
						Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantQualified: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name: ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}, { // TODO: Decide whether batch Lookup should validate that data matches subscribed paths
		desc: "non-subscribed paths",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port2Path},
		wantQualified: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Path: &gpb.Path{
					Origin: "openconfig",
				},
				Timestamp: time.Unix(0, 100),
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name: ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}, {
		desc: "two leaves with wildcard subcription",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: 100,
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 10}},
					}, {
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 12}},
					}}}).
				Sync()
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.InterfaceAny().Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{wildcardPath},
		wantQualified: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: time.Unix(0, 100),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			},
		}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name: ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(10),
					},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(12),
					},
				},
			},
		}),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			batch := dut.Telemetry().NewBatch()
			tt.addPaths(t, dut.Telemetry(), batch)
			got := batch.Lookup(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath...)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantQualified, got, cmpopts.EquateErrors(), cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Lookup() return unexpected device diff(-want,+got):\n %s", diff)
			}
		})
	}
}

func TestBatchWatch(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_arista")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/out-octets", iname)
	}

	port1Name := "Ethernet3/1/1"
	port2Name := "Ethernet3/1/2"
	port1Path := gnmiPath(t, getStrPath(port1Name))
	port2Path := gnmiPath(t, getStrPath(port2Name))
	wildcardPath := gnmiPath(t, getStrPath("*"))

	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		addPaths             func(testing.TB, *device.DevicePath, *telemetry.Batch)
		wantSubscriptionPath []*gpb.Path
		wantPredicateArgs    []*telemetry.QualifiedDevice
		wantWatchRet         *telemetry.QualifiedDevice
		wantPredStatus       bool
	}{{
		desc: "multiple notifications before sync and predicate never true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 30}},
					}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: port1Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 50}},
						}, {
							Path: port2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
						}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantWatchRet: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name: ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(50),
					},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(80),
					},
				},
			},
		}),
		wantPredicateArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(50),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(80),
						},
					},
				},
			}),
		},
	}, {
		desc: "collection with delete",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Delete:    []*gpb.Path{port1Path},
						Update: []*gpb.Update{{
							Path: port2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantWatchRet: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name:     ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(120),
					},
				},
			},
		}),
		wantPredicateArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
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
					port1Name: {
						Name:     ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "predicate becomes true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: port2Path,
							Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 120}},
						}}}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(time.Minute).UnixNano(),
					})
		},
		addPaths: func(t testing.TB, dev *device.DevicePath, b *telemetry.Batch) {
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantPredStatus:       true,
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantWatchRet: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name: ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(75),
					},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(120),
					},
				},
			},
		}),
		wantPredicateArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
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
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "wildcard predicate becomes true",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}}}).
				Sync().
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
						Update: []*gpb.Update{{
							Path: port2Path,
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
		wantPredStatus:       true,
		wantSubscriptionPath: []*gpb.Path{wildcardPath},
		wantWatchRet: (&telemetry.QualifiedDevice{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),

				Path: &gpb.Path{
					Origin: "openconfig",
				},
			}}).SetVal(&telemetry.Device{
			Interface: map[string]*telemetry.Interface{
				port1Name: {
					Name: ygot.String(port1Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(75),
					},
				},
				port2Name: {
					Name: ygot.String(port2Name),
					Counters: &telemetry.Interface_Counters{
						OutOctets: ygot.Uint64(120),
					},
				},
			},
		}),
		wantPredicateArgs: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				}}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
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
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
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
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantWatchRet:         nil,
		wantPredicateArgs:    []*telemetry.QualifiedDevice{},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			batch := dut.Telemetry().NewBatch()
			tt.addPaths(t, dut.Telemetry(), batch)
			i := 0

			got, predStatus := batch.Watch(t, 250*time.Millisecond, func(val *telemetry.QualifiedDevice) bool {
				checkJustReceived(t, val.RecvTimestamp)
				tt.wantPredicateArgs[i].RecvTimestamp = val.RecvTimestamp
				if diff := cmp.Diff(tt.wantPredicateArgs[i], val, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
					t.Errorf("Predicate received unexpected device diff (-want,+got):\n %s", diff)
				}
				i++
				return val.IsPresent() && val.Val(t).GetInterface(port1Name).GetCounters().GetOutOctets() > 60 &&
					val.Val(t).GetInterface(port2Name).GetCounters().GetOutOctets() > 100
			}).Await(t)
			if i != len(tt.wantPredicateArgs) {
				t.Errorf("Predicate didn't receive all qualified devices: got %d, want %d", i, len(tt.wantPredicateArgs))
			}
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath...)
			if tt.wantWatchRet != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantWatchRet.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantWatchRet, got, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Watch() returned unexpected device diff (-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantPredStatus {
				t.Errorf("Watch() returned unexpected status: got %v, want %v ", predStatus, tt.wantPredStatus)
			}
		})
	}
}

func TestBatchCollect(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_cisco")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters/out-octets", iname)
	}

	port1Name := "Ethernet3/1/1"
	port2Name := "Ethernet3/1/2"
	port1Path := gnmiPath(t, getStrPath(port1Name))
	port2Path := gnmiPath(t, getStrPath(port2Name))
	wildcardPath := gnmiPath(t, getStrPath("*"))

	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		addPaths             func(testing.TB, *device.DevicePath, *telemetry.Batch)
		wantSubscriptionPath []*gpb.Path
		wantQualified        []*telemetry.QualifiedDevice
	}{{
		desc: "collection with delete",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Delete:    []*gpb.Path{port1Path},
					Update: []*gpb.Update{{
						Path: port2Path,
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
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantQualified: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
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
					port1Name: {
						Name:     ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "multiple paths",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: port2Path,
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
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantQualified: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,

					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
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
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "multiple notifications before sync",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: port2Path,
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
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantQualified: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
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
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "wildcard subscription",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: port1Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 75}},
					}, {
						Path: port2Path,
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 80}},
					}},
				},
			).Sync().Notification(
				&gpb.Notification{
					Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
					Update: []*gpb.Update{{
						Path: port2Path,
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
		wantSubscriptionPath: []*gpb.Path{wildcardPath},
		wantQualified: []*telemetry.QualifiedDevice{
			(&telemetry.QualifiedDevice{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path: &gpb.Path{
						Origin: "openconfig",
					},
				},
			}).SetVal(&telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
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
					port1Name: {
						Name: ygot.String(port1Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(75),
						},
					},
					port2Name: {
						Name: ygot.String(port2Name),
						Counters: &telemetry.Interface_Counters{
							OutOctets: ygot.Uint64(120),
						},
					},
				},
			}),
		},
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
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
			dev.Interface(port1Name).Counters().OutOctets().Batch(t, b)
			dev.Interface(port2Name).Counters().OutOctets().Batch(t, b)
		},
		wantSubscriptionPath: []*gpb.Path{port1Path, port2Path},
		wantQualified:        nil,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			batch := dut.Telemetry().NewBatch()
			tt.addPaths(t, dut.Telemetry(), batch)

			got := batch.Collect(t, 250*time.Millisecond).Await(t)
			for i, qt := range got {
				checkJustReceived(t, qt.RecvTimestamp)
				tt.wantQualified[i].RecvTimestamp = qt.RecvTimestamp
			}

			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath...)
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedDevice{}), protocmp.Transform()); diff != "" {
				t.Errorf("Collect() returned unexpected device diff (-want,+got):\n %s", diff)
			}
		})
	}
}

func TestNonleafWatch(t *testing.T) {
	initTelemetryFakes()
	dut := DUT(t, "dut_juniper")

	getStrPath := func(iname string) string {
		return fmt.Sprintf("interfaces/interface[name=%s]/state/counters", iname)
	}

	port1Name := "Ethernet3/1/1"
	containerPath := gnmiPath(t, getStrPath(port1Name))
	inOctPath := gnmiPath(t, getStrPath(port1Name)+"/in-octets")
	outOctPath := gnmiPath(t, getStrPath(port1Name)+"/out-octets")

	startTime := time.Now()

	// All of stubbed responses need to end with a notification timestamped
	// after the collect has timed out; otherwise the fake encounters an EOF.
	tests := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		wantSubscriptionPath *gpb.Path
		wantPredicateArgs    []*telemetry.QualifiedInterface_Counters
		wantWatchRet         *telemetry.QualifiedInterface_Counters
		wantPredStatus       bool
	}{{
		desc: "multiple notifications before sync and predicate never true",
		stub: func(s *fakegnmi.Stubber) {
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
		wantSubscriptionPath: containerPath,
		wantPredicateArgs: []*telemetry.QualifiedInterface_Counters{
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      containerPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets:  ygot.Uint64(50),
				OutOctets: ygot.Uint64(80),
			}),
		},
		wantWatchRet: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      containerPath,
			},
		}).SetVal(&telemetry.Interface_Counters{
			InOctets:  ygot.Uint64(50),
			OutOctets: ygot.Uint64(80),
		}),
	}, {
		desc: "with delete",
		stub: func(s *fakegnmi.Stubber) {
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
		wantSubscriptionPath: containerPath,
		wantPredicateArgs: []*telemetry.QualifiedInterface_Counters{
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      containerPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets: ygot.Uint64(75),
			}),
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      containerPath,
				},
			}).SetVal(&telemetry.Interface_Counters{}),
		},
		wantWatchRet: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      containerPath,
			},
		}).SetVal(&telemetry.Interface_Counters{}),
	}, {
		desc: "predicate becomes true",
		stub: func(s *fakegnmi.Stubber) {
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
		wantPredStatus:       true,
		wantSubscriptionPath: containerPath,
		wantPredicateArgs: []*telemetry.QualifiedInterface_Counters{
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime,
					Path:      containerPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets:  ygot.Uint64(75),
				OutOctets: ygot.Uint64(80),
			}),
			(&telemetry.QualifiedInterface_Counters{
				Metadata: &genutil.Metadata{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Path:      containerPath,
				},
			}).SetVal(&telemetry.Interface_Counters{
				InOctets:  ygot.Uint64(120),
				OutOctets: ygot.Uint64(80),
			}),
		},
		wantWatchRet: (&telemetry.QualifiedInterface_Counters{
			Metadata: &genutil.Metadata{
				Timestamp: startTime.Add(10 * time.Millisecond),
				Path:      containerPath,
			},
		}).SetVal(&telemetry.Interface_Counters{
			InOctets:  ygot.Uint64(120),
			OutOctets: ygot.Uint64(80),
		}),
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
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
		wantSubscriptionPath: containerPath,
		wantWatchRet:         nil,
		wantPredicateArgs:    []*telemetry.QualifiedInterface_Counters{},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			dut.Telemetry()
			i := 0

			got, predStatus := dut.Telemetry().Interface(port1Name).Counters().Watch(t, 250*time.Millisecond, func(val *telemetry.QualifiedInterface_Counters) bool {
				checkJustReceived(t, val.RecvTimestamp)
				tt.wantPredicateArgs[i].RecvTimestamp = val.RecvTimestamp
				if diff := cmp.Diff(tt.wantPredicateArgs[i], val, cmp.AllowUnexported(telemetry.QualifiedInterface_Counters{}), protocmp.Transform()); diff != "" {
					t.Errorf("Predicate received unexpected device diff (-want,+got):\n %s", diff)
				}
				i++
				return val.IsPresent() && val.Val(t).GetOutOctets() > 60 && val.Val(t).GetInOctets() > 100
			}).Await(t)
			if i != len(tt.wantPredicateArgs) {
				t.Errorf("Predicate didn't receive all qualified devices: got %d, want %d", i, len(tt.wantPredicateArgs))
			}
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			if tt.wantWatchRet != nil {
				checkJustReceived(t, got.RecvTimestamp)
				tt.wantWatchRet.RecvTimestamp = got.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantWatchRet, got, cmp.AllowUnexported(telemetry.QualifiedInterface_Counters{}), protocmp.Transform()); diff != "" {
				t.Errorf("Watch() returned unexpected device diff (-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantPredStatus {
				t.Errorf("Watch() returned unexpected status: got %v, want %v ", predStatus, tt.wantPredStatus)
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
