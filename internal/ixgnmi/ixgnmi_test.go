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
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ygot/uexampleoc"
	"github.com/openconfig/ygot/ygot"
	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/testing/protocmp"

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

	devWithAttr1 := &oc.Root{}
	devWithAttr1.GetOrCreateNetworkInstance("foo").GetOrCreateProtocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "1").
		GetOrCreateBgp().GetOrCreateRib().GetOrCreateAttrSet(1)

	devWithAttr2 := &oc.Root{}
	devWithAttr2.GetOrCreateNetworkInstance("foo").GetOrCreateProtocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "1").
		GetOrCreateBgp().GetOrCreateRib().GetOrCreateAttrSet(2)

	devWithPartial := &oc.Root{}
	devWithPartial.GetOrCreateNetworkInstance("foo")

	tests := []struct {
		name        string
		mode        gpb.SubscriptionList_Mode
		path        *gpb.Path
		stats       []map[string]ygot.GoStruct
		learnedInfo []*oc.Root
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
		learnedInfo: []*oc.Root{
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
		learnedInfo: []*oc.Root{
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
		learnedInfo: []*oc.Root{
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
	getRsps  map[string][]string
	getErrs  map[string]error
	postRsps map[string][]string
	postErrs map[string]error
}

func (f *fakeSession) Get(_ context.Context, p string, v any) error {
	if rsps := f.getRsps[p]; len(rsps) > 0 {
		rsp := rsps[0]
		f.getRsps[p] = rsps[1:]
		return json.Unmarshal([]byte(rsp), v)
	}
	return f.getErrs[p]
}

func (f *fakeSession) Post(_ context.Context, p string, _, v any) error {
	if rsps := f.postRsps[p]; len(rsps) > 0 {
		rsp := rsps[0]
		f.postRsps[p] = rsps[1:]
		return json.Unmarshal([]byte(rsp), v)
	}
	return f.postErrs[p]
}

func TestProtocolReader(t *testing.T) {
	inPath := &gpb.Path{Elem: []*gpb.PathElem{
		{Name: "network-instances"},
		{Name: "network-instance", Key: map[string]string{"name": "foo"}},
	}}
	activeBGP4Peer := &ixconfig.TopologyBgpIpv4Peer{Active: ixconfig.MultivalueTrue()}
	activeBGP6Peer := &ixconfig.TopologyBgpIpv6Peer{Active: ixconfig.MultivalueTrue()}
	activeISIS := &ixconfig.TopologyIsisL3{} // active despite the absence of the Active field
	activeRSVPLSP := &ixconfig.TopologyRsvpP2PIngressLsps{Active: ixconfig.MultivalueTrue()}

	var gotNodes *cachedNodes
	var readFn func(*oc.NetworkInstance) error
	fetch := func(_ context.Context, _ cfgClient, netInst *oc.NetworkInstance, nodes *cachedNodes) error {
		gotNodes = nodes
		return readFn(netInst)
	}
	key, prevKey := cacheKeys(fetch, "foo")
	reader := protocolReader(fetch)

	tests := []struct {
		desc      string
		initCache map[string]any
		cfg       *ixconfig.Ixnetwork
		readFn    func(*oc.NetworkInstance) error
		want      *gpb.Notification
		wantNodes *cachedNodes
		wantErr   string
	}{{
		desc:    "no network config",
		wantErr: "no IxNetwork config",
	}, {
		desc:      "data is fresh",
		initCache: map[string]any{key: true},
	}, {
		desc:    "failed to get ixia data",
		cfg:     &ixconfig.Ixnetwork{},
		readFn:  func(*oc.NetworkInstance) error { return errors.New("ixia error") },
		wantErr: "ixia error",
	}, {
		desc: "success",
		// Contrived example where there is garbage data in the cache.
		// Used to verify delete notifications are created.
		initCache: map[string]any{
			prevKey: &oc.Root{
				Interface: map[string]*oc.Interface{
					"fake": {Name: ygot.String("fake")},
				},
			},
		},
		cfg: &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					Name: ixconfig.String("Device Group on foo"),
					Ethernet: []*ixconfig.TopologyEthernet{{
						Ipv4: []*ixconfig.TopologyIpv4{{
							BgpIpv4Peer: []*ixconfig.TopologyBgpIpv4Peer{
								activeBGP4Peer,
								// Inactive peer that should be excluded.
								{Active: ixconfig.MultivalueFalse()},
							},
						}},
						Ipv6: []*ixconfig.TopologyIpv6{{
							BgpIpv6Peer: []*ixconfig.TopologyBgpIpv6Peer{
								activeBGP6Peer,
								// Inactive peer that should be excluded.
								{Active: ixconfig.MultivalueFalse()},
							},
						}},
						IsisL3: []*ixconfig.TopologyIsisL3{
							activeISIS,
							// Inactive isis that should be excluded.
							{Active: ixconfig.MultivalueFalse()},
						},
					}},
					NetworkGroup: []*ixconfig.TopologyNetworkGroup{{
						DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
							Ipv4Loopback: []*ixconfig.TopologyIpv4Loopback{{
								RsvpteLsps: []*ixconfig.TopologyRsvpteLsps{
									{RsvpP2PIngressLsps: activeRSVPLSP},
									// Inactive lsp that should be excluded.
									{RsvpP2PIngressLsps: &ixconfig.TopologyRsvpP2PIngressLsps{Active: ixconfig.MultivalueFalse()}},
								},
							}},
						}},
					}},
				}},
			}},
		},
		readFn: func(netInst *oc.NetworkInstance) error {
			netInst.GetOrCreateInterface("bar")
			return nil
		},
		want: &gpb.Notification{
			Update: []*gpb.Update{{
				Path: &gpb.Path{Elem: []*gpb.PathElem{
					{Name: "network-instances"},
					{Name: "network-instance", Key: map[string]string{"name": "foo"}},
					{Name: "interfaces"},
					{Name: "interface", Key: map[string]string{"id": "bar"}},
					{Name: "id"},
				}},
				Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"bar"}},
			}, {
				Path: &gpb.Path{Elem: []*gpb.PathElem{
					{Name: "network-instances"},
					{Name: "network-instance", Key: map[string]string{"name": "foo"}},
					{Name: "interfaces"},
					{Name: "interface", Key: map[string]string{"id": "bar"}},
					{Name: "state"},
					{Name: "id"},
				}},
				Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"bar"}},
			}, {
				Path: &gpb.Path{Elem: []*gpb.PathElem{
					{Name: "network-instances"},
					{Name: "network-instance", Key: map[string]string{"name": "foo"}},
					{Name: "name"},
				}},
				Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"foo"}},
			}, {
				Path: &gpb.Path{Elem: []*gpb.PathElem{
					{Name: "network-instances"},
					{Name: "network-instance", Key: map[string]string{"name": "foo"}},
					{Name: "state"},
					{Name: "name"},
				}},
				Val: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{"foo"}},
			}},
			Delete: []*gpb.Path{{
				Elem: []*gpb.PathElem{
					{Name: "interfaces"},
					{Name: "interface", Key: map[string]string{"name": "fake"}},
					{Name: "state"},
					{Name: "name"},
				}}, {
				Elem: []*gpb.PathElem{
					{Name: "interfaces"},
					{Name: "interface", Key: map[string]string{"name": "fake"}},
					{Name: "name"},
				}},
			},
		},
		wantNodes: &cachedNodes{
			bgp4Peers: []*ixconfig.TopologyBgpIpv4Peer{activeBGP4Peer},
			bgp6Peers: []*ixconfig.TopologyBgpIpv6Peer{activeBGP6Peer},
			isisL3s:   []*ixconfig.TopologyIsisL3{activeISIS},
			rsvpLSPs:  []*ixconfig.TopologyRsvpP2PIngressLsps{activeRSVPLSP},
		},
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{
				fresh:  cache.New(30*time.Second, cache.NoExpiration),
				client: &fakeCfgClient{cfg: test.cfg},
			}
			for k, v := range test.initCache {
				c.fresh.Set(k, v, -1)
			}
			readFn = test.readFn

			gotNotifs, gotErr := reader.read(context.Background(), c, inPath)
			if diff := errdiff.Substring(gotErr, test.wantErr); diff != "" {
				t.Errorf("protocolReader() got unexpected error diff\n%s", diff)
			}
			if gotErr != nil {
				return
			}
			if len(gotNotifs) > 1 {
				t.Fatalf("protocolReader() got %v notifications, want <= 1: %v", len(gotNotifs), gotNotifs)
			}
			var got *gpb.Notification
			if len(gotNotifs) > 0 {
				got = gotNotifs[0]
				got.Timestamp = 0
			}
			if diff := cmp.Diff(test.want, got, protocmp.Transform(), protocmp.SortRepeatedFields(&gpb.Notification{}, "update")); diff != "" {
				t.Errorf("protocolReader() got unexpected response diff (-want,+got)\n%s", diff)
			}
			if diff := cmp.Diff(test.wantNodes, gotNodes, cmp.AllowUnexported(cachedNodes{}, ixconfig.TopologyBgpIpv4Peer{}, ixconfig.TopologyBgpIpv6Peer{}, ixconfig.TopologyIsisL3{})); diff != "" {
				t.Errorf("protocolReader() got unexpected nodes diff (-want,+got)\n%s", diff)
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
