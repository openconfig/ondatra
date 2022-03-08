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

func newStack(idx int, stackAlias string, fieldAliasToIdx map[string]int) TrafficTrafficItemConfigElementStack {
	stack := TrafficTrafficItemConfigElementStack{
		Xpath: &XPath{
			parentXPath: "/traffic/trafficItem[1]/configElement[1]",
			objectName:  "stack",
			alias:       fmt.Sprintf("%s-%d", stackAlias, idx+1),
		},
	}
	xp := stack.XPath().String()
	stack.Field = make([]*TrafficTrafficItemConfigElementStackField, len(fieldAliasToIdx))
	for fa, i := range fieldAliasToIdx {
		stack.Field[i] = &TrafficTrafficItemConfigElementStackField{
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

type HTTP_GETStack TrafficTrafficItemConfigElementStack

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

func (s *HTTP_GETStack) RequestRequest_Method() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Request Method-1"]]
}
func (s *HTTP_GETStack) RequestSpace1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Space1-2"]]
}
func (s *HTTP_GETStack) RequestRequest_URI() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Request URI-3"]]
}
func (s *HTTP_GETStack) RequestSpace2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Space2-4"]]
}
func (s *HTTP_GETStack) RequestRequest_Version() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.Request version-5"]]
}
func (s *HTTP_GETStack) RequestCRLF1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Request.CRLF1-6"]]
}
func (s *HTTP_GETStack) Host() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Host-7"]]
}
func (s *HTTP_GETStack) User_Agent() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.User-Agent-8"]]
}
func (s *HTTP_GETStack) Accept() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-9"]]
}
func (s *HTTP_GETStack) Accept_Language() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-Language-10"]]
}
func (s *HTTP_GETStack) Accept_Encoding() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-Encoding-11"]]
}
func (s *HTTP_GETStack) Accept_Charset() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Accept-Charset-12"]]
}
func (s *HTTP_GETStack) Keep_Alive() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Keep-Alive-13"]]
}
func (s *HTTP_GETStack) Connection() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Connection-14"]]
}
func (s *HTTP_GETStack) Referer() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.Referer-15"]]
}
func (s *HTTP_GETStack) CRLF() *TrafficTrafficItemConfigElementStackField {
	return s.Field[HTTP_GETAliasToFieldIdx["header.CRLF-16"]]
}

func (s *HTTP_GETStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewHTTP_GETStack(idx int) *HTTP_GETStack {
	stack := HTTP_GETStack(newStack(idx, "HTTP_GET", HTTP_GETAliasToFieldIdx))
	return &stack
}

type Icmpv1Stack TrafficTrafficItemConfigElementStack

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

func (s *Icmpv1Stack) MessageType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.messageType-1"]]
}
func (s *Icmpv1Stack) CodeOptionsDestUnreachableCodeOptions() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.destUnreachableCodeOptions-2"]]
}
func (s *Icmpv1Stack) CodeOptionsSrcQuenchOption() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.srcQuenchOption-3"]]
}
func (s *Icmpv1Stack) CodeOptionsInfoRequestOption() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.infoRequestOption-4"]]
}
func (s *Icmpv1Stack) CodeOptionsInfoResponseOption() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.infoResponseOption-5"]]
}
func (s *Icmpv1Stack) CodeOptionsTimeExceededOptions() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.timeExceededOptions-6"]]
}
func (s *Icmpv1Stack) CodeOptionsRedirectMessageOptions() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.codeOptions.redirectMessageOptions-7"]]
}
func (s *Icmpv1Stack) IcmpChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.icmpChecksum-8"]]
}
func (s *Icmpv1Stack) Next4BytesUnusedBitsInMsgType3() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.unusedBitsInMsgType3-9"]]
}
func (s *Icmpv1Stack) Next4BytesUnusedBitsInMsgType4() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.unusedBitsInMsgType4-10"]]
}
func (s *Icmpv1Stack) Next4BytesUnusedBitsInMsgType11() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.unusedBitsInMsgType11-11"]]
}
func (s *Icmpv1Stack) Next4BytesNextFieldsForParameterProblemPointer() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.nextFieldsForParameterProblem.pointer-12"]]
}
func (s *Icmpv1Stack) Next4BytesNextFieldsForParameterProblemUnused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.nextFieldsForParameterProblem.unused-13"]]
}
func (s *Icmpv1Stack) Next4BytesGatewayInternetAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv1AliasToFieldIdx["message.next4Bytes.gatewayInternetAddress-14"]]
}

func (s *Icmpv1Stack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewIcmpv1Stack(idx int) *Icmpv1Stack {
	stack := Icmpv1Stack(newStack(idx, "icmpv1", icmpv1AliasToFieldIdx))
	return &stack
}

type Icmpv2Stack TrafficTrafficItemConfigElementStack

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

func (s *Icmpv2Stack) MessageType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.messageType-1"]]
}
func (s *Icmpv2Stack) CodeValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.codeValue-2"]]
}
func (s *Icmpv2Stack) IcmpChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.icmpChecksum-3"]]
}
func (s *Icmpv2Stack) Identifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.identifier-4"]]
}
func (s *Icmpv2Stack) SequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.sequenceNumber-5"]]
}
func (s *Icmpv2Stack) NextFieldsNone() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-6"]]
}
func (s *Icmpv2Stack) NextFieldsNone_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-7"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampMsgOrigTmpStmp1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampMsg.origTmpStmp1-8"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampMsgRcvTmpStmp1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampMsg.rcvTmpStmp1-9"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampMsgTransTmpStmp1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampMsg.transTmpStmp1-10"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampReplyOrigTmpStmp2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampReply.origTmpStmp2-11"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampReplyRcvTmpStmp2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampReply.rcvTmpStmp2-12"]]
}
func (s *Icmpv2Stack) NextFieldsFieldsForTimeStampReplyTransTmpStmp2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.fieldsForTimeStampReply.transTmpStmp2-13"]]
}
func (s *Icmpv2Stack) NextFieldsNone_3() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-14"]]
}
func (s *Icmpv2Stack) NextFieldsNone_4() *TrafficTrafficItemConfigElementStackField {
	return s.Field[icmpv2AliasToFieldIdx["message.nextFields.none-15"]]
}

func (s *Icmpv2Stack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewIcmpv2Stack(idx int) *Icmpv2Stack {
	stack := Icmpv2Stack(newStack(idx, "icmpv2", icmpv2AliasToFieldIdx))
	return &stack
}

type Ipv4Stack TrafficTrafficItemConfigElementStack

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

func (s *Ipv4Stack) Version() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.version-1"]]
}
func (s *Ipv4Stack) HeaderLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.headerLength-2"]]
}
func (s *Ipv4Stack) PriorityRaw() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.raw-3"]]
}
func (s *Ipv4Stack) PriorityTosPrecedence() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.precedence-4"]]
}
func (s *Ipv4Stack) PriorityTosDelay() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.delay-5"]]
}
func (s *Ipv4Stack) PriorityTosThroughput() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.throughput-6"]]
}
func (s *Ipv4Stack) PriorityTosReliability() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.reliability-7"]]
}
func (s *Ipv4Stack) PriorityTosMonetary() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.monetary-8"]]
}
func (s *Ipv4Stack) PriorityTosUnused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.tos.unused-9"]]
}
func (s *Ipv4Stack) PriorityDsPhbDefaultPHBDefaultPHB() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.defaultPHB.defaultPHB-10"]]
}
func (s *Ipv4Stack) PriorityDsPhbDefaultPHBUnused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.defaultPHB.unused-11"]]
}
func (s *Ipv4Stack) PriorityDsPhbClassSelectorPHBClassSelectorPHB() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.classSelectorPHB.classSelectorPHB-12"]]
}
func (s *Ipv4Stack) PriorityDsPhbClassSelectorPHBUnused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.classSelectorPHB.unused-13"]]
}
func (s *Ipv4Stack) PriorityDsPhbAssuredForwardingPHBAssuredForwardingPHB() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.assuredForwardingPHB.assuredForwardingPHB-14"]]
}
func (s *Ipv4Stack) PriorityDsPhbAssuredForwardingPHBUnused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.assuredForwardingPHB.unused-15"]]
}
func (s *Ipv4Stack) PriorityDsPhbExpeditedForwardingPHBExpeditedForwardingPHB() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.expeditedForwardingPHB.expeditedForwardingPHB-16"]]
}
func (s *Ipv4Stack) PriorityDsPhbExpeditedForwardingPHBUnused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.priority.ds.phb.expeditedForwardingPHB.unused-17"]]
}
func (s *Ipv4Stack) TotalLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.totalLength-18"]]
}
func (s *Ipv4Stack) Identification() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.identification-19"]]
}
func (s *Ipv4Stack) FlagsReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.flags.reserved-20"]]
}
func (s *Ipv4Stack) FlagsFragment() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.flags.fragment-21"]]
}
func (s *Ipv4Stack) FlagsLastFragment() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.flags.lastFragment-22"]]
}
func (s *Ipv4Stack) FragmentOffset() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.fragmentOffset-23"]]
}
func (s *Ipv4Stack) Ttl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.ttl-24"]]
}
func (s *Ipv4Stack) Protocol() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.protocol-25"]]
}
func (s *Ipv4Stack) Checksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.checksum-26"]]
}
func (s *Ipv4Stack) SrcIp() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.srcIp-27"]]
}
func (s *Ipv4Stack) DstIp() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.dstIp-28"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionNop() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.nop-29"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.type-30"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.length-31"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecuritySecurity() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.security-32"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityCompartments() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.compartments-33"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityHandling() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.handling-34"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSecurityTcc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.security.tcc-35"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionLsrrType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.lsrr.type-36"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionLsrrLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.lsrr.length-37"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionPointer() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.pointer-38"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouteData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routeData-39"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSsrrType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.ssrr.type-40"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionSsrrLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.ssrr.length-41"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRecordRouteType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.recordRoute.type-42"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRecordRouteLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.recordRoute.length-43"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionStreamIdType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.streamId.type-44"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionStreamIdLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.streamId.length-45"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionStreamIdId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.streamId.id-46"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.type-47"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.length-48"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampPointer() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.pointer-49"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampOverflow() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.overflow-50"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampFlags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.flags-51"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampPairAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.pair.address-52"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionTimestampPairTimestamp() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.timestamp.pair.timestamp-53"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionLast() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.last-54"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouterAlertType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routerAlert.type-55"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouterAlertLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routerAlert.length-56"]]
}
func (s *Ipv4Stack) OptionsNextOptionOptionRouterAlertValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.nextOption.option.routerAlert.value-57"]]
}
func (s *Ipv4Stack) OptionsPad() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv4AliasToFieldIdx["header.options.pad-58"]]
}

