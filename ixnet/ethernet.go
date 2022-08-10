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

// NewEthernet returns a new ethernet configuration.
// Tests must not call this method directly.
func NewEthernet(pb *opb.EthernetConfig) *Ethernet {
	return &Ethernet{pb}
}

// Ethernet is a representation of Ethernet config on the ATE.
type Ethernet struct {
	pb *opb.EthernetConfig
}

// WithMTU specifies the MTU in bytes for the interface.
func (e *Ethernet) WithMTU(mtu uint16) *Ethernet {
	e.pb.Mtu = uint32(mtu)
	return e
}

// WithVLANID enables a VLAN with a 12-bit specified ID on the interface.
func (e *Ethernet) WithVLANID(id uint16) *Ethernet {
	e.pb.VlanId = uint32(id)
	return e
}

// MACsec creates a MACsec config or returns the existing config.
// The default config params are:
// Encrypted Traffic: Stateless L2-3
// Cipher Suite: GCM-AES-128
// Confidentiality Offset: 0
// Encrypted VLANs Enabled: false
func (e *Ethernet) MACsec() *MACsec {
	if e.pb.Macsec == nil {
		e.pb.Macsec = &opb.MacSec{
			CipherSuite: opb.MacSec_AES_128,
		}
	}
	return &MACsec{pb: e.pb.Macsec}
}

// MACsec is a representation of MACsec config on an ATE.
type MACsec struct {
	pb *opb.MacSec
}

// MKA is a representation of the MACSec Key Agreement protocol on an ATE.
type MKA struct {
	pb *opb.MacSec_MKA
}

// ConnectivityAssociation is a representation of a connectivity association for an MKA protocol.
type ConnectivityAssociation struct {
	pb *opb.MacSec_MKA_ConnectivityAssociation
}

// WithCipherSuiteAES128 sets the cipher suite to be GCM-AES-128.
func (m *MACsec) WithCipherSuiteAES128() *MACsec {
	m.pb.CipherSuite = opb.MacSec_AES_128
	return m
}

// WithCipherSuiteAES256 sets the cipher suite to be GCM-AES-256.
func (m *MACsec) WithCipherSuiteAES256() *MACsec {
	m.pb.CipherSuite = opb.MacSec_AES_256
	return m
}

// WithCipherSuiteAESXPN128 sets the cipher suite to be GCM-AES-XPN-128.
func (m *MACsec) WithCipherSuiteAESXPN128() *MACsec {
	m.pb.CipherSuite = opb.MacSec_AES_XPN_128
	return m
}

// WithCipherSuiteAESXPN256 sets the cipher suite to be GCM-AES-XPN-256.
func (m *MACsec) WithCipherSuiteAESXPN256() *MACsec {
	m.pb.CipherSuite = opb.MacSec_AES_XPN_256
	return m
}

// WithEncryptedVLANsEnabled sets whether encrypted VLANS are enabled.
func (m *MACsec) WithEncryptedVLANsEnabled(enabled bool) *MACsec {
	m.pb.EncryptedVlansEnabled = enabled
	return m
}

// RxSAKPool create an Rx SAK Pool or returns the existing pool.
func (m *MACsec) RxSAKPool() *RxSAKPool {
	if m.pb.RxSakPool == nil {
		m.pb.RxSakPool = &opb.RxSakPool{}
	}
	return &RxSAKPool{pb: m.pb.RxSakPool}
}

// RxSAKPool is a representation of an Rx SAK Pool in MACsec.
type RxSAKPool struct {
	pb *opb.RxSakPool
}

// WithSAK128 specifies the 128-bit receive SAK in the Rx SAK pool.
func (r *RxSAKPool) WithSAK128(sak string) *RxSAKPool {
	r.pb.Sak128 = sak
	return r
}

// WithSAK256 specifies the 256-bit receive SAK in the Rx SAK pool.
func (r *RxSAKPool) WithSAK256(sak string) *RxSAKPool {
	r.pb.Sak256 = sak
	return r
}

// MKA creates an MKA protocol for a MACsec configuration or returns the existing config.
// The default config params are:
//
//	Capability: Not implemented
//	Cipher Suite: GCM-AES-128
//	Confidentiality Offset: No confidentiality
func (m *MACsec) MKA() *MKA {
	if m.pb.Mka == nil {
		m.pb.Mka = &opb.MacSec_MKA{
			Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
			CipherSuite:           opb.MacSec_AES_128,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_NO_CONFIDENTIALITY,
		}
	}
	return &MKA{pb: m.pb.Mka}
}

// WithVersion specifies the MKA protocol version.
func (m *MKA) WithVersion(version uint32) *MKA {
	m.pb.Version = version
	return m
}

