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

package ixconfig

import (
	"fmt"
)

func aliasesToFieldIdx(aliases []string) map[string]int {
	aToIdx := make(map[string]int, len(aliases))
	for i, a := range aliases {
		aToIdx[a] = i
	}
	return aToIdx
}

func newStack(idx int, stackAlias string, fieldAliasToIdx map[string]int) TrafficStack {
	stack := TrafficStack{
		Xpath: &XPath{
			parentXPath: "/traffic/trafficItem[1]/configElement[1]",
			objectName:  "stack",
			alias:       fmt.Sprintf("%s-%d", stackAlias, idx+1),
		},
	}
	xp := stack.XPath().String()
	stack.Field = make([]*TrafficField, len(fieldAliasToIdx))
	for fa, i := range fieldAliasToIdx {
		stack.Field[i] = &TrafficField{
			Xpath: &XPath{
				parentXPath: xp,
				objectName:  "field",
				alias:       fmt.Sprintf("%s.%s", stackAlias, fa),
			},
			ValueList: []string{},
		}
	}
	return stack
}

type HTTP_GETStack TrafficStack

var HTTP_GETAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.Request.Request Method-1",
	"header.Request.Space1-2",
	"header.Request.Request URI-3",
	"header.Request.Space2-4",
	"header.Request.Request version-5",
	"header.Request.CRLF1-6",
	"header.Host-7",
	"header.User-Agent-8",
	"header.Accept-9",
	"header.Accept-Language-10",
	"header.Accept-Encoding-11",
	"header.Accept-Charset-12",
	"header.Keep-Alive-13",
	"header.Connection-14",
	"header.Referer-15",
	"header.CRLF-16",
})

func (s *HTTP_GETStack) RequestRequest_Method() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Request Method-1"]]
}
func (s *HTTP_GETStack) RequestSpace1() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Space1-2"]]
}
func (s *HTTP_GETStack) RequestRequest_URI() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Request URI-3"]]
}
func (s *HTTP_GETStack) RequestSpace2() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Space2-4"]]
}
func (s *HTTP_GETStack) RequestRequest_Version() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Request version-5"]]
}
func (s *HTTP_GETStack) RequestCRLF1() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.CRLF1-6"]]
}
func (s *HTTP_GETStack) Host() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Host-7"]]
}
func (s *HTTP_GETStack) User_Agent() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.User-Agent-8"]]
}
func (s *HTTP_GETStack) Accept() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-9"]]
}
func (s *HTTP_GETStack) Accept_Language() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-Language-10"]]
}
func (s *HTTP_GETStack) Accept_Encoding() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-Encoding-11"]]
}
func (s *HTTP_GETStack) Accept_Charset() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-Charset-12"]]
}
func (s *HTTP_GETStack) Keep_Alive() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Keep-Alive-13"]]
}
func (s *HTTP_GETStack) Connection() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Connection-14"]]
}
func (s *HTTP_GETStack) Referer() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Referer-15"]]
}
func (s *HTTP_GETStack) CRLF() *TrafficField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.CRLF-16"]]
}

func (s *HTTP_GETStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewHTTP_GETStack(idx int) *HTTP_GETStack {
	stack := HTTP_GETStack(newStack(idx, "HTTP_GET", HTTP_GETAliasToFieldIdx))
	return &stack
}

type Icmpv1Stack TrafficStack

var icmpv1AliasToFieldIdx = aliasesToFieldIdx([]string{
	"message.messageType-1",
	"message.codeOptions.destUnreachableCodeOptions-2",
	"message.codeOptions.srcQuenchOption-3",
	"message.codeOptions.infoRequestOption-4",
	"message.codeOptions.infoResponseOption-5",
	"message.codeOptions.timeExceededOptions-6",
	"message.codeOptions.redirectMessageOptions-7",
	"message.icmpChecksum-8",
	"message.next4Bytes.unusedBitsInMsgType3-9",
	"message.next4Bytes.unusedBitsInMsgType4-10",
	"message.next4Bytes.unusedBitsInMsgType11-11",
	"message.next4Bytes.nextFieldsForParameterProblem.pointer-12",
	"message.next4Bytes.nextFieldsForParameterProblem.unused-13",
	"message.next4Bytes.gatewayInternetAddress-14",
})

func (s *Icmpv1Stack) MessageType() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.messageType-1"]]
}
func (s *Icmpv1Stack) CodeOptionsDestUnreachableCodeOptions() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.destUnreachableCodeOptions-2"]]
}
func (s *Icmpv1Stack) CodeOptionsSrcQuenchOption() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.srcQuenchOption-3"]]
}
func (s *Icmpv1Stack) CodeOptionsInfoRequestOption() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.infoRequestOption-4"]]
}
func (s *Icmpv1Stack) CodeOptionsInfoResponseOption() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.infoResponseOption-5"]]
}
func (s *Icmpv1Stack) CodeOptionsTimeExceededOptions() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.timeExceededOptions-6"]]
}
func (s *Icmpv1Stack) CodeOptionsRedirectMessageOptions() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.redirectMessageOptions-7"]]
}
func (s *Icmpv1Stack) IcmpChecksum() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.icmpChecksum-8"]]
}
func (s *Icmpv1Stack) Next4BytesUnusedBitsInMsgType3() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.unusedBitsInMsgType3-9"]]
}
func (s *Icmpv1Stack) Next4BytesUnusedBitsInMsgType4() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.unusedBitsInMsgType4-10"]]
}
func (s *Icmpv1Stack) Next4BytesUnusedBitsInMsgType11() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.unusedBitsInMsgType11-11"]]
}
func (s *Icmpv1Stack) Next4BytesNextFieldsForParameterProblemPointer() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.nextFieldsForParameterProblem.pointer-12"]]
}
func (s *Icmpv1Stack) Next4BytesNextFieldsForParameterProblemUnused() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.nextFieldsForParameterProblem.unused-13"]]
}
func (s *Icmpv1Stack) Next4BytesGatewayInternetAddress() *TrafficField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.gatewayInternetAddress-14"]]
}

func (s *Icmpv1Stack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewIcmpv1Stack(idx int) *Icmpv1Stack {
	stack := Icmpv1Stack(newStack(idx, "icmpv1", icmpv1AliasToFieldIdx))
	return &stack
}

type Icmpv2Stack TrafficStack

var icmpv2AliasToFieldIdx = aliasesToFieldIdx([]string{
	"message.messageType-1",
	"message.codeValue-2",
	"message.icmpChecksum-3",
	"message.identifier-4",
	"message.sequenceNumber-5",
	"message.nextFields.none-6",
	"message.nextFields.none-7",
	"message.nextFields.fieldsForTimeStampMsg.origTmpStmp1-8",
	"message.nextFields.fieldsForTimeStampMsg.rcvTmpStmp1-9",
	"message.nextFields.fieldsForTimeStampMsg.transTmpStmp1-10",
	"message.nextFields.fieldsForTimeStampReply.origTmpStmp2-11",
	"message.nextFields.fieldsForTimeStampReply.rcvTmpStmp2-12",
	"message.nextFields.fieldsForTimeStampReply.transTmpStmp2-13",
	"message.nextFields.none-14",
	"message.nextFields.none-15",
})

func (s *Icmpv2Stack) MessageType() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.messageType-1"]]
}
func (s *Icmpv2Stack) CodeValue() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.codeValue-2"]]
}
func (s *Icmpv2Stack) IcmpChecksum() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.icmpChecksum-3"]]
}
func (s *Icmpv2Stack) Identifier() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.identifier-4"]]
}
func (s *Icmpv2Stack) SequenceNumber() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.sequenceNumber-5"]]
}
func (s *Icmpv2Stack) NextFieldsNone() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-6"]]
}
func (s *Icmpv2Stack) NextFieldsNone_2() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-7"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampMsgOrigTmpStmp1() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampMsg.origTmpStmp1-8"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampMsgRcvTmpStmp1() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampMsg.rcvTmpStmp1-9"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampMsgTransTmpStmp1() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampMsg.transTmpStmp1-10"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampReplyOrigTmpStmp2() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampReply.origTmpStmp2-11"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampReplyRcvTmpStmp2() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampReply.rcvTmpStmp2-12"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampReplyTransTmpStmp2() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampReply.transTmpStmp2-13"]]
}
func (s *Icmpv2Stack) NextFieldsNone_3() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-14"]]
}
func (s *Icmpv2Stack) NextFieldsNone_4() *TrafficField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-15"]]
}

func (s *Icmpv2Stack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewIcmpv2Stack(idx int) *Icmpv2Stack {
	stack := Icmpv2Stack(newStack(idx, "icmpv2", icmpv2AliasToFieldIdx))
	return &stack
}

type Ipv4Stack TrafficStack

var ipv4AliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.version-1",
	"header.headerLength-2",
	"header.priority.raw-3",
	"header.priority.tos.precedence-4",
	"header.priority.tos.delay-5",
	"header.priority.tos.throughput-6",
	"header.priority.tos.reliability-7",
	"header.priority.tos.monetary-8",
	"header.priority.tos.unused-9",
	"header.priority.ds.phb.defaultPHB.defaultPHB-10",
	"header.priority.ds.phb.defaultPHB.unused-11",
	"header.priority.ds.phb.classSelectorPHB.classSelectorPHB-12",
	"header.priority.ds.phb.classSelectorPHB.unused-13",
	"header.priority.ds.phb.assuredForwardingPHB.assuredForwardingPHB-14",
	"header.priority.ds.phb.assuredForwardingPHB.unused-15",
	"header.priority.ds.phb.expeditedForwardingPHB.expeditedForwardingPHB-16",
	"header.priority.ds.phb.expeditedForwardingPHB.unused-17",
	"header.totalLength-18",
	"header.identification-19",
	"header.flags.reserved-20",
	"header.flags.fragment-21",
	"header.flags.lastFragment-22",
	"header.fragmentOffset-23",
	"header.ttl-24",
	"header.protocol-25",
	"header.checksum-26",
	"header.srcIp-27",
	"header.dstIp-28",
	"header.options.nextOption.option.nop-29",
	"header.options.nextOption.option.security.type-30",
	"header.options.nextOption.option.security.length-31",
	"header.options.nextOption.option.security.security-32",
	"header.options.nextOption.option.security.compartments-33",
	"header.options.nextOption.option.security.handling-34",
	"header.options.nextOption.option.security.tcc-35",
	"header.options.nextOption.option.lsrr.type-36",
	"header.options.nextOption.option.lsrr.length-37",
	"header.options.nextOption.option.pointer-38",
	"header.options.nextOption.option.routeData-39",
	"header.options.nextOption.option.ssrr.type-40",
	"header.options.nextOption.option.ssrr.length-41",
	"header.options.nextOption.option.recordRoute.type-42",
	"header.options.nextOption.option.recordRoute.length-43",
	"header.options.nextOption.option.streamId.type-44",
	"header.options.nextOption.option.streamId.length-45",
	"header.options.nextOption.option.streamId.id-46",
	"header.options.nextOption.option.timestamp.type-47",
	"header.options.nextOption.option.timestamp.length-48",
	"header.options.nextOption.option.timestamp.pointer-49",
	"header.options.nextOption.option.timestamp.overflow-50",
	"header.options.nextOption.option.timestamp.flags-51",
	"header.options.nextOption.option.timestamp.pair.address-52",
	"header.options.nextOption.option.timestamp.pair.timestamp-53",
	"header.options.nextOption.option.last-54",
	"header.options.nextOption.option.routerAlert.type-55",
	"header.options.nextOption.option.routerAlert.length-56",
	"header.options.nextOption.option.routerAlert.value-57",
	"header.options.pad-58",
})

