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
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ondatra/internal/ixgnmi"
	"github.com/openconfig/ondatra/internal/ixweb"
	"github.com/openconfig/ondatra/internal/usererr"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
)

type operState uint
type routeTableFormat string

const (
	operStateOff operState = iota
	operStateProtocolsOn
	operStateTrafficOn
	routeTableFormatCisco   routeTableFormat = "cisco"
	routeTableFormatJuniper routeTableFormat = "juniper"

	importRetries = 5
)

var (
	syncedOpArgs = ixweb.OpArgs{"sync"}

	macRE         = regexp.MustCompile(`^([0-9a-f]{2}:){5}([0-9a-f]{2})$`)
	resolveMacsFn = resolveMacs

	importTimeout = 3 * time.Minute

	sleepFn                        = time.Sleep
	syncRouteTableFilesAndImportFn = syncRouteTableFilesAndImport
	validateProtocolStartFn        = validateProtocolStart
	resetIxiaTrafficCfgFn          = resetIxiaTrafficCfg
	genTrafficFn                   = genTraffic
	updateFlowsFn                  = updateFlows
	startTrafficFn                 = startTraffic
)

type cfgClient interface {
	Session() session
	GetNode(context.Context, ixconfig.IxiaCfgNode, interface{}) error
	ExportConfig(context.Context) (*ixconfig.Ixnetwork, error)
	ImportConfig(context.Context, *ixconfig.Ixnetwork, ixconfig.IxiaCfgNode, bool) error
	UpdateIDs(context.Context, *ixconfig.Ixnetwork, ...ixconfig.IxiaCfgNode) error
}

type session interface {
	ID() int
	AbsPath(string) string
	Delete(context.Context, string) error
	Get(context.Context, string, interface{}) error
	Patch(context.Context, string, interface{}) error
	Post(context.Context, string, interface{}, interface{}) error
	Files() files
	Stats() stats
}

type files interface {
	List(context.Context, string) ([]string, error)
	Upload(context.Context, string, []byte) error
	Delete(context.Context, string) error
}

type stats interface {
	Views(context.Context) (map[string]view, error)
	ConfigEgressView(context.Context, []string) (*ixweb.StatView, error)
}

type view interface {
	FetchTable(context.Context) (ixweb.StatTable, error)
}

type clientWrapper struct {
	*ixconfig.Client
}

func (cw *clientWrapper) Session() session {
	return &sessionWrapper{cw.Client.Session()}
}

func unwrapClient(c cfgClient) *ixconfig.Client {
	return c.(*clientWrapper).Client
}

type sessionWrapper struct {
	*ixweb.Session
}

func (sw *sessionWrapper) Files() files {
	return sw.Session.Files()
}

func (sw *sessionWrapper) Stats() stats {
	return &statsWrapper{sw.Session.Stats()}
}

type statsWrapper struct {
	*ixweb.Stats
}

func (sw *statsWrapper) Views(ctx context.Context) (map[string]view, error) {
	views, err := sw.Stats.Views(ctx)
	if err != nil {
		return nil, err
	}
	m := make(map[string]view, len(views))
	for c, v := range views {
		m[c] = &viewWrapper{v}
	}
	return m, nil
}

type viewWrapper struct {
	*ixweb.StatView
}

func (vw *viewWrapper) FetchTable(ctx context.Context) (ixweb.StatTable, error) {
	return vw.StatView.FetchTable(ctx)
}

// IPv4/6 route tables specified by local path.
type routeTables struct {
	format     routeTableFormat
	ipv4, ipv6 string
}

type intf struct {
	deviceGroup       *ixconfig.TopologyDeviceGroup
	ipv4              *ixconfig.TopologyIpv4
	ipv6              *ixconfig.TopologyIpv6
	ipv4Loopback      *ixconfig.TopologyIpv4Loopback
	ingressLSPs       []*ixconfig.TopologyRsvpP2PIngressLsps
	link              ixconfig.IxiaCfgNode
	isrToNetworkGroup map[string]*ixconfig.TopologyNetworkGroup
	netToNetworkGroup map[string]*ixconfig.TopologyNetworkGroup
	netToRouteTables  map[string]*routeTables
	ethMac            string
	resolvedIpv4Mac   string
	resolvedIpv6Mac   string
}

// New creates an ixia object for interacting with the specified session.
func New(ctx context.Context, name string, ixn *binding.IxNetwork) (*IxiaCfgClient, error) {
	ix := &IxiaCfgClient{
		name:        name,
		c:           &clientWrapper{ixconfig.New(ixn.Session)},
		chassisHost: ixn.ChassisHost,
		syslogHost:  ixn.SyslogHost,
	}
	if ix.chassisHost == "" {
		ix.chassisHost = name
	}
	ix.resetClientCfg()
	// Merge the initial config to set global preferences like syslog streaming,
	// but do not overwrite to avoid intefering with interactive debugging.
	if err := ix.importConfig(ctx, ix.cfg, ix.cfg, false); err != nil {
		return nil, fmt.Errorf("could not set minimal initial config for session: %w", err)
	}
	return ix, nil
}

