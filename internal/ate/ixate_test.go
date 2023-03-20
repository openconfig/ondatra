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
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ondatra/internal/ixgnmi"
	"google.golang.org/protobuf/encoding/prototext"

	opb "github.com/openconfig/ondatra/proto"
)

const (
	testDataPath  = "testdata"
	sessionPrefix = "/api/v1/sessions/0"
)

var (
	simpleTop = &Topology{
		Interfaces: []*opb.InterfaceConfig{{
			Name: "intf",
			Link: &opb.InterfaceConfig_Port{"port"},
		}},
	}
)

// updateXPaths sets fake XPaths only for those nodes that require it for config
// generation: vports, lags, and any node that may be a traffic item endpoint.
func updateXPaths(cfg *ixconfig.Ixnetwork) error {
	toXPath := func(format string, args ...any) *ixconfig.XPath {
		xp, err := ixconfig.ParseXPath(fmt.Sprintf(format, args...))
		if err != nil {
			log.Fatalf("Impossible XPath parsing error: %v", err)
		}
		return xp
	}
	for i, p := range cfg.Vport {
		p.Xpath = toXPath("/vport[%d]", i+1)
	}
	for i, l := range cfg.Lag {
		l.Xpath = toXPath("/lag[%d]", i+1)
	}
	for i, t := range cfg.Topology {
		for j, dg := range t.DeviceGroup {
			dg.Xpath = toXPath("/topology[%d]/deviceGroup[%d]", i+1, j+1)
			for _, eth := range dg.Ethernet {
				for _, ipv4 := range eth.Ipv4 {
					for k, bgpv4 := range ipv4.BgpIpv4Peer {
						bgpv4.Active.SingleValue.Xpath = toXPath("/topology[%d]/deviceGroup[%d]/bgpv4Peer[%d]", i+1, j+1, k+1)
					}
				}
				for _, ipv6 := range eth.Ipv6 {
					for k, bgpv6 := range ipv6.BgpIpv6Peer {
						bgpv6.Active.SingleValue.Xpath = toXPath("/topology[%d]/deviceGroup[%d]/bgpv6Peer[%d]", i+1, j+1, k+1)
					}
				}
			}
			for k, ng := range dg.NetworkGroup {
				ng.Xpath = toXPath("/topology[%d]/deviceGroup[%d]/networkGroup[%d]", i+1, j+1, k+1)
				for l, ndg := range ng.DeviceGroup {
					for m, lb := range ndg.Ipv4Loopback {
						for n, lsps := range lb.RsvpteLsps {
							baseRsvpPath := fmt.Sprintf("/topology[%d]/deviceGroup[%d]/networkGroup[%d]/deviceGroup[%d]/ipv4Loopback[%d]/rsvpteLsps[%d]/", i+1, j+1, k+1, l+1, m+1, n+1)
							lsps.RsvpP2PEgressLsps.Xpath = toXPath(path.Join(baseRsvpPath, "rsvpP2PEgressLsps"))
							lsps.RsvpP2PIngressLsps.Xpath = toXPath(path.Join(baseRsvpPath, "rsvpP2PIngressLsps"))
						}
					}
				}
			}
		}
	}
	return nil
}

type fakeCfgClient struct {
	cfgClient
	xPathToID     map[string]string
	updateID      func(ixconfig.IxiaCfgNode) (string, string)
	updateIDErr   error
	importErrs    []error
	lastImportCfg ixconfig.IxiaCfgNode
	session       *fakeSession
}

func (c *fakeCfgClient) Session() session {
	return c.session
}

func (c *fakeCfgClient) NodeID(node ixconfig.IxiaCfgNode) (string, error) {
	xpObj := node.XPath()
	if xpObj == nil {
		return "", fmt.Errorf("no stub XPath for node of type %T", node)
	}
	xp := xpObj.String()
	id, ok := c.xPathToID[xp]
	if !ok {
		return "", fmt.Errorf("no stub ID for node at XPath %q", xp)
	}
	return id, nil
}

func (c *fakeCfgClient) ImportConfig(ctx context.Context, cfg *ixconfig.Ixnetwork, node ixconfig.IxiaCfgNode, _ bool) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	importErr := c.importErrs[0]
	c.importErrs = c.importErrs[1:]
	if importErr != nil {
		return importErr
	}
	var err error
	c.lastImportCfg, err = cfg.ResolvedConfig(node)
	if err != nil {
		return fmt.Errorf("Error resolving xpaths/refs for config node: %w", err)
	}
	return updateXPaths(cfg)
}

type fakeDelayImportClient struct {
	cfgClient
	delays    int
	importErr error
}

func (c *fakeDelayImportClient) ImportConfig(ctx context.Context, _ *ixconfig.Ixnetwork, _ ixconfig.IxiaCfgNode, _ bool) error {
	if c.delays == 0 {
		return c.importErr
	}
	c.delays--
	<-ctx.Done()
	return ctx.Err()
}

func (c *fakeDelayImportClient) Session() session {
	return &fakeSession{}
}

func (c *fakeCfgClient) UpdateIDs(_ context.Context, cfg *ixconfig.Ixnetwork, cns ...ixconfig.IxiaCfgNode) error {
	if err := updateXPaths(cfg); err != nil {
		return err
	}
	return c.updateIDErr
}

type fakeSession struct {
	session
	deleteErrs map[string]error
	getRsps    map[string]string
	getErrs    map[string]error
	patchErrs  map[string]error
	postRsps   map[string]string
	postErrs   map[string]error
	files      *fakeFiles
	stats      *fakeStats
}

func (s *fakeSession) AbsPath(relPath string) string {
	return relPath
}

func (s *fakeSession) Delete(_ context.Context, p string) error {
	return s.deleteErrs[p]
}

func (s *fakeSession) Get(_ context.Context, p string, v any) error {
	if s.getRsps[p] != "" {
		json.Unmarshal([]byte(s.getRsps[p]), v)
	}
	return s.getErrs[p]
}

func (s *fakeSession) Patch(_ context.Context, p string, _ any) error {
	return s.patchErrs[p]
}

func (s *fakeSession) Post(_ context.Context, p string, _, out any) error {
	if s.postRsps[p] != "" && out != nil {
		if err := json.Unmarshal([]byte(s.postRsps[p]), out); err != nil {
			return err
		}
	}
	return s.postErrs[p]
}

func (s *fakeSession) Errors(_ context.Context) ([]*ixweb.Error, error) {
	return nil, nil
}

func (s *fakeSession) Files() files {
	return s.files
}

func (s *fakeSession) Stats() stats {
	return s.stats
}

type fakeFiles struct {
	files
	listRes   []string
	listErr   error
	uploadErr error
	deleteErr error
}

func (f *fakeFiles) List(context.Context, string) ([]string, error) {
	return f.listRes, f.listErr
}

func (f *fakeFiles) Upload(context.Context, string, []byte) error {
	return f.uploadErr
}

func (f *fakeFiles) Delete(context.Context, string) error {
	return f.deleteErr
}

type fakeStats struct {
	stats
	viewsOut  map[string]view
	viewsErr  error
	egressErr error
}

func (s *fakeStats) Views(context.Context) (map[string]view, error) {
	return s.viewsOut, s.viewsErr
}

func (s *fakeStats) ConfigEgressView(context.Context, []string, int) (*ixweb.StatView, error) {
	return nil, s.egressErr
}

type fakeView struct {
	tableOut ixweb.StatTable
	tableErr error
}

func (v *fakeView) FetchTable(context.Context, ...string) (ixweb.StatTable, error) {
	return v.tableOut, v.tableErr
}

func jsonCfgDiff(t *testing.T, wantCfg, gotCfg ixconfig.IxiaCfgNode) string {
	wantBytes, err := json.Marshal(wantCfg)
	if err != nil {
		t.Fatalf("Could not marshal expected config to JSON: %v", err)
	}
	gotBytes, err := json.Marshal(gotCfg)
	if err != nil {
		t.Fatalf("Could not marshal actual config to JSON: %v", err)
	}

	var want, got map[string]any
	if err := json.Unmarshal(wantBytes, &want); err != nil {
		t.Fatalf("Could not unmarshal expected config to a map: %v", err)
	}
	if err := json.Unmarshal(gotBytes, &got); err != nil {
		t.Fatalf("Could not unmarshal actual config to a map: %v", err)
	}
	return cmp.Diff(want, got)
}

func toFlows(t *testing.T, filename string) []*opb.Flow {
	b, err := os.ReadFile(filepath.Join(testDataPath, filename))
	if err != nil {
		t.Fatalf("Could not read file %q: %v", filename, err)
	}
	traf := &opb.Traffic{}
	if err := prototext.Unmarshal(b, traf); err != nil {
		t.Fatalf("Could not unmarshal proto from file %q: %v", filename, err)
	}
	return traf.GetFlows()
}

func toCfg(t *testing.T, filename string) *ixconfig.Ixnetwork {
	b, err := os.ReadFile(filepath.Join(testDataPath, filename))
	if err != nil {
		t.Fatalf("Could not read file %q: %v", filename, err)
	}
	cfg := &ixconfig.Ixnetwork{}
	if err := json.Unmarshal(b, cfg); err != nil {
		t.Fatalf("Could not unmarshal IxNetwork config from file %q: %v", filename, err)
	}
	return cfg
}