func (s *Ipv4Stack) Version() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.version-1"]]
}
func (s *Ipv4Stack) HeaderLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.headerLength-2"]]
}
func (s *Ipv4Stack) PriorityRaw() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.raw-3"]]
}
func (s *Ipv4Stack) PriorityTosPrecedence() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.precedence-4"]]
}
func (s *Ipv4Stack) PriorityTosDelay() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.delay-5"]]
}
func (s *Ipv4Stack) PriorityTosThroughput() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.throughput-6"]]
}
func (s *Ipv4Stack) PriorityTosReliability() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.reliability-7"]]
}
func (s *Ipv4Stack) PriorityTosMonetary() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.monetary-8"]]
}
func (s *Ipv4Stack) PriorityTosUnused() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.unused-9"]]
}
func (s *Ipv4Stack) PriorityDsPhbDefaultPHBDefaultPHB() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.defaultPHB.defaultPHB-10"]]
}
func (s *Ipv4Stack) PriorityDsPhbDefaultPHBUnused() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.defaultPHB.unused-11"]]
}
func (s *Ipv4Stack) PriorityDsPhbClassSelectorPHBClassSelectorPHB() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.classSelectorPHB.classSelectorPHB-12"]]
}
func (s *Ipv4Stack) PriorityDsPhbClassSelectorPHBUnused() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.classSelectorPHB.unused-13"]]
}
func (s *Ipv4Stack) PriorityDsPhbAssuredForwardingPHBAssuredForwardingPHB() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.assuredForwardingPHB.assuredForwardingPHB-14"]]
}
func (s *Ipv4Stack) PriorityDsPhbAssuredForwardingPHBUnused() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.assuredForwardingPHB.unused-15"]]
}
func (s *Ipv4Stack) PriorityDsPhbExpeditedForwardingPHBExpeditedForwardingPHB() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.expeditedForwardingPHB.expeditedForwardingPHB-16"]]
}
func (s *Ipv4Stack) PriorityDsPhbExpeditedForwardingPHBUnused() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.expeditedForwardingPHB.unused-17"]]
}
func (s *Ipv4Stack) TotalLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.totalLength-18"]]
}
func (s *Ipv4Stack) Identification() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.identification-19"]]
}
func (s *Ipv4Stack) FlagsReserved() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.flags.reserved-20"]]
}
func (s *Ipv4Stack) FlagsFragment() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.flags.fragment-21"]]
}
func (s *Ipv4Stack) FlagsLastFragment() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.flags.lastFragment-22"]]
}
func (s *Ipv4Stack) FragmentOffset() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.fragmentOffset-23"]]
}
func (s *Ipv4Stack) Ttl() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.ttl-24"]]
}
func (s *Ipv4Stack) Protocol() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.protocol-25"]]
}
func (s *Ipv4Stack) Checksum() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.checksum-26"]]
}
func (s *Ipv4Stack) SrcIp() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.srcIp-27"]]
}
func (s *Ipv4Stack) DstIp() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.dstIp-28"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionNop() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.nop-29"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityType() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.type-30"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.length-31"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecuritySecurity() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.security-32"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityCompartments() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.compartments-33"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityHandling() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.handling-34"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityTcc() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.tcc-35"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionLsrrType() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.lsrr.type-36"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionLsrrLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.lsrr.length-37"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionPointer() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.pointer-38"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouteData() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routeData-39"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSsrrType() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.ssrr.type-40"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSsrrLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.ssrr.length-41"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRecordRouteType() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.recordRoute.type-42"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRecordRouteLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.recordRoute.length-43"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionStreamIdType() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.streamId.type-44"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionStreamIdLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.streamId.length-45"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionStreamIdId() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.streamId.id-46"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampType() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.type-47"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.length-48"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampPointer() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.pointer-49"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampOverflow() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.overflow-50"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampFlags() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.flags-51"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampPairAddress() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.pair.address-52"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampPairTimestamp() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.pair.timestamp-53"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionLast() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.last-54"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouterAlertType() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routerAlert.type-55"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouterAlertLength() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routerAlert.length-56"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouterAlertValue() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routerAlert.value-57"]]
}
func (s *Ipv4Stack) OptionsPad() *TrafficField {
	return s.Field[ipv4AliasToFieldIdx["header.options.pad-58"]]
}

func (s *Ipv4Stack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewIpv4Stack(idx int) *Ipv4Stack {
	stack := Ipv4Stack(newStack(idx, "ipv4", ipv4AliasToFieldIdx))
	return &stack
}

type Ipv6Stack TrafficStack

var ipv6AliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.versionTrafficClassFlowLabel.version-1",
	"header.versionTrafficClassFlowLabel.trafficClass-2",
	"header.versionTrafficClassFlowLabel.flowLabel-3",
	"header.payloadLength-4",
	"header.nextHeader-5",
	"header.hopLimit-6",
	"header.srcIP-7",
	"header.dstIP-8",
})

func (s *Ipv6Stack) VersionTrafficClassFlowLabelVersion() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.versionTrafficClassFlowLabel.version-1"]]
}
func (s *Ipv6Stack) VersionTrafficClassFlowLabelTrafficClass() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.versionTrafficClassFlowLabel.trafficClass-2"]]
}
func (s *Ipv6Stack) VersionTrafficClassFlowLabelFlowLabel() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.versionTrafficClassFlowLabel.flowLabel-3"]]
}
func (s *Ipv6Stack) PayloadLength() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.payloadLength-4"]]
}
func (s *Ipv6Stack) NextHeader() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.nextHeader-5"]]
}
func (s *Ipv6Stack) HopLimit() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.hopLimit-6"]]
}
func (s *Ipv6Stack) SrcIP() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.srcIP-7"]]
}
func (s *Ipv6Stack) DstIP() *TrafficField {
	return s.Field[ipv6AliasToFieldIdx["header.dstIP-8"]]
}

func (s *Ipv6Stack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewIpv6Stack(idx int) *Ipv6Stack {
	stack := Ipv6Stack(newStack(idx, "ipv6", ipv6AliasToFieldIdx))
	return &stack
}

type LdpHelloStack TrafficStack

var ldpHelloAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.version-1",
	"header.pduLengthinOctets-2",
	"header.lsrID-3",
	"header.labelSpace-4",
	"header.uBit-5",
	"header.type-6",
	"header.length-7",
	"header.messageID-8",
	"header.commonHelloParametersTLV.uBit-9",
	"header.commonHelloParametersTLV.fBit-10",
	"header.commonHelloParametersTLV.type-11",
	"header.commonHelloParametersTLV.length-12",
	"header.commonHelloParametersTLV.holdTime-13",
	"header.commonHelloParametersTLV.tBit-14",
	"header.commonHelloParametersTLV.rBit-15",
	"header.commonHelloParametersTLV.reserved-16",
	"header.optionalParameter.ipv4TransportAddressTLV.uBit-17",
	"header.optionalParameter.ipv4TransportAddressTLV.fBit-18",
	"header.optionalParameter.ipv4TransportAddressTLV.type-19",
	"header.optionalParameter.ipv4TransportAddressTLV.length-20",
	"header.optionalParameter.ipv4TransportAddressTLV.ipv4Address-21",
	"header.optionalParameter.configurationSequenceNumberTLV.uBit-22",
	"header.optionalParameter.configurationSequenceNumberTLV.fBit-23",
	"header.optionalParameter.configurationSequenceNumberTLV.type-24",
	"header.optionalParameter.configurationSequenceNumberTLV.length-25",
	"header.optionalParameter.configurationSequenceNumberTLV.sequenceNumber-26",
	"header.optionalParameter.ipv6TransportAddressTLV.uBit-27",
	"header.optionalParameter.ipv6TransportAddressTLV.fBit-28",
	"header.optionalParameter.ipv6TransportAddressTLV.type-29",
	"header.optionalParameter.ipv6TransportAddressTLV.length-30",
	"header.optionalParameter.ipv6TransportAddressTLV.ipv6Address-31",
})

func (s *LdpHelloStack) Version() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.version-1"]]
}
func (s *LdpHelloStack) PduLengthinOctets() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.pduLengthinOctets-2"]]
}
func (s *LdpHelloStack) LsrID() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.lsrID-3"]]
}
func (s *LdpHelloStack) LabelSpace() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.labelSpace-4"]]
}
func (s *LdpHelloStack) UBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.uBit-5"]]
}
func (s *LdpHelloStack) Type() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.type-6"]]
}
func (s *LdpHelloStack) Length() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.length-7"]]
}
func (s *LdpHelloStack) MessageID() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.messageID-8"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVUBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.uBit-9"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVFBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.fBit-10"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVType() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.type-11"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVLength() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.length-12"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVHoldTime() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.holdTime-13"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVTBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.tBit-14"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVRBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.rBit-15"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVReserved() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.reserved-16"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVUBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.uBit-17"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVFBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.fBit-18"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVType() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.type-19"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVLength() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.length-20"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVIpv4Address() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.ipv4Address-21"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVUBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.uBit-22"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVFBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.fBit-23"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVType() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.type-24"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVLength() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.length-25"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVSequenceNumber() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.sequenceNumber-26"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVUBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.uBit-27"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVFBit() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.fBit-28"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVType() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.type-29"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVLength() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.length-30"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVIpv6Address() *TrafficField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.ipv6Address-31"]]
}

func (s *LdpHelloStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewLdpHelloStack(idx int) *LdpHelloStack {
	stack := LdpHelloStack(newStack(idx, "ldpHello", ldpHelloAliasToFieldIdx))
	return &stack
}

type PimHelloMessageStack TrafficStack

var pimHelloMessageAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.version-1",
	"header.type-2",
	"header.reserved-3",
	"header.headerChecksum-4",
	"header.helloOptionsFields.options.holdtime.type-5",
	"header.helloOptionsFields.options.holdtime.length-6",
	"header.helloOptionsFields.options.holdtime.holdtimesec-7",
	"header.helloOptionsFields.options.lanPruneDelay.type-8",
	"header.helloOptionsFields.options.lanPruneDelay.length-9",
	"header.helloOptionsFields.options.lanPruneDelay.tBit-10",
	"header.helloOptionsFields.options.lanPruneDelay.lanDelay-11",
	"header.helloOptionsFields.options.lanPruneDelay.overrideInterval-12",
	"header.helloOptionsFields.options.drPriority.type-13",
	"header.helloOptionsFields.options.drPriority.length-14",
	"header.helloOptionsFields.options.drPriority.drPriority-15",
	"header.helloOptionsFields.options.generationID.type-16",
	"header.helloOptionsFields.options.generationID.length-17",
	"header.helloOptionsFields.options.generationID.generationID-18",
	"header.helloOptionsFields.options.bidirCapable.type-19",
	"header.helloOptionsFields.options.bidirCapable.length-20",
	"header.helloOptionsFields.options.privateUsageField.type-21",
	"header.helloOptionsFields.options.privateUsageField.length-22",
	"header.helloOptionsFields.options.privateUsageField.value-23",
	"header.helloOptionsFields.options.userDefinedField.type-24",
	"header.helloOptionsFields.options.userDefinedField.length-25",
	"header.helloOptionsFields.options.userDefinedField.value-26",
})

func (s *PimHelloMessageStack) Version() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.version-1"]]
}
func (s *PimHelloMessageStack) Type() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.type-2"]]
}
func (s *PimHelloMessageStack) Reserved() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.reserved-3"]]
}
func (s *PimHelloMessageStack) HeaderChecksum() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.headerChecksum-4"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsHoldtimeType() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.holdtime.type-5"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsHoldtimeLength() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.holdtime.length-6"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsHoldtimeHoldtimesec() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.holdtime.holdtimesec-7"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayType() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.type-8"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayLength() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.length-9"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayTBit() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.tBit-10"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayLanDelay() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.lanDelay-11"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayOverrideInterval() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.overrideInterval-12"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsDrPriorityType() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.drPriority.type-13"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsDrPriorityLength() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.drPriority.length-14"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsDrPriorityDrPriority() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.drPriority.drPriority-15"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsGenerationIDType() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.generationID.type-16"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsGenerationIDLength() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.generationID.length-17"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsGenerationIDGenerationID() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.generationID.generationID-18"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsBidirCapableType() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.bidirCapable.type-19"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsBidirCapableLength() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.bidirCapable.length-20"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsPrivateUsageFieldType() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.privateUsageField.type-21"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsPrivateUsageFieldLength() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.privateUsageField.length-22"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsPrivateUsageFieldValue() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.privateUsageField.value-23"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsUserDefinedFieldType() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.userDefinedField.type-24"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsUserDefinedFieldLength() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.userDefinedField.length-25"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsUserDefinedFieldValue() *TrafficField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.userDefinedField.value-26"]]
}

