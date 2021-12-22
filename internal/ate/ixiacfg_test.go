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
	"golang.org/x/net/context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ondatra/internal/ixgnmi"
	"github.com/openconfig/ondatra/internal/ixweb"

	opb "github.com/openconfig/ondatra/proto"
)

const (
	testDataPath  = "testdata"
	sessionPrefix = "/api/v1/sessions/0"
)

var (
	lastMVID          int
	origImportTimeout = importTimeout
)

func updateXPaths(cfg *ixconfig.Ixnetwork) {
	toXPath := func(p string) *ixconfig.XPath {
		var xp *ixconfig.XPath
		json.Unmarshal([]byte(strconv.Quote(p)), &xp)
		return xp
	}
	for i, vport := range cfg.Vport {
		vport.Xpath = toXPath(fmt.Sprintf("/vport[%d]", i+1))
	}
	for i, lag := range cfg.Lag {
		lag.Xpath = toXPath(fmt.Sprintf("/lag[%d]", i+1))
	}
	const (
		dgXPathFmt            = "/topology[%d]/deviceGroup[%d]"
		ngXPathFmt            = dgXPathFmt + "/networkGroup[%d]"
		bgpRoutePropPathFmt   = ngXPathFmt + "/ipv4PrefixPools[%d]/bgpIPRouteProperty[%d]"
		bgpV6RoutePropPathFmt = ngXPathFmt + "/ipv6PrefixPools[%d]/bgpV6IPRouteProperty[%d]"
		ethXPathFmt           = dgXPathFmt + "/ethernet[%d]"
		ipv4XPathFmt          = ethXPathFmt + "/ipv4[%d]"
		ipv6XPathFmt          = ethXPathFmt + "/ipv6[%d]"
	)
	for i, topo := range cfg.Topology {
		for j, dg := range topo.DeviceGroup {
			dg.Xpath = toXPath(fmt.Sprintf(dgXPathFmt, i+1, j+1))
			for k, ng := range dg.NetworkGroup {
				ng.Xpath = toXPath(fmt.Sprintf(ngXPathFmt, i+1, j+1, k+1))
				for l, pool := range ng.Ipv4PrefixPools {
					for m, brp := range pool.BgpIPRouteProperty {
						brp.Xpath = toXPath(fmt.Sprintf(bgpRoutePropPathFmt, i+1, j+1, k+1, l+1, m+1))
					}
				}
				for l, pool := range ng.Ipv6PrefixPools {
					for m, brp := range pool.BgpV6IPRouteProperty {
						brp.Xpath = toXPath(fmt.Sprintf(bgpV6RoutePropPathFmt, i+1, j+1, k+1, l+1, m+1))
					}
				}
			}
			for k, eth := range dg.Ethernet {
				eth.Xpath = toXPath(fmt.Sprintf(ethXPathFmt, i+1, j+1, k+1))
				if eth.Mac != nil {
					eth.Mac.Xpath = toXPath(fmt.Sprintf("/multivalue[@source = '%s mac']", eth.Xpath.String()))
				}
				for l, ip := range eth.Ipv4 {
					ip.Xpath = toXPath(fmt.Sprintf(ipv4XPathFmt, i+1, j+1, k+1, l+1))
				}
				for l, ip := range eth.Ipv6 {
					ip.Xpath = toXPath(fmt.Sprintf(ipv6XPathFmt, i+1, j+1, k+1, l+1))
				}
			}
		}
	}
	if cfg.Traffic != nil {
		for i, ti := range cfg.Traffic.TrafficItem {
			ti.Xpath = toXPath(fmt.Sprintf("/traffic/trafficItem[%d]", i+1))
		}
	}
}

type fakeCfgClient struct {
	cfgClient
	updateIDErr   error
	pushErrs      []error
	lastPushedCfg ixconfig.IxiaCfgNode
	session       *fakeSession
}

