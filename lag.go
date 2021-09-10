package ondatra

import (
	opb "github.com/openconfig/ondatra/proto"
)

// LAG is a link aggregation group.
type LAG struct {
	pb *opb.Lag
}

// WithPorts sets the LAG ports to the specified ports.
func (l *LAG) WithPorts(ports ...*Port) *LAG {
	l.pb.Ports = nil
	for _, p := range ports {
		l.pb.Ports = append(l.pb.Ports, p.Name())
	}
	return l
}
