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
	"sync/atomic"
	"time"

	opb "github.com/openconfig/ondatra/proto"
	dpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var peerCount uint32

// NewBGP returns a new BGP configuration.
// Tests must not call this method directly.
func NewBGP(pb *opb.BgpConfig) *BGP {
	return &BGP{pb}
}

// BGP is a representation of a BGP config on the ATE.
type BGP struct {
	pb *opb.BgpConfig
}

// BGPPeer is a representation of a BGP peer on the ATE.
type BGPPeer struct {
	pb *opb.BgpPeer
}

// AddPeer adds a new BGP peer to the interface.  By default, the peer will have
// the following configuration:
//
//	Active: true
//	Peer type: external
//	Hold time: 90 seconds (per RFC 4271)
//	Keepalive time: 30 second (per RFC 4271)
//	Restart time: 90 seconds (per RFC 4724)
//	Stale time: 300 seconds (per JunOS defaults)
//	Capabilities:
//	  IPv4 unicast
//	  IPv4 multicast
//	  IPv4 MPLS VPN
//	  IPv6 unicast
//	  IPv6 multicast
//	  IPv6 MPLS VPN
//	  VPLS
//	  Route refresh
func (b *BGP) AddPeer() *BGPPeer {
	atomic.AddUint32(&peerCount, 1)
	ppb := &opb.BgpPeer{
		Id:               peerCount,
		Active:           true,
		Type:             opb.BgpPeer_TYPE_EXTERNAL,
		HoldTimeSec:      90,                         // go/rfc/4271#section-10
		KeepaliveTimeSec: 30,                         // go/rfc/4271#section-10
		RestartTime:      dpb.New(90 * time.Second),  // go/rfc/4724#section-4
		StaleTime:        dpb.New(300 * time.Second), // https://www.juniper.net/documentation/us/en/software/junos/bgp/topics/ref/statement/stale-routes-time-edit-protocols-bgp.html
		Capabilities: &opb.BgpPeer_Capabilities{
			Ipv4Unicast:   true,
			Ipv4Multicast: true,
			Ipv4MplsVpn:   true,
			Ipv6Unicast:   true,
			Ipv6Multicast: true,
			Ipv6MplsVpn:   true,
			Vpls:          true,
			RouteRefresh:  true,
		}}
	b.pb.BgpPeers = append(b.pb.BgpPeers, ppb)
	return &BGPPeer{pb: ppb}
}

// ClearPeers clears BGP peers from the interface.
func (b *BGP) ClearPeers() *BGP {
	b.pb.BgpPeers = nil
	return b
}

// ID returns the id of the peer.
func (b *BGPPeer) ID() uint32 {
	return b.pb.Id
}

// WithActive sets whether the peering is active.
func (b *BGPPeer) WithActive(active bool) *BGPPeer {
	b.pb.Active = active
	return b
}

// WithOnLoopback sets whether the peering should configured on the loopback IP.
func (b *BGPPeer) WithOnLoopback(enable bool) *BGPPeer {
	b.pb.OnLoopback = enable
	return b
}

// WithTypeExternal sets the peering type to external.
func (b *BGPPeer) WithTypeExternal() *BGPPeer {
	b.pb.Type = opb.BgpPeer_TYPE_EXTERNAL
	return b
}

// WithTypeInternal sets the peering type to internal.
func (b *BGPPeer) WithTypeInternal() *BGPPeer {
	b.pb.Type = opb.BgpPeer_TYPE_INTERNAL
	return b
}

// WithPeerAddress sets the IP address (IPv4 or IPv6) of the ATE's peer, which
// is typically the IP address of the DUT.
func (b *BGPPeer) WithPeerAddress(peerAddress string) *BGPPeer {
	b.pb.PeerAddress = peerAddress
	return b
}

// WithLocalASN sets the local ASN.
func (b *BGPPeer) WithLocalASN(localASN uint32) *BGPPeer {
	b.pb.LocalAsn = localASN
	return b
}

// WithHoldTime sets the hold time in seconds.
func (b *BGPPeer) WithHoldTime(holdTimeSec uint16) *BGPPeer {
	b.pb.HoldTimeSec = uint32(holdTimeSec)
	return b
}

// WithRestartTime sets the graceful restart time.
func (b *BGPPeer) WithRestartTime(restartTime time.Duration) *BGPPeer {
	b.pb.RestartTime = dpb.New(restartTime)
	return b
}

// WithStaleTime sets the graceful restart stale route time.
func (b *BGPPeer) WithStaleTime(staleTime time.Duration) *BGPPeer {
	b.pb.StaleTime = dpb.New(staleTime)
	return b
}

