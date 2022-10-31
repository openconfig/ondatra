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

package ixgnmi

import (
	"golang.org/x/net/context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/openconfig/ygot/uexampleoc"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ondatra/telemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

func TestSubscribe(t *testing.T) {
	var stubStats []map[string]ygot.GoStruct
	readStatsFn = func(_ context.Context, _ *Client, _ string, captions []string) (ygot.GoStruct, error) {
		var data ygot.GoStruct = &uexampleoc.Device{}
		if len(stubStats) == 0 {
			return data, nil
		}
		stub := stubStats[0]
		stubStats = stubStats[1:]
		for _, cap := range captions {
			t, ok := stub[cap]
			if !ok {
				continue
			}
			var err error
			data, err = ygot.MergeStructs(data, t)
			if err != nil {
				return nil, err
			}
		}
		return data, nil
	}

	devWithAttr1 := &telemetry.Device{}
	devWithAttr1.GetOrCreateNetworkInstance("foo").GetOrCreateProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "1").
		GetOrCreateBgp().GetOrCreateRib().GetOrCreateAttrSet(1)

	devWithAttr2 := &telemetry.Device{}
	devWithAttr2.GetOrCreateNetworkInstance("foo").GetOrCreateProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "1").
		GetOrCreateBgp().GetOrCreateRib().GetOrCreateAttrSet(2)

	devWithPartial := &telemetry.Device{}
	devWithPartial.GetOrCreateNetworkInstance("foo")

	tests := []struct {
		name        string
		mode        gpb.SubscriptionList_Mode
		path        *gpb.Path
		stats       []map[string]ygot.GoStruct
		learnedInfo []*telemetry.Device
		want        []*gpb.SubscribeResponse
	}{{
		name: "no such path",
		path: &gpb.Path{Elem: []*gpb.PathElem{{Name: "bogus"}}},
		mode: gpb.SubscriptionList_ONCE,
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
		},
	}, {
		name: "components once",
		mode: gpb.SubscriptionList_ONCE,
		path: &gpb.Path{
			Origin: "openconfig",
			Elem:   []*gpb.PathElem{{Name: "components"}},
		},
		stats: []map[string]ygot.GoStruct{{
			portCPUStatsCaption: func() ygot.GoStruct {
				root := &uexampleoc.Device{}
				root.GetOrCreateComponents().GetOrCreateComponent("fakeComp")
				return root
			}(),
		}},
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Path: &gpb.Path{Elem: []*gpb.PathElem{
						{Name: "components"},
						{Name: "component", Key: map[string]string{"name": "fakeComp"}},
						{Name: "name"},
					}},
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"fakeComp"}}},
				},
			}}},
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
		},
	}, {
		name: "components stream",
		mode: gpb.SubscriptionList_STREAM,
		path: &gpb.Path{
			Origin: "openconfig",
			Elem:   []*gpb.PathElem{{Name: "components"}},
		},
		stats: []map[string]ygot.GoStruct{{
			portCPUStatsCaption: func() ygot.GoStruct {
				root := &uexampleoc.Device{}
				root.GetOrCreateComponents().GetOrCreateComponent("fakeComp")
				return root
			}(),
		}, {
			portCPUStatsCaption: func() ygot.GoStruct {
				root := &uexampleoc.Device{}
				root.GetOrCreateComponents().GetOrCreateComponent("fakeComp2")
				return root
			}(),
		}},
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Path: &gpb.Path{Elem: []*gpb.PathElem{
						{Name: "components"},
						{Name: "component", Key: map[string]string{"name": "fakeComp"}},
						{Name: "name"},
					}},
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"fakeComp"}}},
				},
			}}},
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Path: &gpb.Path{Elem: []*gpb.PathElem{
						{Name: "components"},
						{Name: "component", Key: map[string]string{"name": "fakeComp2"}},
						{Name: "name"},
					}},
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"fakeComp2"}}},
				},
			}}},
		},
	}, {
		name: "interfaces once",
		mode: gpb.SubscriptionList_ONCE,
		path: &gpb.Path{
			Origin: "openconfig",
			Elem:   []*gpb.PathElem{{Name: "interfaces"}},
		},
		stats: []map[string]ygot.GoStruct{{
			portStatsCaption: func() ygot.GoStruct {
				root := &uexampleoc.Device{}
				root.GetOrCreateInterfaces().GetOrCreateInterface("fakeIntf")
				return root
			}(),
		}},
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Path: &gpb.Path{Elem: []*gpb.PathElem{
						{Name: "interfaces"},
						{Name: "interface", Key: map[string]string{"name": "fakeIntf"}},
						{Name: "name"},
					}},
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"fakeIntf"}}},
				},
			}}},
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
		},
	}, {
		name: "interfaces stream",
		mode: gpb.SubscriptionList_STREAM,
		path: &gpb.Path{
			Origin: "openconfig",
			Elem:   []*gpb.PathElem{{Name: "interfaces"}},
		},
		stats: []map[string]ygot.GoStruct{{
			portStatsCaption: func() ygot.GoStruct {
				root := &uexampleoc.Device{}
				root.GetOrCreateInterfaces().GetOrCreateInterface("fakeIntf")
				return root
			}(),
		}, {
			portStatsCaption: func() ygot.GoStruct {
				root := &uexampleoc.Device{}
				root.GetOrCreateInterfaces().GetOrCreateInterface("fakeIntf2")
				return root
			}(),
		}},
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Path: &gpb.Path{Elem: []*gpb.PathElem{
						{Name: "interfaces"},
						{Name: "interface", Key: map[string]string{"name": "fakeIntf"}},
						{Name: "name"},
					}},
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"fakeIntf"}}},
				},
			}}},
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Path: &gpb.Path{Elem: []*gpb.PathElem{
						{Name: "interfaces"},
						{Name: "interface", Key: map[string]string{"name": "fakeIntf2"}},
						{Name: "name"},
					}},
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"fakeIntf2"}}},
				},
			}}},
		},
	}, {
		name: "bgp rib once",
		mode: gpb.SubscriptionList_ONCE,
		learnedInfo: []*telemetry.Device{
			devWithAttr1,
		},
		path: &gpb.Path{
			Origin: "openconfig",
			Elem: []*gpb.PathElem{
				{Name: "network-instances"},
				{Name: "network-instance", Key: map[string]string{"name": "foo"}},
				{Name: "protocols"},
				{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
				{Name: "bgp"},
				{Name: "rib"},
				{Name: "attr-sets"},
				{Name: "attr-set", Key: map[string]string{"index": "1"}},
				{Name: "index"},
			},
		},
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{1}},
					Path: &gpb.Path{
						Elem: []*gpb.PathElem{
							{Name: "network-instances"},
							{Name: "network-instance", Key: map[string]string{"name": "foo"}},
							{Name: "protocols"},
							{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
							{Name: "bgp"},
							{Name: "rib"},
							{Name: "attr-sets"},
							{Name: "attr-set", Key: map[string]string{"index": "1"}},
							{Name: "index"},
						},
					},
				}},
			}}},
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
		},
	}, {
		name: "bgp rib streaming",
		mode: gpb.SubscriptionList_STREAM,
		learnedInfo: []*telemetry.Device{
			devWithAttr1,
			devWithAttr2,
		},
		path: &gpb.Path{
			Origin: "openconfig",
			Elem: []*gpb.PathElem{
				{Name: "network-instances"},
				{Name: "network-instance", Key: map[string]string{"name": "foo"}},
				{Name: "protocols"},
				{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
				{Name: "bgp"},
				{Name: "rib"},
				{Name: "attr-sets"},
				{Name: "attr-set", Key: map[string]string{"index": "*"}},
				{Name: "state"},
				{Name: "index"},
			},
		},
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{1}},
					Path: &gpb.Path{
						Elem: []*gpb.PathElem{
							{Name: "network-instances"},
							{Name: "network-instance", Key: map[string]string{"name": "foo"}},
							{Name: "protocols"},
							{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
							{Name: "bgp"},
							{Name: "rib"},
							{Name: "attr-sets"},
							{Name: "attr-set", Key: map[string]string{"index": "1"}},
							{Name: "state"},
							{Name: "index"},
						},
					},
				}},
			}}},
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{2}},
					Path: &gpb.Path{
						Elem: []*gpb.PathElem{
							{Name: "network-instances"},
							{Name: "network-instance", Key: map[string]string{"name": "foo"}},
							{Name: "protocols"},
							{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
							{Name: "bgp"},
							{Name: "rib"},
							{Name: "attr-sets"},
							{Name: "attr-set", Key: map[string]string{"index": "2"}},
							{Name: "state"},
							{Name: "index"},
						},
					},
				}},
			}}},
		},
	}, {
		name: "bgp rib streaming with no data",
		mode: gpb.SubscriptionList_STREAM,
		learnedInfo: []*telemetry.Device{
			nil,
			nil,
			devWithAttr1,
		},
		path: &gpb.Path{
			Origin: "openconfig",
			Elem: []*gpb.PathElem{
				{Name: "network-instances"},
				{Name: "network-instance", Key: map[string]string{"name": "foo"}},
				{Name: "protocols"},
				{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
				{Name: "bgp"},
				{Name: "rib"},
				{Name: "attr-sets"},
				{Name: "attr-set", Key: map[string]string{"index": "*"}},
				{Name: "state"},
				{Name: "index"},
			},
		},
		want: []*gpb.SubscribeResponse{
			{Response: &gpb.SubscribeResponse_SyncResponse{true}},
			{Response: &gpb.SubscribeResponse_Update{&gpb.Notification{
				Prefix: &gpb.Path{Origin: "openconfig", Target: "fakeIxia"},
				Update: []*gpb.Update{{
					Val: &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{1}},
					Path: &gpb.Path{
						Elem: []*gpb.PathElem{
							{Name: "network-instances"},
							{Name: "network-instance", Key: map[string]string{"name": "foo"}},
							{Name: "protocols"},
							{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
							{Name: "bgp"},
							{Name: "rib"},
							{Name: "attr-sets"},
							{Name: "attr-set", Key: map[string]string{"index": "1"}},
							{Name: "state"},
							{Name: "index"},
						},
					},
				}},
			}}},
		},
	}}

	/*
	 * TODO: Add test cases for the flows root. Unfortunately today that
	 * requires dependency on the Ondatra Telemetry API generation, which could
	 * create a circulate dependency back on this module. If and when the
	 * Telemetry API is an independent library, add these test cases.
	 */

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client, err := NewClient(ctx, "fakeIxia", nil, nil)
			if err != nil {
				t.Fatalf("NewClient() got unexpected error: %v", err)
			}

			prefixToReader[bgpRIBPath] = &prefixReader{read: func(_ context.Context, _ *Client, p *gpb.Path) ([]*gpb.Notification, error) {
				if len(tt.learnedInfo) == 0 {
					return nil, nil
				}
				defer func() { tt.learnedInfo = tt.learnedInfo[1:] }()
				if tt.learnedInfo[0] == nil {
					return nil, nil
				}
				ns, err := ygot.TogNMINotifications(tt.learnedInfo[0], time.Now().UnixNano(), ygot.GNMINotificationsConfig{UsePathElem: true})
				if err != nil {
					t.Fatalf("failed to create gNMI notifications: %v", err)
				}
				return ns, nil
			}}

			sub, err := client.Subscribe(ctx)
			if err != nil {
				t.Fatalf("Subscribe() got unexpected error: %v", err)
			}
			stubStats = tt.stats
			req := &gpb.SubscribeRequest{Request: &gpb.SubscribeRequest_Subscribe{&gpb.SubscriptionList{
				Mode:         tt.mode,
				Subscription: []*gpb.Subscription{{Path: tt.path}}},
			}}
			if err := sub.Send(req); err != nil {
				t.Fatalf("Send(%v) got unexpected error: %v", req, err)
			}
			// Sleep to ensure that notifications are received in the correct order.
			time.Sleep(100 * time.Millisecond)
			for _, want := range tt.want {
				// Flush the freshness cache to force the data to be repopulated.
				client.fresh.Flush()
				got, err := sub.Recv()
				if err != nil {
					t.Fatalf("Recv() got unexpected error: %v", err)
				}
				// Zero out the timestamps when comparing.
				if up, ok := got.GetResponse().(*gpb.SubscribeResponse_Update); ok {
					up.Update.Timestamp = 0
				}
				if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
					t.Errorf("Recv() got unexpected response diff (-want,+got): %s", diff)
				} else {
					t.Logf("Recv() correctly got %v", got)
				}
			}
			if len(stubStats) > 0 {
				t.Errorf("Unused stub statistics: %v", stubStats)
			}
			if len(tt.learnedInfo) > 0 {
				t.Errorf("Unused stub statistics: %v", stubStats)
			}
		})
	}
}

