// Copyright 2020 Google LLC
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

// NewUIntRange returns a new UIntRange.
// Tests should not call this method directly.
func NewUIntRange(pb *opb.UIntRange) *UIntRange {
	return &UIntRange{pb}
}

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

// NewAddressRange returns a new Address Range.
// Tests should not call this method directly.
func NewAddressRange(pb *opb.AddressRange) *AddressRange {
	return &AddressRange{AddressIncRange{pb}}
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

// NewStringIncRange returns a new StringIncRange.
// Tests should not call this method directly.
func NewStringIncRange(pb *opb.StringIncRange) *StringIncRange {
	return &StringIncRange{pb}
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

// NewUInt32IncRange returns a new UInt32IncRange.
// Tests should not call this method directly.
func NewUInt32IncRange(pb *opb.UInt32IncRange) *UInt32IncRange {
	return &UInt32IncRange{pb}
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