func (s *PimHelloMessageStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewPimHelloMessageStack(idx int) *PimHelloMessageStack {
	stack := PimHelloMessageStack(newStack(idx, "pimHelloMessage", pimHelloMessageAliasToFieldIdx))
	return &stack
}

type RsvpStack TrafficStack

var rsvpAliasToFieldIdx = aliasesToFieldIdx([]string{
	"rsvpMessege.version-1",
	"rsvpMessege.flag-2",
	"rsvpMessege.messegeType-3",
	"rsvpMessege.rsvpChecksum-4",
	"rsvpMessege.ttl-5",
	"rsvpMessege.reserved-6",
	"rsvpMessege.rsvpLength-7",
	"rsvpMessege.objects.object.bundleMsgOptionalHeader.version-8",
	"rsvpMessege.objects.object.bundleMsgOptionalHeader.flag-9",
	"rsvpMessege.objects.object.bundleMsgOptionalHeader.messegeType-10",
	"rsvpMessege.objects.object.bundleMsgOptionalHeader.rsvpChecksum-11",
	"rsvpMessege.objects.object.bundleMsgOptionalHeader.ttl-12",
	"rsvpMessege.objects.object.bundleMsgOptionalHeader.reserved-13",
	"rsvpMessege.objects.object.bundleMsgOptionalHeader.rsvpLength-14",
	"rsvpMessege.objects.object.objectLength-15",
	"rsvpMessege.objects.object.class-16",
	"rsvpMessege.objects.object.type-17",
	"rsvpMessege.objects.object.objectBody.dataMessege.dataLength-18",
	"rsvpMessege.objects.object.objectBody.dataMessege.data-19",
	"rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.destAddress-20",
	"rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.protocolId-21",
	"rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.flags-22",
	"rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.destPort-23",
	"rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.destAddress-24",
	"rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.protocolId-25",
	"rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.flags-26",
	"rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.destPort-27",
	"rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.destAddress-28",
	"rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.protocolId-29",
	"rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.flags-30",
	"rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.destPort-31",
	"rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.destAddress-32",
	"rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.protocolId-33",
	"rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.flags-34",
	"rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.destPort-35",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.ipv4TunnelEndPointAddress-36",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.reserved-37",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.tunnelId-38",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.extendedTunnelId-39",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.ipv6TunnelEndPointAddress-40",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.reserved-41",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.tunnelId-42",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.extendedTunnelId-43",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.ipv4SessionAddress-44",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.reserved-45",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.flag-46",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.unused-47",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.dscp-48",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.ipv6SessionAddress-49",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.reserved-50",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.flag-51",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.unused-52",
	"rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.dscp-53",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.p2mpId-54",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.reserved-55",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.tunnelId-56",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.extendedTunnelId-57",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.p2mpId-58",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.reserved-59",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.tunnelId-60",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.extendedTunnelId-61",
	"rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType1.ipv4NextPreviousHopAddress-62",
	"rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType1.logicalInterfaceHandle-63",
	"rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType2.ipv6NextPreviousHopAddress-64",
	"rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType2.logicalInterfaceHandle-65",
	"rsvpMessege.objects.object.objectBody.integrityclass4CType1.flags-66",
	"rsvpMessege.objects.object.objectBody.integrityclass4CType1.reserved-67",
	"rsvpMessege.objects.object.objectBody.integrityclass4CType1.keyId-68",
	"rsvpMessege.objects.object.objectBody.integrityclass4CType1.sequenceNumber-69",
	"rsvpMessege.objects.object.objectBody.integrityclass4CType1.msgLength-70",
	"rsvpMessege.objects.object.objectBody.integrityclass4CType1.keyedMessege-71",
	"rsvpMessege.objects.object.objectBody.timevaluesClassClass5CType1.refreshPeriodR-72",
	"rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.ipv4ErrorNodeAddress-73",
	"rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.flags-74",
	"rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.errorCode-75",
	"rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.errorValue-76",
	"rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.ipv6ErrorNodeAddress-77",
	"rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.flags-78",
	"rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.errorCode-79",
	"rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.errorValue-80",
	"rsvpMessege.objects.object.objectBody.ipv4ScopeClassClass7CType1.ipv4SrcAddress-81",
	"rsvpMessege.objects.object.objectBody.ipv6ScopeClassClass7CType2.ipv6SrcAddress-82",
	"rsvpMessege.objects.object.objectBody.styleClassClass8CType1.flags-83",
	"rsvpMessege.objects.object.objectBody.styleClassClass8CType1.reserved-84",
	"rsvpMessege.objects.object.objectBody.styleClassClass8CType1.sharingControl-85",
	"rsvpMessege.objects.object.objectBody.styleClassClass8CType1.senderSelectionControl-86",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.signalType-87",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.rcc-88",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.ncc-89",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.nvc-90",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.multiplier-91",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.transparency-92",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.profile-93",
	"rsvpMessege.objects.object.objectBody.g709Class9CType5.signalType-94",
	"rsvpMessege.objects.object.objectBody.g709Class9CType5.reserved-95",
	"rsvpMessege.objects.object.objectBody.g709Class9CType5.nmc-96",
	"rsvpMessege.objects.object.objectBody.g709Class9CType5.nvc-97",
	"rsvpMessege.objects.object.objectBody.g709Class9CType5.multiplier-98",
	"rsvpMessege.objects.object.objectBody.g709Class9CType5.reserved-99",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType1.ipv4SrcAddress-100",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType1.unused-101",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType1.srcPort-102",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType2.ipv6SrcAddress-103",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType2.unused-104",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType2.srcPort-105",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType3.ipv6SrcAddress-106",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType3.unused-107",
	"rsvpMessege.objects.object.objectBody.filterspecclass10CType3.flowLabel-108",
	"rsvpMessege.objects.object.objectBody.ipv4GPIFILTERSPECClass10CType4.ipv4SrcAddress-109",
	"rsvpMessege.objects.object.objectBody.ipv4GPIFILTERSPECClass10CType4.gpi-110",
	"rsvpMessege.objects.object.objectBody.ipv6GPIFILTERSPECClass10CType5.ipv6SrcAddress-111",
	"rsvpMessege.objects.object.objectBody.ipv6GPIFILTERSPECClass10CType5.gpi-112",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.ipv4TunnelSenderAddress-113",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.unused-114",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.lspID-115",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.ipv6TunnelSenderAddress-116",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.unused-117",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.lspID-118",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.ipv4TunnelSenderAddress-119",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.unused-120",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.lspID-121",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.subGroupOriginatorID-122",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.unused-123",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.subGroupID-124",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.ipv6TunnelSenderAddress-125",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.unused-126",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.lspID-127",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.subGroupOriginatorID-128",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.unused-129",
	"rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.subGroupID-130",
	"rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.ipv4SrcAddress-131",
	"rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.unused-132",
	"rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.srcPort-133",
	"rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.ipv6SrcAddress-134",
	"rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.unused-135",
	"rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.srcPort-136",
	"rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.ipv6SrcAddress-137",
	"rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.unused-138",
	"rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.flowLabel-139",
	"rsvpMessege.objects.object.objectBody.ipv4GPISENDERTEMPLATEClass11CType4.ipv4SrcAddress-140",
	"rsvpMessege.objects.object.objectBody.ipv4GPISENDERTEMPLATEClass11CType4.gpi-141",
	"rsvpMessege.objects.object.objectBody.ipv6GPISENDERTEMPLATEClass11CType5.ipv6SrcAddress-142",
	"rsvpMessege.objects.object.objectBody.ipv6GPISENDERTEMPLATEClass11CType5.gpi-143",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.ipv4TunnelSenderAddress-144",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.unused-145",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.lspID-146",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.ipv6TunnelSenderAddress-147",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.unused-148",
	"rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.lspID-149",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.ipv4TunnelSenderAddress-150",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.unused-151",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.lspID-152",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.subGroupOriginatorID-153",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.unused-154",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.subGroupID-155",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.ipv6TunnelSenderAddress-156",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.unused-157",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.lspID-158",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.subGroupOriginatorID-159",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.unused-160",
	"rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.subGroupID-161",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.versionNumber-162",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.reserved1-163",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.overallLength-164",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.serviceHeader-165",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.zeroBit-166",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.reserved2-167",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.lengthOfService1Data-168",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameterIDTokenBucketTSpec-169",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameter127Flag-170",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameter127Length-171",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.tokenBucketRate-172",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.tokenBucketSize-173",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.peakDataRate-174",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.minimumPolicedUnit-175",
	"rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.maximumPacketSize-176",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.signalType-177",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.rcc-178",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.ncc-179",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.nvc-180",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.multiplier-181",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.transparency-182",
	"rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.profile-183",
	"rsvpMessege.objects.object.objectBody.g709Class12CType5.signalType-184",
	"rsvpMessege.objects.object.objectBody.g709Class12CType5.reserved-185",
	"rsvpMessege.objects.object.objectBody.g709Class12CType5.nmc-186",
	"rsvpMessege.objects.object.objectBody.g709Class12CType5.nvc-187",
	"rsvpMessege.objects.object.objectBody.g709Class12CType5.multiplier-188",
	"rsvpMessege.objects.object.objectBody.g709Class12CType5.reserved-189",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.messageFormatVersionNumber-190",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.reserved-191",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.msgLength-192",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.perServiceHeaderServiceNumber1-193",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.x-194",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.reserved3-195",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.globalBreakBit-196",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID4-197",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter4FlagByte-198",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter4Length-199",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.isHopCnt-200",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID6-201",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter6FlagByte-202",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter6Length-203",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.pathBwEstimate-204",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID8-205",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter8FlagByte-206",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter8Length-207",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.minimumPathLatency-208",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID10-209",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter10FlagByte-210",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter10Length-211",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.composedMTU-212",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.perServiceHeaderServiceNumber2-213",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.xBit-214",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.reserved-215",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.breakBitAndLengthOfPerserviceData-216",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter133ComposedCtot-217",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter133FlagByte-218",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter133Length-219",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.endtoendComposedValueForCCtot-220",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter134ComposedDtot-221",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter134FlagByte-222",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter134Length-223",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.endtoendComposedValueForDDtot-224",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter135ComposedCsum-225",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter135FlagByte-226",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter135Length-227",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.sincelastreshapingPointComposedCCsum-228",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter136ComposedDsum-229",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter136FlagByte-230",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter136Length-231",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.sincelastreshapingPointComposedDDsum-232",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.servicespecificGeneralParameterHeadersvalues.servicespecificGeneralParameterHeadervalue-233",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.perServiceHeaderServiceNumber5-234",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.xBit-235",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.breakBit-236",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.lengthOfPerserviceData-237",
	"rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.servicespecificGeneralParameterHeaders.servicespecificGeneralParameterHeader-238",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.length-239",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policydata-240",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.unused-241",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.dataOffset-242",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.unused-243",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.optionDataLength-244",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.optionData-245",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policyDataLength-246",
	"rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policyData-247",
	"rsvpMessege.objects.object.objectBody.ipv4RESVCONFIRMClass15CType1.ipv4ReceiverAddress-248",
	"rsvpMessege.objects.object.objectBody.ipv6RESVCONFIRMClass15CType2.ipv6ReceiverAddress-249",
	"rsvpMessege.objects.object.objectBody.labelObjectClass16CType1.topLabel-250",
	"rsvpMessege.objects.object.objectBody.labelRequestWithoutLabelRangeClass19CType1.reserved-251",
	"rsvpMessege.objects.object.objectBody.labelRequestWithoutLabelRangeClass19CType1.l3pid-252",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-253",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.l3pid-254",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.mBit-255",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-256",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.minimumVPI-257",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.minimumVCI-258",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-259",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.maximumVPI-260",
	"rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.maximumVCI-261",
	"rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.reserved-262",
	"rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.l3pid-263",
	"rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.res-264",
	"rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.dli-265",
	"rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.minimumDLCI-266",
	"rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.res-267",
	"rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.maximumDLCI-268",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.lBit-269",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.type-270",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.length-271",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.ipv4Address-272",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.prefixLength-273",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.padding-274",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.lBit-275",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.type-276",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.length-277",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.ipv6Address-278",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.prefixLength-279",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.padding-280",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.lBit-281",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.type-282",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.length-283",
	"rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.asNumber-284",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.type-285",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.length-286",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.ipv4Address-287",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.prefixLength-288",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.flags-289",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.type-290",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.length-291",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.ipv6Address-292",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.prefixLength-293",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.flags-294",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.type-295",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.length-296",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.flags-297",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.ctype-298",
	"rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.contentsOfLabelObject-299",
	"rsvpMessege.objects.object.objectBody.helloREQUESTAckClass22CType12.srcInstance-300",
	"rsvpMessege.objects.object.objectBody.helloREQUESTAckClass22CType12.destInstance-301",
	"rsvpMessege.objects.object.objectBody.messageidclass23CType1.flags-302",
	"rsvpMessege.objects.object.objectBody.messageidclass23CType1.epoch-303",
	"rsvpMessege.objects.object.objectBody.messageidclass23CType1.messageIdentifier-304",
	"rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.flags-305",
	"rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.epoch-306",
	"rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.messageIdentifier-307",
	"rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.flags-308",
	"rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.epoch-309",
	"rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.messageidlist.messageIdentifier-310",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.flags-311",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.epoch-312",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.messageidlist.messageIdentifier-313",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.messageidlist.sourceIPAddress-314",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.flags-315",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.epoch-316",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.messageidlist.messageIdentifier-317",
	"rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.messageidlist.sourceIPAddress-318",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.flags-319",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.epoch-320",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.messageIdentifier-321",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.sourceIPAddress-322",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.destinationIPAddress-323",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.flags-324",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.epoch-325",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.messageIdentifier-326",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.sourceIPAddress-327",
	"rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.destinationIPAddress-328",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.maxRSVPhops-329",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.rsvphopcount-330",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.reserved-331",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.mfBit-332",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.requestIdentifier-333",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.pathMTU-334",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.fragmentOffset-335",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.lasthopaddress-336",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.senderTemplateObject.ipv4SrcAddress-337",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.senderTemplateObject.generalizedPortIdentifier-338",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.filterSpecTemplateObject.ipv4SrcAddress-339",
	"rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.filterSpecTemplateObject.generalizedPortIdentifier-340",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.maxRSVPhops-341",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.rsvphopcount-342",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.reserved-343",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.mfBit-344",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.requestIdentifier-345",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.pathMTU-346",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.fragmentOffset-347",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.lasthopaddress-348",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.senderTemplateObject.ipv6SrcAddress-349",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.senderTemplateObject.generalizedPortIdentifier-350",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.filterSpecTemplateObject.ipv6SrcAddress-351",
	"rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.filterSpecTemplateObject.generalizedPortIdentifier-352",
	"rsvpMessege.objects.object.objectBody.diagselectclass33CType1.classType.class-353",
	"rsvpMessege.objects.object.objectBody.diagselectclass33CType1.classType.ctype-354",
	"rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.reserved-355",
	"rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.rPointer-356",
	"rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.nodeAddressList.rsvpNodeAddress-357",
	"rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.reserved-358",
	"rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.rPointer-359",
	"rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.nodeAddressList.rsvpNodeAddress-360",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.dreqArrivalTime-361",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.incomingInterfaceAddress-362",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.outgoingInterfaceAddress-363",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.previousRSVPHopRouterAddress-364",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.dttl-365",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.mBit-366",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.rerr-367",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.k-368",
	"rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.timerValue-369",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.dreqArrivalTime-370",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.incomingInterfaceAddress-371",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.outgoingInterfaceAddress-372",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.previousRSVPHopRouterAddress-373",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.dttl-374",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.mBit-375",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.rerr-376",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.k-377",
	"rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.timerValue-378",
	"rsvpMessege.objects.object.objectBody.s2lsublspipv4Class50CType1.ipv4S2LSubLSPDestinationAddress-379",
	"rsvpMessege.objects.object.objectBody.s2lsublspipv6Class50CType2.ipv6S2LSubLSPDestinationAddress-380",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.lengthbytes-381",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.classNum-382",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.ctype-383",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.plrAddressList.plrID-384",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.plrAddressList.avoidNodeID-385",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.lengthbytes-386",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.classNum-387",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.ctype-388",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.plrAddressList.plrID-389",
	"rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.plrAddressList.avoidNodeID-390",
	"rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.reserved-391",
	"rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.keyId-392",
	"rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.challengeCookie-393",
	"rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.reserved-394",
	"rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapnb-395",
	"rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.reserved-396",
	"rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.exp-397",
	"rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.phbid-398",
	"rsvpMessege.objects.object.objectBody.diffservLLSPclass65CType2.reserved-399",
	"rsvpMessege.objects.object.objectBody.diffservLLSPclass65CType2.psc-400",
	"rsvpMessege.objects.object.objectBody.classtypeclass66CType1.reserved-401",
	"rsvpMessege.objects.object.objectBody.classtypeclass66CType1.ct-402",
	"rsvpMessege.objects.object.objectBody.lsptunnelinterfaceidclass193CType1.lsrId-403",
	"rsvpMessege.objects.object.objectBody.lsptunnelinterfaceidclass193CType1.interfaceID-404",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.lBit-405",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.type-406",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.length-407",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.ipv4Address-408",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.prefixLength-409",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.padding-410",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.lBit-411",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.type-412",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.length-413",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.ipv6Address-414",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.prefixLength-415",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.padding-416",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.lBit-417",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.type-418",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.length-419",
	"rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.asNumber-420",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.type-421",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.length-422",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.ipv4Address-423",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.prefixLength-424",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.flags-425",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.type-426",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.length-427",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.ipv6Address-428",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.prefixLength-429",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.flags-430",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.type-431",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.length-432",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.flags-433",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.ctype-434",
	"rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.contentsOfLabelObject-435",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.lengthbytes-436",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.classNum-437",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.ctype-438",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.setupPrio-439",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.holdingPrio-440",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.hoplimit-441",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.flags-442",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.bandwidth-443",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.includeany-444",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.excludeany-445",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.includeall-446",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.lengthbytes-447",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.classNum-448",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.ctype-449",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.setupPrio-450",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.holdingPrio-451",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.hoplimit-452",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.reserved-453",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.bandwidth-454",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.includeany-455",
	"rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.excludeany-456",
	"rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.setupPrio-457",
	"rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.holdingPrio-458",
	"rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.flags-459",
	"rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.nameLength-460",
	"rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.sessionName-461",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.excludeany-462",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.includeany-463",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.includeall-464",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.setupPrio-465",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.holdingPrio-466",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.flags-467",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.nameLength-468",
	"rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.sessionName-469",
	"rsvpMessege.objects.object.objectBody.atmserviceclassclass227CType1.reserved-470",
	"rsvpMessege.objects.object.objectBody.atmserviceclassclass227CType1.flags-471",
	"rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.lengthbytes-472",
	"rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.classNum-473",
	"rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.ctype-474",
	"rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.resv-475",
	"rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.callOpsFlag-476",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.uBit-477",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.fBit-478",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.sourceIDType-479",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.length-480",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.ipv4Address-481",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.logicalPortId-482",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.uBit-483",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.fBit-484",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.sourceIDType-485",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.length-486",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.ipv6Address-487",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.logicalPortId-488",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.uBit-489",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.fBit-490",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.sourceIDType-491",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.length-492",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.dataLength-493",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.nsap-494",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.logicalPortId-495",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.uBit-496",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.fBit-497",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.destIDType-498",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.length-499",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.ipv4Address-500",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.logicalPortId-501",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.uBit-502",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.fBit-503",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.destIDType-504",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.length-505",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.ipv6Address-506",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.logicalPortId-507",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.uBit-508",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.fBit-509",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.destIDType-510",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.length-511",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.dataLength-512",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.nsap-513",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.logicalPortId-514",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.uBit-515",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.fBit-516",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.egressIDType-517",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.length-518",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.reserved-519",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.lbit-520",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.logicalPortId-521",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.labelLength-522",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.label-523",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.uBit-524",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.fBit-525",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.connectionIDType-526",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.length-527",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.reserved-528",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.cbit-529",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.logicalConnectionId-530",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.uBit-531",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.fBit-532",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.diversityIDType-533",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.length-534",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.localConnectionID-535",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.reserved-536",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.divT-537",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.uBit-538",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.fBit-539",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.contractIDType-540",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.length-541",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.contractID-542",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.uBit-543",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.fBit-544",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.serviceLevelType-545",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.length-546",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.reserved-547",
	"rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.serviceLevel-548",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.lengthbytes-549",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.classNum-550",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.ctype-551",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.type-552",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.resv-553",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.srcLSRAddress-554",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.localId-555",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.type-556",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.resv-557",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.srcLSRAddress-558",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.localId-559",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.type-560",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.resv-561",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.srcLSRAddress-562",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.localId-563",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.type-564",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.resv-565",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.srcLSRAddress-566",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.localId-567",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.type-568",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.resv-569",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressLength-570",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressValue-571",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.localId-572",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.lengthbytes-573",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.classNum-574",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.ctype-575",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.type-576",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.is-577",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.ns-578",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.srcLSRAddress-579",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.localId-580",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.type-581",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.is-582",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.ns-583",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.srcLSRAddress-584",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.localId-585",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.type-586",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.is-587",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.ns-588",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.srcLSRAddress-589",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.localId-590",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.type-591",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.is-592",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.ns-593",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.srcLSRAddress-594",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.localId-595",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.type-596",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.is-597",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.ns-598",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressLength-599",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressValue-600",
	"rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.localId-601",
})

func (s *RsvpStack) Version() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.version-1"]]
}
func (s *RsvpStack) Flag() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.flag-2"]]
}
func (s *RsvpStack) MessegeType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.messegeType-3"]]
}
func (s *RsvpStack) RsvpChecksum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.rsvpChecksum-4"]]
}
func (s *RsvpStack) Ttl() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.ttl-5"]]
}
func (s *RsvpStack) Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.reserved-6"]]
}
func (s *RsvpStack) RsvpLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.rsvpLength-7"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderVersion() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.version-8"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderFlag() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.flag-9"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderMessegeType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.messegeType-10"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderRsvpChecksum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.rsvpChecksum-11"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderTtl() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.ttl-12"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderReserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.reserved-13"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderRsvpLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.rsvpLength-14"]]
}
func (s *RsvpStack) ObjectsObjectObjectLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectLength-15"]]
}
func (s *RsvpStack) ObjectsObjectClass() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.class-16"]]
}
func (s *RsvpStack) ObjectsObjectType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.type-17"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDataMessegeDataLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.dataMessege.dataLength-18"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDataMessegeData() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.dataMessege.data-19"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1DestAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.destAddress-20"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1ProtocolId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.protocolId-21"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.flags-22"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1DestPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.destPort-23"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2DestAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.destAddress-24"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2ProtocolId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.protocolId-25"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.flags-26"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2DestPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.destPort-27"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3DestAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.destAddress-28"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3ProtocolId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.protocolId-29"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.flags-30"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3DestPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.destPort-31"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4DestAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.destAddress-32"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4ProtocolId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.protocolId-33"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.flags-34"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4DestPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.destPort-35"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7Ipv4TunnelEndPointAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.ipv4TunnelEndPointAddress-36"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.reserved-37"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7TunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.tunnelId-38"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7ExtendedTunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.extendedTunnelId-39"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8Ipv6TunnelEndPointAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.ipv6TunnelEndPointAddress-40"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.reserved-41"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8TunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.tunnelId-42"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8ExtendedTunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.extendedTunnelId-43"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Ipv4SessionAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.ipv4SessionAddress-44"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.reserved-45"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Flag() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.flag-46"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.unused-47"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Dscp() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.dscp-48"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Ipv6SessionAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.ipv6SessionAddress-49"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.reserved-50"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Flag() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.flag-51"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.unused-52"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Dscp() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.dscp-53"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13P2mpId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.p2mpId-54"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.reserved-55"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13TunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.tunnelId-56"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13ExtendedTunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.extendedTunnelId-57"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14P2mpId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.p2mpId-58"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.reserved-59"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14TunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.tunnelId-60"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14ExtendedTunnelId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.extendedTunnelId-61"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType1Ipv4NextPreviousHopAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType1.ipv4NextPreviousHopAddress-62"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType1LogicalInterfaceHandle() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType1.logicalInterfaceHandle-63"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType2Ipv6NextPreviousHopAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType2.ipv6NextPreviousHopAddress-64"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType2LogicalInterfaceHandle() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType2.logicalInterfaceHandle-65"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.flags-66"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.reserved-67"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1KeyId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.keyId-68"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1SequenceNumber() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.sequenceNumber-69"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1MsgLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.msgLength-70"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1KeyedMessege() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.keyedMessege-71"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyTimevaluesClassClass5CType1RefreshPeriodR() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.timevaluesClassClass5CType1.refreshPeriodR-72"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1Ipv4ErrorNodeAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.ipv4ErrorNodeAddress-73"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.flags-74"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1ErrorCode() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.errorCode-75"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1ErrorValue() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.errorValue-76"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2Ipv6ErrorNodeAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.ipv6ErrorNodeAddress-77"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.flags-78"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2ErrorCode() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.errorCode-79"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2ErrorValue() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.errorValue-80"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ScopeClassClass7CType1Ipv4SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ScopeClassClass7CType1.ipv4SrcAddress-81"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ScopeClassClass7CType2Ipv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ScopeClassClass7CType2.ipv6SrcAddress-82"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.flags-83"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.reserved-84"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1SharingControl() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.sharingControl-85"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1SenderSelectionControl() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.senderSelectionControl-86"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4SignalType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.signalType-87"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Rcc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.rcc-88"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Ncc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.ncc-89"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Nvc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.nvc-90"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Multiplier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.multiplier-91"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Transparency() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.transparency-92"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Profile() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.profile-93"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5SignalType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.signalType-94"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.reserved-95"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Nmc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.nmc-96"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Nvc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.nvc-97"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Multiplier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.multiplier-98"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Reserved_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.reserved-99"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType1Ipv4SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType1.ipv4SrcAddress-100"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType1Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType1.unused-101"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType1SrcPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType1.srcPort-102"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType2Ipv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType2.ipv6SrcAddress-103"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType2Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType2.unused-104"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType2SrcPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType2.srcPort-105"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType3Ipv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType3.ipv6SrcAddress-106"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType3Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType3.unused-107"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType3FlowLabel() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType3.flowLabel-108"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPIFILTERSPECClass10CType4Ipv4SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPIFILTERSPECClass10CType4.ipv4SrcAddress-109"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPIFILTERSPECClass10CType4Gpi() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPIFILTERSPECClass10CType4.gpi-110"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPIFILTERSPECClass10CType5Ipv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPIFILTERSPECClass10CType5.ipv6SrcAddress-111"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPIFILTERSPECClass10CType5Gpi() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPIFILTERSPECClass10CType5.gpi-112"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4FILTERSPECClass10CType7Ipv4TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.ipv4TunnelSenderAddress-113"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4FILTERSPECClass10CType7Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.unused-114"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4FILTERSPECClass10CType7LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.lspID-115"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6FILTERSPECClass10CType8Ipv6TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.ipv6TunnelSenderAddress-116"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6FILTERSPECClass10CType8Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.unused-117"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6FILTERSPECClass10CType8LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.lspID-118"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12Ipv4TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.ipv4TunnelSenderAddress-119"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.unused-120"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.lspID-121"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12SubGroupOriginatorID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.subGroupOriginatorID-122"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12Unused_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.unused-123"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12SubGroupID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.subGroupID-124"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13Ipv6TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.ipv6TunnelSenderAddress-125"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.unused-126"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.lspID-127"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13SubGroupOriginatorID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.subGroupOriginatorID-128"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13Unused_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.unused-129"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13SubGroupID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.subGroupID-130"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4SENDERTEMPLATEClass11CType1Ipv4SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.ipv4SrcAddress-131"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4SENDERTEMPLATEClass11CType1Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.unused-132"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4SENDERTEMPLATEClass11CType1SrcPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.srcPort-133"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6SENDERTEMPLATEClass11CType2Ipv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.ipv6SrcAddress-134"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6SENDERTEMPLATEClass11CType2Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.unused-135"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6SENDERTEMPLATEClass11CType2SrcPort() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.srcPort-136"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4FlowlabelSENDERTEMPLATEClass11CType3Ipv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.ipv6SrcAddress-137"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4FlowlabelSENDERTEMPLATEClass11CType3Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.unused-138"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4FlowlabelSENDERTEMPLATEClass11CType3FlowLabel() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.flowLabel-139"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISENDERTEMPLATEClass11CType4Ipv4SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISENDERTEMPLATEClass11CType4.ipv4SrcAddress-140"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISENDERTEMPLATEClass11CType4Gpi() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISENDERTEMPLATEClass11CType4.gpi-141"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISENDERTEMPLATEClass11CType5Ipv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISENDERTEMPLATEClass11CType5.ipv6SrcAddress-142"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISENDERTEMPLATEClass11CType5Gpi() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISENDERTEMPLATEClass11CType5.gpi-143"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4SENDERTEMPLATEClass11CType7Ipv4TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.ipv4TunnelSenderAddress-144"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4SENDERTEMPLATEClass11CType7Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.unused-145"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4SENDERTEMPLATEClass11CType7LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.lspID-146"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6SENDERTEMPLATEClass11CType8Ipv6TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.ipv6TunnelSenderAddress-147"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6SENDERTEMPLATEClass11CType8Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.unused-148"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6SENDERTEMPLATEClass11CType8LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.lspID-149"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12Ipv4TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.ipv4TunnelSenderAddress-150"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.unused-151"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.lspID-152"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12SubGroupOriginatorID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.subGroupOriginatorID-153"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12Unused_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.unused-154"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12SubGroupID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.subGroupID-155"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13Ipv6TunnelSenderAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.ipv6TunnelSenderAddress-156"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.unused-157"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13LspID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.lspID-158"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13SubGroupOriginatorID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.subGroupOriginatorID-159"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13Unused_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.unused-160"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13SubGroupID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.subGroupID-161"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2VersionNumber() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.versionNumber-162"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Reserved1() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.reserved1-163"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2OverallLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.overallLength-164"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2ServiceHeader() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.serviceHeader-165"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2ZeroBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.zeroBit-166"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Reserved2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.reserved2-167"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2LengthOfService1Data() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.lengthOfService1Data-168"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2ParameterIDTokenBucketTSpec() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameterIDTokenBucketTSpec-169"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Parameter127Flag() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameter127Flag-170"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Parameter127Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameter127Length-171"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2TokenBucketRate() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.tokenBucketRate-172"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2TokenBucketSize() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.tokenBucketSize-173"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2PeakDataRate() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.peakDataRate-174"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2MinimumPolicedUnit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.minimumPolicedUnit-175"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2MaximumPacketSize() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.maximumPacketSize-176"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4SignalType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.signalType-177"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Rcc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.rcc-178"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Ncc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.ncc-179"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Nvc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.nvc-180"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Multiplier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.multiplier-181"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Transparency() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.transparency-182"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Profile() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.profile-183"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5SignalType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.signalType-184"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.reserved-185"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Nmc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.nmc-186"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Nvc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.nvc-187"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Multiplier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.multiplier-188"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Reserved_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.reserved-189"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2MessageFormatVersionNumber() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.messageFormatVersionNumber-190"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.reserved-191"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2MsgLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.msgLength-192"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1PerServiceHeaderServiceNumber1() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.perServiceHeaderServiceNumber1-193"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1X() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.x-194"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Reserved3() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.reserved3-195"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1GlobalBreakBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.globalBreakBit-196"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID4() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID4-197"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter4FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter4FlagByte-198"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter4Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter4Length-199"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1IsHopCnt() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.isHopCnt-200"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID6() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID6-201"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter6FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter6FlagByte-202"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter6Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter6Length-203"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1PathBwEstimate() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.pathBwEstimate-204"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID8() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID8-205"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter8FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter8FlagByte-206"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter8Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter8Length-207"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1MinimumPathLatency() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.minimumPathLatency-208"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID10() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID10-209"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter10FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter10FlagByte-210"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter10Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter10Length-211"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ComposedMTU() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.composedMTU-212"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2PerServiceHeaderServiceNumber2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.perServiceHeaderServiceNumber2-213"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2XBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.xBit-214"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.reserved-215"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2BreakBitAndLengthOfPerserviceData() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.breakBitAndLengthOfPerserviceData-216"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter133ComposedCtot() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter133ComposedCtot-217"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter133FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter133FlagByte-218"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter133Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter133Length-219"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsEndtoendComposedValueForCCtot() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.endtoendComposedValueForCCtot-220"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter134ComposedDtot() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter134ComposedDtot-221"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter134FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter134FlagByte-222"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter134Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter134Length-223"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsEndtoendComposedValueForDDtot() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.endtoendComposedValueForDDtot-224"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter135ComposedCsum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter135ComposedCsum-225"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter135FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter135FlagByte-226"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter135Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter135Length-227"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsSincelastreshapingPointComposedCCsum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.sincelastreshapingPointComposedCCsum-228"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter136ComposedDsum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter136ComposedDsum-229"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter136FlagByte() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter136FlagByte-230"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter136Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter136Length-231"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsSincelastreshapingPointComposedDDsum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.sincelastreshapingPointComposedDDsum-232"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsServicespecificGeneralParameterHeadersvaluesServicespecificGeneralParameterHeadervalue() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.servicespecificGeneralParameterHeadersvalues.servicespecificGeneralParameterHeadervalue-233"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentPerServiceHeaderServiceNumber5() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.perServiceHeaderServiceNumber5-234"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentXBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.xBit-235"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentBreakBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.breakBit-236"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentLengthOfPerserviceData() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.lengthOfPerserviceData-237"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentServicespecificGeneralParameterHeadersServicespecificGeneralParameterHeader() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.servicespecificGeneralParameterHeaders.servicespecificGeneralParameterHeader-238"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.length-239"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Policydata() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policydata-240"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Unused() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.unused-241"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1DataOffset() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.dataOffset-242"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Unused_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.unused-243"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1OptionDataLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.optionDataLength-244"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1OptionData() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.optionData-245"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1PolicyDataLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policyDataLength-246"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1PolicyData() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policyData-247"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4RESVCONFIRMClass15CType1Ipv4ReceiverAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4RESVCONFIRMClass15CType1.ipv4ReceiverAddress-248"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6RESVCONFIRMClass15CType2Ipv6ReceiverAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6RESVCONFIRMClass15CType2.ipv6ReceiverAddress-249"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectClass16CType1TopLabel() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectClass16CType1.topLabel-250"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelRequestWithoutLabelRangeClass19CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelRequestWithoutLabelRangeClass19CType1.reserved-251"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelRequestWithoutLabelRangeClass19CType1L3pid() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelRequestWithoutLabelRangeClass19CType1.l3pid-252"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-253"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2L3pid() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.l3pid-254"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.mBit-255"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2Reserved_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-256"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MinimumVPI() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.minimumVPI-257"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MinimumVCI() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.minimumVCI-258"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2Reserved_3() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-259"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MaximumVPI() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.maximumVPI-260"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MaximumVCI() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.maximumVCI-261"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.reserved-262"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3L3pid() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.l3pid-263"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Res() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.res-264"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Dli() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.dli-265"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3MinimumDLCI() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.minimumDLCI-266"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Res_2() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.res-267"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3MaximumDLCI() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.maximumDLCI-268"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1LBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.lBit-269"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.type-270"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.length-271"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Ipv4Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.ipv4Address-272"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.prefixLength-273"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Padding() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.padding-274"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2LBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.lBit-275"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.type-276"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.length-277"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Ipv6Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.ipv6Address-278"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.prefixLength-279"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Padding() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.padding-280"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32LBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.lBit-281"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.type-282"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.length-283"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32AsNumber() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.asNumber-284"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.type-285"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.length-286"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Ipv4Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.ipv4Address-287"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.prefixLength-288"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.flags-289"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.type-290"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.length-291"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Ipv6Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.ipv6Address-292"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.prefixLength-293"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.flags-294"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.type-295"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.length-296"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.flags-297"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.ctype-298"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3ContentsOfLabelObject() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.contentsOfLabelObject-299"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyHelloREQUESTAckClass22CType12SrcInstance() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.helloREQUESTAckClass22CType12.srcInstance-300"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyHelloREQUESTAckClass22CType12DestInstance() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.helloREQUESTAckClass22CType12.destInstance-301"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidclass23CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidclass23CType1.flags-302"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidclass23CType1Epoch() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidclass23CType1.epoch-303"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidclass23CType1MessageIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidclass23CType1.messageIdentifier-304"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidAckNackclass24CType12Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.flags-305"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidAckNackclass24CType12Epoch() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.epoch-306"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidAckNackclass24CType12MessageIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.messageIdentifier-307"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidlistclass25CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.flags-308"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidlistclass25CType1Epoch() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.epoch-309"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidlistclass25CType1MessageidlistMessageIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.messageidlist.messageIdentifier-310"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.flags-311"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2Epoch() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.epoch-312"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2MessageidlistMessageIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.messageidlist.messageIdentifier-313"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2MessageidlistSourceIPAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.messageidlist.sourceIPAddress-314"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.flags-315"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3Epoch() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.epoch-316"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3MessageidlistMessageIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.messageidlist.messageIdentifier-317"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3MessageidlistSourceIPAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.messageidlist.sourceIPAddress-318"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.flags-319"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4Epoch() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.epoch-320"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4MessageidlistMessageIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.messageIdentifier-321"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4MessageidlistSourceIPAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.sourceIPAddress-322"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4MessageidlistDestinationIPAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.destinationIPAddress-323"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.flags-324"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5Epoch() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.epoch-325"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5MessageidlistMessageIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.messageIdentifier-326"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5MessageidlistSourceIPAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.sourceIPAddress-327"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5MessageidlistDestinationIPAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.destinationIPAddress-328"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1MaxRSVPhops() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.maxRSVPhops-329"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1Rsvphopcount() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.rsvphopcount-330"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.reserved-331"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1MfBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.mfBit-332"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1RequestIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.requestIdentifier-333"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1PathMTU() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.pathMTU-334"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1FragmentOffset() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.fragmentOffset-335"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1Lasthopaddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.lasthopaddress-336"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1SenderTemplateObjectIpv4SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.senderTemplateObject.ipv4SrcAddress-337"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1SenderTemplateObjectGeneralizedPortIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.senderTemplateObject.generalizedPortIdentifier-338"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1FilterSpecTemplateObjectIpv4SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.filterSpecTemplateObject.ipv4SrcAddress-339"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1FilterSpecTemplateObjectGeneralizedPortIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.filterSpecTemplateObject.generalizedPortIdentifier-340"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2MaxRSVPhops() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.maxRSVPhops-341"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2Rsvphopcount() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.rsvphopcount-342"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.reserved-343"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2MfBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.mfBit-344"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2RequestIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.requestIdentifier-345"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2PathMTU() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.pathMTU-346"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2FragmentOffset() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.fragmentOffset-347"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2Lasthopaddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.lasthopaddress-348"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2SenderTemplateObjectIpv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.senderTemplateObject.ipv6SrcAddress-349"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2SenderTemplateObjectGeneralizedPortIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.senderTemplateObject.generalizedPortIdentifier-350"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2FilterSpecTemplateObjectIpv6SrcAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.filterSpecTemplateObject.ipv6SrcAddress-351"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2FilterSpecTemplateObjectGeneralizedPortIdentifier() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.filterSpecTemplateObject.generalizedPortIdentifier-352"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagselectclass33CType1ClassTypeClass() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagselectclass33CType1.classType.class-353"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagselectclass33CType1ClassTypeCtype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagselectclass33CType1.classType.ctype-354"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv4Objectclass31CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.reserved-355"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv4Objectclass31CType1RPointer() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.rPointer-356"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv4Objectclass31CType1NodeAddressListRsvpNodeAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.nodeAddressList.rsvpNodeAddress-357"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv6Objectclass31CType2Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.reserved-358"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv6Objectclass31CType2RPointer() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.rPointer-359"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv6Objectclass31CType2NodeAddressListRsvpNodeAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.nodeAddressList.rsvpNodeAddress-360"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1DreqArrivalTime() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.dreqArrivalTime-361"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1IncomingInterfaceAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.incomingInterfaceAddress-362"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1OutgoingInterfaceAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.outgoingInterfaceAddress-363"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1PreviousRSVPHopRouterAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.previousRSVPHopRouterAddress-364"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1Dttl() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.dttl-365"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1MBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.mBit-366"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1Rerr() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.rerr-367"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1K() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.k-368"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1TimerValue() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.timerValue-369"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2DreqArrivalTime() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.dreqArrivalTime-370"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2IncomingInterfaceAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.incomingInterfaceAddress-371"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2OutgoingInterfaceAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.outgoingInterfaceAddress-372"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2PreviousRSVPHopRouterAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.previousRSVPHopRouterAddress-373"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2Dttl() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.dttl-374"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2MBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.mBit-375"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2Rerr() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.rerr-376"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2K() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.k-377"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2TimerValue() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.timerValue-378"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyS2lsublspipv4Class50CType1Ipv4S2LSubLSPDestinationAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.s2lsublspipv4Class50CType1.ipv4S2LSubLSPDestinationAddress-379"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyS2lsublspipv6Class50CType2Ipv6S2LSubLSPDestinationAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.s2lsublspipv6Class50CType2.ipv6S2LSubLSPDestinationAddress-380"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7Lengthbytes() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.lengthbytes-381"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7ClassNum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.classNum-382"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.ctype-383"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7PlrAddressListPlrID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.plrAddressList.plrID-384"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7PlrAddressListAvoidNodeID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.plrAddressList.avoidNodeID-385"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8Lengthbytes() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.lengthbytes-386"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8ClassNum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.classNum-387"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.ctype-388"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8PlrAddressListPlrID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.plrAddressList.plrID-389"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8PlrAddressListAvoidNodeID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.plrAddressList.avoidNodeID-390"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyChallengeObjectclass64CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.reserved-391"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyChallengeObjectclass64CType1KeyId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.keyId-392"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyChallengeObjectclass64CType1ChallengeCookie() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.challengeCookie-393"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.reserved-394"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1Mapnb() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapnb-395"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1MapListReserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.reserved-396"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1MapListExp() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.exp-397"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1MapListPhbid() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.phbid-398"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservLLSPclass65CType2Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservLLSPclass65CType2.reserved-399"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservLLSPclass65CType2Psc() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservLLSPclass65CType2.psc-400"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyClasstypeclass66CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.classtypeclass66CType1.reserved-401"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyClasstypeclass66CType1Ct() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.classtypeclass66CType1.ct-402"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelinterfaceidclass193CType1LsrId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelinterfaceidclass193CType1.lsrId-403"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelinterfaceidclass193CType1InterfaceID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelinterfaceidclass193CType1.interfaceID-404"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1LBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.lBit-405"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.type-406"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.length-407"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Ipv4Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.ipv4Address-408"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.prefixLength-409"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Padding() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.padding-410"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2LBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.lBit-411"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.type-412"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.length-413"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Ipv6Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.ipv6Address-414"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.prefixLength-415"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Padding() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.padding-416"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32LBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.lBit-417"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.type-418"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.length-419"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32AsNumber() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.asNumber-420"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.type-421"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.length-422"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Ipv4Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.ipv4Address-423"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.prefixLength-424"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.flags-425"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.type-426"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.length-427"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Ipv6Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.ipv6Address-428"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2PrefixLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.prefixLength-429"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.flags-430"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Type() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.type-431"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Length() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.length-432"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.flags-433"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.ctype-434"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3ContentsOfLabelObject() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.contentsOfLabelObject-435"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Lengthbytes() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.lengthbytes-436"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1ClassNum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.classNum-437"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.ctype-438"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1SetupPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.setupPrio-439"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1HoldingPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.holdingPrio-440"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Hoplimit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.hoplimit-441"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.flags-442"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Bandwidth() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.bandwidth-443"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Includeany() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.includeany-444"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Excludeany() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.excludeany-445"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Includeall() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.includeall-446"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Lengthbytes() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.lengthbytes-447"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7ClassNum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.classNum-448"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.ctype-449"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7SetupPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.setupPrio-450"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7HoldingPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.holdingPrio-451"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Hoplimit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.hoplimit-452"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.reserved-453"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Bandwidth() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.bandwidth-454"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Includeany() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.includeany-455"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Excludeany() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.excludeany-456"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7SetupPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.setupPrio-457"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7HoldingPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.holdingPrio-458"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.flags-459"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7NameLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.nameLength-460"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7SessionName() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.sessionName-461"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Excludeany() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.excludeany-462"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Includeany() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.includeany-463"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Includeall() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.includeall-464"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1SetupPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.setupPrio-465"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1HoldingPrio() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.holdingPrio-466"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.flags-467"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1NameLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.nameLength-468"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1SessionName() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.sessionName-469"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyAtmserviceclassclass227CType1Reserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.atmserviceclassclass227CType1.reserved-470"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyAtmserviceclassclass227CType1Flags() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.atmserviceclassclass227CType1.flags-471"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2Lengthbytes() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.lengthbytes-472"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2ClassNum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.classNum-473"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.ctype-474"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2Resv() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.resv-475"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2CallOpsFlag() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.callOpsFlag-476"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.uBit-477"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.fBit-478"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDSourceIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.sourceIDType-479"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.length-480"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDIpv4Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.ipv4Address-481"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDLogicalPortId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.logicalPortId-482"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.uBit-483"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.fBit-484"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDSourceIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.sourceIDType-485"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.length-486"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDIpv6Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.ipv6Address-487"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDLogicalPortId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.logicalPortId-488"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.uBit-489"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.fBit-490"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDSourceIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.sourceIDType-491"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.length-492"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDDataLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.dataLength-493"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDNsap() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.nsap-494"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDLogicalPortId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.logicalPortId-495"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.uBit-496"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.fBit-497"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDDestIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.destIDType-498"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.length-499"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDIpv4Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.ipv4Address-500"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDLogicalPortId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.logicalPortId-501"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.uBit-502"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.fBit-503"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDDestIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.destIDType-504"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.length-505"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDIpv6Address() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.ipv6Address-506"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDLogicalPortId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.logicalPortId-507"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.uBit-508"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.fBit-509"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDDestIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.destIDType-510"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.length-511"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDDataLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.dataLength-512"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDNsap() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.nsap-513"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDLogicalPortId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.logicalPortId-514"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.uBit-515"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.fBit-516"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVEgressIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.egressIDType-517"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.length-518"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVReserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.reserved-519"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLbit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.lbit-520"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLogicalPortId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.logicalPortId-521"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLabelLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.labelLength-522"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLabel() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.label-523"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.uBit-524"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.fBit-525"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDConnectionIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.connectionIDType-526"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.length-527"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDReserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.reserved-528"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDCbit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.cbit-529"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDLogicalConnectionId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.logicalConnectionId-530"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.uBit-531"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.fBit-532"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityDiversityIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.diversityIDType-533"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.length-534"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityIteratingListLocalConnectionID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.localConnectionID-535"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityIteratingListReserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.reserved-536"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityIteratingListDivT() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.divT-537"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.uBit-538"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.fBit-539"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdContractIDType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.contractIDType-540"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.length-541"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdContractID() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.contractID-542"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelUBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.uBit-543"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelFBit() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.fBit-544"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelServiceLevelType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.serviceLevelType-545"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.length-546"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelReserved() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.reserved-547"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelServiceLevel() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.serviceLevel-548"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1Lengthbytes() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.lengthbytes-549"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1ClassNum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.classNum-550"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.ctype-551"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.type-552"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesResv() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.resv-553"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.srcLSRAddress-554"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.localId-555"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.type-556"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesResv() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.resv-557"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.srcLSRAddress-558"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.localId-559"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.type-560"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesResv() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.resv-561"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.srcLSRAddress-562"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.localId-563"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.type-564"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesResv() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.resv-565"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.srcLSRAddress-566"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.localId-567"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.type-568"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedResv() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.resv-569"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressLength-570"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressValue() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressValue-571"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.localId-572"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2Lengthbytes() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.lengthbytes-573"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2ClassNum() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.classNum-574"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2Ctype() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.ctype-575"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.type-576"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesIs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.is-577"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesNs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.ns-578"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.srcLSRAddress-579"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.localId-580"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.type-581"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesIs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.is-582"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesNs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.ns-583"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.srcLSRAddress-584"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.localId-585"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.type-586"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesIs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.is-587"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesNs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.ns-588"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.srcLSRAddress-589"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.localId-590"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.type-591"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesIs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.is-592"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesNs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.ns-593"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesSrcLSRAddress() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.srcLSRAddress-594"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.localId-595"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedType() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.type-596"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedIs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.is-597"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedNs() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.ns-598"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressLength() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressLength-599"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressValue() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressValue-600"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedLocalId() *TrafficField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.localId-601"]]
}