type fakeCfgClient struct {
	sess      session
	cfg       *ixconfig.Ixnetwork
	xpathToID map[string]string
	updateErr error
}

func (f *fakeCfgClient) Session() session {
	return f.sess
}

func (f *fakeCfgClient) NodeID(node ixconfig.IxiaCfgNode) (string, error) {
	xp := node.XPath()
	if xp == nil {
		return "", fmt.Errorf("no stub XPath for node of type %T", node)
	}
	id, ok := f.xpathToID[xp.String()]
	if !ok {
		return "", fmt.Errorf("no stub ID for node at %q", xp)
	}
	return id, nil
}

func (f *fakeCfgClient) LastImportedConfig() *ixconfig.Ixnetwork {
	return f.cfg
}

func (f *fakeCfgClient) UpdateIDs(context.Context, *ixconfig.Ixnetwork, ...ixconfig.IxiaCfgNode) error {
	return f.updateErr
}

type fakeSession struct {
	getRsps  map[string]string
	getErrs  map[string]error
	postRsps map[string]string
	postErrs map[string]error
}

func (f *fakeSession) Get(_ context.Context, p string, v any) error {
	if f.getRsps[p] != "" {
		json.Unmarshal([]byte(f.getRsps[p]), v)
	}
	return f.getErrs[p]
}

