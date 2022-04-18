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

	log "github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/openconfig/ondatra/binding/ixweb"
)

func parseRow(r ixweb.StatRow, nameKey string, strKeys []string, intKeys []string, floatKeys []string) (map[string]*uint64, map[string]*float32, error) {
	name := r[nameKey]
	if name == "" {
		return nil, nil, errors.Errorf("row %v did not include value for required key %q", r, nameKey)
	}
	lookup := func(key string) string {
		v, ok := r[key]
		// Warn if the stat column is present but empty.
		if ok && v == "" {
			log.Warningf("got empty stat %q for key %q", key, name)
		}
		return v
	}

	for _, k := range strKeys {
		lookup(k) // log only
	}
	ints := make(map[string]*uint64)
	for _, k := range intKeys {
		v := lookup(k)
		if v == "" {
			continue
		}
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "invalid value %q for int stat %q for key %q", v, k, name)
		}
		ints[k] = &i
	}
	floats := make(map[string]*float32)
	for _, k := range floatKeys {
		v := lookup(k)
		if v == "" {
			continue
		}
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "invalid value %q for float stat %s for key %s", v, k, name)
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
	strKeys := []string{lineSpeedKey, linkStateKey}
	intKeys := []string{framesTxKey, framesRxKey, bytesTxKey, bytesRxKey, crcErrsKey}
	floatKeys := []string{txRateKey, rxRateKey}

	var portRows []*portRow
	for _, row := range t {
		intVals, floatVals, err := parseRow(row, nameKey, strKeys, intKeys, floatKeys)
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
	var strKeys []string
	intKeys := []string{totalMemKey, freeMemKey, cpuLoadKey}
	var floatKeys []string

	var portCPURows []*portCPURow
	for _, row := range t {
		intVals, _, err := parseRow(row, portNameKey, strKeys, intKeys, floatKeys)
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
	)
	strKeys := []string{rxPortKey, txPortKey, srcIPv4Key, dstIPv4Key, srcIPv6Key, dstIPv6Key}
	intKeys := []string{rxBytesKey, txFramesKey, rxFramesKey, mplsLabelKey, vlanIDKey}
	floatKeys := []string{lossPctKey, txRateKey, rxRateKey, txFrameRateKey, rxFrameRateKey}

	nameKey := trafficItemKey
	if overrideNameKey != "" {
		nameKey = overrideNameKey
	}

	var flowRows []*flowRow
	for _, row := range t {
		intVals, floatVals, err := parseRow(row, nameKey, strKeys, intKeys, floatKeys)
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

func rowString(row interface{}) string {
	var fields []string
	val := reflect.ValueOf(row).Elem()
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		var valStr string
		if fieldVal := val.Field(i); !fieldVal.IsZero() {
			if fieldVal.Type().Kind() == reflect.Pointer {
				fieldVal = fieldVal.Elem()
			}
			valStr = fmt.Sprintf("%v", fieldVal.Interface())
		}
		fields = append(fields, fmt.Sprintf("%s:%s", fieldName, valStr))
	}
	return fmt.Sprintf("{%s}", strings.Join(fields, ", "))
}
