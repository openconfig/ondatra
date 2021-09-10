package ondatra

import (
	opb "github.com/openconfig/ondatra/proto"
)

// IP is an representation of IP config on the ATE.
type IP struct {
	pb *opb.IpConfig
}

// WithAddress sets the IP address in CIDR notation.
func (i *IP) WithAddress(address string) *IP {
	i.pb.AddressCidr = address
	return i
}

// WithDefaultGateway sets the default gateway.
func (i *IP) WithDefaultGateway(gateway string) *IP {
	i.pb.DefaultGateway = gateway
	return i
}