func (s *Ipv4Stack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewIpv4Stack(idx int) *Ipv4Stack {
	stack := Ipv4Stack(newStack(idx, "ipv4", ipv4AliasToFieldIdx))
	return &stack
}

type Ipv6Stack TrafficTrafficItemConfigElementStack

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

func (s *Ipv6Stack) VersionTrafficClassFlowLabelVersion() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.versionTrafficClassFlowLabel.version-1"]]
}
func (s *Ipv6Stack) VersionTrafficClassFlowLabelTrafficClass() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.versionTrafficClassFlowLabel.trafficClass-2"]]
}
func (s *Ipv6Stack) VersionTrafficClassFlowLabelFlowLabel() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.versionTrafficClassFlowLabel.flowLabel-3"]]
}
func (s *Ipv6Stack) PayloadLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.payloadLength-4"]]
}
func (s *Ipv6Stack) NextHeader() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.nextHeader-5"]]
}
func (s *Ipv6Stack) HopLimit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.hopLimit-6"]]
}
func (s *Ipv6Stack) SrcIP() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.srcIP-7"]]
}
func (s *Ipv6Stack) DstIP() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ipv6AliasToFieldIdx["header.dstIP-8"]]
}

func (s *Ipv6Stack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewIpv6Stack(idx int) *Ipv6Stack {
	stack := Ipv6Stack(newStack(idx, "ipv6", ipv6AliasToFieldIdx))
	return &stack
}

type LdpHelloStack TrafficTrafficItemConfigElementStack

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

func (s *LdpHelloStack) Version() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.version-1"]]
}
func (s *LdpHelloStack) PduLengthinOctets() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.pduLengthinOctets-2"]]
}
func (s *LdpHelloStack) LsrID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.lsrID-3"]]
}
func (s *LdpHelloStack) LabelSpace() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.labelSpace-4"]]
}
func (s *LdpHelloStack) UBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.uBit-5"]]
}
func (s *LdpHelloStack) Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.type-6"]]
}
func (s *LdpHelloStack) Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.length-7"]]
}
func (s *LdpHelloStack) MessageID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.messageID-8"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.uBit-9"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.fBit-10"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.type-11"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.length-12"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVHoldTime() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.holdTime-13"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVTBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.tBit-14"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVRBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.rBit-15"]]
}
func (s *LdpHelloStack) CommonHelloParametersTLVReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.commonHelloParametersTLV.reserved-16"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.uBit-17"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.fBit-18"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.type-19"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.length-20"]]
}
func (s *LdpHelloStack) OptionalParameterIpv4TransportAddressTLVIpv4Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv4TransportAddressTLV.ipv4Address-21"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.uBit-22"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.fBit-23"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.type-24"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.length-25"]]
}
func (s *LdpHelloStack) OptionalParameterConfigurationSequenceNumberTLVSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.configurationSequenceNumberTLV.sequenceNumber-26"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.uBit-27"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.fBit-28"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.type-29"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.length-30"]]
}
func (s *LdpHelloStack) OptionalParameterIpv6TransportAddressTLVIpv6Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ldpHelloAliasToFieldIdx["header.optionalParameter.ipv6TransportAddressTLV.ipv6Address-31"]]
}

func (s *LdpHelloStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewLdpHelloStack(idx int) *LdpHelloStack {
	stack := LdpHelloStack(newStack(idx, "ldpHello", ldpHelloAliasToFieldIdx))
	return &stack
}

type PimHelloMessageStack TrafficTrafficItemConfigElementStack

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

func (s *PimHelloMessageStack) Version() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.version-1"]]
}
func (s *PimHelloMessageStack) Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.type-2"]]
}
func (s *PimHelloMessageStack) Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.reserved-3"]]
}
func (s *PimHelloMessageStack) HeaderChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.headerChecksum-4"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsHoldtimeType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.holdtime.type-5"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsHoldtimeLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.holdtime.length-6"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsHoldtimeHoldtimesec() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.holdtime.holdtimesec-7"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.type-8"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.length-9"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayTBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.tBit-10"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayLanDelay() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.lanDelay-11"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsLanPruneDelayOverrideInterval() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.lanPruneDelay.overrideInterval-12"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsDrPriorityType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.drPriority.type-13"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsDrPriorityLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.drPriority.length-14"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsDrPriorityDrPriority() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.drPriority.drPriority-15"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsGenerationIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.generationID.type-16"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsGenerationIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.generationID.length-17"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsGenerationIDGenerationID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.generationID.generationID-18"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsBidirCapableType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.bidirCapable.type-19"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsBidirCapableLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.bidirCapable.length-20"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsPrivateUsageFieldType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.privateUsageField.type-21"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsPrivateUsageFieldLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.privateUsageField.length-22"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsPrivateUsageFieldValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.privateUsageField.value-23"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsUserDefinedFieldType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.userDefinedField.type-24"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsUserDefinedFieldLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.userDefinedField.length-25"]]
}
func (s *PimHelloMessageStack) HelloOptionsFieldsOptionsUserDefinedFieldValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[pimHelloMessageAliasToFieldIdx["header.helloOptionsFields.options.userDefinedField.value-26"]]
}