func (f *fakeSession) Post(_ context.Context, p string, _, v any) error {
	if f.postRsps[p] != "" {
		json.Unmarshal([]byte(f.postRsps[p]), v)
	}
	return f.postErrs[p]
}

func TestBGP4RIBFromIxia(t *testing.T) {
	bgp4XP := parseXPath(t, "/xpath/to/bgpv4")
	inCfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Name: ixconfig.String("Device Group on eth0"),
				Ethernet: []*ixconfig.TopologyEthernet{{
					Ipv4: []*ixconfig.TopologyIpv4{{
						BgpIpv4Peer: []*ixconfig.TopologyBgpIpv4Peer{{
							DutIp: ixconfig.MultivalueStr("1.2.3.4"),
							Xpath: bgp4XP,
						}},
					}},
				}},
			}},
		}},
	}

	tests := []struct {
		desc      string
		intf      string
		cfg       *ixconfig.Ixnetwork
		postErr   map[string]error
		getErr    map[string]error
		getRsps   map[string]string
		updateErr error
		want      map[string]*table
		wantErr   string
	}{{
		desc:    "get config error",
		wantErr: "no IxNetwork config found",
	}, {
		desc:      "update ID error",
		cfg:       inCfg,
		updateErr: errors.New("fake"),
		wantErr:   "failed to update IDs",
	}, {
		desc: "run op error",
		intf: "eth0",
		cfg:  inCfg,
		postErr: map[string]error{
			"topology/deviceGroup/ethernet/ipv4/bgpIpv4Peer/operations/getAllLearnedInfo": errors.New("run fail"),
		},
		wantErr: "run fail",
	}, {
		desc: "get error",
		intf: "eth0",
		cfg:  inCfg,
		getErr: map[string]error{
			"/api/v1/sessions/0/topology/1/deviceGroup/1/ethernet/1/ipv4/1/bgpIpv4Peer/1/learnedInfo/1/table/1": errors.New("get fail"),
		},
		wantErr: "get fail",
	}, {
		desc: "no data for interface",
		intf: "eth1",
		cfg:  inCfg,
	}, {
		desc: "success ipv4",
		intf: "eth0",
		cfg:  inCfg,
		getRsps: map[string]string{
			"/api/v1/sessions/0/topology/1/deviceGroup/1/ethernet/1/ipv4/1/bgpIpv4Peer/1/learnedInfo/1/table/1": `{"id": 1}`,
		},
		want: map[string]*table{"1.2.3.4": {ID: 1}},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{
				client: &fakeCfgClient{
					sess: &fakeSession{
						postErrs: test.postErr,
						getErrs:  test.getErr,
						getRsps:  test.getRsps,
					},
					cfg:       test.cfg,
					updateErr: test.updateErr,
					xpathToID: map[string]string{
						bgp4XP.String(): "/api/v1/sessions/0/topology/1/deviceGroup/1/ethernet/1/ipv4/1/bgpIpv4Peer/1",
					},
				},
			}
			got, err := c.bgp4RIBFromIxia(context.Background(), test.intf)
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("bgp4RIBFromIxia got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(test.want, got); d != "" {
				t.Errorf("bgp4RIBFromIxia got unexpected diff (-want +got)\n%s", d)
			}
		})
	}
}