func toTrafficCfg(t *testing.T, filename string) *ixconfig.Traffic {
	b, err := os.ReadFile(filepath.Join(testDataPath, filename))
	if err != nil {
		t.Fatalf("Could not read file %q: %v", filename, err)
	}
	cfg := &ixconfig.Traffic{}
	if err := json.Unmarshal(b, cfg); err != nil {
		t.Fatalf("Could not unmarshal IxNetwork config from file %q: %v", filename, err)
	}
	return cfg
}

func stubLogOperationResult() {
	// TODO(team): Implement operation logging for open-source Ondatra.
}

func restoreStubs() {
	resolveMacsFn = resolveMacs
	sleepFn = time.Sleep
	syncRouteTableFilesAndImportFn = syncRouteTableFilesAndImport
	validateProtocolStartFn = validateProtocolStart
	resetIxiaTrafficCfgFn = resetIxiaTrafficCfg
	genTrafficFn = genTraffic
	updateFlowsFn = updateFlows
	startTrafficFn = startTraffic
}

func TestImportConfig(t *testing.T) {
	defer restoreStubs()
	stubLogOperationResult()
	sleepFn = func(time.Duration) {}
	tests := []struct {
		desc      string
		timeout   bool
		delays    int
		importErr error
		wantErr   string
	}{{
		desc:      "non-timeout error on import",
		importErr: errors.New("some import error"),
		wantErr:   "some import error",
	}, {
		desc:    "timeout on all config import attempts",
		timeout: true,
		delays:  importRetries,
		wantErr: "timeout importing",
	}, {
		desc:    "timeout on initial config import",
		timeout: true,
		delays:  1,
	}, {
		desc: "successful import",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			importTimeout := time.Minute
			if test.timeout {
				importTimeout = 0
			}
			c := &ixATE{
				c: &fakeDelayImportClient{
					delays:    test.delays,
					importErr: test.importErr,
				},
				cfg: &ixconfig.Ixnetwork{},
			}
			gotErr := c.importConfig(context.Background(), c.cfg, false, importTimeout, "import test")
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("importConfig: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestPushTopology(t *testing.T) {
	const (
		ixiaName    = "ixia1"
		chassisHost = "192.168.1.1"
	)
	defer restoreStubs()
	stubLogOperationResult()
	sleepFn = func(time.Duration) {}
	resetIxiaTrafficCfgFn = func(context.Context, *ixATE) error {
		return nil
	}

	tests := []struct {
		desc, reqFile, wantCfgFile string
		top                        *Topology
		importErrs                 []error
		routeTableImportErr        error
		opErr                      error
		wantErr                    string
	}{{
		desc:       "Error if no interfaces",
		top:        &Topology{},
		importErrs: []error{nil, nil},
		wantErr:    "zero interfaces",
	}, {
		desc: "Error if too many interfaces",
		top: func() *Topology {
			top := &Topology{}
			for i := 0; i <= maxIntfs; i++ {
				top.Interfaces = append(top.Interfaces, &opb.InterfaceConfig{})
			}
			return top
		}(),
		importErrs: []error{nil, nil},
		wantErr:    "must be at most",
	}, {
		desc:       "Error on port config push",
		top:        simpleTop,
		importErrs: []error{errors.New("ports push err")},
		wantErr:    "ports push",
	}, {
		desc:       "Error on topo config push",
		top:        simpleTop,
		importErrs: []error{nil, errors.New("topo push err")},
		wantErr:    "topo push",
	}, {
		desc:                "Route table import err",
		top:                 simpleTop,
		importErrs:          []error{nil, nil},
		routeTableImportErr: errors.New("failed route table import"),
		wantErr:             "failed route",
	}, {
		desc: "IS-IS config with no traffic",
		top: &Topology{
			Interfaces: []*opb.InterfaceConfig{{
				Name: "intf",
				Link: &opb.InterfaceConfig_Port{"12/1"},
				Ethernet: &opb.EthernetConfig{
					Mtu: 9000,
					Fec: &opb.Fec{Enabled: true},
				},
				Ipv4: &opb.IpConfig{
					AddressCidr:    "192.168.31.2/30",
					DefaultGateway: "192.168.31.1",
				},
				Isis: &opb.ISISConfig{
					Level:              opb.ISISConfig_L2,
					Metric:             10,
					AreaId:             "490001",
					EnableWideMetric:   true,
					NetworkType:        opb.ISISConfig_POINT_TO_POINT,
					AuthType:           opb.ISISConfig_MD5,
					AuthKey:            "md5password",
					EnableHelloPadding: true,
					TeRouterId:         "0.0.0.0",
					CapabilityRouterId: "0.0.0.1",
					HelloIntervalSec:   10,
					DeadIntervalSec:    30,
				},
				Networks: []*opb.Network{{
					Name:          "net1",
					InterfaceName: "intf",
					Ipv4: &opb.NetworkIp{
						AddressCidr: "30.0.0.0/30",
						Count:       15000,
					},
					Isis: &opb.IPReachability{
						Metric:      10,
						RouteOrigin: opb.IPReachability_INTERNAL,
					},
				}, {
					Name:          "net2",
					InterfaceName: "intf",
					Ipv4: &opb.NetworkIp{
						AddressCidr: "40.0.0.0/30",
						Count:       35000,
					},
					Isis: &opb.IPReachability{
						Metric:      10,
						RouteOrigin: opb.IPReachability_INTERNAL,
					},
				}, {
					Name:          "net3",
					InterfaceName: "intf",
					Ipv6: &opb.NetworkIp{
						AddressCidr: "2001:4860:0:1::444:1/126",
						Count:       15000,
					},
					Isis: &opb.IPReachability{
						Metric:      10,
						RouteOrigin: opb.IPReachability_INTERNAL,
					},
				}},
			}},
		},
		importErrs:  []error{nil, nil},
		wantCfgFile: "isis_no_traffic_cfg.json",
	}, {
		desc: "FEC disabled",
		top: &Topology{
			Interfaces: []*opb.InterfaceConfig{{
				Name: "intf",
				Link: &opb.InterfaceConfig_Port{"12/1"},
				Ethernet: &opb.EthernetConfig{
					Mtu: 1500,
					Fec: &opb.Fec{Enabled: false},
				},
			}},
		},
		importErrs:  []error{nil, nil},
		wantCfgFile: "fec_disabled.json",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fc := &fakeCfgClient{
				importErrs: test.importErrs,
			}
			c := &ixATE{
				c:           fc,
				name:        ixiaName,
				chassisHost: chassisHost,
			}
			syncRouteTableFilesAndImportFn = func(context.Context, *ixATE) error {
				return test.routeTableImportErr
			}

			gotErr := c.PushTopology(context.Background(), test.top)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("PushTopology: unexpected error result, got err: %v, want %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			wantCfg := toCfg(t, test.wantCfgFile)
			if diff := jsonCfgDiff(t, wantCfg, fc.lastImportCfg); diff != "" {
				t.Errorf("PushTopology: Unexpected topology config pushed, diff (-want, +got)\n%s", diff)
			}
		})
	}
}

// Config generation is tested in TestPushTopology, this test is only for Ixnetwork interactions.
func TestUpdateTopology(t *testing.T) {
	const vportID = "/fake/vport/1"
	vportXP := parseXPath(t, "/fake/vport/1")

	tests := []struct {
		desc                 string
		vportState           string
		top                  *Topology
		operState            operState
		importErr            error
		routeTableImportErr  error
		startProtocolsErr    error
		validateProtocolsErr error
		genErr               error
		startErr             error
		wantErr              string
	}{{
		desc:    "error if no interfaces",
		top:     &Topology{},
		wantErr: "zero interfaces",
	}, {
		desc:      "error on config push",
		top:       simpleTop,
		importErr: errors.New("push error"),
		wantErr:   "push error",
	}, {
		desc:                "error on route table import",
		top:                 simpleTop,
		routeTableImportErr: errors.New("route table import error"),
		wantErr:             "route table import error",
	}, {
		desc: "traffic/protocols not yet running",
		top:  simpleTop,
	}, {
		desc:       "error vport down",
		top:        simpleTop,
		operState:  operStateProtocolsOn,
		vportState: "connectedLinkDown",
		wantErr:    "connectedLinkDown",
	}, {
		desc:                 "error starting protocols",
		top:                  simpleTop,
		operState:            operStateProtocolsOn,
		vportState:           "connectedLinkUp",
		startProtocolsErr:    errors.New("could not start protocols"),
		validateProtocolsErr: errors.New("protocols not up"),
		// TODO(team); Revert to checking start protocols error (see comment in 'startProtocols' section.)
		wantErr: "protocols not up",
	}, {
		desc:                 "error waiting for protocols",
		top:                  simpleTop,
		operState:            operStateProtocolsOn,
		vportState:           "connectedLinkUp",
		validateProtocolsErr: errors.New("protocols not up"),
		wantErr:              "protocols not up",
	}, {
		desc:       "error starting traffic",
		top:        simpleTop,
		operState:  operStateTrafficOn,
		vportState: "connectedLinkUp",
		genErr:     errors.New("could not generate traffic"),
		wantErr:    "could not generate traffic",
	}, {
		desc:       "error starting traffic",
		top:        simpleTop,
		operState:  operStateTrafficOn,
		vportState: "connectedLinkUp",
		startErr:   errors.New("could not start traffic"),
		wantErr:    "could not start traffic",
	}, {
		desc:       "successful protocol/traffic start",
		top:        simpleTop,
		operState:  operStateTrafficOn,
		vportState: "connectedLinkUp",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			fc := &fakeCfgClient{
				importErrs: []error{test.importErr},
				session: &fakeSession{
					postErrs: map[string]error{"/operations/startallprotocols": test.startProtocolsErr},
					getRsps:  map[string]string{vportID: fmt.Sprintf("{\"connectionState\": \"%s\"}", test.vportState)},
				},
				xPathToID: map[string]string{vportXP.String(): vportID},
			}
			c := &ixATE{
				c:         fc,
				cfg:       &ixconfig.Ixnetwork{Traffic: &ixconfig.Traffic{}},
				operState: test.operState,
				ports: map[string]*ixconfig.Vport{
					"port": &ixconfig.Vport{
						Name:     ixconfig.String("1/2"),
						Xpath:    vportXP,
						L1Config: &ixconfig.VportL1Config{},
					},
				},
				intfs: make(map[string]*intf),
			}
			syncRouteTableFilesAndImportFn = func(context.Context, *ixATE) error {
				return test.routeTableImportErr
			}
			validateProtocolStartFn = func(context.Context, *ixATE) ([]string, error) {
				return nil, test.validateProtocolsErr
			}
			genTrafficFn = func(context.Context, *ixATE) error {
				return test.genErr
			}
			startTrafficFn = func(context.Context, *ixATE) error {
				return test.startErr
			}
			gotErr := c.UpdateTopology(context.Background(), test.top)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("UpdateTopology: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestSyncRouteTableFilesAndImport(t *testing.T) {
	const (
		net            = "someNet"
		localFile      = "juniper_route_table"
		localFileCrc32 = "3163592075"
		v4ImportOp     = "topology/deviceGroup/networkGroup/ipv4PrefixPools/bgpIPRouteProperty/operations/importbgproutes"
		v6ImportOp     = "topology/deviceGroup/networkGroup/ipv6PrefixPools/bgpV6IPRouteProperty/operations/importbgproutes"
	)
	rtFile := filepath.Join(testDataPath, localFile)
	bgp4XP := parseXPath(t, "/fake/xpath/bgp4")
	bgp6XP := parseXPath(t, "/fake/xpath/bgp6")
	baseClient := func() *ixATE {
		ng := &ixconfig.TopologyNetworkGroup{
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{Xpath: bgp4XP}},
			}},
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{Xpath: bgp6XP}},
			}},
		}
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					NetworkGroup: []*ixconfig.TopologyNetworkGroup{ng},
				}},
			}},
		}
		return &ixATE{
			cfg: cfg,
			intfs: map[string]*intf{
				"someIntf": &intf{
					netToNetworkGroup: map[string]*ixconfig.TopologyNetworkGroup{net: ng},
					netToRouteTables: map[string]*routeTables{
						net: &routeTables{
							format: routeTableFormatJuniper,
							ipv4:   rtFile,
							ipv6:   rtFile,
						},
					},
				},
			},
		}
	}
	tests := []struct {
		desc         string
		syncedFiles  map[string]string
		fileInList   string
		listErr      error
		uploadErr    error
		deleteErr    error
		updateIDsErr error
		v4ImportErr  error
		v6ImportErr  error
		wantErr      string
	}{{
		desc:        "files already synced",
		syncedFiles: map[string]string{rtFile: "existing_file"},
	}, {
		desc:        "could not list files",
		syncedFiles: map[string]string{},
		listErr:     errors.New("error listing chassis files"),
		wantErr:     "error listing chassis files",
	}, {
		desc:        "bad format for route table file",
		syncedFiles: map[string]string{},
		fileInList:  "badfile",
		wantErr:     "unexpected format",
	}, {
		desc:        "bad timestamp in route table file",
		syncedFiles: map[string]string{},
		fileInList:  "ondatra_route_table_zzz_aaa",
		wantErr:     "unexpected timestamp format",
	}, {
		desc:        "error deleting old file",
		syncedFiles: map[string]string{},
		fileInList:  fmt.Sprintf("ondatra_route_table_0_%s", localFileCrc32),
		deleteErr:   errors.New("error deleting file"),
		wantErr:     "error deleting file",
	}, {
		desc:        "error uploading new file",
		syncedFiles: map[string]string{},
		fileInList:  fmt.Sprintf("ondatra_route_table_0_%s", localFileCrc32),
		uploadErr:   errors.New("error uploading file"),
		wantErr:     "error uploading file",
	}, {
		desc:         "error updating object IDs",
		syncedFiles:  map[string]string{},
		fileInList:   fmt.Sprintf("ondatra_route_table_0_%s", localFileCrc32),
		updateIDsErr: errors.New("error updating IDs"),
		wantErr:      "error updating IDs",
	}, {
		desc:        "error on V4 route import",
		syncedFiles: map[string]string{},
		fileInList:  fmt.Sprintf("ondatra_route_table_0_%s", localFileCrc32),
		v4ImportErr: errors.New("failed to import v4 routes"),
		wantErr:     "failed to import v4 routes",
	}, {
		desc:        "error on V6 route import",
		syncedFiles: map[string]string{},
		fileInList:  fmt.Sprintf("ondatra_route_table_0_%s", localFileCrc32),
		v6ImportErr: errors.New("failed to import v6 routes"),
		wantErr:     "failed to import v6 routes",
	}, {
		desc:        "file already uploaded",
		syncedFiles: map[string]string{},
		fileInList:  fmt.Sprintf("ondatra_route_table_%d_%s", time.Now().Unix(), localFileCrc32),
	}, {
		desc:        "successful delete, file successfully uploaded",
		syncedFiles: map[string]string{},
		fileInList:  fmt.Sprintf("ondatra_route_table_0_%s", localFileCrc32),
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			sleepFn = func(time.Duration) {}

			c := baseClient()
			c.routeTableToIxFile = test.syncedFiles
			c.c = &fakeCfgClient{
				xPathToID: map[string]string{
					bgp4XP.String(): "/id/to/bgpv4",
					bgp6XP.String(): "/id/to/bgpv6",
				},
				updateIDErr: test.updateIDsErr,
				session: &fakeSession{
					postErrs: map[string]error{
						v4ImportOp: test.v4ImportErr,
						v6ImportOp: test.v6ImportErr,
					},
					files: &fakeFiles{
						listRes:   []string{test.fileInList},
						listErr:   test.listErr,
						uploadErr: test.uploadErr,
						deleteErr: test.deleteErr,
					},
				},
			}
			gotErr := syncRouteTableFilesAndImport(context.Background(), c)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("syncRouteTableFilesAndImport: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestValidateProtocolStart(t *testing.T) {
	const (
		ethID  = "id/to/eth"
		ipv4ID = "id/to/ipv4"
		ipv6ID = "id/to/ipv6"
		lspID  = "id/tp/lsp"
	)
	ethXP := parseXPath(t, "/fake/xpath/eth")
	ipv4XP := parseXPath(t, "/fake/xpath/ipv4")
	ipv6XP := parseXPath(t, "/fake/xpath/ipv6")
	lspXP := parseXPath(t, "/fake/xpath/lsp")
	baseClient := func(withIsisAndLAG bool) *ixATE {
		lsp := &ixconfig.TopologyRsvpteLsps{
			RsvpP2PIngressLsps: &ixconfig.TopologyRsvpP2PIngressLsps{
				Xpath:  lspXP,
				Active: ixconfig.MultivalueTrue(),
			},
		}
		dg := &ixconfig.TopologyDeviceGroup{
			Ethernet: []*ixconfig.TopologyEthernet{{
				Xpath: ethXP,
				Mac:   &ixconfig.Multivalue{},
				Ipv4: []*ixconfig.TopologyIpv4{{
					Xpath:      ipv4XP,
					RsvpteLsps: []*ixconfig.TopologyRsvpteLsps{lsp},
				}},
				Ipv6: []*ixconfig.TopologyIpv6{{Xpath: ipv6XP}},
			}},
		}
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{dg},
			}},
			Lag: []*ixconfig.Lag{{}},
			Vport: []*ixconfig.Vport{
				{Name: ixconfig.String("1/1")},
				{Name: ixconfig.String("2/2")},
			},
		}
		updateXPaths(cfg) // Only needed for IS-IS/LAG test case

		var isr map[string]*ixconfig.TopologyNetworkGroup
		var link ixconfig.IxiaCfgNode
		link = cfg.Vport[0]
		if withIsisAndLAG {
			link = cfg.Lag[0]
			isr = map[string]*ixconfig.TopologyNetworkGroup{"isr": nil}
		}

		return &ixATE{
			cfg: cfg,
			intfs: map[string]*intf{
				"someIntf": &intf{
					deviceGroup:       dg,
					rsvpLSPs:          map[string]*ixconfig.TopologyRsvpteLsps{"lspName": lsp},
					link:              link,
					isrToNetworkGroup: isr,
				},
			},
			lagPorts: map[*ixconfig.Lag][]*ixconfig.Vport{
				cfg.Lag[0]: []*ixconfig.Vport{cfg.Vport[0], cfg.Vport[1]},
			},
		}
	}
	tests := []struct {
		desc           string
		withIsisAndLAG bool
		topoStatus     string
		topoErr        error
		errorsResult   string
		idErr          error
		protocolStatus string
		protocolErr    error
		lspStatus      string
		lspErr         error
		ethRestartErr  error
		wantErr        string
		wantPorts      []string
	}{{
		desc:    "error fetching protocol status",
		topoErr: errors.New("error on querying global protocol status"),
		wantErr: "could not fetch global topology",
	}, {
		desc:       "protocols never started",
		topoStatus: "notStarted",
		wantErr:    "protocols did not start in time",
	}, {
		desc:       "error updating protocol IDs",
		topoStatus: "started",
		idErr:      errors.New("error updating IDs"),
		wantErr:    "could not update IDs",
	}, {
		desc:        "error fetching individual protocol statuses",
		topoStatus:  "started",
		protocolErr: errors.New("error querying protocol"),
		wantErr:     "could not fetch element",
	}, {
		desc:           "individual protocols never start",
		topoStatus:     "started",
		protocolStatus: "down",
		wantErr:        "some protocol instances did not start",
		wantPorts:      []string{"1/1"},
	}, {
		desc:           "individual protocols never start on LAG interface",
		withIsisAndLAG: true,
		topoStatus:     "started",
		protocolStatus: "down",
		wantErr:        "some protocol instances did not start",
		wantPorts:      []string{"1/1", "2/2"},
	}, {
		desc:           "error restarting eth protocol",
		topoStatus:     "started",
		protocolStatus: "down",
		ethRestartErr:  errors.New("error restarting protocol"),
		wantErr:        "could not restart down protocol",
	}, {
		desc:           "error fetching ingress LSP status",
		topoStatus:     "started",
		protocolStatus: "up",
		lspErr:         errors.New("error querying ingress LSP"),
		wantErr:        "could not fetch element",
	}, {
		desc:           "ingress LSP never starts",
		topoStatus:     "started",
		protocolStatus: "up",
		lspStatus:      "down",
		wantErr:        "some ingress LSP instances did not start",
		wantPorts:      []string{"1/1"},
	}, {
		desc:           "all up",
		topoStatus:     "started",
		protocolStatus: "up",
		lspStatus:      "up",
	}, {
		desc:           "index out of bounds with IS-IS and LAG",
		withIsisAndLAG: true,
		topoStatus:     "notStarted",
		errorsResult:   `[{"Name": "Index was outside the bounds of the array"}]`,
		protocolStatus: "up",
		lspStatus:      "up",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			sleepFn = func(time.Duration) {}

			c := baseClient(test.withIsisAndLAG)
			c.c = &fakeCfgClient{
				xPathToID: map[string]string{
					ethXP.String():  ethID,
					ipv4XP.String(): ipv4ID,
					ipv6XP.String(): ipv6ID,
					lspXP.String():  lspID,
				},
				updateIDErr: test.idErr,
				session: &fakeSession{
					getRsps: map[string]string{
						"globals/topology":        fmt.Sprintf(`{"status": "%s"}`, test.topoStatus),
						"globals/appErrors/error": test.errorsResult,
						ethID:                     fmt.Sprintf(`{"sessionStatus": ["%s"]}`, test.protocolStatus),
						ipv4ID:                    fmt.Sprintf(`{"sessionStatus": ["%s"]}`, test.protocolStatus),
						ipv6ID:                    fmt.Sprintf(`{"sessionStatus": ["%s"]}`, test.protocolStatus),
						lspID:                     fmt.Sprintf(`{"state": ["%s"]}`, test.lspStatus),
					},
					getErrs: map[string]error{
						"globals/topology": test.topoErr,
						ethID:              test.protocolErr,
						ipv4ID:             test.protocolErr,
						ipv6ID:             test.protocolErr,
						lspID:              test.lspErr,
					},
					postErrs: map[string]error{
						"topology/deviceGroup/ethernet/operations/restartdown": test.ethRestartErr,
					},
				},
			}
			gotPorts, gotErr := validateProtocolStart(context.Background(), c)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("validateProtocolStart: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.wantPorts, gotPorts, cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("validateProtocolStart: unexpected difference in failed port set (-want, +got): %v", diff)
			}
		})
	}
}

