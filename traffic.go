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
	"fmt"
	"testing"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/ate"
	"github.com/openconfig/ondatra/internal/events"

	opb "github.com/openconfig/ondatra/proto"
)

// Traffic is the ATE Traffic API.
type Traffic struct {
	ate binding.ATE
}

func (tr *Traffic) String() string {
	return fmt.Sprintf("Traffic%+v", *tr)
}

// NewFlow returns a new Traffic flow. By default the flow will have the following properties:
//
//	frame size: max(64, sizeof(headers) + (4 * number of ethernet headers))
//	frame rate: 10% of the line rate
//	header addresses inferred: true
func (tr *Traffic) NewFlow(name string) *Flow {
	return &Flow{pb: &opb.Flow{Name: name}}
}

// Start pushes and starts provided traffic flows on the ATE.
// If no flows are provided, starts the previously pushed flows.
func (tr *Traffic) Start(t testing.TB, flows ...*Flow) {
	t.Helper()
	t = events.ActionStarted(t, "Starting traffic on %s", tr.ate)
	if err := tr.start(flows); err != nil {
		t.Fatalf("Start(t) on %s: %v", tr, err)
	}
}

func (tr *Traffic) start(flows []*Flow) error {
	var pbs []*opb.Flow
	for _, f := range flows {
		pbs = append(pbs, f.pb)
	}
	return ate.StartTraffic(context.Background(), tr.ate, pbs)
}

// Update modifies already-running traffic on the ATE.
// Only frame rate and frame size are updated while traffic is running.
// Frame size can only be updated to a frame size of the same type.
func (tr *Traffic) Update(t testing.TB, flows ...*Flow) {
	t.Helper()
	t = events.ActionStarted(t, "Updating traffic on %s", tr.ate)
	if err := tr.update(flows); err != nil {
		t.Fatalf("Update(t) on %s: %v", tr, err)
	}
}

func (tr *Traffic) update(flows []*Flow) error {
	var pbs []*opb.Flow
	for _, f := range flows {
		pbs = append(pbs, f.pb)
	}
	return ate.UpdateTraffic(context.Background(), tr.ate, pbs)
}

// Stop stops all traffic flows on the ATE.
func (tr *Traffic) Stop(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Stopping traffic on %s", tr.ate)
	if err := ate.StopTraffic(context.Background(), tr.ate); err != nil {
		t.Fatalf("Stop(t) on %s: %v", tr, err)
	}
}

// IMIXCustom is an representation of custom IMIX entries to be configured for a flow on the ATE.
type IMIXCustom struct {
	pb *opb.FrameSize_ImixCustom
}

// Flow is an representation of a flow that is to be configured on the ATE.
type Flow struct {
	headers []Header
	pb      *opb.Flow
}

// Endpoint is a potential source or destination of a flow.
// There are two types of endpoints: Interfaces and Networks.
type Endpoint interface {
	EndpointPB() *opb.Flow_Endpoint
}

// Transmission represents the transmission pattern of a flow.
type Transmission struct {
	pb *opb.Transmission
}

func (f *Flow) String() string {
	return f.pb.String()
}

// Name returns the name of the flow.
func (f *Flow) Name() string {
	return f.pb.Name
}

// Headers returns the headers of the flow.
func (f *Flow) Headers() []Header {
	return f.headers
}

// WithSrcEndpoints sets the flow's source to the specified endpoints.
func (f *Flow) WithSrcEndpoints(eps ...Endpoint) *Flow {
	f.pb.SrcEndpoints = endpointPBs(eps)
	return f
}

// WithDstEndpoints sets the flow's destination to the specified endpoints.
func (f *Flow) WithDstEndpoints(eps ...Endpoint) *Flow {
	f.pb.DstEndpoints = endpointPBs(eps)
	return f
}

func endpointPBs(eps []Endpoint) []*opb.Flow_Endpoint {
	var epPBs []*opb.Flow_Endpoint
	for _, ep := range eps {
		if ep == nil {
			log.Fatalf("nil endpoint not allowed: %v", eps)
		}
		epPBs = append(epPBs, ep.EndpointPB())
	}
	return epPBs
}

// WithHeaders sets the flow's packet headers.
func (f *Flow) WithHeaders(headers ...Header) *Flow {
	f.headers = headers
	f.pb.Headers = nil
	for _, h := range headers {
		f.pb.Headers = append(f.pb.Headers, h.asPB())
	}
	return f
}

// WithFrameRatePct sets the flow's rate to be pct% of the line rate.
func (f *Flow) WithFrameRatePct(pct float64) *Flow {
	f.pb.FrameRate = &opb.FrameRate{Type: &opb.FrameRate_Percent{Percent: pct}}
	return f
}

// WithFrameRateFPS sets the flow's rate to be n frames per second.
func (f *Flow) WithFrameRateFPS(n uint64) *Flow {
	f.pb.FrameRate = &opb.FrameRate{Type: &opb.FrameRate_Fps{Fps: n}}
	return f
}

// WithFrameRateBPS sets the flow's rate to be n bits per second.
func (f *Flow) WithFrameRateBPS(n uint64) *Flow {
	f.pb.FrameRate = &opb.FrameRate{Type: &opb.FrameRate_Bps{Bps: n}}
	return f
}