func (s *RsvpStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewRsvpStack(idx int) *RsvpStack {
	stack := RsvpStack(newStack(idx, "rsvp", rsvpAliasToFieldIdx))
	return &stack
}

type EthernetStack TrafficStack

var ethernetAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.destinationAddress-1",
	"header.sourceAddress-2",
	"header.etherType-3",
	"header.pfcQueue-4",
})

func (s *EthernetStack) DestinationAddress() *TrafficField {
	return s.Field[ethernetAliasToFieldIdx["header.destinationAddress-1"]]
}
func (s *EthernetStack) SourceAddress() *TrafficField {
	return s.Field[ethernetAliasToFieldIdx["header.sourceAddress-2"]]
}
func (s *EthernetStack) EtherType() *TrafficField {
	return s.Field[ethernetAliasToFieldIdx["header.etherType-3"]]
}
func (s *EthernetStack) PfcQueue() *TrafficField {
	return s.Field[ethernetAliasToFieldIdx["header.pfcQueue-4"]]
}

func (s *EthernetStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewEthernetStack(idx int) *EthernetStack {
	stack := EthernetStack(newStack(idx, "ethernet", ethernetAliasToFieldIdx))
	return &stack
}

type GreStack TrafficStack

var greAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.checksumPresent-1",
	"header.reserved1-2",
	"header.keyPresent-3",
	"header.sequencePresent-4",
	"header.reserved2-5",
	"header.version-6",
	"header.protocol-7",
	"header.checksumHolder.withChecksum.checksum-8",
	"header.checksumHolder.withChecksum.reserved-9",
	"header.checksumHolder.noChecksum-10",
	"header.keyHolder.key-11",
	"header.keyHolder.noKey-12",
	"header.sequenceHolder.sequenceNum-13",
	"header.sequenceHolder.noSequenceNum-14",
})