// IxiaCfgClient implements an ATE interface for a specific IxNetwork session using JSON-based configuration.
type IxiaCfgClient struct {
	c                    cfgClient
	name                 string
	syslogHost           string
	chassisHost          string
	cfgPushCount         int
	cfg                  *ixconfig.Ixnetwork
	routeTableToIxFile   map[string]string // Mapping of route tables (by local path) to IxNetwork file name.
	ports                map[string]*ixconfig.Vport
	lags                 map[string]*ixconfig.Lag
	lagPorts             map[*ixconfig.Lag][]*ixconfig.Vport
	intfs                map[string]*intf
	flowToTrafficItem    map[string]*ixconfig.TrafficTrafficItem
	ingressTrackingFlows []string
	egressTrackingFlows  []string
	// Operational state is updated as needed on successful API calls.
	operState operState

	mu      sync.Mutex
	gclient *ixgnmi.Client
}

// Resets traffic configuration on this client. Does not make any changes on the Ixia.
func (ix *IxiaCfgClient) resetClientTrafficCfg() {
	ix.cfg.Traffic = nil
	ix.flowToTrafficItem = make(map[string]*ixconfig.TrafficTrafficItem)
	ix.ingressTrackingFlows = nil
	ix.egressTrackingFlows = nil
}

// Removes traffic configuration on the Ixia.
func resetIxiaTrafficCfg(ctx context.Context, ix *IxiaCfgClient) error {
	if err := ix.stopAllTraffic(ctx); err != nil {
		return err
	}
	if err := ix.c.Session().Post(ctx, "operations/clearstats", ixweb.OpArgs{}, nil); err != nil {
		return errors.Wrap(err, "could not clear stats")
	}
	if err := ix.c.Session().Delete(ctx, "traffic/trafficItem"); err != nil {
		return errors.Wrap(err, "could not delete IxNetwork traffic items")
	}
	return nil
}

// Resets all configuration on this client. Does not make any changes on the Ixia.
func (ix *IxiaCfgClient) resetClientCfg() {
	ix.cfg = &ixconfig.Ixnetwork{}
	if ix.syslogHost != "" {
		ix.cfg.Globals = &ixconfig.Globals{
			Preferences: &ixconfig.GlobalsPreferences{
				StreamLogsToSyslogServer: ixconfig.Bool(true),
				SyslogHost:               ixconfig.String(ix.syslogHost),
			},
		}
	}
	ix.routeTableToIxFile = make(map[string]string)
	ix.ports = make(map[string]*ixconfig.Vport)
	ix.lags = make(map[string]*ixconfig.Lag)
	ix.lagPorts = make(map[*ixconfig.Lag][]*ixconfig.Vport)
	ix.intfs = make(map[string]*intf)
	ix.resetClientTrafficCfg()
}

// importConfig is a wrapper around the config client ImportConfig method.
// It writes configs as test artifacts before pushing.
func (ix *IxiaCfgClient) importConfig(ctx context.Context, cfg *ixconfig.Ixnetwork, node ixconfig.IxiaCfgNode, overwrite bool) error {
	ix.cfgPushCount++
	fileSuffix := ".json"
	if overwrite {
		fileSuffix = "_overwrite" + fileSuffix
	}
	filePath := fmt.Sprintf("ixnetwork-config-%s-%02d%s", ix.name, ix.cfgPushCount, fileSuffix)
	jsonStr, err := json.MarshalIndent(node, "", "   ")
	if err != nil {
		return fmt.Errorf("could not marshal IxNetwork config for logging to file: %w", err)
	}
	if err := ioutil.WriteFile(filePath, jsonStr, 0644); err != nil {
		return fmt.Errorf("could not log IxNetwork config to file: %w", err)
	}
	log.Infof("IxNetwork config logged to file %s", filePath)

	const importDelay = 15 * time.Second
	importCtx, cancel := context.WithTimeout(ctx, importTimeout)
	defer cancel()

	for i := 0; i < importRetries; i++ {
		err := ix.c.ImportConfig(importCtx, cfg, node, overwrite)
		// If no error or if there is an error and the request did not timeout.
		if err == nil || importCtx.Err() != context.DeadlineExceeded {
			return err
		}
		log.Warningf("Slow import config, canceling and retrying ...")
		sleepFn(importDelay)
		importCtx, cancel = context.WithTimeout(ctx, importTimeout)
		defer cancel()
	}
	return errors.New("timeout importing config")
}