func TestBGP6RIBFromIxia(t *testing.T) {
	bgp6XP := parseXPath(t, "/xpath/to/bgpv6")
	inCfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Name: ixconfig.String("Device Group on eth0"),
				Ethernet: []*ixconfig.TopologyEthernet{{
					Ipv6: []*ixconfig.TopologyIpv6{{
						BgpIpv6Peer: []*ixconfig.TopologyBgpIpv6Peer{{
							DutIp: ixconfig.MultivalueStr("::1"),
							Xpath: bgp6XP,
						}},
					}},
				}},
			}},
		}},
	}

	tests := []struct {
		desc      string
		intf      string
		cfg       *ixconfig.Ixnetwork
		postErr   map[string]error
		getErr    map[string]error
		getRsps   map[string]string
		updateErr error
		want      map[string]*table
		wantErr   string
	}{{
		desc:    "get config error",
		wantErr: "no IxNetwork config found",
	}, {
		desc:      "update ID error",
		cfg:       inCfg,
		updateErr: errors.New("fake"),
		wantErr:   "failed to update IDs",
	}, {
		desc: "run op error",
		intf: "eth0",
		cfg:  inCfg,
		postErr: map[string]error{
			"topology/deviceGroup/ethernet/ipv6/bgpIpv6Peer/operations/getAllLearnedInfo": errors.New("run fail"),
		},
		wantErr: "run fail",
	}, {
		desc: "get error",
		intf: "eth0",
		cfg:  inCfg,
		getErr: map[string]error{
			"/api/v1/sessions/0/topology/1/deviceGroup/1/ethernet/1/ipv6/1/bgpIpv6Peer/1/learnedInfo/1/table/1": errors.New("get fail"),
		},
		wantErr: "get fail",
	}, {
		desc: "no data for interface",
		intf: "eth1",
		cfg:  inCfg,
	}, {
		desc: "success ipv6",
		intf: "eth0",
		getRsps: map[string]string{
			"/api/v1/sessions/0/topology/1/deviceGroup/1/ethernet/1/ipv6/1/bgpIpv6Peer/1/learnedInfo/1/table/1": `{"id": 1}`,
		},
		want: map[string]*table{"::1": {ID: 1}},
		cfg:  inCfg,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{
				client: &fakeCfgClient{
					sess: &fakeSession{
						postErrs: test.postErr,
						getErrs:  test.getErr,
						getRsps:  test.getRsps,
					},
					cfg:       test.cfg,
					updateErr: test.updateErr,
					xpathToID: map[string]string{
						bgp6XP.String(): "/api/v1/sessions/0/topology/1/deviceGroup/1/ethernet/1/ipv6/1/bgpIpv6Peer/1",
					},
				},
			}
			got, err := c.bgp6RIBFromIxia(context.Background(), test.intf)
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("bgp6RIBFromIxia got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(test.want, got); d != "" {
				t.Errorf("bgp6RIBFromIxia got unexpected diff (-want +got)\n%s", d)
			}
		})
	}
}

