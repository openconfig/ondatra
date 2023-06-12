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
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"strings"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	asnSetModeToStr = map[opb.BgpAsnSetMode]string{
		opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE:       "dontincludelocalas",
		opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ:               "includelocalasasasseq",
		opb.BgpAsnSetMode_ASN_SET_MODE_AS_SET:               "includelocalasasasset",
		opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ_CONFEDERATION: "includelocalasasasseqconfederation",
		opb.BgpAsnSetMode_ASN_SET_MODE_AS_SET_CONFEDERATION: "includelocalasasassetconfederation",
		opb.BgpAsnSetMode_ASN_SET_MODE_PREPEND:              "prependlocalastofirstsegment",
	}
)

// Returns the string representing the given cipher suite and the length of keys in bits.
func cipherSuiteStrAndLen(cpb opb.MacSec_CipherSuite) (string, uint32, error) {
	switch cpb {
	case opb.MacSec_CIPHER_SUITE_UNSPECIFIED:
		return "", 0, fmt.Errorf("cipher suite not specified")
	case opb.MacSec_AES_128:
		return "aes128", 128, nil
	case opb.MacSec_AES_256:
		return "aes256", 256, nil
	case opb.MacSec_AES_XPN_128:
		return "aesxpn128", 128, nil
	case opb.MacSec_AES_XPN_256:
		return "aesxpn256", 256, nil
	default:
		return "", 0, fmt.Errorf("unrecognized cipher suite %v", cpb)
	}
}

func confOffsetVal(cpb opb.MacSec_MKA_ConfidentialityOffset) (string, error) {
	switch cpb {
	case opb.MacSec_MKA_OFFSET_UNSPECIFIED:
		return "", fmt.Errorf("confidentiality offset not specified")
	case opb.MacSec_MKA_OFFSET_NO_CONFIDENTIALITY:
		return "noconfidentiality", nil
	case opb.MacSec_MKA_OFFSET_0:
		return "noconfidentialityoffset", nil
	case opb.MacSec_MKA_OFFSET_30:
		return "confidentialityoffset30octets", nil
	case opb.MacSec_MKA_OFFSET_50:
		return "confidentialityoffset50octets", nil
	default:
		return "", fmt.Errorf("unrecognized confidentiality offset %v", cpb)
	}
}

func mka(mkapb *opb.MacSec_MKA) (*ixconfig.TopologyMka, error) {
	cipherSuite, keyLen, err := cipherSuiteStrAndLen(mkapb.GetCipherSuite())
	if err != nil {
		return nil, err
	}

	confOffset, err := confOffsetVal(mkapb.GetConfidentialityOffset())
	if err != nil {
		return nil, err
	}

	var capability string
	switch mkapb.GetCapability() {
	case opb.MacSec_MKA_CAPABILITY_UNSPECIFIED:
		return nil, fmt.Errorf("capability not specified")
	case opb.MacSec_MKA_NOT_IMPLEMENTED:
		capability = "macsecnotimplemented"
	case opb.MacSec_MKA_INTEGRITY_WITHOUT_CONFIDENTIALITY:
		capability = "macsecintegritywithoutconfidentiality"
	case opb.MacSec_MKA_INTEGRITY_WITH_NO_CONFIDENTIALITY_OFFSET:
		capability = "macsecintegritywithnoconfidentialityoffset"
	case opb.MacSec_MKA_INTEGRITY_WITH_CONFIDENTIALITY_OFFSET:
		capability = "macsecintegritywithconfidentialityoffset"
	default:
		return nil, fmt.Errorf("unrecognized capability %v", mkapb.GetCapability())
	}

	cakCache := &ixconfig.TopologyCakCache{}
	if capb := mkapb.GetConnectivityAssociation(); capb != nil {
		checkHex := func(s string, maxLen uint32) error {
			if _, err := hex.DecodeString(s); err != nil {
				return fmt.Errorf("could not parse %q as hex string: %w", s, err)
			}
			if len(s) > int(maxLen) {
				return fmt.Errorf("hex string %q exceeds limit of %d characters", s, maxLen)
			}
			return nil
		}
		if err := checkHex(capb.GetCkn(), 32); err != nil {
			return nil, err
		}
		if err := checkHex(capb.GetCak(), keyLen/4); err != nil {
			return nil, err
		}
		cakCache.CakName = ixconfig.MultivalueStr(capb.GetCkn())
		switch keyLen {
		case 128:
			cakCache.CakValue128 = ixconfig.MultivalueStr(capb.GetCak())
		case 256:
			cakCache.CakValue256 = ixconfig.MultivalueStr(capb.GetCak())
		default:
			return nil, fmt.Errorf("invalid key length %d for CAK value", keyLen)
		}
	}

	return &ixconfig.TopologyMka{
		MkaVersion:            ixconfig.MultivalueUint32(mkapb.GetVersion()),
		CipherSuite:           ixconfig.MultivalueStr(cipherSuite),
		ConfidentialityOffset: ixconfig.MultivalueStr(confOffset),
		MacsecCapability:      ixconfig.MultivalueStr(capability),
		CakCache:              cakCache,
	}, nil
}

func (ix *ixATE) addMACsecProtocol(ifc *opb.InterfaceConfig) error {
	macsec := ifc.GetEthernet().GetMacsec()
	if macsec == nil {
		return nil
	}

	cipherSuite, _, err := cipherSuiteStrAndLen(macsec.GetCipherSuite())
	if err != nil {
		return err
	}

	var rxSAKPool *ixconfig.TopologyRxSakPool
	if rsp := macsec.GetRxSakPool(); rsp != nil {
		rxSAKPool = &ixconfig.TopologyRxSakPool{
			RxSak128: ixconfig.MultivalueStr(rsp.GetSak128()),
			RxSak256: ixconfig.MultivalueStr(rsp.GetSak256()),
		}
	}

	eth := ix.intfs[ifc.GetName()].deviceGroup.Ethernet[0]
	eth.Macsec = append(eth.Macsec, &ixconfig.TopologyMacsec{
		Name:                 ixconfig.String(fmt.Sprintf("MACsec on %s", ifc.GetName())),
		CipherSuite:          ixconfig.MultivalueStr(cipherSuite),
		EnableEncryptedVlans: ixconfig.MultivalueBool(macsec.GetEncryptedVlansEnabled()),
		RxSakPool:            rxSAKPool,
	})
	if mkapb := macsec.GetMka(); mkapb != nil {
		m, err := mka(mkapb)
		if err != nil {
			return err
		}
		eth.Mka = []*ixconfig.TopologyMka{m}
	}
	return nil
}

// addIPProtocols adds IxNetwork IP protocols, assuming the device group for the given interface already exsts.
// Returns an error if the IP configuration does not validate.
func (ix *ixATE) addIPProtocols(ifc *opb.InterfaceConfig) error {
	intf := ix.intfs[ifc.GetName()]
	eth := intf.deviceGroup.Ethernet[0]
	hasMacSec := len(eth.Macsec) > 0
	if ipv4 := ifc.GetIpv4(); ipv4 != nil {
		if ipv4.GetAddressCidr() == "" || ipv4.GetDefaultGateway() == "" {
			return fmt.Errorf("need both address and default gateway defined for IP V4 layer")
		}
		ip, netw, err := net.ParseCIDR(ipv4.GetAddressCidr())
		if err != nil {
			return err
		}
		mask, _ := netw.Mask.Size()
		intf.ipv4 = &ixconfig.TopologyIpv4{
			Name:           ixconfig.String(fmt.Sprintf("IPv4 on %s", ifc.GetName())),
			Address:        ixconfig.MultivalueStr(ip.String()),
			Prefix:         ixconfig.MultivalueUint32(uint32(mask)),
			GatewayIp:      ixconfig.MultivalueStr(ipv4.GetDefaultGateway()),
			ResolveGateway: ixconfig.MultivalueTrue(),
		}
		if hasMacSec {
			eth.Macsec[0].Ipv4 = append(eth.Macsec[0].Ipv4, intf.ipv4)
		} else {
			eth.Ipv4 = append(eth.Ipv4, intf.ipv4)
		}
	}

	if ipv6 := ifc.GetIpv6(); ipv6 != nil {
		if ipv6.GetAddressCidr() == "" || ipv6.GetDefaultGateway() == "" {
			return fmt.Errorf("need both address and default gateway defined for IP V6 layer")
		}
		ip, netw, err := net.ParseCIDR(ipv6.GetAddressCidr())
		if err != nil {
			return err
		}
		mask, _ := netw.Mask.Size()
		intf.ipv6 = &ixconfig.TopologyIpv6{
			Name:           ixconfig.String(fmt.Sprintf("IPv6 on %s", ifc.GetName())),
			Address:        ixconfig.MultivalueStr(ip.String()),
			Prefix:         ixconfig.MultivalueUint32(uint32(mask)),
			GatewayIp:      ixconfig.MultivalueStr(ipv6.GetDefaultGateway()),
			ResolveGateway: ixconfig.MultivalueTrue(),
		}
		if hasMacSec {
			eth.Macsec[0].Ipv6 = append(eth.Macsec[0].Ipv6, intf.ipv6)
		} else {
			eth.Ipv6 = append(eth.Ipv6, intf.ipv6)
		}
	}
	return nil
}

// addIPLoopbackProtocols adds IxNetwork IP loopback protocols, assuming the
// device group for the given interface already exsts.
// Returns an error if the IP configuration does not validate.
func (ix *ixATE) addIPLoopbackProtocols(ifc *opb.InterfaceConfig) error {
	intf := ix.intfs[ifc.GetName()]
	dg := intf.deviceGroup
	if ipv4 := ifc.GetIpv4LoopbackCidr(); ipv4 != "" {
		ip, mask, isV6, err := parseCIDR(ipv4)
		if err != nil || isV6 {
			return fmt.Errorf("invalid IPv4 CIDR address: %q", ipv4)
		}
		dg.Ipv4Loopback = append(dg.Ipv4Loopback, &ixconfig.TopologyIpv4Loopback{
			Name:    ixconfig.String(fmt.Sprintf("IPv4 loopback on %s", ifc.GetName())),
			Address: ixconfig.MultivalueStr(ip),
			Prefix:  ixconfig.MultivalueUint32(mask),
		})
		intf.ipv4Loopback = dg.Ipv4Loopback[0]
	}
	if ipv6 := ifc.GetIpv6LoopbackCidr(); ipv6 != "" {
		ip, mask, isV6, err := parseCIDR(ipv6)
		if err != nil || !isV6 {
			return fmt.Errorf("invalid IPv6 CIDR address: %q", ipv6)
		}
		dg.Ipv6Loopback = append(dg.Ipv6Loopback, &ixconfig.TopologyIpv6Loopback{
			Name:    ixconfig.String(fmt.Sprintf("IPv6 loopback on %s", ifc.GetName())),
			Address: ixconfig.MultivalueStr(ip),
			Prefix:  ixconfig.MultivalueUint32(mask),
		})
		intf.ipv6Loopback = dg.Ipv6Loopback[0]
	}
	return nil
}

