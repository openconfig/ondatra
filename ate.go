package ondatra

import "github.com/openconfig/ondatra/internal/reservation"

// ATEDevice is an automated test equipment.
type ATEDevice struct {
	*Device
}

// Topology returns a handle to the topology API.
func (a *ATEDevice) Topology() *Topology {
	return &Topology{a.res.(*reservation.ATE)}
}

// Traffic returns a handle to the traffic API.
func (a *ATEDevice) Traffic() *Traffic {
	return &Traffic{a.res.(*reservation.ATE)}
}