func (s *PimHelloMessageStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewPimHelloMessageStack(idx int) *PimHelloMessageStack {
	stack := PimHelloMessageStack(newStack(idx, "pimHelloMessage", pimHelloMessageAliasToFieldIdx))
	return &stack
}

type RsvpStack TrafficTrafficItemConfigElementStack

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

func (s *RsvpStack) Version() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.version-1"]]
}
func (s *RsvpStack) Flag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.flag-2"]]
}
func (s *RsvpStack) MessegeType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.messegeType-3"]]
}
func (s *RsvpStack) RsvpChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.rsvpChecksum-4"]]
}
func (s *RsvpStack) Ttl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.ttl-5"]]
}
func (s *RsvpStack) Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.reserved-6"]]
}
func (s *RsvpStack) RsvpLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.rsvpLength-7"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderVersion() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.version-8"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderFlag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.flag-9"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderMessegeType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.messegeType-10"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderRsvpChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.rsvpChecksum-11"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderTtl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.ttl-12"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.reserved-13"]]
}
func (s *RsvpStack) ObjectsObjectBundleMsgOptionalHeaderRsvpLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.bundleMsgOptionalHeader.rsvpLength-14"]]
}
func (s *RsvpStack) ObjectsObjectObjectLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectLength-15"]]
}
func (s *RsvpStack) ObjectsObjectClass() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.class-16"]]
}
func (s *RsvpStack) ObjectsObjectType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.type-17"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDataMessegeDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.dataMessege.dataLength-18"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDataMessegeData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.dataMessege.data-19"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1DestAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.destAddress-20"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1ProtocolId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.protocolId-21"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.flags-22"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4UDPSESSIONClass1CType1DestPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4UDPSESSIONClass1CType1.destPort-23"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2DestAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.destAddress-24"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2ProtocolId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.protocolId-25"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.flags-26"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6UDPSESSIONClass1CType2DestPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6UDPSESSIONClass1CType2.destPort-27"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3DestAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.destAddress-28"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3ProtocolId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.protocolId-29"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.flags-30"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISESSIONClass1CType3DestPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISESSIONClass1CType3.destPort-31"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4DestAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.destAddress-32"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4ProtocolId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.protocolId-33"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.flags-34"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISESSIONClass1CType4DestPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISESSIONClass1CType4.destPort-35"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7Ipv4TunnelEndPointAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.ipv4TunnelEndPointAddress-36"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.reserved-37"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7TunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.tunnelId-38"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4Class1CType7ExtendedTunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4Class1CType7.extendedTunnelId-39"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8Ipv6TunnelEndPointAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.ipv6TunnelEndPointAddress-40"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.reserved-41"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8TunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.tunnelId-42"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6Class1CType8ExtendedTunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6Class1CType8.extendedTunnelId-43"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Ipv4SessionAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.ipv4SessionAddress-44"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.reserved-45"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Flag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.flag-46"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.unused-47"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip4class1CType9Dscp() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip4class1CType9.dscp-48"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Ipv6SessionAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.ipv6SessionAddress-49"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.reserved-50"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Flag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.flag-51"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.unused-52"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvpaggregateip6class1CType10Dscp() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvpaggregateip6class1CType10.dscp-53"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13P2mpId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.p2mpId-54"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.reserved-55"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13TunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.tunnelId-56"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4Class1CType13ExtendedTunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4Class1CType13.extendedTunnelId-57"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14P2mpId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.p2mpId-58"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.reserved-59"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14TunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.tunnelId-60"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6Class1CType14ExtendedTunnelId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6Class1CType14.extendedTunnelId-61"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType1Ipv4NextPreviousHopAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType1.ipv4NextPreviousHopAddress-62"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType1LogicalInterfaceHandle() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType1.logicalInterfaceHandle-63"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType2Ipv6NextPreviousHopAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType2.ipv6NextPreviousHopAddress-64"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRsvphopClassClass3CType2LogicalInterfaceHandle() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.rsvphopClassClass3CType2.logicalInterfaceHandle-65"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.flags-66"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.reserved-67"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1KeyId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.keyId-68"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1SequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.sequenceNumber-69"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1MsgLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.msgLength-70"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntegrityclass4CType1KeyedMessege() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.integrityclass4CType1.keyedMessege-71"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyTimevaluesClassClass5CType1RefreshPeriodR() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.timevaluesClassClass5CType1.refreshPeriodR-72"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1Ipv4ErrorNodeAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.ipv4ErrorNodeAddress-73"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.flags-74"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1ErrorCode() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.errorCode-75"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ERRORSPECClass6CType1ErrorValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ERRORSPECClass6CType1.errorValue-76"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2Ipv6ErrorNodeAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.ipv6ErrorNodeAddress-77"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.flags-78"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2ErrorCode() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.errorCode-79"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ERRORSPECClass6CType2ErrorValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ERRORSPECClass6CType2.errorValue-80"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4ScopeClassClass7CType1Ipv4SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4ScopeClassClass7CType1.ipv4SrcAddress-81"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6ScopeClassClass7CType2Ipv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6ScopeClassClass7CType2.ipv6SrcAddress-82"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.flags-83"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.reserved-84"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1SharingControl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.sharingControl-85"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyStyleClassClass8CType1SenderSelectionControl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.styleClassClass8CType1.senderSelectionControl-86"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4SignalType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.signalType-87"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Rcc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.rcc-88"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Ncc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.ncc-89"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Nvc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.nvc-90"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Multiplier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.multiplier-91"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Transparency() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.transparency-92"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass9CType4Profile() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass9CType4.profile-93"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5SignalType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.signalType-94"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.reserved-95"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Nmc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.nmc-96"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Nvc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.nvc-97"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Multiplier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.multiplier-98"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class9CType5Reserved_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class9CType5.reserved-99"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType1Ipv4SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType1.ipv4SrcAddress-100"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType1Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType1.unused-101"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType1SrcPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType1.srcPort-102"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType2Ipv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType2.ipv6SrcAddress-103"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType2Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType2.unused-104"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType2SrcPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType2.srcPort-105"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType3Ipv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType3.ipv6SrcAddress-106"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType3Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType3.unused-107"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFilterspecclass10CType3FlowLabel() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.filterspecclass10CType3.flowLabel-108"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPIFILTERSPECClass10CType4Ipv4SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPIFILTERSPECClass10CType4.ipv4SrcAddress-109"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPIFILTERSPECClass10CType4Gpi() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPIFILTERSPECClass10CType4.gpi-110"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPIFILTERSPECClass10CType5Ipv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPIFILTERSPECClass10CType5.ipv6SrcAddress-111"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPIFILTERSPECClass10CType5Gpi() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPIFILTERSPECClass10CType5.gpi-112"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4FILTERSPECClass10CType7Ipv4TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.ipv4TunnelSenderAddress-113"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4FILTERSPECClass10CType7Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.unused-114"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4FILTERSPECClass10CType7LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4FILTERSPECClass10CType7.lspID-115"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6FILTERSPECClass10CType8Ipv6TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.ipv6TunnelSenderAddress-116"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6FILTERSPECClass10CType8Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.unused-117"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6FILTERSPECClass10CType8LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6FILTERSPECClass10CType8.lspID-118"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12Ipv4TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.ipv4TunnelSenderAddress-119"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.unused-120"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.lspID-121"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12SubGroupOriginatorID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.subGroupOriginatorID-122"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12Unused_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.unused-123"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv4FILTERSPECClass10CType12SubGroupID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv4FILTERSPECClass10CType12.subGroupID-124"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13Ipv6TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.ipv6TunnelSenderAddress-125"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.unused-126"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.lspID-127"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13SubGroupOriginatorID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.subGroupOriginatorID-128"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13Unused_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.unused-129"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mpLSPIPv6FILTERSPECClass10CType13SubGroupID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mpLSPIPv6FILTERSPECClass10CType13.subGroupID-130"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4SENDERTEMPLATEClass11CType1Ipv4SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.ipv4SrcAddress-131"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4SENDERTEMPLATEClass11CType1Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.unused-132"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4SENDERTEMPLATEClass11CType1SrcPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4SENDERTEMPLATEClass11CType1.srcPort-133"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6SENDERTEMPLATEClass11CType2Ipv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.ipv6SrcAddress-134"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6SENDERTEMPLATEClass11CType2Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.unused-135"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6SENDERTEMPLATEClass11CType2SrcPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6SENDERTEMPLATEClass11CType2.srcPort-136"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4FlowlabelSENDERTEMPLATEClass11CType3Ipv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.ipv6SrcAddress-137"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4FlowlabelSENDERTEMPLATEClass11CType3Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.unused-138"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4FlowlabelSENDERTEMPLATEClass11CType3FlowLabel() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4FlowlabelSENDERTEMPLATEClass11CType3.flowLabel-139"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISENDERTEMPLATEClass11CType4Ipv4SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISENDERTEMPLATEClass11CType4.ipv4SrcAddress-140"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4GPISENDERTEMPLATEClass11CType4Gpi() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4GPISENDERTEMPLATEClass11CType4.gpi-141"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISENDERTEMPLATEClass11CType5Ipv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISENDERTEMPLATEClass11CType5.ipv6SrcAddress-142"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6GPISENDERTEMPLATEClass11CType5Gpi() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6GPISENDERTEMPLATEClass11CType5.gpi-143"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4SENDERTEMPLATEClass11CType7Ipv4TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.ipv4TunnelSenderAddress-144"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4SENDERTEMPLATEClass11CType7Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.unused-145"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv4SENDERTEMPLATEClass11CType7LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv4SENDERTEMPLATEClass11CType7.lspID-146"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6SENDERTEMPLATEClass11CType8Ipv6TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.ipv6TunnelSenderAddress-147"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6SENDERTEMPLATEClass11CType8Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.unused-148"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelipv6SENDERTEMPLATEClass11CType8LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelipv6SENDERTEMPLATEClass11CType8.lspID-149"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12Ipv4TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.ipv4TunnelSenderAddress-150"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.unused-151"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.lspID-152"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12SubGroupOriginatorID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.subGroupOriginatorID-153"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12Unused_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.unused-154"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv4SENDERTEMPLATEClass11CType12SubGroupID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv4SENDERTEMPLATEClass11CType12.subGroupID-155"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13Ipv6TunnelSenderAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.ipv6TunnelSenderAddress-156"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.unused-157"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13LspID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.lspID-158"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13SubGroupOriginatorID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.subGroupOriginatorID-159"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13Unused_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.unused-160"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyP2mplsptunnelipv6SENDERTEMPLATEClass11CType13SubGroupID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.p2mplsptunnelipv6SENDERTEMPLATEClass11CType13.subGroupID-161"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2VersionNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.versionNumber-162"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Reserved1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.reserved1-163"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2OverallLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.overallLength-164"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2ServiceHeader() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.serviceHeader-165"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2ZeroBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.zeroBit-166"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Reserved2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.reserved2-167"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2LengthOfService1Data() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.lengthOfService1Data-168"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2ParameterIDTokenBucketTSpec() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameterIDTokenBucketTSpec-169"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Parameter127Flag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameter127Flag-170"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2Parameter127Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.parameter127Length-171"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2TokenBucketRate() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.tokenBucketRate-172"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2TokenBucketSize() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.tokenBucketSize-173"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2PeakDataRate() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.peakDataRate-174"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2MinimumPolicedUnit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.minimumPolicedUnit-175"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservSENDERTSPECTEMPLATEClass12CType2MaximumPacketSize() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservSENDERTSPECTEMPLATEClass12CType2.maximumPacketSize-176"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4SignalType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.signalType-177"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Rcc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.rcc-178"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Ncc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.ncc-179"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Nvc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.nvc-180"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Multiplier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.multiplier-181"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Transparency() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.transparency-182"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySonetsdhClass12CType4Profile() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.sonetsdhClass12CType4.profile-183"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5SignalType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.signalType-184"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.reserved-185"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Nmc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.nmc-186"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Nvc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.nvc-187"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Multiplier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.multiplier-188"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyG709Class12CType5Reserved_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.g709Class12CType5.reserved-189"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2MessageFormatVersionNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.messageFormatVersionNumber-190"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.reserved-191"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2MsgLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.msgLength-192"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1PerServiceHeaderServiceNumber1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.perServiceHeaderServiceNumber1-193"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1X() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.x-194"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Reserved3() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.reserved3-195"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1GlobalBreakBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.globalBreakBit-196"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID4() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID4-197"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter4FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter4FlagByte-198"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter4Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter4Length-199"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1IsHopCnt() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.isHopCnt-200"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID6() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID6-201"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter6FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter6FlagByte-202"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter6Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter6Length-203"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1PathBwEstimate() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.pathBwEstimate-204"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID8() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID8-205"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter8FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter8FlagByte-206"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter8Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter8Length-207"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1MinimumPathLatency() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.minimumPathLatency-208"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ParameterID10() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameterID10-209"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter10FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter10FlagByte-210"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1Parameter10Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.parameter10Length-211"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GeneralParametersFragmentService1ComposedMTU() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.generalParametersFragmentService1.composedMTU-212"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2PerServiceHeaderServiceNumber2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.perServiceHeaderServiceNumber2-213"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2XBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.xBit-214"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.reserved-215"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2BreakBitAndLengthOfPerserviceData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.breakBitAndLengthOfPerserviceData-216"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter133ComposedCtot() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter133ComposedCtot-217"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter133FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter133FlagByte-218"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter133Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter133Length-219"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsEndtoendComposedValueForCCtot() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.endtoendComposedValueForCCtot-220"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter134ComposedDtot() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter134ComposedDtot-221"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter134FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter134FlagByte-222"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter134Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter134Length-223"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsEndtoendComposedValueForDDtot() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.endtoendComposedValueForDDtot-224"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter135ComposedCsum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter135ComposedCsum-225"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter135FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter135FlagByte-226"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter135Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter135Length-227"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsSincelastreshapingPointComposedCCsum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.sincelastreshapingPointComposedCCsum-228"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameterIDParameter136ComposedDsum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameterIDParameter136ComposedDsum-229"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter136FlagByte() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter136FlagByte-230"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsParameter136Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.parameter136Length-231"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsSincelastreshapingPointComposedDDsum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.sincelastreshapingPointComposedDDsum-232"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2GuaranteedServiceFragmentService2OptionalFieldsServicespecificGeneralParameterHeadersvaluesServicespecificGeneralParameterHeadervalue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.guaranteedServiceFragmentService2.optionalFields.servicespecificGeneralParameterHeadersvalues.servicespecificGeneralParameterHeadervalue-233"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentPerServiceHeaderServiceNumber5() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.perServiceHeaderServiceNumber5-234"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentXBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.xBit-235"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentBreakBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.breakBit-236"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentLengthOfPerserviceData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.lengthOfPerserviceData-237"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIntservADSPECTEMPLATEClass13CType2ControlledLoadServiceDataFragmentServicespecificGeneralParameterHeadersServicespecificGeneralParameterHeader() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.intservADSPECTEMPLATEClass13CType2.controlledLoadServiceDataFragment.servicespecificGeneralParameterHeaders.servicespecificGeneralParameterHeader-238"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.length-239"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Policydata() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policydata-240"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Unused() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.unused-241"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1DataOffset() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.dataOffset-242"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1Unused_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.unused-243"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1OptionDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.optionDataLength-244"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1OptionData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.optionData-245"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1PolicyDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policyDataLength-246"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyPolicydataObjectClass14CType1PolicyData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.policydataObjectClass14CType1.policyData-247"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4RESVCONFIRMClass15CType1Ipv4ReceiverAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4RESVCONFIRMClass15CType1.ipv4ReceiverAddress-248"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6RESVCONFIRMClass15CType2Ipv6ReceiverAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6RESVCONFIRMClass15CType2.ipv6ReceiverAddress-249"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectClass16CType1TopLabel() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectClass16CType1.topLabel-250"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelRequestWithoutLabelRangeClass19CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelRequestWithoutLabelRangeClass19CType1.reserved-251"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelRequestWithoutLabelRangeClass19CType1L3pid() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelRequestWithoutLabelRangeClass19CType1.l3pid-252"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-253"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2L3pid() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.l3pid-254"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.mBit-255"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2Reserved_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-256"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MinimumVPI() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.minimumVPI-257"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MinimumVCI() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.minimumVCI-258"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2Reserved_3() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.reserved-259"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MaximumVPI() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.maximumVPI-260"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithATMLabelRangeClass19CType2MaximumVCI() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithATMLabelRangeClass19CType2.maximumVCI-261"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.reserved-262"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3L3pid() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.l3pid-263"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Res() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.res-264"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Dli() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.dli-265"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3MinimumDLCI() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.minimumDLCI-266"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3Res_2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.res-267"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLabelObjectWithFrameRelayLabelRangeClass19CType3MaximumDLCI() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.labelObjectWithFrameRelayLabelRangeClass19CType3.maximumDLCI-268"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1LBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.lBit-269"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.type-270"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.length-271"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Ipv4Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.ipv4Address-272"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.prefixLength-273"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype1Padding() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype1.padding-274"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2LBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.lBit-275"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.type-276"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.length-277"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Ipv6Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.ipv6Address-278"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.prefixLength-279"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype2Padding() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype2.padding-280"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32LBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.lBit-281"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.type-282"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.length-283"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyExplicitrouteClass20SubTypeCtype32AsNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.explicitrouteClass20.subType.ctype32.asNumber-284"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.type-285"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.length-286"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Ipv4Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.ipv4Address-287"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.prefixLength-288"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype1.flags-289"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.type-290"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.length-291"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Ipv6Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.ipv6Address-292"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.prefixLength-293"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype2Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype2.flags-294"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.type-295"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.length-296"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.flags-297"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.ctype-298"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRecordrouteClass21SubTypeCtype3ContentsOfLabelObject() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.recordrouteClass21.subType.ctype3.contentsOfLabelObject-299"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyHelloREQUESTAckClass22CType12SrcInstance() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.helloREQUESTAckClass22CType12.srcInstance-300"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyHelloREQUESTAckClass22CType12DestInstance() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.helloREQUESTAckClass22CType12.destInstance-301"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidclass23CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidclass23CType1.flags-302"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidclass23CType1Epoch() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidclass23CType1.epoch-303"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidclass23CType1MessageIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidclass23CType1.messageIdentifier-304"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidAckNackclass24CType12Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.flags-305"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidAckNackclass24CType12Epoch() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.epoch-306"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidAckNackclass24CType12MessageIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidAckNackclass24CType12.messageIdentifier-307"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidlistclass25CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.flags-308"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidlistclass25CType1Epoch() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.epoch-309"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidlistclass25CType1MessageidlistMessageIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidlistclass25CType1.messageidlist.messageIdentifier-310"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.flags-311"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2Epoch() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.epoch-312"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2MessageidlistMessageIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.messageidlist.messageIdentifier-313"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv4class25CType2MessageidlistSourceIPAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv4class25CType2.messageidlist.sourceIPAddress-314"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.flags-315"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3Epoch() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.epoch-316"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3MessageidlistMessageIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.messageidlist.messageIdentifier-317"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidsourceListIPv6class25CType3MessageidlistSourceIPAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidsourceListIPv6class25CType3.messageidlist.sourceIPAddress-318"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.flags-319"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4Epoch() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.epoch-320"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4MessageidlistMessageIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.messageIdentifier-321"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4MessageidlistSourceIPAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.sourceIPAddress-322"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv4class25CType4MessageidlistDestinationIPAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv4class25CType4.messageidlist.destinationIPAddress-323"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.flags-324"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5Epoch() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.epoch-325"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5MessageidlistMessageIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.messageIdentifier-326"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5MessageidlistSourceIPAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.sourceIPAddress-327"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyMessageidmcastlistIPv6class25CType5MessageidlistDestinationIPAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.messageidmcastlistIPv6class25CType5.messageidlist.destinationIPAddress-328"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1MaxRSVPhops() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.maxRSVPhops-329"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1Rsvphopcount() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.rsvphopcount-330"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.reserved-331"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1MfBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.mfBit-332"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1RequestIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.requestIdentifier-333"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1PathMTU() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.pathMTU-334"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1FragmentOffset() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.fragmentOffset-335"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1Lasthopaddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.lasthopaddress-336"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1SenderTemplateObjectIpv4SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.senderTemplateObject.ipv4SrcAddress-337"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1SenderTemplateObjectGeneralizedPortIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.senderTemplateObject.generalizedPortIdentifier-338"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1FilterSpecTemplateObjectIpv4SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.filterSpecTemplateObject.ipv4SrcAddress-339"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv4DIAGNOSTICclass30CType1FilterSpecTemplateObjectGeneralizedPortIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv4DIAGNOSTICclass30CType1.filterSpecTemplateObject.generalizedPortIdentifier-340"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2MaxRSVPhops() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.maxRSVPhops-341"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2Rsvphopcount() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.rsvphopcount-342"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.reserved-343"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2MfBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.mfBit-344"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2RequestIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.requestIdentifier-345"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2PathMTU() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.pathMTU-346"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2FragmentOffset() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.fragmentOffset-347"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2Lasthopaddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.lasthopaddress-348"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2SenderTemplateObjectIpv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.senderTemplateObject.ipv6SrcAddress-349"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2SenderTemplateObjectGeneralizedPortIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.senderTemplateObject.generalizedPortIdentifier-350"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2FilterSpecTemplateObjectIpv6SrcAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.filterSpecTemplateObject.ipv6SrcAddress-351"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyIpv6DIAGNOSTICclass30CType2FilterSpecTemplateObjectGeneralizedPortIdentifier() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.ipv6DIAGNOSTICclass30CType2.filterSpecTemplateObject.generalizedPortIdentifier-352"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagselectclass33CType1ClassTypeClass() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagselectclass33CType1.classType.class-353"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagselectclass33CType1ClassTypeCtype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagselectclass33CType1.classType.ctype-354"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv4Objectclass31CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.reserved-355"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv4Objectclass31CType1RPointer() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.rPointer-356"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv4Objectclass31CType1NodeAddressListRsvpNodeAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv4Objectclass31CType1.nodeAddressList.rsvpNodeAddress-357"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv6Objectclass31CType2Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.reserved-358"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv6Objectclass31CType2RPointer() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.rPointer-359"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyRouteIPv6Objectclass31CType2NodeAddressListRsvpNodeAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.routeIPv6Objectclass31CType2.nodeAddressList.rsvpNodeAddress-360"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1DreqArrivalTime() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.dreqArrivalTime-361"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1IncomingInterfaceAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.incomingInterfaceAddress-362"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1OutgoingInterfaceAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.outgoingInterfaceAddress-363"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1PreviousRSVPHopRouterAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.previousRSVPHopRouterAddress-364"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1Dttl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.dttl-365"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1MBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.mBit-366"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1Rerr() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.rerr-367"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1K() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.k-368"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseclass32CType1TimerValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseclass32CType1.timerValue-369"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2DreqArrivalTime() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.dreqArrivalTime-370"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2IncomingInterfaceAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.incomingInterfaceAddress-371"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2OutgoingInterfaceAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.outgoingInterfaceAddress-372"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2PreviousRSVPHopRouterAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.previousRSVPHopRouterAddress-373"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2Dttl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.dttl-374"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2MBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.mBit-375"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2Rerr() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.rerr-376"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2K() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.k-377"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiagresponseIPv6class32CType2TimerValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diagresponseIPv6class32CType2.timerValue-378"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyS2lsublspipv4Class50CType1Ipv4S2LSubLSPDestinationAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.s2lsublspipv4Class50CType1.ipv4S2LSubLSPDestinationAddress-379"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyS2lsublspipv6Class50CType2Ipv6S2LSubLSPDestinationAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.s2lsublspipv6Class50CType2.ipv6S2LSubLSPDestinationAddress-380"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7Lengthbytes() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.lengthbytes-381"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7ClassNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.classNum-382"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.ctype-383"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7PlrAddressListPlrID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.plrAddressList.plrID-384"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv4class63CType7PlrAddressListAvoidNodeID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv4class63CType7.plrAddressList.avoidNodeID-385"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8Lengthbytes() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.lengthbytes-386"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8ClassNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.classNum-387"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.ctype-388"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8PlrAddressListPlrID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.plrAddressList.plrID-389"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDetourObjectIPv6class63CType8PlrAddressListAvoidNodeID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.detourObjectIPv6class63CType8.plrAddressList.avoidNodeID-390"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyChallengeObjectclass64CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.reserved-391"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyChallengeObjectclass64CType1KeyId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.keyId-392"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyChallengeObjectclass64CType1ChallengeCookie() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.challengeObjectclass64CType1.challengeCookie-393"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.reserved-394"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1Mapnb() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapnb-395"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1MapListReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.reserved-396"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1MapListExp() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.exp-397"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservELSPclass65CType1MapListPhbid() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservELSPclass65CType1.mapList.phbid-398"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservLLSPclass65CType2Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservLLSPclass65CType2.reserved-399"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyDiffservLLSPclass65CType2Psc() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.diffservLLSPclass65CType2.psc-400"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyClasstypeclass66CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.classtypeclass66CType1.reserved-401"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyClasstypeclass66CType1Ct() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.classtypeclass66CType1.ct-402"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelinterfaceidclass193CType1LsrId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelinterfaceidclass193CType1.lsrId-403"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelinterfaceidclass193CType1InterfaceID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelinterfaceidclass193CType1.interfaceID-404"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1LBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.lBit-405"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.type-406"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.length-407"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Ipv4Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.ipv4Address-408"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.prefixLength-409"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype1Padding() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype1.padding-410"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2LBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.lBit-411"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.type-412"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.length-413"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Ipv6Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.ipv6Address-414"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.prefixLength-415"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype2Padding() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype2.padding-416"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32LBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.lBit-417"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.type-418"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.length-419"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryexplicitrouteClass200CType2SubTypeCtype32AsNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryexplicitrouteClass200CType2.subType.ctype32.asNumber-420"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.type-421"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.length-422"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Ipv4Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.ipv4Address-423"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.prefixLength-424"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype1.flags-425"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.type-426"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.length-427"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Ipv6Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.ipv6Address-428"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2PrefixLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.prefixLength-429"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype2Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype2.flags-430"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Type() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.type-431"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.length-432"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.flags-433"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.ctype-434"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodySecondaryrecordrouteClass201CType2SubTypeCtype3ContentsOfLabelObject() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.secondaryrecordrouteClass201CType2.subType.ctype3.contentsOfLabelObject-435"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Lengthbytes() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.lengthbytes-436"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1ClassNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.classNum-437"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.ctype-438"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1SetupPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.setupPrio-439"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1HoldingPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.holdingPrio-440"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Hoplimit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.hoplimit-441"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.flags-442"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Bandwidth() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.bandwidth-443"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Includeany() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.includeany-444"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Excludeany() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.excludeany-445"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType1Includeall() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType1.includeall-446"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Lengthbytes() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.lengthbytes-447"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7ClassNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.classNum-448"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.ctype-449"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7SetupPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.setupPrio-450"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7HoldingPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.holdingPrio-451"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Hoplimit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.hoplimit-452"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.reserved-453"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Bandwidth() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.bandwidth-454"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Includeany() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.includeany-455"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyFastrerouteclass205CType7Excludeany() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.fastrerouteclass205CType7.excludeany-456"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7SetupPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.setupPrio-457"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7HoldingPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.holdingPrio-458"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.flags-459"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7NameLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.nameLength-460"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelsessionattributeclass207CType7SessionName() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelsessionattributeclass207CType7.sessionName-461"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Excludeany() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.excludeany-462"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Includeany() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.includeany-463"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Includeall() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.includeall-464"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1SetupPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.setupPrio-465"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1HoldingPrio() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.holdingPrio-466"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.flags-467"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1NameLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.nameLength-468"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyLsptunnelrasessionattributeclass207CType1SessionName() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.lsptunnelrasessionattributeclass207CType1.sessionName-469"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyAtmserviceclassclass227CType1Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.atmserviceclassclass227CType1.reserved-470"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyAtmserviceclassclass227CType1Flags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.atmserviceclassclass227CType1.flags-471"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2Lengthbytes() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.lengthbytes-472"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2ClassNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.classNum-473"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.ctype-474"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2Resv() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.resv-475"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallCapabilityObjectclass228CType2CallOpsFlag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callCapabilityObjectclass228CType2.callOpsFlag-476"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.uBit-477"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.fBit-478"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDSourceIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.sourceIDType-479"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.length-480"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDIpv4Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.ipv4Address-481"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4SourceIDLogicalPortId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4SourceID.logicalPortId-482"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.uBit-483"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.fBit-484"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDSourceIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.sourceIDType-485"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.length-486"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDIpv6Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.ipv6Address-487"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6SourceIDLogicalPortId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6SourceID.logicalPortId-488"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.uBit-489"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.fBit-490"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDSourceIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.sourceIDType-491"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.length-492"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.dataLength-493"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDNsap() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.nsap-494"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapSourceIDLogicalPortId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapSourceID.logicalPortId-495"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.uBit-496"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.fBit-497"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDDestIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.destIDType-498"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.length-499"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDIpv4Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.ipv4Address-500"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv4DestIDLogicalPortId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv4DestID.logicalPortId-501"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.uBit-502"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.fBit-503"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDDestIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.destIDType-504"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.length-505"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDIpv6Address() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.ipv6Address-506"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeIpv6DestIDLogicalPortId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.ipv6DestID.logicalPortId-507"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.uBit-508"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.fBit-509"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDDestIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.destIDType-510"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.length-511"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.dataLength-512"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDNsap() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.nsap-513"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeNsapDestIDLogicalPortId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.nsapDestID.logicalPortId-514"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.uBit-515"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.fBit-516"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVEgressIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.egressIDType-517"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.length-518"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.reserved-519"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLbit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.lbit-520"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLogicalPortId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.logicalPortId-521"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLabelLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.labelLength-522"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeEgressLabelTLVLabel() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.egressLabelTLV.label-523"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.uBit-524"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.fBit-525"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDConnectionIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.connectionIDType-526"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.length-527"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.reserved-528"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDCbit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.cbit-529"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeLocalConnectionIDLogicalConnectionId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.localConnectionID.logicalConnectionId-530"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.uBit-531"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.fBit-532"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityDiversityIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.diversityIDType-533"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.length-534"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityIteratingListLocalConnectionID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.localConnectionID-535"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityIteratingListReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.reserved-536"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeDiversityIteratingListDivT() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.diversity.iteratingList.divT-537"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.uBit-538"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.fBit-539"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdContractIDType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.contractIDType-540"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.length-541"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeContractIdContractID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.contractId.contractID-542"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelUBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.uBit-543"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelFBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.fBit-544"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelServiceLevelType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.serviceLevelType-545"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.length-546"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.reserved-547"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyUniSignalingClass229CType1SubTypeUniServiceLevelServiceLevel() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.uniSignalingClass229CType1.subType.uniServiceLevel.serviceLevel-548"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1Lengthbytes() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.lengthbytes-549"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1ClassNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.classNum-550"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.ctype-551"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.type-552"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesResv() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.resv-553"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.srcLSRAddress-554"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength4BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength4Bytes.localId-555"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.type-556"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesResv() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.resv-557"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.srcLSRAddress-558"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength16BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength16Bytes.localId-559"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.type-560"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesResv() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.resv-561"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.srcLSRAddress-562"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength20BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength20Bytes.localId-563"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.type-564"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesResv() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.resv-565"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.srcLSRAddress-566"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLength6BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLength6Bytes.localId-567"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.type-568"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedResv() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.resv-569"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressLength-570"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressValue-571"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType1CallIdentifiersSrcLSRAddressLengthVendorDefinedLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType1.callIdentifiers.srcLSRAddressLengthVendorDefined.localId-572"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2Lengthbytes() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.lengthbytes-573"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2ClassNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.classNum-574"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2Ctype() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.ctype-575"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.type-576"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesIs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.is-577"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesNs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.ns-578"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.srcLSRAddress-579"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength4BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength4Bytes.localId-580"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.type-581"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesIs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.is-582"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesNs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.ns-583"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.srcLSRAddress-584"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength16BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength16Bytes.localId-585"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.type-586"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesIs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.is-587"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesNs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.ns-588"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.srcLSRAddress-589"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength20BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength20Bytes.localId-590"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.type-591"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesIs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.is-592"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesNs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.ns-593"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesSrcLSRAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.srcLSRAddress-594"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLength6BytesLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLength6Bytes.localId-595"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.type-596"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedIs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.is-597"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedNs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.ns-598"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressLength-599"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedSrcLSRAddressAddressValue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.srcLSRAddress.addressValue-600"]]
}
func (s *RsvpStack) ObjectsObjectObjectBodyCallIdentifierObjectclass230CType2CallIdentifiersSrcLSRAddressLengthVendorDefinedLocalId() *TrafficTrafficItemConfigElementStackField {
	return s.Field[rsvpAliasToFieldIdx["rsvpMessege.objects.object.objectBody.callIdentifierObjectclass230CType2.callIdentifiers.srcLSRAddressLengthVendorDefined.localId-601"]]
}