func parseXPath(t *testing.T, str string) *ixconfig.XPath {
	xp, err := ixconfig.ParseXPath(str)
	if err != nil {
		t.Fatalf("Error parsing XPath %q: %v", str, err)
	}
	return xp
}

func createSampleDev(t testing.TB, deletes ...*gpb.Path) *gpb.Notification {
	dev := &telemetry.Device{}
	rib := dev.GetOrCreateNetworkInstance("foo").
		GetOrCreateProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "1").
		GetOrCreateBgp().GetOrCreateRib()

	route4 := rib.GetOrCreateAfiSafi(telemetry.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST).
		GetOrCreateIpv4Unicast().GetOrCreateNeighbor("1.2.3.4").
		GetOrCreateAdjRibInPre().GetOrCreateRoute("192.168.1.1/0", 0)
	route4.AttrIndex = ygot.Uint64(0)
	route4.CommunityIndex = ygot.Uint64(0)
	rib.GetOrCreateCommunity(0)
	attr0 := rib.GetOrCreateAttrSet(0)
	attr0.Origin = telemetry.BgpTypes_BgpOriginAttrType_IGP
	attr0.NextHop = ygot.String("")
	attr0.Aigp = ygot.Uint64(0)
	attr0.LocalPref = ygot.Uint32(0)
	attr0.Med = ygot.Uint32(0)

	route6 := rib.GetOrCreateAfiSafi(telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST).
		GetOrCreateIpv6Unicast().GetOrCreateNeighbor("1:2:3:4:5:6:7:8").
		GetOrCreateAdjRibInPre().GetOrCreateRoute("::1/0", 0)
	route6.AttrIndex = ygot.Uint64(1)
	route6.CommunityIndex = ygot.Uint64(1)
	rib.GetOrCreateCommunity(1)
	attr1 := rib.GetOrCreateAttrSet(1)
	attr1.Origin = telemetry.BgpTypes_BgpOriginAttrType_EGP
	attr1.NextHop = ygot.String("")
	attr1.Aigp = ygot.Uint64(0)
	attr1.LocalPref = ygot.Uint32(0)
	attr1.Med = ygot.Uint32(0)

	ns, err := ygot.TogNMINotifications(dev, 0, ygot.GNMINotificationsConfig{UsePathElem: true})
	if err != nil {
		t.Fatal("error creating notifications")
	}
	ns[0].Delete = deletes
	return ns[0]
}