// addISISProtocols adds IxNetwork IS-IS protocols, assuming the device group for the given interface already exists.
// Each interface config must apply to the same set of ports. Returns an error if the
// interface configs are not validated as a set.
func (ix *ixATE) addISISProtocols(ifc *opb.InterfaceConfig) error {
	isis := ifc.GetIsis()
	if isis == nil {
		return nil
	}

	var level, networkType string
	var enable3WayHandshake bool
	switch isis.GetLevel() {
	case opb.ISISConfig_LEVEL_UNSPECIFIED:
		return fmt.Errorf("level not specified")
	case opb.ISISConfig_L1:
		level = "level1"
	case opb.ISISConfig_L2:
		level = "level2"
	case opb.ISISConfig_L1L2:
		level = "l1l2"
	default:
		return fmt.Errorf("unrecognized level %s", isis.GetLevel())
	}

	switch isis.GetNetworkType() {
	case opb.ISISConfig_NETWORK_TYPE_UNSPECIFIED:
		return fmt.Errorf("network type not specified")
	case opb.ISISConfig_BROADCAST:
		networkType = "broadcast"
	case opb.ISISConfig_POINT_TO_POINT:
		networkType = "pointpoint"
		enable3WayHandshake = true
	default:
		return fmt.Errorf("unrecognized network type %s", isis.GetNetworkType())
	}

	authType, err := authTypeString(isis.GetAuthType())
	if err != nil {
		return err
	}
	areaAuthType, err := authTypeString(isis.GetAreaAuthType())
	if err != nil {
		return err
	}
	domainAuthType, err := authTypeString(isis.GetDomainAuthType())
	if err != nil {
		return err
	}

	areaID, err := areaIDToIxHex(isis.GetAreaId())
	if err != nil {
		return err
	}
	isisIntf := &ixconfig.TopologyIsisL3{
		Name:                           ixconfig.String(fmt.Sprintf("IS-IS on %s", ifc.GetName())),
		LevelType:                      ixconfig.MultivalueStr(level),
		NetworkType:                    ixconfig.MultivalueStr(networkType),
		Enable3WayHandshake:            ixconfig.MultivalueBool(enable3WayHandshake),
		AuthType:                       ixconfig.MultivalueStr(authType),
		CircuitTranmitPasswordOrMD5Key: ixconfig.MultivalueStr(isis.GetAuthKey()),
		InterfaceMetric:                ixconfig.MultivalueUint32(isis.GetMetric()),
		Level2Priority:                 ixconfig.MultivalueUint32(isis.GetInterfacePriority()),
		Level2HelloInterval:            ixconfig.MultivalueUint32(isis.GetHelloIntervalSec()),
		Level2DeadInterval:             ixconfig.MultivalueUint32(isis.GetDeadIntervalSec()),
	}
	isisRtr := &ixconfig.TopologyIsisL3Router{
		Name:                           ixconfig.String(fmt.Sprintf("IS-IS Router on %s", ifc.GetName())),
		EnableHelloPadding:             ixconfig.MultivalueBool(isis.GetEnableHelloPadding()),
		EnableWideMetric:               ixconfig.MultivalueBool(isis.GetEnableWideMetric()),
		EnableTE:                       ixconfig.MultivalueBool(isis.GetEnableTe()),
		DiscardLSPs:                    ixconfig.MultivalueBool(isis.GetDiscardLsps()),
		AreaAddresses:                  ixconfig.MultivalueStr(areaID),
		TERouterId:                     ixconfig.MultivalueStr(isis.GetTeRouterId()),
		AreaAuthenticationType:         ixconfig.MultivalueStr(areaAuthType),
		AreaTransmitPasswordOrMD5Key:   ixconfig.MultivalueStr(isis.GetAreaAuthKey()),
		DomainAuthenticationType:       ixconfig.MultivalueStr(domainAuthType),
		DomainTransmitPasswordOrMD5Key: ixconfig.MultivalueStr(isis.GetDomainAuthKey()),
	}
	if isis.GetCapabilityRouterId() != "" {
		isisRtr.RtrcapId = ixconfig.MultivalueStr(isis.GetCapabilityRouterId())
	}

	if err = isisSegmentRouting(isisIntf, isisRtr, isis.GetSegmentRouting()); err != nil {
		return err
	}
	isisNetwGrps, err := isisReachability(ifc.GetName(), isis.GetIsReachabilities())
	if err != nil {
		return err
	}
	ix.intfs[ifc.GetName()].isrToNetworkGroup = isisNetwGrps

	dg := ix.intfs[ifc.GetName()].deviceGroup
	dg.Ethernet[0].IsisL3 = append(dg.Ethernet[0].IsisL3, isisIntf)
	dg.IsisL3Router = append(dg.IsisL3Router, isisRtr)
	for _, ng := range isisNetwGrps {
		dg.NetworkGroup = append(dg.NetworkGroup, ng)
	}
	return nil
}