func (s *GreStack) ChecksumPresent() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.checksumPresent-1"]]
}
func (s *GreStack) Reserved1() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.reserved1-2"]]
}
func (s *GreStack) KeyPresent() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.keyPresent-3"]]
}
func (s *GreStack) SequencePresent() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.sequencePresent-4"]]
}
func (s *GreStack) Reserved2() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.reserved2-5"]]
}
func (s *GreStack) Version() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.version-6"]]
}
func (s *GreStack) Protocol() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.protocol-7"]]
}
func (s *GreStack) ChecksumHolderWithChecksumChecksum() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.checksumHolder.withChecksum.checksum-8"]]
}
func (s *GreStack) ChecksumHolderWithChecksumReserved() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.checksumHolder.withChecksum.reserved-9"]]
}
func (s *GreStack) ChecksumHolderNoChecksum() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.checksumHolder.noChecksum-10"]]
}
func (s *GreStack) KeyHolderKey() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.keyHolder.key-11"]]
}
func (s *GreStack) KeyHolderNoKey() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.keyHolder.noKey-12"]]
}
func (s *GreStack) SequenceHolderSequenceNum() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.sequenceHolder.sequenceNum-13"]]
}
func (s *GreStack) SequenceHolderNoSequenceNum() *TrafficField {
	return s.Field[greAliasToFieldIdx["header.sequenceHolder.noSequenceNum-14"]]
}

