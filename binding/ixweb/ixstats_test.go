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

package ixweb

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
)

func TestStatViews(t *testing.T) {
	sess := &Session{ixweb: &IxWeb{
		hostname: "fakeHost",
		client: &fakeHTTPClient{
			doResps: []*http.Response{fakeResponse(200, `[
				{"id": 1, "caption": "caption1"},
				{"id": 2, "caption": "caption2"}
			]`)},
		},
	}}
	wantViews := map[string]*StatView{
		"caption1": &StatView{sess: sess, id: 1, caption: "caption1"},
		"caption2": &StatView{sess: sess, id: 2, caption: "caption2"},
	}
	gotViews, err := sess.Stats().Views(context.Background())
	if err != nil {
		t.Fatalf("StatViews got unexpected err: %v", err)
	}
	cmpView := func(v1, v2 *StatView) bool {
		return v1.sess == v2.sess && v1.id == v2.id && v1.caption == v2.caption
	}
	if diff := cmp.Diff(wantViews, gotViews, cmp.Comparer(cmpView)); diff != "" {
		t.Fatalf("StatViews got unexpected diff (-want, +got): %s", diff)
	}
}

func TestConfigEgressStatView(t *testing.T) {
	fakeClient := &fakeHTTPClient{}
	stats := &Stats{sess: &Session{ixweb: &IxWeb{
		hostname: "fakeHost",
		client:   fakeClient,
	}}}

	tests := []struct {
		name       string
		egPageSize int
		doResps    []*http.Response
		wantErr    string
	}{{
		name:       "no existing",
		egPageSize: 1,
		doResps: []*http.Response{
			fakeResponse(200, "[]"),
			fakeResponse(200, `{"links": [{"href": "/api/v1/sessions/1/ixnetwork/statistics/view/11"}]}`),
			fakeResponse(200, `[{"name": "portFilter1", "links": [{"href": "pf1"}]}]`),
			fakeResponse(200, `[{"name": "trafficItemFilter1", "links": [{"href": "tif1"}]}]`),
			fakeResponse(200, "set layer23TrafficFlowFilter"),
			fakeResponse(200, `[{"name": "trackingFilter1", "links": [{"href": "tf1"}]}]`),
			fakeResponse(200, "set enumerationFilter"),
			fakeResponse(200, `[{"name": "statisticFilter1", "links": [{"href": "sf1"}]}]`),
			fakeResponse(200, "enabled statisticFilter"),
			fakeResponse(200, "enabled EgressStatView"),
			fakeResponse(200, "set page sizes"),
		},
	}, {
		name:       "delete existing",
		egPageSize: 1,
		doResps: []*http.Response{
			fakeResponse(200, `[{"id": 11, "caption": "EgressStatView"}]`),
			fakeResponse(200, "deleted EgressStatView"),
			fakeResponse(200, `{"links": [{"href": "/api/v1/sessions/1/ixnetwork/statistics/view/11"}]}`),
			fakeResponse(200, `[{"name": "portFilter1", "links": [{"href": "pf1"}]}]`),
			fakeResponse(200, `[{"name": "trafficItemFilter1", "links": [{"href": "tif1"}]}]`),
			fakeResponse(200, "set layer23TrafficFlowFilter"),
			fakeResponse(200, `[{"name": "trackingFilter1", "links": [{"href": "tf1"}]}]`),
			fakeResponse(200, "set enumerationFilter"),
			fakeResponse(200, `[{"name": "statisticFilter1", "links": [{"href": "sf1"}]}]`),
			fakeResponse(200, "enabled statisticFilter"),
			fakeResponse(200, "enabled EgressStatView"),
			fakeResponse(200, "set page sizes"),
		},
	}, {
		name:       "bad egress page size",
		egPageSize: 100,
		wantErr:    "not in allowed range",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fakeClient.doResps = test.doResps
			view, err := stats.ConfigEgressView(context.Background(), []string{"item1"}, test.egPageSize)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Fatalf("ConfigEgressStatView unexpected err: got %v, want %v", err, test.wantErr)
			}
			if test.wantErr != "" {
				return
			}
			if view.id != 11 {
				t.Errorf("ConfigEgressStatView unexpected id, got %v, want %v", view.id, 11)
			}
			if view.caption != EgressStatsCaption {
				t.Errorf("ConfigEgressStatView unexpected caption, got %v, want %v", view.caption, EgressStatsCaption)
			}
		})
	}
}

