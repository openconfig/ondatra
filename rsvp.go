// Copyright 2021 Google LLC
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
	opb "github.com/openconfig/ondatra/proto"
)

// RSVP represents an RSVP config on an ATE.
type RSVP struct {
	pb *opb.RsvpConfig
}

// WithISISReachability configures which IS-IS reachability config defines the
// routes available for LSPs.
func (r *RSVP) WithISISReachability(isr string) *RSVP {
	r.pb.IsReachabilityName = isr
	return r
}

// WithBundleMessageSending enables or disables bundle messaging for the RSVP interface.
func (r *RSVP) WithBundleMessageSending(enable bool) *RSVP {
	r.pb.BundleMessageSending = enable
	return r
}

// WithRefreshReduction enables or disables refresh reduction for the RSVP interface.
func (r *RSVP) WithRefreshReduction(enable bool) *RSVP {
	r.pb.RefreshReduction = enable
	return r
}

// AddLoopback adds a loopback address for the RSVP instance.
func (r *RSVP) AddLoopback() *RSVPLoopback {
	lb := &RSVPLoopback{pb: &opb.RsvpConfig_Loopback{}}
	r.pb.Loopbacks = append(r.pb.Loopbacks, lb.pb)
	return lb
}

// RSVPLoopback represents a simulated loopback interface for RSVP configuration.
type RSVPLoopback struct {
	pb *opb.RsvpConfig_Loopback
}

// WithLocalIP configures the local loopback address for the RSVP loopback instance.
func (r *RSVPLoopback) WithLocalIP(ipv4CIDR string) *RSVPLoopback {
	r.pb.LocalIpCidr = ipv4CIDR
	return r
}

// AddIngressLSP adds an ingress LSP for the RSVP loopback.
func (r *RSVPLoopback) AddIngressLSP() *RSVPIngressLSP {
	lsp := &RSVPIngressLSP{pb: &opb.RsvpConfig_Loopback_IngressLSP{}}
	r.pb.IngressLsps = append(r.pb.IngressLsps, lsp.pb)
	return lsp
}

// RSVPIngressLSP respresents an ingress LSP on an RSVP configuration
type RSVPIngressLSP struct {
	pb *opb.RsvpConfig_Loopback_IngressLSP
}

// WithRemoteIP configures the remote address for the RSVP ingress LSP.
func (r *RSVPIngressLSP) WithRemoteIP(ipv4CIDR string) *RSVPIngressLSP {
	r.pb.RemoteIpCidr = ipv4CIDR
	return r
}

// WithLocalProtection enables or disables local protection for the RSVP ingress LSP.
func (r *RSVPIngressLSP) WithLocalProtection(enable bool) *RSVPIngressLSP {
	r.pb.LocalProtection = enable
	return r
}

// WithBandwidthProtection enables or disables bandwidth protection for the RSVP ingress LSP.
func (r *RSVPIngressLSP) WithBandwidthProtection(enable bool) *RSVPIngressLSP {
	r.pb.BandwidthProtection = enable
	return r
}

// WithFastReroute enables or disables fast reroute for the RSVP ingress LSP.
func (r *RSVPIngressLSP) WithFastReroute(enable bool) *RSVPIngressLSP {
	r.pb.FastReroute = enable
	return r
}

// WithPathReoptimization enables or disables path re-optimization for the RSVP ingress LSP.
func (r *RSVPIngressLSP) WithPathReoptimization(enable bool) *RSVPIngressLSP {
	r.pb.PathReoptimization = enable
	return r
}

// AddERO adds an ERO for the RSVP ingress LSP.
func (r *RSVPIngressLSP) AddERO() *RSVPIngressLSPERO {
	ero := &RSVPIngressLSPERO{pb: &opb.RsvpConfig_Loopback_IngressLSP_ERO{}}
	r.pb.Eros = append(r.pb.Eros, ero.pb)
	return ero
}

// RSVPIngressLSPERO represents an explicit route object for an RSVP ingress LSP.
type RSVPIngressLSPERO struct {
	pb *opb.RsvpConfig_Loopback_IngressLSP_ERO
}

// WithIP configures the address for the RSVP ingress LSP ERO.
func (r *RSVPIngressLSPERO) WithIP(ipv4CIDR string) *RSVPIngressLSPERO {
	r.pb.Ipv4Cidr = ipv4CIDR
	return r
}

// AddRRO adds an RRO for the RSVP ingress LSP.
func (r *RSVPIngressLSP) AddRRO() *RSVPIngressLSPRRO {
	rro := &RSVPIngressLSPRRO{pb: &opb.RsvpConfig_Loopback_IngressLSP_RRO{}}
	r.pb.Rros = append(r.pb.Rros, rro.pb)
	return rro
}

// RSVPIngressLSPRRO represents a record route object for an RSVP ingress LSP.
type RSVPIngressLSPRRO struct {
	pb *opb.RsvpConfig_Loopback_IngressLSP_RRO
}

// WithIP configures the address for the RSVP ingress LSP RRO.
func (r *RSVPIngressLSPRRO) WithIP(ipv4 string) *RSVPIngressLSPRRO {
	r.pb.Ipv4 = ipv4
	return r
}
