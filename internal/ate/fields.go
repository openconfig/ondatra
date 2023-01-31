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

package ate

import (
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"strings"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	// To be stubbed out by tests.
	randUInt32 = rand.Uint32
	one        = big.NewInt(1)
)

type addrType string

const (
	mac48len = 6

	mac48AddrType addrType = "MAC-48"
	ipv4AddrType  addrType = "IPv4"
	ipv6AddrType  addrType = "IPv6"
)

func uintToStr(v uint32) *string {
	s := strconv.FormatUint(uint64(v), 10)
	return &s
}

func uintToHexStr(v uint32) *string {
	s := strconv.FormatUint(uint64(v), 16)
	return &s
}

func setSingleValue(field *ixconfig.TrafficTrafficItemConfigElementStackField, val *string) {
	field.ValueType = ixconfig.String("singleValue")
	field.Auto = ixconfig.Bool(false)
	field.SingleValue = val
}

func setRandomRange(field *ixconfig.TrafficTrafficItemConfigElementStackField, min, max, step *string, count uint32) {
	field.ValueType = ixconfig.String("repeatableRandomRange")
	field.Auto = ixconfig.Bool(false)
	field.FullMesh = ixconfig.Bool(false)
	field.Seed = uintToStr(randUInt32())
	field.CountValue = uintToStr(count)
	field.MinValue = min
	field.MaxValue = max
	field.StepValue = step
}

func setIncrement(field *ixconfig.TrafficTrafficItemConfigElementStackField, start, step *string, count uint32) {
	field.ValueType = ixconfig.String("increment")
	field.Auto = ixconfig.Bool(false)
	field.FullMesh = ixconfig.Bool(false)
	field.CountValue = uintToStr(count)
	field.StartValue = start
	field.StepValue = step
}

func setList(field *ixconfig.TrafficTrafficItemConfigElementStackField, vals []string) {
	field.ValueType = ixconfig.String("valueList")
	field.Auto = ixconfig.Bool(false)
	field.FullMesh = ixconfig.Bool(false)
	field.ValueList = vals
}

func setUintRangeField(field *ixconfig.TrafficTrafficItemConfigElementStackField, r *opb.UIntRange) error {
	if r == nil {
		field.Auto = ixconfig.Bool(true)
		return nil
	}
	if r.GetCount() == 0 {
		return fmt.Errorf("count in range is not set or zero")
	}
	if r.GetMin() > r.GetMax() {
		return fmt.Errorf("min %d in range is greater than max %d", r.GetMin(), r.GetMax())
	}

	width := r.GetMax() - r.GetMin()
	if r.GetStep() == 0 {
		r.Step = 1
		maxStep := (width + 1) / r.GetCount()
		if maxStep == 0 {
			return fmt.Errorf(
				"count %d in range cannot fit in [%d, %d]",
				r.GetCount(), r.GetMin(), r.GetMax())
		}
		if !r.GetRandom() {
			r.Step = maxStep
		}
	} else {
		maxCount := (width / r.GetStep()) + 1
		if r.GetCount() > maxCount {
			return fmt.Errorf(
				"count %d with step %d cannot fit in [%d, %d]",
				r.GetCount(), r.GetStep(), r.GetMin(), r.GetMax())
		}
	}

	if r.GetRandom() {
		setRandomRange(field, uintToStr(r.GetMin()), uintToStr(r.GetMax()), uintToStr(r.GetStep()), r.GetCount())
	} else if r.GetCount() == 1 {
		setSingleValue(field, uintToStr(r.GetMin()))
	} else {
		setIncrement(field, uintToStr(r.GetMin()), uintToStr(r.GetStep()), r.GetCount())
	}
	return nil
}

func setAddrRangeField(field *ixconfig.TrafficTrafficItemConfigElementStackField, t addrType, r *opb.AddressRange) error {
	step, err := addrRangeToStep(r, t)
	if err != nil {
		return err
	}

	if r.GetRandom() {
		setRandomRange(field, ixconfig.String(r.GetMin()), ixconfig.String(r.GetMax()), ixconfig.String(step), r.GetCount())
	} else if r.GetCount() == 1 {
		setSingleValue(field, ixconfig.String(r.GetMin()))
	} else {
		setIncrement(field, ixconfig.String(r.GetMin()), ixconfig.String(step), r.GetCount())
	}
	return nil
}

