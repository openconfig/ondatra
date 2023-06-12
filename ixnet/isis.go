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

package ixnet

import (
	opb "github.com/openconfig/ondatra/proto"
)

// NewISIS constructs a new ISIS configuration.
// Tests should not call this directly; call Interface.ISIS() instead.
func NewISIS(pb *opb.ISISConfig) *ISIS {
	return &ISIS{pb: pb}
}

// ISIS is a representation of a IS-IS config on the ATE.
// Must be constructed by calling NewISIS().
type ISIS struct {
	pb *opb.ISISConfig
}

// IPReachabilityConfig is the IS-IS config for a simulated network pool.
// Must be constructed by calling NewIPReachability().
type IPReachabilityConfig struct {
	pb *opb.IPReachability
}

// ISReachabilityConfig is a representation of the simulated topology of IS-IS nodes.
type ISReachabilityConfig struct {
	pb *opb.ISReachability
}

// WithName assigns a name to the IS-IS reachability config.
// It should be unique among all reachability configs on the interface.
func (ir *ISReachabilityConfig) WithName(name string) *ISReachabilityConfig {
	ir.pb.Name = name
	return ir
}

// ISISNode is a representation of a simulated IS-IS node.
type ISISNode struct {
	pb *opb.ISReachability_Node
}

// ISISNodeLink is a representation of a simulated IS-IS node link.
type ISISNodeLink struct {
	pb *opb.ISReachability_Node_Link
}

// ISISRoutes represents the routes exported by an IS-IS node.
type ISISRoutes struct {
	pb *opb.ISReachability_Node_Routes
}

// WithLevelL1 sets the IS-IS level to L1.
func (i *ISIS) WithLevelL1() *ISIS {
	i.pb.Level = opb.ISISConfig_L1
	return i
}

// WithLevelL2 sets the IS-IS level to L2.
func (i *ISIS) WithLevelL2() *ISIS {
	i.pb.Level = opb.ISISConfig_L2
	return i
}

// WithLevelL1L2 sets the IS-IS level to L1/L2.
func (i *ISIS) WithLevelL1L2() *ISIS {
	i.pb.Level = opb.ISISConfig_L1L2
	return i
}

// WithNetworkTypeBroadcast sets the IS-IS network type to broadcast.
func (i *ISIS) WithNetworkTypeBroadcast() *ISIS {
	i.pb.NetworkType = opb.ISISConfig_BROADCAST
	return i
}

// WithNetworkTypePointToPoint sets the IS-IS network type to point-to-point.
func (i *ISIS) WithNetworkTypePointToPoint() *ISIS {
	i.pb.NetworkType = opb.ISISConfig_POINT_TO_POINT
	return i
}

// WithMetric sets the IS-IS link metric.
func (i *ISIS) WithMetric(metric uint32) *ISIS {
	i.pb.Metric = metric
	return i
}

// WithAreaID sets the area id for the device.
func (i *ISIS) WithAreaID(areaID string) *ISIS {
	i.pb.AreaId = areaID
	return i
}

// WithWideMetricEnabled sets whether the wide metric is enabled.
func (i *ISIS) WithWideMetricEnabled(enabled bool) *ISIS {
	i.pb.EnableWideMetric = enabled
	return i
}

// WithHelloPaddingEnabled sets whether hello padding is enabled.
func (i *ISIS) WithHelloPaddingEnabled(enabled bool) *ISIS {
	i.pb.EnableHelloPadding = enabled
	return i
}

// WithAuthMD5 sets md5 authentication.
func (i *ISIS) WithAuthMD5(key string) *ISIS {
	i.pb.AuthType = opb.ISISConfig_MD5
	i.pb.AuthKey = key
	return i
}

// WithAuthPassword sets password authentication.
func (i *ISIS) WithAuthPassword(key string) *ISIS {
	i.pb.AuthType = opb.ISISConfig_PASSWORD
	i.pb.AuthKey = key
	return i
}

// WithAuthDisabled disables authentication.
func (i *ISIS) WithAuthDisabled() *ISIS {
	i.pb.AuthType = opb.ISISConfig_AUTH_TYPE_UNSPECIFIED
	i.pb.AuthKey = ""
	return i
}

// WithAreaAuthMD5 sets md5 authentication for area.
func (i *ISIS) WithAreaAuthMD5(key string) *ISIS {
	i.pb.AreaAuthType = opb.ISISConfig_MD5
	i.pb.AreaAuthKey = key
	return i
}

// WithAreaAuthPassword sets password authentication for area.
func (i *ISIS) WithAreaAuthPassword(key string) *ISIS {
	i.pb.AreaAuthType = opb.ISISConfig_PASSWORD
	i.pb.AreaAuthKey = key
	return i
}