func (s *RsvpStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewRsvpStack(idx int) *RsvpStack {
	stack := RsvpStack(newStack(idx, "rsvp", rsvpAliasToFieldIdx))
	return &stack
}

type EthernetStack TrafficTrafficItemConfigElementStack

var ethernetAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.destinationAddress-1",
	"header.sourceAddress-2",
	"header.etherType-3",
	"header.pfcQueue-4",
})

func (s *EthernetStack) DestinationAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ethernetAliasToFieldIdx["header.destinationAddress-1"]]
}
func (s *EthernetStack) SourceAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ethernetAliasToFieldIdx["header.sourceAddress-2"]]
}
func (s *EthernetStack) EtherType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ethernetAliasToFieldIdx["header.etherType-3"]]
}
func (s *EthernetStack) PfcQueue() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ethernetAliasToFieldIdx["header.pfcQueue-4"]]
}

func (s *EthernetStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewEthernetStack(idx int) *EthernetStack {
	stack := EthernetStack(newStack(idx, "ethernet", ethernetAliasToFieldIdx))
	return &stack
}

type GreStack TrafficTrafficItemConfigElementStack

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

func (s *GreStack) ChecksumPresent() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.checksumPresent-1"]]
}
func (s *GreStack) Reserved1() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.reserved1-2"]]
}
func (s *GreStack) KeyPresent() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.keyPresent-3"]]
}
func (s *GreStack) SequencePresent() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.sequencePresent-4"]]
}
func (s *GreStack) Reserved2() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.reserved2-5"]]
}
func (s *GreStack) Version() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.version-6"]]
}
func (s *GreStack) Protocol() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.protocol-7"]]
}
func (s *GreStack) ChecksumHolderWithChecksumChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.checksumHolder.withChecksum.checksum-8"]]
}
func (s *GreStack) ChecksumHolderWithChecksumReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.checksumHolder.withChecksum.reserved-9"]]
}
func (s *GreStack) ChecksumHolderNoChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.checksumHolder.noChecksum-10"]]
}
func (s *GreStack) KeyHolderKey() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.keyHolder.key-11"]]
}
func (s *GreStack) KeyHolderNoKey() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.keyHolder.noKey-12"]]
}
func (s *GreStack) SequenceHolderSequenceNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.sequenceHolder.sequenceNum-13"]]
}
func (s *GreStack) SequenceHolderNoSequenceNum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[greAliasToFieldIdx["header.sequenceHolder.noSequenceNum-14"]]
}

