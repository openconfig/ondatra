package ondatra

import (
	opb "github.com/openconfig/ondatra/proto"
)

const (
	maxFlowLabel uint32 = (1 << 20) - 1
	maxPort      uint32 = (1 << 16) - 1
)

// UIntRange is a range of unsigned integers.
type UIntRange struct {
	pb *opb.UIntRange
}

// WithMin sets the minimum value of the range.
func (r *UIntRange) WithMin(min uint32) *UIntRange {
	r.pb.Min = min
	return r
}

// WithMax sets the maximum value of the range.
func (r *UIntRange) WithMax(max uint32) *UIntRange {
	r.pb.Max = max
	return r
}

// WithStep sets the distance between adjacent values in the range; must not be zero.
func (r *UIntRange) WithStep(step uint32) *UIntRange {
	r.pb.Step = step
	return r
}

// WithCount sets the number of values in the range; must not be zero.
func (r *UIntRange) WithCount(count uint32) *UIntRange {
	r.pb.Count = count
	return r
}

// WithRandom sets the values in the range to be chosen randomly.
func (r *UIntRange) WithRandom() *UIntRange {
	r.pb.Random = true
	return r
}

// AddressRange is a range of addresses.
type AddressRange struct {
	AddressIncRange
}

// WithMin sets the minimum value of the range.
func (r *AddressRange) WithMin(min string) *AddressRange {
	r.AddressIncRange.WithMin(min)
	return r
}

// WithMax sets the maximum value of the range.
func (r *AddressRange) WithMax(max string) *AddressRange {
	r.AddressIncRange.WithMax(max)
	return r
}

// WithCount sets the number of values in the range; must not be zero.
func (r *AddressRange) WithCount(count uint32) *AddressRange {
	r.AddressIncRange.WithCount(count)
	return r
}

// WithStep sets the step address between values in the range.
// If not specified, the default depends on whether the range is random.
// For non-random ranges, step defaults to the largest step that fits the specified count of values.
// For random ranges, the step defaults to a single address.
func (r *AddressRange) WithStep(step string) *AddressRange {
	r.AddressIncRange.WithStep(step)
	return r
}

// WithRandom sets the values in the range to be chosen randomly.
func (r *AddressRange) WithRandom() *AddressRange {
	r.pb.Random = true
	return r
}

// AddressIncRange is a range network addresses that increment by a fixed step.
type AddressIncRange struct {
	pb *opb.AddressRange
}

// WithMin sets the minimum value of the range.
func (r *AddressIncRange) WithMin(addr string) *AddressIncRange {
	r.pb.Min = addr
	return r
}

// WithMax sets the maximum value of the range.
func (r *AddressIncRange) WithMax(max string) *AddressIncRange {
	r.pb.Max = max
	return r
}

// WithCount sets the number of values in the range; must not be zero.
func (r *AddressIncRange) WithCount(count uint32) *AddressIncRange {
	r.pb.Count = count
	return r
}

// WithStep sets the step address between values in the range.
// If not specified, defaults to the largest step that fits the specified count of values.
func (r *AddressIncRange) WithStep(step string) *AddressIncRange {
	r.pb.Step = step
	return r
}

// StringIncRange is an range of strings that increments by a fixed step.
type StringIncRange struct {
	pb *opb.StringIncRange
}

// WithStart sets the start of the range.
func (r *StringIncRange) WithStart(start string) *StringIncRange {
	r.pb.Start = start
	return r
}

// WithStep sets the step between values in the range.
func (r *StringIncRange) WithStep(step string) *StringIncRange {
	r.pb.Step = step
	return r
}

// UInt32IncRange is an range of 32-bit integers that increments by a fixed
// step.
type UInt32IncRange struct {
	pb *opb.UInt32IncRange
}

// WithStart sets the start of the range.
func (r *UInt32IncRange) WithStart(start uint32) *UInt32IncRange {
	r.pb.Start = start
	return r
}

// WithStep sets the step between values in the range.
func (r *UInt32IncRange) WithStep(step uint32) *UInt32IncRange {
	r.pb.Step = step
	return r
}

func intRangeSingle(i uint32) *opb.UIntRange {
	return &opb.UIntRange{Min: i, Max: i, Count: 1}
}

func addrRangeSingle(a string) *opb.AddressRange {
	return &opb.AddressRange{Min: a, Max: a, Count: 1}
}

func newPortRange() *opb.UIntRange {
	return &opb.UIntRange{Min: 1, Max: maxPort}
}

func newFlowLabelRange() *opb.UIntRange {
	return &opb.UIntRange{Max: maxFlowLabel}
}

func newMACAddrRange() *opb.AddressRange {
	return &opb.AddressRange{Min: "00:00:00:00:00:01", Max: "FF:FF:FF:FF:FF:FE"}
}

func newIPv4AddrRange() *opb.AddressRange {
	return &opb.AddressRange{Min: "0.0.0.1", Max: "255.255.255.254"}
}

func newIPv6AddrRange() *opb.AddressRange {
	return &opb.AddressRange{Min: "::1", Max: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"}
}