// WithAreaAuthDisabled disables area authentication.
func (i *ISIS) WithAreaAuthDisabled() *ISIS {
	i.pb.AreaAuthType = opb.ISISConfig_AUTH_TYPE_UNSPECIFIED
	i.pb.AreaAuthKey = ""
	return i
}

// WithDomainAuthMD5 sets md5 authentication for domain.
func (i *ISIS) WithDomainAuthMD5(key string) *ISIS {
	i.pb.DomainAuthType = opb.ISISConfig_MD5
	i.pb.DomainAuthKey = key
	return i
}

// WithDomainAuthPassword sets password authentication for domain.
func (i *ISIS) WithDomainAuthPassword(key string) *ISIS {
	i.pb.DomainAuthType = opb.ISISConfig_PASSWORD
	i.pb.DomainAuthKey = key
	return i
}

// WithDomainAuthDisabled disables domain authentication.
func (i *ISIS) WithDomainAuthDisabled() *ISIS {
	i.pb.DomainAuthType = opb.ISISConfig_AUTH_TYPE_UNSPECIFIED
	i.pb.DomainAuthKey = ""
	return i
}

// WithTEEnabled sets whether traffic engineering is enabled.
func (i *ISIS) WithTEEnabled(enabled bool) *ISIS {
	i.pb.EnableTe = enabled
	return i
}

// WithTERouterID sets the TE router id.
func (i *ISIS) WithTERouterID(routerID string) *ISIS {
	i.pb.TeRouterId = routerID
	return i
}

// WithCapabilityRouterID sets the ISIS capability router id.
func (i *ISIS) WithCapabilityRouterID(routerID string) *ISIS {
	i.pb.CapabilityRouterId = routerID
	return i
}

// WithLSPsDiscarded sets whether to discard learned LSP info.
func (i *ISIS) WithLSPsDiscarded(discard bool) *ISIS {
	i.pb.DiscardLsps = discard
	return i
}

// WithPriority sets the priority of the interface.
func (i *ISIS) WithPriority(priority uint32) *ISIS {
	i.pb.InterfacePriority = priority
	return i
}

// WithHelloInterval sets the interval in seconds between hello packets.
func (i *ISIS) WithHelloInterval(intervalSec uint32) *ISIS {
	i.pb.HelloIntervalSec = intervalSec
	return i
}

// WithDeadInterval sets the interval in seconds before considering that the adjacency is down.
func (i *ISIS) WithDeadInterval(intervalSec uint32) *ISIS {
	i.pb.DeadIntervalSec = intervalSec
	return i
}

// SegmentRouting creates or returns the ISIS Segment Routing configuration.
func (i *ISIS) SegmentRouting() *ISISSegmentRouting {
	if i.pb.SegmentRouting == nil {
		i.pb.SegmentRouting = &opb.ISISSegmentRouting{}
	}
	return &ISISSegmentRouting{pb: i.pb.SegmentRouting}
}

// AddISReachability adds an ISReachability config to the ISIS config.
func (i *ISIS) AddISReachability() *ISReachabilityConfig {
	isR := &ISReachabilityConfig{pb: &opb.ISReachability{}}
	i.pb.IsReachabilities = append(i.pb.IsReachabilities, isR.pb)
	return isR
}

// ClearISReachabilities clears ISReachability configs from the ISIS config.
func (i *ISIS) ClearISReachabilities() *ISIS {
	i.pb.IsReachabilities = nil
	return i
}

// WithActive sets whether the prefixes are active.
func (ip *IPReachabilityConfig) WithActive(active bool) *IPReachabilityConfig {
	ip.pb.Active = active
	return ip
}

// WithIPReachabilityInternal sets route origin as internal.
func (ip *IPReachabilityConfig) WithIPReachabilityInternal() *IPReachabilityConfig {
	ip.pb.RouteOrigin = opb.IPReachability_INTERNAL
	return ip
}

// WithIPReachabilityExternal sets route origin as external.
func (ip *IPReachabilityConfig) WithIPReachabilityExternal() *IPReachabilityConfig {
	ip.pb.RouteOrigin = opb.IPReachability_EXTERNAL
	return ip
}

// WithIPReachabilityMetric sets metric for the reachable IPs.
func (ip *IPReachabilityConfig) WithIPReachabilityMetric(metric uint32) *IPReachabilityConfig {
	ip.pb.Metric = metric
	return ip
}

// WithIPReachabilityAlgorithm sets SR algorithm for the reachable IPs.
func (ip *IPReachabilityConfig) WithIPReachabilityAlgorithm(algo uint32) *IPReachabilityConfig {
	ip.pb.Algorithm = algo
	return ip
}