// EgressTracking represents the egress tracking configuration of a flow.
type EgressTracking struct {
	pb *opb.EgressTracking
}

// EgressTracking returns the egress tracking configuration of the flow.
// By default the configuration has the following properties:
// enabled: true
// offset: 32
// width: 3
// count: 8
func (f *Flow) EgressTracking() *EgressTracking {
	if f.pb.EgressTracking == nil {
		f.pb.EgressTracking = &opb.EgressTracking{
			Enabled: true,
			Offset:  32,
			Width:   3,
			Count:   8,
		}
	}
	return &EgressTracking{f.pb.EgressTracking}
}

// WithEnabled sets whether egress tracking is enabled.
func (et *EgressTracking) WithEnabled(enabled bool) *EgressTracking {
	et.pb.Enabled = enabled
	return et
}

// WithOffset sets the egress tracking bit offset.
func (et *EgressTracking) WithOffset(offset uint32) *EgressTracking {
	et.pb.Offset = offset
	return et
}

// WithWidth sets the egress tracking bit width.
func (et *EgressTracking) WithWidth(width uint32) *EgressTracking {
	et.pb.Width = width
	return et
}

// WithCount sets the number of unique values to egress track.
func (et *EgressTracking) WithCount(count uint32) *EgressTracking {
	et.pb.Count = count
	return et
}

// WithIngressTrackingByPorts enables ingress tracking by rx/tx ports.
func (f *Flow) WithIngressTrackingByPorts(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.Ports = enable
	return f
}

// WithIngressTrackingByMplsLabel enables ingress tracking by mpls label.
func (f *Flow) WithIngressTrackingByMplsLabel(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.MplsLabel = enable
	return f
}

// WithIngressTrackingBySrcEndpoint enables ingress tracking by source endpoint.
func (f *Flow) WithIngressTrackingBySrcEndpoint(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.SrcEndpoint = enable
	return f
}

// WithIngressTrackingByDstEndpoint enables ingress tracking by destination endpoint.
func (f *Flow) WithIngressTrackingByDstEndpoint(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.DstEndpoint = enable
	return f
}

// WithIngressTrackingBySrcIPV4 enables ingress tracking by src ipv4.
func (f *Flow) WithIngressTrackingBySrcIPV4(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.SrcIpv4 = enable
	return f
}

// WithIngressTrackingByDstIPV4 enables ingress tracking by dst ipv4.
func (f *Flow) WithIngressTrackingByDstIPV4(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.DstIpv4 = enable
	return f
}

// WithIngressTrackingBySrcIPV6 enables ingress tracking by src ipv6.
func (f *Flow) WithIngressTrackingBySrcIPV6(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.SrcIpv6 = enable
	return f
}

// WithIngressTrackingByDstIPV6 enables ingress tracking by dst ipv6.
func (f *Flow) WithIngressTrackingByDstIPV6(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.DstIpv6 = enable
	return f
}

// WithIngressTrackingByVLANID enables ingress tracking by VLAN ID.
func (f *Flow) WithIngressTrackingByVLANID(enable bool) *Flow {
	if f.pb.IngressTrackingFilters == nil {
		f.pb.IngressTrackingFilters = &opb.Flow_IngressTrackingFilters{}
	}
	f.pb.IngressTrackingFilters.VlanId = enable
	return f
}

// WithFrameSize sets the frame size for this flow to a fixed value.
// To generate a range of frame sizes, use FrameSizeRange() instead.
func (f *Flow) WithFrameSize(n uint32) *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_Fixed{Fixed: n},
	}
	return f
}

// WithFrameSizeRandom sets the frame size for this flow to a random value in the given range.
func (f *Flow) WithFrameSizeRandom(min, max uint32) *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_Random_{
			Random: &opb.FrameSize_Random{Min: min, Max: max},
		},
	}
	return f
}

// WithFrameSizeIMIXCisco configures this IMIX to use the Cisco preset (64:7, 594:4 and 1518:1).
func (f *Flow) WithFrameSizeIMIXCisco() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_CISCO,
		},
	}
	return f
}

// WithFrameSizeIMIXDefault configures this IMIX to use the default preset (64:7, 570:4, and 1518:1).
func (f *Flow) WithFrameSizeIMIXDefault() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_DEFAULT,
		},
	}
	return f
}

// WithFrameSizeIMIXIPSec configures this IMIX to use the IPSec preset (90:58.67, 92:2, 594:23.66 and 1418:15.67).
func (f *Flow) WithFrameSizeIMIXIPSec() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_IPSEC,
		},
	}
	return f
}

// WithFrameSizeIMIXIPv6 configures this IMIX to use the IPv6 IMIX preset (60:58.67, 496:2, 594:23.66 and 1518:15.67).
func (f *Flow) WithFrameSizeIMIXIPv6() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_IPV6,
		},
	}
	return f
}