// WithActAsRestarted sets whether the Restart State bit is advertised in the BGP session graceful restart flags.
func (b *BGPPeer) WithActAsRestarted(enabled bool) *BGPPeer {
	b.pb.ActAsRestarted = enabled
	return b
}

// WithAdvertiseEndOfRIB sets whether the End-of-RIB marker is advertised in the BGP session.
func (b *BGPPeer) WithAdvertiseEndOfRIB(enabled bool) *BGPPeer {
	b.pb.AdvertiseEndOfRib = enabled
	return b
}

// WithKeepaliveTime sets the keepalive time in seconds configured on the ATE.
func (b *BGPPeer) WithKeepaliveTime(keepaliveTimeSec uint16) *BGPPeer {
	b.pb.KeepaliveTimeSec = uint32(keepaliveTimeSec)
	return b
}

// WithMD5Key enables authentication and sets the MD5 key configured on the ATE.
func (b *BGPPeer) WithMD5Key(key string) *BGPPeer {
	b.pb.Md5Key = key
	return b
}

// BGPCapabilities is a representation of BGP capabilities on the ATE.
type BGPCapabilities struct {
	pb *opb.BgpPeer_Capabilities
}

// Capabilities returns the existing capabilities config.
func (b *BGPPeer) Capabilities() *BGPCapabilities {
	return &BGPCapabilities{pb: b.pb.Capabilities}
}

// WithIPv4UnicastEnabled sets whether the IPv4 unicast capability is enabled on
// the ATE.
func (c *BGPCapabilities) WithIPv4UnicastEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4Unicast = enabled
	return c
}

// WithIPv4MulticastEnabled sets whether the IPv4 multicast capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithIPv4MulticastEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4Multicast = enabled
	return c
}

// WithIPv4MPLSVPNEnabled sets whether the IPv4 MPLS VPN capability is enabled
// on the ATE.
func (c *BGPCapabilities) WithIPv4MPLSVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4MplsVpn = enabled
	return c
}

// WithIPv6UnicastEnabled sets whether the IPv6 unicast capability is enabled on
// the ATE.
func (c *BGPCapabilities) WithIPv6UnicastEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6Unicast = enabled
	return c
}

// WithIPv6MulticastEnabled sets whether the IPv6 multicast capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithIPv6MulticastEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6Multicast = enabled
	return c
}

// WithIPv6MPLSVPNEnabled sets whether the IPv6 MPLS VPN capability is enabled
// on the ATE.
func (c *BGPCapabilities) WithIPv6MPLSVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6MplsVpn = enabled
	return c
}

// WithIPv4MDTEnabled sets whether the IPv4 MDT capability is enabled on the
// ATE.
func (c *BGPCapabilities) WithIPv4MDTEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4Mdt = enabled
	return c
}

// WithVPLSEnabled sets whether the VPLS capability is enabled on the ATE.
func (c *BGPCapabilities) WithVPLSEnabled(enabled bool) *BGPCapabilities {
	c.pb.Vpls = enabled
	return c
}

// WithIPv4MulticastVPNEnabled sets whether the IPv4 multicast VPN capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithIPv4MulticastVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4MulticastVpn = enabled
	return c
}

// WithIPv6MulticastVPNEnabled sets whether the IPv6 multicast VPN capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithIPv6MulticastVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6MulticastVpn = enabled
	return c
}

// WithRouteRefreshEnabled sets whether the route refresh capability is enabled
// on the ATE.
func (c *BGPCapabilities) WithRouteRefreshEnabled(enabled bool) *BGPCapabilities {
	c.pb.RouteRefresh = enabled
	return c
}

// WithRouteConstraintEnabled sets whether the route constraint capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithRouteConstraintEnabled(enabled bool) *BGPCapabilities {
	c.pb.RouteConstraint = enabled
	return c
}

// WithLinkStateNonVPNEnabled sets whether the link state non-VPN capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithLinkStateNonVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.LinkStateNonVpn = enabled
	return c
}

// WithEVPNEnabled sets whether the EVPN capability is enabled on the ATE.
func (c *BGPCapabilities) WithEVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.Evpn = enabled
	return c
}

// WithIPv4MulticastBGPMPLSVPNEnabled sets whether the IPv4 multicast BGP MPLS
// VPN capability is enabled on the ATE.
func (c *BGPCapabilities) WithIPv4MulticastBGPMPLSVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4MulticastBgpMplsVpn = enabled
	return c
}

// WithIPv6MulticastBGPMPLSVPNEnabled sets whether the IPv6 multicast BGP MPLS
// VPN capability is enabled on the ATE.
func (c *BGPCapabilities) WithIPv6MulticastBGPMPLSVPNEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6MulticastBgpMplsVpn = enabled
	return c
}