func addrRangeToStep(r *opb.AddressRange, t addrType) (string, error) {
	stepInt, err := addrRangeToStepInt(r, t)
	if err != nil {
		return "", err
	}
	step, err := toAddr(stepInt, t)
	if err != nil {
		return "", fmt.Errorf("could not convert step value %v to address of type %q: %w", stepInt, t, err)
	}
	return step, nil
}

func addrRangeToStepInt(r *opb.AddressRange, t addrType) (*big.Int, error) {
	if r.GetCount() == 0 {
		return nil, fmt.Errorf("count in range is not set or zero")
	}
	min, err := parseAddr(r.GetMin(), t)
	if err != nil {
		return nil, fmt.Errorf("min %q in range is invalid: %w", r.GetMin(), err)
	}
	max, err := parseAddr(r.GetMax(), t)
	if err != nil {
		return nil, fmt.Errorf("max %q in range is invalid: %w", r.GetMax(), err)
	}
	if min.Cmp(max) > 0 {
		return nil, fmt.Errorf("min %q in range is greater than max %q", r.GetMin(), r.GetMax())
	}

	count := big.NewInt(int64(r.GetCount()))
	width := new(big.Int)
	width.Set(max).Sub(width, min)

	if r.GetStep() == "" {
		maxStep := width.Add(width, one).Div(width, count)
		if maxStep.Sign() == 0 {
			return nil, fmt.Errorf(
				"count %d in range cannot fit in [%q, %q]",
				r.GetCount(), r.GetMin(), r.GetMax())
		}
		if r.GetRandom() {
			return one, nil
		}
		return maxStep, nil
	}

	step, err := parseAddr(r.GetStep(), t)
	if err != nil {
		return nil, fmt.Errorf("step %q in range is invalid: %w", r.GetStep(), err)
	}
	if step.Sign() == 0 {
		return nil, fmt.Errorf("step %q in range is zero", r.GetStep())
	}
	maxCount := width.Div(width, step).Add(width, one)
	if count.Cmp(maxCount) > 0 {
		return nil, fmt.Errorf(
			"count %d with step %q in range cannot fit in [%q, %q]",
			r.GetCount(), r.GetStep(), r.GetMin(), r.GetMax())
	}
	return step, nil
}

func parseAddr(s string, t addrType) (*big.Int, error) {
	var b []byte
	var err error

	switch t {
	case mac48AddrType:
		b, err = net.ParseMAC(s)
		if err != nil {
			return nil, fmt.Errorf("not a MAC address: %q", s)
		}
		if len(b) != mac48len {
			return nil, fmt.Errorf("MAC address not in MAC-48 format: %q", s)
		}
	case ipv4AddrType:
		i, isV6 := parseIP(s)
		if i == nil || isV6 {
			return nil, fmt.Errorf("not an IPv4 address: %q", s)
		}
		b = i.To4()
	case ipv6AddrType:
		i, isV6 := parseIP(s)
		if i == nil || !isV6 {
			return nil, fmt.Errorf("not an IPv6 address: %q", s)
		}
		b = i
	default:
		return nil, fmt.Errorf("unknown address type: %q", t)
	}

	return new(big.Int).SetBytes(b), nil
}

// Returns parsed IP and whether the address is a V6 address.
func parseIP(s string) (net.IP, bool) {
	ip := net.ParseIP(s)
	return ip, ip != nil && strings.Contains(s, ":")
}

func toAddr(i *big.Int, t addrType) (string, error) {
	switch t {
	case mac48AddrType:
		b, err := padToLen(i.Bytes(), mac48len)
		if err != nil {
			return "", err
		}
		return net.HardwareAddr(b).String(), nil
	case ipv4AddrType:
		b, err := padToLen(i.Bytes(), net.IPv4len)
		if err != nil {
			return "", err
		}
		return net.IP(b).String(), nil
	case ipv6AddrType:
		b, err := padToLen(i.Bytes(), net.IPv6len)
		if err != nil {
			return "", err
		}
		return net.IP(b).String(), nil
	default:
		return "", fmt.Errorf("unknown address type: %q", t)
	}
}

func padToLen(b []byte, toLen int) ([]byte, error) {
	padLen := toLen - len(b)
	if padLen < 0 {
		return nil, fmt.Errorf("slice %v already longer than %d", b, toLen)
	}
	if padLen > 0 {
		b = append(make([]byte, padLen), b...)
	}
	return b, nil
}
