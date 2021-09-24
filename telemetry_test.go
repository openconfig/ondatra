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
	"github.com/openconfig/ondatra/internal/fakegnmi"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/negtest"
	"github.com/openconfig/ondatra/telemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var fakeGNMI = func() *fakegnmi.FakeGNMI {
	fg, err := fakegnmi.Start(0)
	if err != nil {
		log.Fatalf("could not start fake gNMI server: %v", err)
	}
	return fg
}()

func initTelemetryFakes(t *testing.T) {
	t.Helper()
	initFakeBinding(t)
	reserveFakeTestbed(t)
	fakeBind.GNMIDialer = func(ctx context.Context, _ *reservation.DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
		return fakeGNMI.Dial(ctx, opts...)
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
	initTelemetryFakes(t)
	dut := DUT(t, "dut")
	connPath := gnmiPath(t, "meta/connected")
	syncPath := gnmiPath(t, "meta/sync")
	leavesAddedPath := gnmiPath(t, "meta/targetLeavesAdded")

	type qualifiedType interface {
		GetTimestamp() time.Time
		GetRecvTimestamp() time.Time
	}

	testsPass := []struct {
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		getFunc              func(t testing.TB) qualifiedType
		wantSubscriptionPath *gpb.Path
		wantQualified        qualifiedType
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
		getFunc: func(t testing.TB) qualifiedType {
			return dut.Telemetry().Meta().Connected().GetFull(t)
		},
		wantSubscriptionPath: connPath,
		wantQualified: (&telemetry.QualifiedBool{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		getFunc: func(t testing.TB) qualifiedType {
			return dut.Telemetry().Meta().Connected().GetFull(t)
		},
		wantSubscriptionPath: connPath,
		wantQualified: (&telemetry.QualifiedBool{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		getFunc: func(t testing.TB) qualifiedType {
			return dut.Telemetry().Meta().Sync().GetFull(t)
		},
		wantSubscriptionPath: syncPath,
		wantQualified: (&telemetry.QualifiedBool{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		getFunc: func(t testing.TB) qualifiedType {
			return dut.Telemetry().Meta().Sync().GetFull(t)
		},
		wantSubscriptionPath: syncPath,
		wantQualified: (&telemetry.QualifiedBool{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		getFunc: func(t testing.TB) qualifiedType {
			return dut.Telemetry().Meta().TargetLeavesAdded().GetFull(t)
		},
		wantSubscriptionPath: leavesAddedPath,
		wantQualified: (&telemetry.QualifiedInt64{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 200),
				Present:   true,
				Path:      leavesAddedPath,
			}}).SetVal(42),
	}, {
		desc: "no values for leaves added",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(&gpb.Notification{}).Sync()
		},
		getFunc: func(t testing.TB) qualifiedType {
			return dut.Telemetry().Meta().TargetLeavesAdded().GetFull(t)
		},
		wantSubscriptionPath: leavesAddedPath,
		wantQualified: &(telemetry.QualifiedInt64{
			QualifiedType: &genutil.QualifiedType{
				Present: false,
				Path:    leavesAddedPath,
			}}),
	}, {
		desc: "no values for leaves added without connected return value",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		getFunc: func(t testing.TB) qualifiedType {
			return dut.Telemetry().Meta().TargetLeavesAdded().GetFull(t)
		},
		wantSubscriptionPath: leavesAddedPath,
		wantQualified: &(telemetry.QualifiedInt64{
			QualifiedType: &genutil.QualifiedType{
				Present: false,
				Path:    leavesAddedPath,
			}}),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := tt.getFunc(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)

			checkJustReceived(t, got.GetRecvTimestamp())
			if diff := cmp.Diff(tt.wantQualified, got, cmpopts.IgnoreFields(telemetry.QualifiedBool{}, "RecvTimestamp"), cmpopts.IgnoreFields(telemetry.QualifiedInt64{}, "RecvTimestamp"), cmp.AllowUnexported(telemetry.QualifiedBool{}, telemetry.QualifiedInt64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 101),
				Present:   true,
				Path:      statusPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: statusPath,
		wantQualified: &(telemetry.QualifiedE_Interface_OperStatus{
			QualifiedType: &genutil.QualifiedType{
				Present: false,
				Path:    statusPath,
			}}),
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 102),
				Present:   true,
				Path:      statusPath,
			}}).SetVal(telemetry.Interface_OperStatus_UP),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Telemetry().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(tt.inPortKey).OperStatus().GetFull(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedE_Interface_OperStatus{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantQualified.Present {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
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
		wantFatalMsg: "2 data points",
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
		wantFatalMsg: "query-noncompliant",
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
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				dut.Telemetry().Interface(staticPortName).OperStatus().GetFull(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("GetFull failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestGetDefault(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")

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
		inGetCall:            dut.Telemetry().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(staticPortName).Enabled().GetFull,
		wantSubscriptionPath: enabledPath,
		wantQualified: (&telemetry.QualifiedBool{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
				Path:      enabledPath,
			}}).SetVal(false),
	}, {
		desc: "telemetry path API: zero notifications on a leaf that has a default value",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inGetCall:            dut.Telemetry().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(staticPortName).Enabled().GetFull,
		wantSubscriptionPath: enabledPath,
		wantQualified: (&telemetry.QualifiedBool{
			QualifiedType: &genutil.QualifiedType{
				Present: true,
				Path:    enabledPath,
			}}).SetVal(true),
	}, {
		desc: "config path API: zero notifications on a leaf that has a default value",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inGetCall:            dut.Config().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(staticPortName).Enabled().GetFull,
		wantSubscriptionPath: enabledConfigPath,
		wantQualified: (&telemetry.QualifiedBool{
			QualifiedType: &genutil.QualifiedType{
				Present: true,
				Path:    enabledConfigPath,
			}}).SetVal(true),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := tt.inGetCall(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedBool{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got status Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantQualified.Present {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
				Path:      resolvedPath,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 101),
				Present:   true,
				Path:      descPath,
			}}).SetVal("foo"),
	}, {
		desc: "no values",
		stub: func(s *fakegnmi.Stubber) {
			s.Sync()
		},
		inPortKey:            staticPortName,
		wantSubscriptionPath: descPath,
		wantQualified: &(telemetry.QualifiedString{
			QualifiedType: &genutil.QualifiedType{
				Present: false,
				Path:    descPath,
			}}),
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Config().WithReplica(5).WithSubscriptionMode(gpb.SubscriptionMode_ON_CHANGE).Interface(tt.inPortKey).Description().GetFull(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedString{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got desc Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantQualified.Present {
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
		wantFatalMsg: "2 data points",
	}, {
		desc: "returned prefix + path is incompatible",
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
		wantFatalMsg: "query-noncompliant",
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
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				dut.Config().Interface(staticPortName).Description().GetFull(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("GetFull failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestGetNonleaf(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")
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
		desc                 string
		stub                 func(s *fakegnmi.Stubber)
		inGetCall            func(t testing.TB) *telemetry.QualifiedInterface
		wantSubscriptionPath *gpb.Path
		wantQualified        *telemetry.QualifiedInterface
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
				Path:      staticIntfPath,
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		inGetCall:            dut.Config().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		inGetCall:            dut.Config().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		inGetCall:            dut.Config().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				// This value should not be present, but is due to https://github.com/openconfig/ygot/issues/544
				Present: true,
				Path:    staticIntfPath,
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
		inGetCall:            dut.Telemetry().Interface(genericPortName).GetFull,
		wantSubscriptionPath: resolvedIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 200),
				Present:   true,
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 400),
				Present:   true,
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
		inGetCall:            dut.Telemetry().Interface(staticPortName).GetFull,
		wantSubscriptionPath: staticIntfPath,
		wantQualified: &telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Present: false,
				Path:    staticIntfPath,
			}},
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := tt.inGetCall(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s\nComplianceErrors:\n%v", diff, got.ComplianceErrors)
			}
			// Test Val(t) API.
			if tt.wantQualified.Present {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
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
		inGetCall:            dut.Telemetry().GetFull,
		wantSubscriptionPath: rootPath,
		wantQualified: (&telemetry.QualifiedDevice{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 100),
				Present:   true,
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
			if tt.wantQualified.Present {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t)); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}
		})
	}

	testsFail := []struct {
		desc                  string
		stub                  func(s *fakegnmi.Stubber)
		inInterfacePath       func() *telemetry.InterfacePath
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
		inInterfacePath: func() *telemetry.InterfacePath {
			return dut.Telemetry().Interface(staticPortName)
		},
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Present: false,
				Path:    staticIntfPath,
			}}),
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
		inInterfacePath: func() *telemetry.InterfacePath {
			return dut.Telemetry().Interface(genericPortName)
		},
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Timestamp: time.Unix(0, 200),
				Present:   true,
				Path:      resolvedIntfPath,
			}}).SetVal(&telemetry.Interface{
			Enabled: ygot.Bool(true),
		}),
		wantTypeErrSubstr: "failed to unmarshal",
	}, {
		desc: "returned prefix + path is incompatible",
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
		inInterfacePath: func() *telemetry.InterfacePath {
			return dut.Telemetry().Interface(staticPortName)
		},
		wantQualified: (&telemetry.QualifiedInterface{
			QualifiedType: &genutil.QualifiedType{
				Present: false,
				Path:    staticIntfPath,
			}}),
		wantPathErrSubstr: "no match found",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			if tt.wantFatalMsg != "" {
				got := negtest.ExpectFatal(t, func(t testing.TB) {
					tt.inInterfacePath().GetFull(t)
				})
				if !strings.Contains(got, tt.wantFatalMsg) {
					t.Errorf("GetFull failed with message %q, want %q", got, tt.wantFatalMsg)
				}
				return
			}

			got := tt.inInterfacePath().GetFull(t)
			checkJustReceived(t, got.RecvTimestamp)
			tt.wantQualified.RecvTimestamp = got.RecvTimestamp
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedInterface{}), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got Qualified type different from expected (-want,+got):\n %s", diff)
			}
			// Test Val(t) API.
			if tt.wantQualified.Present {
				if diff := cmp.Diff(tt.wantQualified.Val(t), got.Val(t), cmpopts.IgnoreTypes(&genutil.ComplianceErrors{})); diff != "" {
					t.Errorf("Got status val different from expected (-want,+got):\n %s", diff)
				}
			}

			verifyTelemetryErrSubstr(t, got.ComplianceErrors, tt.wantPathErrSubstr, tt.wantTypeErrSubstr, tt.wantValidateErrSubstr)
		})
	}
}