func (c *fakeCfgClient) Session() session {
	return c.session
}

func (c *fakeCfgClient) ImportConfig(ctx context.Context, cfg *ixconfig.Ixnetwork, cfgNode ixconfig.IxiaCfgNode, _ bool) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	updateXPaths(cfg)
	c.lastPushedCfg = cfgNode
	pushErr := c.pushErrs[0]
	c.pushErrs = c.pushErrs[1:]
	return pushErr
}

func (c *fakeCfgClient) GetNode(ctx context.Context, n ixconfig.IxiaCfgNode, v interface{}) error {
	return c.session.Get(ctx, n.GetRestID(), v)
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

// Computes a plausible Ixnetwork REST API ID based on the object's XPath.
// Does not currently work foralias-based XPaths.
func toID(cn ixconfig.IxiaCfgNode) string {
	if cn.GetRestID() != "" {
		return cn.GetRestID()
	}
	var id string
	xp := cn.XPath().String()
	if strings.Contains(xp, "multivalue") {
		id = fmt.Sprintf("/multivalue/%d", lastMVID)
		lastMVID++
	} else {
		id = strings.ReplaceAll(cn.XPath().String(), "[", "/")
		id = strings.ReplaceAll(id, "]", "/")
	}
	return path.Join(sessionPrefix, id)
}

func (c *fakeCfgClient) UpdateIDs(_ context.Context, _ *ixconfig.Ixnetwork, cns ...ixconfig.IxiaCfgNode) error {
	for _, cn := range cns {
		switch c := cn.(type) {
		case *ixconfig.Multivalue:
			c.RestID = toID(c)
		case *ixconfig.TopologyEthernet:
			c.RestID = toID(c)
		case *ixconfig.TopologyIpv4:
			c.RestID = toID(c)
		case *ixconfig.TopologyIpv6:
			c.RestID = toID(c)
		case *ixconfig.TopologyBgpIpRouteProperty:
			c.RestID = toID(c)
		case *ixconfig.TopologyBgpV6IpRouteProperty:
			c.RestID = toID(c)
		case *ixconfig.TrafficTrafficItem:
			c.RestID = toID(c)
		case *ixconfig.Vport:
			c.RestID = toID(c)
		default:
			return fmt.Errorf("id update not supported for node type %T", c)
		}
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

func (s *fakeSession) Get(_ context.Context, p string, v interface{}) error {
	if s.getRsps[p] != "" {
		json.Unmarshal([]byte(s.getRsps[p]), v)
	}
	return s.getErrs[p]
}

func (s *fakeSession) Patch(_ context.Context, p string, _ interface{}) error {
	return s.patchErrs[p]
}

func (s *fakeSession) Post(_ context.Context, p string, _, out interface{}) error {
	if s.postRsps[p] != "" && out != nil {
		if err := json.Unmarshal([]byte(s.postRsps[p]), out); err != nil {
			return err
		}
	}
	return s.postErrs[p]
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

func (s *fakeStats) ConfigEgressView(context.Context, []string) (*ixweb.StatView, error) {
	return nil, s.egressErr
}

type fakeView struct {
	tableOut ixweb.StatTable
	tableErr error
}

func (v *fakeView) FetchTable(ctx context.Context) (ixweb.StatTable, error) {
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

	var want, got map[string]interface{}
	if err := json.Unmarshal(wantBytes, &want); err != nil {
		t.Fatalf("Could not unmarshal expected config to a map: %v", err)
	}
	if err := json.Unmarshal(gotBytes, &got); err != nil {
		t.Fatalf("Could not unmarshal actual config to a map: %v", err)
	}
	return cmp.Diff(want, got)
}

func readTopology(t *testing.T, filename string) *opb.Topology {
	b, err := ioutil.ReadFile(filepath.Join(testDataPath, filename))
	if err != nil {
		t.Fatalf("Could not read file %q: %v", filename, err)
	}
	topo := &opb.Topology{}
	if err := prototext.Unmarshal(b, topo); err != nil {
		t.Fatalf("Could not unmarshal proto from file %q: %v", filename, err)
	}
	return topo
}

func toFlows(t *testing.T, filename string) []*opb.Flow {
	b, err := ioutil.ReadFile(filepath.Join(testDataPath, filename))
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
	b, err := ioutil.ReadFile(filepath.Join(testDataPath, filename))
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
	b, err := ioutil.ReadFile(filepath.Join(testDataPath, filename))
	if err != nil {
		t.Fatalf("Could not read file %q: %v", filename, err)
	}
	cfg := &ixconfig.Traffic{}
	if err := json.Unmarshal(b, cfg); err != nil {
		t.Fatalf("Could not unmarshal IxNetwork config from file %q: %v", filename, err)
	}
	return cfg
}

func restoreStubs() {
	importTimeout = origImportTimeout
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
			importTimeout = origImportTimeout
			if test.timeout {
				importTimeout = 0
			}
			c := &IxiaCfgClient{
				c: &fakeDelayImportClient{
					delays:    test.delays,
					importErr: test.importErr,
				},
				cfg: &ixconfig.Ixnetwork{},
			}
			gotErr := c.importConfig(context.Background(), c.cfg, c.cfg, false)
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
	sleepFn = func(time.Duration) {}
	resetIxiaTrafficCfgFn = func(context.Context, *IxiaCfgClient) error {
		return nil
	}
	tests := []struct {
		desc, reqFile, wantCfgFile string
		pushErrs                   []error
		routeTableImportErr        error
		opErr                      error
		wantErr                    bool
	}{{
		desc:        "Error on port config push",
		reqFile:     "no_intfs.textproto",
		pushErrs:    []error{errors.New("ports push err")},
		wantCfgFile: "no_intfs_cfg.json",
		wantErr:     true,
	}, {
		desc:        "Error on topo config push",
		reqFile:     "no_intfs.textproto",
		pushErrs:    []error{nil, errors.New("topo push err")},
		wantCfgFile: "no_intfs_cfg.json",
		wantErr:     true,
	}, {
		desc:                "Route table import err",
		reqFile:             "no_intfs.textproto",
		pushErrs:            []error{nil, nil},
		routeTableImportErr: errors.New("failed route table import"),
		wantErr:             true,
	}, {
		desc:        "Base config push",
		reqFile:     "no_intfs.textproto",
		pushErrs:    []error{nil, nil},
		wantCfgFile: "no_intfs_cfg.json",
	}, {
		desc:        "IS-IS config with no traffic",
		reqFile:     "isis_no_traffic.textproto",
		pushErrs:    []error{nil, nil},
		wantCfgFile: "isis_no_traffic_cfg.json",
	}, {
		desc:        "FEC disabled",
		reqFile:     "fec_disabled.textproto",
		pushErrs:    []error{nil, nil},
		wantCfgFile: "fec_disabled.json",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fc := &fakeCfgClient{
				pushErrs: test.pushErrs,
			}
			c := IxiaCfgClient{
				c:           fc,
				name:        ixiaName,
				chassisHost: chassisHost,
			}
			syncRouteTableFilesAndImportFn = func(context.Context, *IxiaCfgClient) error {
				return test.routeTableImportErr
			}

			top := readTopology(t, test.reqFile)
			gotErr := c.PushTopology(context.Background(), top)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("PushTopology: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}

			if !test.wantErr {
				wantCfg := toCfg(t, test.wantCfgFile)
				if diff := jsonCfgDiff(t, wantCfg, fc.lastPushedCfg); diff != "" {
					t.Fatalf("PushTopology: Unexpected topology config pushed, diff (-want, +got)\n%s", diff)
				}
			}
		})
	}
}

// Config generation is tested in TestPushTopology, this test is only for Ixnetwork interactions.
func TestUpdateTopology(t *testing.T) {
	tests := []struct {
		desc                 string
		operState            operState
		pushErr              error
		routeTableImportErr  error
		startProtocolsErr    error
		validateProtocolsErr error
		genErr               error
		startErr             error
		wantErr              string
	}{{
		desc:    "Error on config push",
		pushErr: errors.New("push error"),
		wantErr: "push error",
	}, {
		desc:                "Error on route table import",
		routeTableImportErr: errors.New("route table import error"),
		wantErr:             "route table import error",
	}, {
		desc: "traffic/protocols not yet running",
	}, {
		desc:                 "error starting protocols",
		operState:            operStateProtocolsOn,
		startProtocolsErr:    errors.New("could not start protocols"),
		validateProtocolsErr: errors.New("protocols not up"),
		// TODO; Revert to checking start protocols error (see comment in 'startProtocols' section.)
		wantErr: "protocols not up",
	}, {
		desc:                 "error waiting for protocols",
		operState:            operStateProtocolsOn,
		validateProtocolsErr: errors.New("protocols not up"),
		wantErr:              "protocols not up",
	}, {
		desc:      "error starting traffic",
		operState: operStateTrafficOn,
		genErr:    errors.New("could not generate traffic"),
		wantErr:   "could not generate traffic",
	}, {
		desc:      "error starting traffic",
		operState: operStateTrafficOn,
		startErr:  errors.New("could not start traffic"),
		wantErr:   "could not start traffic",
	}, {
		desc:      "successful protocol/traffic start",
		operState: operStateTrafficOn,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			fc := &fakeCfgClient{
				pushErrs: []error{test.pushErr},
				session: &fakeSession{
					postErrs: map[string]error{"/operations/startallprotocols": test.startProtocolsErr},
				},
			}
			c := IxiaCfgClient{
				c:         fc,
				cfg:       &ixconfig.Ixnetwork{Traffic: &ixconfig.Traffic{}},
				operState: test.operState,
			}
			syncRouteTableFilesAndImportFn = func(context.Context, *IxiaCfgClient) error {
				return test.routeTableImportErr
			}
			validateProtocolStartFn = func(context.Context, *IxiaCfgClient) error {
				return test.validateProtocolsErr
			}
			genTrafficFn = func(context.Context, *IxiaCfgClient) error {
				return test.genErr
			}
			startTrafficFn = func(context.Context, *IxiaCfgClient) error {
				return test.startErr
			}

			gotErr := c.UpdateTopology(context.Background(), &opb.Topology{})
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
	baseClient := func() *IxiaCfgClient {
		ng := &ixconfig.TopologyNetworkGroup{
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{}},
			}},
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{}},
			}},
		}
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					NetworkGroup: []*ixconfig.TopologyNetworkGroup{ng},
				}},
			}},
		}
		updateXPaths(cfg)
		return &IxiaCfgClient{
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
			sleepFn = func(time.Duration) {}

			c := baseClient()
			c.routeTableToIxFile = test.syncedFiles
			c.c = &fakeCfgClient{
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
	baseClient := func() *IxiaCfgClient {
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					Ethernet: []*ixconfig.TopologyEthernet{{
						Mac:  &ixconfig.Multivalue{},
						Ipv4: []*ixconfig.TopologyIpv4{{}},
						Ipv6: []*ixconfig.TopologyIpv6{{}},
					}},
				}},
			}},
		}
		updateXPaths(cfg)
		return &IxiaCfgClient{
			cfg: cfg,
			intfs: map[string]*intf{
				"someIntf": &intf{
					deviceGroup: cfg.Topology[0].DeviceGroup[0],
				},
			},
		}
	}
	tests := []struct {
		desc           string
		topoStatus     string
		topoErr        error
		idErr          error
		protocolStatus string
		protocolErr    error
		ethRestartErr  error
		wantErr        string
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
		idErr:      errors.New("error fetching IDs"),
		wantErr:    "could not fetch IDs",
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
	}, {
		desc:           "error restarting eth protocol",
		topoStatus:     "started",
		protocolStatus: "down",
		ethRestartErr:  errors.New("error restarting protocol"),
		wantErr:        "could not restart down protocol",
	}, {
		desc:           "protocols all up",
		topoStatus:     "started",
		protocolStatus: "up",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			sleepFn = func(time.Duration) {}

			c := baseClient()
			eth := c.cfg.Topology[0].DeviceGroup[0].Ethernet[0]
			c.c = &fakeCfgClient{
				updateIDErr: test.idErr,
				session: &fakeSession{
					getRsps: map[string]string{
						"globals/topology": fmt.Sprintf(`{"status": "%s"}`, test.topoStatus),
						toID(eth):          fmt.Sprintf(`{"sessionStatus": ["%s"]}`, test.protocolStatus),
						toID(eth.Ipv4[0]):  fmt.Sprintf(`{"sessionStatus": ["%s"]}`, test.protocolStatus),
						toID(eth.Ipv6[0]):  fmt.Sprintf(`{"sessionStatus": ["%s"]}`, test.protocolStatus),
					},
					getErrs: map[string]error{
						"globals/topology": test.topoErr,
						toID(eth):          test.protocolErr,
						toID(eth.Ipv4[0]):  test.protocolErr,
						toID(eth.Ipv6[0]):  test.protocolErr,
					},
					postErrs: map[string]error{
						"topology/deviceGroup/ethernet/operations/restartdown": test.ethRestartErr,
					},
				},
			}
			gotErr := validateProtocolStart(context.Background(), c)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("validateProtocolStart: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestStartProtocols(t *testing.T) {
	const op = "operations/startallprotocols"
	tests := []struct {
		desc                string
		opErr, protocolsErr error
		wantErr             bool
	}{{
		// Currently a failure is only reported if protocols are not up, even if the operation failed.
		// TODO; Revert to checking that an error is produced(see comment in 'startProtocols' section.)
		desc:  "Error from op",
		opErr: errors.New("someError"),
	}, {
		desc:         "Error waiting for protocols",
		protocolsErr: errors.New("protocols not up"),
		wantErr:      true,
	}, {
		desc: "No error",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			c := IxiaCfgClient{c: &fakeCfgClient{session: &fakeSession{
				postErrs: map[string]error{op: test.opErr},
			}}}
			validateProtocolStartFn = func(context.Context, *IxiaCfgClient) error {
				return test.protocolsErr
			}

			gotErr := c.StartProtocols(context.Background())
			if (gotErr != nil) != test.wantErr {
				t.Errorf("StartProtocols: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
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
			c := IxiaCfgClient{
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
	updateXPaths(cfg)

	tests := []struct {
		desc         string
		port         string
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
		desc: "successfully set port state",
		port: port,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &IxiaCfgClient{
				cfg:   cfg,
				ports: map[string]*ixconfig.Vport{port: cfg.Vport[0]},
				c: &fakeCfgClient{
					session: &fakeSession{
						postErrs: map[string]error{"vport/operations/linkupdn": test.opErr},
					},
					updateIDErr: test.idErr,
				},
			}

			gotErr := c.SetPortState(context.Background(), test.port, true)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("SetPortState: got err: %v, want err %q", gotErr, test.wantErr)
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
	)
	cfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Ethernet: []*ixconfig.TopologyEthernet{{
					Mac:  &ixconfig.Multivalue{},
					Ipv4: []*ixconfig.TopologyIpv4{{}},
					Ipv6: []*ixconfig.TopologyIpv6{{}},
				}},
			}},
		}},
	}
	dg := cfg.Topology[0].DeviceGroup[0]
	eth := dg.Ethernet[0]
	updateXPaths(cfg)

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
			c := &IxiaCfgClient{
				cfg: cfg,
				c: &fakeCfgClient{
					updateIDErr: test.idErr,
					session: &fakeSession{
						getRsps: map[string]string{
							toID(eth.Ipv4[0]): fmt.Sprintf("{\"resolvedGatewayMac\": [\"%s\"]}", test.ipv4MacRsp),
							toID(eth.Ipv6[0]): fmt.Sprintf("{\"resolvedGatewayMac\": [\"%s\"]}", test.ipv6MacRsp),
						},
						getErrs: map[string]error{
							toID(eth.Ipv4[0]): test.ipv4Err,
							toID(eth.Ipv6[0]): test.ipv6Err,
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
			c := &IxiaCfgClient{
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
			c := &IxiaCfgClient{
				cfg: &ixconfig.Ixnetwork{
					Traffic: &ixconfig.Traffic{
						TrafficItem: []*ixconfig.TrafficTrafficItem{{RestID: "path/to/item"}},
					},
				},
				c: &fakeCfgClient{session: &fakeSession{
					postErrs: map[string]error{
						generateOp: test.generateErr,
					},
				}},
				egressTrackingFlows: []string{"someFlow"},
			}
			gotErr := genTraffic(context.Background(), c)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("genTraffic: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestUpdateFlows(t *testing.T) {
	const flowName = "someFlow"
	baseClient := func(c cfgClient) *IxiaCfgClient {
		cfg := &ixconfig.Ixnetwork{
			Traffic: &ixconfig.Traffic{
				TrafficItem: []*ixconfig.TrafficTrafficItem{{
					ConfigElement: []*ixconfig.TrafficConfigElement{{
						FrameSize: &ixconfig.TrafficFrameSize{
							Type_:     ixconfig.String("fixed"),
							FixedSize: ixconfig.NumberInt(1000),
						},
						FrameRate: &ixconfig.TrafficFrameRate{
							Type_: ixconfig.String("percentLineRate"),
							Rate:  ixconfig.NumberUint32(20),
						},
					}},
				}},
			},
		}
		updateXPaths(cfg)
		c.UpdateIDs(context.Background(), cfg, cfg.Traffic.TrafficItem[0])
		return &IxiaCfgClient{
			cfg: cfg,
			c:   c,
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
			fc := &fakeCfgClient{session: &fakeSession{}}
			c := baseClient(fc)
			ti := c.flowToTrafficItem[flowName]
			fc.session.patchErrs = map[string]error{
				path.Join(ti.GetRestID(), "highLevelStream/1", "frameSize"): test.patchSizeErr,
				path.Join(ti.GetRestID(), "highLevelStream/1", "frameRate"): test.patchRateErr,
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
		desc                          string
		applyErr, startErr, egressErr error
		wantErr                       string
	}{{
		desc:     "failed traffic apply op",
		applyErr: errors.New("failed apply op on traffic items"),
		wantErr:  "could not apply",
	}, {
		desc:     "failed start traffic op",
		startErr: errors.New("failed to start traffic"),
		wantErr:  "could not start traffic",
	}, {
		desc:      "failed to configure egress tracking",
		egressErr: errors.New("failed to config egress tracking"),
		wantErr:   "could not set egress stat tracking",
	}, {
		desc: "no errors",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			c := &IxiaCfgClient{
				cfg: &ixconfig.Ixnetwork{
					Traffic: &ixconfig.Traffic{
						TrafficItem: []*ixconfig.TrafficTrafficItem{{RestID: "path/to/item"}},
					},
				},
				c: &fakeCfgClient{session: &fakeSession{
					postErrs: map[string]error{
						applyTrafficOp: test.applyErr,
						startTrafficOp: test.startErr,
					},
					stats: &fakeStats{
						egressErr: test.egressErr,
					},
				}},
				egressTrackingFlows: []string{"someFlow"},
			}
			gotErr := startTraffic(context.Background(), c)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("startTraffic: got err: %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestStartTraffic(t *testing.T) {
	const (
		intfName = "someIntf"
	)

	tests := []struct {
		desc                               string
		operState                          operState
		intfsFile, reqFile, wantCfgFile    string
		resetTrafficCfgErr, resolveMacsErr error
		pushErr                            error
		updateIDsErr                       error
		genErr                             error
		flowsErr                           error
		startErr                           error
		wantErr                            bool
	}{{
		desc:      "protocols not running",
		intfsFile: "no_intfs.textproto",
		reqFile:   "no_flows.textproto",
		wantErr:   true,
	}, {
		desc:      "traffic already started",
		operState: operStateTrafficOn,
		intfsFile: "no_intfs.textproto",
		reqFile:   "no_flows.textproto",
	}, {
		desc:               "failed traffic reset",
		operState:          operStateProtocolsOn,
		resetTrafficCfgErr: errors.New("failed to reset traffic config"),
		intfsFile:          "no_intfs.textproto",
		reqFile:            "no_flows.textproto",
		wantErr:            true,
	}, {
		desc:           "failed mac resolution",
		operState:      operStateProtocolsOn,
		resolveMacsErr: errors.New("failed to resolve macs"),
		intfsFile:      "no_intfs.textproto",
		reqFile:        "no_flows.textproto",
		wantErr:        true,
	}, {
		desc:      "failed traffic config push",
		operState: operStateProtocolsOn,
		pushErr:   errors.New("traffic config push failed"),
		intfsFile: "no_intfs.textproto",
		reqFile:   "no_flows.textproto",
		wantErr:   true,
	}, {
		desc:         "failed traffic item ID update",
		operState:    operStateProtocolsOn,
		updateIDsErr: errors.New("traffic item IDs not updated"),
		intfsFile:    "no_intfs.textproto",
		reqFile:      "no_flows.textproto",
		wantErr:      true,
	}, {
		desc:      "failed to generate traffic",
		operState: operStateProtocolsOn,
		genErr:    errors.New("failed to generate traffic items"),
		intfsFile: "no_intfs.textproto",
		reqFile:   "no_flows.textproto",
		wantErr:   true,
	}, {
		desc:      "failed to update traffic flows",
		operState: operStateProtocolsOn,
		flowsErr:  errors.New("failed to update traffic flows"),
		intfsFile: "no_intfs.textproto",
		reqFile:   "no_flows.textproto",
		wantErr:   true,
	}, {
		desc:      "failed to start traffic flows",
		operState: operStateProtocolsOn,
		startErr:  errors.New("failed to start traffic flows"),
		intfsFile: "no_intfs.textproto",
		reqFile:   "no_flows.textproto",
		wantErr:   true,
	}, {
		desc:        "no flows configured",
		operState:   operStateProtocolsOn,
		intfsFile:   "no_intfs.textproto",
		reqFile:     "no_flows.textproto",
		wantCfgFile: "no_flows_cfg.json",
	}, {
		desc:        "ipv4 with egress tracking",
		operState:   operStateProtocolsOn,
		intfsFile:   "intfs_for_flows.textproto",
		reqFile:     "ipv4_flow_with_egress.textproto",
		wantCfgFile: "ipv4_flow_with_egress.json",
	}, {
		desc:        "ipv6 with flow labels",
		operState:   operStateProtocolsOn,
		intfsFile:   "intfs_for_flows.textproto",
		reqFile:     "ipv6_flow_with_labels.textproto",
		wantCfgFile: "ipv6_flow_with_labels.json",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			defer restoreStubs()
			sleepFn = func(time.Duration) {}
			resetIxiaTrafficCfgFn = func(_ context.Context, ix *IxiaCfgClient) error {
				ix.flowToTrafficItem = map[string]*ixconfig.TrafficTrafficItem{}
				return nil
			}
			c := IxiaCfgClient{
				c: &fakeCfgClient{pushErrs: []error{nil, nil, nil}},
			}
			top := readTopology(t, test.intfsFile)
			if err := c.PushTopology(context.Background(), top); err != nil {
				t.Fatalf("Failed to configure initial topology for StartTraffic test: %v", err)
			}

			// Swap out no-op config client with the one to use for StartTraffic testing.
			fc := &fakeCfgClient{
				pushErrs:    []error{test.pushErr},
				updateIDErr: test.updateIDsErr,
			}
			c.c = fc
			c.operState = test.operState
			// Initialize stub functions for StartTraffic testing.
			resetIxiaTrafficCfgFn = func(context.Context, *IxiaCfgClient) error {
				return test.resetTrafficCfgErr
			}
			resolveMacsFn = func(context.Context, *IxiaCfgClient) error {
				return test.resolveMacsErr
			}
			genTrafficFn = func(context.Context, *IxiaCfgClient) error {
				return test.genErr
			}
			updateFlowsFn = func(context.Context, *IxiaCfgClient, []*opb.Flow) error {
				return test.flowsErr
			}
			startTrafficFn = func(context.Context, *IxiaCfgClient) error {
				return test.startErr
			}

			flows := toFlows(t, test.reqFile)
			gotErr := c.StartTraffic(context.Background(), flows)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("StartTraffic: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}

			if !test.wantErr && test.wantCfgFile != "" {
				wantCfg := toTrafficCfg(t, test.wantCfgFile)
				if diff := jsonCfgDiff(t, wantCfg, fc.lastPushedCfg); diff != "" {
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
			c := &IxiaCfgClient{operState: test.operState}
			updateFlowsFn = func(context.Context, *IxiaCfgClient, []*opb.Flow) error {
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
		wantErr   bool
	}{{
		desc: "traffic not started",
	}, {
		desc:      "Error from stop traffic op",
		operState: operStateTrafficOn,
		opErr:     errors.New("someError"),
		wantErr:   true,
	}, {
		desc:      "Successfully stopped",
		operState: operStateTrafficOn,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := IxiaCfgClient{
				c: &fakeCfgClient{
					session: &fakeSession{postErrs: map[string]error{op: test.opErr}},
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
			c := IxiaCfgClient{
				egressTrackingFlows: []string{"someFlow"},
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
		desc     string
		ifc      *opb.InterfaceConfig
		pushErrs []error
		applyErr error
		wantErr  string
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
		pushErrs: []error{errors.New("error pushing config")},
		wantErr:  "could not update BGP v4 peer",
	}, {
		desc: "BGP v6 peer config failure",
		ifc: &opb.InterfaceConfig{
			Name:     intfName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
			Ipv6: &opb.IpConfig{
				AddressCidr:    "aa::/127",
				DefaultGateway: "aa::2",
			},
			Bgp: &opb.BgpConfig{
				BgpPeers: []*opb.BgpPeer{{PeerAddress: "bb::"}},
			},
		},
		pushErrs: []error{errors.New("error pushing config")},
		wantErr:  "could not update BGP v6 peer",
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
				AddressCidr:    "aa::/127",
				DefaultGateway: "aa::2",
			},
			Bgp: &opb.BgpConfig{
				BgpPeers: []*opb.BgpPeer{{PeerAddress: "1.1.1.1"}, {PeerAddress: "bb::"}},
			},
		},
		pushErrs: []error{nil, nil},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			cfg := &ixconfig.Ixnetwork{
				Vport: []*ixconfig.Vport{{
					Name:     ixconfig.String(port),
					L1Config: &ixconfig.VportL1Config{},
				}},
			}
			updateXPaths(cfg)
			c := &IxiaCfgClient{
				cfg:   cfg,
				intfs: map[string]*intf{},
				ports: map[string]*ixconfig.Vport{port: cfg.Vport[0]},
				c: &fakeCfgClient{
					pushErrs: test.pushErrs,
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