func TestFetchTable(t *testing.T) {
	view := &StatView{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{doResps: []*http.Response{
		fakeResponse(200, "enabled CSV logging"),
		fakeResponse(200, `{"isReady": true}`),
		fakeResponse(200, "ran takeviewcsvsnapshot"),
		fakeResponse(200, "col1,col2,col3\na1,a2,a3\nb1,b2,b3"),
	}}}}}
	wantTable := StatTable{
		StatRow{"col1": "a1", "col2": "a2", "col3": "a3"},
		StatRow{"col1": "b1", "col2": "b2", "col3": "b3"},
	}
	gotTable, err := view.FetchTable(context.Background())
	if err != nil {
		t.Fatalf("FetchTable got unexpected error: %v", err)
	}
	if diff := cmp.Diff(wantTable, gotTable); diff != "" {
		t.Fatalf("FetchTable got unexpected diff (-want +got): %s", diff)
	}
}

func TestFetchEgressStatsTable(t *testing.T) {
	view := &StatView{
		caption: EgressStatsCaption,
		sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{doResps: []*http.Response{
			fakeResponse(200, `{"isReady": true}`),
			fakeResponse(200, "set current page"),
			fakeResponse(200, `{
				"columnCaptions": ["col1", "col2", "col3"],
				"pageValues": [[["a1", "a2", "a3"]]],
				"totalPages": 2
			}`),
			fakeResponse(200, "set current page"),
			fakeResponse(200, `{
				"columnCaptions": ["col1", "col2", "col3"],
				"pageValues": [[["b1", "b2", "b3"]]],
				"totalPages": 2
			}`),
		}}}},
	}
	wantTable := StatTable{
		StatRow{"col1": "a1", "col2": "a2", "col3": "a3"},
		StatRow{"col1": "b1", "col2": "b2", "col3": "b3"},
	}
	gotTable, err := view.FetchTable(context.Background())
	if err != nil {
		t.Fatalf("FetchTable got unexpected error: %v", err)
	}
	if diff := cmp.Diff(wantTable, gotTable); diff != "" {
		t.Fatalf("FetchTable got unexpected diff (-want +got): %s", diff)
	}
}

func TestFetchTableNotReady(t *testing.T) {
	view := &StatView{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{doResps: append(
		[]*http.Response{fakeResponse(200, "enabled CSV logging")},
		repeatResponses(200, `{"isReady": false}`, viewReadyAttempts)...,
	)}}}}
	gotTable, err := view.FetchTable(context.Background())
	if err == nil {
		t.Fatalf("FetchTable unexpectedly succeeded, got: %v", gotTable)
	}
	if wantSubstr := "timeout"; !strings.Contains(err.Error(), wantSubstr) {
		t.Fatalf("FetchTable unexpected error, got %v, want substring %q", err, wantSubstr)
	}
}