func TestWildcardGet(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")

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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      intf1Path,
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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      intf1Path,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
					Path:      intf2Path,
				}}).SetVal(456),
		},
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := dut.Telemetry().InterfaceAny().Counters().InOctets().GetFull(t)
			verifySubscriptionPathsSent(t, wantSubscriptionPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
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
				t.Errorf("GetFull(wildcard leaf): Got values different from expected (-want,+got):\n %s", diff)
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
		wantFatalMsg: "query-noncompliant",
	}, {
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
	}, {
		desc: "schema noncompliance (update ignored for groups that are entirely noncompliant): one notification that doesn't unmarshal because type doesn't match",
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
		wantFatalMsg: "failed to unmarshal",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				dut.Telemetry().InterfaceAny().Counters().InOctets().GetFull(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Wildcard Get failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestWildcardNonleafGet(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")

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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: time.Unix(0, 100),
					Present:   true,
					Path:      staticIntfPath,
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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: time.Unix(0, 200),
					Present:   true,
					Path:      resolvedIntfPath,
				}}).SetVal(&telemetry.Interface{
				OperStatus: telemetry.Interface_OperStatus_UP,
				Enabled:    ygot.Bool(true),
			}),
			(&telemetry.QualifiedInterface{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: time.Unix(0, 400),
					Present:   true,
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
			got := dut.Telemetry().InterfaceAny().GetFull(t)
			verifySubscriptionPathsSent(t, wantSubscriptionPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
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
				t.Errorf("GetFull(wildcard non-leaf): Got values different from expected (-want,+got):\n %s", diff)
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
				got := negtest.ExpectFatal(t, func(t testing.TB) {
					dut.Telemetry().InterfaceAny().GetFull(t)
				})
				if !strings.Contains(got, tt.wantFatalMsg) {
					t.Errorf("GetFull failed with message %q, want %q", got, tt.wantFatalMsg)
				}
				return
			}
			got := dut.Telemetry().InterfaceAny().GetFull(t)
			if len(got) != 1 {
				t.Fatalf("Expected exactly one piece of data for this test, but got %v: %v", len(got), got)
			}

			verifyTelemetryErrSubstr(t, got[0].GetComplianceErrors(), tt.wantPathErrSubstr, tt.wantTypeErrSubstr, tt.wantValidateErrSubstr)
		})
	}
}