func (ix *IxiaCfgClient) configureTopology(ics []*opb.InterfaceConfig) error {
	ix.cfg.Topology = nil
	ifsByLink := groupByLink(ics)
	for _, ifs := range ifsByLink {
		ix.addTopology(ifs)
		for _, ifc := range ifs {
			// TODO: Add MACsec to the 'golden' ixiajsoncfg PushTopology tests.
			if err := ix.addMACsecProtocol(ifc); err != nil {
				return err
			}
			if err := ix.addIPProtocols(ifc); err != nil {
				return err
			}
			if err := ix.addIPLoopbackProtocols(ifc); err != nil {
				return err
			}
			if err := ix.addISISProtocols(ifc); err != nil {
				return err
			}
			if err := ix.addBGPProtocols(ifc); err != nil {
				return err
			}
			if err := ix.addRSVPProtocols(ifc); err != nil {
				return err
			}
			if err := ix.addNetworks(ifc); err != nil {
				return err
			}
		}
	}
	return nil
}

func groupByLink(ics []*opb.InterfaceConfig) [][]*opb.InterfaceConfig {
	linkToIntfs := make(map[string][]*opb.InterfaceConfig)
	for _, ic := range ics {
		var linkStr string
		switch link := ic.GetLink().(type) {
		// Add "port/" or "link/" prefix so they are sorted deterministically below.
		case *opb.InterfaceConfig_Port:
			linkStr = "port/" + link.Port
		case *opb.InterfaceConfig_Lag:
			linkStr = "link/" + link.Lag
		}
		linkToIntfs[linkStr] = append(linkToIntfs[linkStr], ic)
	}
	var links []string
	for link := range linkToIntfs {
		links = append(links, link)
	}
	sort.Strings(links) // Deterministic ordering for deterministic configs.
	var icGroups [][]*opb.InterfaceConfig
	for _, fp := range links {
		intfs := linkToIntfs[fp]
		sort.Slice(intfs, func(i, j int) bool {
			return strings.Compare(intfs[i].GetName(), intfs[j].GetName()) < 0
		})
		icGroups = append(icGroups, intfs)
	}
	return icGroups
}

func importBgpRoutes(ctx context.Context, ix *IxiaCfgClient, node ixconfig.IxiaCfgNode, rt *routeTables, isV6 bool) error {
	rtFile := rt.ipv4
	pools := "ipv4PrefixPools"
	rp := "bgpIPRouteProperty"
	if isV6 {
		rtFile = rt.ipv6
		pools = "ipv6PrefixPools"
		rp = "bgpV6IPRouteProperty"
	}
	importPath := fmt.Sprintf("topology/deviceGroup/networkGroup/%s/%s/operations/importbgproutes", pools, rp)
	return ix.c.Session().Post(ctx, importPath, ixweb.OpArgs{
		node.GetRestID(),
		"replicate",               // Replicate routes on each configured peer.
		false,                     // Import all routes, not just the 'best' routes.
		"overwriteTestersAddress", // Overwrite next hops in the file to use the device group's IPv4 address.
		rt.format,
		ix.routeTableToIxFile[rtFile],
	}, nil)
}

// Uploads needed route table files to the Ixia chassis as necessary and then imports routes into the corresponding network groups.
func syncRouteTableFilesAndImport(ctx context.Context, ix *IxiaCfgClient) error {
	// Uploaded Ondatra route table file names take the format "ondatra_route_table_<epoch seconds>_<CRC32 hash>
	const rtFilePrefix = "ondatra_route_table_"

	toSync := map[string]bool{}
	brps := []ixconfig.IxiaCfgNode{}
	for _, intf := range ix.intfs {
		for net, rt := range intf.netToRouteTables {
			ng := intf.netToNetworkGroup[net]
			if rt.ipv4 != "" && ix.routeTableToIxFile[rt.ipv4] == "" {
				brps = append(brps, ng.Ipv4PrefixPools[0].BgpIPRouteProperty[0])
				toSync[rt.ipv4] = true
			}
			if rt.ipv6 != "" && ix.routeTableToIxFile[rt.ipv6] == "" {
				brps = append(brps, ng.Ipv6PrefixPools[0].BgpV6IPRouteProperty[0])
				toSync[rt.ipv6] = true
			}
		}
	}

	if len(toSync) == 0 {
		// Return early to avoid further lookups.
		return nil
	}

	hashToFile := map[string]string{}
	for f := range toSync {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return errors.Wrapf(err, "could not read route table file at %s for hashing", f)
		}
		hashToFile[fmt.Sprintf("%d", crc32.ChecksumIEEE(b))] = f
	}

	rtFiles, err := ix.c.Session().Files().List(ctx, rtFilePrefix+"*")
	if err != nil {
		return errors.Wrap(err, "could not query for files on Ixia chassis")
	}

	now := time.Now()
	for _, f := range rtFiles {
		fdata := strings.Split(strings.TrimPrefix(f, rtFilePrefix), "_")
		if len(fdata) != 2 {
			return errors.Errorf("unexpected format for Ixia route table file %s", f)
		}
		ftime, err := strconv.Atoi(fdata[0])
		if err != nil {
			return errors.Errorf("unexpected timestamp format for Ixia route table file %s", f)
		}
		if time.Unix(int64(ftime), 0).Before(now.Add(-24 * 7 * time.Hour)) {
			if err := ix.c.Session().Files().Delete(ctx, f); err != nil {
				return errors.Wrapf(err, "could not delete out-of-date route table file %s", f)
			}
			continue
		}
		if hf := hashToFile[fdata[1]]; hf != "" {
			delete(hashToFile, fdata[1])
			ix.routeTableToIxFile[hf] = f
		}
	}

	for h, f := range hashToFile {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return errors.Wrapf(err, "could not read route table file at %s for upload", f)
		}
		fn := fmt.Sprintf("%s%d_%s", rtFilePrefix, now.Unix(), h)
		if err := ix.c.Session().Files().Upload(ctx, fn, b); err != nil {
			return errors.Wrapf(err, "could not upload route table file at %s to %s", f, fn)
		}
		ix.routeTableToIxFile[f] = fn
	}

	if err := ix.c.UpdateIDs(ctx, ix.cfg, brps...); err != nil {
		return errors.Wrap(err, "could not fetch IDs for importing routes")
	}
	for _, intf := range ix.intfs {
		for net, rt := range intf.netToRouteTables {
			ng := intf.netToNetworkGroup[net]
			if rt.ipv4 != "" {
				if err := importBgpRoutes(ctx, ix, ng.Ipv4PrefixPools[0].BgpIPRouteProperty[0], rt, false); err != nil {
					return errors.Wrapf(err, "errors importing BGP IPv4 routes for network %s", net)
				}
			}
			if rt.ipv6 != "" {
				if err := importBgpRoutes(ctx, ix, ng.Ipv6PrefixPools[0].BgpV6IPRouteProperty[0], rt, true); err != nil {
					return errors.Wrapf(err, "errors importing BGP IPv6 routes for network %s", net)
				}
			}
		}
	}
	return nil
}