// WithIPv4UnicastFlowSpecEnabled sets whether the IPv4 unicast flow spec
// capability is enabled on the ATE.
func (c *BGPCapabilities) WithIPv4UnicastFlowSpecEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4UnicastFlowSpec = enabled
	return c
}

// WithIPv6UnicastFlowSpecEnabled sets whether the IPv6 unicast flow spec
// capability is enabled on the ATE.
func (c *BGPCapabilities) WithIPv6UnicastFlowSpecEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6UnicastFlowSpec = enabled
	return c
}

// WithIPv4UnicastAddPathEnabled sets whether the IPv4 unicast add path
// capability is enabled on the ATE.
func (c *BGPCapabilities) WithIPv4UnicastAddPathEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4UnicastAddPath = enabled
	return c
}

// WithIPv6UnicastAddPathEnabled sets whether the IPv6 unicast add path
// capability is enabled on the ATE.
func (c *BGPCapabilities) WithIPv6UnicastAddPathEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6UnicastAddPath = enabled
	return c
}

// WithExtendedNextHopEncodingEnabled sets whether the extended next hop
// encoding capability is enabled on the ATE.  This is only applicable to IPv6
// peers.
func (c *BGPCapabilities) WithExtendedNextHopEncodingEnabled(enabled bool) *BGPCapabilities {
	c.pb.ExtendedNextHopEncoding = enabled
	return c
}

// WithIPv4SRTEPolicyEnabled sets whether the IPv4 SR-TE policy capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithIPv4SRTEPolicyEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv4SrtePolicy = enabled
	return c
}

// WithIPv6SRTEPolicyEnabled sets whether the IPv6 SR-TE policy capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithIPv6SRTEPolicyEnabled(enabled bool) *BGPCapabilities {
	c.pb.Ipv6SrtePolicy = enabled
	return c
}

// WithGracefulRestart sets whether the Graceful restart capability is
// enabled on the ATE.
func (c *BGPCapabilities) WithGracefulRestart(enabled bool) *BGPCapabilities {
	c.pb.GracefulRestart = enabled
	return c
}

// BGPAttributes is a representation of BGP attributes on the ATE.
type BGPAttributes struct {
	pb *opb.BgpAttributes
}

// WithActive sets whether the prefixes are active.
func (a *BGPAttributes) WithActive(active bool) *BGPAttributes {
	a.pb.Active = active
	return a
}

// WithNextHopAddress sets the next hop address for the advertised prefixes.
func (a *BGPAttributes) WithNextHopAddress(address string) *BGPAttributes {
	a.pb.NextHopAddress = address
	return a
}

// WithOriginIGP sets the origin to IGP for the advertised prefixes.
func (a *BGPAttributes) WithOriginIGP() *BGPAttributes {
	a.pb.Origin = opb.BgpAttributes_ORIGIN_IGP
	return a
}

// WithOriginEGP sets the origin to EGP for the advertised prefixes.
func (a *BGPAttributes) WithOriginEGP() *BGPAttributes {
	a.pb.Origin = opb.BgpAttributes_ORIGIN_EGP
	return a
}

// WithOriginIncomplete sets the origin to incomplete for the advertised prefixes.
func (a *BGPAttributes) WithOriginIncomplete() *BGPAttributes {
	a.pb.Origin = opb.BgpAttributes_ORIGIN_INCOMPLETE
	return a
}

// WithLocalPreference sets the local preference for the advertised prefixes.
func (a *BGPAttributes) WithLocalPreference(localPref uint32) *BGPAttributes {
	a.pb.LocalPreference = localPref
	return a
}

// BGPCommunities is a representation of BGP communities on the ATE.
type BGPCommunities struct {
	pb *opb.BgpCommunities
}

// Communities creates a BGP communities config for the advertised prefixes or
// returns the existing config.
func (a *BGPAttributes) Communities() *BGPCommunities {
	if a.pb.Communities == nil {
		a.pb.Communities = &opb.BgpCommunities{}
	}
	return &BGPCommunities{pb: a.pb.Communities}
}

// WithNoExport includes the well-known community type NO_EXPORT.
func (c *BGPCommunities) WithNoExport() *BGPCommunities {
	c.pb.NoExport = true
	return c
}

// WithNoAdvertise includes the well-known community type NO_ADVERTISE.
func (c *BGPCommunities) WithNoAdvertise() *BGPCommunities {
	c.pb.NoAdvertise = true
	return c
}

// WithNoExportSubconfed includes the well-known community type
// NO_EXPORT_SUBCONFED.
func (c *BGPCommunities) WithNoExportSubconfed() *BGPCommunities {
	c.pb.NoExportSubconfed = true
	return c
}