func (s *GreStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewGreStack(idx int) *GreStack {
	stack := GreStack(newStack(idx, "gre", greAliasToFieldIdx))
	return &stack
}

type MplsStack TrafficTrafficItemConfigElementStack

var mplsAliasToFieldIdx = aliasesToFieldIdx([]string{
	"label.value-1",
	"label.experimental-2",
	"label.bottomOfStack-3",
	"label.ttl-4",
	"label.tracker-5",
})

func (s *MplsStack) Value() *TrafficTrafficItemConfigElementStackField {
	return s.Field[mplsAliasToFieldIdx["label.value-1"]]
}
func (s *MplsStack) Experimental() *TrafficTrafficItemConfigElementStackField {
	return s.Field[mplsAliasToFieldIdx["label.experimental-2"]]
}
func (s *MplsStack) BottomOfStack() *TrafficTrafficItemConfigElementStackField {
	return s.Field[mplsAliasToFieldIdx["label.bottomOfStack-3"]]
}
func (s *MplsStack) Ttl() *TrafficTrafficItemConfigElementStackField {
	return s.Field[mplsAliasToFieldIdx["label.ttl-4"]]
}
func (s *MplsStack) Tracker() *TrafficTrafficItemConfigElementStackField {
	return s.Field[mplsAliasToFieldIdx["label.tracker-5"]]
}

func (s *MplsStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewMplsStack(idx int) *MplsStack {
	stack := MplsStack(newStack(idx, "mpls", mplsAliasToFieldIdx))
	return &stack
}

type Ospfv2DatabaseDescriptionStack TrafficTrafficItemConfigElementStack

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

func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderOspfVersion() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderPacketType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderPacketLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAreaID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2DatabaseDescriptionStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyInterfaceMTU() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.interfaceMTU-15"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyOptions() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.options-16"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.reserved-17"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyDatabaseDescriptionFlags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.databaseDescriptionFlags-18"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyDdSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.ddSequenceNumber-19"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateAge() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateAge-20"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderOptions() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.options-21"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateType-22"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateID-23"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateAdvertisingRouter() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateAdvertisingRouter-24"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateSequenceNumber-25"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateChecksum-26"]]
}
func (s *Ospfv2DatabaseDescriptionStack) DatabaseDescriptionBodyLsaHeadersListLinkStateAdvertisementHeaderVariableHeaderLinkStateLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2DatabaseDescriptionAliasToFieldIdx["header.databaseDescriptionBody.lsaHeadersList.linkStateAdvertisementHeader.variableHeader.linkStateLength-27"]]
}

