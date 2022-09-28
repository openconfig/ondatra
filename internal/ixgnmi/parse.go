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
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding/ixweb"
)

// rowParser encapsulates per-key parsing metadata, including custom parsing
// for some int-valued keys.
type rowParser struct {
	nameKey                     string
	strKeys, intKeys, floatKeys []string
	customIntKeys               map[string]func(string) (uint64, error)
}

func (p *rowParser) parseRow(r ixweb.StatRow) (map[string]*uint64, map[string]*float32, error) {
	name := r[p.nameKey]
	if name == "" {
		return nil, nil, fmt.Errorf("row %v did not include value for required key %q", r, p.nameKey)
	}
	lookup := func(key string) string {
		v, ok := r[key]
		// Warn if the stat column is present but empty.
		if ok && v == "" {
			log.Warningf("got empty stat %q for key %q", key, name)
		}
		return v
	}

	for _, k := range p.strKeys {
		lookup(k) // log only
	}
	ints := make(map[string]*uint64)
	for _, k := range p.intKeys {
		v := lookup(k)
		if v == "" {
			continue
		}
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid value %q for int stat %q for key %q: %w", v, k, name, err)
		}
		ints[k] = &i
	}
	for k, parseFn := range p.customIntKeys {
		v := lookup(k)
		if v == "" {
			continue
		}
		i, err := parseFn(v)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid value %q for int stat %q for key %q: %w", v, k, name, err)
		}
		ints[k] = &i
	}
	floats := make(map[string]*float32)
	for _, k := range p.floatKeys {
		v := lookup(k)
		if v == "" {
			continue
		}
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid value %q for float stat %s for key %s: %w", v, k, name, err)
		}
		f32 := float32(f)
		floats[k] = &f32
	}
	return ints, floats, nil
}

type portRow struct {
	PortName, LineSpeed, LinkState                string
	FramesTx, FramesRx, BytesTx, BytesRx, CRCErrs *uint64
	TxRate, RxRate                                *float32
}

func (r *portRow) String() string {
	return fmt.Sprintf("%+v", *r)
}

func parsePortStats(t ixweb.StatTable) ([]*portRow, error) {
	const (
		nameKey      = "Port Name"
		lineSpeedKey = "Line Speed"
		linkStateKey = "Link State"
		framesTxKey  = "Frames Tx."
		framesRxKey  = "Valid Frames Rx."
		bytesTxKey   = "Bytes Tx."
		bytesRxKey   = "Bytes Rx."
		txRateKey    = "Tx. Rate (bps)"
		rxRateKey    = "Rx. Rate (bps)"
		crcErrsKey   = "CRC Errors"
	)
	rp := &rowParser{
		nameKey:   nameKey,
		strKeys:   []string{lineSpeedKey, linkStateKey},
		intKeys:   []string{framesTxKey, framesRxKey, bytesTxKey, bytesRxKey, crcErrsKey},
		floatKeys: []string{txRateKey, rxRateKey},
	}

	var portRows []*portRow
	for _, row := range t {
		intVals, floatVals, err := rp.parseRow(row)
		if err != nil {
			return nil, err
		}
		portRows = append(portRows, &portRow{
			PortName:  row[nameKey],
			LineSpeed: row[lineSpeedKey],
			LinkState: row[linkStateKey],
			FramesTx:  intVals[framesTxKey],
			FramesRx:  intVals[framesRxKey],
			BytesTx:   intVals[bytesTxKey],
			BytesRx:   intVals[bytesRxKey],
			CRCErrs:   intVals[crcErrsKey],
			TxRate:    floatVals[txRateKey],
			RxRate:    floatVals[rxRateKey],
		})
	}
	return portRows, nil
}

type portCPURow struct {
	PortName                         string
	TotalMemory, FreeMemory, CPULoad *uint64
}

func (r *portCPURow) String() string {
	return rowString(r)
}

func parsePortCPUStats(t ixweb.StatTable) ([]*portCPURow, error) {
	const (
		portNameKey = "Port Name"
		totalMemKey = "Total Memory(KB)"
		freeMemKey  = "Free Memory(KB)"
		cpuLoadKey  = "%CPU Load"
	)
	rp := &rowParser{
		nameKey: portNameKey,
		intKeys: []string{totalMemKey, freeMemKey, cpuLoadKey},
	}

	var portCPURows []*portCPURow
	for _, row := range t {
		intVals, _, err := rp.parseRow(row)
		if err != nil {
			return nil, err
		}
		portCPURows = append(portCPURows, &portCPURow{
			PortName:    row[portNameKey],
			TotalMemory: intVals[totalMemKey],
			FreeMemory:  intVals[freeMemKey],
			CPULoad:     intVals[cpuLoadKey],
		})
	}
	return portCPURows, nil
}

type flowRow struct {
	TrafficItem                                       string
	RxBytes, TxFrames, RxFrames, MPLSLabel, VLANID    *uint64
	LossPct, TxRate, RxRate, TxFrameRate, RxFrameRate *float32
	// Optional ingress-tracking fields.
	RxPort, TxPort, SrcIPv4, DstIPv4, SrcIPv6, DstIPv6 string
	// Optional convergence stats fields.
	FirstPacketTime, ConvergenceTime *uint64
}

func (r *flowRow) String() string {
	return rowString(r)
}

func parseFlowStats(t ixweb.StatTable) ([]*flowRow, error) {
	return parseFlowStatsHelp(t, "")
}