func (ix *IxiaCfgClient) updateTopology(ctx context.Context, ics []*opb.InterfaceConfig) error {
	if err := ix.configureTopology(ics); err != nil {
		return err
	}
	if err := ix.importConfig(ctx, ix.cfg, ix.cfg, false); err != nil {
		return err
	}
	return syncRouteTableFilesAndImportFn(ctx, ix)
}

// PushTopology configures the IxNetwork session with the specified topology.
func (ix *IxiaCfgClient) PushTopology(ctx context.Context, top *opb.Topology) error {
	ix.resetClientCfg()
	if err := ix.addPorts(top); err != nil {
		return err
	}
	if err := ix.importConfig(ctx, ix.cfg, ix.cfg, true); err != nil {
		return err
	}
	if lags := top.GetLags(); len(lags) > 0 {
		ix.addLAGs(lags)
		if err := ix.importConfig(ctx, ix.cfg, ix.cfg, false); err != nil {
			return err
		}
	}
	// Avoid a possible race condition with repeated config imports (b/191984474).
	sleepFn(45 * time.Second)
	if err := ix.updateTopology(ctx, top.GetInterfaces()); err != nil {
		return err
	}
	ix.operState = operStateOff
	return nil
}

// UpdateTopology updates IxNetwork session to the specified topology.
func (ix *IxiaCfgClient) UpdateTopology(ctx context.Context, top *opb.Topology) error {
	if err := ix.updateTopology(ctx, top.GetInterfaces()); err != nil {
		return err
	}
	// Protocols/traffic are stopped after updating topology, restart as needed.
	if ix.operState != operStateOff {
		if err := ix.startProtocols(ctx); err != nil {
			return err
		}
		if ix.operState == operStateTrafficOn {
			if err := genTrafficFn(ctx, ix); err != nil {
				return err
			}
			if err := startTrafficFn(ctx, ix); err != nil {
				return err
			}
		}
	}
	return nil
}

type stateRsp interface {
	Up() bool
}

type protocolRsp struct {
	SessionStatus []string
}

func (r *protocolRsp) Up() bool {
	return len(r.SessionStatus) > 0 && r.SessionStatus[0] == "up"
}

type lspRsp struct {
	State []string
}

func (r *lspRsp) Up() bool {
	if len(r.State) == 0 {
		return false
	}
	for _, s := range r.State {
		if s != "up" {
			return false
		}
	}
	return true
}

func checkUp(ctx context.Context, ix *IxiaCfgClient, nodes []ixconfig.IxiaCfgNode, r stateRsp, retries int) ([]ixconfig.IxiaCfgNode, error) {
	const retryWait = 10 * time.Second
	notUp := func(nodes []ixconfig.IxiaCfgNode) ([]ixconfig.IxiaCfgNode, error) {
		var notUp []ixconfig.IxiaCfgNode
		for _, n := range nodes {
			if err := ix.c.GetNode(ctx, n, r); err != nil {
				return nil, errors.Wrapf(err, "could not fetch element at %q to check session status", n.GetRestID())
			}
			if !r.Up() {
				notUp = append(notUp, n)
			}
		}
		return notUp, nil
	}
	nodes, err := notUp(nodes)
	for i := 0; i < retries && err == nil && len(nodes) > 0; i++ {
		sleepFn(retryWait)
		nodes, err = notUp(nodes)
	}
	return nodes, err
}