func TestStartProtocols(t *testing.T) {
	const vportID = "/fake/vport/1"
	vportXP := parseXPath(t, "/fake/vport/1")

	tests := []struct {
		desc, vportState, wantErr string
		opErr, protocolsErr       error
	}{{
		desc:       "error vport down",
		vportState: "connectedLinkDown",
		wantErr:    "connectedLinkDown",
	}, {
		// Currently a failure is only reported if protocols are not up, even if the operation failed.
		// TODO(team): Revert to checking that an error is produced(see comment in 'startProtocols' section.)
		desc:       "error from op",
		vportState: "connectedLinkUp",
		opErr:      errors.New("someError"),
	}, {
		desc:         "error waiting for protocol status validation",
		vportState:   "connectedLinkUp",
		protocolsErr: errors.New("protocols not up"),
		wantErr:      "not up",
	}, {
		// Errors from 'startallprotcols' op should be reported if protocols fail to come up.
		desc:         "error on op and error on validation",
		vportState:   "connectedLinkUp",
		opErr:        errors.New("someError"),
		protocolsErr: errors.New("protocols not up"),
		wantErr:      "someError",
	}, {
		desc:       "no error",
		vportState: "connectedLinkUp",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			c := &ixATE{
				c: &fakeCfgClient{
					session: &fakeSession{
						postErrs: map[string]error{"operations/startallprotocols": test.opErr},
						getRsps:  map[string]string{vportID: fmt.Sprintf("{\"connectionState\": \"%s\"}", test.vportState)},
					},
					xPathToID: map[string]string{vportXP.String(): vportID},
				},
				cfg: &ixconfig.Ixnetwork{},
				ports: map[string]*ixconfig.Vport{
					"port1": &ixconfig.Vport{Xpath: vportXP},
				},
			}
			validateProtocolStartFn = func(context.Context, *ixATE) ([]string, error) {
				return nil, test.protocolsErr
			}

			gotErr := c.StartProtocols(context.Background())
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("StartProtocols: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestStopProtocols(t *testing.T) {
	const op = "operations/stopallprotocols"
	tests := []struct {
		desc      string
		operState operState
		opErr     error
		wantErr   bool
	}{{
		desc: "protocols not started",
	}, {
		desc:      "Error from op",
		operState: operStateProtocolsOn,
		opErr:     errors.New("someError"),
		wantErr:   true,
	}, {
		desc:      "No error",
		operState: operStateProtocolsOn,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stubLogOperationResult()
			c := &ixATE{
				c: &fakeCfgClient{
					session: &fakeSession{postErrs: map[string]error{op: test.opErr}},
				},
				operState: test.operState,
			}
			gotErr := c.StopProtocols(context.Background())
			if (gotErr != nil) != test.wantErr {
				t.Errorf("StopProtocols: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}
		})
	}
}

func TestSetPortState(t *testing.T) {
	const port = "1/1"
	cfg := &ixconfig.Ixnetwork{
		Vport: []*ixconfig.Vport{{Name: ixconfig.String(port)}},
	}

	tests := []struct {
		desc         string
		port         string
		enabled      bool
		idErr, opErr error
		wantErr      string
	}{{
		desc:    "invalid port",
		port:    "2/2",
		wantErr: "does not exist in current config",
	}, {
		desc:    "error updating ID",
		port:    port,
		idErr:   errors.New("update ID error"),
		wantErr: "could not fetch ID",
	}, {
		desc:    "error setting link up/down",
		port:    port,
		opErr:   errors.New("link up/down error"),
		wantErr: "error setting port state",
	}, {
		desc:    "successfully set port state up",
		port:    port,
		enabled: true,
	}, {
		desc:    "successfully set port state down",
		port:    port,
		enabled: false,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stubLogOperationResult()
			c := &ixATE{
				cfg:   cfg,
				ports: map[string]*ixconfig.Vport{port: cfg.Vport[0]},
				c: &fakeCfgClient{
					session: &fakeSession{
						postErrs: map[string]error{"vport/operations/linkupdn": test.opErr},
					},
					xPathToID:   map[string]string{"/vport[1]": "/id/to/vport"},
					updateIDErr: test.idErr,
				},
			}

			gotErr := c.SetPortState(context.Background(), test.port, &test.enabled)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("SetPortState: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestSetLACPState(t *testing.T) {
	const (
		port1 = "1/1"
		port2 = "1/2"
	)
	lacpXP := parseXPath(t, "/fake/xpath/lacp")
	lag := &ixconfig.Lag{
		ProtocolStack: &ixconfig.LagProtocolStack{
			Ethernet: []*ixconfig.LagEthernet{{
				Lagportlacp: []*ixconfig.LagLagportlacp{{
					Xpath: lacpXP,
				}},
			}},
		}}
	cfg := &ixconfig.Ixnetwork{
		Vport: []*ixconfig.Vport{
			{Name: ixconfig.String(port1)},
			{Name: ixconfig.String(port2)},
		},
	}

	tests := []struct {
		desc         string
		port         string
		enabled      bool
		idErr, opErr error
		wantErr      string
	}{{
		desc:    "port not in config",
		port:    "2/2",
		wantErr: "does not exist in current config",
	}, {
		desc:    "port not in lag",
		port:    port2,
		wantErr: "not a member of a LAG",
	}, {
		desc:    "error updating ID",
		port:    port1,
		idErr:   errors.New("update ID error"),
		wantErr: "could not fetch ID",
	}, {
		desc:    "error setting LACP state",
		port:    port1,
		opErr:   errors.New("lacp state error"),
		wantErr: "lacp state error",
	}, {
		desc:    "successfully set lacp state on",
		port:    port1,
		enabled: true,
	}, {
		desc:    "successfully set lacp state off",
		port:    port1,
		enabled: false,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stubLogOperationResult()
			c := &ixATE{
				cfg: cfg,
				ports: map[string]*ixconfig.Vport{
					port1: cfg.Vport[0],
					port2: cfg.Vport[1],
				},
				lagPorts: map[*ixconfig.Lag][]*ixconfig.Vport{
					lag: {cfg.Vport[0]},
				},
				c: &fakeCfgClient{
					session: &fakeSession{
						postErrs: map[string]error{
							startLACPOp: test.opErr,
							stopLACPOp:  test.opErr,
						},
					},
					xPathToID:   map[string]string{lacpXP.String(): "/id/to/lacp"},
					updateIDErr: test.idErr,
				},
			}

			gotErr := c.SetLACPState(context.Background(), test.port, &test.enabled)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("SetLACPState: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestResolveMacs(t *testing.T) {
	const (
		intfName = "someIntf"
		ethMac   = "11:11:11:11:11:11"
		ipv4Mac  = "22:22:22:22:22:22"
		ipv6Mac  = "33:33:33:33:33:33"
		valuesOp = "multivalue/operations/getvalues"
		ipv4ID   = "id/to/ipv4"
		ipv6ID   = "id/to/ipv6"
	)
	macXP := parseXPath(t, "/fake/xpath/mac")
	ipv4XP := parseXPath(t, "/fake/xpath/ipv4")
	ipv6XP := parseXPath(t, "/fake/xpath/ipv6")
	cfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Ethernet: []*ixconfig.TopologyEthernet{{
					Mac:  &ixconfig.Multivalue{Xpath: macXP},
					Ipv4: []*ixconfig.TopologyIpv4{{Xpath: ipv4XP}},
					Ipv6: []*ixconfig.TopologyIpv6{{Xpath: ipv6XP}},
				}},
			}},
		}},
	}
	dg := cfg.Topology[0].DeviceGroup[0]
	eth := dg.Ethernet[0]

	tests := []struct {
		desc                                 string
		ethMacRsp, ipv4MacRsp, ipv6MacRsp    string
		idErr, ethErr, ipv4Err, ipv6Err      error
		wantErr                              bool
		wantEthMac, wantIpv4Mac, wantIpv6Mac string
	}{{
		desc:    "failed ID update",
		idErr:   errors.New("failed to update IDs"),
		wantErr: true,
	}, {
		desc:    "failed eth mac resolution",
		ethErr:  errors.New("failed eth mac multivalue fetch"),
		wantErr: true,
	}, {
		desc:    "failed IPv4 mac resolution",
		ipv4Err: errors.New("failed IPv4 mac resolution"),
		wantErr: true,
	}, {
		desc:    "failed IPv6 mac resolution",
		ipv6Err: errors.New("failed IPv6 mac resolution"),
		wantErr: true,
	}, {
		desc:        "ignore invalid mac",
		ethMacRsp:   ethMac,
		ipv4MacRsp:  "removePacket[Unresolved]",
		ipv6MacRsp:  ipv6Mac,
		wantEthMac:  ethMac,
		wantIpv6Mac: ipv6Mac,
	}, {
		desc:        "successful mac resolution",
		ethMacRsp:   ethMac,
		ipv4MacRsp:  ipv4Mac,
		ipv6MacRsp:  ipv6Mac,
		wantEthMac:  ethMac,
		wantIpv4Mac: ipv4Mac,
		wantIpv6Mac: ipv6Mac,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stubLogOperationResult()
			c := &ixATE{
				cfg: cfg,
				c: &fakeCfgClient{
					xPathToID: map[string]string{
						macXP.String():  "id/to/mac",
						ipv4XP.String(): ipv4ID,
						ipv6XP.String(): ipv6ID,
					},
					updateIDErr: test.idErr,
					session: &fakeSession{
						getRsps: map[string]string{
							ipv4ID: fmt.Sprintf("{\"resolvedGatewayMac\": [\"%s\"]}", test.ipv4MacRsp),
							ipv6ID: fmt.Sprintf("{\"resolvedGatewayMac\": [\"%s\"]}", test.ipv6MacRsp),
						},
						getErrs: map[string]error{
							ipv4ID: test.ipv4Err,
							ipv6ID: test.ipv6Err,
						},
						postRsps: map[string]string{
							valuesOp: fmt.Sprintf("[\"%s\"]", test.ethMacRsp),
						},
						postErrs: map[string]error{
							valuesOp: test.ethErr,
						},
					},
				},
				intfs: map[string]*intf{
					intfName: &intf{
						deviceGroup: dg,
						ipv4:        eth.Ipv4[0],
						ipv6:        eth.Ipv6[0],
					},
				},
			}

			gotErr := resolveMacs(context.Background(), c)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("resolveMacs: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}
			intf := c.intfs[intfName]
			if intf.ethMac != test.wantEthMac {
				t.Errorf("resolveMacs: wrong value for resolved eth mac: got %q, want %q", intf.ethMac, test.wantEthMac)
			}
			if intf.resolvedIpv4Mac != test.wantIpv4Mac {
				t.Errorf("resolveMacs: wrong value for resolved IPv4 mac: got %q, want %q", intf.resolvedIpv4Mac, test.wantIpv4Mac)
			}
			if intf.resolvedIpv6Mac != test.wantIpv6Mac {
				t.Errorf("resolveMacs: wrong value for resolved IPv6 mac: got %q, want %q", intf.resolvedIpv6Mac, test.wantIpv6Mac)
			}
		})
	}
}

func TestResetTrafficCfg(t *testing.T) {
	const (
		stopTrafficOp     = "traffic/operations/stop"
		clearStatsOp      = "operations/clearstats"
		deleteTrafficPath = "traffic/trafficItem"
	)
	defer restoreStubs()
	sleepFn = func(time.Duration) {}

	tests := []struct {
		desc                          string
		stopTrafficErr, clearStatsErr error
		deleteTrafficErr              error
		wantErr                       bool
	}{{
		desc:           "failed to stop traffic",
		stopTrafficErr: errors.New("stop traffic error"),
		wantErr:        true,
	}, {
		desc:          "failed to clear stats",
		clearStatsErr: errors.New("clear stats error"),
		wantErr:       true,
	}, {
		desc:             "failed to delete traffic config",
		deleteTrafficErr: errors.New("delete traffic error"),
		wantErr:          true,
	}, {
		desc: "successful traffic reset",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stubLogOperationResult()
			c := &ixATE{
				c: &fakeCfgClient{
					session: &fakeSession{
						deleteErrs: map[string]error{
							deleteTrafficPath: test.deleteTrafficErr,
						},
						postErrs: map[string]error{
							stopTrafficOp: test.stopTrafficErr,
							clearStatsOp:  test.clearStatsErr,
						},
					},
				},
				cfg:       &ixconfig.Ixnetwork{},
				operState: operStateTrafficOn,
			}

			gotErr := resetIxiaTrafficCfg(context.Background(), c)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("resetIxiaTrafficCfg: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}
		})
	}
}

func TestGenTraffic(t *testing.T) {
	const generateOp = "traffic/trafficItem/operations/generate"
	tiXP := parseXPath(t, "/fake/xpath/trafficItem")
	cfg := &ixconfig.Ixnetwork{
		Traffic: &ixconfig.Traffic{
			TrafficItem: []*ixconfig.TrafficTrafficItem{{Xpath: tiXP}},
		},
	}

	tests := []struct {
		desc        string
		generateErr error
		wantErr     string
	}{{
		desc:        "failed traffic item generate op",
		generateErr: errors.New("failed generate op on traffic items"),
		wantErr:     "could not generate",
	}, {
		desc: "no errors",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			c := &ixATE{
				cfg: cfg,
				c: &fakeCfgClient{
					xPathToID: map[string]string{tiXP.String(): "/id/to/trafficitem"},
					session: &fakeSession{
						postErrs: map[string]error{generateOp: test.generateErr},
					},
				},
				egressTrackFlowCounts: map[string]int{"someFlow": 1},
			}
			gotErr := genTraffic(context.Background(), c)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("genTraffic: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestUpdateFlows(t *testing.T) {
	const (
		flowName = "someFlow"
		tiID     = "id/to/traffic/item"
	)
	tiXP := parseXPath(t, "/fake/xpath/trafficItem")
	baseClient := func() *ixATE {
		cfg := &ixconfig.Ixnetwork{
			Traffic: &ixconfig.Traffic{
				TrafficItem: []*ixconfig.TrafficTrafficItem{{
					Xpath: tiXP,
					ConfigElement: []*ixconfig.TrafficTrafficItemConfigElement{{
						FrameSize: &ixconfig.TrafficTrafficItemConfigElementFrameSize{
							Type_:     ixconfig.String("fixed"),
							FixedSize: ixconfig.NumberInt(1000),
						},
						FrameRate: &ixconfig.TrafficTrafficItemConfigElementFrameRate{
							Type_: ixconfig.String("percentLineRate"),
							Rate:  ixconfig.NumberUint32(20),
						},
					}},
				}},
			},
		}
		return &ixATE{
			cfg: cfg,
			flowToTrafficItem: map[string]*ixconfig.TrafficTrafficItem{
				flowName: cfg.Traffic.TrafficItem[0],
			},
		}
	}

	tests := []struct {
		desc         string
		flow         *opb.Flow
		patchSizeErr error
		patchRateErr error
		wantErr      string
	}{{
		desc:    "non-existent flow",
		flow:    &opb.Flow{Name: "invalid"},
		wantErr: "does not exist",
	}, {
		desc: "bad frame size for flow",
		flow: &opb.Flow{
			Name: flowName,
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_ImixPreset_{
					ImixPreset: opb.FrameSize_IMIX_DEFAULT,
				},
			},
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Bps{Bps: 1500},
			},
		},
		wantErr: "cannot switch to",
	}, {
		desc: "frame size patch error",
		flow: &opb.Flow{
			Name: flowName,
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_Fixed{
					Fixed: 1500,
				},
			},
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Bps{Bps: 1500},
			},
		},
		patchSizeErr: errors.New("error patching size"),
		wantErr:      "error patching size",
	}, {
		desc: "frame rate patch error",
		flow: &opb.Flow{
			Name: flowName,
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_Fixed{
					Fixed: 1500,
				},
			},
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Bps{Bps: 1500},
			},
		},
		patchRateErr: errors.New("error patching rate"),
		wantErr:      "error patching rate",
	}, {
		desc: "successful update",
		flow: &opb.Flow{
			Name: flowName,
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_Fixed{
					Fixed: 1500,
				},
			},
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Bps{Bps: 1500},
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			c := baseClient()
			c.c = &fakeCfgClient{
				session: &fakeSession{
					patchErrs: map[string]error{
						path.Join(tiID, "highLevelStream/1", "frameSize"): test.patchSizeErr,
						path.Join(tiID, "highLevelStream/1", "frameRate"): test.patchRateErr,
					},
				},
				xPathToID: map[string]string{tiXP.String(): tiID},
			}

			gotErr := updateFlows(context.Background(), c, []*opb.Flow{test.flow})
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("updateFlows: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestStart(t *testing.T) {
	const (
		applyTrafficOp = "traffic/operations/apply"
		startTrafficOp = "traffic/operations/start"
	)
	tests := []struct {
		desc                string
		startErr, egressErr error
		wantErr             string
	}{{
		desc:     "failed start traffic op",
		startErr: errors.New("failed to start traffic"),
		wantErr:  "could not start traffic",
	}, {
		desc:      "failed to configure egress tracking",
		egressErr: errors.New("failed to config egress tracking"),
		wantErr:   "could not configure egress tracking view",
	}, {
		desc: "no errors",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			c := &ixATE{
				cfg: &ixconfig.Ixnetwork{
					Traffic: &ixconfig.Traffic{
						TrafficItem: []*ixconfig.TrafficTrafficItem{{}},
					},
				},
				c: &fakeCfgClient{session: &fakeSession{
					postErrs: map[string]error{
						startTrafficOp: test.startErr,
					},
					stats: &fakeStats{
						egressErr: test.egressErr,
					},
				}},
				egressTrackFlowCounts: map[string]int{"someFlow": 1},
			}
			gotErr := startTraffic(context.Background(), c)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("startTraffic: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

// TODO(team): Use a Traffic struct in place of a reqFile and expand the
// tests to cover more behaviors of validateInterfaces.
func TestStartTraffic(t *testing.T) {
	tests := []struct {
		desc                                 string
		operState                            operState
		existingFlows                        map[string]*ixconfig.TrafficTrafficItem
		reqFile, wantCfgFile                 string
		resetTrafficCfgErr, resolveMacsErr   error
		patchLatencyErr, patchConvergenceErr error
		importErr                            error
		updateIDsErr                         error
		genErr                               error
		flowsErr                             error
		applyErr                             error
		startErr                             error
		wantErr                              bool
	}{{
		desc:    "protocols not running",
		reqFile: "no_flows.textproto",
		wantErr: true,
	}, {
		desc:      "traffic already started",
		operState: operStateTrafficOn,
		reqFile:   "no_flows.textproto",
	}, {
		desc:               "failed traffic reset",
		operState:          operStateProtocolsOn,
		resetTrafficCfgErr: errors.New("failed to reset traffic config"),
		reqFile:            "ipv4_flow_with_egress.textproto",
		wantErr:            true,
	}, {
		desc:           "failed mac resolution",
		operState:      operStateProtocolsOn,
		resolveMacsErr: errors.New("failed to resolve macs"),
		reqFile:        "ipv4_flow_with_egress.textproto",
		wantErr:        true,
	}, {
		desc:      "failed traffic config push",
		operState: operStateProtocolsOn,
		importErr: errors.New("traffic config push failed"),
		reqFile:   "ipv4_flow_with_egress.textproto",
		wantErr:   true,
	}, {
		desc:         "failed traffic item ID update",
		operState:    operStateProtocolsOn,
		updateIDsErr: errors.New("traffic item IDs not updated"),
		reqFile:      "ipv4_flow_with_egress.textproto",
		wantErr:      true,
	}, {
		desc:      "failed to generate traffic",
		operState: operStateProtocolsOn,
		genErr:    errors.New("failed to generate traffic items"),
		reqFile:   "ipv4_flow_with_egress.textproto",
		wantErr:   true,
	}, {
		desc:      "failed to update traffic flows",
		operState: operStateProtocolsOn,
		flowsErr:  errors.New("failed to update traffic flows"),
		reqFile:   "ipv4_flow_with_egress.textproto",
		wantErr:   true,
	}, {
		desc:      "failed to apply traffic flows",
		operState: operStateProtocolsOn,
		applyErr:  errors.New("failed to apply traffic flows"),
		reqFile:   "ipv4_flow_with_egress.textproto",
		wantErr:   true,
	}, {
		desc:      "failed to start traffic flows",
		operState: operStateProtocolsOn,
		startErr:  errors.New("failed to start traffic flows"),
		reqFile:   "no_flows.textproto",
		wantErr:   true,
	}, {
		desc:      "no flows configured",
		operState: operStateProtocolsOn,
		reqFile:   "no_flows.textproto",
		wantErr:   true,
	}, {
		desc:          "flows previously configured, no new flows specified",
		operState:     operStateProtocolsOn,
		existingFlows: map[string]*ixconfig.TrafficTrafficItem{"someFlow": {}},
		reqFile:       "no_flows.textproto",
	}, {
		desc:        "ipv4 with egress tracking",
		operState:   operStateProtocolsOn,
		reqFile:     "ipv4_flow_with_egress.textproto",
		wantCfgFile: "ipv4_flow_with_egress.json",
	}, {
		desc:        "ipv6 with flow labels",
		operState:   operStateProtocolsOn,
		reqFile:     "ipv6_flow_with_labels.textproto",
		wantCfgFile: "ipv6_flow_with_labels.json",
	}, {
		desc:            "ipv4 with convergence tracking, failed latency disable",
		operState:       operStateProtocolsOn,
		patchLatencyErr: errors.New("failed to disable latency tracking"),
		reqFile:         "ipv4_flow_with_convergence.textproto",
		wantErr:         true,
	}, {
		desc:                "ipv4 with convergence tracking, failed convergence enable",
		operState:           operStateProtocolsOn,
		patchConvergenceErr: errors.New("failed to enable convergence tracking"),
		reqFile:             "ipv4_flow_with_convergence.textproto",
		wantErr:             true,
	}, {
		desc:        "ipv4 with convergence tracking",
		operState:   operStateProtocolsOn,
		reqFile:     "ipv4_flow_with_convergence.textproto",
		wantCfgFile: "ipv4_flow_with_convergence.json",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			sleepFn = func(time.Duration) {}
			resetIxiaTrafficCfgFn = func(_ context.Context, ix *ixATE) error {
				ix.flowToTrafficItem = map[string]*ixconfig.TrafficTrafficItem{}
				return nil
			}
			c := &ixATE{
				c: &fakeCfgClient{importErrs: []error{nil, nil, nil}},
			}

			top := &Topology{
				Interfaces: []*opb.InterfaceConfig{{
					Name: "intf1",
					Link: &opb.InterfaceConfig_Port{"1/18"},
					Ipv6: &opb.IpConfig{
						AddressCidr:    "2a00:1450:100f:1:192:168:130:2/126",
						DefaultGateway: "2a00:1450:100f:1:192:168:130:1",
					},
				}, {
					Name: "intf2",
					Link: &opb.InterfaceConfig_Port{"2/8"},
					Ipv6: &opb.IpConfig{
						AddressCidr:    "2a00:1450:100f:1:192:168:135:2/126",
						DefaultGateway: "2a00:1450:100f:1:192:168:135:1",
					},
				}},
			}
			if err := c.PushTopology(context.Background(), top); err != nil {
				t.Fatalf("Failed to configure initial topology for StartTraffic test: %v", err)
			}
			// Set existing flows here because they are deleted by PushTopology.
			c.flowToTrafficItem = test.existingFlows

			// Swap out no-op config client with the one to use for StartTraffic testing.
			fc := &fakeCfgClient{
				importErrs:  []error{test.importErr},
				updateIDErr: test.updateIDsErr,
				session: &fakeSession{
					patchErrs: map[string]error{
						"/traffic/statistics/latency":         test.patchLatencyErr,
						"/traffic/statistics/cpdpConvergence": test.patchConvergenceErr,
					},
				},
			}
			c.c = fc
			c.operState = test.operState
			// Initialize stub functions for StartTraffic testing.
			resetIxiaTrafficCfgFn = func(context.Context, *ixATE) error {
				return test.resetTrafficCfgErr
			}
			resolveMacsFn = func(context.Context, *ixATE) error {
				return test.resolveMacsErr
			}
			genTrafficFn = func(context.Context, *ixATE) error {
				return test.genErr
			}
			updateFlowsFn = func(context.Context, *ixATE, []*opb.Flow) error {
				return test.flowsErr
			}
			applyTrafficFn = func(context.Context, *ixATE) error {
				return test.applyErr
			}
			startTrafficFn = func(context.Context, *ixATE) error {
				return test.startErr
			}

			flows := toFlows(t, test.reqFile)
			gotErr := c.StartTraffic(context.Background(), flows)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("StartTraffic: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}

			if !test.wantErr {
				if test.wantCfgFile == "" {
					if fc.lastImportCfg != nil {
						t.Fatalf("StartTraffic: Did not want traffic config pushed, got\n%v", fc.lastImportCfg)
					}
					return
				}
				wantCfg := toTrafficCfg(t, test.wantCfgFile)
				if diff := jsonCfgDiff(t, wantCfg, fc.lastImportCfg); diff != "" {
					t.Fatalf("StartTraffic: Unexpected traffic config pushed, diff (-want, +got)\n%s", diff)
				}
			}
		})
	}
}

func TestUpdateTraffic(t *testing.T) {
	tests := []struct {
		desc      string
		operState operState
		flowsErr  error
		wantErr   string
	}{{
		desc:    "traffic not started",
		wantErr: "cannot update traffic before it has been started",
	}, {
		desc:      "non-existent flow",
		operState: operStateTrafficOn,
		flowsErr:  errors.New("failed to update traffic flows"),
		wantErr:   "failed to update",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			c := &ixATE{operState: test.operState}
			updateFlowsFn = func(context.Context, *ixATE, []*opb.Flow) error {
				return test.flowsErr
			}
			gotErr := c.UpdateTraffic(context.Background(), nil)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("UpdateTraffic: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestStopAllTraffic(t *testing.T) {
	const op = "traffic/operations/stop"
	defer restoreStubs()
	sleepFn = func(time.Duration) {}
	tests := []struct {
		desc      string
		operState operState
		opErr     error
		viewsOut  map[string]view
		wantErr   bool
	}{{
		desc: "traffic not started",
	}, {
		desc:      "Error from stop traffic op",
		operState: operStateTrafficOn,
		opErr:     errors.New("someError"),
		wantErr:   true,
	}, {
		desc:      "Stats never become available",
		operState: operStateTrafficOn,
		viewsOut:  map[string]view{},
		wantErr:   true,
	}, {
		desc:      "Successfully stopped",
		operState: operStateTrafficOn,
		viewsOut:  map[string]view{ixweb.TrafficItemStatsCaption: nil},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stubLogOperationResult()
			c := &ixATE{
				c: &fakeCfgClient{
					session: &fakeSession{
						postErrs: map[string]error{op: test.opErr},
						stats: &fakeStats{
							viewsOut: test.viewsOut,
						},
					},
				},
				operState: test.operState,
			}
			gotErr := c.StopAllTraffic(context.Background())
			if (gotErr != nil) != test.wantErr {
				t.Errorf("StopAllTraffic: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}
		})
	}
}

func TestReadStats(t *testing.T) {
	defer restoreStubs()
	captions := []string{"view1", "view2"}

	tests := []struct {
		desc     string
		viewsOut map[string]view
		viewsErr error
		want     *ixgnmi.Stats
		wantErr  bool
	}{{
		desc:     "error fetching views",
		viewsErr: errors.New("someError"),
		wantErr:  true,
	}, {
		desc: "error fetching table",
		viewsOut: map[string]view{
			"view1": &fakeView{tableOut: ixweb.StatTable{{"k1": "v1"}}},
			"view2": &fakeView{tableErr: errors.New("someError")},
		},
		wantErr: true,
	}, {
		desc: "no error",
		viewsOut: map[string]view{
			"view1": &fakeView{tableOut: ixweb.StatTable{{"k1": "v1"}}},
			"view2": &fakeView{tableOut: ixweb.StatTable{{"k2": "v2"}}},
		},
		want: &ixgnmi.Stats{
			EgressTrackFlows: []string{"someFlow"},
			Tables: map[string]ixweb.StatTable{
				"view1": ixweb.StatTable{{"k1": "v1"}},
				"view2": ixweb.StatTable{{"k2": "v2"}},
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stubLogOperationResult()
			c := &ixATE{
				egressTrackFlowCounts: map[string]int{"someFlow": 1},
				c: &fakeCfgClient{
					session: &fakeSession{stats: &fakeStats{
						viewsOut: test.viewsOut,
						viewsErr: test.viewsErr,
					}},
				},
			}
			got, gotErr := c.readStats(context.Background(), captions)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("readStats unexpected error, got err: %v, want err? %t", gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("readStats unexpected diff, got (-want +got): %s", diff)
			}
		})
	}
}

func TestUpdateBGPPeerStates(t *testing.T) {
	const (
		intfName = "someIntf"
		port     = "1/1"
	)
	tests := []struct {
		desc       string
		ifc        *opb.InterfaceConfig
		importErrs []error
		applyErr   error
		wantErr    string
	}{{
		desc: "BGP v4 peer config failure",
		ifc: &opb.InterfaceConfig{
			Name:     intfName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
			Ipv4: &opb.IpConfig{
				AddressCidr:    "192.168.1.1/30",
				DefaultGateway: "192.168.1.2",
			},
			Bgp: &opb.BgpConfig{
				BgpPeers: []*opb.BgpPeer{{PeerAddress: "1.1.1.1"}},
			},
		},
		importErrs: []error{errors.New("error pushing config")},
		wantErr:    "could not update BGP v4 peer",
	}, {
		desc: "BGP v6 peer config failure",
		ifc: &opb.InterfaceConfig{
			Name:     intfName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
			Ipv6: &opb.IpConfig{
				AddressCidr:    "aa::/126",
				DefaultGateway: "aa::2",
			},
			Bgp: &opb.BgpConfig{
				BgpPeers: []*opb.BgpPeer{{PeerAddress: "bb::"}},
			},
		},
		importErrs: []error{errors.New("error pushing config")},
		wantErr:    "could not update BGP v6 peer",
	}, {
		desc: "BGP v4 loopback peer config failure",
		ifc: &opb.InterfaceConfig{
			Name:             intfName,
			Link:             &opb.InterfaceConfig_Port{port},
			Ethernet:         &opb.EthernetConfig{Mtu: 1500},
			Ipv4LoopbackCidr: "192.168.2.1/30",
			Bgp: &opb.BgpConfig{
				BgpPeers: []*opb.BgpPeer{{
					PeerAddress: "2.2.2.2",
					OnLoopback:  true,
				}},
			},
		},
		importErrs: []error{errors.New("error pushing config")},
		wantErr:    "could not update BGP v4 loopback peer",
	}, {
		desc: "BGP v6 loopback peer config failure",
		ifc: &opb.InterfaceConfig{
			Name:             intfName,
			Link:             &opb.InterfaceConfig_Port{port},
			Ethernet:         &opb.EthernetConfig{Mtu: 1500},
			Ipv6LoopbackCidr: "cafe:beef::1/128",
			Bgp: &opb.BgpConfig{
				BgpPeers: []*opb.BgpPeer{{
					PeerAddress: "cafe::1",
					OnLoopback:  true,
				}},
			},
		},
		importErrs: []error{errors.New("error pushing config")},
		wantErr:    "could not update BGP v6 loopback peer",
	}, {
		desc: "config apply failure",
		ifc: &opb.InterfaceConfig{
			Name:     intfName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		},
		applyErr: errors.New("apply on the fly failure"),
		wantErr:  "could not apply",
	}, {
		desc: "successful update",
		ifc: &opb.InterfaceConfig{
			Name:     intfName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
			Ipv4: &opb.IpConfig{
				AddressCidr:    "192.168.1.1/30",
				DefaultGateway: "192.168.1.2",
			},
			Ipv6: &opb.IpConfig{
				AddressCidr:    "aa::/126",
				DefaultGateway: "aa::2",
			},
			Bgp: &opb.BgpConfig{
				BgpPeers: []*opb.BgpPeer{{PeerAddress: "1.1.1.1"}, {PeerAddress: "bb::"}},
			},
		},
		importErrs: []error{nil, nil},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			cfg := &ixconfig.Ixnetwork{
				Vport: []*ixconfig.Vport{{
					Name:     ixconfig.String(port),
					L1Config: &ixconfig.VportL1Config{},
				}},
			}
			updateXPaths(cfg)
			c := &ixATE{
				cfg:   cfg,
				intfs: map[string]*intf{},
				ports: map[string]*ixconfig.Vport{port: cfg.Vport[0]},
				c: &fakeCfgClient{
					importErrs: test.importErrs,
					session: &fakeSession{postErrs: map[string]error{
						"globals/topology/operations/applyonthefly": test.applyErr,
					}},
				},
			}
			gotErr := c.UpdateBGPPeerStates(context.Background(), []*opb.InterfaceConfig{test.ifc})
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("UpdateBGPPeerStates: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestUpdateNetworkGroups(t *testing.T) {
	const (
		intfName = "someIntf"
		port     = "1/1"
	)
	tests := []struct {
		desc       string
		ifc        *opb.InterfaceConfig
		importErrs []error
		applyErr   error
		wantErr    string
	}{{
		desc: "import failure",
		ifc: &opb.InterfaceConfig{
			Name: intfName,
			Link: &opb.InterfaceConfig_Port{port},
			Ipv4: &opb.IpConfig{
				AddressCidr:    "192.168.31.2/30",
				DefaultGateway: "192.168.31.1",
			},
			Isis: &opb.ISISConfig{
				Level:            opb.ISISConfig_L2,
				Metric:           10,
				AreaId:           "490001",
				EnableWideMetric: true,
				NetworkType:      opb.ISISConfig_POINT_TO_POINT,
			},
			Networks: []*opb.Network{{
				Name:          "net1",
				InterfaceName: intfName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "30.0.0.0/30",
					Count:       15000,
				},
				Isis: &opb.IPReachability{
					Metric:      10,
					RouteOrigin: opb.IPReachability_INTERNAL,
				},
			}},
		},
		importErrs: []error{errors.New("error pushing config")},
		wantErr:    "could not update network groups",
	}, {
		desc: "config apply failure",
		ifc: &opb.InterfaceConfig{
			Name:     intfName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		},
		applyErr: errors.New("apply on the fly failure"),
		wantErr:  "could not apply",
	}, {
		desc: "successful update",
		ifc: &opb.InterfaceConfig{
			Name: intfName,
			Link: &opb.InterfaceConfig_Port{port},
			Ipv4: &opb.IpConfig{
				AddressCidr:    "192.168.31.2/30",
				DefaultGateway: "192.168.31.1",
			},
			Isis: &opb.ISISConfig{
				Level:            opb.ISISConfig_L2,
				Metric:           10,
				AreaId:           "490001",
				EnableWideMetric: true,
				NetworkType:      opb.ISISConfig_POINT_TO_POINT,
			},
			Networks: []*opb.Network{{
				Name:          "net1",
				InterfaceName: intfName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "30.0.0.0/30",
					Count:       15000,
				},
				Isis: &opb.IPReachability{
					Metric:      10,
					RouteOrigin: opb.IPReachability_INTERNAL,
				},
			}},
		},
		importErrs: []error{nil, nil},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			cfg := &ixconfig.Ixnetwork{
				Vport: []*ixconfig.Vport{{
					Name:     ixconfig.String(port),
					L1Config: &ixconfig.VportL1Config{},
				}},
			}
			updateXPaths(cfg)
			c := &ixATE{
				cfg:   cfg,
				intfs: map[string]*intf{},
				ports: map[string]*ixconfig.Vport{port: cfg.Vport[0]},
				c: &fakeCfgClient{
					importErrs: test.importErrs,
					session: &fakeSession{postErrs: map[string]error{
						"globals/topology/operations/applyonthefly": test.applyErr,
					}},
				},
			}
			gotErr := c.UpdateNetworkGroups(context.Background(), []*opb.InterfaceConfig{test.ifc})
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("UpdateNetworkGroups: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func parseXPath(t *testing.T, str string) *ixconfig.XPath {
	xp, err := ixconfig.ParseXPath(str)
	if err != nil {
		t.Fatalf("Error parsing XPath %q: %v", str, err)
	}
	return xp
}

func TestSendBGPGracefulRestart(t *testing.T) {
	const (
		peerV4ID = 4
		peerV6ID = 6
	)
	bgp4XP := parseXPath(t, "/fake/xpath/bgp4")
	bgp6XP := parseXPath(t, "/fake/xpath/bgp6")
	bgpV4 := &ixconfig.TopologyBgpIpv4Peer{Xpath: bgp4XP, Active: ixconfig.MultivalueBool(true)}
	bgpV6 := &ixconfig.TopologyBgpIpv6Peer{Xpath: bgp6XP, Active: ixconfig.MultivalueBool(true)}
	cfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Ethernet: []*ixconfig.TopologyEthernet{{
					Ipv4: []*ixconfig.TopologyIpv4{{
						BgpIpv4Peer: []*ixconfig.TopologyBgpIpv4Peer{bgpV4},
					}},
					Ipv6: []*ixconfig.TopologyIpv6{{
						BgpIpv6Peer: []*ixconfig.TopologyBgpIpv6Peer{bgpV6},
					}},
				}},
			}},
		}},
	}
	intfs := map[string]*intf{"myInt": &intf{
		idToBGPv4Peer: map[uint32]*ixconfig.TopologyBgpIpv4Peer{peerV4ID: bgpV4},
		idToBGPv6Peer: map[uint32]*ixconfig.TopologyBgpIpv6Peer{peerV6ID: bgpV6},
	}}

	tests := []struct {
		desc    string
		peerIDs []uint32
		v4Err   error
		v6Err   error
		wantErr string
	}{{
		desc:    "no bgp peers",
		wantErr: "no bgp peers",
	}, {
		desc:    "v4 send error",
		peerIDs: []uint32{peerV4ID},
		v4Err:   errors.New("error sending bgp v4"),
		wantErr: "error sending bgp v4",
	}, {
		desc:    "v6 send error",
		peerIDs: []uint32{peerV6ID},
		v6Err:   errors.New("error sending bgp v6"),
		wantErr: "error sending bgp v6",
	}, {
		desc:    "success sending graceful restart for bgp v4 peer",
		peerIDs: []uint32{peerV4ID},
	}, {
		desc:    "success sending graceful restart for bgp v6 peer",
		peerIDs: []uint32{peerV6ID},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			c := &ixATE{
				cfg:   cfg,
				intfs: intfs,
				c: &fakeCfgClient{
					xPathToID: map[string]string{
						bgp4XP.String(): "/id/to/bgpv4",
						bgp6XP.String(): "/id/to/bgpv6",
					},
					session: &fakeSession{postErrs: map[string]error{
						bgpPeerV4GracefulRestartOp: test.v4Err,
						bgpPeerV6GracefulRestartOp: test.v6Err,
					}},
				},
			}
			gotErr := c.SendBGPGracefulRestart(context.Background(), test.peerIDs, 0)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("SendBGPGracefulRestart: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestSendBGPPeerNotification(t *testing.T) {
	const (
		peerV4ID = 4
		peerV6ID = 6
	)
	bgp4XP := parseXPath(t, "/fake/xpath/bgp4")
	bgp6XP := parseXPath(t, "/fake/xpath/bgp6")
	bgpV4 := &ixconfig.TopologyBgpIpv4Peer{Xpath: bgp4XP, Active: ixconfig.MultivalueBool(true)}
	bgpV6 := &ixconfig.TopologyBgpIpv6Peer{Xpath: bgp6XP, Active: ixconfig.MultivalueBool(true)}
	cfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Ethernet: []*ixconfig.TopologyEthernet{{
					Ipv4: []*ixconfig.TopologyIpv4{{
						BgpIpv4Peer: []*ixconfig.TopologyBgpIpv4Peer{bgpV4},
					}},
					Ipv6: []*ixconfig.TopologyIpv6{{
						BgpIpv6Peer: []*ixconfig.TopologyBgpIpv6Peer{bgpV6},
					}},
				}},
			}},
		}},
	}
	intfs := map[string]*intf{"myInt": &intf{
		idToBGPv4Peer: map[uint32]*ixconfig.TopologyBgpIpv4Peer{peerV4ID: bgpV4},
		idToBGPv6Peer: map[uint32]*ixconfig.TopologyBgpIpv6Peer{peerV6ID: bgpV6},
	}}

	tests := []struct {
		desc    string
		peerIDs []uint32
		v4Err   error
		v6Err   error
		wantErr string
	}{{
		desc:    "no bgp peers",
		wantErr: "no bgp peers",
	}, {
		desc:    "v4 send error",
		peerIDs: []uint32{peerV4ID},
		v4Err:   errors.New("error sending bgp v4"),
		wantErr: "error sending bgp v4",
	}, {
		desc:    "v6 send error",
		peerIDs: []uint32{peerV6ID},
		v6Err:   errors.New("error sending bgp v6"),
		wantErr: "error sending bgp v6",
	}, {
		desc:    "success sending notification for bgp v6 peer",
		peerIDs: []uint32{peerV4ID},
	}, {
		desc:    "success sending notification for bgp v4 peer",
		peerIDs: []uint32{peerV6ID},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			stubLogOperationResult()
			c := &ixATE{
				cfg:   cfg,
				intfs: intfs,
				c: &fakeCfgClient{
					xPathToID: map[string]string{
						bgp4XP.String(): "/id/to/bgpv4",
						bgp6XP.String(): "/id/to/bgpv6",
					},
					session: &fakeSession{postErrs: map[string]error{
						bgpPeerV4NotifyOp: test.v4Err,
						bgpPeerV6NotifyOp: test.v6Err,
					}},
				},
			}
			gotErr := c.SendBGPPeerNotification(context.Background(), test.peerIDs, 0, 0)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("SendBGPPeerNotification: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}