func parseFlowStatsHelp(t ixweb.StatTable, overrideNameKey string) ([]*flowRow, error) {
	const (
		trafficItemKey = "Traffic Item"
		txFramesKey    = "Tx Frames"
		rxFramesKey    = "Rx Frames"
		lossPctKey     = "Loss %"
		txFrameRateKey = "Tx Frame Rate"
		rxFrameRateKey = "Rx Frame Rate"
		rxBytesKey     = "Rx Bytes"
		txRateKey      = "Tx Rate (bps)"
		rxRateKey      = "Rx Rate (bps)"
		// Optional ingress tracking fields.
		rxPortKey    = "Rx Port"
		txPortKey    = "Tx Port"
		mplsLabelKey = "MPLS:Label Value"
		srcIPv4Key   = "IPv4 :Source Address"
		dstIPv4Key   = "IPv4 :Destination Address"
		srcIPv6Key   = "IPv6 :Source Address"
		dstIPv6Key   = "IPv6 :Destination Address"
		vlanIDKey    = "VLAN:VLAN-ID"
		// Optional convergence tracking fields
		firstTimestampKey = "First TimeStamp"
		rampUpTimeUsKey   = "Ramp-up Convergence Time (us)"
	)
	nameKey := trafficItemKey
	if overrideNameKey != "" {
		nameKey = overrideNameKey
	}
	rp := &rowParser{
		nameKey:   nameKey,
		strKeys:   []string{rxPortKey, txPortKey, srcIPv4Key, dstIPv4Key, srcIPv6Key, dstIPv6Key},
		intKeys:   []string{rxBytesKey, txFramesKey, rxFramesKey, mplsLabelKey, vlanIDKey},
		floatKeys: []string{lossPctKey, txRateKey, rxRateKey, txFrameRateKey, rxFrameRateKey},
		customIntKeys: map[string]func(string) (uint64, error){
			firstTimestampKey: func(ts string) (uint64, error) {
				fmtErr := func(ts string) error { return fmt.Errorf("bad format for timestamp %q (expected HH:MM:SS.NNN)", ts) }
				// Expected format: HH:MM:SS.NNN
				hms := strings.Split(ts, ":")
				if len(hms) != 3 {
					return 0, fmtErr(ts)
				}
				sms := strings.Split(hms[2], ".")
				if len(sms) != 2 {
					return 0, fmtErr(ts)
				}
				hrs, err := strconv.ParseUint(hms[0], 10, 64)
				if err != nil {
					return 0, err
				}
				min, err := strconv.ParseUint(hms[1], 10, 64)
				if err != nil {
					return 0, err
				}
				s, err := strconv.ParseUint(sms[0], 10, 64)
				if err != nil {
					return 0, err
				}
				ms, err := strconv.ParseUint(sms[1], 10, 64)
				if err != nil {
					return 0, err
				}
				return uint64(time.Hour)*hrs + uint64(time.Minute)*min + uint64(time.Second)*s + uint64(time.Millisecond)*ms, nil
			},
			rampUpTimeUsKey: func(ts string) (uint64, error) {
				tsUs, err := strconv.ParseUint(strings.ReplaceAll(ts, ",", ""), 10, 64)
				if err != nil {
					return 0, err
				}
				return uint64(time.Microsecond) * tsUs, nil
			},
		},
	}

	var flowRows []*flowRow
	for _, row := range t {
		intVals, floatVals, err := rp.parseRow(row)
		if err != nil {
			return nil, err
		}
		flowRows = append(flowRows, &flowRow{
			TrafficItem: row[trafficItemKey],
			RxBytes:     intVals[rxBytesKey],
			TxFrames:    intVals[txFramesKey],
			RxFrames:    intVals[rxFramesKey],
			LossPct:     floatVals[lossPctKey],
			TxRate:      floatVals[txRateKey],
			RxRate:      floatVals[rxRateKey],
			TxFrameRate: floatVals[txFrameRateKey],
			RxFrameRate: floatVals[rxFrameRateKey],
			// Optional ingress tracking fields.
			RxPort:    row[rxPortKey],
			TxPort:    row[txPortKey],
			MPLSLabel: intVals[mplsLabelKey],
			SrcIPv4:   row[srcIPv4Key],
			DstIPv4:   row[dstIPv4Key],
			SrcIPv6:   row[srcIPv6Key],
			DstIPv6:   row[dstIPv6Key],
			VLANID:    intVals[vlanIDKey],
			// Optional convergence stats.
			FirstPacketTime: intVals[firstTimestampKey],
			ConvergenceTime: intVals[rampUpTimeUsKey],
		})
	}
	return flowRows, nil
}

type egressRow struct {
	filter string
	*flowRow
}

func (r *egressRow) String() string {
	return fmt.Sprintf("{Filter:%s, FlowRow:%s}", r.filter, r.flowRow)
}

func parseEgressStats(t ixweb.StatTable) ([]*egressRow, error) {
	const egressKey = "Egress Tracking"
	flowRows, err := parseFlowStatsHelp(t, egressKey)
	if err != nil {
		return nil, err
	}
	var egressRows []*egressRow
	for i, row := range t {
		egressRows = append(egressRows, &egressRow{
			flowRow: flowRows[i],
			filter:  row[egressKey],
		})
	}
	return egressRows, nil
}

func rowString(row any) string {
	var fields []string
	val := reflect.ValueOf(row).Elem()
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		var valStr string
		if fieldVal := val.Field(i); !fieldVal.IsZero() {
			if fieldVal.Type().Kind() == reflect.Ptr {
				fieldVal = fieldVal.Elem()
			}
			valStr = fmt.Sprintf("%v", fieldVal.Interface())
		}
		fields = append(fields, fmt.Sprintf("%s:%s", fieldName, valStr))
	}
	return fmt.Sprintf("{%s}", strings.Join(fields, ", "))
}