// Waits for protocols to start, restarting any that initially fail and rechecking as necessary.
func validateProtocolStart(ctx context.Context, ix *IxiaCfgClient) error {
	const (
		retryWait      = 10 * time.Second
		maxRetriesLSPs = 60
	)
	maxRetriesProtocols := 5
	if len(ix.routeTableToIxFile) > 0 {
		// Protocols may need extra time to start on large-scale route imports.
		maxRetriesProtocols = 15
	}
	topo := struct {
		ProtocolActionsInProgress []string
		Status                    string
	}{}
	started := func() (bool, error) {
		if err := ix.c.Session().Get(ctx, "globals/topology", &topo); err != nil {
			return false, errors.Wrap(err, "could not fetch global topology to check protocol status")
		}
		return topo.Status == "started" && len(topo.ProtocolActionsInProgress) == 0, nil
	}
	// Wait for protocols to start.
	start, err := started()
	for i := 0; i < maxRetriesProtocols && err == nil && !start; i++ {
		sleepFn(retryWait)
		start, err = started()
	}
	if err != nil {
		return err
	}
	if !start {
		return errors.New("protocols did not start in time")
	}

	var protocolNodes []ixconfig.IxiaCfgNode
	var ingressLSPNodes []ixconfig.IxiaCfgNode
	for _, intf := range ix.intfs {
		protocolNodes = append(protocolNodes, intf.deviceGroup.Ethernet[0])
		if intf.ipv4 != nil {
			protocolNodes = append(protocolNodes, intf.ipv4)
		}
		if intf.ipv6 != nil {
			protocolNodes = append(protocolNodes, intf.ipv6)
		}
		for _, lsp := range intf.ingressLSPs {
			ingressLSPNodes = append(ingressLSPNodes, lsp)
		}
	}
	var cfgNodes []ixconfig.IxiaCfgNode
	cfgNodes = append(cfgNodes, protocolNodes...)
	cfgNodes = append(cfgNodes, ingressLSPNodes...)
	if err := ix.c.UpdateIDs(ctx, ix.cfg, cfgNodes...); err != nil {
		return errors.Wrap(err, "could not update IDs for checking protocol sessions")
	}

	protocolNodes, err = checkUp(ctx, ix, protocolNodes, &protocolRsp{}, maxRetriesProtocols)
	if err != nil {
		return err
	}
	for _, n := range protocolNodes {
		// Restart any still down protocols individually.
		var op string
		switch cn := n.(type) {
		case *ixconfig.TopologyEthernet:
			op = "topology/deviceGroup/ethernet/operations/restartdown"
		case *ixconfig.TopologyIpv4:
			op = "topology/deviceGroup/ethernet/ipv4/operations/restartdown"
		case *ixconfig.TopologyIpv6:
			op = "topology/deviceGroup/ethernet/ipv6/operations/restartdown"
		default:
			return fmt.Errorf("tried to restart invalid config node type %T", cn)
		}
		if err := ix.c.Session().Post(ctx, op, ixweb.OpArgs{[]string{n.GetRestID()}}, nil); err != nil {
			return fmt.Errorf("could not restart down protocol at %q", n.GetRestID())
		}
	}
	protocolNodes, err = checkUp(ctx, ix, protocolNodes, &protocolRsp{}, maxRetriesProtocols)
	if err != nil {
		return err
	}
	if len(protocolNodes) > 0 {
		var ids []string
		for _, n := range protocolNodes {
			ids = append(ids, n.GetRestID())
		}
		return fmt.Errorf("some protocol instances did not start: %v", ids)
	}

	// Validate ingress LSP state
	ingressLSPNodes, err = checkUp(ctx, ix, ingressLSPNodes, &lspRsp{}, maxRetriesLSPs)
	if err != nil {
		return err
	}
	if len(ingressLSPNodes) > 0 {
		var ids []string
		for _, n := range ingressLSPNodes {
			ids = append(ids, n.GetRestID())
		}
		return fmt.Errorf("some ingress LSP instances did not start: %v", ids)
	}
	return nil
}

func (ix *IxiaCfgClient) startProtocols(ctx context.Context) error {
	if err := ix.c.Session().Post(ctx, "operations/startallprotocols", syncedOpArgs, nil); err != nil {
		log.Warningf("First attempted startallprotocols op failed: %v", err)
		if err := ix.c.Session().Post(ctx, "operations/startallprotocols", syncedOpArgs, nil); err != nil {
			log.Warningf("Second attempted startallprotocols op failed: %v", err)
		}
	}
	// Protocols may have started even if 'startallprotocols' reported a failure,
	// so return based on actual protocol state.
	// TODO: This may be related to behavior on 9.12, try reverting
	// to failing tests on a single 'startallprotocols' error and check the
	// behavior for a topology with RSVP protocols configured after updating to
	// IxNetwork 9.10update2 everywhere.
	if err := validateProtocolStartFn(ctx, ix); err != nil {
		return errors.Wrap(err, "failed to start protocols, also check the log for errors/warnings for the 'startallprotocols' op")
	}
	return nil
}