func (s *GreStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewGreStack(idx int) *GreStack {
	stack := GreStack(newStack(idx, "gre", greAliasToFieldIdx))
	return &stack
}

type MplsStack TrafficStack

var mplsAliasToFieldIdx = aliasesToFieldIdx([]string{
	"label.value-1",
	"label.experimental-2",
	"label.bottomOfStack-3",
	"label.ttl-4",
	"label.tracker-5",
})

func (s *MplsStack) Value() *TrafficField {
	return s.Field[mplsAliasToFieldIdx["label.value-1"]]
}
func (s *MplsStack) Experimental() *TrafficField {
	return s.Field[mplsAliasToFieldIdx["label.experimental-2"]]
}
func (s *MplsStack) BottomOfStack() *TrafficField {
	return s.Field[mplsAliasToFieldIdx["label.bottomOfStack-3"]]
}
func (s *MplsStack) Ttl() *TrafficField {
	return s.Field[mplsAliasToFieldIdx["label.ttl-4"]]
}
func (s *MplsStack) Tracker() *TrafficField {
	return s.Field[mplsAliasToFieldIdx["label.tracker-5"]]
}

func (s *MplsStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewMplsStack(idx int) *MplsStack {
	stack := MplsStack(newStack(idx, "mpls", mplsAliasToFieldIdx))
	return &stack
}

type Ospfv2DatabaseDescriptionStack TrafficStack

var ospfv2DatabaseDescriptionAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.ospfv2PacketHeader.ospfVersion-1",
	"header.ospfv2PacketHeader.packetType-2",
	"header.ospfv2PacketHeader.packetLength-3",
	"header.ospfv2PacketHeader.routerID-4",
	"header.ospfv2PacketHeader.areaID-5",
	"header.ospfv2PacketHeader.checksum-6",
	"header.ospfv2PacketHeader.authenticationType-7",
	"header.ospfv2PacketHeader.authenticationData.nullAuthentication-8",
	"header.ospfv2PacketHeader.authenticationData.simplePassword-9",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13",
	"header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14",
	"header.databaseDescriptionBody.interfaceMTU-15",
	"header.databaseDescriptionBody.options-16",
	"header.databaseDescriptionBody.reserved-17",
	"header.databaseDescriptionBody.databaseDescriptionFlags-18",
	"header.databaseDescriptionBody.ddSequenceNumber-19",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateAge-20",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.options-21",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateType-22",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateID-23",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateAdvertisingRouter-24",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateSequenceNumber-25",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateChecksum-26",
	"header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateLength-27",
})

