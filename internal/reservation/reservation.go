// Package reservation provides common utilities for using reservation.
package reservation

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/openconfig/ondatra/internal/usererr/usererr"
	"github.com/pkg/errors"

	opb "github.com/openconfig/ondatra/proto/ondatra_go_proto"
)

// Reservation holds the reserved DUTs and ATEs as an id map.
type Reservation struct {
	ID   string
	DUTs map[string]*DUT
	ATEs map[string]*ATE
}

// Device is a reserved DUT or ATE.
type Device interface {
	Dimensions() *Dims
}

// Dims contains the dimensions of reserved DUT or ATE.
type Dims struct {
	Name            string
	Vendor          opb.Device_Vendor
	HardwareModel   string
	SoftwareVersion string
	Ports           map[string]*Port
}

func (d *Dims) String() string {
	return fmt.Sprintf("Dims%+v", *d)
}

// DUT is a reserved DUT
type DUT struct {
	*Dims
}

// Dimensions returns the dimensions of the device.
func (d *DUT) Dimensions() *Dims {
	return d.Dims
}

func (d *DUT) String() string {
	return fmt.Sprintf("DUT%+v", *d)
}

// ATE is a reserved ATE.
type ATE struct {
	*Dims
}

// Dimensions returns the dimensions of the device.
func (a *ATE) Dimensions() *Dims {
	return a.Dims
}

func (a *ATE) String() string {
	return fmt.Sprintf("ATE%+v", *a)
}

// Port is a reserved Port.
type Port struct {
	Name string
}

func (p *Port) String() string {
	return fmt.Sprintf("Port%+v", *p)
}

// Device returns the reserved device with the given ID.
func (r *Reservation) Device(id string) (Device, error) {
	if d, err := r.DUT(id); err == nil { // if NO error
		return d, nil
	}
	if a, err := r.ATE(id); err == nil { // if NO error
		return a, nil
	}
	return nil, errors.Errorf("device ID %s not found in the reservation", id)
}

// DUT returns the reserved DUT with the given device ID.
func (r *Reservation) DUT(id string) (*DUT, error) {
	if d, ok := r.DUTs[id]; ok {
		return d, nil
	}
	return nil, errors.Errorf("DUT ID %s not found in the reservation", id)
}

// ATE returns the reserved ATE with the given device ID.
func (r *Reservation) ATE(id string) (*ATE, error) {
	if a, ok := r.ATEs[id]; ok {
		return a, nil
	}
	return nil, errors.Errorf("ATE ID %s not found in the reservation", id)
}

// Port returns the reserved port with the given ID.
func (d *Dims) Port(id string) (*Port, error) {
	if p, ok := d.Ports[id]; ok {
		return p, nil
	}
	return nil, errors.Errorf("port ID %s not found in reserved device %s", id, d.Name)
}

// InterpolateConfig substitutes templated variables in device config text.
// The following Go template functions are allowed in config:
// - {{ port "<portID>" }}: replaced with the physical port name
// - {{ secrets "<arg1>" "<arg2>" }}: left untouched, returned as-is
func (d *Dims) InterpolateConfig(config string) (string, error) {
	template := template.New(d.Name).
		Funcs(map[string]interface{}{
			// "secrets" function should be a noop
			"secrets": func(secrets ...string) string {
				var args []string
				for _, s := range secrets {
					args = append(args, fmt.Sprintf("%q", s))
				}
				return fmt.Sprintf("{{ secrets %s }}", strings.Join(args, " "))
			},
			"port": func(portID string) (string, error) {
				port, err := d.Port(portID)
				if err != nil {
					return "", usererr.Wrap(err)
				}
				return port.Name, nil
			},
		})
	template, err := template.Parse(config)
	if err != nil {
		return "", usererr.Wrapf(err, "invalid port template in config: %q", config)
	}
	var b strings.Builder
	if err = template.Execute(&b, nil); err != nil {
		return "", usererr.Wrapf(err, "invalid port template in config: %q", config)
	}
	return b.String(), nil
}