func (s *Ospfv2DatabaseDescriptionStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewOspfv2DatabaseDescriptionStack(idx int) *Ospfv2DatabaseDescriptionStack {
	stack := Ospfv2DatabaseDescriptionStack(newStack(idx, "ospfv2DatabaseDescription", ospfv2DatabaseDescriptionAliasToFieldIdx))
	return &stack
}

type Ospfv2HelloStack TrafficTrafficItemConfigElementStack

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

func (s *Ospfv2HelloStack) Ospfv2PacketHeaderOspfVersion() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderPacketType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderPacketLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAreaID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2HelloStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2HelloStack) NetworkMask() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.networkMask-15"]]
}
func (s *Ospfv2HelloStack) HelloInterval() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.helloInterval-16"]]
}
func (s *Ospfv2HelloStack) Options() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.options-17"]]
}
func (s *Ospfv2HelloStack) RouterPriority() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.routerPriority-18"]]
}
func (s *Ospfv2HelloStack) RouterDeadInterval() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.routerDeadInterval-19"]]
}
func (s *Ospfv2HelloStack) DesignatedRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.designatedRouterID-20"]]
}
func (s *Ospfv2HelloStack) BackupDesignatedRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.backupDesignatedRouterID-21"]]
}
func (s *Ospfv2HelloStack) HelloNeighborListNeighborRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2HelloAliasToFieldIdx["header.helloNeighborList.neighborRouterID-22"]]
}