func authTypeString(authType opb.ISISConfig_AuthType) (string, error) {
	switch authType {
	case opb.ISISConfig_AUTH_TYPE_UNSPECIFIED:
		return "none", nil
	case opb.ISISConfig_MD5:
		return "md5", nil
	case opb.ISISConfig_PASSWORD:
		return "password", nil
	default:
		return "", fmt.Errorf("unrecognized auth type %s", authType)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// isisSegmentRouting updates isisIntf/isisRtr based on the given segment routing config.
func isisSegmentRouting(isisIntf *ixconfig.TopologyIsisL3, isisRtr *ixconfig.TopologyIsisL3Router, sr *opb.ISISSegmentRouting) error {
	if sr == nil || sr.GetEnable() == false {
		isisRtr.EnableSR = ixconfig.Bool(false)
		isisIntf.EnableAdjSID = ixconfig.MultivalueFalse()
		return nil
	}

	// Configure IS-IS router
	isisRtr.EnableSR = ixconfig.Bool(true)
	isisRtr.SIDIndexLabel = ixconfig.MultivalueUint32(sr.GetSidIndexLabel())
	isisRtr.RFlag = ixconfig.MultivalueBool(sr.GetFlagReadvertise())
	isisRtr.NFlag = ixconfig.MultivalueBool(sr.GetFlagNodeSid())
	isisRtr.PFlag = ixconfig.MultivalueBool(sr.GetFlagNoPhp())
	isisRtr.EFlag = ixconfig.MultivalueBool(sr.GetFlagExplicitNull())
	isisRtr.VFlag = ixconfig.MultivalueBool(sr.GetFlagValue())
	isisRtr.LFlag = ixconfig.MultivalueBool(sr.GetFlagLocal())
	isisRtr.SRAlgorithmCount = ixconfig.NumberInt(max(1, len(sr.GetAlgorithms())))
	for _, alg := range sr.GetAlgorithms() {
		isisRtr.IsisSRAlgorithmList = append(isisRtr.IsisSRAlgorithmList, &ixconfig.TopologyIsisSrAlgorithmList{
			IsisSrAlgorithm: ixconfig.MultivalueUint32(alg),
		})
	}

	if sr.GetPrefixSid() != "" {
		nodePrefix, nodeMask, isV6, err := parseCIDR(sr.GetPrefixSid())
		if err != nil || isV6 {
			return fmt.Errorf("invalid prefix SID: %q, want an IPv4 address in CIDR format", sr.GetPrefixSid())
		}
		isisRtr.NodePrefix = ixconfig.MultivalueStr(nodePrefix)
		isisRtr.Mask = ixconfig.MultivalueUint32(nodeMask)
	}

	var lbCounts, lbLabels []uint32
	for _, srlb := range sr.GetSrlbRanges() {
		lbCounts = append(lbCounts, srlb.GetSidCount())
		lbLabels = append(lbLabels, srlb.GetSidStartLabel())
	}
	isisRtr.SrlbDescriptorCount = ixconfig.NumberInt(max(1, len(lbCounts)))
	for i := range lbCounts {
		isisRtr.AdvertiseSRLB = ixconfig.MultivalueTrue()
		isisRtr.IsisSRLBDescriptorList = append(isisRtr.IsisSRLBDescriptorList, &ixconfig.TopologyIsisSrlbDescriptorList{
			SIDCount:      ixconfig.MultivalueUint32(lbCounts[i]),
			StartSIDLabel: ixconfig.MultivalueUint32(lbLabels[i]),
		})
	}

	var gbCounts, gbLabels []uint32
	for _, srgb := range sr.GetSrgbRanges() {
		gbCounts = append(gbCounts, srgb.GetSidCount())
		gbLabels = append(gbLabels, srgb.GetSidStartLabel())
	}
	isisRtr.SRGBRangeCount = ixconfig.NumberInt(len(gbCounts))
	for i := range gbCounts {
		isisRtr.IsisSRGBRangeSubObjectsList = append(isisRtr.IsisSRGBRangeSubObjectsList, &ixconfig.TopologyIsisSrgbRangeSubObjectsList{
			SIDCount:      ixconfig.MultivalueUint32(gbCounts[i]),
			StartSIDLabel: ixconfig.MultivalueUint32(gbLabels[i]),
		})
	}

	// Configure IS-IS interface
	as := sr.GetAdjacencySid()
	if as == nil {
		isisIntf.EnableAdjSID = ixconfig.MultivalueFalse()
		return nil
	}
	isisIntf.EnableAdjSID = ixconfig.MultivalueTrue()
	isisIntf.AdjSID = ixconfig.MultivalueStr(as.GetSid())
	isisIntf.OverrideFFlag = ixconfig.MultivalueBool(as.GetOverrideFlagAddressFamily())
	isisIntf.FFlag = ixconfig.MultivalueBool(as.GetFlagAddressFamily())
	isisIntf.BFlag = ixconfig.MultivalueBool(as.GetFlagBackup())
	isisIntf.VFlag = ixconfig.MultivalueBool(as.GetFlagValue())
	isisIntf.LFlag = ixconfig.MultivalueBool(as.GetFlagLocal())
	isisIntf.SFlag = ixconfig.MultivalueBool(as.GetFlagSet())
	isisIntf.PFlag = ixconfig.MultivalueBool(as.GetFlagPersistent())
	return nil
}

func parseCIDR(cidr string) (string, uint32, bool, error) {
	_, netw, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", 0, false, fmt.Errorf("could not parse %q as CIDR: %w", cidr, err)
	}
	maskLen, _ := netw.Mask.Size()
	return strings.SplitN(cidr, "/", 2)[0], uint32(maskLen), strings.Contains(cidr, ":"), nil
}

func toIxHex(s string, lenBytes int) (string, error) {
	b, err := hex.DecodeString(strings.NewReplacer(":", "", ".", "", " ", "").Replace(s))
	if err != nil {
		return "", fmt.Errorf("could not decode %q as hex: %w", s, err)
	}
	if len(b) > lenBytes {
		return "", fmt.Errorf("%q is too large (max size is %d bytes)", s, lenBytes)
	}

	words := make([]string, lenBytes)
	fillLen := lenBytes - len(b)
	for i := 0; i < len(b); i++ {
		words[fillLen+i] = fmt.Sprintf("%02x", b[i])
	}
	for i := 0; i < fillLen; i++ {
		words[i] = "00"
	}
	return strings.Join(words, " "), nil
}

func areaIDToIxHex(id string) (string, error) {
	return toIxHex(id, 3)
}

func sysIDToIxHex(id string) (string, error) {
	return toIxHex(id, 6)
}

func parseLink(toCIDR, fromCIDR string, asV6 bool) (string, string, uint32, error) {
	if (toCIDR == "") != (fromCIDR == "") {
		return "", "", 0, fmt.Errorf("either both or neither IS-IS node link to/from addresses must be set (to: %q, from: %q)", toCIDR, fromCIDR)
	}

	vers := "IPv4"
	if asV6 {
		vers = "IPv6"
	}

	toIP, toIPMask, isV6, err := parseCIDR(toCIDR)
	if err != nil || isV6 != asV6 {
		return "", "", 0, fmt.Errorf("error parsing IS-IS node link 'to' address %q as %s: %w", toCIDR, vers, err)
	}

	fromIP, fromIPMask, isV6, err := parseCIDR(fromCIDR)
	if err != nil || isV6 != asV6 {
		return "", "", 0, fmt.Errorf("error parsing IS-IS node link 'from' address %q as %s: %w", fromCIDR, vers, err)
	}

	if toIPMask != fromIPMask {
		return "", "", 0, fmt.Errorf("unequal masks for 'to' address %q and 'from' address %q for IS-IS node link", toCIDR, fromCIDR)
	}
	return toIP, fromIP, toIPMask, nil
}

// isisReachability returns configured TopologyNetworkGroups based on the given reachability configuration.
func isisReachability(ifcName string, isrs []*opb.ISReachability) (map[string]*ixconfig.TopologyNetworkGroup, error) {
	isisNetwGrps := make(map[string]*ixconfig.TopologyNetworkGroup)
	for i, cfg := range isrs {
		var maxLinks, maxAlgos, numSRGBRanges, numSRLBRanges int
		for j, node := range cfg.GetNodes() {
			if numLinks := len(node.GetLinks()); numLinks > maxLinks {
				maxLinks = numLinks
			}
			if sr := node.GetSegmentRouting(); sr != nil {
				if numAlgos := len(sr.GetAlgorithms()); numAlgos > maxAlgos {
					maxAlgos = numAlgos
				}
				if j == 0 {
					numSRGBRanges = len(sr.GetSrgbRanges())
					numSRLBRanges = len(sr.GetSrlbRanges())
				} else if numSRGBRanges != len(sr.GetSrgbRanges()) {
					return nil, fmt.Errorf("all IS-IS nodes in reachability config need the same number of SRGB ranges")
				} else if numSRLBRanges != len(sr.GetSrlbRanges()) {
					return nil, fmt.Errorf("all IS-IS nodes in reachability config need the same number of SRLB ranges")
				}
			} else if numSRGBRanges != 0 {
				return nil, fmt.Errorf("all IS-IS nodes in reachability config need the same number of SRGB ranges")
			} else if numSRLBRanges != 0 {
				return nil, fmt.Errorf("all IS-IS nodes in reachability config need the same number of SRLB ranges")
			}
		}
		var algoLists []*ixconfig.TopologyIsisSrAlgorithmList
		for j := 0; j < maxAlgos; j++ {
			algoLists = append(algoLists, &ixconfig.TopologyIsisSrAlgorithmList{})
		}
		var srgbRangeLists []*ixconfig.TopologyIsisSrgbRangeSubObjectsList
		for j := 0; j < numSRGBRanges; j++ {
			srgbRangeLists = append(srgbRangeLists, &ixconfig.TopologyIsisSrgbRangeSubObjectsList{})
		}
		var srlbDescriptorLists []*ixconfig.TopologyIsisSrlbDescriptorList
		for j := 0; j < numSRLBRanges; j++ {
			srlbDescriptorLists = append(srlbDescriptorLists, &ixconfig.TopologyIsisSrlbDescriptorList{})
		}
		ipv4Routes := &ixconfig.TopologyIPv4PseudoNodeRoutes{}
		ipv6Routes := &ixconfig.TopologyIPv6PseudoNodeRoutes{}
		isisRtr := &ixconfig.TopologyIsisL3PseudoRouter{
			// Even if we're not configuring anything, IxNetwork requires the counts to be at least 1.
			SRAlgorithmCount:            ixconfig.NumberInt(max(maxAlgos, 1)),
			IsisSRAlgorithmList:         algoLists,
			SRGBRangeCount:              ixconfig.NumberInt(max(numSRGBRanges, 1)),
			IsisSRGBRangeSubObjectsList: srgbRangeLists,
			SrlbDescriptorCount:         ixconfig.NumberInt(max(numSRLBRanges, 1)),
			IsisSRLBDescriptorList:      srlbDescriptorLists,
			IPv4PseudoNodeRoutes:        []*ixconfig.TopologyIPv4PseudoNodeRoutes{ipv4Routes},
			IPv6PseudoNodeRoutes:        []*ixconfig.TopologyIPv6PseudoNodeRoutes{ipv6Routes},
		}
		simRtr := &ixconfig.TopologySimRouter{IsisL3PseudoRouter: []*ixconfig.TopologyIsisL3PseudoRouter{isisRtr}}
		simIntfIPv4 := &ixconfig.TopologySimInterfaceIPv4Config{}
		simIntfIPv6 := &ixconfig.TopologySimInterfaceIPv6Config{}
		attPoint1 := &ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{}
		attPoint2 := &ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{}
		isisIntf := &ixconfig.TopologyIsisL3PseudoInterface{
			NoOfTeProfile:                    ixconfig.NumberUint32(1),
			IsisL3PseudoIfaceAttPoint1Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{attPoint1},
			IsisL3PseudoIfaceAttPoint2Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{attPoint2},
		}
		for _, node := range cfg.GetNodes() {
			sysID, err := sysIDToIxHex(node.GetSystemId())
			if err != nil {
				return nil, err
			}
			simRtr.SystemId = appendStrToMultivalueList(simRtr.SystemId, sysID)

			isisRtr.Enable = appendBoolToMultivalueList(isisRtr.Enable, node.GetEnableTe())
			isisRtr.EnableIpV6TE = appendBoolToMultivalueList(isisRtr.EnableIpV6TE, node.GetEnableTe())
			isisRtr.EnableWideMetric = appendBoolToMultivalueList(isisRtr.EnableWideMetric, node.GetEnableWideMetric())
			isisRtr.EnableWMforTEinNetworkGroup = appendBoolToMultivalueList(isisRtr.EnableWMforTEinNetworkGroup, node.GetEnableWideMetric())
			capRtrID := "0.0.0.0"
			if id := node.GetCapabilityRouterId(); id != "" {
				ip, isV6 := parseIP(id)
				if ip == nil || isV6 {
					return nil, fmt.Errorf("invalid capability router ID: %q", id)
				}
				capRtrID = ip.String()
			}
			isisRtr.RtrcapId = appendStrToMultivalueList(isisRtr.RtrcapId, capRtrID)
			teRtrID := "0.0.0.0"
			if id := node.GetTeRouterId(); id != "" {
				ip, isV6 := parseIP(id)
				if ip == nil || isV6 {
					return nil, fmt.Errorf("invalid TE router ID: %q", id)
				}
				teRtrID = ip.String()
			}
			isisRtr.TERouterId = appendStrToMultivalueList(isisRtr.TERouterId, teRtrID)

			// Node links.
			links := node.GetLinks()
			for k := 0; k < maxLinks; k++ {
				toIPv4, fromIPv4, toIPv6, fromIPv6 := "0.0.0.0", "0.0.0.0", "::", "::"
				maskV4, maskV6 := uint32(1), uint32(1) // IxNetwork will reject '0' as a value even if the address is not enabled.
				var enableIPv4, enableIPv6 bool
				if k < len(links) {
					if to, from := links[k].GetToIpv4(), links[k].GetFromIpv4(); to != "" || from != "" {
						enableIPv4 = true
						toIPv4, fromIPv4, maskV4, err = parseLink(to, from, false)
						if err != nil {
							return nil, err
						}
					}
					if to, from := links[k].GetToIpv6(), links[k].GetFromIpv6(); to != "" || from != "" {
						enableIPv6 = true
						toIPv6, fromIPv6, maskV6, err = parseLink(to, from, true)
						if err != nil {
							return nil, err
						}
					}
				}

				simIntfIPv4.EnableIp = appendBoolToMultivalueList(simIntfIPv4.EnableIp, enableIPv4)
				simIntfIPv4.ToIP = appendStrToMultivalueList(simIntfIPv4.ToIP, toIPv4)
				simIntfIPv4.FromIP = appendStrToMultivalueList(simIntfIPv4.FromIP, fromIPv4)
				simIntfIPv4.SubnetPrefixLength = appendUintToMultivalueList(simIntfIPv4.SubnetPrefixLength, maskV4)

				simIntfIPv6.EnableIp = appendBoolToMultivalueList(simIntfIPv6.EnableIp, enableIPv6)
				simIntfIPv6.ToIP = appendStrToMultivalueList(simIntfIPv6.ToIP, toIPv6)
				simIntfIPv6.FromIP = appendStrToMultivalueList(simIntfIPv6.FromIP, fromIPv6)
				simIntfIPv6.SubnetPrefixLength = appendUintToMultivalueList(simIntfIPv6.SubnetPrefixLength, maskV6)

				inMetric := node.GetIngressMetric()
				if inMetric < 1 || inMetric > 63 {
					return nil, fmt.Errorf("invalid value %d for node ingress metric", inMetric)
				}
				egMetric := node.GetEgressMetric()
				if egMetric < 1 || egMetric > 63 {
					return nil, fmt.Errorf("invalid value %d for node egress metric", egMetric)
				}
				attPoint1.LinkMetric = appendUintToMultivalueList(attPoint1.LinkMetric, inMetric)
				attPoint2.LinkMetric = appendUintToMultivalueList(attPoint2.LinkMetric, egMetric)
			}

			// Route export
			if routes := node.GetRoutesIpv4(); routes != nil {
				if err = isisIPv4RouteExport(routes, ipv4Routes); err != nil {
					return nil, err
				}
			}
			if routes := node.GetRoutesIpv6(); routes != nil {
				if err = isisIPv6RouteExport(routes, ipv6Routes); err != nil {
					return nil, err
				}
			}

			// Segment routing.
			var enableSR, srRFlag, srNFlag, srPFlag, srEFlag, srVFlag, srLFlag bool
			var srSIDIdxLbl uint32
			var srgbRanges, srlbRanges []*opb.ISISSegmentRouting_SIDRange
			var algos []uint32
			nodePrefix, nodeMask := "0.0.0.0", uint32(1)
			sr := node.GetSegmentRouting()
			if sr != nil {
				enableSR = sr.GetEnable()
				srSIDIdxLbl = sr.GetSidIndexLabel()
				srRFlag = sr.GetFlagReadvertise()
				srNFlag = sr.GetFlagNodeSid()
				srPFlag = sr.GetFlagNoPhp()
				srEFlag = sr.GetFlagExplicitNull()
				srVFlag = sr.GetFlagValue()
				srLFlag = sr.GetFlagLocal()
				srgbRanges = sr.GetSrgbRanges()
				srlbRanges = sr.GetSrlbRanges()
				algos = sr.GetAlgorithms()

				if sr.GetPrefixSid() != "" {
					var isV6 bool
					nodePrefix, nodeMask, isV6, err = parseCIDR(sr.GetPrefixSid())
					if err != nil || isV6 {
						return nil, fmt.Errorf("invalid prefix SID: %q", sr.GetPrefixSid())
					}
				}
			}
			if enableSR {
				// Enable SR if any node requires it
				isisRtr.EnableSR = ixconfig.Bool(true)
			}
			isisRtr.SIDIndexLabel = appendUintToMultivalueList(isisRtr.SIDIndexLabel, srSIDIdxLbl)
			isisRtr.RFlag = appendBoolToMultivalueList(isisRtr.RFlag, srRFlag)
			isisRtr.NFlag = appendBoolToMultivalueList(isisRtr.NFlag, srNFlag)
			isisRtr.PFlag = appendBoolToMultivalueList(isisRtr.PFlag, srPFlag)
			isisRtr.EFlag = appendBoolToMultivalueList(isisRtr.EFlag, srEFlag)
			isisRtr.VFlag = appendBoolToMultivalueList(isisRtr.VFlag, srVFlag)
			isisRtr.LFlag = appendBoolToMultivalueList(isisRtr.LFlag, srLFlag)
			for k := 0; k < numSRGBRanges; k++ {
				srgbRangeLists[k].StartSIDLabel = appendUintToMultivalueList(srgbRangeLists[k].StartSIDLabel, srgbRanges[k].GetSidStartLabel())
				srgbRangeLists[k].SIDCount = appendUintToMultivalueList(srgbRangeLists[k].SIDCount, srgbRanges[k].GetSidCount())
			}
			for k := 0; k < numSRLBRanges; k++ {
				srlbDescriptorLists[k].StartSIDLabel = appendUintToMultivalueList(srlbDescriptorLists[k].StartSIDLabel, srlbRanges[k].GetSidStartLabel())
				srlbDescriptorLists[k].SIDCount = appendUintToMultivalueList(srlbDescriptorLists[k].SIDCount, srlbRanges[k].GetSidCount())
			}
			for k := 0; k < maxAlgos; k++ {
				var algo uint32
				if k < len(algos) {
					algo = algos[k]
				}
				algoLists[k].IsisSrAlgorithm = appendUintToMultivalueList(algoLists[k].IsisSrAlgorithm, algo)
			}
			isisRtr.NodePrefix = appendStrToMultivalueList(isisRtr.NodePrefix, nodePrefix)
			isisRtr.Mask = appendUintToMultivalueList(isisRtr.Mask, nodeMask)

			var adj *opb.ISISSegmentRouting_AdjacencySID
			if sr != nil {
				adj = sr.GetAdjacencySid()
			}
			var enableAdjSID, adjOverrideFFlag, adjFFlag, adjBFlag, adjVFlag, adjLFlag, adjSFlag, adjPFlag bool
			adjSID := uint32(1) // IxNetwork will reject a value of '0' even if adjacencies are disabled.
			if adj != nil {
				enableAdjSID = true
				adjSIDNum, err := strconv.ParseUint(adj.GetSid(), 10, 32)
				if err != nil {
					return nil, fmt.Errorf("adjacency SID value %q is not a number", adj.GetSid())
				}
				adjSID = uint32(adjSIDNum)
				adjOverrideFFlag = adj.GetOverrideFlagAddressFamily()
				adjFFlag = adj.GetFlagAddressFamily()
				adjBFlag = adj.GetFlagBackup()
				adjVFlag = adj.GetFlagValue()
				adjLFlag = adj.GetFlagLocal()
				adjSFlag = adj.GetFlagSet()
				adjPFlag = adj.GetFlagPersistent()
			}
			for k := 0; k < maxLinks; k++ {
				isisIntf.EnableAdjSID = appendBoolToMultivalueList(isisIntf.EnableAdjSID, enableAdjSID)
				isisIntf.AdjSID = appendUintToMultivalueList(isisIntf.AdjSID, adjSID)
				isisIntf.OverrideFFlag = appendBoolToMultivalueList(isisIntf.OverrideFFlag, adjOverrideFFlag)
				isisIntf.FFlag = appendBoolToMultivalueList(isisIntf.FFlag, adjFFlag)
				isisIntf.BFlag = appendBoolToMultivalueList(isisIntf.BFlag, adjBFlag)
				isisIntf.VFlag = appendBoolToMultivalueList(isisIntf.VFlag, adjVFlag)
				isisIntf.LFlag = appendBoolToMultivalueList(isisIntf.LFlag, adjLFlag)
				isisIntf.SFlag = appendBoolToMultivalueList(isisIntf.SFlag, adjSFlag)
				isisIntf.PFlag = appendBoolToMultivalueList(isisIntf.PFlag, adjPFlag)
			}
		}

		name := cfg.GetName()
		if name == "" {
			name = fmt.Sprintf("IS-IS Topology %d for %s", i, ifcName)
		}
		if isisNetwGrps[name] != nil {
			return nil, fmt.Errorf("name %q reused for multiple IS-IS reachability groups on interface %q", name, ifcName)
		}
		isisNetwGrps[name] = &ixconfig.TopologyNetworkGroup{
			Name: ixconfig.String(name),
			NetworkTopology: &ixconfig.TopologyNetworkTopology{
				NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
					Nodes:          ixconfig.NumberInt(len(cfg.GetNodes())),
					LinkMultiplier: ixconfig.NumberInt(maxLinks),
				},
				SimRouter: []*ixconfig.TopologySimRouter{simRtr},
				SimInterface: []*ixconfig.TopologySimInterface{{
					IsisL3PseudoInterface:  []*ixconfig.TopologyIsisL3PseudoInterface{isisIntf},
					SimInterfaceIPv4Config: []*ixconfig.TopologySimInterfaceIPv4Config{simIntfIPv4},
					SimInterfaceIPv6Config: []*ixconfig.TopologySimInterfaceIPv6Config{simIntfIPv6},
				}},
			},
		}
	}
	return isisNetwGrps, nil
}

