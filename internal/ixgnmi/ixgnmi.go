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
	"golang.org/x/net/context"
	"errors"
	"fmt"
	"net"
	"path"
	"strings"
	"sync"
	"time"

	log "github.com/golang/glog"
	"google.golang.org/grpc/credentials/local"
	"google.golang.org/grpc"
	"github.com/patrickmn/go-cache"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	gcache "github.com/openconfig/gnmi/cache"
	"github.com/openconfig/gnmi/subscribe"
	"github.com/openconfig/gocloser"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ondatra/telemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	ribOCPath  = "/network-instances/network-instance/protocols/protocol/bgp/rib"
	lspsOCPath = "/network-instances/network-instance/mpls/signaling-protocols/rsvp-te"

	portStatsCaption    = "Port Statistics"
	portCPUStatsCaption = "Port CPU Statistics"
	flowStatsCaption    = "Flow Statistics"
)

var (
	prefixToReader = map[string]*prefixReader{
		"/components": statViewReader(portCPUStatsCaption),
		"/flows":      statViewReader(ixweb.TrafficItemStatsCaption, flowStatsCaption, ixweb.EgressStatsCaption),
		"/interfaces": statViewReader(portStatsCaption),
		ribOCPath: &prefixReader{read: func(ctx context.Context, c *Client, p *gpb.Path) ([]*gpb.Notification, error) {
			n, err := c.pathToOCRIB(ctx, p)
			if n != nil {
				return []*gpb.Notification{n}, err
			}
			return nil, err
		}},
		lspsOCPath: &prefixReader{read: func(ctx context.Context, c *Client, _ *gpb.Path) ([]*gpb.Notification, error) {
			n, err := c.pathToOCLSPs(ctx)
			if n != nil {
				return []*gpb.Notification{n}, err
			}
			return nil, err
		}},
	}

	// To be stubbed out by tests.
	readStatsFn = func(ctx context.Context, c *Client, cacheKey string, captions []string) (ygot.GoStruct, error) {
		return c.readStats(ctx, cacheKey, captions)
	}
	ribFromIxiaFn  = (*Client).ribFromIxia
	lspsFromIxiaFn = (*Client).lspsFromIxia
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
	// mu guards access to peerCache.
	mu        sync.Mutex
	peerCache map[string]map[string]ixconfig.IxiaCfgNode
	// fresh tracks whether stats in the Cache are fresh.
	// Keys are the prefix strings enumerated in prefixToReader.
	// A key's presence in the map indicates that the data at that path is fresh.
	fresh *cache.Cache
}

