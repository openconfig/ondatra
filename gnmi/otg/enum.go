//go:build go1.18

/*
Package otg is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

This package was generated by ygnmi version: v0.1.0: (ygot: v0.23.1)
using the following YANG input files:
  - models-yang/models/isis/open-traffic-generator-isis.yang
  - models-yang/models/types/open-traffic-generator-types.yang
  - models-yang/models/flow/open-traffic-generator-flow.yang
  - models-yang/models/discovery/open-traffic-generator-discovery.yang
  - models-yang/models/interface/open-traffic-generator-port.yang
  - models-yang/models/bgp/open-traffic-generator-bgp.yang

Imported modules were sourced from:
  - models-yang/models/...
*/
package otg

import (
	"github.com/openconfig/ygot/ygot"
)

// E_BgpPeer_SessionState is a derived int64 type which is used to represent
// the enumerated node BgpPeer_SessionState. An additional value named
// BgpPeer_SessionState_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_BgpPeer_SessionState int64

// IsYANGGoEnum ensures that BgpPeer_SessionState implements the yang.GoEnum
// interface. This ensures that BgpPeer_SessionState can be identified as a
// mapped type for a YANG enumeration.
func (E_BgpPeer_SessionState) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  BgpPeer_SessionState.
func (E_BgpPeer_SessionState) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_BgpPeer_SessionState.
func (e E_BgpPeer_SessionState) String() string {
	return ygot.EnumLogString(e, int64(e), "E_BgpPeer_SessionState")
}

const (
	// BgpPeer_SessionState_UNSET corresponds to the value UNSET of BgpPeer_SessionState
	BgpPeer_SessionState_UNSET E_BgpPeer_SessionState = 0
	// BgpPeer_SessionState_IDLE corresponds to the value IDLE of BgpPeer_SessionState
	BgpPeer_SessionState_IDLE E_BgpPeer_SessionState = 1
	// BgpPeer_SessionState_CONNECT corresponds to the value CONNECT of BgpPeer_SessionState
	BgpPeer_SessionState_CONNECT E_BgpPeer_SessionState = 2
	// BgpPeer_SessionState_ACTIVE corresponds to the value ACTIVE of BgpPeer_SessionState
	BgpPeer_SessionState_ACTIVE E_BgpPeer_SessionState = 3
	// BgpPeer_SessionState_OPEN_SENT corresponds to the value OPEN_SENT of BgpPeer_SessionState
	BgpPeer_SessionState_OPEN_SENT E_BgpPeer_SessionState = 4
	// BgpPeer_SessionState_OPEN_CONFIRM corresponds to the value OPEN_CONFIRM of BgpPeer_SessionState
	BgpPeer_SessionState_OPEN_CONFIRM E_BgpPeer_SessionState = 5
	// BgpPeer_SessionState_ESTABLISHED corresponds to the value ESTABLISHED of BgpPeer_SessionState
	BgpPeer_SessionState_ESTABLISHED E_BgpPeer_SessionState = 6
)

// E_Port_Link is a derived int64 type which is used to represent
// the enumerated node Port_Link. An additional value named
// Port_Link_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_Port_Link int64

// IsYANGGoEnum ensures that Port_Link implements the yang.GoEnum
// interface. This ensures that Port_Link can be identified as a
// mapped type for a YANG enumeration.
func (E_Port_Link) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  Port_Link.
func (E_Port_Link) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_Port_Link.
func (e E_Port_Link) String() string {
	return ygot.EnumLogString(e, int64(e), "E_Port_Link")
}

const (
	// Port_Link_UNSET corresponds to the value UNSET of Port_Link
	Port_Link_UNSET E_Port_Link = 0
	// Port_Link_UP corresponds to the value UP of Port_Link
	Port_Link_UP E_Port_Link = 1
	// Port_Link_DOWN corresponds to the value DOWN of Port_Link
	Port_Link_DOWN E_Port_Link = 2
)

// E_State_CommunityType is a derived int64 type which is used to represent
// the enumerated node State_CommunityType. An additional value named
// State_CommunityType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_State_CommunityType int64