// WithLLGRStale includes the well-known community type LLGR_STALE.
func (c *BGPCommunities) WithLLGRStale() *BGPCommunities {
	c.pb.LlgrStale = true
	return c
}

// WithNoLLGR includes the well-known community type NO_LLGR.
func (c *BGPCommunities) WithNoLLGR() *BGPCommunities {
	c.pb.NoLlgr = true
	return c
}

// WithPrivateCommunities includes private communities using an "AS:value"
// format.
func (c *BGPCommunities) WithPrivateCommunities(values ...string) *BGPCommunities {
	c.pb.PrivateCommunities = values
	return c
}

// BGPExtendedCommunityColor is a representation of the BGP color extended
// community on the ATE.
type BGPExtendedCommunityColor struct {
	pb *opb.BgpAttributes_ExtendedCommunity_Color
}

// AddExtendedCommunityColor adds a color extended community to the advertised
// prefixes and returns it for further configuration.  By default, the color
// extended community will have the following configuration:
//
//	CO bits: 00
func (a *BGPAttributes) AddExtendedCommunityColor() *BGPExtendedCommunityColor {
	eccpb := &opb.BgpAttributes_ExtendedCommunity_Color{
		CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
	}
	a.pb.ExtendedCommunities = append(
		a.pb.ExtendedCommunities,
		&opb.BgpAttributes_ExtendedCommunity{
			Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
				Color: eccpb,
			}})
	return &BGPExtendedCommunityColor{pb: eccpb}
}

// ClearExtendedCommunityColors clears color extended communities from the
// advertised prefixes.
func (a *BGPAttributes) ClearExtendedCommunityColors() *BGPAttributes {
	a.pb.ExtendedCommunities = nil
	return a
}

// WithCOBits00 sets the CO bits to 00 for the color extended community.
func (c *BGPExtendedCommunityColor) WithCOBits00() *BGPExtendedCommunityColor {
	c.pb.CoBits = opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00
	return c
}

// WithCOBits01 sets the CO bits to 01 for the color extended community.
func (c *BGPExtendedCommunityColor) WithCOBits01() *BGPExtendedCommunityColor {
	c.pb.CoBits = opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_01
	return c
}

// WithCOBits10 sets the CO bits to 10 for the color extended community.
func (c *BGPExtendedCommunityColor) WithCOBits10() *BGPExtendedCommunityColor {
	c.pb.CoBits = opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_10
	return c
}

// WithCOBits11 sets the CO bits to 11 for the color extended community.
func (c *BGPExtendedCommunityColor) WithCOBits11() *BGPExtendedCommunityColor {
	c.pb.CoBits = opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_11
	return c
}

// WithReservedBits sets the reserved bits for the color extended community.
func (c *BGPExtendedCommunityColor) WithReservedBits(reservedBits uint16) *BGPExtendedCommunityColor {
	c.pb.ReservedBits = uint32(reservedBits)
	return c
}

// WithValue sets the value for the color extended community.
func (c *BGPExtendedCommunityColor) WithValue(value uint32) *BGPExtendedCommunityColor {
	c.pb.Value = value
	return c
}

// WithASNSetModeDoNotInclude sets the ASN set mode to not include the local ASN
// for the advertised prefixes.
func (a *BGPAttributes) WithASNSetModeDoNotInclude() *BGPAttributes {
	a.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE
	return a
}

// WithASNSetModeSEQ sets the ASN set mode to include the local ASN as AS-SEQ
// for the advertised prefixes.
func (a *BGPAttributes) WithASNSetModeSEQ() *BGPAttributes {
	a.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ
	return a
}

// WithASNSetModeSET sets the ASN set mode to include the local ASN as AS-SET
// for the advertised prefixes.
func (a *BGPAttributes) WithASNSetModeSET() *BGPAttributes {
	a.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SET
	return a
}

// WithASNSetModeSEQConfederation sets the ASN set mode to include the local ASN
// as AS-SEQ-CONFEDERATION for the advertised prefixes.
func (a *BGPAttributes) WithASNSetModeSEQConfederation() *BGPAttributes {
	a.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ_CONFEDERATION
	return a
}

// WithASNSetModeSETConfederation sets the ASN set mode to include the local ASN
// as AS-SET-CONFEDERATION for the advertised prefixes.
func (a *BGPAttributes) WithASNSetModeSETConfederation() *BGPAttributes {
	a.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SET_CONFEDERATION
	return a
}

// WithASNSetModePrepend sets the ASN set mode to prepend the local ASN to the
// first segment for the advertised prefixes.
func (a *BGPAttributes) WithASNSetModePrepend() *BGPAttributes {
	a.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_PREPEND
	return a
}

// BGPASPathSegment is a representation of an BGP AS path segment on the ATE.
type BGPASPathSegment struct {
	pb *opb.BgpAttributes_AsPathSegment
}