func isisIPv4RouteExport(routes *opb.ISReachability_Node_Routes, ipv4Routes *ixconfig.TopologyIPv4PseudoNodeRoutes) error {
	var exportRoutes, routeConfigureSID, routeRFlag, routeNFlag, routePFlag, routeEFlag, routeVFlag, routeLFlag bool
	var routeMetric, routeAlgo, routeSIDIndexLabel uint32
	routeOrigin := "internal"
	exportRoutes = true
	routeAddr, routeMask, isV6, err := parseCIDR(routes.GetPrefix())
	if err != nil || isV6 {
		return fmt.Errorf("invalid value %q for IPv4 route prefix", routes.GetPrefix())
	}
	routeCount := routes.GetNumRoutes()
	if ipr := routes.GetReachability(); ipr != nil {
		routeOrigin, err = routeOriginStr(ipr.GetRouteOrigin())
		if err != nil {
			return err
		}
		exportRoutes = ipr.GetActive()
		routeMetric = ipr.GetMetric()
		routeAlgo = ipr.GetAlgorithm()
		routeConfigureSID = ipr.GetEnableSidIndexLabel()
		routeSIDIndexLabel = ipr.GetSidIndexLabel()
		routeRFlag = ipr.GetFlagReadvertise()
		routeNFlag = ipr.GetFlagNodeSid()
		routePFlag = ipr.GetFlagNoPhp()
		routeEFlag = ipr.GetFlagExplicitNull()
		routeVFlag = ipr.GetFlagValue()
		routeLFlag = ipr.GetFlagLocal()
	}
	ipv4Routes.Active = appendBoolToMultivalueList(ipv4Routes.Active, exportRoutes)
	ipv4Routes.NetworkAddress = appendStrToMultivalueList(ipv4Routes.NetworkAddress, routeAddr)
	ipv4Routes.PrefixLength = appendUintToMultivalueList(ipv4Routes.PrefixLength, routeMask)
	ipv4Routes.RangeSize = appendUint64ToMultivalueList(ipv4Routes.RangeSize, routeCount)
	ipv4Routes.Ipv4Metric = appendUintToMultivalueList(ipv4Routes.Ipv4Metric, routeMetric)
	ipv4Routes.Algorithm = appendUintToMultivalueList(ipv4Routes.Algorithm, routeAlgo)
	ipv4Routes.Ipv4RouteOrigin = appendStrToMultivalueList(ipv4Routes.Ipv4RouteOrigin, routeOrigin)
	ipv4Routes.ConfigureSIDIndexLabel = appendBoolToMultivalueList(ipv4Routes.ConfigureSIDIndexLabel, routeConfigureSID)
	ipv4Routes.SIDIndexLabel = appendUintToMultivalueList(ipv4Routes.SIDIndexLabel, routeSIDIndexLabel)
	ipv4Routes.Ipv4RFlag = appendBoolToMultivalueList(ipv4Routes.Ipv4RFlag, routeRFlag)
	ipv4Routes.Ipv4NFlag = appendBoolToMultivalueList(ipv4Routes.Ipv4NFlag, routeNFlag)
	ipv4Routes.Ipv4PFlag = appendBoolToMultivalueList(ipv4Routes.Ipv4PFlag, routePFlag)
	ipv4Routes.Ipv4EFlag = appendBoolToMultivalueList(ipv4Routes.Ipv4EFlag, routeEFlag)
	ipv4Routes.Ipv4VFlag = appendBoolToMultivalueList(ipv4Routes.Ipv4VFlag, routeVFlag)
	ipv4Routes.Ipv4LFlag = appendBoolToMultivalueList(ipv4Routes.Ipv4LFlag, routeLFlag)
	return nil
}

func isisIPv6RouteExport(routes *opb.ISReachability_Node_Routes, ipv6Routes *ixconfig.TopologyIPv6PseudoNodeRoutes) error {
	var exportRoutes, routeConfigureSID, routeRFlag, routeNFlag, routePFlag, routeEFlag, routeVFlag, routeLFlag bool
	var routeMetric, routeAlgo, routeSIDIndexLabel uint32
	routeOrigin := "internal"
	exportRoutes = true
	routeAddr, routeMask, isV6, err := parseCIDR(routes.GetPrefix())
	if err != nil || !isV6 {
		return fmt.Errorf("invalid value %q for IPv6 route prefix", routes.GetPrefix())
	}
	routeCount := routes.GetNumRoutes()
	if ipr := routes.GetReachability(); ipr != nil {
		routeOrigin, err = routeOriginStr(ipr.GetRouteOrigin())
		if err != nil {
			return err
		}
		exportRoutes = ipr.GetActive()
		routeMetric = ipr.GetMetric()
		routeAlgo = ipr.GetAlgorithm()
		routeConfigureSID = ipr.GetEnableSidIndexLabel()
		routeSIDIndexLabel = ipr.GetSidIndexLabel()
		routeRFlag = ipr.GetFlagReadvertise()
		routeNFlag = ipr.GetFlagNodeSid()
		routePFlag = ipr.GetFlagNoPhp()
		routeEFlag = ipr.GetFlagExplicitNull()
		routeVFlag = ipr.GetFlagValue()
		routeLFlag = ipr.GetFlagLocal()
	}
	ipv6Routes.Active = appendBoolToMultivalueList(ipv6Routes.Active, exportRoutes)
	ipv6Routes.NetworkAddress = appendStrToMultivalueList(ipv6Routes.NetworkAddress, routeAddr)
	ipv6Routes.Prefix = appendUintToMultivalueList(ipv6Routes.Prefix, routeMask)
	ipv6Routes.RangeSize = appendUint64ToMultivalueList(ipv6Routes.RangeSize, routeCount)
	ipv6Routes.Ipv6Metric = appendUintToMultivalueList(ipv6Routes.Ipv6Metric, routeMetric)
	ipv6Routes.Algorithm = appendUintToMultivalueList(ipv6Routes.Algorithm, routeAlgo)
	ipv6Routes.Ipv6RouteOrigin = appendStrToMultivalueList(ipv6Routes.Ipv6RouteOrigin, routeOrigin)
	ipv6Routes.ConfigureSIDIndexLabel = appendBoolToMultivalueList(ipv6Routes.ConfigureSIDIndexLabel, routeConfigureSID)
	ipv6Routes.SIDIndexLabel = appendUintToMultivalueList(ipv6Routes.SIDIndexLabel, routeSIDIndexLabel)
	ipv6Routes.Ipv6RFlag = appendBoolToMultivalueList(ipv6Routes.Ipv6RFlag, routeRFlag)
	ipv6Routes.Ipv6NFlag = appendBoolToMultivalueList(ipv6Routes.Ipv6NFlag, routeNFlag)
	ipv6Routes.Ipv6PFlag = appendBoolToMultivalueList(ipv6Routes.Ipv6PFlag, routePFlag)
	ipv6Routes.Ipv6EFlag = appendBoolToMultivalueList(ipv6Routes.Ipv6EFlag, routeEFlag)
	ipv6Routes.Ipv6VFlag = appendBoolToMultivalueList(ipv6Routes.Ipv6VFlag, routeVFlag)
	ipv6Routes.Ipv6LFlag = appendBoolToMultivalueList(ipv6Routes.Ipv6LFlag, routeLFlag)
	return nil
}

func (ix *ixATE) addBGPProtocols(ifc *opb.InterfaceConfig) error {
	const maxNumPeersIxNetwork = 2
	if ifc.GetBgp() == nil {
		return nil
	}
	if peerCount := len(ifc.GetBgp().GetBgpPeers()); peerCount > maxNumPeersIxNetwork {
		return fmt.Errorf("specified number of peers %d exceeds maximum of %d for IxNetwork device group", peerCount, maxNumPeersIxNetwork)
	}
	v4Peers, v6Peers, v4LoopbackPeers, v6LoopbackPeers, err := splitBGPPeers(ifc.GetBgp().GetBgpPeers())
	if err != nil {
		return err
	}
	intf := ix.intfs[ifc.GetName()]
	if len(v4Peers) > 0 {
		if intf.ipv4 == nil {
			return fmt.Errorf("specified V4 peers without IPv4 configured on interface %q", ifc.GetName())
		}
		peers, err := bgpV4Peers(v4Peers)
		if err != nil {
			return err
		}
		intf.ipv4.BgpIpv4Peer = peers
		intf.idToBGPv4Peer = make(map[uint32]*ixconfig.TopologyBgpIpv4Peer)
		for i, peer := range peers {
			intf.idToBGPv4Peer[v4Peers[i].GetId()] = peer
		}
	}
	if len(v6Peers) > 0 {
		if intf.ipv6 == nil {
			return fmt.Errorf("specified V6 peers without IPv6 configured on interface %q", ifc.GetName())
		}
		peers, err := bgpV6Peers(v6Peers)
		if err != nil {
			return err
		}
		intf.ipv6.BgpIpv6Peer = peers
		intf.idToBGPv6Peer = make(map[uint32]*ixconfig.TopologyBgpIpv6Peer)
		for i, peer := range peers {
			intf.idToBGPv6Peer[v6Peers[i].GetId()] = peer
		}
	}
	if len(v4LoopbackPeers) > 0 {
		if intf.ipv4Loopback == nil {
			return fmt.Errorf("specified V4 loopback peers without IPv4 loopback configured on interface %q", ifc.GetName())
		}
		peers, err := bgpV4Peers(v4LoopbackPeers)
		if err != nil {
			return err
		}
		intf.ipv4Loopback.BgpIpv4Peer = peers
	}
	if len(v6LoopbackPeers) > 0 {
		if intf.ipv6Loopback == nil {
			return fmt.Errorf("specified V6 loopback peers without IPv6 loopback configured on interface %q", ifc.GetName())
		}
		peers, err := bgpV6Peers(v6LoopbackPeers)
		if err != nil {
			return err
		}
		intf.ipv6Loopback.BgpIpv6Peer = peers
	}
	return nil
}