// IsYANGGoEnum ensures that State_CommunityType implements the yang.GoEnum
// interface. This ensures that State_CommunityType can be identified as a
// mapped type for a YANG enumeration.
func (E_State_CommunityType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  State_CommunityType.
func (E_State_CommunityType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_State_CommunityType.
func (e E_State_CommunityType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_State_CommunityType")
}

const (
	// State_CommunityType_UNSET corresponds to the value UNSET of State_CommunityType
	State_CommunityType_UNSET E_State_CommunityType = 0
	// State_CommunityType_MANUAL_AS_NUMBER corresponds to the value MANUAL_AS_NUMBER of State_CommunityType
	State_CommunityType_MANUAL_AS_NUMBER E_State_CommunityType = 1
	// State_CommunityType_NO_EXPORT corresponds to the value NO_EXPORT of State_CommunityType
	State_CommunityType_NO_EXPORT E_State_CommunityType = 2
	// State_CommunityType_NO_ADVERTISE corresponds to the value NO_ADVERTISE of State_CommunityType
	State_CommunityType_NO_ADVERTISE E_State_CommunityType = 3
	// State_CommunityType_NO_EXPORT_SUBCONFED corresponds to the value NO_EXPORT_SUBCONFED of State_CommunityType
	State_CommunityType_NO_EXPORT_SUBCONFED E_State_CommunityType = 4
	// State_CommunityType_LLGR_STALE corresponds to the value LLGR_STALE of State_CommunityType
	State_CommunityType_LLGR_STALE E_State_CommunityType = 5
	// State_CommunityType_NO_LLGR corresponds to the value NO_LLGR of State_CommunityType
	State_CommunityType_NO_LLGR E_State_CommunityType = 6
)

// E_State_SegmentType is a derived int64 type which is used to represent
// the enumerated node State_SegmentType. An additional value named
// State_SegmentType_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_State_SegmentType int64

// IsYANGGoEnum ensures that State_SegmentType implements the yang.GoEnum
// interface. This ensures that State_SegmentType can be identified as a
// mapped type for a YANG enumeration.
func (E_State_SegmentType) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  State_SegmentType.
func (E_State_SegmentType) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_State_SegmentType.
func (e E_State_SegmentType) String() string {
	return ygot.EnumLogString(e, int64(e), "E_State_SegmentType")
}

const (
	// State_SegmentType_UNSET corresponds to the value UNSET of State_SegmentType
	State_SegmentType_UNSET E_State_SegmentType = 0
	// State_SegmentType_AS_SEQUENCE corresponds to the value AS_SEQUENCE of State_SegmentType
	State_SegmentType_AS_SEQUENCE E_State_SegmentType = 1
	// State_SegmentType_AS_SET corresponds to the value AS_SET of State_SegmentType
	State_SegmentType_AS_SET E_State_SegmentType = 2
	// State_SegmentType_AS_CONFED_SEQUENCE corresponds to the value AS_CONFED_SEQUENCE of State_SegmentType
	State_SegmentType_AS_CONFED_SEQUENCE E_State_SegmentType = 3
	// State_SegmentType_AS_CONFED_SET corresponds to the value AS_CONFED_SET of State_SegmentType
	State_SegmentType_AS_CONFED_SET E_State_SegmentType = 4
)

// E_UnicastIpv4Prefix_Origin is a derived int64 type which is used to represent
// the enumerated node UnicastIpv4Prefix_Origin. An additional value named
// UnicastIpv4Prefix_Origin_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_UnicastIpv4Prefix_Origin int64

// IsYANGGoEnum ensures that UnicastIpv4Prefix_Origin implements the yang.GoEnum
// interface. This ensures that UnicastIpv4Prefix_Origin can be identified as a
// mapped type for a YANG enumeration.
func (E_UnicastIpv4Prefix_Origin) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  UnicastIpv4Prefix_Origin.
func (E_UnicastIpv4Prefix_Origin) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_UnicastIpv4Prefix_Origin.
func (e E_UnicastIpv4Prefix_Origin) String() string {
	return ygot.EnumLogString(e, int64(e), "E_UnicastIpv4Prefix_Origin")
}

const (
	// UnicastIpv4Prefix_Origin_UNSET corresponds to the value UNSET of UnicastIpv4Prefix_Origin
	UnicastIpv4Prefix_Origin_UNSET E_UnicastIpv4Prefix_Origin = 0
	// UnicastIpv4Prefix_Origin_IGP corresponds to the value IGP of UnicastIpv4Prefix_Origin
	UnicastIpv4Prefix_Origin_IGP E_UnicastIpv4Prefix_Origin = 1
	// UnicastIpv4Prefix_Origin_EGP corresponds to the value EGP of UnicastIpv4Prefix_Origin
	UnicastIpv4Prefix_Origin_EGP E_UnicastIpv4Prefix_Origin = 2
)

// E_UnicastIpv6Prefix_Origin is a derived int64 type which is used to represent
// the enumerated node UnicastIpv6Prefix_Origin. An additional value named
// UnicastIpv6Prefix_Origin_UNSET is added to the enumeration which is used as
// the nil value, indicating that the enumeration was not explicitly set by
// the program importing the generated structures.
type E_UnicastIpv6Prefix_Origin int64

// IsYANGGoEnum ensures that UnicastIpv6Prefix_Origin implements the yang.GoEnum
// interface. This ensures that UnicastIpv6Prefix_Origin can be identified as a
// mapped type for a YANG enumeration.
func (E_UnicastIpv6Prefix_Origin) IsYANGGoEnum() {}

// ΛMap returns the value lookup map associated with  UnicastIpv6Prefix_Origin.
func (E_UnicastIpv6Prefix_Origin) ΛMap() map[string]map[int64]ygot.EnumDefinition { return ΛEnum }

// String returns a logging-friendly string for E_UnicastIpv6Prefix_Origin.
func (e E_UnicastIpv6Prefix_Origin) String() string {
	return ygot.EnumLogString(e, int64(e), "E_UnicastIpv6Prefix_Origin")
}

const (
	// UnicastIpv6Prefix_Origin_UNSET corresponds to the value UNSET of UnicastIpv6Prefix_Origin
	UnicastIpv6Prefix_Origin_UNSET E_UnicastIpv6Prefix_Origin = 0
	// UnicastIpv6Prefix_Origin_IGP corresponds to the value IGP of UnicastIpv6Prefix_Origin
	UnicastIpv6Prefix_Origin_IGP E_UnicastIpv6Prefix_Origin = 1
	// UnicastIpv6Prefix_Origin_EGP corresponds to the value EGP of UnicastIpv6Prefix_Origin
	UnicastIpv6Prefix_Origin_EGP E_UnicastIpv6Prefix_Origin = 2
)