func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderOspfVersion() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderPacketType() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderPacketLength() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderRouterID() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAreaID() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderChecksum() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationType() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyInterfaceMTU() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.interfaceMTU-15"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyOptions() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.options-16"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyReserved() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.reserved-17"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyDatabaseDescriptionFlags() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.databaseDescriptionFlags-18"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyDdSequenceNumber() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.ddSequenceNumber-19"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateAge() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateAge-20"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderOptions() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.options-21"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateType() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateType-22"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateID() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateID-23"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateAdvertisingRouter() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateAdvertisingRouter-24"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateSequenceNumber() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateSequenceNumber-25"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateChecksum() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateChecksum-26"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateLength() *TrafficField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateLength-27"]]
}

func (s *Ospfv2DatabaseDescriptionStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewOspfv2DatabaseDescriptionStack(idx int) *Ospfv2DatabaseDescriptionStack {
	stack := Ospfv2DatabaseDescriptionStack(newStack(idx, "ospfv2DatabaseDescription", ospfv2DatabaseDescriptionAliasToFieldIdx))
	return &stack
}

type Ospfv2HelloStack TrafficStack

var ospfv2HelloAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.ospfv2PacketHeader.ospfVersion-1",
	"header.ospfv2PacketHeader.packetType-2",
	"header.ospfv2PacketHeader.packetLength-3",
	"header.ospfv2PacketHeader.routerID-4",
	"header.ospfv2PacketHeader.areaID-5",
	"header.ospfv2PacketHeader.checksum-6",
	"header.ospfv2PacketHeader.authenticationType-7",
	"header.ospfv2PacketHeader.authenticationData.nullAuthentication-8",
	"header.ospfv2PacketHeader.authenticationData.simplePassword-9",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13",
	"header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14",
	"header.networkMask-15",
	"header.helloInterval-16",
	"header.options-17",
	"header.routerPriority-18",
	"header.routerDeadInterval-19",
	"header.designatedRouterID-20",
	"header.backupDesignatedRouterID-21",
	"header.helloNeighborList.neighborRouterID-22",
})

func (s *Ospfv2HelloStack) Ospfv2PacketHeaderOspfVersion() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderPacketType() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderPacketLength() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderRouterID() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAreaID() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderChecksum() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationType() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2HelloStack) NetworkMask() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.networkMask-15"]]
}
func (s *Ospfv2HelloStack) HelloInterval() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.helloInterval-16"]]
}
func (s *Ospfv2HelloStack) Options() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.options-17"]]
}
func (s *Ospfv2HelloStack) RouterPriority() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.routerPriority-18"]]
}
func (s *Ospfv2HelloStack) RouterDeadInterval() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.routerDeadInterval-19"]]
}
func (s *Ospfv2HelloStack) DesignatedRouterID() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.designatedRouterID-20"]]
}
func (s *Ospfv2HelloStack) BackupDesignatedRouterID() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.backupDesignatedRouterID-21"]]
}
func (s *Ospfv2HelloStack) HelloNeighborListNeighborRouterID() *TrafficField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.helloNeighborList.neighborRouterID-22"]]
}

func (s *Ospfv2HelloStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewOspfv2HelloStack(idx int) *Ospfv2HelloStack {
	stack := Ospfv2HelloStack(newStack(idx, "ospfv2Hello", ospfv2HelloAliasToFieldIdx))
	return &stack
}

type Ospfv2LSAAcknowledgementStack TrafficStack

var ospfv2LSAAcknowledgementAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.ospfv2PacketHeader.ospfVersion-1",
	"header.ospfv2PacketHeader.packetType-2",
	"header.ospfv2PacketHeader.packetLength-3",
	"header.ospfv2PacketHeader.routerID-4",
	"header.ospfv2PacketHeader.areaID-5",
	"header.ospfv2PacketHeader.checksum-6",
	"header.ospfv2PacketHeader.authenticationType-7",
	"header.ospfv2PacketHeader.authenticationData.nullAuthentication-8",
	"header.ospfv2PacketHeader.authenticationData.simplePassword-9",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13",
	"header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14",
	"header.linkStateAdvertisementHeader.variableHeader.linkStateAge-15",
	"header.linkStateAdvertisementHeader.variableHeader.options-16",
	"header.linkStateAdvertisementHeader.variableHeader.linkStateType-17",
	"header.linkStateAdvertisementHeader.variableHeader.linkStateID-18",
	"header.linkStateAdvertisementHeader.variableHeader.linkStateAdvertisingRouter-19",
	"header.linkStateAdvertisementHeader.variableHeader.linkStateSequenceNumber-20",
	"header.linkStateAdvertisementHeader.variableHeader.linkStateChecksum-21",
	"header.linkStateAdvertisementHeader.variableHeader.linkStateLength-22",
})

func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderOspfVersion() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderPacketType() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderPacketLength() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderRouterID() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAreaID() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderChecksum() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationType() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateAge() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateAge-15"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderOptions() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.options-16"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateType() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateType-17"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateID() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateID-18"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateAdvertisingRouter() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateAdvertisingRouter-19"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateSequenceNumber() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateSequenceNumber-20"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateChecksum() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateChecksum-21"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateLength() *TrafficField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateLength-22"]]
}

func (s *Ospfv2LSAAcknowledgementStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewOspfv2LSAAcknowledgementStack(idx int) *Ospfv2LSAAcknowledgementStack {
	stack := Ospfv2LSAAcknowledgementStack(newStack(idx, "ospfv2LSAAcknowledgement", ospfv2LSAAcknowledgementAliasToFieldIdx))
	return &stack
}

type Ospfv2LSARequestStack TrafficStack

var ospfv2LSARequestAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.ospfv2PacketHeader.ospfVersion-1",
	"header.ospfv2PacketHeader.packetType-2",
	"header.ospfv2PacketHeader.packetLength-3",
	"header.ospfv2PacketHeader.routerID-4",
	"header.ospfv2PacketHeader.areaID-5",
	"header.ospfv2PacketHeader.checksum-6",
	"header.ospfv2PacketHeader.authenticationType-7",
	"header.ospfv2PacketHeader.authenticationData.nullAuthentication-8",
	"header.ospfv2PacketHeader.authenticationData.simplePassword-9",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13",
	"header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14",
	"header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.reserved-15",
	"header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateType-16",
	"header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateID-17",
	"header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateAdvertisingRouter-18",
})

func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderOspfVersion() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderPacketType() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderPacketLength() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderRouterID() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAreaID() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderChecksum() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationType() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionReserved() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.reserved-15"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateType() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateType-16"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateID() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateID-17"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateAdvertisingRouter() *TrafficField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateAdvertisingRouter-18"]]
}

func (s *Ospfv2LSARequestStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewOspfv2LSARequestStack(idx int) *Ospfv2LSARequestStack {
	stack := Ospfv2LSARequestStack(newStack(idx, "ospfv2LSARequest", ospfv2LSARequestAliasToFieldIdx))
	return &stack
}

type Ospfv2LSAUpdateStack TrafficStack

var ospfv2LSAUpdateAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.ospfv2PacketHeader.ospfVersion-1",
	"header.ospfv2PacketHeader.packetType-2",
	"header.ospfv2PacketHeader.packetLength-3",
	"header.ospfv2PacketHeader.routerID-4",
	"header.ospfv2PacketHeader.areaID-5",
	"header.ospfv2PacketHeader.checksum-6",
	"header.ospfv2PacketHeader.authenticationType-7",
	"header.ospfv2PacketHeader.authenticationData.nullAuthentication-8",
	"header.ospfv2PacketHeader.authenticationData.simplePassword-9",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12",
	"header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13",
	"header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14",
	"header.linkStateUpdateBody.numberOfLSAs-15",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAge-16",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.options-17",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateType-18",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateID-19",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisingRouter-20",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateSequenceNumber-21",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateChecksum-22",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateLength-23",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.reserved-24",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerLSAFlags-25",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.pad-26",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.numberOfRouterInterfaceLinks-27",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.routerInterfaceLinkID-28",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.routerInterfaceLinkData-29",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.interfaceType-30",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.numberOfTOSEntries-31",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.interfaceMetric-32",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.typeOfService-33",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.reserved-34",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.metricForCorrespondingTypeOfService-35",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkLSA.networkMask-36",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkLSA.attachedRouterList.attachedRouterID-37",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.networkMask-38",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.reserved-39",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.routeMetric-40",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.tosList.tosEntry.typeOfService-41",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-42",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.networkMask-43",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.reserved-44",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.routeMetric-45",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.tosList.tosEntry.typeOfService-46",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-47",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.networkMask-48",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.ebit-49",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.reserved-50",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.routeMetric-51",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.forwardingAddress-52",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.externalRouteTag-53",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.ebit-54",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.typeOfService-55",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-56",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.forwardingAddress-57",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.networkMask-58",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.ebit-59",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.reserved-60",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.routeMetric-61",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.forwardingAddress-62",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.externalRouteTag-63",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.ebit-64",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.typeOfService-65",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-66",
	"header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.forwardingAddress-67",
})

func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderOspfVersion() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderPacketType() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderPacketLength() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderRouterID() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAreaID() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderChecksum() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationType() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyNumberOfLSAs() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.numberOfLSAs-15"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAge() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAge-16"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderOptions() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.options-17"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateType() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateType-18"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateID() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateID-19"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisingRouter() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisingRouter-20"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateSequenceNumber() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateSequenceNumber-21"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateChecksum() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateChecksum-22"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateLength() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateLength-23"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSAReserved() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.reserved-24"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterLSAFlags() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerLSAFlags-25"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSAPad() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.pad-26"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSANumberOfRouterInterfaceLinks() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.numberOfRouterInterfaceLinks-27"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceRouterInterfaceLinkID() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.routerInterfaceLinkID-28"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceRouterInterfaceLinkData() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.routerInterfaceLinkData-29"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceInterfaceType() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.interfaceType-30"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceNumberOfTOSEntries() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.numberOfTOSEntries-31"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceInterfaceMetric() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.interfaceMetric-32"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceTosListTosEntryTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.typeOfService-33"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceTosListTosEntryReserved() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.reserved-34"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.metricForCorrespondingTypeOfService-35"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkLSANetworkMask() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkLSA.networkMask-36"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkLSAAttachedRouterListAttachedRouterID() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkLSA.attachedRouterList.attachedRouterID-37"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteNetworkMask() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.networkMask-38"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteReserved() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.reserved-39"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteRouteMetric() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.routeMetric-40"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteTosListTosEntryTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.tosList.tosEntry.typeOfService-41"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-42"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteNetworkMask() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.networkMask-43"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteReserved() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.reserved-44"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteRouteMetric() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.routeMetric-45"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteTosListTosEntryTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.tosList.tosEntry.typeOfService-46"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-47"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteNetworkMask() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.networkMask-48"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteEbit() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.ebit-49"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteReserved() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.reserved-50"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteRouteMetric() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.routeMetric-51"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteForwardingAddress() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.forwardingAddress-52"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteExternalRouteTag() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.externalRouteTag-53"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryEbit() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.ebit-54"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.typeOfService-55"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-56"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryForwardingAddress() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.forwardingAddress-57"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteNetworkMask() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.networkMask-58"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteEbit() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.ebit-59"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteReserved() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.reserved-60"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteRouteMetric() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.routeMetric-61"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteForwardingAddress() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.forwardingAddress-62"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteExternalRouteTag() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.externalRouteTag-63"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryEbit() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.ebit-64"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.typeOfService-65"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-66"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryForwardingAddress() *TrafficField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.forwardingAddress-67"]]
}