// StartProtocols starts running protocols for the IxNetwork session.
func (ix *IxiaCfgClient) StartProtocols(ctx context.Context) error {
	if ix.operState != operStateOff {
		log.Infof("Protocols already started, not running operation on Ixia.")
		return nil
	}
	if err := ix.startProtocols(ctx); err != nil {
		return err
	}
	ix.operState = operStateProtocolsOn
	return nil
}

// StopProtocols stops running protocols for the IxNetwork session.
func (ix *IxiaCfgClient) StopProtocols(ctx context.Context) error {
	if ix.operState == operStateOff {
		log.Infof("Protocols already stopped, not running operation on Ixia.")
		return nil
	}
	if err := ix.c.Session().Post(ctx, "operations/stopallprotocols", syncedOpArgs, nil); err != nil {
		return errors.Wrap(err, "could not stop protocols")
	}
	ix.operState = operStateOff
	return nil
}

// SetPortState enables/disables the given Ixia port.
func (ix *IxiaCfgClient) SetPortState(ctx context.Context, port string, enabled bool) error {
	vport, ok := ix.ports[port]
	if !ok {
		return usererr.New("port %q does not exist in current configuration", port)
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, vport); err != nil {
		return errors.Wrapf(err, "could not fetch ID for vport for %q", port)
	}

	state := "down"
	if enabled {
		state = "up"
	}
	if err := ix.c.Session().Post(ctx, "vport/operations/linkupdn", ixweb.OpArgs{[]string{vport.GetRestID()}, state}, nil); err != nil {
		return errors.Wrapf(err, "error setting port state for %q", port)
	}
	return nil
}

func resolveMacs(ctx context.Context, ix *IxiaCfgClient) error {
	var nodes []ixconfig.IxiaCfgNode
	for _, intf := range ix.intfs {
		nodes = append(nodes, intf.deviceGroup.Ethernet[0].Mac)
		if intf.ipv4 != nil {
			nodes = append(nodes, intf.ipv4)
		}
		if intf.ipv6 != nil {
			nodes = append(nodes, intf.ipv6)
		}
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, nodes...); err != nil {
		return errors.Wrap(err, "could not fetch IDs for resolving MACs")
	}

	firstValidMac := func(macs []string) string {
		if len(macs) > 0 && macRE.MatchString(macs[0]) && macs[0] != "00:00:00:00:00:00" {
			return macs[0]
		}
		return ""
	}
	ipData := struct {
		ResolvedGatewayMAC []string `json:"resolvedGatewayMac"`
	}{}

	for _, intf := range ix.intfs {
		eth := intf.deviceGroup.Ethernet[0]
		var macs []string
		err := ix.c.Session().Post(ctx, "multivalue/operations/getvalues", ixweb.OpArgs{eth.Mac.GetRestID(), 0, 1}, &macs)
		if err != nil {
			return errors.Wrapf(err, "could not fetch value for MAC multivalue %q", eth.Mac.GetRestID())
		}
		intf.ethMac = firstValidMac(macs)

		if intf.ipv4 != nil {
			if err := ix.c.GetNode(ctx, intf.ipv4, &ipData); err != nil {
				return errors.Wrapf(err, "could not fetch IPv4 object at %q", intf.ipv4.GetRestID())
			}
			if mac := firstValidMac(ipData.ResolvedGatewayMAC); mac != "" {
				intf.resolvedIpv4Mac = mac
			}
		}

		if intf.ipv6 != nil {
			if err := ix.c.GetNode(ctx, intf.ipv6, &ipData); err != nil {
				return errors.Wrapf(err, "could not fetch IPv6 object at %q", intf.ipv6.GetRestID())
			}
			if mac := firstValidMac(ipData.ResolvedGatewayMAC); mac != "" {
				intf.resolvedIpv6Mac = mac
			}
		}
	}
	return nil
}

// Expects all traffic items to have up-to-date REST IDs.
func genTraffic(ctx context.Context, ix *IxiaCfgClient) error {
	if ix.cfg.Traffic == nil || len(ix.cfg.Traffic.TrafficItem) == 0 {
		return nil
	}
	var tiRefs []string
	for _, ti := range ix.cfg.Traffic.TrafficItem {
		tiRefs = append(tiRefs, ti.GetRestID())
	}
	if err := ix.c.Session().Post(ctx, "traffic/trafficItem/operations/generate", ixweb.OpArgs{tiRefs}, nil); err != nil {
		return errors.Wrap(err, "could not generate traffic flows")
	}
	return nil
}