func (s *Ospfv2HelloStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewOspfv2HelloStack(idx int) *Ospfv2HelloStack {
	stack := Ospfv2HelloStack(newStack(idx, "ospfv2Hello", ospfv2HelloAliasToFieldIdx))
	return &stack
}

type Ospfv2LSAAcknowledgementStack TrafficTrafficItemConfigElementStack

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

func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderOspfVersion() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderPacketType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderPacketLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAreaID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2LSAAcknowledgementStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateAge() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateAge-15"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderOptions() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.options-16"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateType-17"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateID-18"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateAdvertisingRouter() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateAdvertisingRouter-19"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateSequenceNumber-20"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateChecksum-21"]]
}
func (s *Ospfv2LSAAcknowledgementStack) LinkStateAdvertisementHeaderVariableHeaderLinkStateLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAAcknowledgementAliasToFieldIdx["header.linkStateAdvertisementHeader.variableHeader.linkStateLength-22"]]
}

func (s *Ospfv2LSAAcknowledgementStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewOspfv2LSAAcknowledgementStack(idx int) *Ospfv2LSAAcknowledgementStack {
	stack := Ospfv2LSAAcknowledgementStack(newStack(idx, "ospfv2LSAAcknowledgement", ospfv2LSAAcknowledgementAliasToFieldIdx))
	return &stack
}

type Ospfv2LSARequestStack TrafficTrafficItemConfigElementStack

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

func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderOspfVersion() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderPacketType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderPacketLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAreaID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2LSARequestStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.reserved-15"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateType-16"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateID-17"]]
}
func (s *Ospfv2LSARequestStack) LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateAdvertisingRouter() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSARequestAliasToFieldIdx["header.linkStateRequestBody.requestedLSAsList.requestedLSADescription.linkStateAdvertisingRouter-18"]]
}

func (s *Ospfv2LSARequestStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewOspfv2LSARequestStack(idx int) *Ospfv2LSARequestStack {
	stack := Ospfv2LSARequestStack(newStack(idx, "ospfv2LSARequest", ospfv2LSARequestAliasToFieldIdx))
	return &stack
}

type Ospfv2LSAUpdateStack TrafficTrafficItemConfigElementStack

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

func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderOspfVersion() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.ospfVersion-1"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderPacketType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.packetType-2"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderPacketLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.packetLength-3"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.routerID-4"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAreaID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.areaID-5"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.checksum-6"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationType-7"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataNullAuthentication() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.nullAuthentication-8"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataSimplePassword() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.simplePassword-9"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.reserved-10"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataKeyID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.keyID-11"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataAuthenticationDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.authenticationDataLength-12"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataCryptographicAuthenticationDataCryptographicSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.cryptographicAuthenticationData.cryptographicSequenceNumber-13"]]
}
func (s *Ospfv2LSAUpdateStack) Ospfv2PacketHeaderAuthenticationDataUserDefinedAuthenticationDataUserDefinedAuthData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.ospfv2PacketHeader.authenticationData.userDefinedAuthenticationData.userDefinedAuthData-14"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyNumberOfLSAs() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.numberOfLSAs-15"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAge() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAge-16"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderOptions() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.options-17"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateType-18"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateID-19"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisingRouter() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisingRouter-20"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateSequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateSequenceNumber-21"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateChecksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateChecksum-22"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateLength-23"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSAReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.reserved-24"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterLSAFlags() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerLSAFlags-25"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSAPad() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.pad-26"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSANumberOfRouterInterfaceLinks() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.numberOfRouterInterfaceLinks-27"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceRouterInterfaceLinkID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.routerInterfaceLinkID-28"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceRouterInterfaceLinkData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.routerInterfaceLinkData-29"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceInterfaceType() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.interfaceType-30"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceNumberOfTOSEntries() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.numberOfTOSEntries-31"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceInterfaceMetric() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.interfaceMetric-32"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceTosListTosEntryTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.typeOfService-33"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceTosListTosEntryReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.reserved-34"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyRouterLSARouterInterfacesListRouterInterfaceTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.routerLSA.routerInterfacesList.routerInterface.tosList.tosEntry.metricForCorrespondingTypeOfService-35"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkLSANetworkMask() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkLSA.networkMask-36"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkLSAAttachedRouterListAttachedRouterID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkLSA.attachedRouterList.attachedRouterID-37"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteNetworkMask() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.networkMask-38"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.reserved-39"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteRouteMetric() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.routeMetric-40"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteTosListTosEntryTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.tosList.tosEntry.typeOfService-41"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyNetworkSummaryLSASummaryRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.networkSummaryLSA.summaryRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-42"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteNetworkMask() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.networkMask-43"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.reserved-44"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteRouteMetric() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.routeMetric-45"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteTosListTosEntryTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.tosList.tosEntry.typeOfService-46"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyAsBorderRouterSummaryLSASummaryRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.asBorderRouterSummaryLSA.summaryRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-47"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteNetworkMask() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.networkMask-48"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteEbit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.ebit-49"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.reserved-50"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteRouteMetric() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.routeMetric-51"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteForwardingAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.forwardingAddress-52"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteExternalRouteTag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.externalRouteTag-53"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryEbit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.ebit-54"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.typeOfService-55"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-56"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyExternalLSAExternalRouteTosListTosEntryForwardingAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.externalLSA.externalRoute.tosList.tosEntry.forwardingAddress-57"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteNetworkMask() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.networkMask-58"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteEbit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.ebit-59"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteReserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.reserved-60"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteRouteMetric() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.routeMetric-61"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteForwardingAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.forwardingAddress-62"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteExternalRouteTag() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.externalRouteTag-63"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryEbit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.ebit-64"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.typeOfService-65"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryMetricForCorrespondingTypeOfService() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.metricForCorrespondingTypeOfService-66"]]
}
func (s *Ospfv2LSAUpdateStack) LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisementBodyType7NSSAExternalLSAExternalRouteTosListTosEntryForwardingAddress() *TrafficTrafficItemConfigElementStackField {
	return s.Field[ospfv2LSAUpdateAliasToFieldIdx["header.linkStateUpdateBody.lsaList.linkStateAdvertisement.variableHeader.linkStateAdvertisementBody.type7NSSAExternalLSA.externalRoute.tosList.tosEntry.forwardingAddress-67"]]
}