func (s *Ospfv2LSAUpdateStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewOspfv2LSAUpdateStack(idx int) *Ospfv2LSAUpdateStack {
	stack := Ospfv2LSAUpdateStack(newStack(idx, "ospfv2LSAUpdate", ospfv2LSAUpdateAliasToFieldIdx))
	return &stack
}

type TcpStack TrafficStack

var tcpAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.srcPort-1",
	"header.dstPort-2",
	"header.sequenceNumber-3",
	"header.acknowledgementNumber-4",
	"header.dataOffset-5",
	"header.reserved-6",
	"header.ecn.nsBit-7",
	"header.ecn.cwrBit-8",
	"header.ecn.ecnEchoBit-9",
	"header.controlBits.urgBit-10",
	"header.controlBits.ackBit-11",
	"header.controlBits.pshBit-12",
	"header.controlBits.rstBit-13",
	"header.controlBits.synBit-14",
	"header.controlBits.finBit-15",
	"header.window-16",
	"header.checksum-17",
	"header.urgentPtr-18",
	"header.options.option.type.userDefined.kind-19",
	"header.options.option.type.userDefined.length-20",
	"header.options.option.type.userDefined.data-21",
	"header.options.option.type.endOfOptionList.kind-22",
	"header.options.option.type.noOperation.kind-23",
	"header.options.option.type.maximumSegmentSize.kind-24",
	"header.options.option.type.maximumSegmentSize.length-25",
	"header.options.option.type.maximumSegmentSize.data-26",
	"header.options.option.type.wsopt.kind-27",
	"header.options.option.type.wsopt.length-28",
	"header.options.option.type.wsopt.data-29",
	"header.options.option.type.sackPermitted.kind-30",
	"header.options.option.type.sackPermitted.length-31",
	"header.options.option.type.sack.kind-32",
	"header.options.option.type.sack.length-33",
	"header.options.option.type.sack.data-34",
	"header.options.option.type.echo.kind-35",
	"header.options.option.type.echo.length-36",
	"header.options.option.type.echo.data-37",
	"header.options.option.type.echoReply.kind-38",
	"header.options.option.type.echoReply.length-39",
	"header.options.option.type.echoReply.data-40",
	"header.options.option.type.tsopt.kind-41",
	"header.options.option.type.tsopt.length-42",
	"header.options.option.type.tsopt.data-43",
	"header.options.option.type.partialOrderConnectionPermitted.kind-44",
	"header.options.option.type.partialOrderConnectionPermitted.length-45",
	"header.options.option.type.partialOrderServiceProfile.kind-46",
	"header.options.option.type.partialOrderServiceProfile.length-47",
	"header.options.option.type.partialOrderServiceProfile.data-48",
	"header.options.option.type.cc.kind-49",
	"header.options.option.type.ccNew.kind-50",
	"header.options.option.type.ccEcho.kind-51",
	"header.options.option.type.alternateChecksumRequest.kind-52",
	"header.options.option.type.alternateChecksumRequest.length-53",
	"header.options.option.type.alternateChecksumRequest.data-54",
	"header.options.option.type.alternateChecksumData.kind-55",
	"header.options.option.type.alternateChecksumData.length-56",
	"header.options.option.type.alternateChecksumData.data-57",
	"header.options.option.type.skeeter.kind-58",
	"header.options.option.type.bubba.kind-59",
	"header.options.option.type.trailerChecksum.kind-60",
	"header.options.option.type.trailerChecksum.length-61",
	"header.options.option.type.trailerChecksum.data-62",
	"header.options.option.type.md5Signature.kind-63",
	"header.options.option.type.md5Signature.length-64",
	"header.options.option.type.md5Signature.data-65",
	"header.options.option.type.scpsCapabilities.kind-66",
	"header.options.option.type.selectiveNegativeAck.kind-67",
	"header.options.option.type.recordBoundaries.kind-68",
	"header.options.option.type.corruptionExperienced.kind-69",
	"header.options.option.type.snap.kind-70",
	"header.options.option.type.unassigned1.kind-71",
	"header.options.option.type.compressionFilter.kind-72",
	"header.options.option.type.quickStartResponse.kind-73",
	"header.options.option.type.quickStartResponse.length-74",
	"header.options.option.type.quickStartResponse.data-75",
	"header.options.option.type.unassigned2.kind-76",
	"header.options.option.type.rfc3692StypeExperiment1.kind-77",
	"header.options.option.type.rfc3692StypeExperiment1.length-78",
	"header.options.option.type.rfc3692StypeExperiment1.data-79",
	"header.options.option.type.rfc3692StypeExperiment2.kind-80",
	"header.options.option.type.rfc3692StypeExperiment2.length-81",
	"header.options.option.type.rfc3692StypeExperiment2.data-82",
	"header.options.pad-83",
})

func (s *TcpStack) SrcPort() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.srcPort-1"]]
}
func (s *TcpStack) DstPort() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.dstPort-2"]]
}
func (s *TcpStack) SequenceNumber() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.sequenceNumber-3"]]
}
func (s *TcpStack) AcknowledgementNumber() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.acknowledgementNumber-4"]]
}
func (s *TcpStack) DataOffset() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.dataOffset-5"]]
}
func (s *TcpStack) Reserved() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.reserved-6"]]
}
func (s *TcpStack) EcnNsBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.ecn.nsBit-7"]]
}
func (s *TcpStack) EcnCwrBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.ecn.cwrBit-8"]]
}
func (s *TcpStack) EcnEcnEchoBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.ecn.ecnEchoBit-9"]]
}
func (s *TcpStack) ControlBitsUrgBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.urgBit-10"]]
}
func (s *TcpStack) ControlBitsAckBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.ackBit-11"]]
}
func (s *TcpStack) ControlBitsPshBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.pshBit-12"]]
}
func (s *TcpStack) ControlBitsRstBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.rstBit-13"]]
}
func (s *TcpStack) ControlBitsSynBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.synBit-14"]]
}
func (s *TcpStack) ControlBitsFinBit() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.finBit-15"]]
}
func (s *TcpStack) Window() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.window-16"]]
}
func (s *TcpStack) Checksum() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.checksum-17"]]
}
func (s *TcpStack) UrgentPtr() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.urgentPtr-18"]]
}
func (s *TcpStack) OptionsOptionTypeUserDefinedKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.userDefined.kind-19"]]
}
func (s *TcpStack) OptionsOptionTypeUserDefinedLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.userDefined.length-20"]]
}
func (s *TcpStack) OptionsOptionTypeUserDefinedData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.userDefined.data-21"]]
}
func (s *TcpStack) OptionsOptionTypeEndOfOptionListKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.endOfOptionList.kind-22"]]
}
func (s *TcpStack) OptionsOptionTypeNoOperationKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.noOperation.kind-23"]]
}
func (s *TcpStack) OptionsOptionTypeMaximumSegmentSizeKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.maximumSegmentSize.kind-24"]]
}
func (s *TcpStack) OptionsOptionTypeMaximumSegmentSizeLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.maximumSegmentSize.length-25"]]
}
func (s *TcpStack) OptionsOptionTypeMaximumSegmentSizeData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.maximumSegmentSize.data-26"]]
}
func (s *TcpStack) OptionsOptionTypeWsoptKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.wsopt.kind-27"]]
}
func (s *TcpStack) OptionsOptionTypeWsoptLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.wsopt.length-28"]]
}
func (s *TcpStack) OptionsOptionTypeWsoptData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.wsopt.data-29"]]
}
func (s *TcpStack) OptionsOptionTypeSackPermittedKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sackPermitted.kind-30"]]
}
func (s *TcpStack) OptionsOptionTypeSackPermittedLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sackPermitted.length-31"]]
}
func (s *TcpStack) OptionsOptionTypeSackKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sack.kind-32"]]
}
func (s *TcpStack) OptionsOptionTypeSackLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sack.length-33"]]
}
func (s *TcpStack) OptionsOptionTypeSackData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sack.data-34"]]
}
func (s *TcpStack) OptionsOptionTypeEchoKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echo.kind-35"]]
}
func (s *TcpStack) OptionsOptionTypeEchoLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echo.length-36"]]
}
func (s *TcpStack) OptionsOptionTypeEchoData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echo.data-37"]]
}
func (s *TcpStack) OptionsOptionTypeEchoReplyKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echoReply.kind-38"]]
}
func (s *TcpStack) OptionsOptionTypeEchoReplyLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echoReply.length-39"]]
}
func (s *TcpStack) OptionsOptionTypeEchoReplyData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echoReply.data-40"]]
}
func (s *TcpStack) OptionsOptionTypeTsoptKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.tsopt.kind-41"]]
}
func (s *TcpStack) OptionsOptionTypeTsoptLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.tsopt.length-42"]]
}
func (s *TcpStack) OptionsOptionTypeTsoptData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.tsopt.data-43"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderConnectionPermittedKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderConnectionPermitted.kind-44"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderConnectionPermittedLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderConnectionPermitted.length-45"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderServiceProfileKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderServiceProfile.kind-46"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderServiceProfileLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderServiceProfile.length-47"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderServiceProfileData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderServiceProfile.data-48"]]
}
func (s *TcpStack) OptionsOptionTypeCcKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.cc.kind-49"]]
}
func (s *TcpStack) OptionsOptionTypeCcNewKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.ccNew.kind-50"]]
}
func (s *TcpStack) OptionsOptionTypeCcEchoKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.ccEcho.kind-51"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumRequestKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumRequest.kind-52"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumRequestLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumRequest.length-53"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumRequestData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumRequest.data-54"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumDataKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumData.kind-55"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumDataLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumData.length-56"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumDataData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumData.data-57"]]
}
func (s *TcpStack) OptionsOptionTypeSkeeterKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.skeeter.kind-58"]]
}
func (s *TcpStack) OptionsOptionTypeBubbaKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.bubba.kind-59"]]
}
func (s *TcpStack) OptionsOptionTypeTrailerChecksumKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.trailerChecksum.kind-60"]]
}
func (s *TcpStack) OptionsOptionTypeTrailerChecksumLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.trailerChecksum.length-61"]]
}
func (s *TcpStack) OptionsOptionTypeTrailerChecksumData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.trailerChecksum.data-62"]]
}
func (s *TcpStack) OptionsOptionTypeMd5SignatureKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.md5Signature.kind-63"]]
}
func (s *TcpStack) OptionsOptionTypeMd5SignatureLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.md5Signature.length-64"]]
}
func (s *TcpStack) OptionsOptionTypeMd5SignatureData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.md5Signature.data-65"]]
}
func (s *TcpStack) OptionsOptionTypeScpsCapabilitiesKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.scpsCapabilities.kind-66"]]
}
func (s *TcpStack) OptionsOptionTypeSelectiveNegativeAckKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.selectiveNegativeAck.kind-67"]]
}
func (s *TcpStack) OptionsOptionTypeRecordBoundariesKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.recordBoundaries.kind-68"]]
}
func (s *TcpStack) OptionsOptionTypeCorruptionExperiencedKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.corruptionExperienced.kind-69"]]
}
func (s *TcpStack) OptionsOptionTypeSnapKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.snap.kind-70"]]
}
func (s *TcpStack) OptionsOptionTypeUnassigned1Kind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.unassigned1.kind-71"]]
}
func (s *TcpStack) OptionsOptionTypeCompressionFilterKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.compressionFilter.kind-72"]]
}
func (s *TcpStack) OptionsOptionTypeQuickStartResponseKind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.quickStartResponse.kind-73"]]
}
func (s *TcpStack) OptionsOptionTypeQuickStartResponseLength() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.quickStartResponse.length-74"]]
}
func (s *TcpStack) OptionsOptionTypeQuickStartResponseData() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.quickStartResponse.data-75"]]
}
func (s *TcpStack) OptionsOptionTypeUnassigned2Kind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.unassigned2.kind-76"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment1Kind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment1.kind-77"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment1Length() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment1.length-78"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment1Data() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment1.data-79"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment2Kind() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment2.kind-80"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment2Length() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment2.length-81"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment2Data() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment2.data-82"]]
}
func (s *TcpStack) OptionsPad() *TrafficField {
	return s.Field[tcpAliasToFieldIdx["header.options.pad-83"]]
}

func (s *TcpStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewTcpStack(idx int) *TcpStack {
	stack := TcpStack(newStack(idx, "tcp", tcpAliasToFieldIdx))
	return &stack
}

type UdpStack TrafficStack

var udpAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.srcPort-1",
	"header.dstPort-2",
	"header.length-3",
	"header.checksum-4",
})

func (s *UdpStack) SrcPort() *TrafficField {
	return s.Field[udpAliasToFieldIdx["header.srcPort-1"]]
}
func (s *UdpStack) DstPort() *TrafficField {
	return s.Field[udpAliasToFieldIdx["header.dstPort-2"]]
}
func (s *UdpStack) Length() *TrafficField {
	return s.Field[udpAliasToFieldIdx["header.length-3"]]
}
func (s *UdpStack) Checksum() *TrafficField {
	return s.Field[udpAliasToFieldIdx["header.checksum-4"]]
}

func (s *UdpStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewUdpStack(idx int) *UdpStack {
	stack := UdpStack(newStack(idx, "udp", udpAliasToFieldIdx))
	return &stack
}

type VlanStack TrafficStack

var vlanAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.vlanTag.vlanUserPriority-1",
	"header.vlanTag.cfi-2",
	"header.vlanTag.vlanID-3",
	"header.protocolID-4",
})

func (s *VlanStack) VlanTagVlanUserPriority() *TrafficField {
	return s.Field[vlanAliasToFieldIdx["header.vlanTag.vlanUserPriority-1"]]
}
func (s *VlanStack) VlanTagCfi() *TrafficField {
	return s.Field[vlanAliasToFieldIdx["header.vlanTag.cfi-2"]]
}
func (s *VlanStack) VlanTagVlanID() *TrafficField {
	return s.Field[vlanAliasToFieldIdx["header.vlanTag.vlanID-3"]]
}
func (s *VlanStack) ProtocolID() *TrafficField {
	return s.Field[vlanAliasToFieldIdx["header.protocolID-4"]]
}

func (s *VlanStack) TrafficStack() *TrafficStack {
	ts := TrafficStack(*s)
	return &ts
}

func NewVlanStack(idx int) *VlanStack {
	stack := VlanStack(newStack(idx, "vlan", vlanAliasToFieldIdx))
	return &stack
}
