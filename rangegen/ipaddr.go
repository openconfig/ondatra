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

package rangegen

import (
	"errors"
	"fmt"
	"math/big"
	"net"
	"strings"
)

var (
	maxIPv4 = maxVal(net.IPv4len)
	maxIPv6 = maxVal(net.IPv6len)
)

// network represents an IPv4 or IPv6 network.
type network struct {
	ip      []byte
	maskLen uint
	isIPv6  bool
}

// parseCIDR returns the network represented by the given CIDR.
// Returns an error if the CIDR is not valid.
func parseCIDR(cidr string) (*network, error) {
	_, netw, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	maskLen, _ := netw.Mask.Size()
	if strings.Contains(cidr, ":") {
		return &network{
			ip:      netw.IP,
			maskLen: uint(maskLen),
			isIPv6:  true,
		}, nil
	}
	return &network{
		ip:      netw.IP.To4(),
		maskLen: uint(maskLen),
	}, nil
}

// CIDRs returns a channel of the specified number of CIDR strings starting from the given CIDR.
func CIDRs(cidr string, count uint32) (<-chan string, error) {
	const (
		maxIPv4StrLen = len("255.255.255.255")
		maxIPv6StrLen = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	)
	n, err := parseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	ipLen := net.IPv4len
	maxIPStrLen := maxIPv4StrLen
	maxIP := maxIPv4
	toStr := toIPv4String
	if n.isIPv6 {
		ipLen = net.IPv6len
		maxIPStrLen = maxIPv6StrLen
		maxIP = maxIPv6
		toStr = toIPv6String
	}

	ip := new(big.Int).SetBytes(n.ip)
	inc := new(big.Int).Lsh(one, 8*uint(ipLen)-n.maskLen)
	if err := checkMax(ip, inc, maxIP, count); err != nil {
		return nil, err
	}
	ipBytes := make([]byte, ipLen)
	maskBytes := []byte(fmt.Sprintf("/%d", n.maskLen))
	strBytes := make([]byte, maxIPStrLen+len(maskBytes))

	netsCh := make(chan string, 100)
	go func() {
		defer close(netsCh)
		for i := 0; i < int(count); i++ {
			ip.FillBytes(ipBytes)
			ip.Add(ip, inc)
			netsCh <- toStr(ipBytes, strBytes, maskBytes)
		}
	}()
	return netsCh, nil
}

func checkMax(start, inc, max *big.Int, numVals uint32) error {
	num := big.NewInt(int64(numVals))
	length := num.Mul(num, inc)
	last := length.Add(start, length)
	if last.Cmp(max) == 1 {
		return errors.New("requested network range exceeds max IP value")
	}
	return nil
}

// Logic copied from net.IP.String(), but modified to accept a preallocated output slice and mask.
func toIPv4String(ip, str, mask []byte) string {
	n := ubtoa(str, 0, ip[0])
	str[n] = '.'
	n++

	n += ubtoa(str, n, ip[1])
	str[n] = '.'
	n++

	n += ubtoa(str, n, ip[2])
	str[n] = '.'
	n++

	n += ubtoa(str, n, ip[3])
	return string(append(str[:n], mask...))
}

// Copied from the 'net' package.
// Encodes the string form of the integer v to dst[start:] and returns the number of bytes written to dst.
func ubtoa(dst []byte, start int, v byte) int {
	if v < 10 {
		dst[start] = v + '0'
		return 1
	} else if v < 100 {
		dst[start+1] = v%10 + '0'
		dst[start] = v/10 + '0'
		return 2
	}

	dst[start+2] = v%10 + '0'
	dst[start+1] = (v/10)%10 + '0'
	dst[start] = v/100 + '0'
	return 3
}

// Logic copied from net.Ip.String(), but modified to accept a preallocated output slice and mask.
func toIPv6String(ip, str, mask []byte) string {
	str = str[:0]
	// Find longest run of zeros.
	e0 := -1
	e1 := -1
	for i := 0; i < net.IPv6len; i += 2 {
		j := i
		for j < net.IPv6len && ip[j] == 0 && ip[j+1] == 0 {
			j += 2
		}
		if j > i && j-i > e1-e0 {
			e0 = i
			e1 = j
			i = j
		}
	}
	// The symbol "::" MUST NOT be used to shorten just one 16 bit 0 field.
	if e1-e0 <= 2 {
		e0 = -1
		e1 = -1
	}

	// Print with possible :: in place of run of zeros
	for i := 0; i < net.IPv6len; i += 2 {
		if i == e0 {
			str = append(str, ':', ':')
			i = e1
			if i >= net.IPv6len {
				break
			}
		} else if i > 0 {
			str = append(str, ':')
		}
		str = appendHex(str, (uint32(ip[i])<<8)|uint32(ip[i+1]))
	}
	return string(append(str, mask...))
}

// Copied from the 'net' package. Converts i to a hexadecimal string. Leading zeros are not printed.
func appendHex(dst []byte, i uint32) []byte {
	const hexDigit = "0123456789abcdef"
	if i == 0 {
		return append(dst, '0')
	}
	for j := 7; j >= 0; j-- {
		v := i >> uint(j*4)
		if v > 0 {
			dst = append(dst, hexDigit[v&0xf])
		}
	}
	return dst
}