// WithSIDIndexLabelEnabled enables or disables SID/Index/Label for the reachable IPs.
func (ip *IPReachabilityConfig) WithSIDIndexLabelEnabled(enabled bool) *IPReachabilityConfig {
	ip.pb.EnableSidIndexLabel = enabled
	return ip
}

// WithIPReachabilitySIDIndexLabel sets SID/Index/Label for the reachable IPs.
func (ip *IPReachabilityConfig) WithIPReachabilitySIDIndexLabel(label uint32) *IPReachabilityConfig {
	ip.pb.SidIndexLabel = label
	return ip
}

// WithFlagReadvertise sets the Readvertise(R) flag.
func (ip *IPReachabilityConfig) WithFlagReadvertise(enabled bool) *IPReachabilityConfig {
	ip.pb.FlagReadvertise = enabled
	return ip
}

// WithFlagNodeSID sets the NodeSID(N) flag.
func (ip *IPReachabilityConfig) WithFlagNodeSID(enabled bool) *IPReachabilityConfig {
	ip.pb.FlagNodeSid = enabled
	return ip
}

// WithFlagNoPHP sets the NoPHP(P) flag.
func (ip *IPReachabilityConfig) WithFlagNoPHP(enabled bool) *IPReachabilityConfig {
	ip.pb.FlagNoPhp = enabled
	return ip
}

// WithFlagExplicitNull sets the ExplicitNull(E) flag.
func (ip *IPReachabilityConfig) WithFlagExplicitNull(enabled bool) *IPReachabilityConfig {
	ip.pb.FlagExplicitNull = enabled
	return ip
}

// WithFlagValue sets the Value(V) flag.
func (ip *IPReachabilityConfig) WithFlagValue(enabled bool) *IPReachabilityConfig {
	ip.pb.FlagValue = enabled
	return ip
}

// WithFlagLocal sets the Local(L) flag.
func (ip *IPReachabilityConfig) WithFlagLocal(enabled bool) *IPReachabilityConfig {
	ip.pb.FlagLocal = enabled
	return ip
}

// AddISISNode adds a simulated IS-IS node with ingress/egress metrics defaulted to 10.
func (ir *ISReachabilityConfig) AddISISNode() *ISISNode {
	node := &ISISNode{pb: &opb.ISReachability_Node{
		EgressMetric:  10,
		IngressMetric: 10,
	}}
	ir.pb.Nodes = append(ir.pb.Nodes, node.pb)
	return node
}

// ClearISISNodes clears simulated IS-IS nodes.
func (ir *ISReachabilityConfig) ClearISISNodes() *ISReachabilityConfig {
	ir.pb.Nodes = nil
	return ir
}

// WithIngressMetric sets the metric on the ingress link.
func (node *ISISNode) WithIngressMetric(metric uint32) *ISISNode {
	node.pb.IngressMetric = metric
	return node
}

// WithEgressMetric sets the metric on the egress link.
func (node *ISISNode) WithEgressMetric(metric uint32) *ISISNode {
	node.pb.EgressMetric = metric
	return node
}

// WithSystemID sets the system id for the simulated node.
func (node *ISISNode) WithSystemID(id string) *ISISNode {
	node.pb.SystemId = id
	return node
}

// WithTEEnabled enables TE on the simulated IS-IS node.
func (node *ISISNode) WithTEEnabled(enabled bool) *ISISNode {
	node.pb.EnableTe = enabled
	return node
}

// WithWideMetricEnabled enables wide metric on the simulated IS-IS node.
func (node *ISISNode) WithWideMetricEnabled(enabled bool) *ISISNode {
	node.pb.EnableWideMetric = enabled
	return node
}

// WithTERouterID sets the TE router ID for the node.
func (node *ISISNode) WithTERouterID(id string) *ISISNode {
	node.pb.TeRouterId = id
	return node
}

// WithCapabilityRouterID sets the capability router ID for the node.
func (node *ISISNode) WithCapabilityRouterID(id string) *ISISNode {
	node.pb.CapabilityRouterId = id
	return node
}

// SegmentRouting creates or returns the ISIS Segment Routing configuration for the node.
func (node *ISISNode) SegmentRouting() *ISISSegmentRouting {
	if node.pb.SegmentRouting == nil {
		node.pb.SegmentRouting = &opb.ISISSegmentRouting{}
	}
	return &ISISSegmentRouting{pb: node.pb.SegmentRouting}
}

// ISISSegmentRouting holds the segment routing configuration.
type ISISSegmentRouting struct {
	pb *opb.ISISSegmentRouting
}

