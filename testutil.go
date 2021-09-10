package ondatra

import (
	"path/filepath"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/fakebind"
	"github.com/openconfig/ondatra/internal/reservation"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	fakeBind   *fakebind.Binding
	fakeTBPath = filepath.Join("testdata", "testbed.pb.txt")

	fakeRes = &reservation.Reservation{
		DUTs: map[string]*reservation.DUT{
			"dut": &reservation.DUT{&reservation.Dims{
				Name:            "pf01.xxx01",
				Vendor:          opb.Device_ARISTA,
				HardwareModel:   "aristaModel",
				SoftwareVersion: "aristaVersion",
				Ports: map[string]*reservation.Port{
					"port1": &reservation.Port{Name: "Et1/2/3"},
					"port2": &reservation.Port{Name: "Et4/5/6"},
				},
			}},
			"dut_cisco": &reservation.DUT{&reservation.Dims{
				Name:            "pf02.xxx01",
				Vendor:          opb.Device_CISCO,
				HardwareModel:   "ciscoModel",
				SoftwareVersion: "ciscoVersion",
			}},
			"dut_juniper": &reservation.DUT{&reservation.Dims{
				Name:            "pf03.xxx01",
				Vendor:          opb.Device_JUNIPER,
				HardwareModel:   "juniperModel",
				SoftwareVersion: "juniperVersion",
			}},
		},
		ATEs: map[string]*reservation.ATE{
			"ate": &reservation.ATE{&reservation.Dims{
				Name:            "ix1",
				Vendor:          opb.Device_IXIA,
				HardwareModel:   "ixiaModel",
				SoftwareVersion: "ixiaVersion",
				Ports: map[string]*reservation.Port{
					"port1": &reservation.Port{Name: "1/1"},
					"port2": &reservation.Port{Name: "1/2"},
				},
			}},
		},
	}
)

// Initializes Ondatra with a fake binding implementation.
func initFakeBinding() {
	fakeBind = &fakebind.Binding{Reservation: fakeRes}
	binding.Init(fakeBind)
}

// Reserves the fake testbed.
// Should only be called within a test's init() function.
func reserveFakeTestbed() {
	if err := reserve(fakeTBPath, time.Hour, 0); err != nil {
		log.Fatal("Failed to reserve testbed: %v", err)
	}
}