func splitBGPPeers(peers []*opb.BgpPeer) ([]*opb.BgpPeer, []*opb.BgpPeer, []*opb.BgpPeer, []*opb.BgpPeer, error) {
	var v4Peers, v6Peers, v4LoopbackPeers, v6LoopbackPeers []*opb.BgpPeer
	for _, peer := range peers {
		ip, isV6 := parseIP(peer.GetPeerAddress())
		if ip == nil {
			return nil, nil, nil, nil, fmt.Errorf("invalid peer address %q", peer.GetPeerAddress())
		}
		if peer.GetOnLoopback() {
			if isV6 {
				v6LoopbackPeers = append(v6LoopbackPeers, peer)
			} else {
				v4LoopbackPeers = append(v4LoopbackPeers, peer)
			}
		} else if isV6 {
			v6Peers = append(v6Peers, peer)
		} else {
			v4Peers = append(v4Peers, peer)
		}
	}
	return v4Peers, v6Peers, v4LoopbackPeers, v6LoopbackPeers, nil
}

func bgpType(pt opb.BgpPeer_Type) *ixconfig.Multivalue {
	t := "external"
	if pt == opb.BgpPeer_TYPE_INTERNAL {
		t = "internal"
	}
	return ixconfig.MultivalueStr(t)
}

type policyCounts struct {
	numPolicies, numCommunities, maxSegListCount, maxSegCount uint32
}

func srtePolicyCounts(grps []*opb.BgpPeer_SrtePolicyGroup) (*policyCounts, error) {
	const (
		maxNumCommunitiesIxNetwork  = 32
		maxNumSegmentListsIxNetwork = 64
		maxNumSegmentsIxNetwork     = 20
	)
	var numPolicies, numCommunities, maxSegListCount, maxSegCount uint32 = 0, 0, 1, 1
	for i, grp := range grps {
		numPolicies += grp.GetCount()

		numCommsForGrp := uint32(numBGPSRTEComms(grp.GetCommunities()))
		if i > 0 && numCommunities != numCommsForGrp {
			return nil, fmt.Errorf("all policy groups for a BGP peer must have the same number of communities")
		} else if numCommsForGrp > maxNumCommunitiesIxNetwork {
			return nil, fmt.Errorf("only up to %d communities are supported per policy", maxNumCommunitiesIxNetwork)
		}
		numCommunities = numCommsForGrp

		if n := uint32(len(grp.GetSegmentLists())); n > maxNumSegmentListsIxNetwork {
			return nil, fmt.Errorf("only up to %d segment lists are supported per policy", maxNumSegmentListsIxNetwork)
		} else if n > maxSegListCount {
			maxSegListCount = n
		}
		for _, sl := range grp.GetSegmentLists() {
			if n := uint32(len(sl.GetSegments())); n > maxNumSegmentsIxNetwork {
				return nil, fmt.Errorf("only up to %d segments lists are supported per segment list", maxNumSegmentsIxNetwork)
			} else if n > maxSegCount {
				maxSegCount = n
			}
		}
	}
	return &policyCounts{
		numPolicies:     numPolicies,
		numCommunities:  numCommunities,
		maxSegListCount: maxSegListCount,
		maxSegCount:     maxSegCount,
	}, nil
}

type policyGroup struct {
	active                       bool
	distinguisher                uint32
	colorStart, colorStep        uint32
	policyType, v4EP, v6EP       string
	nhType, nhIPType, v4NH, v6NH string

	asmOverride bool
	asm         string

	enablePref bool
	pref       uint32

	enableENLP bool
	enlp       uint32

	enableBinding, asMPLS bool
	bindingType, v6SID    string
	octetStart, octetStep uint32

	enableOriginatorID bool
	originatorIDs      []string
}

func (pg *policyGroup) setTypeAndEPs(ep string) error {
	parsedEP, isV6 := parseIP(ep)
	if parsedEP == nil {
		return fmt.Errorf("invalid policy group endpoint %q", ep)
	}
	pg.policyType = "ipv4"
	pg.v4EP, pg.v6EP = ep, "::"
	if isV6 {
		pg.policyType = "ipv6"
		pg.v4EP, pg.v6EP = "0.0.0.0", ep
	}
	return nil
}

func (pg *policyGroup) setNHTypesAndAddrs(nh string) error {
	pg.nhType = "sameaslocalip"
	pg.v4NH, pg.v6NH = "0.0.0.0", "::"
	if nh != "" {
		parsedNH, isV6 := parseIP(nh)
		if parsedNH == nil {
			return fmt.Errorf("invalid policy group next hop address %q", nh)
		}
		pg.nhType = "manually"
		if isV6 {
			pg.nhIPType = "ipv6"
			pg.v6NH = nh
		} else {
			pg.nhIPType = "ipv4"
			pg.v4NH = nh
		}
	}
	return nil
}

func (pg *policyGroup) setASNSetModeAndOverride(asnSetMode opb.BgpAsnSetMode) error {
	if asnSetMode == opb.BgpAsnSetMode_ASN_SET_MODE_UNSPECIFIED {
		return nil
	}
	mode, ok := asnSetModeToStr[asnSetMode]
	if !ok {
		return fmt.Errorf("invalid policy group ASN set mode %s", asnSetMode)
	}
	pg.asm = mode
	pg.asmOverride = true
	return nil
}

func (pg *policyGroup) setBinding(bdg *opb.BgpPeer_SrtePolicyGroup_Binding) error {
	pg.v6SID = "::"
	if bdg != nil {
		pg.enableBinding = true
		switch bdg.GetType().(type) {
		case *opb.BgpPeer_SrtePolicyGroup_Binding_NoBinding:
			pg.bindingType = "nobinding"
		case *opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSid:
			pg.bindingType = "sid4"
			pg.octetStart = bdg.GetFourOctetSid().GetStart()
			pg.octetStep = bdg.GetFourOctetSid().GetStep()
		case *opb.BgpPeer_SrtePolicyGroup_Binding_FourOctetSidAsMplsLabel:
			pg.bindingType = "sid4"
			pg.octetStart = bdg.GetFourOctetSidAsMplsLabel().GetStart()
			pg.octetStep = bdg.GetFourOctetSidAsMplsLabel().GetStep()
			pg.asMPLS = true
		case *opb.BgpPeer_SrtePolicyGroup_Binding_Ipv6Sid:
			pg.bindingType = "ipv6sid"
			pg.v6SID = bdg.GetIpv6Sid()
		default:
			return fmt.Errorf("unrecognized binding type %s", bdg.GetType())
		}
	}
	return nil
}

func srtePolicyGroup(grp *opb.BgpPeer_SrtePolicyGroup) (*policyGroup, error) {
	pg := &policyGroup{
		active:        grp.GetActive(),
		distinguisher: grp.GetDistinguisher(),

		colorStart: grp.GetPolicyColor().GetStart(),
		colorStep:  grp.GetPolicyColor().GetStep(),
	}

	if grp.GetPreference() != nil {
		pg.enablePref = true
		pg.pref = grp.GetPreference().GetPreference()
	}
	if grp.GetEnlp() != nil {
		pg.enableENLP = true
		pg.enlp = grp.GetEnlp().GetEnlp()
	}

	if err := pg.setTypeAndEPs(grp.GetEndpoint()); err != nil {
		return nil, err
	}
	if err := pg.setNHTypesAndAddrs(grp.GetNextHopAddress()); err != nil {
		return nil, err
	}
	if err := pg.setASNSetModeAndOverride(grp.GetAsnSetMode()); err != nil {
		return nil, err
	}
	if err := pg.setBinding(grp.GetBinding()); err != nil {
		return nil, err
	}
	if origID := grp.GetOriginatorId(); origID != nil {
		pg.enableOriginatorID = true
		startIP, isV6 := parseIP(origID.GetStart())
		if startIP == nil || isV6 {
			return nil, fmt.Errorf("originator ID start IP %q is not a valid IPv4 address", origID.GetStart())
		}
		stepIP, isV6 := parseIP(origID.GetStep())
		if stepIP == nil || isV6 {
			return nil, fmt.Errorf("originator ID step %q is not a valid IPv4 address", origID.GetStep())
		}
		ip := new(big.Int).SetBytes(startIP[len(startIP)-4:])
		step := new(big.Int).SetBytes(stepIP[len(stepIP)-4:])
		for i := 0; i < int(grp.GetCount()); i++ {
			pg.originatorIDs = append(pg.originatorIDs, net.IP(ip.Bytes()).String())
			ip = ip.Add(ip, step)
		}
	}
	return pg, nil
}

func segment(seg *opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment) (string, uint32, uint32, uint32, bool, string, error) {
	var bos bool
	var label, tc, ttl uint32
	var segType string
	v6SID := "::"

	switch seg.GetType().(type) {
	case *opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid_:
		segType = "mplssid"
		segMpls := seg.GetMplsSid()
		label = segMpls.GetLabel()
		tc = segMpls.GetTc()
		bos = segMpls.GetS()
		ttl = segMpls.GetTtl()
	case *opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_Ipv6Sid:
		segType = "ipv6sid"
		v6SID = seg.GetIpv6Sid()
	default:
		return "", 0, 0, 0, false, "", fmt.Errorf("unrecognized segment type %s", seg.GetType())
	}
	return segType, label, tc, ttl, bos, v6SID, nil
}

func numBGPSRTEComms(communities *opb.BgpCommunities) int {
	count := 0
	if communities.GetNoExport() {
		count++
	}
	if communities.GetNoAdvertise() {
		count++
	}
	if communities.GetNoExportSubconfed() {
		count++
	}
	if communities.GetLlgrStale() {
		count++
	}
	if communities.GetNoLlgr() {
		count++
	}
	return count + len(communities.GetPrivateCommunities())
}

func bgpSRTEComms(communities *opb.BgpCommunities, commLists []*ixconfig.TopologyBgpCommunitiesList) error {
	i := 0
	setComm := func(name, asn, lto string) {
		commList := commLists[i]
		commList.Type_ = appendStrToMultivalueList(commList.Type_, name)
		commList.AsNumber = appendStrToMultivalueList(commList.AsNumber, asn)
		commList.LastTwoOctets = appendStrToMultivalueList(commList.LastTwoOctets, lto)
		i++
	}
	if communities.GetNoExport() {
		setComm("noexport", "0", "0")
	}
	if communities.GetNoAdvertise() {
		setComm("noadvertised", "0", "0")
	}
	if communities.GetNoExportSubconfed() {
		setComm("noexport_subconfed", "0", "0")
	}
	if communities.GetLlgrStale() {
		setComm("llgr_stale", "0", "0")
	}
	if communities.GetNoLlgr() {
		setComm("no_llgr", "0", "0")
	}
	for _, comm := range communities.GetPrivateCommunities() {
		commSplit := strings.SplitN(comm, ":", 2)
		if len(commSplit) != 2 {
			return fmt.Errorf("invalid format for BGP community %q", comm)
		}
		setComm("manual", commSplit[0], commSplit[1])
	}
	return nil
}