func TestCollect(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")
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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(456),
		},
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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
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
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(123),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(456),
			{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(20 * time.Millisecond),
					Present:   false,
					Path:      staticOctetsPath,
				},
			},
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(30 * time.Millisecond),
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(789),
			{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(40 * time.Millisecond),
					Present:   false,
					Path:      staticOctetsPath,
				},
			},
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(50 * time.Millisecond),
					Present:   true,
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
				tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
		})
	}

	testsFail := []struct {
		desc         string
		stub         func(s *fakegnmi.Stubber)
		wantFatalMsg string
	}{{
		desc: "multiple values with prefixes but the last one is wrong",
		stub: func(s *fakegnmi.Stubber) {
			s.Notification(
				&gpb.Notification{
					Timestamp: startTime.UnixNano(),
					Update: []*gpb.Update{{
						Path: &gpb.Path{Elem: staticOctetsPath.Elem[len(staticOctetsPath.Elem)-1:]},
						Val:  &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 123}},
					}},
					Prefix: &gpb.Path{Elem: staticOctetsPath.Elem[:len(staticOctetsPath.Elem)-1]},
				}).
				Notification(
					&gpb.Notification{
						Timestamp: startTime.Add(10 * time.Millisecond).UnixNano(),
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
		wantFatalMsg: "no match found",
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
		wantFatalMsg: "failed to unmarshal",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				dut.Telemetry().Interface(staticPortName).Counters().OutOctets().Collect(t, time.Second).Await(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Collect failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestCollectUntil(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")

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
		wantSubscriptionPath *gpb.Path
		wantQualified        []*telemetry.QualifiedUint64
		wantStatus           bool
	}{{
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
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(50),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(80),
		},
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
		inPortKey:            staticPortName,
		wantSubscriptionPath: staticOctetsPath,
		wantStatus:           true,
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(50),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
					Path:      staticOctetsPath,
				}}).SetVal(120),
		},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())
			got, predStatus := dut.Telemetry().Interface(tt.inPortKey).Counters().OutOctets().CollectUntil(t, time.Second, func(val *telemetry.QualifiedUint64) bool {
				if val.IsPresent() {
					return val.Val(t) > 100
				}
				return false
			}).Await(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
			}
		})
	}
}