// AddASPathSegment adds an AS path segment with the given ASNs.  By default,
// the AS path segment will have the following configuration:
//
//	Segment type: AS-SEQ
func (a *BGPAttributes) AddASPathSegment(asns ...uint32) *BGPASPathSegment {
	apspb := &opb.BgpAttributes_AsPathSegment{
		Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
		Asns: asns,
	}
	a.pb.AsPathSegments = append(a.pb.AsPathSegments, apspb)
	return &BGPASPathSegment{pb: apspb}
}

// ClearASPathSegments clears AS path segments.
func (a *BGPAttributes) ClearASPathSegments() *BGPAttributes {
	a.pb.AsPathSegments = nil
	return a
}

// WithTypeSEQ sets the type of the AS path segment to AS-SEQ.
func (a *BGPASPathSegment) WithTypeSEQ() *BGPASPathSegment {
	a.pb.Type = opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ
	return a
}

// WithTypeSET sets the type of the AS path segment to AS-SET.
func (a *BGPASPathSegment) WithTypeSET() *BGPASPathSegment {
	a.pb.Type = opb.BgpAttributes_AsPathSegment_TYPE_AS_SET
	return a
}

// WithTypeSEQConfederation sets the type of the AS path segment to
// AS-SEQ-CONFEDERATION.
func (a *BGPASPathSegment) WithTypeSEQConfederation() *BGPASPathSegment {
	a.pb.Type = opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ_CONFEDERATION
	return a
}

// WithTypeSETConfederation sets the type of the AS path segment to
// AS-SET-CONFEDERATION.
func (a *BGPASPathSegment) WithTypeSETConfederation() *BGPASPathSegment {
	a.pb.Type = opb.BgpAttributes_AsPathSegment_TYPE_AS_SET_CONFEDERATION
	return a
}

// WithOriginatorID sets the originator ID for the routes.
func (a *BGPAttributes) WithOriginatorID(id string) *BGPAttributes {
	a.pb.OriginatorId = &opb.StringIncRange{
		Start: id,
		Step:  "0.0.0.0",
	}
	return a
}

// OriginatorIDRange sets the originator ID of the routes to a range of values
// and returns the range. By default, the range will have the following
// configuration:
//
//	Start: 0.0.0.1
//	Step: 0.0.0.1
func (a *BGPAttributes) OriginatorIDRange() *StringIncRange {
	if a.pb.OriginatorId == nil {
		a.pb.OriginatorId = &opb.StringIncRange{
			Start: "0.0.0.1",
			Step:  "0.0.0.1",
		}
	}
	return NewStringIncRange(a.pb.OriginatorId)
}

// WithClusterIDs sets the given cluster IDs for the associated routes.
func (a *BGPAttributes) WithClusterIDs(values ...string) *BGPAttributes {
	a.pb.ClusterIds = values
	return a
}

// WithAdvertisementProtocolV4 sets the routes to be advertised over a BGP V4
// peer.
func (a *BGPAttributes) WithAdvertisementProtocolV4() *BGPAttributes {
	a.pb.AdvertisementProtocol = opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V4
	return a
}

// WithAdvertisementProtocolV6 sets the routes to be advertised over a BGP V6
// peer.
func (a *BGPAttributes) WithAdvertisementProtocolV6() *BGPAttributes {
	a.pb.AdvertisementProtocol = opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V6
	return a
}

// WithAdvertisementProtocolV4AndV6 sets the routes to be advertised over both
// BGP V4 and V6 peers.
func (a *BGPAttributes) WithAdvertisementProtocolV4AndV6() *BGPAttributes {
	a.pb.AdvertisementProtocol = opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V4_AND_V6
	return a
}

// WithAdvertisementProtocolSameAsRoute sets the routes to be exported over the
// BGP peer matching the routes' IP protocol.
func (a *BGPAttributes) WithAdvertisementProtocolSameAsRoute() *BGPAttributes {
	a.pb.AdvertisementProtocol = opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_SAME_AS_ROUTE
	return a
}

// BGPSRTEPolicyGroup is a representation of BGP SR-TE policies on the ATE.
type BGPSRTEPolicyGroup struct {
	pb *opb.BgpPeer_SrtePolicyGroup
}

// AddSRTEPolicyGroup adds SR-TE policies to the peer.  Multiple policies that
// share the same configuration can be added by using WithCount.  By default,
// the SR-TE policies will have the following configuration:
//
//	Count: 1
//	Active: true
//	Distinguisher: 1
//	Policy color: 100
func (b *BGPPeer) AddSRTEPolicyGroup() *BGPSRTEPolicyGroup {
	spgpb := &opb.BgpPeer_SrtePolicyGroup{
		Count:         1,
		Active:        true,
		Distinguisher: 1,
		PolicyColor: &opb.UInt32IncRange{
			Start: 100,
			Step:  0,
		}}
	b.pb.SrtePolicyGroups = append(b.pb.SrtePolicyGroups, spgpb)
	return &BGPSRTEPolicyGroup{pb: spgpb}
}