// AdjacencySID holds the Adjacency SID configuration.
type AdjacencySID struct {
	pb *opb.ISISSegmentRouting_AdjacencySID
}

// SIDRange holds the SR range configuration.
type SIDRange struct {
	pb *opb.ISISSegmentRouting_SIDRange
}

// WithEnabled sets whether segment routing is enabled.
func (sr *ISISSegmentRouting) WithEnabled(enabled bool) *ISISSegmentRouting {
	sr.pb.Enable = enabled
	return sr
}

// AdjacencySID gets or creates an AdjacencySID configuration with Local(L) and Value(V) flags set.
func (sr *ISISSegmentRouting) AdjacencySID() *AdjacencySID {
	if sr.pb.AdjacencySid == nil {
		sr.pb.AdjacencySid = &opb.ISISSegmentRouting_AdjacencySID{FlagValue: true, FlagLocal: true}
	}
	return &AdjacencySID{pb: sr.pb.AdjacencySid}
}

// WithAdjacencySID sets SID for the adjacency.
func (as *AdjacencySID) WithAdjacencySID(sid string) *AdjacencySID {
	as.pb.Sid = sid
	return as
}

// WithFlagAddressFamily sets the AddressFamily(F) flag.
func (as *AdjacencySID) WithFlagAddressFamily(enabled bool) *AdjacencySID {
	as.pb.OverrideFlagAddressFamily = true
	as.pb.FlagAddressFamily = enabled
	return as
}

// WithFlagBackup sets the Backup(B) flag.
func (as *AdjacencySID) WithFlagBackup(enabled bool) *AdjacencySID {
	as.pb.FlagBackup = enabled
	return as
}

// WithFlagValue sets the Value(V) flag. [Default = true]
func (as *AdjacencySID) WithFlagValue(enabled bool) *AdjacencySID {
	as.pb.FlagValue = enabled
	return as
}

// WithFlagLocal sets the Local(L) flag. [Default = true]
func (as *AdjacencySID) WithFlagLocal(enabled bool) *AdjacencySID {
	as.pb.FlagLocal = enabled
	return as
}

// WithFlagSet sets the Set(S) flag.
func (as *AdjacencySID) WithFlagSet(enabled bool) *AdjacencySID {
	as.pb.FlagSet = enabled
	return as
}

// WithFlagPersistent sets the Persistent(P) flag.
func (as *AdjacencySID) WithFlagPersistent(enabled bool) *AdjacencySID {
	as.pb.FlagPersistent = enabled
	return as
}

// WithSIDIndexLabel sets the SID index label.
func (sr *ISISSegmentRouting) WithSIDIndexLabel(label uint32) *ISISSegmentRouting {
	sr.pb.SidIndexLabel = label
	return sr
}

// WithFlagReadvertise sets the Readvertise(R) flag.
func (sr *ISISSegmentRouting) WithFlagReadvertise(enabled bool) *ISISSegmentRouting {
	sr.pb.FlagReadvertise = enabled
	return sr
}

// WithFlagNodeSID sets the NodeSID(N) flag.
func (sr *ISISSegmentRouting) WithFlagNodeSID(enabled bool) *ISISSegmentRouting {
	sr.pb.FlagNodeSid = enabled
	return sr
}

// WithFlagNoPHP sets the NoPHP(P) flag.
func (sr *ISISSegmentRouting) WithFlagNoPHP(enabled bool) *ISISSegmentRouting {
	sr.pb.FlagNoPhp = enabled
	return sr
}

// WithFlagExplicitNull sets the ExplicitNull(E) flag.
func (sr *ISISSegmentRouting) WithFlagExplicitNull(enabled bool) *ISISSegmentRouting {
	sr.pb.FlagExplicitNull = enabled
	return sr
}

// WithFlagValue sets the Value(V) flag.
func (sr *ISISSegmentRouting) WithFlagValue(enabled bool) *ISISSegmentRouting {
	sr.pb.FlagValue = enabled
	return sr
}

// WithFlagLocal sets the Local(L) flag.
func (sr *ISISSegmentRouting) WithFlagLocal(enabled bool) *ISISSegmentRouting {
	sr.pb.FlagLocal = enabled
	return sr
}

// WithAlgorithms sets the SR algorithms.
func (sr *ISISSegmentRouting) WithAlgorithms(algos ...uint32) *ISISSegmentRouting {
	sr.pb.Algorithms = algos
	return sr
}

// WithPrefixSID sets the prefix SID, specified in CIDR notation.
func (sr *ISISSegmentRouting) WithPrefixSID(sid string) *ISISSegmentRouting {
	sr.pb.PrefixSid = sid
	return sr
}

