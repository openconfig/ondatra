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

// Package ixgnmi provides a GNMI client for Ixia.
package ixgnmi

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	gcache "github.com/openconfig/gnmi/cache"
	"github.com/openconfig/gnmi/subscribe"
	closer "github.com/openconfig/gocloser"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	bgpRIBPath   = "/network-instances/network-instance/protocols/protocol/bgp/rib"
	isisLSDBPath = "/network-instances/network-instance/protocols/protocol/isis/levels/level/link-state-database"
	rsvpTEPath   = "/network-instances/network-instance/mpls/signaling-protocols/rsvp-te"

	portStatsCaption    = "Port Statistics"
	portCPUStatsCaption = "Port CPU Statistics"
	flowStatsCaption    = "Flow Statistics"
)

var (
	prefixToReader = map[string]*prefixReader{
		"/components": statViewReader(portCPUStatsCaption),
		"/flows":      statViewReader(ixweb.TrafficItemStatsCaption, flowStatsCaption, ixweb.EgressStatsCaption),
		"/interfaces": statViewReader(portStatsCaption),
		bgpRIBPath:    protocolReader(bgpRIBFromIxia),
		isisLSDBPath:  protocolReader(isisLSDBFromIxia),
		rsvpTEPath:    protocolReader(rsvpTEFromIxia),
	}

	// To be stubbed out by tests.
	readStatsFn = func(ctx context.Context, c *Client, cacheKey string, captions []string) (ygot.GoStruct, error) {
		return c.readStats(ctx, cacheKey, captions)
	}
)

type prefixReader struct {
	mu   sync.Mutex
	read func(context.Context, *Client, *gpb.Path) ([]*gpb.Notification, error)
}

// StatReader reads Ixia stats for a specified set of views.
type StatReader func(context.Context, []string) (*Stats, error)

// Stats are Ixia statistics collected from ixia.
type Stats struct {
	Tables            map[string]ixweb.StatTable
	IngressTrackFlows []string
	EgressTrackFlows  []string
}

type cfgClient interface {
	Session() session
	NodeID(ixconfig.IxiaCfgNode) (string, error)
	LastImportedConfig() *ixconfig.Ixnetwork
	UpdateIDs(context.Context, *ixconfig.Ixnetwork, ...ixconfig.IxiaCfgNode) error
}

type session interface {
	Get(context.Context, string, any) error
	Post(context.Context, string, any, any) error
}

type clientWrapper struct {
	*ixconfig.Client
}

func (cw *clientWrapper) Session() session {
	return cw.Client.Session()
}

func statViewReader(captions ...string) *prefixReader {
	return &prefixReader{read: func(ctx context.Context, c *Client, p *gpb.Path) ([]*gpb.Notification, error) {
		ys, err := readStatsFn(ctx, c, p.GetElem()[0].GetName(), captions)
		if err != nil {
			return nil, err
		}
		if ys == nil {
			return nil, nil
		}
		ns, err := ygot.TogNMINotifications(
			ys,
			time.Now().UnixNano(),
			ygot.GNMINotificationsConfig{UsePathElem: true},
		)
		if err != nil {
			return nil, fmt.Errorf("cannot render telemetry Notifications: %w", err)
		}
		return ns, nil
	}}
}

func protocolReader(fetch func(context.Context, cfgClient, *oc.NetworkInstance, *cachedNodes) error) *prefixReader {
	return &prefixReader{read: func(ctx context.Context, c *Client, p *gpb.Path) ([]*gpb.Notification, error) {
		intf := p.GetElem()[1].GetKey()["name"]
		key, prevKey := cacheKeys(fetch, intf)
		if _, isFresh := c.fresh.Get(key); isFresh {
			return nil, nil
		}

		dev := new(oc.Root)
		nodes, err := c.fetchCachedNodes(ctx, intf)
		if err != nil {
			return nil, err
		}
		err = fetch(ctx, c.client, dev.GetOrCreateNetworkInstance(intf), nodes)
		if err != nil {
			return nil, err
		}

		var prevDev *oc.Root
		if prev, ok := c.fresh.Get(prevKey); ok {
			prevDev = prev.(*oc.Root)
		}
		notif, err := ygot.Diff(prevDev, dev)
		if err != nil {
			return nil, fmt.Errorf("failed to render notifications: %w", err)
		}
		notif.Timestamp = time.Now().UnixNano()
		c.fresh.Set(prevKey, dev, cache.NoExpiration)
		c.fresh.SetDefault(key, true)
		return []*gpb.Notification{notif}, err
	}}
}

func cacheKeys(fn any, intf string) (string, string) {
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	key := fmt.Sprintf("%s %s", fnName, intf)
	return key, "prev " + key
}