// ClearSRTEPolicyGroups clears SR-TE policies from the peer.
func (b *BGPPeer) ClearSRTEPolicyGroups() *BGPPeer {
	b.pb.SrtePolicyGroups = nil
	return b
}

// WithCount sets the number of policies that will be added with the same
// configuration.
func (p *BGPSRTEPolicyGroup) WithCount(count uint32) *BGPSRTEPolicyGroup {
	p.pb.Count = count
	return p
}

// WithActive sets whether the policies are active.
func (p *BGPSRTEPolicyGroup) WithActive(active bool) *BGPSRTEPolicyGroup {
	p.pb.Active = active
	return p
}

// WithDistinguisher sets the distinguisher of the policies.
func (p *BGPSRTEPolicyGroup) WithDistinguisher(distinguisher uint32) *BGPSRTEPolicyGroup {
	p.pb.Distinguisher = distinguisher
	return p
}

// WithPolicyColor sets the policy color of the policies.
func (p *BGPSRTEPolicyGroup) WithPolicyColor(policyColor uint32) *BGPSRTEPolicyGroup {
	p.pb.PolicyColor = &opb.UInt32IncRange{
		Start: policyColor,
		Step:  0,
	}
	return p
}

// PolicyColorRange sets the policy color of the policies to a range of values
// and returns the range.  By default, the range will have the following
// configuration:
//
//	Start: 100
//	Step: 1
func (p *BGPSRTEPolicyGroup) PolicyColorRange() *UInt32IncRange {
	if p.pb.PolicyColor == nil {
		p.pb.PolicyColor = &opb.UInt32IncRange{
			Start: 100,
			Step:  1,
		}
	}
	return NewUInt32IncRange(p.pb.PolicyColor)
}

// WithEndpoint sets the endpoint of the policies.
func (p *BGPSRTEPolicyGroup) WithEndpoint(endpoint string) *BGPSRTEPolicyGroup {
	p.pb.Endpoint = endpoint
	return p
}

// WithNextHopAddress sets the next hop address of the policies.
func (p *BGPSRTEPolicyGroup) WithNextHopAddress(address string) *BGPSRTEPolicyGroup {
	p.pb.NextHopAddress = address
	return p
}

// WithOriginatorID sets the originator ID of the policies.
func (p *BGPSRTEPolicyGroup) WithOriginatorID(id string) *BGPSRTEPolicyGroup {
	p.pb.OriginatorId = &opb.StringIncRange{
		Start: id,
		Step:  "0.0.0.0",
	}
	return p
}

// OriginatorIDRange sets the originator ID of the policies to a range of values
// and returns the range.  By default, the range will have the following
// configuration:
//
//	Start: 0.0.0.1
//	Step: 0.0.0.1
func (p *BGPSRTEPolicyGroup) OriginatorIDRange() *StringIncRange {
	if p.pb.OriginatorId == nil {
		p.pb.OriginatorId = &opb.StringIncRange{
			Start: "0.0.0.1",
			Step:  "0.0.0.1",
		}
	}
	return NewStringIncRange(p.pb.OriginatorId)
}

// Communities creates a BGP communities config for the advertised prefixes or
// returns the existing config.
func (p *BGPSRTEPolicyGroup) Communities() *BGPCommunities {
	if p.pb.Communities == nil {
		p.pb.Communities = &opb.BgpCommunities{}
	}
	return &BGPCommunities{pb: p.pb.Communities}
}

// WithASNSetModeDoNotInclude sets the ASN set mode to not include the local ASN
// of the policies.
func (p *BGPSRTEPolicyGroup) WithASNSetModeDoNotInclude() *BGPSRTEPolicyGroup {
	p.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE
	return p
}

// WithASNSetModeSEQ sets the ASN set mode to include the local ASN as AS-SEQ
// of the policies.
func (p *BGPSRTEPolicyGroup) WithASNSetModeSEQ() *BGPSRTEPolicyGroup {
	p.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ
	return p
}

// WithASNSetModeSET sets the ASN set mode to include the local ASN as AS-SET
// of the policies.
func (p *BGPSRTEPolicyGroup) WithASNSetModeSET() *BGPSRTEPolicyGroup {
	p.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SET
	return p
}