// Transmission returns the transmission configuration for this flow.
// By default the transmission will have the following properties:
// pattern: continuous
// min packet gap: 12 bytes.
func (f *Flow) Transmission() *Transmission {
	if f.pb.Transmission == nil {
		f.pb.Transmission = &opb.Transmission{
			Pattern:     opb.Transmission_CONTINUOUS,
			MinGapBytes: 12,
		}
	}
	return &Transmission{pb: f.pb.Transmission}
}

// WithPatternContinuous configures the transmission to send packets continuously.
func (t *Transmission) WithPatternContinuous() *Transmission {
	t.pb.Pattern = opb.Transmission_CONTINUOUS
	return t
}

// WithPatternBurst configures the transmission to send packets in bursts.
func (t *Transmission) WithPatternBurst() *Transmission {
	t.pb.Pattern = opb.Transmission_BURST
	return t
}

// WithPatternFixedPacketCount configures the transmission to send a specified number of packets.
func (t *Transmission) WithPatternFixedPacketCount(packetsCount uint32) *Transmission {
	t.pb.Pattern = opb.Transmission_FIXED_FRAME_COUNT
	t.pb.FrameCount = packetsCount
	return t
}

// WithPatternFixedDuration configures the transmission to send a specified number of seconds.
func (t *Transmission) WithPatternFixedDuration(secs uint32) *Transmission {
	t.pb.Pattern = opb.Transmission_FIXED_DURATION
	t.pb.DurationSecs = secs
	return t
}

// WithMinGapBytes sets the minimum gap between transmitted packets.
func (t *Transmission) WithMinGapBytes(minPacketGapBytes uint32) *Transmission {
	t.pb.MinGapBytes = minPacketGapBytes
	return t
}

// WithPacketsPerBurst sets the number of packets to transmit per burst.
// May only be set when the transmission pattern is burst.
func (t *Transmission) WithPacketsPerBurst(packetsPerBurst uint32) *Transmission {
	t.pb.PacketsPerBurst = packetsPerBurst
	return t
}

// WithInterburstGapNanoseconds sets the gap (in nanoseconds) between transmission bursts.
// May only be set when the transmission pattern is burst.
func (t *Transmission) WithInterburstGapNanoseconds(nanoseconds uint32) *Transmission {
	t.pb.InterburstGap = &opb.Transmission_Nanoseconds{Nanoseconds: nanoseconds}
	return t
}

// WithInterburstGapBytes sets the gap (in bytes) between transmission bursts.
// May only be set when the transmission pattern is burst.
func (t *Transmission) WithInterburstGapBytes(bytes uint32) *Transmission {
	t.pb.InterburstGap = &opb.Transmission_Bytes{Bytes: bytes}
	return t
}

// WithFrameSizeIMIXRPRQuadmodal configures this IMIX to use the RPR Quadmodal preset (64:50, 512:15, 1518:15 and 9216:20).
func (f *Flow) WithFrameSizeIMIXRPRQuadmodal() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_RPR_QUADMODAL,
		},
	}
	return f
}

// WithFrameSizeIMIXRPRTrimodal configures this IMIX to use the RPR Trimodal preset (64:60, 512:20 and 1518:20).
func (f *Flow) WithFrameSizeIMIXRPRTrimodal() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_RPR_TRIMODAL,
		},
	}
	return f
}

// WithFrameSizeIMIXStandard configures this IMIX to use the Standard preset (58:58.67, 62:2, 594:23.66 and 1518:15.67).
func (f *Flow) WithFrameSizeIMIXStandard() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_STANDARD,
		},
	}
	return f
}

// WithFrameSizeIMIXTCP configures this IMIX to use the TCP preset (90:58.67, 92:2, 594:23.66 and 1518:15.67).
func (f *Flow) WithFrameSizeIMIXTCP() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_TCP,
		},
	}
	return f
}

// WithFrameSizeIMIXTolly configures this IMIX to use the Tolly preset (64:55, 78:5, 576:17 and 1518:23).
func (f *Flow) WithFrameSizeIMIXTolly() *Flow {
	f.pb.FrameSize = &opb.FrameSize{
		Type: &opb.FrameSize_ImixPreset_{
			ImixPreset: opb.FrameSize_IMIX_TOLLY,
		},
	}
	return f
}

// FrameSizeIMIXCustom configures this IMIX to use custom size and weight pairs.
func (f *Flow) FrameSizeIMIXCustom() *IMIXCustom {
	if f.pb.FrameSize.GetImixCustom() == nil {
		f.pb.FrameSize = &opb.FrameSize{
			Type: &opb.FrameSize_ImixCustom_{
				ImixCustom: &opb.FrameSize_ImixCustom{},
			},
		}
	}
	return &IMIXCustom{pb: f.pb.FrameSize.GetImixCustom()}
}

// AddEntry adds a custom IMIX entry.
func (ic *IMIXCustom) AddEntry(size, weight uint32) *IMIXCustom {
	ic.pb.Entries = append(ic.pb.Entries, &opb.FrameSize_ImixCustomEntry{
		Size:   size,
		Weight: weight,
	})
	return ic
}

// WithConvergenceTracking enables convergence tracking for the flow.
func (f *Flow) WithConvergenceTracking(enable bool) *Flow {
	f.pb.ConvergenceTracking = true
	return f
}