// NewClient creates a new Ixia gNMI client.
func NewClient(ctx context.Context, name string, reader StatReader, client *ixconfig.Client, opts ...grpc.DialOption) (_ *Client, rerr error) {
	const statExpiration = 30 * time.Second
	gc := gcache.New([]string{name})
	srv := grpc.NewServer()
	subSrv, err := subscribe.NewServer(gc)
	if err != nil {
		return nil, fmt.Errorf("cannot create new subscribe server: %w", err)
	}
	gc.SetClient(subSrv.Update)
	gpb.RegisterGNMIServer(srv, subSrv)
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, fmt.Errorf("cannot listen on an available port: %w", err)
	}
	defer closer.CloseOnErr(&rerr, lis.Close, "error closing listener")
	go srv.Serve(lis)
	addr := lis.Addr().String()
	opts = append(opts, grpc.WithTransportCredentials(local.NewCredentials()))
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("DialContext(%s, nil, %v): %w", addr, opts, err)
	}
	defer closer.CloseOnErr(&rerr, conn.Close, "error closing gRPC connection")
	return &Client{
		ignmi:  gpb.NewGNMIClient(conn),
		target: gc.GetTarget(name),
		reader: reader,
		client: &clientWrapper{client},
		fresh:  cache.New(statExpiration, cache.NoExpiration),
	}, nil
}

// unexported interface that can be embedded without exporting
type ignmi interface {
	gpb.GNMIClient
}

// Client is an Ixia GNMI client.
type Client struct {
	ignmi
	target *gcache.Target
	reader StatReader
	client cfgClient
	// mu guards access to intfCache.
	mu        sync.Mutex
	intfCache map[string]*cachedNodes
	// fresh tracks whether stats in the Cache are fresh.
	// Keys are the prefix strings enumerated in prefixToReader.
	// A key's presence in the map indicates that the data at that path is fresh.
	fresh *cache.Cache
}

type cachedNodes struct {
	bgp4Peers []*ixconfig.TopologyBgpIpv4Peer
	bgp6Peers []*ixconfig.TopologyBgpIpv6Peer
	isisL3s   []*ixconfig.TopologyIsisL3
	rsvpLSPs  []*ixconfig.TopologyRsvpP2PIngressLsps
}

// Flush flushes the GNMI data for the Ixia.
func (c *Client) Flush() {
	c.fresh.Flush()
	c.target.Reset()
	c.target.Connect()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.intfCache = nil
}