// WithASNSetModeSEQConfederation sets the ASN set mode to include the local ASN
// as AS-SEQ-CONFEDERATION of the policies.
func (p *BGPSRTEPolicyGroup) WithASNSetModeSEQConfederation() *BGPSRTEPolicyGroup {
	p.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ_CONFEDERATION
	return p
}

// WithASNSetModeSETConfederation sets the ASN set mode to include the local ASN
// as AS-SET-CONFEDERATION of the policies.
func (p *BGPSRTEPolicyGroup) WithASNSetModeSETConfederation() *BGPSRTEPolicyGroup {
	p.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_AS_SET_CONFEDERATION
	return p
}

// WithASNSetModePrepend sets the ASN set mode to prepend the local ASN to the
// first segment of the policies.
func (p *BGPSRTEPolicyGroup) WithASNSetModePrepend() *BGPSRTEPolicyGroup {
	p.pb.AsnSetMode = opb.BgpAsnSetMode_ASN_SET_MODE_PREPEND
	return p
}

// WithENLP sets the ENLP value of the policies.
func (p *BGPSRTEPolicyGroup) WithENLP(enlp uint32) *BGPSRTEPolicyGroup {
	p.pb.Enlp = &opb.BgpPeer_SrtePolicyGroup_Enlp{
		Enlp: enlp,
	}
	return p
}

// WithPreference sets the preference of the policies.
func (p *BGPSRTEPolicyGroup) WithPreference(preference uint32) *BGPSRTEPolicyGroup {
	p.pb.Preference = &opb.BgpPeer_SrtePolicyGroup_Preference{
		Preference: preference,
	}
	return p
}

// WithBindingSIDNone sets the binding SID of the polices to no binding.
func (p *BGPSRTEPolicyGroup) WithBindingSIDNone() *BGPSRTEPolicyGroup {
	p.pb.Binding = &opb.BgpPeer_SrtePolicyGroup_Binding{
		Type: &opb.BgpPeer_SrtePolicyGroup_Binding_NoBinding{
			NoBinding: &emptypb.Empty{},
		}}
	return p
}

// WithBindingSID4Octet sets the binding SID of the polices to the 4-octet
// value.
func (p *BGPSRTEPolicyGroup) WithBindingSID4Octet(bsid uint32) *BGPSRTEPolicyGroup {
	p.pb.Binding = &opb.BgpPeer_SrtePolicyGroup_Binding{
		Type: &opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSid{
			FourOctetSid: &opb.UInt32IncRange{
				Start: bsid,
				Step:  0,
			}}}
	return p
}

// WithBindingSID4OctetMPLS sets the binding SID of the polices to the 4-octet
// value as an MPLS label.
func (p *BGPSRTEPolicyGroup) WithBindingSID4OctetMPLS(bsid uint32) *BGPSRTEPolicyGroup {
	p.pb.Binding = &opb.BgpPeer_SrtePolicyGroup_Binding{
		Type: &opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSidAsMplsLabel{
			FourOctetSidAsMplsLabel: &opb.UInt32IncRange{
				Start: bsid,
				Step:  0,
			}}}
	return p
}

// BindingSID4OctetRange sets the binding SID of the polices to a range of
// values and returns the range.  By default, the range will have the following
// configuration:
//
//	Start: 0
//	Step: 1
func (p *BGPSRTEPolicyGroup) BindingSID4OctetRange() *UInt32IncRange {
	if p.pb.Binding == nil {
		p.pb.Binding = &opb.BgpPeer_SrtePolicyGroup_Binding{}
	}
	var rpb *opb.UInt32IncRange
	switch t := p.pb.Binding.Type.(type) {
	case *opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSid:
		rpb = t.FourOctetSid
	default:
		rpb = &opb.UInt32IncRange{
			Start: 0,
			Step:  1,
		}
		p.pb.Binding.Type = &opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSid{
			FourOctetSid: rpb,
		}
	}
	return NewUInt32IncRange(rpb)
}

// BindingSID4OctetMPLSRange sets the binding SID of the polices to a range of
// values as MPLS labels and returns the range.  By default, the range will have
// the following configuration:
//
//	Start: 0
//	Step: 1
func (p *BGPSRTEPolicyGroup) BindingSID4OctetMPLSRange() *UInt32IncRange {
	if p.pb.Binding == nil {
		p.pb.Binding = &opb.BgpPeer_SrtePolicyGroup_Binding{}
	}
	var rpb *opb.UInt32IncRange
	switch t := p.pb.Binding.Type.(type) {
	case *opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSidAsMplsLabel:
		rpb = t.FourOctetSidAsMplsLabel
	default:
		rpb = &opb.UInt32IncRange{
			Start: 0,
			Step:  1,
		}
		p.pb.Binding.Type = &opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSidAsMplsLabel{
			FourOctetSidAsMplsLabel: rpb,
		}
	}
	return NewUInt32IncRange(rpb)
}