func TestWildcardCollectUntil(t *testing.T) {
	initTelemetryFakes(t)
	dut := DUT(t, "dut")

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
		wantQualified        []*telemetry.QualifiedUint64
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
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      ethPath,
				}}).SetVal(150),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
					Path:      loPath,
				}}).SetVal(50),
		},
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
		wantQualified: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime,
					Present:   true,
					Path:      ethPath,
				}}).SetVal(50),
			(&telemetry.QualifiedUint64{
				QualifiedType: &genutil.QualifiedType{
					Timestamp: startTime.Add(10 * time.Millisecond),
					Present:   true,
					Path:      loPath,
				}}).SetVal(150),
		},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.stub(fakeGNMI.Stub())

			var ethCond, loCond bool
			got, predStatus := dut.Telemetry().InterfaceAny().Counters().OutOctets().CollectUntil(t, time.Second, func(val *telemetry.QualifiedUint64) bool {
				if !ethCond {
					ethCond = val.IsPresent() && val.Path.String() == ethPath.String() && val.Val(t) < 100
				}
				if !loCond {
					loCond = val.IsPresent() && val.Path.String() == loPath.String() && val.Val(t) > 100
				}
				return ethCond && loCond
			}).Await(t)
			verifySubscriptionPathsSent(t, tt.wantSubscriptionPath)
			for i, q := range got {
				checkJustReceived(t, q.RecvTimestamp)
				tt.wantQualified[i].RecvTimestamp = q.RecvTimestamp
			}
			if diff := cmp.Diff(tt.wantQualified, got, cmp.AllowUnexported(telemetry.QualifiedUint64{}), protocmp.Transform()); diff != "" {
				t.Errorf("Got out octets different from expected, diff(-want,+got):\n %s", diff)
			}
			if predStatus != tt.wantStatus {
				t.Errorf("Got different predicate status, got %v want %v ", predStatus, tt.wantStatus)
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