// WithCapabilityNotImplemented sets the MKA protocol to not indicate any MacSec capability.
func (m *MKA) WithCapabilityNotImplemented() *MKA {
	m.pb.Capability = opb.MacSec_MKA_NOT_IMPLEMENTED
	return m
}

// WithCapabilityIntegrityWithoutConfidentiality sets the MKA protocol to indicate the integrity without confidentiality MacSec capability.
func (m *MKA) WithCapabilityIntegrityWithoutConfidentiality() *MKA {
	m.pb.Capability = opb.MacSec_MKA_INTEGRITY_WITHOUT_CONFIDENTIALITY
	return m
}

// WithCapabilityIntegrityWithNoConfidentialityOffset sets the MKA protocol to indicate the confidentiality without offset MacSec capability.
func (m *MKA) WithCapabilityIntegrityWithNoConfidentialityOffset() *MKA {
	m.pb.Capability = opb.MacSec_MKA_INTEGRITY_WITH_NO_CONFIDENTIALITY_OFFSET
	return m
}

// WithCapabilityIntegrityWithConfidentialityOffset sets the MKA protocol to indicate the confidentiality with offset MacSec capability.
func (m *MKA) WithCapabilityIntegrityWithConfidentialityOffset() *MKA {
	m.pb.Capability = opb.MacSec_MKA_INTEGRITY_WITH_CONFIDENTIALITY_OFFSET
	return m
}

// WithConfidentialityOffsetNoConfidentiality sets the confidentiality offset to indicate no confidentiality.
func (m *MKA) WithConfidentialityOffsetNoConfidentiality() *MKA {
	m.pb.ConfidentialityOffset = opb.MacSec_MKA_OFFSET_NO_CONFIDENTIALITY
	return m
}

// WithConfidentialityOffset0 sets the confidentiality offset to be 0.
func (m *MKA) WithConfidentialityOffset0() *MKA {
	m.pb.ConfidentialityOffset = opb.MacSec_MKA_OFFSET_0
	return m
}

// WithConfidentialityOffset30 sets the confidentiality offset to be 30.
func (m *MKA) WithConfidentialityOffset30() *MKA {
	m.pb.ConfidentialityOffset = opb.MacSec_MKA_OFFSET_30
	return m
}

// WithConfidentialityOffset50 sets the confidentiality offset to be 50.
func (m *MKA) WithConfidentialityOffset50() *MKA {
	m.pb.ConfidentialityOffset = opb.MacSec_MKA_OFFSET_50
	return m
}

// WithCipherSuiteAES128 sets the cipher suite to be GCM-AES-128.
func (m *MKA) WithCipherSuiteAES128() *MKA {
	m.pb.CipherSuite = opb.MacSec_AES_128
	return m
}

// WithCipherSuiteAES256 sets the cipher suite to be GCM-AES-256.
func (m *MKA) WithCipherSuiteAES256() *MKA {
	m.pb.CipherSuite = opb.MacSec_AES_256
	return m
}

// WithCipherSuiteAESXPN128 sets the cipher suite to be GCM-AES-XPN-128.
func (m *MKA) WithCipherSuiteAESXPN128() *MKA {
	m.pb.CipherSuite = opb.MacSec_AES_XPN_128
	return m
}

// WithCipherSuiteAESXPN256 sets the cipher suite to be GCM-AES-XPN-256.
func (m *MKA) WithCipherSuiteAESXPN256() *MKA {
	m.pb.CipherSuite = opb.MacSec_AES_XPN_256
	return m
}

// ConnectivityAssociation creates a ConnectivityAssociation config for an MKA protocol or returns the existing config.
func (m *MKA) ConnectivityAssociation() *ConnectivityAssociation {
	if m.pb.ConnectivityAssociation == nil {
		m.pb.ConnectivityAssociation = &opb.MacSec_MKA_ConnectivityAssociation{}
	}
	return &ConnectivityAssociation{pb: m.pb.ConnectivityAssociation}
}

// WithCKN sets the key name, specified as a hex string.
func (c *ConnectivityAssociation) WithCKN(name string) *ConnectivityAssociation {
	c.pb.Ckn = name
	return c
}

// WithCAK sets the key, specified as a hex string.
func (c *ConnectivityAssociation) WithCAK(key string) *ConnectivityAssociation {
	c.pb.Cak = key
	return c
}

// FEC creates a an FEC config or returns the existing config.
func (e *Ethernet) FEC() *FEC {
	return &FEC{pb: e.pb.Fec}
}

// FEC is a representation of forward error correction settings on an ATE.
type FEC struct {
	pb *opb.Fec
}

// WithEnabled sets whether FEC is enabled.
func (f *FEC) WithEnabled(enabled bool) *FEC {
	f.pb.Enabled = enabled
	return f
}