// Flush flushes the GNMI data for the Ixia.
func (c *Client) Flush() {
	c.fresh.Flush()
	c.target.Reset()
	c.target.Connect()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.peerCache = nil
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

func (c *Client) publishRoot(ctx context.Context, root string, path *gpb.Path) (bool, error) {
	reader := prefixToReader[root]
	if reader == nil {
		log.V(2).Info("No reader for path")
		return true, nil
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

func (c *Client) fetchPeerCache(ctx context.Context) (map[string]map[string]ixconfig.IxiaCfgNode, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.peerCache != nil {
		return c.peerCache, nil
	}
	c.peerCache = make(map[string]map[string]ixconfig.IxiaCfgNode)
	cfg := c.client.LastImportedConfig()
	if cfg == nil {
		return nil, errors.New("no IxNetwork config found")
	}

	var allPeers []ixconfig.IxiaCfgNode
	for _, topo := range cfg.Topology {
		for _, dg := range topo.DeviceGroup {
			if len(dg.Ethernet) < 1 {
				continue
			}
			var iface string
			cnt, err := fmt.Sscanf(*dg.Name, "Device Group on %s", &iface)
			if err != nil || cnt != 1 {
				continue
			}
			c.peerCache[iface] = make(map[string]ixconfig.IxiaCfgNode)
			if len(dg.Ethernet[0].Ipv4) > 0 {
				for _, p := range dg.Ethernet[0].Ipv4[0].BgpIpv4Peer {
					c.peerCache[iface][*p.DutIp.SingleValue.Value] = p
					allPeers = append(allPeers, p)
				}
			}
			if len(dg.Ethernet[0].Ipv6) > 0 {
				for _, p := range dg.Ethernet[0].Ipv6[0].BgpIpv6Peer {
					c.peerCache[iface][*p.DutIp.SingleValue.Value] = p
					allPeers = append(allPeers, p)
				}
			}
		}
	}
	if err := c.client.UpdateIDs(ctx, cfg, allPeers...); err != nil {
		return nil, fmt.Errorf("failed to update IDs for bgp peers: %v", allPeers)
	}
	return c.peerCache, nil
}

// ribFromIxia gets the BGP RIB from an IXIA device.
func (c *Client) ribFromIxia(ctx context.Context, pi peerInfo) (*table, error) {
	const (
		bgpV4OpPath = "topology/deviceGroup/ethernet/ipv4/bgpIpv4Peer/operations/getAllLearnedInfo"
		bgpV6OpPath = "topology/deviceGroup/ethernet/ipv6/bgpIpv6Peer/operations/getAllLearnedInfo"
	)
	peerCache, err := c.fetchPeerCache(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update cache: %w", err)
	}
	node, ok := peerCache[pi.intf][pi.neighbor]
	if !ok {
		return nil, fmt.Errorf("no peer %q on interface %q", pi.neighbor, pi.intf)
	}

	opPath := bgpV6OpPath
	if pi.isIPV4 {
		opPath = bgpV4OpPath
	}
	nodeID, err := c.client.NodeID(node)
	if err != nil {
		return nil, err
	}
	if err := c.client.Session().Post(ctx, opPath, ixweb.OpArgs{[]string{nodeID}}, nil); err != nil {
		return nil, fmt.Errorf("failed to run op: %w", err)
	}
	table := &table{}
	if err := c.client.Session().Get(ctx, path.Join(nodeID, "learnedInfo/1/table/1"), table); err != nil {
		return nil, fmt.Errorf("failed to get learned info: %w", err)
	}
	return table, nil
}

type peerInfo struct {
	intf     string
	protocol string
	neighbor string
	isIPV4   bool
}

const (
	peerInfoCacheKey = "bgpLearnedInfoCacheKey"
	oldRibCacheKey   = "bgpRIBCacheKey"
)

func (c *Client) pathToOCRIB(ctx context.Context, p *gpb.Path) (*gpb.Notification, error) {
	const (
		attrSetOCPath      = ribOCPath + "/attr-sets"
		communityOCPath    = ribOCPath + "/communities"
		bgpV4UnicastOCPath = ribOCPath + "/afi-safis/afi-safi/ipv4-unicast/neighbors/neighbor/adj-rib-in-pre/"
		bgpV6UnicastOCPath = ribOCPath + "/afi-safis/afi-safi/ipv6-unicast/neighbors/neighbor/adj-rib-in-pre/"
	)

	schemaPath, err := ygot.PathToSchemaPath(p)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema path: %w", err)
	}

	var cachePI peerInfo
	cached, hasCachedPeer := c.fresh.Get(peerInfoCacheKey)
	if hasCachedPeer {
		cachePI = cached.(peerInfo)
	}

	// Attr set and communities are only valid and updated on calls to Adj RIB paths.
	if strings.HasPrefix(schemaPath, attrSetOCPath) || strings.HasPrefix(schemaPath, communityOCPath) {
		if !hasCachedPeer {
			return nil, errors.New("need to read the attr index or comm index before reading attr sets or communities")
		}
		return nil, nil
	}

	// Ignore unknown paths.
	if !strings.HasPrefix(schemaPath, bgpV4UnicastOCPath) && !strings.HasPrefix(schemaPath, bgpV6UnicastOCPath) {
		return nil, nil
	}

	pi := peerInfo{
		intf:     p.GetElem()[1].GetKey()["name"],
		protocol: p.GetElem()[3].GetKey()["name"],
		neighbor: p.GetElem()[10].GetKey()["neighbor-address"],
		isIPV4:   strings.HasPrefix(schemaPath, bgpV4UnicastOCPath),
	}

	_, hasFreshInfo := c.fresh.Get(ribOCPath)

	// Getting a new path ignores the cache duration and forces a refresh.
	if hasFreshInfo && cachePI == pi {
		return nil, nil
	}

	cachePI = pi
	c.fresh.Set(peerInfoCacheKey, cachePI, -1)

	table, err := ribFromIxiaFn(c, ctx, pi)
	if err != nil {
		return nil, fmt.Errorf("failed to read Ixia table: %w", err)
	}

	var info []bgpLearnedInfo
	if err := unmarshalTable(table, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Ixia table: %w", err)
	}

	dev, err := learnedBGPToOC(info, cachePI.intf, cachePI.protocol, cachePI.neighbor, cachePI.isIPV4)
	if err != nil {
		return nil, err
	}

	var oldRIB *telemetry.Device
	if old, ok := c.fresh.Get(oldRibCacheKey); ok {
		oldRIB = old.(*telemetry.Device)
	}

	notif, err := ygot.Diff(oldRIB, dev)
	if err != nil {
		return nil, fmt.Errorf("failed to render notifications: %w", err)
	}
	notif.Timestamp = time.Now().UnixNano()

	c.fresh.Set(oldRibCacheKey, dev, -1)
	c.fresh.SetDefault(ribOCPath, true)

	return notif, nil
}

func dgToLSPs(dg *ixconfig.TopologyDeviceGroup) []*ixconfig.TopologyRsvpteLsps {
	var lsps []*ixconfig.TopologyRsvpteLsps
	for _, ng := range dg.NetworkGroup {
		for _, ndg := range ng.DeviceGroup {
			for _, lb := range ndg.Ipv4Loopback {
				lsps = append(lsps, lb.RsvpteLsps...)
			}
		}
	}
	for _, ndg := range dg.DeviceGroup {
		lsps = append(lsps, dgToLSPs(ndg)...)
	}
	return lsps
}

func (c *Client) lspsFromIxia(ctx context.Context) (map[string][]*ingressLSP, error) {
	const mvEP = "multivalue/operations/getvalues"
	cfg := c.client.LastImportedConfig()
	if cfg == nil {
		return nil, errors.New("no IxNetwork config found")
	}

	var lsps []*ixconfig.TopologyRsvpteLsps
	for _, topo := range cfg.Topology {
		for _, dg := range topo.DeviceGroup {
			lsps = append(lsps, dgToLSPs(dg)...)
		}
	}

	var ingressLSPNodes []ixconfig.IxiaCfgNode
	nameToIngressLSP := map[string]*ixconfig.TopologyRsvpP2PIngressLsps{}
	for _, lspNode := range lsps {
		if iLSP := lspNode.RsvpP2PIngressLsps; iLSP != nil && *(iLSP.Active.SingleValue.Value) == "true" {
			ingressLSPNodes = append(ingressLSPNodes, iLSP)
			nameToIngressLSP[*(lspNode.Name)] = iLSP
		}
	}

	if err := c.client.UpdateIDs(ctx, cfg, ingressLSPNodes...); err != nil {
		return nil, fmt.Errorf("failed to update IDs for ingress LSPs: %w", err)
	}

	ingressLSPsByPrefix := map[string][]*ingressLSP{}
	for n, lspCfg := range nameToIngressLSP {
		nodeID, err := c.client.NodeID(lspCfg)
		if err != nil {
			return nil, err
		}
		ixLSPs := new(struct {
			State    []string
			SourceIP string
			RemoteIP string
		})
		if err := c.client.Session().Get(ctx, nodeID, ixLSPs); err != nil {
			return nil, fmt.Errorf("failed to fetch ingress LSPs config: %w", err)
		}
		var srcIPs, dstIPs []string
		if err := c.client.Session().Post(ctx, mvEP, ixweb.OpArgs{ixLSPs.SourceIP, 0, len(ixLSPs.State)}, &srcIPs); err != nil {
			return nil, fmt.Errorf("failed to fetch source IPs for LSP: %w", err)
		}
		if err := c.client.Session().Post(ctx, mvEP, ixweb.OpArgs{ixLSPs.RemoteIP, 0, len(ixLSPs.State)}, &dstIPs); err != nil {
			return nil, fmt.Errorf("failed to fetch destination IPs for LSP: %w", err)
		}
		var lsps []*ingressLSP
		for i, state := range ixLSPs.State {
			lsps = append(lsps, &ingressLSP{
				up:  state == "up",
				src: srcIPs[i],
				dst: dstIPs[i],
			})
		}
		ingressLSPsByPrefix[n] = lsps
	}

	return ingressLSPsByPrefix, nil
}

const oldLSPCacheKey = "lspCacheKey"

func (c *Client) pathToOCLSPs(ctx context.Context) (*gpb.Notification, error) {
	if _, hasFreshInfo := c.fresh.Get(lspsOCPath); hasFreshInfo {
		return nil, nil
	}

	ingressLSPs, err := lspsFromIxiaFn(c, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read LSPs from Ixia: %w", err)
	}

	dev := &telemetry.Device{}
	rsvpTE := dev.GetOrCreateNetworkInstance("default").GetOrCreateMpls().GetOrCreateSignalingProtocols().GetOrCreateRsvpTe()
	if err := lsps(ingressLSPs, rsvpTE); err != nil {
		return nil, err
	}

	var oldLSPs *telemetry.Device
	if old, ok := c.fresh.Get(oldLSPCacheKey); ok {
		oldLSPs = old.(*telemetry.Device)
	}

	notif, err := ygot.Diff(oldLSPs, dev)
	if err != nil {
		return nil, fmt.Errorf("failed to render notifications: %w", err)
	}
	notif.Timestamp = time.Now().UnixNano()

	c.fresh.Set(oldLSPCacheKey, dev, -1)
	c.fresh.SetDefault(lspsOCPath, true)

	return notif, nil
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