// AddSRGBRange adds a SRGB range.
func (sr *ISISSegmentRouting) AddSRGBRange() *SIDRange {
	srr := &SIDRange{pb: &opb.ISISSegmentRouting_SIDRange{}}
	sr.pb.SrgbRanges = append(sr.pb.SrgbRanges, srr.pb)
	return srr
}

// ClearSRGBRanges clears SRGB ranges.
func (sr *ISISSegmentRouting) ClearSRGBRanges() *ISISSegmentRouting {
	sr.pb.SrgbRanges = nil
	return sr
}

// AddSRLBRange adds a SRLB range.
func (sr *ISISSegmentRouting) AddSRLBRange() *SIDRange {
	srr := &SIDRange{pb: &opb.ISISSegmentRouting_SIDRange{}}
	sr.pb.SrlbRanges = append(sr.pb.SrlbRanges, srr.pb)
	return srr
}

// ClearSRLBRanges clears SRLB ranges.
func (sr *ISISSegmentRouting) ClearSRLBRanges() *ISISSegmentRouting {
	sr.pb.SrlbRanges = nil
	return sr
}

// WithSIDStartLabel sets the SID start label.
func (srr *SIDRange) WithSIDStartLabel(label uint32) *SIDRange {
	srr.pb.SidStartLabel = label
	return srr
}

// WithSIDCount sets the count of the SID labels.
func (srr *SIDRange) WithSIDCount(c uint32) *SIDRange {
	srr.pb.SidCount = c
	return srr
}

// AddLink adds a simulated IS-IS node link.
func (node *ISISNode) AddLink() *ISISNodeLink {
	link := &ISISNodeLink{pb: &opb.ISReachability_Node_Link{}}
	node.pb.Links = append(node.pb.Links, link.pb)
	return link
}

// ClearLinks clears simulated links for an IS-IS node.
func (node *ISISNode) ClearLinks() *ISISNode {
	node.pb.Links = nil
	return node
}

// WithFromIPv4 sets the IPv4 'from' address for the link.
func (link *ISISNodeLink) WithFromIPv4(ip string) *ISISNodeLink {
	link.pb.FromIpv4 = ip
	return link
}

// WithToIPv4 sets the IPv4 'to' address for the link.
func (link *ISISNodeLink) WithToIPv4(ip string) *ISISNodeLink {
	link.pb.ToIpv4 = ip
	return link
}

// WithFromIPv6 sets the IPv6 'from' address for the link.
func (link *ISISNodeLink) WithFromIPv6(ip string) *ISISNodeLink {
	link.pb.FromIpv6 = ip
	return link
}

// WithToIPv6 sets the IPv6 'to' address for the link.
func (link *ISISNodeLink) WithToIPv6(ip string) *ISISNodeLink {
	link.pb.ToIpv6 = ip
	return link
}

// RoutesIPv4 creates or returns the ISIS IPv4 route configuration.
func (node *ISISNode) RoutesIPv4() *ISISRoutes {
	if node.pb.RoutesIpv4 == nil {
		node.pb.RoutesIpv4 = &opb.ISReachability_Node_Routes{}
	}
	return &ISISRoutes{pb: node.pb.RoutesIpv4}
}

// RoutesIPv6 creates or returns the ISIS IPv6 route configuration.
func (node *ISISNode) RoutesIPv6() *ISISRoutes {
	if node.pb.RoutesIpv6 == nil {
		node.pb.RoutesIpv6 = &opb.ISReachability_Node_Routes{}
	}
	return &ISISRoutes{pb: node.pb.RoutesIpv6}
}

// WithPrefix sets the (CIDR-string) prefix for the exported routes.
func (routes *ISISRoutes) WithPrefix(prefix string) *ISISRoutes {
	routes.pb.Prefix = prefix
	return routes
}

// WithNumRoutes sets the number of exported routes.
func (routes *ISISRoutes) WithNumRoutes(numRoutes uint64) *ISISRoutes {
	routes.pb.NumRoutes = numRoutes
	return routes
}

// IPReachability creates an IP reachability configuration for the network or
// returns the existing config. The default config params are:
//
//	Active: True
//	Route Origin: Internal
//	Metric: 10
func (routes *ISISRoutes) IPReachability() *IPReachabilityConfig {
	if routes.pb.Reachability == nil {
		routes.pb.Reachability = &opb.IPReachability{Active: true, Metric: 10, RouteOrigin: opb.IPReachability_INTERNAL}
	}
	return &IPReachabilityConfig{routes.pb.Reachability}
}