func srtePolicyListV4(grps []*opb.BgpPeer_SrtePolicyGroup) (*ixconfig.TopologyBgpSrtePoliciesListV4, uint32, error) {
	policyCounts, err := srtePolicyCounts(grps)
	if err != nil {
		return nil, 0, err
	}
	if policyCounts.numPolicies == 0 {
		return nil, 0, nil
	}

	segments := &ixconfig.TopologyBgpSrtePoliciesSegmentsCollectionV4{}
	segmentList := &ixconfig.TopologyBgpSrtePoliciesSegmentListV4{
		NumberOfSegmentsV4:                  ixconfig.NumberUint32(policyCounts.maxSegCount),
		BgpSRTEPoliciesSegmentsCollectionV4: segments,
	}
	tunnel := &ixconfig.TopologyBgpSrtePoliciesTunnelEncapsulationListV4{
		NumberOfSegmentListV4:        ixconfig.NumberUint32(policyCounts.maxSegListCount),
		BgpSRTEPoliciesSegmentListV4: segmentList,
	}
	policyList := &ixconfig.TopologyBgpSrtePoliciesListV4{
		NoOfCommunities:                          ixconfig.NumberUint32(policyCounts.numCommunities),
		BgpSRTEPoliciesTunnelEncapsulationListV4: tunnel,
	}
	for i := 0; i < int(policyCounts.numCommunities); i++ {
		policyList.BgpCommunitiesList = append(policyList.BgpCommunitiesList, &ixconfig.TopologyBgpCommunitiesList{})
	}

	policyIdx := uint32(1)
	for _, grp := range grps {
		pg, err := srtePolicyGroup(grp)
		if err != nil {
			return nil, 0, err
		}

		for i := 0; i < int(grp.GetCount()); i++ {
			policyList.Active = appendBoolToMultivalueList(policyList.Active, pg.active)
			policyList.Distinguisher = appendUintToMultivalueList(policyList.Distinguisher, pg.distinguisher)
			policyList.PolicyColor = appendUintToMultivalueList(policyList.PolicyColor, pg.colorStart+uint32(i)*pg.colorStep)

			policyList.PolicyType = appendStrToMultivalueList(policyList.PolicyType, pg.policyType)
			policyList.EndPointV4 = appendStrToMultivalueList(policyList.EndPointV4, pg.v4EP)
			policyList.EndPointV6 = appendStrToMultivalueList(policyList.EndPointV6, pg.v6EP)

			policyList.EnableNextHop = appendBoolToMultivalueList(policyList.EnableNextHop, true)
			policyList.SetNextHop = appendStrToMultivalueList(policyList.SetNextHop, pg.nhType)
			policyList.SetNextHopIpType = appendStrToMultivalueList(policyList.SetNextHopIpType, pg.nhIPType)
			policyList.Ipv4NextHop = appendStrToMultivalueList(policyList.Ipv4NextHop, pg.v4NH)
			policyList.Ipv6NextHop = appendStrToMultivalueList(policyList.Ipv6NextHop, pg.v6NH)

			policyList.OverridePeerAsSetMode = appendBoolToMultivalueList(policyList.OverridePeerAsSetMode, pg.asmOverride)
			policyList.AsSetMode = appendStrToMultivalueList(policyList.AsSetMode, pg.asm)

			origID := ""
			if pg.enableOriginatorID {
				origID = pg.originatorIDs[i]
			}
			policyList.EnableOriginatorId = appendBoolToMultivalueList(policyList.EnableOriginatorId, pg.enableOriginatorID)
			policyList.OriginatorId = appendStrToMultivalueList(policyList.OriginatorId, origID)

			tunnel.Active = appendBoolToMultivalueList(tunnel.Active, true)
			tunnel.EnPrefTLV = appendBoolToMultivalueList(tunnel.EnPrefTLV, pg.enablePref)
			tunnel.PrefValue = appendUintToMultivalueList(tunnel.PrefValue, pg.pref)
			tunnel.EnENLPTLV = appendBoolToMultivalueList(tunnel.EnENLPTLV, pg.enableENLP)
			tunnel.ENLPValue = appendUintToMultivalueList(tunnel.ENLPValue, pg.enlp)

			tunnel.EnBindingTLV = appendBoolToMultivalueList(tunnel.EnBindingTLV, pg.enableBinding)
			tunnel.BindingSIDType = appendStrToMultivalueList(tunnel.BindingSIDType, pg.bindingType)
			tunnel.UseAsMPLSLabel = appendBoolToMultivalueList(tunnel.UseAsMPLSLabel, pg.asMPLS)
			tunnel.SID4Octet = appendUintToMultivalueList(tunnel.SID4Octet, pg.octetStart+uint32(i)*pg.octetStep)
			tunnel.IPv6SID = appendStrToMultivalueList(tunnel.IPv6SID, pg.v6SID)

			policyList.EnableCommunity = appendBoolToMultivalueList(policyList.EnableCommunity, policyCounts.numCommunities > 0)
			if err := bgpSRTEComms(grp.GetCommunities(), policyList.BgpCommunitiesList); err != nil {
				return nil, 0, err
			}

			segLists := grp.GetSegmentLists()
			for j := 0; j < int(policyCounts.maxSegListCount); j++ {
				var active, enableWeight bool
				var weight uint32
				if j < len(segLists) {
					active = segLists[j].GetActive()
					if segLists[j].GetWeight() != nil {
						enableWeight = true
						weight = segLists[j].GetWeight().GetWeight()
					}
				}
				segmentList.Active = appendBoolToMultivalueList(segmentList.Active, active)
				segmentList.EnWeight = appendBoolToMultivalueList(segmentList.EnWeight, enableWeight)
				segmentList.Weight = appendUintToMultivalueList(segmentList.Weight, weight)

				var segs []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment
				if j < len(segLists) {
					segs = segLists[j].GetSegments()
				}
				for k := 0; k < int(policyCounts.maxSegCount); k++ {
					var active, bos bool
					var label, tc, ttl uint32
					segType := "ipv4nodeaddress" // Needs to be some valid value.
					v6SID := "::"
					if k < len(segs) {
						active = true
						segType, label, tc, ttl, bos, v6SID, err = segment(segs[k])
						if err != nil {
							return nil, 0, err
						}
					}
					segments.Active = appendBoolToMultivalueList(segments.Active, active)
					segments.SegmentType = appendStrToMultivalueList(segments.SegmentType, segType)
					segments.Label = appendUintToMultivalueList(segments.Label, label)
					segments.TrafficClass = appendUintToMultivalueList(segments.TrafficClass, tc)
					segments.BottomOfStack = appendBoolToMultivalueList(segments.BottomOfStack, bos)
					segments.TimeToLive = appendUintToMultivalueList(segments.TimeToLive, ttl)
					segments.Ipv6SID = appendStrToMultivalueList(segments.Ipv6SID, v6SID)
				}
			}

		}
		policyIdx += grp.GetCount()
	}
	return policyList, policyCounts.numPolicies, nil
}

func srtePolicyListV6(grps []*opb.BgpPeer_SrtePolicyGroup) (*ixconfig.TopologyBgpSrtePoliciesListV6, uint32, error) {
	policyCounts, err := srtePolicyCounts(grps)
	if err != nil {
		return nil, 0, err
	}
	if policyCounts.numPolicies == 0 {
		return nil, 0, nil
	}

	segments := &ixconfig.TopologyBgpSrtePoliciesSegmentsCollectionV6{}
	segmentList := &ixconfig.TopologyBgpSrtePoliciesSegmentListV6{
		NumberOfSegmentsV6:                  ixconfig.NumberUint32(policyCounts.maxSegCount),
		BgpSRTEPoliciesSegmentsCollectionV6: segments,
	}
	tunnel := &ixconfig.TopologyBgpSrtePoliciesTunnelEncapsulationListV6{
		NumberOfSegmentListV6:        ixconfig.NumberUint32(policyCounts.maxSegListCount),
		BgpSRTEPoliciesSegmentListV6: segmentList,
	}
	policyList := &ixconfig.TopologyBgpSrtePoliciesListV6{
		NoOfCommunities:                          ixconfig.NumberUint32(policyCounts.numCommunities),
		BgpSRTEPoliciesTunnelEncapsulationListV6: tunnel,
	}
	for i := 0; i < int(policyCounts.numCommunities); i++ {
		policyList.BgpCommunitiesList = append(policyList.BgpCommunitiesList, &ixconfig.TopologyBgpCommunitiesList{})
	}

	policyIdx := uint32(1)
	for _, grp := range grps {
		pg, err := srtePolicyGroup(grp)
		if err != nil {
			return nil, 0, err
		}

		for i := 0; i < int(grp.GetCount()); i++ {
			policyList.Active = appendBoolToMultivalueList(policyList.Active, pg.active)
			policyList.Distinguisher = appendUintToMultivalueList(policyList.Distinguisher, pg.distinguisher)
			policyList.PolicyColor = appendUintToMultivalueList(policyList.PolicyColor, pg.colorStart+uint32(i)*pg.colorStep)

			policyList.PolicyType = appendStrToMultivalueList(policyList.PolicyType, pg.policyType)
			policyList.EndPointV4 = appendStrToMultivalueList(policyList.EndPointV4, pg.v4EP)
			policyList.EndPointV6 = appendStrToMultivalueList(policyList.EndPointV6, pg.v6EP)

			policyList.EnableNextHop = appendBoolToMultivalueList(policyList.EnableNextHop, true)
			policyList.SetNextHop = appendStrToMultivalueList(policyList.SetNextHop, pg.nhType)
			policyList.SetNextHopIpType = appendStrToMultivalueList(policyList.SetNextHopIpType, pg.nhIPType)
			policyList.Ipv4NextHop = appendStrToMultivalueList(policyList.Ipv4NextHop, pg.v4NH)
			policyList.Ipv6NextHop = appendStrToMultivalueList(policyList.Ipv6NextHop, pg.v6NH)

			policyList.OverridePeerAsSetMode = appendBoolToMultivalueList(policyList.OverridePeerAsSetMode, pg.asmOverride)
			policyList.AsSetMode = appendStrToMultivalueList(policyList.AsSetMode, pg.asm)

			origID := ""
			if pg.enableOriginatorID {
				origID = pg.originatorIDs[i]
			}
			policyList.EnableOriginatorId = appendBoolToMultivalueList(policyList.EnableOriginatorId, pg.enableOriginatorID)
			policyList.OriginatorId = appendStrToMultivalueList(policyList.OriginatorId, origID)

			tunnel.Active = appendBoolToMultivalueList(tunnel.Active, true)
			tunnel.EnPrefTLV = appendBoolToMultivalueList(tunnel.EnPrefTLV, pg.enablePref)
			tunnel.PrefValue = appendUintToMultivalueList(tunnel.PrefValue, pg.pref)
			tunnel.EnENLPTLV = appendBoolToMultivalueList(tunnel.EnENLPTLV, pg.enableENLP)
			tunnel.ENLPValue = appendUintToMultivalueList(tunnel.ENLPValue, pg.enlp)

			tunnel.EnBindingTLV = appendBoolToMultivalueList(tunnel.EnBindingTLV, pg.enableBinding)
			tunnel.BindingSIDType = appendStrToMultivalueList(tunnel.BindingSIDType, pg.bindingType)
			tunnel.UseAsMPLSLabel = appendBoolToMultivalueList(tunnel.UseAsMPLSLabel, pg.asMPLS)
			tunnel.SID4Octet = appendUintToMultivalueList(tunnel.SID4Octet, pg.octetStart+uint32(i)*pg.octetStep)
			tunnel.IPv6SID = appendStrToMultivalueList(tunnel.IPv6SID, pg.v6SID)

			policyList.EnableCommunity = appendBoolToMultivalueList(policyList.EnableCommunity, policyCounts.numCommunities > 0)
			if err := bgpSRTEComms(grp.GetCommunities(), policyList.BgpCommunitiesList); err != nil {
				return nil, 0, err
			}

			segLists := grp.GetSegmentLists()
			for j := 0; j < int(policyCounts.maxSegListCount); j++ {
				var active, enableWeight bool
				var weight uint32
				if j < len(segLists) {
					active = segLists[j].GetActive()
					if segLists[j].GetWeight() != nil {
						enableWeight = true
						weight = segLists[j].GetWeight().GetWeight()
					}
				}
				segmentList.Active = appendBoolToMultivalueList(segmentList.Active, active)
				segmentList.EnWeight = appendBoolToMultivalueList(segmentList.EnWeight, enableWeight)
				segmentList.Weight = appendUintToMultivalueList(segmentList.Weight, weight)

				var segs []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment
				if j < len(segLists) {
					segs = segLists[j].GetSegments()
				}
				for k := 0; k < int(policyCounts.maxSegCount); k++ {
					var active, bos bool
					var label, tc, ttl uint32
					segType := "ipv4nodeaddress" // Needs to be some valid value.
					v6SID := "::"
					if k < len(segs) {
						active = true
						segType, label, tc, ttl, bos, v6SID, err = segment(segs[k])
						if err != nil {
							return nil, 0, err
						}
					}
					segments.Active = appendBoolToMultivalueList(segments.Active, active)
					segments.SegmentType = appendStrToMultivalueList(segments.SegmentType, segType)
					segments.Label = appendUintToMultivalueList(segments.Label, label)
					segments.TrafficClass = appendUintToMultivalueList(segments.TrafficClass, tc)
					segments.BottomOfStack = appendBoolToMultivalueList(segments.BottomOfStack, bos)
					segments.TimeToLive = appendUintToMultivalueList(segments.TimeToLive, ttl)
					segments.Ipv6SID = appendStrToMultivalueList(segments.Ipv6SID, v6SID)
				}
			}
		}
		policyIdx += grp.GetCount()
	}
	return policyList, policyCounts.numPolicies, nil
}