func (c *Client) Subscribe(ctx context.Context, opts ...grpc.CallOption) (gpb.GNMI_SubscribeClient, error) {
	gsub, err := c.ignmi.Subscribe(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &subClient{
		GNMI_SubscribeClient: gsub,
		parent:               c,
	}, nil
}

func (c *Client) publishRoot(ctx context.Context, prefix string, path *gpb.Path) (bool, error) {
	reader := prefixToReader[prefix]
	if reader == nil {
		return false, fmt.Errorf("no reader for prefix: %q", prefix)
	}

	// Ensure we do not have concurrent writes to the same root.
	// This effectively queues up concurrent gnmi requests for the same root,
	// so that one sees the cached result of another for improved performance.
	reader.mu.Lock()
	defer reader.mu.Unlock()

	c.fresh.DeleteExpired()
	ns, err := reader.read(ctx, c, path)
	if err != nil {
		return false, err
	}
	hasData := len(ns) > 0
	for _, n := range ns {
		n.Prefix = &gpb.Path{
			Target: c.target.Name(),
			Origin: "openconfig",
		}
		log.V(2).Infof("Pushing notif to cache: %+v", n)
		if err := c.target.GnmiUpdate(n); err != nil {
			return false, fmt.Errorf("failed to update gNMI cache for target %s: %w", c.target.Name(), err)
		}
	}
	return hasData, nil
}

func (c *Client) readStats(ctx context.Context, cacheKey string, captions []string) (ygot.GoStruct, error) {
	if _, ok := c.fresh.Get(cacheKey); ok {
		return nil, nil
	}
	s, err := c.reader(ctx, captions)
	if err != nil {
		return nil, fmt.Errorf("error retrieving statistics for views %v: %w", captions, err)
	}
	rs, err := translate(s)
	if err != nil {
		return nil, fmt.Errorf("cannot translate statistics: %w", err)
	}
	c.fresh.SetDefault(cacheKey, true)
	return rs, nil
}

func (c *Client) fetchCachedNodes(ctx context.Context, intf string) (*cachedNodes, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.intfCache == nil {
		if err := c.updateIntfCache(ctx); err != nil {
			return nil, err
		}
	}
	nodes, ok := c.intfCache[intf]
	if !ok {
		nodes = new(cachedNodes)
	}
	return nodes, nil
}

func (c *Client) updateIntfCache(ctx context.Context) error {
	c.intfCache = make(map[string]*cachedNodes)
	cfg := c.client.LastImportedConfig()
	if cfg == nil {
		return errors.New("no IxNetwork config found")
	}

	var allNodes []ixconfig.IxiaCfgNode
	for _, topo := range cfg.Topology {
		for _, dg := range topo.DeviceGroup {
			var intf string
			cnt, err := fmt.Sscanf(*dg.Name, "Device Group on %s", &intf)
			if err != nil || cnt != 1 {
				continue
			}
			nodes := new(cachedNodes)
			c.intfCache[intf] = nodes
			if len(dg.Ethernet) > 0 {
				eth := dg.Ethernet[0]
				if len(eth.Ipv4) > 0 {
					for _, peer := range eth.Ipv4[0].BgpIpv4Peer {
						if isActive(peer.Active) {
							nodes.bgp4Peers = append(nodes.bgp4Peers, peer)
							allNodes = append(allNodes, peer)
						}
					}
				}
				if len(eth.Ipv6) > 0 {
					for _, peer := range eth.Ipv6[0].BgpIpv6Peer {
						if isActive(peer.Active) {
							nodes.bgp6Peers = append(nodes.bgp6Peers, peer)
							allNodes = append(allNodes, peer)
						}
					}
				}
				for _, isis := range eth.IsisL3 {
					if isActive(isis.Active) {
						nodes.isisL3s = append(nodes.isisL3s, isis)
						allNodes = append(allNodes, isis)
					}
				}
			}
			for _, dg := range allDevGroups(dg) {
				for _, lb := range dg.Ipv4Loopback {
					for _, lsp := range lb.RsvpteLsps {
						if ilsp := lsp.RsvpP2PIngressLsps; ilsp != nil && isActive(ilsp.Active) {
							nodes.rsvpLSPs = append(nodes.rsvpLSPs, ilsp)
							allNodes = append(allNodes, ilsp)
						}
					}
				}
			}
		}
	}
	if err := c.client.UpdateIDs(ctx, cfg, allNodes...); err != nil {
		return fmt.Errorf("failed to update IDs for nodes: %v", allNodes)
	}
	return nil
}

func isActive(mv *ixconfig.Multivalue) bool {
	if mv == nil {
		// If there is no value, assume active.
		return true
	}
	return *(mv.SingleValue.Value) == "true"
}

func allDevGroups(dg *ixconfig.TopologyDeviceGroup) []*ixconfig.TopologyDeviceGroup {
	dgs := []*ixconfig.TopologyDeviceGroup{dg}
	for _, ng := range dg.NetworkGroup {
		for _, ndg := range ng.DeviceGroup {
			dgs = append(dgs, allDevGroups(ndg)...)
		}
	}
	for _, dg := range dg.DeviceGroup {
		dgs = append(dgs, allDevGroups(dg)...)
	}
	return dgs
}

type subClient struct {
	gpb.GNMI_SubscribeClient
	parent *Client
	mode   gpb.SubscriptionList_Mode
	roots  []string
	paths  []*gpb.Path
}

func (s *subClient) Send(req *gpb.SubscribeRequest) error {
	sub := req.GetSubscribe()
	s.mode = sub.GetMode()
	prefix := sub.GetPrefix()
	if prefix == nil {
		prefix = &gpb.Path{}
		sub.Prefix = prefix
	}
	prefix.Target = s.parent.target.Name()

	for _, subscript := range sub.GetSubscription() {
		path, err := util.JoinPaths(prefix, subscript.GetPath())
		if err != nil {
			return err
		}
		if path.GetOrigin() != "openconfig" {
			continue
		}
		sp, err := ygot.PathToSchemaPath(path)
		if err != nil {
			return err
		}
		for root := range prefixToReader {
			if strings.HasPrefix(sp, root) {
				s.roots = append(s.roots, root)
				s.paths = append(s.paths, path)
				break
			}
		}

		// gNMI cache will not match entries if the origin is set in the path
		// instead of the prefix, and the subscribe logic complains if it is
		// set in both places. So after confirming the origin is "openconfig",
		// erase the origin and set the prefix origin below to "openconfig."
		subscript.GetPath().Origin = ""
	}
	prefix.Origin = "openconfig"

	// For once subscriptions, need to make sure the data is fully populated
	// before even sending the request; otherwise it might not be found.
	if s.mode == gpb.SubscriptionList_ONCE {
		if _, err := s.publishRoots(); err != nil {
			return err
		}
	}
	return s.GNMI_SubscribeClient.SendMsg(req)
}

func (s *subClient) Recv() (*gpb.SubscribeResponse, error) {
	// For non-once subscriptions, refresh the data on every receive.
	if s.mode != gpb.SubscriptionList_ONCE {
		hasData, err := s.publishRoots()
		if err != nil {
			return nil, err
		}
		if hasData {
			return s.GNMI_SubscribeClient.Recv()
		}
		// If there's no data, keep trying get more in the background.
		tick := time.NewTicker(50 * time.Millisecond)
		done := make(chan struct{})
		defer tick.Stop()
		defer close(done)
		go func() {
			for {
				select {
				case <-done:
					return
				case <-tick.C:
					data, err := s.publishRoots()
					if err != nil {
						log.Errorf("Failed to publish roots: %v", err)
						return
					}
					if data {
						return
					}
				}
			}
		}()

	}
	return s.GNMI_SubscribeClient.Recv()
}

func (s *subClient) publishRoots() (bool, error) {
	var ret bool
	for i, root := range s.roots {
		hasData, err := s.parent.publishRoot(s.Context(), root, s.paths[i])
		ret = hasData || ret
		if err != nil {
			return false, err
		}
	}
	return ret, nil
}