func (s *Ospfv2LSAUpdateStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewOspfv2LSAUpdateStack(idx int) *Ospfv2LSAUpdateStack {
	stack := Ospfv2LSAUpdateStack(newStack(idx, "ospfv2LSAUpdate", ospfv2LSAUpdateAliasToFieldIdx))
	return &stack
}

type TcpStack TrafficTrafficItemConfigElementStack

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

func (s *TcpStack) SrcPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.srcPort-1"]]
}
func (s *TcpStack) DstPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.dstPort-2"]]
}
func (s *TcpStack) SequenceNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.sequenceNumber-3"]]
}
func (s *TcpStack) AcknowledgementNumber() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.acknowledgementNumber-4"]]
}
func (s *TcpStack) DataOffset() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.dataOffset-5"]]
}
func (s *TcpStack) Reserved() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.reserved-6"]]
}
func (s *TcpStack) EcnNsBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.ecn.nsBit-7"]]
}
func (s *TcpStack) EcnCwrBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.ecn.cwrBit-8"]]
}
func (s *TcpStack) EcnEcnEchoBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.ecn.ecnEchoBit-9"]]
}
func (s *TcpStack) ControlBitsUrgBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.urgBit-10"]]
}
func (s *TcpStack) ControlBitsAckBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.ackBit-11"]]
}
func (s *TcpStack) ControlBitsPshBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.pshBit-12"]]
}
func (s *TcpStack) ControlBitsRstBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.rstBit-13"]]
}
func (s *TcpStack) ControlBitsSynBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.synBit-14"]]
}
func (s *TcpStack) ControlBitsFinBit() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.controlBits.finBit-15"]]
}
func (s *TcpStack) Window() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.window-16"]]
}
func (s *TcpStack) Checksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.checksum-17"]]
}
func (s *TcpStack) UrgentPtr() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.urgentPtr-18"]]
}
func (s *TcpStack) OptionsOptionTypeUserDefinedKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.userDefined.kind-19"]]
}
func (s *TcpStack) OptionsOptionTypeUserDefinedLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.userDefined.length-20"]]
}
func (s *TcpStack) OptionsOptionTypeUserDefinedData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.userDefined.data-21"]]
}
func (s *TcpStack) OptionsOptionTypeEndOfOptionListKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.endOfOptionList.kind-22"]]
}
func (s *TcpStack) OptionsOptionTypeNoOperationKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.noOperation.kind-23"]]
}
func (s *TcpStack) OptionsOptionTypeMaximumSegmentSizeKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.maximumSegmentSize.kind-24"]]
}
func (s *TcpStack) OptionsOptionTypeMaximumSegmentSizeLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.maximumSegmentSize.length-25"]]
}
func (s *TcpStack) OptionsOptionTypeMaximumSegmentSizeData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.maximumSegmentSize.data-26"]]
}
func (s *TcpStack) OptionsOptionTypeWsoptKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.wsopt.kind-27"]]
}
func (s *TcpStack) OptionsOptionTypeWsoptLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.wsopt.length-28"]]
}
func (s *TcpStack) OptionsOptionTypeWsoptData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.wsopt.data-29"]]
}
func (s *TcpStack) OptionsOptionTypeSackPermittedKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sackPermitted.kind-30"]]
}
func (s *TcpStack) OptionsOptionTypeSackPermittedLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sackPermitted.length-31"]]
}
func (s *TcpStack) OptionsOptionTypeSackKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sack.kind-32"]]
}
func (s *TcpStack) OptionsOptionTypeSackLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sack.length-33"]]
}
func (s *TcpStack) OptionsOptionTypeSackData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.sack.data-34"]]
}
func (s *TcpStack) OptionsOptionTypeEchoKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echo.kind-35"]]
}
func (s *TcpStack) OptionsOptionTypeEchoLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echo.length-36"]]
}
func (s *TcpStack) OptionsOptionTypeEchoData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echo.data-37"]]
}
func (s *TcpStack) OptionsOptionTypeEchoReplyKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echoReply.kind-38"]]
}
func (s *TcpStack) OptionsOptionTypeEchoReplyLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echoReply.length-39"]]
}
func (s *TcpStack) OptionsOptionTypeEchoReplyData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.echoReply.data-40"]]
}
func (s *TcpStack) OptionsOptionTypeTsoptKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.tsopt.kind-41"]]
}
func (s *TcpStack) OptionsOptionTypeTsoptLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.tsopt.length-42"]]
}
func (s *TcpStack) OptionsOptionTypeTsoptData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.tsopt.data-43"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderConnectionPermittedKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderConnectionPermitted.kind-44"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderConnectionPermittedLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderConnectionPermitted.length-45"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderServiceProfileKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderServiceProfile.kind-46"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderServiceProfileLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderServiceProfile.length-47"]]
}
func (s *TcpStack) OptionsOptionTypePartialOrderServiceProfileData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.partialOrderServiceProfile.data-48"]]
}
func (s *TcpStack) OptionsOptionTypeCcKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.cc.kind-49"]]
}
func (s *TcpStack) OptionsOptionTypeCcNewKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.ccNew.kind-50"]]
}
func (s *TcpStack) OptionsOptionTypeCcEchoKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.ccEcho.kind-51"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumRequestKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumRequest.kind-52"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumRequestLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumRequest.length-53"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumRequestData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumRequest.data-54"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumDataKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumData.kind-55"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumDataLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumData.length-56"]]
}
func (s *TcpStack) OptionsOptionTypeAlternateChecksumDataData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.alternateChecksumData.data-57"]]
}
func (s *TcpStack) OptionsOptionTypeSkeeterKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.skeeter.kind-58"]]
}
func (s *TcpStack) OptionsOptionTypeBubbaKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.bubba.kind-59"]]
}
func (s *TcpStack) OptionsOptionTypeTrailerChecksumKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.trailerChecksum.kind-60"]]
}
func (s *TcpStack) OptionsOptionTypeTrailerChecksumLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.trailerChecksum.length-61"]]
}
func (s *TcpStack) OptionsOptionTypeTrailerChecksumData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.trailerChecksum.data-62"]]
}
func (s *TcpStack) OptionsOptionTypeMd5SignatureKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.md5Signature.kind-63"]]
}
func (s *TcpStack) OptionsOptionTypeMd5SignatureLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.md5Signature.length-64"]]
}
func (s *TcpStack) OptionsOptionTypeMd5SignatureData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.md5Signature.data-65"]]
}
func (s *TcpStack) OptionsOptionTypeScpsCapabilitiesKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.scpsCapabilities.kind-66"]]
}
func (s *TcpStack) OptionsOptionTypeSelectiveNegativeAckKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.selectiveNegativeAck.kind-67"]]
}
func (s *TcpStack) OptionsOptionTypeRecordBoundariesKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.recordBoundaries.kind-68"]]
}
func (s *TcpStack) OptionsOptionTypeCorruptionExperiencedKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.corruptionExperienced.kind-69"]]
}
func (s *TcpStack) OptionsOptionTypeSnapKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.snap.kind-70"]]
}
func (s *TcpStack) OptionsOptionTypeUnassigned1Kind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.unassigned1.kind-71"]]
}
func (s *TcpStack) OptionsOptionTypeCompressionFilterKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.compressionFilter.kind-72"]]
}
func (s *TcpStack) OptionsOptionTypeQuickStartResponseKind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.quickStartResponse.kind-73"]]
}
func (s *TcpStack) OptionsOptionTypeQuickStartResponseLength() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.quickStartResponse.length-74"]]
}
func (s *TcpStack) OptionsOptionTypeQuickStartResponseData() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.quickStartResponse.data-75"]]
}
func (s *TcpStack) OptionsOptionTypeUnassigned2Kind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.unassigned2.kind-76"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment1Kind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment1.kind-77"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment1Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment1.length-78"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment1Data() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment1.data-79"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment2Kind() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment2.kind-80"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment2Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment2.length-81"]]
}
func (s *TcpStack) OptionsOptionTypeRfc3692StypeExperiment2Data() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.option.type.rfc3692StypeExperiment2.data-82"]]
}
func (s *TcpStack) OptionsPad() *TrafficTrafficItemConfigElementStackField {
	return s.Field[tcpAliasToFieldIdx["header.options.pad-83"]]
}

func (s *TcpStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewTcpStack(idx int) *TcpStack {
	stack := TcpStack(newStack(idx, "tcp", tcpAliasToFieldIdx))
	return &stack
}

type UdpStack TrafficTrafficItemConfigElementStack

var udpAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.srcPort-1",
	"header.dstPort-2",
	"header.length-3",
	"header.checksum-4",
})

func (s *UdpStack) SrcPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[udpAliasToFieldIdx["header.srcPort-1"]]
}
func (s *UdpStack) DstPort() *TrafficTrafficItemConfigElementStackField {
	return s.Field[udpAliasToFieldIdx["header.dstPort-2"]]
}
func (s *UdpStack) Length() *TrafficTrafficItemConfigElementStackField {
	return s.Field[udpAliasToFieldIdx["header.length-3"]]
}
func (s *UdpStack) Checksum() *TrafficTrafficItemConfigElementStackField {
	return s.Field[udpAliasToFieldIdx["header.checksum-4"]]
}

func (s *UdpStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewUdpStack(idx int) *UdpStack {
	stack := UdpStack(newStack(idx, "udp", udpAliasToFieldIdx))
	return &stack
}

type VlanStack TrafficTrafficItemConfigElementStack

var vlanAliasToFieldIdx = aliasesToFieldIdx([]string{
	"header.vlanTag.vlanUserPriority-1",
	"header.vlanTag.cfi-2",
	"header.vlanTag.vlanID-3",
	"header.protocolID-4",
})

func (s *VlanStack) VlanTagVlanUserPriority() *TrafficTrafficItemConfigElementStackField {
	return s.Field[vlanAliasToFieldIdx["header.vlanTag.vlanUserPriority-1"]]
}
func (s *VlanStack) VlanTagCfi() *TrafficTrafficItemConfigElementStackField {
	return s.Field[vlanAliasToFieldIdx["header.vlanTag.cfi-2"]]
}
func (s *VlanStack) VlanTagVlanID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[vlanAliasToFieldIdx["header.vlanTag.vlanID-3"]]
}
func (s *VlanStack) ProtocolID() *TrafficTrafficItemConfigElementStackField {
	return s.Field[vlanAliasToFieldIdx["header.protocolID-4"]]
}

func (s *VlanStack) TrafficTrafficItemConfigElementStack() *TrafficTrafficItemConfigElementStack {
	ts := TrafficTrafficItemConfigElementStack(*s)
	return &ts
}

func NewVlanStack(idx int) *VlanStack {
	stack := VlanStack(newStack(idx, "vlan", vlanAliasToFieldIdx))
	return &stack
}