func TestPathToBGPRIB(t *testing.T) {
	inPath := &gpb.Path{
		Elem: []*gpb.PathElem{
			{Name: "network-instances"},
			{Name: "network-instance", Key: map[string]string{"name": "foo"}},
			{Name: "protocols"},
			{Name: "protocol", Key: map[string]string{"identifier": "BGP", "name": "1"}},
			{Name: "bgp"},
			{Name: "rib"},
		},
	}

	tests := []struct {
		desc             string
		rib4, rib6       map[string]*table
		rib4Err, rib6Err error
		initCache        map[string]any
		want             *gpb.Notification
		wantErr          string
	}{{
		desc:    "get rib v4 error",
		rib4Err: errors.New("foo"),
		wantErr: "failed to read rib v4",
	}, {
		desc:    "get rib v6 error",
		rib6Err: errors.New("foo"),
		wantErr: "failed to read rib v6",
	}, {
		desc: "unmarshal v4 error",
		rib4: map[string]*table{"1.2.3.4": {
			Columns: []string{"Prefix Length"},
			Values:  [][]string{{"foo"}},
		}},
		wantErr: "failed to unmarshal rib v4 table",
	}, {
		desc: "unmarshal v6 error",
		rib6: map[string]*table{"1:2:3:4:5:6:7:8": {
			Columns: []string{"Prefix Length"},
			Values:  [][]string{{"foo"}},
		}},
		wantErr: "failed to unmarshal rib v6 table",
	}, {
		desc: "success uncached",
		rib4: map[string]*table{"1.2.3.4": {
			Columns: []string{"IPv4 Prefix", "Origin"},
			Values:  [][]string{{"192.168.1.1", "IGP"}},
		}},
		rib6: map[string]*table{"1:2:3:4:5:6:7:8": {
			Columns: []string{"IPv6 Prefix", "Origin"},
			Values:  [][]string{{"::1", "EGP"}},
		}},
		want: createSampleDev(t),
	}, {
		desc: "success cached",
		rib4: map[string]*table{"1.2.3.4": {
			Columns: []string{"IPv4 Prefix", "Origin"},
			Values:  [][]string{{"192.168.1.1", "IGP"}},
		}},
		rib6: map[string]*table{"1:2:3:4:5:6:7:8": {
			Columns: []string{"IPv6 Prefix", "Origin"},
			Values:  [][]string{{"::1", "EGP"}},
		}},
		want: createSampleDev(t,
			&gpb.Path{Elem: []*gpb.PathElem{{Name: "interfaces"}, {Name: "interface", Key: map[string]string{"name": "fake"}}, {Name: "state"}, {Name: "name"}}},
			&gpb.Path{Elem: []*gpb.PathElem{{Name: "interfaces"}, {Name: "interface", Key: map[string]string{"name": "fake"}}, {Name: "name"}}},
		),
		// This is a contrived example where there is garbage data in the cache, used to verify delete notifications are created.
		// In normal usage, only BGP RIB info could be in the cache.
		initCache: map[string]any{
			"old " + bgpRIBPath + ",foo": &telemetry.Device{
				Interface: map[string]*telemetry.Interface{
					"fake": &telemetry.Interface{
						Name: ygot.String("fake"),
					},
				},
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{
				fresh: cache.New(30*time.Second, cache.NoExpiration),
			}
			for k, v := range test.initCache {
				c.fresh.Set(k, v, -1)
			}
			bgp4RIBFromIxiaFn = func(*Client, context.Context, string) (map[string]*table, error) {
				return test.rib4, test.rib4Err
			}
			bgp6RIBFromIxiaFn = func(*Client, context.Context, string) (map[string]*table, error) {
				return test.rib6, test.rib6Err
			}

			got, gotErr := c.pathToBGPRIB(context.Background(), inPath)
			if diff := errdiff.Substring(gotErr, test.wantErr); diff != "" {
				t.Errorf("pathToBGPRIB() got unexpected error diff\n%s", diff)
			}
			if gotErr != nil {
				return
			}
			if got != nil {
				got.Timestamp = 0
			}
			if diff := cmp.Diff(test.want, got, protocmp.Transform(), protocmp.SortRepeatedFields(&gpb.Notification{}, "update")); diff != "" {
				t.Errorf("pathToBGPRIB() got unexpected response diff (-want,+got)\n%s", diff)
			}
		})
	}
}

func TestRSVPTEFromIxia(t *testing.T) {
	const lspID = "/api/v1/sessions/0/topology/1/deviceGroup/1/networkGroup/1/deviceGroup/1/ipv4Loopback/1/rsvpteLsps/1/rsvpP2PIngressLsps"
	lspXP := parseXPath(t, "/fake/xpath/lsp")
	lspCfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Name: ygot.String("Device Group on eth0"),
				NetworkGroup: []*ixconfig.TopologyNetworkGroup{{
					DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
						Ipv4Loopback: []*ixconfig.TopologyIpv4Loopback{{
							RsvpteLsps: []*ixconfig.TopologyRsvpteLsps{{
								Name: ixconfig.String("lsp0"),
								RsvpP2PIngressLsps: &ixconfig.TopologyRsvpP2PIngressLsps{
									Xpath:  lspXP,
									Active: ixconfig.MultivalueTrue(),
								},
							}},
						}},
					}},
				}},
			}},
		}},
	}

	tests := []struct {
		desc         string
		intf         string
		lspRsp       string
		lspErr       error
		mvRsp        string
		mvErr        error
		updateIDsErr error
		initCache    map[string]*cachedNodes
		want         map[string][]*ingressLSP
		wantErr      string
	}{{
		desc:         "update IDs error",
		intf:         "eth0",
		updateIDsErr: errors.New("some error"),
		wantErr:      "failed to update IDs",
	}, {
		desc:    "lsp lookup error",
		intf:    "eth0",
		lspErr:  errors.New("some error"),
		wantErr: "failed to fetch ingress LSPs config",
	}, {
		desc:    "multivalue lookup error",
		intf:    "eth0",
		lspRsp:  `{"state": ["up", "notStarted"], "sourceIP": "/api/v1/sessions/0/multivalue/1", "destIP": "/api/v1/sessions/0/multivalue/2"}`,
		mvErr:   errors.New("some error"),
		wantErr: "failed to fetch source IPs for LSP",
	}, {
		desc: "no data for interface",
		intf: "eth1",
	}, {
		desc:   "success",
		intf:   "eth0",
		lspRsp: `{"state": ["up", "notStarted"], "sourceIP": "/api/v1/sessions/0/multivalue/1", "destIP": "/api/v1/sessions/0/multivalue/2"}`,
		mvRsp:  `["1.1.1.1", "2.2.2.2"]`,
		want: map[string][]*ingressLSP{
			"lsp0": []*ingressLSP{{
				up:  true,
				src: "1.1.1.1",
				dst: "1.1.1.1",
			}, {

				up:  false,
				src: "2.2.2.2",
				dst: "2.2.2.2",
			}},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{
				client: &fakeCfgClient{
					sess: &fakeSession{
						getErrs: map[string]error{
							lspID: test.lspErr,
						},
						getRsps: map[string]string{
							lspID: test.lspRsp,
						},
						postErrs: map[string]error{
							"multivalue/operations/getvalues": test.mvErr,
						},
						postRsps: map[string]string{
							"multivalue/operations/getvalues": test.mvRsp,
						},
					},
					cfg:       lspCfg,
					updateErr: test.updateIDsErr,
					xpathToID: map[string]string{
						lspXP.String(): "/api/v1/sessions/0/topology/1/deviceGroup/1/networkGroup/1/deviceGroup/1/ipv4Loopback/1/rsvpteLsps/1/rsvpP2PIngressLsps",
					},
				},
				intfCache: test.initCache,
			}
			got, err := c.rsvpTEFromIxia(context.Background(), test.intf)
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("rsvpTEFromIxia() got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(test.want, got, cmp.AllowUnexported(ingressLSP{})); d != "" {
				t.Errorf("rsvpTEFromIxia() got unexpected diff (-want +got)\n%s", d)
			}
		})
	}
}