func TestFetchTableWithDrilldowns(t *testing.T) {
	origDrilldownFn := drilldownFn
	defer func() { drilldownFn = origDrilldownFn }()
	tests := []struct {
		desc      string
		resps     []*http.Response
		ddErr     error
		ddView    *StatView
		wantErr   string
		wantTable StatTable
	}{{
		desc: "drilldown error",
		resps: []*http.Response{
			fakeResponse(200, "enabled CSV logging"),
			fakeResponse(200, `{"isReady": true}`),
			fakeResponse(200, "ran takeviewcsvsnapshot"),
			fakeResponse(200, "col1,col2,col3\na1,a2,a3"),
		},
		ddErr:   errors.New("drilldown error"),
		wantErr: "could not perform drilldown",
	}, {
		desc: "no further views",
		resps: []*http.Response{
			fakeResponse(200, "enabled CSV logging"),
			fakeResponse(200, `{"isReady": true}`),
			fakeResponse(200, "ran takeviewcsvsnapshot"),
			fakeResponse(200, "col1,col2,col3\na1,a2,a3"),
		},
		wantTable: StatTable{
			StatRow{"col1": "a1", "col2": "a2", "col3": "a3"},
		},
	}, {
		desc: "too many rows in drilldown",
		resps: []*http.Response{
			fakeResponse(200, "enabled CSV logging"),
			fakeResponse(200, `{"isReady": true}`),
			fakeResponse(200, "ran takeviewcsvsnapshot"),
			fakeResponse(200, "col1,col2,col3\na1,a2,a3"),
		},
		ddView: &StatView{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{doResps: []*http.Response{
			fakeResponse(200, "enabled CSV logging"),
			fakeResponse(200, `{"isReady": true}`),
			fakeResponse(200, "ran takeviewcsvsnapshot"),
			fakeResponse(200, "ddCol\nddVal0\nddVal1"),
		}}}}},
		wantErr: "expected exactly 1",
	}, {
		desc: "update from drilldown",
		resps: []*http.Response{
			fakeResponse(200, "enabled CSV logging"),
			fakeResponse(200, `{"isReady": true}`),
			fakeResponse(200, "ran takeviewcsvsnapshot"),
			fakeResponse(200, "col1,col2,col3\na1,a2,a3"),
		},
		ddView: &StatView{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{doResps: []*http.Response{
			fakeResponse(200, "enabled CSV logging"),
			fakeResponse(200, `{"isReady": true}`),
			fakeResponse(200, "ran takeviewcsvsnapshot"),
			fakeResponse(200, "ddCol\nddVal0"),
		}}}}},
		wantTable: StatTable{
			StatRow{"col1": "a1", "col2": "a2", "col3": "a3", "ddCol": "ddVal0"},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			drilldownFn = func(*StatView, context.Context, int, string) (*StatView, map[string]*StatView, error) {
				var svMap map[string]*StatView
				if test.ddView != nil {
					svMap = map[string]*StatView{test.ddView.caption: test.ddView}
				}
				return test.ddView, svMap, test.ddErr
			}
			view := &StatView{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{doResps: test.resps}}}}
			gotTable, gotErr := view.FetchTable(context.Background(), "some drilldown caption")
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Fatalf("FetchTable: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.wantTable, gotTable); diff != "" {
				t.Fatalf("FetchTable got unexpected diff (-want +got): %s", diff)
			}
		})
	}
}

func TestDrilldown(t *testing.T) {
	const dd = "some drilldown"
	tests := []struct {
		desc            string
		resps           []*http.Response
		respErrs        []error
		wantNextCaption string
		wantViews       map[string]*StatView
		wantErr         string
	}{{
		desc:     "error fetching drill down options",
		respErrs: []error{errors.New("some error")},
		wantErr:  "could not retrieve drilldown options",
	}, {
		desc:  "no valid drilldown option",
		resps: []*http.Response{fakeResponse(200, "[]")},
	}, {
		desc:     "error on drilldown",
		resps:    []*http.Response{fakeResponse(200, fmt.Sprintf("[\"%s\"]", dd))},
		respErrs: []error{nil, errors.New("some error")},
		wantErr:  "could not perform stats drilldown",
	}, {
		desc: "successful drilldown",
		resps: []*http.Response{
			fakeResponse(200, fmt.Sprintf(`["%s"]`, dd)),
			fakeResponse(200, `"some view"`),
			fakeResponse(200, `[{"id": 20, "caption": "some view"}]`),
		},
		wantNextCaption: "some view",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			view := &StatView{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{
				doResps: test.resps,
				doErrs:  test.respErrs,
			}}}}
			gotView, _, gotErr := view.drilldown(context.Background(), 1, dd)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Fatalf("drilldown: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
			if test.wantNextCaption != "" && (gotView == nil || test.wantNextCaption != gotView.caption) {
				t.Fatalf("drilldown: did not return view with expected caption %q", test.wantNextCaption)
			}
		})
	}
}