// updateFlows updates frame size/rate configuration for flows after generation.
// Assumes that IxNetwork traffic items corresponding to the flows have updated
// REST IDs.
func updateFlows(ctx context.Context, ix *IxiaCfgClient, flows []*opb.Flow) error {
	const hlsSuffix = "highLevelStream/1"
	for _, f := range flows {
		ti := ix.flowToTrafficItem[f.GetName()]
		if ti == nil {
			return usererr.New("flow %q does not exist", f.GetName())
		}
		ce := ti.ConfigElement[0]

		fs, fsMap, err := frameSize(f.GetFrameSize())
		if err != nil {
			return errors.Wrapf(err, "could not compute new frame size for flow %q", f.GetName())
		}
		if fs != nil {
			if ce == nil {
				return usererr.New("cannot update frame size for flow %q since it was not originally specified explicitly", f.GetName())
			}
			if *(fs.Type_) != *(ce.FrameSize.Type_) {
				return usererr.New("cannot switch to frame size type %T for flow %q", f.GetFrameSize().GetType(), f.GetName())
			}
			if err := ix.c.Session().Patch(ctx, path.Join(ti.GetRestID(), hlsSuffix, "frameSize"), fsMap); err != nil {
				return errors.Wrapf(err, "could not patch frame size for flow %q", f.GetName())
			}
		}

		_, frMap, err := frameRate(f.GetFrameRate())
		if err != nil {
			return errors.Wrapf(err, "could not compute new frame rate for flow %q", f.GetName())
		}

		if frMap != nil {
			if err := ix.c.Session().Patch(ctx, path.Join(ti.GetRestID(), hlsSuffix, "frameRate"), frMap); err != nil {
				return errors.Wrapf(err, "could not patch frame rate for flow %q", f.GetName())
			}
		}
	}
	return nil
}

func startTraffic(ctx context.Context, ix *IxiaCfgClient) error {
	trafficArgs := ixweb.OpArgs{ix.c.Session().AbsPath("traffic")}
	if err := ix.c.Session().Post(ctx, "traffic/operations/apply", trafficArgs, nil); err != nil {
		return errors.Wrap(err, "could not apply traffic config")
	}
	if err := ix.c.Session().Post(ctx, "traffic/operations/start", trafficArgs, nil); err != nil {
		return errors.Wrap(err, "could not start traffic")
	}

	// TODO: Investigate using JSON-based config instead of the REST API.
	if len(ix.egressTrackingFlows) > 0 {
		if _, err := ix.c.Session().Stats().ConfigEgressView(ctx, ix.egressTrackingFlows); err != nil {
			return errors.Wrap(err, "could not set egress stat tracking")
		}
	}
	return nil
}

// StartTraffic starts traffic for the IxNetwork session based on the given flows.
func (ix *IxiaCfgClient) StartTraffic(ctx context.Context, flows []*opb.Flow) error {
	if ix.operState == operStateOff {
		return usererr.New("cannot start traffic before starting protocols")
	}
	if ix.operState == operStateTrafficOn {
		log.Infof("Traffic already running, not running operation on Ixia.")
		return nil
	}
	ix.resetClientTrafficCfg()
	if err := resetIxiaTrafficCfgFn(ctx, ix); err != nil {
		return errors.Wrap(err, "could not reset traffic config on Ixia")
	}
	// There is a race condition in configuring traffic after clearing traffic, sleep for 30 seconds to avoid.
	sleepFn(30 * time.Second)
	if err := resolveMacsFn(ctx, ix); err != nil {
		return errors.Wrap(err, "could not resolve MAC addresses in config")
	}

	if err := ix.addTraffic(flows); err != nil {
		return errors.Wrap(err, "could not compute traffic configuration")
	}

	if err := ix.importConfig(ctx, ix.cfg, ix.cfg.Traffic, false); err != nil {
		return errors.Wrap(err, "could not push traffic configuration")
	}

	var items []ixconfig.IxiaCfgNode
	for _, ti := range ix.cfg.Traffic.TrafficItem {
		items = append(items, ti)
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, items...); err != nil {
		return errors.Wrap(err, "could not fetch refs for traffic items")
	}

	if err := genTrafficFn(ctx, ix); err != nil {
		return errors.Wrap(err, "could not generate traffic flow")
	}

	// IxNetwork does not necessarily configure frame rate correctly when generating
	// traffic streams (b/204318369).
	if err := updateFlowsFn(ctx, ix, flows); err != nil {
		return errors.Wrap(err, "could not update generated flows")
	}

	if err := startTrafficFn(ctx, ix); err != nil {
		return errors.Wrap(err, "could not start traffic")
	}

	ix.operState = operStateTrafficOn
	return nil
}

// UpdateTraffic updates traffic config for the IxNetwork session based on the given flows.
func (ix *IxiaCfgClient) UpdateTraffic(ctx context.Context, flows []*opb.Flow) error {
	if ix.operState != operStateTrafficOn {
		return usererr.New("cannot update traffic before it has been started")
	}
	if err := updateFlowsFn(ctx, ix, flows); err != nil {
		return errors.Wrap(err, "could not update running traffic flows")
	}
	return nil
}

func (ix *IxiaCfgClient) stopAllTraffic(ctx context.Context) error {
	trafficArgs := ixweb.OpArgs{ix.c.Session().AbsPath("traffic")}
	if err := ix.c.Session().Post(ctx, "traffic/operations/stop", trafficArgs, nil); err != nil {
		return errors.Wrap(err, "could not stop traffic")
	}
	// Wait a sufficient amount of time to ensure that traffic is stopped.
	sleepFn(15 * time.Second)
	return nil
}