func TestPathToRSVPTE(t *testing.T) {
	oldLSPsFn := rsvpTEFromIxiaFn
	defer func() { rsvpTEFromIxiaFn = oldLSPsFn }()

	inPath := &gpb.Path{Elem: []*gpb.PathElem{
		{Name: "network-instances"},
		{Name: "network-instance", Key: map[string]string{"name": "foo"}},
		{Name: "mpls"},
		{Name: "signaling-protocols"},
		{Name: "rsvp-te"},
	}}

	tests := []struct {
		desc      string
		initCache map[string]any
		lsps      map[string][]*ingressLSP
		lspsErr   error
		want      *gpb.Notification
		wantErr   string
	}{{
		desc: "data is fresh",
		initCache: map[string]any{
			rsvpTEPath + ",foo": true,
		},
	}, {
		desc:    "failed to read LSP data from ATE",
		lspsErr: errors.New("some error"),
		wantErr: "failed to read LSPs from Ixia",
	}, {
		desc: "successful update",
		lsps: map[string][]*ingressLSP{
			"lsp0": []*ingressLSP{{
				up:  true,
				src: "1.1.1.1",
				dst: "2.2.2.2",
			}},
		},
		initCache: map[string]any{
			rsvpTEPath + ",bar": true,
			"old " + rsvpTEPath + ",foo": &telemetry.Device{
				NetworkInstance: map[string]*telemetry.NetworkInstance{
					"fake": &telemetry.NetworkInstance{
						Name: ygot.String("fake"),
					},
				},
			},
		},
		want: func(t *testing.T) *gpb.Notification {
			dev := &telemetry.Device{}
			lsp := dev.GetOrCreateNetworkInstance("foo").GetOrCreateMpls().GetOrCreateSignalingProtocols().GetOrCreateRsvpTe().GetOrCreateSession(2765635037960339456)
			lsp.SessionName = ygot.String("lsp0 0")
			lsp.Type = telemetry.MplsTypes_LSP_ROLE_INGRESS
			lsp.SourceAddress = ygot.String("1.1.1.1")
			lsp.DestinationAddress = ygot.String("2.2.2.2")
			lsp.Status = telemetry.Session_Status_UP
			ns, err := ygot.TogNMINotifications(dev, 0, ygot.GNMINotificationsConfig{UsePathElem: true})
			if err != nil {
				t.Fatal("error creating notifications")
			}
			ns[0].Delete = []*gpb.Path{{
				Elem: []*gpb.PathElem{{Name: "network-instances"}, {Name: "network-instance", Key: map[string]string{"name": "fake"}}, {Name: "state"}, {Name: "name"}},
			}, {
				Elem: []*gpb.PathElem{{Name: "network-instances"}, {Name: "network-instance", Key: map[string]string{"name": "fake"}}, {Name: "name"}},
			}}
			return ns[0]
		}(t),
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{
				fresh: cache.New(30*time.Second, cache.NoExpiration),
			}
			for k, v := range test.initCache {
				c.fresh.Set(k, v, -1)
			}
			rsvpTEFromIxiaFn = func(*Client, context.Context, string) (map[string][]*ingressLSP, error) {
				return test.lsps, test.lspsErr
			}

			got, gotErr := c.pathToRSVPTE(context.Background(), inPath)
			if diff := errdiff.Substring(gotErr, test.wantErr); diff != "" {
				t.Errorf("pathToRSVPTE() got unexpected error diff\n%s", diff)
			}
			if gotErr != nil {
				return
			}
			if got != nil {
				got.Timestamp = 0
			}
			if diff := cmp.Diff(test.want, got, protocmp.Transform(), protocmp.SortRepeatedFields(&gpb.Notification{}, "update")); diff != "" {
				t.Errorf("pathToRSVPTE() got unexpected response diff (-want,+got)\n%s", diff)
			}
		})
	}
}