// All peers addresses are assumed to parse as IPv4 addresses
func bgpV4Peers(v4Peers []*opb.BgpPeer) ([]*ixconfig.TopologyBgpIpv4Peer, error) {
	var peers []*ixconfig.TopologyBgpIpv4Peer
	for _, p := range v4Peers {
		peer := &ixconfig.TopologyBgpIpv4Peer{
			Active:            ixconfig.MultivalueBool(p.GetActive()),
			Type_:             bgpType(p.GetType()),
			DutIp:             ixconfig.MultivalueStr(p.GetPeerAddress()),
			HoldTimer:         ixconfig.MultivalueUint32(p.GetHoldTimeSec()),
			KeepaliveTimer:    ixconfig.MultivalueUint32(p.GetKeepaliveTimeSec()),
			FilterIpV4Unicast: ixconfig.MultivalueTrue(),
			ActAsRestarted:    ixconfig.MultivalueBool(p.GetActAsRestarted()),
			AdvertiseEndOfRib: ixconfig.MultivalueBool(p.GetAdvertiseEndOfRib()),
			RestartTime:       ixconfig.MultivalueUint32(uint32(p.GetRestartTime().AsDuration().Seconds())),
			StaleTime:         ixconfig.MultivalueUint32(uint32(p.GetStaleTime().AsDuration().Seconds())),
		}
		caps := p.GetCapabilities()
		if caps != nil {
			peer.CapabilityIpV4Unicast = ixconfig.MultivalueBool(caps.GetIpv4Unicast())
			peer.CapabilityIpV4Multicast = ixconfig.MultivalueBool(caps.GetIpv4Multicast())
			peer.CapabilityIpV4MplsVpn = ixconfig.MultivalueBool(caps.GetIpv4MplsVpn())
			peer.CapabilityIpV6Unicast = ixconfig.MultivalueBool(caps.GetIpv6Unicast())
			peer.CapabilityIpV6Multicast = ixconfig.MultivalueBool(caps.GetIpv6Multicast())
			peer.CapabilityIpV6MplsVpn = ixconfig.MultivalueBool(caps.GetIpv6MplsVpn())
			peer.CapabilityIpV4Mdt = ixconfig.MultivalueBool(caps.GetIpv4Mdt())
			peer.CapabilityVpls = ixconfig.MultivalueBool(caps.GetVpls())
			peer.CapabilityIpV4MulticastVpn = ixconfig.MultivalueBool(caps.GetIpv4MulticastVpn())
			peer.CapabilityIpV6MulticastVpn = ixconfig.MultivalueBool(caps.GetIpv6MulticastVpn())
			peer.CapabilityRouteRefresh = ixconfig.MultivalueBool(caps.GetRouteRefresh())
			peer.CapabilityRouteConstraint = ixconfig.MultivalueBool(caps.GetRouteConstraint())
			peer.CapabilityLinkStateNonVpn = ixconfig.MultivalueBool(caps.GetLinkStateNonVpn())
			peer.Evpn = ixconfig.MultivalueBool(caps.GetEvpn())
			peer.Ipv4MulticastBgpMplsVpn = ixconfig.MultivalueBool(caps.GetIpv4MulticastBgpMplsVpn())
			peer.Ipv6MulticastBgpMplsVpn = ixconfig.MultivalueBool(caps.GetIpv6MulticastBgpMplsVpn())
			peer.Capabilityipv4UnicastFlowSpec = ixconfig.MultivalueBool(caps.GetIpv4UnicastFlowSpec())
			peer.Capabilityipv6UnicastFlowSpec = ixconfig.MultivalueBool(caps.GetIpv6UnicastFlowSpec())
			peer.CapabilityIpv4UnicastAddPath = ixconfig.MultivalueBool(caps.GetIpv4UnicastAddPath())
			peer.CapabilityIpv6UnicastAddPath = ixconfig.MultivalueBool(caps.GetIpv6UnicastAddPath())
			peer.CapabilitySRTEPoliciesV4 = ixconfig.MultivalueBool(caps.GetIpv4SrtePolicy())
			peer.CapabilitySRTEPoliciesV6 = ixconfig.MultivalueBool(caps.GetIpv6SrtePolicy())
			peer.EnableGracefulRestart = ixconfig.MultivalueBool(caps.GetGracefulRestart())
		}

		if p.GetLocalAsn() >= 65536 {
			peer.Enable4ByteAs = ixconfig.MultivalueTrue()
			peer.LocalAs4Bytes = ixconfig.MultivalueUint32(p.GetLocalAsn())
		} else {
			peer.Enable4ByteAs = ixconfig.MultivalueFalse()
			peer.LocalAs2Bytes = ixconfig.MultivalueUint32(p.GetLocalAsn())
		}

		if p.GetMd5Key() != "" {
			peer.Authentication = ixconfig.MultivalueStr("md5")
			peer.Md5Key = ixconfig.MultivalueStr(p.GetMd5Key())
		}

		policyList, numPolicies, err := srtePolicyListV4(p.GetSrtePolicyGroups())
		if err != nil {
			return nil, err
		}
		peer.NumberSRTEPolicies = ixconfig.NumberUint32(numPolicies)
		peer.BgpSRTEPoliciesListV4 = policyList
		peers = append(peers, peer)
	}
	return peers, nil
}

// All peers addresses are assumed to parse as IPv6 addresses
func bgpV6Peers(v6Peers []*opb.BgpPeer) ([]*ixconfig.TopologyBgpIpv6Peer, error) {
	var peers []*ixconfig.TopologyBgpIpv6Peer
	for _, p := range v6Peers {
		peer := &ixconfig.TopologyBgpIpv6Peer{
			Active:            ixconfig.MultivalueBool(p.GetActive()),
			Type_:             bgpType(p.GetType()),
			DutIp:             ixconfig.MultivalueStr(p.GetPeerAddress()),
			HoldTimer:         ixconfig.MultivalueUint32(p.GetHoldTimeSec()),
			KeepaliveTimer:    ixconfig.MultivalueUint32(p.GetKeepaliveTimeSec()),
			FilterIpV6Unicast: ixconfig.MultivalueTrue(),
			ActAsRestarted:    ixconfig.MultivalueBool(p.GetActAsRestarted()),
			AdvertiseEndOfRib: ixconfig.MultivalueBool(p.GetAdvertiseEndOfRib()),
			RestartTime:       ixconfig.MultivalueUint32(uint32(p.GetRestartTime().AsDuration().Seconds())),
			StaleTime:         ixconfig.MultivalueUint32(uint32(p.GetStaleTime().AsDuration().Seconds())),
		}
		caps := p.GetCapabilities()
		if caps != nil {
			peer.CapabilityIpV4Unicast = ixconfig.MultivalueBool(caps.GetIpv4Unicast())
			peer.CapabilityIpV4Multicast = ixconfig.MultivalueBool(caps.GetIpv4Multicast())
			peer.CapabilityIpV4MplsVpn = ixconfig.MultivalueBool(caps.GetIpv4MplsVpn())
			peer.CapabilityIpV6Unicast = ixconfig.MultivalueBool(caps.GetIpv6Unicast())
			peer.CapabilityIpV6Multicast = ixconfig.MultivalueBool(caps.GetIpv6Multicast())
			peer.CapabilityIpV6MplsVpn = ixconfig.MultivalueBool(caps.GetIpv6MplsVpn())
			peer.CapabilityIpV4Mdt = ixconfig.MultivalueBool(caps.GetIpv4Mdt())
			peer.CapabilityVpls = ixconfig.MultivalueBool(caps.GetVpls())
			peer.CapabilityIpV4MulticastVpn = ixconfig.MultivalueBool(caps.GetIpv4MulticastVpn())
			peer.CapabilityIpV6MulticastVpn = ixconfig.MultivalueBool(caps.GetIpv6MulticastVpn())
			peer.CapabilityRouteRefresh = ixconfig.MultivalueBool(caps.GetRouteRefresh())
			peer.CapabilityRouteConstraint = ixconfig.MultivalueBool(caps.GetRouteConstraint())
			peer.CapabilityLinkStateNonVpn = ixconfig.MultivalueBool(caps.GetLinkStateNonVpn())
			peer.Evpn = ixconfig.MultivalueBool(caps.GetEvpn())
			peer.Ipv4MulticastBgpMplsVpn = ixconfig.MultivalueBool(caps.GetIpv4MulticastBgpMplsVpn())
			peer.Ipv6MulticastBgpMplsVpn = ixconfig.MultivalueBool(caps.GetIpv6MulticastBgpMplsVpn())
			peer.Capabilityipv4UnicastFlowSpec = ixconfig.MultivalueBool(caps.GetIpv4UnicastFlowSpec())
			peer.Capabilityipv6UnicastFlowSpec = ixconfig.MultivalueBool(caps.GetIpv6UnicastFlowSpec())
			peer.CapabilityIpv4UnicastAddPath = ixconfig.MultivalueBool(caps.GetIpv4UnicastAddPath())
			peer.CapabilityIpv6UnicastAddPath = ixconfig.MultivalueBool(caps.GetIpv6UnicastAddPath())
			peer.CapabilitySRTEPoliciesV4 = ixconfig.MultivalueBool(caps.GetIpv4SrtePolicy())
			peer.CapabilitySRTEPoliciesV6 = ixconfig.MultivalueBool(caps.GetIpv6SrtePolicy())
			peer.EnableGracefulRestart = ixconfig.MultivalueBool(caps.GetGracefulRestart())
			// V6 peers only
			peer.CapabilityNHEncodingCapabilities = ixconfig.MultivalueBool(caps.GetExtendedNextHopEncoding())
		}

		if p.GetLocalAsn() >= 65536 {
			peer.Enable4ByteAs = ixconfig.MultivalueTrue()
			peer.LocalAs4Bytes = ixconfig.MultivalueUint32(p.GetLocalAsn())
		} else {
			peer.Enable4ByteAs = ixconfig.MultivalueFalse()
			peer.LocalAs2Bytes = ixconfig.MultivalueUint32(p.GetLocalAsn())
		}

		if p.GetMd5Key() != "" {
			peer.Authentication = ixconfig.MultivalueStr("md5")
			peer.Md5Key = ixconfig.MultivalueStr(p.GetMd5Key())
		}

		policyList, numPolicies, err := srtePolicyListV6(p.GetSrtePolicyGroups())
		if err != nil {
			return nil, err
		}
		peer.NumberSRTEPolicies = ixconfig.NumberUint32(numPolicies)
		peer.BgpSRTEPoliciesListV6 = policyList
		peers = append(peers, peer)
	}
	return peers, nil
}

