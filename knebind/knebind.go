// Package knebind provides an Ondatra binding for KNE devices.
package knebind

import (
	"golang.org/x/net/context"
	"crypto/tls"
	"fmt"
	"os/exec"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"github.com/pborman/uuid"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	kpb "github.com/google/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

var (
	// TODO: when Ondatra supports the OS dimension, use it to
	// distinguish CSR from CXR and CEVO from VMX.
	type2VendorMap = map[kpb.Node_Type]opb.Device_Vendor{
		kpb.Node_ARISTA_CEOS:  opb.Device_ARISTA,
		kpb.Node_CISCO_CSR:    opb.Device_CISCO,
		kpb.Node_CISCO_CXR:    opb.Device_CISCO,
		kpb.Node_JUNIPER_CEVO: opb.Device_JUNIPER,
		kpb.Node_JUNIPER_VMX:  opb.Device_JUNIPER,
	}

	fetchTopo = fetchTopology // to be stubbed out by tests
)

// Bind implements the ondatra Binding interface for KNE
type Bind struct {
	binding.Binding
	dut2GNMIAddr map[*reservation.DUT]string
	mu           sync.Mutex
	cfg          *Config
}

// New returns a new KNE bind instance.
func New(cfg *Config) (*Bind, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	return &Bind{
		cfg:          cfg,
		dut2GNMIAddr: make(map[*reservation.DUT]string),
	}, nil
}

// Reserve implements the binding Reserve method by finding nodes and links in
// the topology specified in the config file that match the requested testbed.
func (b *Bind) Reserve(ctx context.Context, tb *opb.Testbed, runTime time.Duration, waitTime time.Duration) (*reservation.Reservation, error) {
	topo, err := fetchTopo(b.cfg)
	if err != nil {
		return nil, err
	}
	a, err := solve(tb, topo)
	if err != nil {
		return nil, err
	}
	res := &reservation.Reservation{
		ID:   uuid.New(),
		DUTs: make(map[string]*reservation.DUT),
	}
	for _, dut := range tb.GetDuts() {
		resDUT, err := b.resolveDUT(dut, a)
		if err != nil {
			return nil, err
		}
		res.DUTs[dut.GetId()] = resDUT
	}
	return res, nil
}

func fetchTopology(cfg *Config) (*kpb.Topology, error) {
	args := []string{"topology", "service", cfg.TopoPath}
	if cfg.KubecfgPath != "" {
		args = append(args, fmt.Sprintf("--kubecfg=%s", cfg.KubecfgPath))
	}
	cmd := exec.Command(cfg.CLIPath, args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrapf(err, "error executing command %v", cmd)
	}
	topo := new(kpb.Topology)
	if err := prototext.Unmarshal(out, topo); err != nil {
		return nil, errors.Wrap(err, "error unmarshalling KNE topology proto")
	}
	return topo, nil
}

func (b *Bind) resolveDUT(dev *opb.Device, a *assign) (*reservation.DUT, error) {
	node := a.dut2Node[dev]
	vendor, ok := type2VendorMap[node.GetType()]
	if !ok {
		return nil, errors.Errorf("No known device vendor for node type: %v", node.GetType())
	}
	typeName := kpb.Node_Type_name[int32(node.GetType())]
	dut := &reservation.DUT{&reservation.Dims{
		Name:   node.GetName(),
		Vendor: vendor,
		// TODO: Determine appropriate hardware model and software version
		HardwareModel:   typeName,
		SoftwareVersion: typeName,
		Ports:           make(map[string]*reservation.Port),
	}}
	for _, p := range dev.GetPorts() {
		dut.Ports[p.GetId()] = &reservation.Port{Name: a.port2Intf[p].name}
	}
	ga, err := gnmiAddr(node)
	if err != nil {
		return nil, err
	}
	b.dut2GNMIAddr[dut] = ga
	return dut, nil
}

func gnmiAddr(node *kpb.Node) (string, error) {
	for _, s := range node.GetServices() {
		if s.GetName() == "gnmi" {
			return fmt.Sprintf("%s:%d", s.GetOutsideIp(), s.GetOutside()), nil
		}
	}
	return "", errors.Errorf("No GNMI service found in node: %v", node)
}

// Release is a no-op because there's need to reserve local VMs.
func (b *Bind) Release(_ context.Context) error {
	return nil
}

// SetTestMetadata is unused for KNE.
func (b *Bind) SetTestMetadata(_ *binding.TestMetadata) error {
	return nil
}

func (b *Bind) DialGNMI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	addr := b.dut2GNMIAddr[dut]
	log.Infof("Dialing GNMI dut %s@%s", dut.Name, addr)
	opts = append(opts,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithPerRPCCredentials(&passCred{
			username: b.cfg.Username,
			password: b.cfg.Password,
		}))
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "DialContext(ctx, %s, %v)", addr, opts)
	}
	return gpb.NewGNMIClient(conn), nil
}

type passCred struct {
	username string
	password string
}

func (c *passCred) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": c.username,
		"password": c.password,
	}, nil
}

func (c *passCred) RequireTransportSecurity() bool {
	return true
}