// StopAllTraffic stops all traffic for the IxNetwork session.
func (ix *IxiaCfgClient) StopAllTraffic(ctx context.Context) error {
	if ix.operState != operStateTrafficOn {
		log.Infof("Traffic already stopped, not running operation on Ixia.")
		return nil
	}
	if err := ix.stopAllTraffic(ctx); err != nil {
		return err
	}
	ix.operState = operStateProtocolsOn
	return nil
}

// DialGNMI constructs and returns a GNMI client for the Ixia.
func (ix *IxiaCfgClient) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	ix.mu.Lock()
	defer ix.mu.Unlock()
	if ix.gclient != nil {
		return nil, errors.Errorf("GNMI client for Ixia %s already dialed", ix.name)
	}
	gclient, err := ixgnmi.NewClient(ctx, ix.name, ix.readStats, unwrapClient(ix.c), opts...)
	if err != nil {
		return nil, err
	}
	ix.gclient = gclient
	return gclient, nil
}

func (ix *IxiaCfgClient) readStats(ctx context.Context, captions []string) (*ixgnmi.Stats, error) {
	if len(ix.egressTrackingFlows) == 0 {
		var cmod []string
		for _, c := range captions {
			if c != ixweb.EgressStatsCaption {
				cmod = append(cmod, c)
			}
		}
		captions = cmod
	}
	views, err := ix.c.Session().Stats().Views(ctx)
	if err != nil {
		return nil, err
	}
	tables := make(map[string]ixweb.StatTable)
	for _, cap := range captions {
		view, ok := views[cap]
		var table ixweb.StatTable
		if ok {
			table, err = view.FetchTable(ctx)
			if err != nil {
				return nil, err
			}
		} else {
			log.Warningf("No view with caption %q, got views %v", cap, views)
		}
		tables[cap] = table
	}
	if err != nil {
		return nil, errors.Wrapf(err, "Error retrieving statistics for views: %v", captions)
	}
	return &ixgnmi.Stats{
		Tables:            tables,
		IngressTrackFlows: ix.ingressTrackingFlows,
		EgressTrackFlows:  ix.egressTrackingFlows,
	}, nil
}

// FlushStats will remove all data from the cache for the Ixia and marks all stats as stale.
// This ensures that future telemetry calls attempt to fetch data and that stale data doesn't persist.
func (ix *IxiaCfgClient) FlushStats() {
	ix.mu.Lock()
	defer ix.mu.Unlock()
	if ix.gclient != nil {
		ix.gclient.Flush()
	}
}

// UpdateBGPPeerStates exists only to match the API of the prior IxNetwork ATE binding.
// It assumes that the only changes in the provided interface configs are updates to
// BGP active states.
// TODO: Remove this method once new Ixia config binding is used.
func (ix *IxiaCfgClient) UpdateBGPPeerStates(ctx context.Context, ifs []*opb.InterfaceConfig) error {
	if err := ix.configureTopology(ifs); err != nil {
		return err
	}
	var peerStates []*ixconfig.MultivalueSingleValue
	for _, intf := range ix.intfs {
		if v4 := intf.ipv4; v4 != nil {
			for _, p := range v4.BgpIpv4Peer {
				peerStates = append(peerStates, p.Active.SingleValue)
			}
		}
		if v6 := intf.ipv6; v6 != nil {
			for _, p := range v6.BgpIpv6Peer {
				peerStates = append(peerStates, p.Active.SingleValue)
			}
		}
		if v4lb := intf.ipv4Loopback; v4lb != nil {
			for _, p := range v4lb.BgpIpv4Peer {
				peerStates = append(peerStates, p.Active.SingleValue)
			}
		}
	}
	if len(peerStates) == 0 {
		return nil
	}
	var psCfgNodes []ixconfig.IxiaCfgNode
	for _, p := range peerStates {
		psCfgNodes = append(psCfgNodes, p)
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, psCfgNodes...); err != nil {
		return errors.Wrap(err, "could not update IDs of BGP peer 'active' multivalue nodes")
	}
	for _, p := range peerStates {
		if err := ix.c.Session().Patch(ctx, p.GetRestID(), map[string]string{"value": *(p.Value)}); err != nil {
			return errors.Wrapf(err, "could not update peer state for node at XPath %v", p.XPath())
		}
	}
	const (
		applyOnTheFlyArg = "globals/topology"
		applyOnTheFlyOp  = "globals/topology/operations/applyonthefly"
	)
	applyOnTheFlyArgs := ixweb.OpArgs{ix.c.Session().AbsPath(applyOnTheFlyArg)}
	if err := ix.c.Session().Post(ctx, applyOnTheFlyOp, applyOnTheFlyArgs, nil); err != nil {
		return errors.Wrap(err, "could not apply topology changes")
	}
	return nil
}