// WithBindingSIDIPv6 sets the binding SID of the polices to the IPv6 value.
func (p *BGPSRTEPolicyGroup) WithBindingSIDIPv6(bsid string) *BGPSRTEPolicyGroup {
	p.pb.Binding = &opb.BgpPeer_SrtePolicyGroup_Binding{
		Type: &opb.BgpPeer_SrtePolicyGroup_Binding_Ipv6Sid{
			Ipv6Sid: bsid,
		}}
	return p
}

// BGPSegmentList is a representation of a segment list on the ATE.
type BGPSegmentList struct {
	pb *opb.BgpPeer_SrtePolicyGroup_SegmentList
}

// AddSegmentList adds a segment list to the policies.  By default, the segment
// list will have the following configuration:
//
//	Active: true
func (p *BGPSRTEPolicyGroup) AddSegmentList() *BGPSegmentList {
	slpb := &opb.BgpPeer_SrtePolicyGroup_SegmentList{
		Active: true,
	}
	p.pb.SegmentLists = append(p.pb.SegmentLists, slpb)
	return &BGPSegmentList{pb: slpb}
}

// ClearSegmentLists clears segment lists from the policies.
func (p *BGPSRTEPolicyGroup) ClearSegmentLists() *BGPSRTEPolicyGroup {
	p.pb.SegmentLists = nil
	return p
}

// WithActive sets whether the segment list is active.
func (l *BGPSegmentList) WithActive(active bool) *BGPSegmentList {
	l.pb.Active = active
	return l
}

// WithWeight sets the weight of the segment list.
func (l *BGPSegmentList) WithWeight(weight uint32) *BGPSegmentList {
	l.pb.Weight = &opb.BgpPeer_SrtePolicyGroup_SegmentList_Weight{
		Weight: weight,
	}
	return l
}

// BGPSegment is a representation of a segment on the ATE.
type BGPSegment struct {
	pb *opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment
}

// AddSegment adds a segment to the segment list.  By default, the segment will
// have the following configuration:
//
//	Active: true
func (l *BGPSegmentList) AddSegment() *BGPSegment {
	spb := &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{
		Active: true,
	}
	l.pb.Segments = append(l.pb.Segments, spb)
	return &BGPSegment{pb: spb}
}

// ClearSegments clears segments from the segment list.
func (l *BGPSegmentList) ClearSegments() *BGPSegmentList {
	l.pb.Segments = nil
	return l
}

// WithActive sets whether the segment is active.
func (s *BGPSegment) WithActive(active bool) *BGPSegment {
	s.pb.Active = active
	return s
}

// BGPSegmentMPLSSID is a representation of an MPLS SID segment on the ATE.
type BGPSegmentMPLSSID struct {
	pb *opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid
}

// BGPSegmentMPLSSID sets the segment type to MPLS SID and returns the segment.
// By default, the segment will have the following configuration:
//
//	Label: 16
//	TC: 0
//	S: false
//	TTL: 255
func (s *BGPSegment) BGPSegmentMPLSSID() *BGPSegmentMPLSSID {
	var mspb *opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid
	switch t := s.pb.Type.(type) {
	case *opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid_:
		mspb = t.MplsSid
	default:
		mspb = &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid{
			Label: 16,
			Tc:    0,
			S:     false,
			Ttl:   255,
		}
		s.pb.Type = &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid_{
			MplsSid: mspb,
		}
	}
	return &BGPSegmentMPLSSID{pb: mspb}
}

// WithLabel sets the label of the MPLS SID segment.
func (m *BGPSegmentMPLSSID) WithLabel(label uint32) *BGPSegmentMPLSSID {
	m.pb.Label = label
	return m
}

// WithTC sets the traffic class of the MPLS SID segment.
func (m *BGPSegmentMPLSSID) WithTC(tc uint8) *BGPSegmentMPLSSID {
	m.pb.Tc = uint32(tc)
	return m
}

// WithS sets the bottom-of-stack bit of the MPLS SID segment.
func (m *BGPSegmentMPLSSID) WithS(s bool) *BGPSegmentMPLSSID {
	m.pb.S = s
	return m
}

// WithTTL sets the TTL of the MPLS SID segment.
func (m *BGPSegmentMPLSSID) WithTTL(ttl uint8) *BGPSegmentMPLSSID {
	m.pb.Ttl = uint32(ttl)
	return m
}

// WithIPv6SID sets the segment type to IPv6 SID and sets the SID to the IPv6
// value.
func (s *BGPSegment) WithIPv6SID(sid string) *BGPSegment {
	s.pb.Type = &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_Ipv6Sid{
		Ipv6Sid: sid,
	}
	return s
}