func (ix *ixATE) addRSVPProtocols(ifc *opb.InterfaceConfig) error {
	if len(ifc.GetRsvps()) == 0 {
		return nil
	}

	intf := ix.intfs[ifc.GetName()]
	intf.rsvpLSPs = make(map[string]*ixconfig.TopologyRsvpteLsps, len(ifc.GetRsvps()))
	for rsvpIdx, rsvp := range ifc.GetRsvps() {
		if intf.ipv4 == nil {
			return fmt.Errorf("could not find IPv4 config to configure RSVP interface for %q", ifc.GetName())
		}
		if len(intf.ipv4.RsvpteIf) == 0 {
			intf.ipv4.RsvpteIf = []*ixconfig.TopologyRsvpteIf{{
				Active:                     ixconfig.MultivalueTrue(),
				EnableHelloExtension:       ixconfig.MultivalueTrue(),
				EnableBundleMessageSending: ixconfig.MultivalueBool(rsvp.GetBundleMessageSending()),
				EnableRefreshReduction:     ixconfig.MultivalueBool(rsvp.GetRefreshReduction()),
			}}
		}

		ng := intf.isrToNetworkGroup[rsvp.GetIsReachabilityName()]
		if ng == nil {
			return fmt.Errorf("could not find IS-IS reachability config %q specified for RSVP config", rsvp.GetIsReachabilityName())
		}

		numLoopbacks := len(rsvp.GetLoopbacks())
		ntl := ng.NetworkTopology.NetTopologyLinear
		// TODO(team): Will need a different check after enabling custom IS-IS topologies.
		if ntl == nil || ntl.Nodes == nil || *(ntl.Nodes) < float32(numLoopbacks) {
			return fmt.Errorf("cannot add more loopbacks than configured IS-IS nodes on reachability config %q", rsvp.GetIsReachabilityName())
		}

		// Compute/validate LSP and route object counts.
		var numIngressLSPs, numEROs, numRROs int
		for i, lb := range rsvp.GetLoopbacks() {
			lsps := lb.GetIngressLsps()
			if lspCnt := len(lsps); i == 0 {
				numIngressLSPs = lspCnt
			} else if numIngressLSPs != lspCnt {
				return fmt.Errorf("ingress LSP count must match on all loopbacks for an RSVP configuration (checks RSVP configs for IS-IS reachability config %q)", rsvp.GetIsReachabilityName())
			}

			for _, lsp := range lsps {
				if erosCnt := len(lsp.GetEros()); erosCnt != 0 && numEROs == 0 {
					numEROs = erosCnt
				} else if erosCnt != 0 && numEROs != erosCnt {
					return fmt.Errorf("ERO count must match on all configured EROs for an RSVP configuration (check RSVP configs for IS-IS reachability config %q)", rsvp.GetIsReachabilityName())
				}
				if rrosCnt := len(lsp.GetRros()); rrosCnt != 0 && numRROs == 0 {
					numRROs = rrosCnt
				} else if rrosCnt != 0 && numRROs != rrosCnt {
					return fmt.Errorf("RRO count must match on all configured RROs for an RSVP configuration (checks RSVP configs for IS-IS reachability config %q)", rsvp.GetIsReachabilityName())
				}
			}
		}

		var eros []*ixconfig.TopologyRsvpEroSubObjectsList
		for i := 0; i < numEROs; i++ {
			eros = append(eros, &ixconfig.TopologyRsvpEroSubObjectsList{})
		}
		var rros []*ixconfig.TopologyRsvpIngressRroSubObjectsList
		for i := 0; i < numRROs; i++ {
			rros = append(rros, &ixconfig.TopologyRsvpIngressRroSubObjectsList{})
		}
		ingressLSPs := &ixconfig.TopologyRsvpP2PIngressLsps{Active: ixconfig.MultivalueFalse()}
		if numIngressLSPs != 0 {
			ingressLSPs = &ixconfig.TopologyRsvpP2PIngressLsps{
				Active:                       ixconfig.MultivalueTrue(),
				NumberOfEroSubObjects:        ixconfig.NumberInt(numEROs),
				RsvpEROSubObjectsList:        eros,
				NumberOfRroSubObjects:        ixconfig.NumberInt(numRROs),
				RsvpIngressRROSubObjectsList: rros,
			}
		}
		name := rsvp.GetName()
		if name == "" {
			name = fmt.Sprintf("RSVP LSPs %d for %s", rsvpIdx, ifc.GetName())
		}
		rsvpTELSPs := &ixconfig.TopologyRsvpteLsps{
			Name:               ixconfig.String(name),
			Active:             ixconfig.MultivalueTrue(),
			EnableP2PEgress:    ixconfig.Bool(true),
			IngressP2PLsps:     ixconfig.NumberInt(numIngressLSPs),
			RsvpP2PEgressLsps:  &ixconfig.TopologyRsvpP2PEgressLsps{}, // Created only to use as a (non-nil) traffic endpoint.
			RsvpP2PIngressLsps: ingressLSPs,
		}
		intf.rsvpLSPs[name] = rsvpTELSPs
		dg := &ixconfig.TopologyDeviceGroup{
			Multiplier: ixconfig.NumberInt(numLoopbacks),
			Ipv4Loopback: []*ixconfig.TopologyIpv4Loopback{{
				RsvpteLsps: []*ixconfig.TopologyRsvpteLsps{rsvpTELSPs},
			}},
		}
		ng.DeviceGroup = append(ng.DeviceGroup, dg)
		lpCfg := dg.Ipv4Loopback[0]

		for _, lb := range rsvp.GetLoopbacks() {
			ip, mask, isV6, err := parseCIDR(lb.GetLocalIpCidr())
			if err != nil || isV6 {
				return fmt.Errorf("could not parse %q as IPv4 address for RSVP loopback: is V6? %t error? %v", lb.GetLocalIpCidr(), isV6, err)
			}
			lpCfg.Address = appendStrToMultivalueList(lpCfg.Address, ip)
			lpCfg.Prefix = appendUintToMultivalueList(lpCfg.Prefix, mask)

			for _, lsp := range lb.GetIngressLsps() {
				ip, mask, isV6, err := parseCIDR(lsp.GetRemoteIpCidr())
				if err != nil || isV6 {
					return fmt.Errorf("could not parse %q as IPv4 address for RSVP ingress LSP: is V6? %t error? %v", lsp.GetRemoteIpCidr(), isV6, err)
				}
				ingressLSPs.RemoteIp = appendStrToMultivalueList(ingressLSPs.RemoteIp, ip)
				ingressLSPs.PrefixLength = appendUintToMultivalueList(ingressLSPs.PrefixLength, mask)

				ingressLSPs.LocalProtectionDesired = appendBoolToMultivalueList(ingressLSPs.LocalProtectionDesired, lsp.GetLocalProtection())
				ingressLSPs.BandwidthProtectionDesired = appendBoolToMultivalueList(ingressLSPs.BandwidthProtectionDesired, lsp.GetBandwidthProtection())
				ingressLSPs.EnableFastReroute = appendBoolToMultivalueList(ingressLSPs.EnableFastReroute, lsp.GetFastReroute())
				ingressLSPs.EnablePathReOptimization = appendBoolToMultivalueList(ingressLSPs.EnablePathReOptimization, lsp.GetPathReoptimization())

				ingressLSPs.TunnelId = appendUintToMultivalueList(ingressLSPs.TunnelId, lsp.GetTunnelId())
				ingressLSPs.LspId = appendUintToMultivalueList(ingressLSPs.LspId, lsp.GetLspId())

				enableEROs := len(lsp.GetEros()) > 0
				ingressLSPs.EnableEro = appendBoolToMultivalueList(ingressLSPs.EnableEro, enableEROs)
				for i := 0; i < numEROs; i++ {
					var ip string
					var mask uint32
					if enableEROs {
						var isV6 bool
						var err error
						cidr := lsp.GetEros()[i].GetIpv4Cidr()
						ip, mask, isV6, err = parseCIDR(cidr)
						if err != nil || isV6 {
							return fmt.Errorf("could not parse %q as IPv4 address for RSVP ERO object: is V6? %t error? %v", cidr, isV6, err)
						}
					}
					eros[i].Ip = appendStrToMultivalueList(eros[i].Ip, ip)
					eros[i].PrefixLength = appendUintToMultivalueList(eros[i].PrefixLength, mask)
				}

				enableRROs := len(lsp.GetRros()) > 0
				ingressLSPs.SendRro = appendBoolToMultivalueList(ingressLSPs.SendRro, enableRROs)
				for i := 0; i < numRROs; i++ {
					var ip, addr string
					if enableRROs {
						ip = lsp.GetRros()[i].GetIpv4()
					}
					if ip != "" {
						ipAddr, isV6 := parseIP(ip)
						if ipAddr == nil || isV6 {
							return fmt.Errorf("invalid IP address %q for RSVP RRO object", ip)
						}
						addr = ipAddr.String()
					}
					rros[i].Ip = appendStrToMultivalueList(rros[i].Ip, addr)
				}
			}
		}
	}
	return nil
}

func (ix *ixATE) addDHCPProtocols(ifc *opb.InterfaceConfig) error {
	intf := ix.intfs[ifc.GetName()]

	if dhcp6c := ifc.GetDhcpv6Client(); dhcp6c != nil {
		eth := intf.deviceGroup.Ethernet[0]
		eth.Dhcpv6client = []*ixconfig.TopologyDhcpv6client{{}}
	}

	if dhcp6s := ifc.GetDhcpv6Server(); dhcp6s != nil {
		step, err := addrRangeToStep(dhcp6s.LeaseAddrs, ipv6AddrType)
		if err != nil {
			return err
		}
		intf.ipv6.Dhcpv6server = []*ixconfig.TopologyDhcpv6server{{
			Dhcp6ServerSessions: &ixconfig.TopologyDhcp6ServerSessions{
				IpAddress:          ixconfig.MultivalueStr(dhcp6s.LeaseAddrs.GetMin()),
				IpAddressIncrement: ixconfig.MultivalueStr(step),
				PoolSize:           ixconfig.MultivalueUint32(dhcp6s.LeaseAddrs.GetCount()),
			},
		}}
	}

	return nil
}

func appendStrToMultivalueList(mv *ixconfig.Multivalue, val string) *ixconfig.Multivalue {
	if mv == nil {
		mv = &ixconfig.Multivalue{ValueList: &ixconfig.MultivalueValueList{}}
		mv.Pattern = ixconfig.String("valueList")
	}
	mv.ValueList.Values = append(mv.ValueList.Values, val)
	return mv
}

func appendBoolToMultivalueList(mv *ixconfig.Multivalue, val bool) *ixconfig.Multivalue {
	return appendStrToMultivalueList(mv, strconv.FormatBool(val))
}

func appendIntToMultivalueList(mv *ixconfig.Multivalue, val int) *ixconfig.Multivalue {
	return appendStrToMultivalueList(mv, strconv.Itoa(val))
}

func appendUintToMultivalueList(mv *ixconfig.Multivalue, val uint32) *ixconfig.Multivalue {
	return appendStrToMultivalueList(mv, strconv.FormatUint(uint64(val), 10))
}

func appendUint64ToMultivalueList(mv *ixconfig.Multivalue, val uint64) *ixconfig.Multivalue {
	return appendStrToMultivalueList(mv, strconv.FormatUint(val, 10))
}
