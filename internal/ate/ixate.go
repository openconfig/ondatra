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
	"hash/crc32"
	"net"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ondatra/internal/ixgnmi"
	"github.com/openconfig/ondatra/internal/rawapis"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
)

type operState uint
type routeTableFormat string

// OperStatus represents the status of a completed IxNetwork operation.
type OperStatus string

const (
	operStateOff operState = iota
	operStateProtocolsOn
	operStateTrafficOn
	routeTableFormatCisco   routeTableFormat = "cisco"
	routeTableFormatJuniper routeTableFormat = "juniper"
	routeTableFormatCSV     routeTableFormat = "csv"

	// OperStatusSuccess indicates a successful operation.
	OperStatusSuccess OperStatus = "success"
	// OperStatusFailure indicates a failed operation.
	OperStatusFailure OperStatus = "failure"

	importRetries = 5

	// IxNetwork has an undocumented maximum number of DeviceGroups.
	maxIntfs = 256

	bgpPeerV4NotifyOp = "topology/deviceGroup/ethernet/ipv4/bgpIpv4Peer/operations/breaktcpsession"
	bgpPeerV6NotifyOp = "topology/deviceGroup/ethernet/ipv6/bgpIpv6Peer/operations/breaktcpsession"

	bgpPeerV4GracefulRestartOp = "topology/deviceGroup/ethernet/ipv4/bgpIpv4Peer/operations/gracefulrestart"
	bgpPeerV6GracefulRestartOp = "topology/deviceGroup/ethernet/ipv6/bgpIpv6Peer/operations/gracefulrestart"

	startLACPOp = "lag/protocolStack/ethernet/lagportlacp/port/operations/start"
	stopLACPOp  = "lag/protocolStack/ethernet/lagportlacp/port/operations/stop"
)

var (
	syncedOpArgs = ixweb.OpArgs{"sync"}

	macRE         = regexp.MustCompile(`^([0-9a-f]{2}:){5}([0-9a-f]{2})$`)
	resolveMacsFn = resolveMacs

	// TODO(team): Lower timeouts after chassis hardware upgrades.
	peersImportTimeout   = time.Minute
	trafficImportTimeout = 4 * time.Minute
	topoImportTimeout    = 3 * time.Minute

	sleepFn                        = time.Sleep
	syncRouteTableFilesAndImportFn = syncRouteTableFilesAndImport
	validateProtocolStartFn        = validateProtocolStart
	resetIxiaTrafficCfgFn          = resetIxiaTrafficCfg
	genTrafficFn                   = genTraffic
	updateFlowsFn                  = updateFlows
	configureTrafficFn             = configureTraffic
	applyTrafficFn                 = applyTraffic
	startTrafficFn                 = startTraffic
)

type cfgClient interface {
	Session() session
	NodeID(ixconfig.IxiaCfgNode) (string, error)
	ExportConfig(context.Context) (*ixconfig.Ixnetwork, error)
	ImportConfig(context.Context, *ixconfig.Ixnetwork, ixconfig.IxiaCfgNode, bool) error
	UpdateIDs(context.Context, *ixconfig.Ixnetwork, ...ixconfig.IxiaCfgNode) error
}

type session interface {
	ID() int
	AbsPath(string) string
	Delete(context.Context, string) error
	Get(context.Context, string, any) error
	Patch(context.Context, string, any) error
	Post(context.Context, string, any, any) error
	Errors(context.Context) ([]*ixweb.Error, error)
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
	ConfigEgressView(context.Context, []string, int) (*ixweb.StatView, error)
}

type view interface {
	FetchTable(context.Context, ...string) (ixweb.StatTable, error)
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

func (vw *viewWrapper) FetchTable(ctx context.Context, drilldowns ...string) (ixweb.StatTable, error) {
	return vw.StatView.FetchTable(ctx, drilldowns...)
}

// CfgComponents represents the physical ATE components in use for the session and
// their association with configured interfaces/protocols.
type CfgComponents struct {
	Host                 string
	Linecards            []uint64
	Ports                []string
	PortToInterfaces     map[string][]string
	InterfaceToProtocols map[string][]string
}

func (c *CfgComponents) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("<error marshaling JSON: %v>", err)
	}
	return string(b)
}

// OperResult represents the status of an operation and the configuration context in which it executed.
type OperResult struct {
	Path        string
	Status      OperStatus
	Start       time.Time
	End         time.Time
	OpErr       error
	SessionErrs []*ixweb.Error
	Components  *CfgComponents
}

func (o *OperResult) String() string {
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return fmt.Sprintf("<error marshaling JSON: %v>", err)
	}
	return string(b)
}

// IPv4/6 route tables specified by local path.
type routeTables struct {
	format           routeTableFormat
	ipv4, ipv6       string
	overwriteNexthop bool
}

type intf struct {
	deviceGroup       *ixconfig.TopologyDeviceGroup
	ipv4              *ixconfig.TopologyIpv4
	ipv6              *ixconfig.TopologyIpv6
	ipv4Loopback      *ixconfig.TopologyIpv4Loopback
	ipv6Loopback      *ixconfig.TopologyIpv6Loopback
	rsvpLSPs          map[string]*ixconfig.TopologyRsvpteLsps
	link              ixconfig.IxiaCfgNode
	isrToNetworkGroup map[string]*ixconfig.TopologyNetworkGroup
	netToNetworkGroup map[string]*ixconfig.TopologyNetworkGroup
	netToRouteTables  map[string]*routeTables
	ethMac            string
	resolvedIpv4Mac   string
	resolvedIpv6Mac   string
	hasVLAN           bool
	idToBGPv4Peer     map[uint32]*ixconfig.TopologyBgpIpv4Peer
	idToBGPv6Peer     map[uint32]*ixconfig.TopologyBgpIpv6Peer
}

func (i *intf) String() string {
	return fmt.Sprintf("%+v", *i)
}

func newIxATE(ctx context.Context, name string, ixn *binding.IxNetwork) (*ixATE, error) {
	ix := &ixATE{
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
	if err := ix.importConfig(ctx, ix.cfg, false, time.Minute, "minimal initial config reset"); err != nil {
		return nil, fmt.Errorf("could not set minimal initial config for session: %w", err)
	}
	return ix, nil
}

// ixATE provides an ATE interface backed by an IxNetwork session.
type ixATE struct {
	c                     cfgClient
	name                  string
	syslogHost            string
	chassisHost           string
	cfgPushCount          int
	cfg                   *ixconfig.Ixnetwork
	routeTableToIxFile    map[string]string // Mapping of route tables (by local path) to IxNetwork file name.
	ports                 map[string]*ixconfig.Vport
	lags                  map[string]*ixconfig.Lag
	lagPorts              map[*ixconfig.Lag][]*ixconfig.Vport
	intfs                 map[string]*intf
	flowToTrafficItem     map[string]*ixconfig.TrafficTrafficItem
	ingressTrackFlows     []string
	egressTrackFlowCounts map[string]int
	convergenceTracking   bool
	// Operational state is updated as needed on successful API calls.
	operState operState

	mu      sync.Mutex
	gclient *ixgnmi.Client
}

func linkToPorts(ix *ixATE, link ixconfig.IxiaCfgNode) ([]string, error) {
	var ports []string
	switch l := link.(type) {
	case *ixconfig.Vport:
		ports = []string{*l.Name}
	case *ixconfig.Lag:
		vports := ix.lagPorts[l]
		for _, vp := range vports {
			ports = append(ports, *vp.Name)
		}
	default:
		return nil, fmt.Errorf("link node is not a Vport or Lag link: %v(%T)", link, link)
	}
	return ports, nil
}

func (ix *ixATE) logOp(ctx context.Context, path string, opErr error, start, end time.Time) {
	components := &CfgComponents{
		Host:                 ix.name,
		PortToInterfaces:     map[string][]string{},
		InterfaceToProtocols: map[string][]string{},
	}
	linecards := map[uint64]bool{}
	for port := range ix.ports {
		components.Ports = append(components.Ports, port)
		lc, err := strconv.ParseUint(strings.Split(port, "/")[0], 10, 64)
		if err != nil {
			log.Errorf("Could not parse linecard from port %q: %v", port, err)
			continue
		}
		linecards[lc] = true
	}
	for lc := range linecards {
		components.Linecards = append(components.Linecards, lc)
	}
	for name, intf := range ix.intfs {
		ports, err := linkToPorts(ix, intf.link)
		if err != nil {
			log.Errorf("Could not compute ports for intf %q", name)
			continue
		}
		for _, port := range ports {
			components.PortToInterfaces[port] = append(components.PortToInterfaces[port], name)
		}
		var protocolsForIntf []string
		if v4 := intf.ipv4; v4 != nil {
			protocolsForIntf = append(protocolsForIntf, "IPv4")
			if len(v4.BgpIpv4Peer) != 0 {
				protocolsForIntf = append(protocolsForIntf, "BGPOnV4")
			}
		}
		if v6 := intf.ipv6; v6 != nil {
			protocolsForIntf = append(protocolsForIntf, "IPv6")
			if len(v6.BgpIpv6Peer) != 0 {
				protocolsForIntf = append(protocolsForIntf, "BGPOnV6")
			}
		}
		if intf.ipv4Loopback != nil {
			protocolsForIntf = append(protocolsForIntf, "IPv4Loopback")
		}
		if intf.ipv6Loopback != nil {
			protocolsForIntf = append(protocolsForIntf, "IPv6Loopback")
		}
		if len(intf.deviceGroup.IsisL3Router) != 0 {
			protocolsForIntf = append(protocolsForIntf, "IS-IS")
		}
		if len(intf.rsvpLSPs) != 0 {
			protocolsForIntf = append(protocolsForIntf, "RSVP")
		}
		components.InterfaceToProtocols[name] = protocolsForIntf
	}

	status := OperStatusSuccess
	if opErr != nil {
		status = OperStatusFailure
	}
	opResult := &OperResult{
		Path:        path,
		Status:      status,
		Start:       start,
		End:         end,
		OpErr:       opErr,
		SessionErrs: ix.sessionErrors(ctx),
		Components:  components,
	}
	log.Infof(opResult.String())
}

func (ix *ixATE) runOp(ctx context.Context, path string, in any, out any) error {
	start := time.Now()
	err := ix.c.Session().Post(ctx, path, in, out)
	end := time.Now()
	ix.logOp(ctx, path, err, start, end)
	return err
}

// Resets traffic configuration on this client. Does not make any changes on the Ixia.
func (ix *ixATE) resetClientTrafficCfg() {
	ix.cfg.Traffic = nil
	ix.flowToTrafficItem = make(map[string]*ixconfig.TrafficTrafficItem)
	ix.ingressTrackFlows = nil
	ix.egressTrackFlowCounts = make(map[string]int)
	ix.convergenceTracking = false
}

// Removes traffic configuration on the Ixia.
func resetIxiaTrafficCfg(ctx context.Context, ix *ixATE) error {
	if err := ix.stopAllTraffic(ctx); err != nil {
		return err
	}
	if err := ix.runOp(ctx, "operations/clearstats", ixweb.OpArgs{}, nil); err != nil {
		return fmt.Errorf("could not clear stats: %w", err)
	}
	if err := ix.c.Session().Delete(ctx, "traffic/trafficItem"); err != nil {
		return fmt.Errorf("could not delete IxNetwork traffic items: %w", err)
	}
	return nil
}

// Resets all configuration on this client. Does not make any changes on the Ixia.
func (ix *ixATE) resetClientCfg() {
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

// sessionErrors returns the errors reported by the IxNetwork session.
// Does not fail if there was an error querying the session ('nil' is
// returned both in this case and if there were no errors.)
func (ix *ixATE) sessionErrors(ctx context.Context) []*ixweb.Error {
	errors, err := ix.c.Session().Errors(ctx)
	if err != nil {
		log.Errorf("Could not fetch errors for IxNetwork session: %v", err)
		return nil
	}
	var sErrs []*ixweb.Error
	for _, e := range errors {
		if e.Level == "kError" {
			sErrs = append(sErrs, e)
		}
	}
	return sErrs
}

// importConfig is a wrapper around the config client ImportConfig method.
// It writes configs as test artifacts before pushing.
func (ix *ixATE) importConfig(ctx context.Context, node ixconfig.IxiaCfgNode, overwrite bool, timeout time.Duration, desc string) error {
	log.Infof("Importing config for %s", desc)
	ix.cfgPushCount++
	fileSuffix := ".json"
	if overwrite {
		fileSuffix = "_overwrite" + fileSuffix
	}
	desc = strings.ReplaceAll(desc, " ", "_")
	fileName := fmt.Sprintf("ixnetwork-config-%s-%02d-%s-%s", ix.name, ix.cfgPushCount, desc, fileSuffix)
	outputDir := os.Getenv("TEST_UNDECLARED_OUTPUTS_DIR")
	if outputDir == "" {
		outputDir = os.TempDir()
	}
	filePath := path.Join(outputDir, fileName)
	defer func() {
		// Record the pushed config after XPaths have been updated by ix.c.ImportConfig.
		jsonStr, err := json.MarshalIndent(node, "", "   ")
		if err != nil {
			log.Errorf("could not marshal IxNetwork config for logging to file: %v", err)
		}
		if err := os.WriteFile(filePath, jsonStr, 0644); err != nil {
			log.Errorf("could not log IxNetwork config to file: %v", err)
		}
		log.Infof("IxNetwork config push attempt logged to file %s", filePath)
	}()

	const importDelay = 15 * time.Second
	importCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for i := 0; i < importRetries; i++ {
		start := time.Now()
		err := ix.c.ImportConfig(importCtx, ix.cfg, node, overwrite)
		end := time.Now()
		ix.logOp(ctx, "importconfig", err, start, end)
		// If no error or if there is an error and the request did not timeout.
		if err == nil || importCtx.Err() != context.DeadlineExceeded {
			return err
		}
		log.Warningf("Slow import config, canceling and retrying ...")
		sleepFn(importDelay)
		importCtx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	return errors.New("timeout importing config")
}

func (ix *ixATE) configureTopology(ics []*opb.InterfaceConfig) error {
	if len(ix.ports) == 0 {
		return errors.New("no ports configured for topology, check if (*Topology).Update called before (*Topology).Push")
	}
	ix.cfg.Topology = nil
	ifsByLink := groupByLink(ics)
	for _, ifs := range ifsByLink {
		if err := ix.addTopology(ifs); err != nil {
			return err
		}
		for _, ifc := range ifs {
			// TODO(greg-dennis): Add MACsec to the 'golden' ixiajsoncfg PushTopology tests.
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
			if err := ix.addDHCPProtocols(ifc); err != nil {
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

func importBgpRoutes(ctx context.Context, ix *ixATE, node ixconfig.IxiaCfgNode, rt *routeTables, isV6 bool) error {
	rtFile := rt.ipv4
	pools := "ipv4PrefixPools"
	rp := "bgpIPRouteProperty"
	if isV6 {
		rtFile = rt.ipv6
		pools = "ipv6PrefixPools"
		rp = "bgpV6IPRouteProperty"
	}
	nextHop := "preserveFromFile"
	if rt.overwriteNexthop {
		nextHop = "overwriteTestersAddress"
	}

	nodeID, err := ix.c.NodeID(node)
	if err != nil {
		return err
	}
	importPath := fmt.Sprintf("topology/deviceGroup/networkGroup/%s/%s/operations/importbgproutes", pools, rp)
	return ix.runOp(ctx, importPath, ixweb.OpArgs{
		nodeID,
		"replicate", // Replicate routes on each configured peer.
		false,       // Import all routes, not just the 'best' routes.
		nextHop,
		rt.format,
		ix.routeTableToIxFile[rtFile],
	}, nil)
}

// Uploads needed route table files to the Ixia chassis as necessary and then imports routes into the corresponding network groups.
func syncRouteTableFilesAndImport(ctx context.Context, ix *ixATE) error {
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
		b, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("could not read route table file at %s for hashing: %w", f, err)
		}
		hashToFile[fmt.Sprintf("%d", crc32.ChecksumIEEE(b))] = f
	}

	rtFiles, err := ix.c.Session().Files().List(ctx, rtFilePrefix+"*")
	if err != nil {
		return fmt.Errorf("could not query for files on Ixia chassis: %w", err)
	}

	now := time.Now()
	for _, f := range rtFiles {
		fdata := strings.Split(strings.TrimPrefix(f, rtFilePrefix), "_")
		if len(fdata) != 2 {
			return fmt.Errorf("unexpected format for Ixia route table file %s", f)
		}
		ftime, err := strconv.Atoi(fdata[0])
		if err != nil {
			return fmt.Errorf("unexpected timestamp format for Ixia route table file %s", f)
		}
		if time.Unix(int64(ftime), 0).Before(now.Add(-24 * 7 * time.Hour)) {
			if err := ix.c.Session().Files().Delete(ctx, f); err != nil {
				return fmt.Errorf("could not delete out-of-date route table file %s: %w", f, err)
			}
			continue
		}
		if hf := hashToFile[fdata[1]]; hf != "" {
			delete(hashToFile, fdata[1])
			ix.routeTableToIxFile[hf] = f
		}
	}

	for h, f := range hashToFile {
		b, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("could not read route table file at %s for upload: %w", f, err)
		}
		fn := fmt.Sprintf("%s%d_%s", rtFilePrefix, now.Unix(), h)
		if err := ix.c.Session().Files().Upload(ctx, fn, b); err != nil {
			return fmt.Errorf("could not upload route table file at %s to %s: %w", f, fn, err)
		}
		ix.routeTableToIxFile[f] = fn
	}

	if err := ix.c.UpdateIDs(ctx, ix.cfg, brps...); err != nil {
		return fmt.Errorf("could not update IDs for importing routes: %w", err)
	}
	for _, intf := range ix.intfs {
		for net, rt := range intf.netToRouteTables {
			ng := intf.netToNetworkGroup[net]
			if rt.ipv4 != "" {
				if err := importBgpRoutes(ctx, ix, ng.Ipv4PrefixPools[0].BgpIPRouteProperty[0], rt, false); err != nil {
					return fmt.Errorf("errors importing BGP IPv4 routes for network %s: %w", net, err)
				}
			}
			if rt.ipv6 != "" {
				if err := importBgpRoutes(ctx, ix, ng.Ipv6PrefixPools[0].BgpV6IPRouteProperty[0], rt, true); err != nil {
					return fmt.Errorf("errors importing BGP IPv6 routes for network %s: %w", net, err)
				}
			}
		}
	}
	return nil
}

func (ix *ixATE) updateTopology(ctx context.Context, ics []*opb.InterfaceConfig) error {
	if err := ix.configureTopology(ics); err != nil {
		return err
	}
	// Full topology updates can take some time, use a 3 minute timeout.
	if err := ix.importConfig(ctx, ix.cfg, false, topoImportTimeout, "topology and protocols configuration"); err != nil {
		return err
	}
	return syncRouteTableFilesAndImportFn(ctx, ix)
}

// PushTopology configures the IxNetwork session with the specified topology.
func (ix *ixATE) PushTopology(ctx context.Context, top *Topology) error {
	if err := validateInterfaces(top.Interfaces); err != nil {
		return err
	}
	ix.resetClientCfg()
	if err := ix.addPorts(top); err != nil {
		return err
	}
	if err := ix.importConfig(ctx, ix.cfg, true, topoImportTimeout, "port configuration"); err != nil {
		return err
	}
	if lags := top.LAGs; len(lags) > 0 {
		if err := ix.addLAGs(lags); err != nil {
			return err
		}
		if err := ix.importConfig(ctx, ix.cfg, false, topoImportTimeout, "LAG configuration"); err != nil {
			return err
		}
	}
	// Avoid a possible race condition with repeated config imports (b/191984474).
	sleepFn(45 * time.Second)
	if err := ix.updateTopology(ctx, top.Interfaces); err != nil {
		return err
	}
	ix.operState = operStateOff
	return nil
}

// UpdateTopology updates IxNetwork session to the specified topology.
func (ix *ixATE) UpdateTopology(ctx context.Context, top *Topology) error {
	if err := validateInterfaces(top.Interfaces); err != nil {
		return err
	}
	if err := ix.updateTopology(ctx, top.Interfaces); err != nil {
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

func checkUp(ctx context.Context, ix *ixATE, nodes []ixconfig.IxiaCfgNode, r stateRsp, retries int) ([]ixconfig.IxiaCfgNode, error) {
	const retryWait = 10 * time.Second
	notUp := func(nodes []ixconfig.IxiaCfgNode) ([]ixconfig.IxiaCfgNode, error) {
		var notUp []ixconfig.IxiaCfgNode
		for _, n := range nodes {
			nodeID, err := ix.c.NodeID(n)
			if err != nil {
				return nil, err
			}
			if err := ix.c.Session().Get(ctx, nodeID, r); err != nil {
				return nil, fmt.Errorf("could not fetch element at %q to check session status: %w", nodeID, err)
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

func failedNodesWithPorts(ix *ixATE, msg string, nodes []ixconfig.IxiaCfgNode, nodeToLink map[ixconfig.IxiaCfgNode]ixconfig.IxiaCfgNode) ([]string, error) {
	if len(nodes) == 0 {
		return nil, nil
	}
	var idsAndPorts []string
	failedPorts := map[string]bool{}
	for _, n := range nodes {
		nodeID, err := ix.c.NodeID(n)
		if err != nil {
			return nil, err
		}

		ports, err := linkToPorts(ix, nodeToLink[n])
		if err != nil {
			return nil, fmt.Errorf("error finding ports for link on interface for node %q: %w", nodeID, err)
		}
		for _, p := range ports {
			failedPorts[p] = true
		}
		idsAndPorts = append(idsAndPorts, fmt.Sprintf("%s (on %v)", nodeID, ports))
	}
	var failedPortsList []string
	for p := range failedPorts {
		failedPortsList = append(failedPortsList, p)
	}
	return failedPortsList, fmt.Errorf("%s: %v", msg, idsAndPorts)
}

// Waits for protocols to start, restarting any that initially fail and rechecking as necessary.
// Returns a list of ports where protocols did not start (if there are no other errors).
func validateProtocolStart(ctx context.Context, ix *ixATE) ([]string, error) {
	const (
		retryWait      = 10 * time.Second
		maxRetriesLSPs = 60
	)
	maxRetriesProtocols := 5
	if len(ix.routeTableToIxFile) > 0 {
		// Protocols may need extra time to start on large-scale route imports.
		maxRetriesProtocols = 30
	}
	topo := struct {
		ProtocolActionsInProgress []string
		Status                    string
	}{}
	started := func() (bool, error, bool) {
		if err := ix.c.Session().Get(ctx, "globals/topology", &topo); err != nil {
			return false, fmt.Errorf("could not fetch global topology to check protocol status: %w", err), false
		}
		if topo.Status == "started" && len(topo.ProtocolActionsInProgress) == 0 {
			return true, nil, false
		}
		// TODO(team): Remove IS-IS/Lag check after Ixias updated to 9.20update1
		for _, intf := range ix.intfs {
			// Checks for the known index out of bounds error if there is any IS-IS config on any LAG.
			// Still returns 'false' as the 'started' result for the function to trigger the retry
			// loop to ensure protocols start.
			if strings.Contains(intf.link.XPath().String(), "lag") && len(intf.isrToNetworkGroup) > 0 {
				var appErrs []struct{ Name string }
				if err := ix.c.Session().Get(ctx, "globals/appErrors/error", &appErrs); err != nil {
					return false, fmt.Errorf("could not fetch global errors: %w", err), false
				}
				for _, e := range appErrs {
					if strings.Contains(e.Name, "Index was outside the bounds of the array") {
						log.Warning("Ignoring protocol status error because of known 'Index was outside the bounds of the array' issue configuring IS-IS on LAGs")
						return false, nil, true
					}
				}
			}
		}
		return false, nil, false
	}
	// Wait for protocols to start.
	start, err, isrAndLagErr := started()
	for i := 0; i < maxRetriesProtocols && err == nil && !start; i++ {
		sleepFn(retryWait)
		start, err, isrAndLagErr = started()
	}
	if err != nil {
		return nil, err
	}
	if !start && !isrAndLagErr {
		return nil, errors.New("protocols did not start in time")
	}

	var protocolNodes []ixconfig.IxiaCfgNode
	var ingressLSPNodes []ixconfig.IxiaCfgNode
	nodeToLink := map[ixconfig.IxiaCfgNode]ixconfig.IxiaCfgNode{}
	for _, intf := range ix.intfs {
		nodeToLink[intf.deviceGroup.Ethernet[0]] = intf.link
		protocolNodes = append(protocolNodes, intf.deviceGroup.Ethernet[0])
		if intf.ipv4 != nil {
			nodeToLink[intf.ipv4] = intf.link
			protocolNodes = append(protocolNodes, intf.ipv4)
		}
		if intf.ipv6 != nil {
			nodeToLink[intf.ipv6] = intf.link
			protocolNodes = append(protocolNodes, intf.ipv6)
		}
		for _, lsps := range intf.rsvpLSPs {
			if *(lsps.RsvpP2PIngressLsps.Active.SingleValue.Value) == "true" {
				nodeToLink[lsps.RsvpP2PIngressLsps] = intf.link
				ingressLSPNodes = append(ingressLSPNodes, lsps.RsvpP2PIngressLsps)
			}
		}
	}
	var cfgNodes []ixconfig.IxiaCfgNode
	cfgNodes = append(cfgNodes, protocolNodes...)
	cfgNodes = append(cfgNodes, ingressLSPNodes...)
	if err := ix.c.UpdateIDs(ctx, ix.cfg, cfgNodes...); err != nil {
		return nil, fmt.Errorf("could not update IDs for checking protocol sessions: %w", err)
	}

	protocolNodes, err = checkUp(ctx, ix, protocolNodes, &protocolRsp{}, maxRetriesProtocols)
	if err != nil {
		return nil, err
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
			return nil, fmt.Errorf("tried to restart invalid config node type %T", cn)
		}
		nodeID, err := ix.c.NodeID(n)
		if err != nil {
			return nil, err
		}
		if err := ix.runOp(ctx, op, ixweb.OpArgs{[]string{nodeID}}, nil); err != nil {
			return nil, fmt.Errorf("could not restart down protocol at %q", nodeID)
		}
	}
	protocolNodes, err = checkUp(ctx, ix, protocolNodes, &protocolRsp{}, maxRetriesProtocols)
	if err != nil {
		return nil, err
	}
	if ports, err := failedNodesWithPorts(ix, "some protocol instances did not start", protocolNodes, nodeToLink); err != nil {
		return ports, err
	}

	// Validate ingress LSP state
	ingressLSPNodes, err = checkUp(ctx, ix, ingressLSPNodes, &lspRsp{}, maxRetriesLSPs)
	if err != nil {
		return nil, err
	}
	return failedNodesWithPorts(ix, "some ingress LSP instances did not start", ingressLSPNodes, nodeToLink)
}

func (ix *ixATE) checkPortsUp(ctx context.Context) error {
	const upState = "connectedLinkUp"
	var vports []ixconfig.IxiaCfgNode
	for _, vport := range ix.ports {
		vports = append(vports, vport)
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, vports...); err != nil {
		return err
	}
	for port, vport := range ix.ports {
		vpid, err := ix.c.NodeID(vport)
		if err != nil {
			return err
		}
		var data struct {
			ConnectionState string `json:"connectionState"`
		}
		if err := ix.c.Session().Get(ctx, vpid, &data); err != nil {
			return err
		}
		if data.ConnectionState != upState {
			return fmt.Errorf("port %q in state %q, not %q", port, data.ConnectionState, upState)
		}
	}
	return nil
}

func (ix *ixATE) startProtocols(ctx context.Context) error {
	if err := ix.checkPortsUp(ctx); err != nil {
		return err
	}
	errStart := ix.runOp(ctx, "operations/startallprotocols", syncedOpArgs, nil)
	if errStart != nil {
		log.Warningf("First attempted startallprotocols op failed: %v", errStart)
		if errStart = ix.runOp(ctx, "operations/startallprotocols", syncedOpArgs, nil); errStart != nil {
			log.Warningf("Second attempted startallprotocols op failed: %v", errStart)
		}
	}
	// Protocols may have started even if 'startallprotocols' reported a failure,
	// so return based on actual protocol state.
	// TODO(team): This may be related to behavior on 9.12, try reverting
	// to failing tests on a single 'startallprotocols' error and check the
	// behavior for a topology with RSVP protocols configured after updating to
	// IxNetwork 9.10update2 everywhere.
	if _, err := validateProtocolStartFn(ctx, ix); err != nil {
		sessErrs := ix.sessionErrors(ctx)
		var errBuilder strings.Builder
		for i, e := range sessErrs {
			errBuilder.WriteString(fmt.Sprintf("name: %q, description: %q, instances: ", e.Name, e.Description))
			for j, row := range e.InstanceRowValues {
				errBuilder.WriteString(fmt.Sprintf("%v", row))
				if j != len(e.InstanceRowValues)-1 {
					errBuilder.WriteString(", ")
				}
			}
			if i != len(sessErrs)-1 {
				errBuilder.WriteRune('\n')
			}
		}
		sErr := errBuilder.String()
		if sErr != "" {
			sErr = fmt.Sprintf("\n%s", sErr)
		}
		if errStart == nil {
			return fmt.Errorf("failed to validate protocols: %w%s", err, sErr)
		}
		return fmt.Errorf("failed to validate protocols after error on start: %v, %w%s", err, errStart, sErr)
	}
	log.Infof("Protocols started successfully")
	return nil
}

// StartProtocols starts running protocols for the IxNetwork session.
func (ix *ixATE) StartProtocols(ctx context.Context) error {
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
func (ix *ixATE) StopProtocols(ctx context.Context) error {
	if ix.operState == operStateOff {
		log.Infof("Protocols already stopped, not running operation on Ixia.")
		return nil
	}
	if err := ix.runOp(ctx, "operations/stopallprotocols", syncedOpArgs, nil); err != nil {
		return fmt.Errorf("could not stop protocols: %w", err)
	}
	ix.operState = operStateOff
	return nil
}

// SetPortState enables/disables the given Ixia port.
func (ix *ixATE) SetPortState(ctx context.Context, port string, enabled *bool) error {
	if port == "" {
		return fmt.Errorf("no port provided in SetPortState action on ATE %q", ix.name)
	}
	if enabled == nil {
		return fmt.Errorf("no enabled state provided in SetPortState action on ATE %q", ix.name)
	}
	vport, ok := ix.ports[port]
	if !ok {
		return fmt.Errorf("port %q does not exist in current configuration", port)
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, vport); err != nil {
		return fmt.Errorf("could not fetch ID for vport for %q: %w", port, err)
	}

	state := "down"
	if *enabled {
		state = "up"
	}
	vportID, err := ix.c.NodeID(vport)
	if err != nil {
		return err
	}
	if err := ix.runOp(ctx, "vport/operations/linkupdn", ixweb.OpArgs{[]string{vportID}, state}, nil); err != nil {
		return fmt.Errorf("error setting port state for %q: %w", port, err)
	}
	return nil
}

// SetLACPState enables/disables LACP on the given Ixia port in a LAG.
func (ix *ixATE) SetLACPState(ctx context.Context, port string, enabled *bool) error {
	if port == "" {
		return fmt.Errorf("no port provided in SetLACPState action on ATE %q", ix.name)
	}
	if enabled == nil {
		return fmt.Errorf("no enabled state provided in SetLACPState action on ATE %q", ix.name)
	}
	op := stopLACPOp
	if *enabled {
		op = startLACPOp
	}

	vport, ok := ix.ports[port]
	if !ok {
		return fmt.Errorf("port %q does not exist in current configuration", port)
	}
	var ixLAG *ixconfig.Lag
	var portIdx int
	for l, vports := range ix.lagPorts {
		for i, p := range vports {
			if p == vport {
				ixLAG = l
				portIdx = i
			}
		}
	}
	if ixLAG == nil {
		return fmt.Errorf("port %q is not a member of a LAG", port)
	}

	lacpNode := ixLAG.ProtocolStack.Ethernet[0].Lagportlacp[0]
	if err := ix.c.UpdateIDs(ctx, ix.cfg, lacpNode); err != nil {
		return fmt.Errorf("could not fetch ID for vport for %q: %w", port, err)
	}
	lacpID, err := ix.c.NodeID(lacpNode)
	if err != nil {
		return err
	}
	portID := fmt.Sprintf("%s/port/%d", lacpID, portIdx+1)
	if err := ix.runOp(ctx, op, ixweb.OpArgs{[]string{portID}}, nil); err != nil {
		return fmt.Errorf("error setting LACP state for %q: %w", port, err)
	}
	return nil
}

func resolveMacs(ctx context.Context, ix *ixATE) error {
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
		return fmt.Errorf("could not fetch IDs for resolving MACs: %w", err)
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
		macID, err := ix.c.NodeID(eth.Mac)
		if err != nil {
			return err
		}
		var macs []string
		err = ix.runOp(ctx, "multivalue/operations/getvalues", ixweb.OpArgs{macID, 0, 1}, &macs)
		if err != nil {
			return fmt.Errorf("could not fetch value for MAC multivalue %q: %w", macID, err)
		}
		intf.ethMac = firstValidMac(macs)

		if intf.ipv4 != nil {
			ipv4ID, err := ix.c.NodeID(intf.ipv4)
			if err != nil {
				return err
			}
			if err := ix.c.Session().Get(ctx, ipv4ID, &ipData); err != nil {
				return fmt.Errorf("could not fetch IPv4 object at %q: %w", ipv4ID, err)
			}
			if mac := firstValidMac(ipData.ResolvedGatewayMAC); mac != "" {
				intf.resolvedIpv4Mac = mac
			}
		}

		if intf.ipv6 != nil {
			ipv6ID, err := ix.c.NodeID(intf.ipv6)
			if err != nil {
				return err
			}
			if err := ix.c.Session().Get(ctx, ipv6ID, &ipData); err != nil {
				return fmt.Errorf("could not fetch IPv6 object at %q: %w", ipv6ID, err)
			}
			if mac := firstValidMac(ipData.ResolvedGatewayMAC); mac != "" {
				intf.resolvedIpv6Mac = mac
			}
		}
	}
	return nil
}

// Expects all traffic items to have up-to-date REST IDs.
func genTraffic(ctx context.Context, ix *ixATE) error {
	if ix.cfg.Traffic == nil || len(ix.cfg.Traffic.TrafficItem) == 0 {
		return nil
	}
	var tiIDs []string
	for _, ti := range ix.cfg.Traffic.TrafficItem {
		tiID, err := ix.c.NodeID(ti)
		if err != nil {
			return err
		}
		tiIDs = append(tiIDs, tiID)
	}
	if err := ix.runOp(ctx, "traffic/trafficItem/operations/generate", ixweb.OpArgs{tiIDs}, nil); err != nil {
		return fmt.Errorf("could not generate traffic flows: %w", err)
	}
	return nil
}

// updateFlows updates frame size/rate configuration for flows after generation.
// Assumes that IxNetwork traffic items corresponding to the flows have updated
// REST IDs.
func updateFlows(ctx context.Context, ix *ixATE, flows []*opb.Flow) error {
	const hlsSuffix = "highLevelStream/1"
	for _, f := range flows {
		ti := ix.flowToTrafficItem[f.GetName()]
		if ti == nil {
			return fmt.Errorf("flow %q does not exist", f.GetName())
		}
		tiID, err := ix.c.NodeID(ti)
		if err != nil {
			return err
		}
		ce := ti.ConfigElement[0]

		fs, fsMap, err := frameSize(f.GetFrameSize())
		if err != nil {
			return fmt.Errorf("could not compute new frame size for flow %q: %w", f.GetName(), err)
		}
		if fs != nil {
			if ce == nil {
				return fmt.Errorf("cannot update frame size for flow %q since it was not originally specified explicitly", f.GetName())
			}
			if *(fs.Type_) != *(ce.FrameSize.Type_) {
				return fmt.Errorf("cannot switch to frame size type %T for flow %q", f.GetFrameSize().GetType(), f.GetName())
			}
			if err := ix.c.Session().Patch(ctx, path.Join(tiID, hlsSuffix, "frameSize"), fsMap); err != nil {
				return fmt.Errorf("could not patch frame size for flow %q: %w", f.GetName(), err)
			}
		}

		_, frMap, err := frameRate(f.GetFrameRate())
		if err != nil {
			return fmt.Errorf("could not compute new frame rate for flow %q: %w", f.GetName(), err)
		}

		if frMap != nil {
			if err := ix.c.Session().Patch(ctx, path.Join(tiID, hlsSuffix, "frameRate"), frMap); err != nil {
				return fmt.Errorf("could not patch frame rate for flow %q: %w", f.GetName(), err)
			}
		}
	}
	return nil
}

func configureTraffic(ctx context.Context, ix *ixATE, flows []*opb.Flow) error {
	if err := validateFlows(flows); err != nil {
		return err
	}
	ix.resetClientTrafficCfg()
	if err := resetIxiaTrafficCfgFn(ctx, ix); err != nil {
		return fmt.Errorf("could not reset traffic config on Ixia: %w", err)
	}
	if err := resolveMacsFn(ctx, ix); err != nil {
		return fmt.Errorf("could not resolve MAC addresses in config: %w", err)
	}

	for _, f := range flows {
		if f.GetConvergenceTracking() {
			ix.convergenceTracking = true
			// Update latency/convergence tracking configuration. This can't be done via config push
			// because convergence tracking will be automatically unset if latency tracking is not
			// separately disabled first.
			if err := ix.c.Session().Patch(ctx, "/traffic/statistics/latency", map[string]any{"enabled": false}); err != nil {
				return fmt.Errorf("could not disable latency statistics: %w", err)
			}
			if err := ix.c.Session().Patch(ctx, "/traffic/statistics/cpdpConvergence", map[string]any{
				"enabled":                          true,
				"enableControlPlaneEvents":         true,
				"enableDataPlaneEventsRateMonitor": true,
			}); err != nil {
				return fmt.Errorf("could not configure convergence measurements: %w", err)
			}
			break
		}
	}

	if err := ix.addTraffic(flows); err != nil {
		return fmt.Errorf("could not compute traffic configuration: %w", err)
	}

	// There is a race condition in configuring traffic after clearing traffic, sleep for 30 seconds to avoid.
	sleepFn(30 * time.Second)
	if err := ix.importConfig(ctx, ix.cfg.Traffic, false, trafficImportTimeout, "traffic configuration"); err != nil {
		return fmt.Errorf("could not push traffic configuration: %w", err)
	}

	var items []ixconfig.IxiaCfgNode
	for _, ti := range ix.cfg.Traffic.TrafficItem {
		items = append(items, ti)
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, items...); err != nil {
		return fmt.Errorf("could not update IDs for traffic items: %w", err)
	}

	if err := genTrafficFn(ctx, ix); err != nil {
		return fmt.Errorf("could not generate traffic flow: %w", err)
	}

	// IxNetwork does not necessarily configure frame rate correctly when generating
	// traffic streams (b/204318369).
	if err := updateFlowsFn(ctx, ix, flows); err != nil {
		return fmt.Errorf("could not update generated flows: %w", err)
	}
	return nil
}

func applyTraffic(ctx context.Context, ix *ixATE) error {
	trafficArgs := ixweb.OpArgs{ix.c.Session().AbsPath("traffic")}
	if err := ix.runOp(ctx, "traffic/operations/apply", trafficArgs, nil); err != nil {
		return fmt.Errorf("could not apply traffic config: %w", err)
	}
	return nil
}

func startTraffic(ctx context.Context, ix *ixATE) error {
	if len(ix.egressTrackFlowCounts) > 0 {
		var maxCount int
		var egressTrackFlows []string
		for flow, count := range ix.egressTrackFlowCounts {
			egressTrackFlows = append(egressTrackFlows, flow)
			if count > maxCount {
				maxCount = count
			}
		}
		if _, err := ix.c.Session().Stats().ConfigEgressView(ctx, egressTrackFlows, maxCount); err != nil {
			return fmt.Errorf("could not configure egress tracking view: %w", err)
		}
	}
	trafficArgs := ixweb.OpArgs{ix.c.Session().AbsPath("traffic")}
	if err := ix.runOp(ctx, "traffic/operations/start", trafficArgs, nil); err != nil {
		return fmt.Errorf("could not start traffic: %w", err)
	}
	return nil
}

// StartTraffic starts traffic for the IxNetwork session based on the given flows.
// If no flows are provided, starts the previously pushed flows.
func (ix *ixATE) StartTraffic(ctx context.Context, flows []*opb.Flow) error {
	if ix.operState == operStateOff {
		return fmt.Errorf("cannot start traffic before starting protocols")
	}
	if ix.operState == operStateTrafficOn {
		log.Infof("Traffic already running, not running operation on Ixia.")
		return nil
	}

	// Configure and apply traffic only when flows are provided.
	if len(flows) > 0 {
		if err := configureTrafficFn(ctx, ix, flows); err != nil {
			return fmt.Errorf("could not configure traffic: %w", err)
		}
		if err := applyTrafficFn(ctx, ix); err != nil {
			return fmt.Errorf("could not apply traffic: %w", err)
		}
	} else if len(ix.flowToTrafficItem) == 0 {
		// No flows have been configured since the last full topology push.
		return errors.New("cannot start traffic, no flows defined")
	}

	if err := startTrafficFn(ctx, ix); err != nil {
		return fmt.Errorf("could not start traffic: %w", err)
	}

	ix.operState = operStateTrafficOn
	return nil
}

// UpdateTraffic updates traffic config for the IxNetwork session based on the given flows.
func (ix *ixATE) UpdateTraffic(ctx context.Context, flows []*opb.Flow) error {
	if err := validateFlows(flows); err != nil {
		return err
	}
	if ix.operState != operStateTrafficOn {
		return fmt.Errorf("cannot update traffic before it has been started")
	}
	if err := updateFlowsFn(ctx, ix, flows); err != nil {
		return fmt.Errorf("could not update running traffic flows: %w", err)
	}
	return nil
}

func (ix *ixATE) stopAllTraffic(ctx context.Context) error {
	trafficArgs := ixweb.OpArgs{ix.c.Session().AbsPath("traffic")}
	if err := ix.runOp(ctx, "traffic/operations/stop", trafficArgs, nil); err != nil {
		return fmt.Errorf("could not stop traffic: %w", err)
	}
	// Wait a sufficient amount of time to ensure that traffic is stopped.
	sleepFn(15 * time.Second)
	return nil
}

func (ix *ixATE) trafficItemStatsAvailable(ctx context.Context) bool {
	views, statErr := ix.c.Session().Stats().Views(ctx)
	if statErr != nil {
		log.Warningf("Could not fetch list ot statistic views: %v", statErr)
	}
	_, avail := views[ixweb.TrafficItemStatsCaption]
	return avail
}

// StopAllTraffic stops all traffic for the IxNetwork session and waits for stats to be populated.
func (ix *ixATE) StopAllTraffic(ctx context.Context) error {
	const (
		retryWait       = 5 * time.Second
		maxRetriesStats = 6
	)
	if ix.operState != operStateTrafficOn {
		log.Infof("Traffic already stopped, not running operation on Ixia.")
		return nil
	}
	if err := ix.stopAllTraffic(ctx); err != nil {
		return err
	}
	ix.operState = operStateProtocolsOn
	log.Infof("Traffic stopped, waiting for stats to populate.")

	statsUpdated := ix.trafficItemStatsAvailable(ctx)
	for i := 0; i < maxRetriesStats && !statsUpdated; i++ {
		sleepFn(retryWait)
		statsUpdated = ix.trafficItemStatsAvailable(ctx)
	}
	if !statsUpdated {
		return errors.New("traffic item statistics did not become available after stopping traffic")
	}
	return nil
}

// FetchGNMI returns the GNMI client for the Ixia.
func (ix *ixATE) FetchGNMI(ctx context.Context) (gpb.GNMIClient, error) {
	ix.mu.Lock()
	defer ix.mu.Unlock()
	if ix.gclient == nil {
		gclient, err := ixgnmi.NewClient(ctx, ix.name, ix.readStats, unwrapClient(ix.c), rawapis.CommonDialOpts...)
		if err != nil {
			return nil, err
		}
		ix.gclient = gclient
	}
	return ix.gclient, nil
}

func (ix *ixATE) readStats(ctx context.Context, captions []string) (*ixgnmi.Stats, error) {
	if len(ix.egressTrackFlowCounts) == 0 {
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
			var drilldowns []string
			if cap == ixweb.TrafficItemStatsCaption && ix.convergenceTracking {
				drilldowns = []string{"Drill down per Dest Endpoint", "Drill Down per Rx Port"}
			}
			table, err = view.FetchTable(ctx, drilldowns...)
			if err != nil {
				return nil, err
			}
		} else {
			log.Warningf("No view with caption %q, got views %v", cap, views)
		}
		tables[cap] = table
	}
	if err != nil {
		return nil, fmt.Errorf("error retrieving statistics for views %v: %w", captions, err)
	}
	var egressTrackFlows []string
	for flow := range ix.egressTrackFlowCounts {
		egressTrackFlows = append(egressTrackFlows, flow)
	}
	return &ixgnmi.Stats{
		Tables:            tables,
		IngressTrackFlows: ix.ingressTrackFlows,
		EgressTrackFlows:  egressTrackFlows,
	}, nil
}

// FlushStats will remove all data from the cache for the Ixia and marks all stats as stale.
// This ensures that future telemetry calls attempt to fetch data and that stale data doesn't persist.
func (ix *ixATE) FlushStats() {
	ix.mu.Lock()
	defer ix.mu.Unlock()
	if ix.gclient != nil {
		ix.gclient.Flush()
	}
}

// UpdateBGPPeerStates exists only to match the API of the prior IxNetwork ATE binding.
// It assumes that the only changes in the provided interface configs are updates to
// BGP active states.
// TODO(team): Remove this method once new Ixia config binding is used.
func (ix *ixATE) UpdateBGPPeerStates(ctx context.Context, ifs []*opb.InterfaceConfig) error {
	if err := validateInterfaces(ifs); err != nil {
		return err
	}
	if err := ix.configureTopology(ifs); err != nil {
		return err
	}
	for _, intf := range ix.intfs {
		if v4 := intf.ipv4; v4 != nil {
			for _, peer := range v4.BgpIpv4Peer {
				if err := ix.importConfig(ctx, peer, false, peersImportTimeout, "BGP v4 peer update"); err != nil {
					return fmt.Errorf("could not update BGP v4 peer: %w", err)
				}
			}
		}
		if v6 := intf.ipv6; v6 != nil {
			for _, peer := range v6.BgpIpv6Peer {
				if err := ix.importConfig(ctx, peer, false, peersImportTimeout, "BGP v6 peer update"); err != nil {
					return fmt.Errorf("could not update BGP v6 peer: %w", err)
				}
			}
		}
		if v4lb := intf.ipv4Loopback; v4lb != nil {
			for _, peer := range v4lb.BgpIpv4Peer {
				if err := ix.importConfig(ctx, peer, false, peersImportTimeout, "BGP v4 loopback peer update"); err != nil {
					return fmt.Errorf("could not update BGP v4 loopback peer: %w", err)
				}
			}
		}
		if v6lb := intf.ipv6Loopback; v6lb != nil {
			for _, peer := range v6lb.BgpIpv6Peer {
				if err := ix.importConfig(ctx, peer, false, peersImportTimeout, "BGP v6 loopback peer update"); err != nil {
					return fmt.Errorf("could not update BGP v6 loopback peer: %w", err)
				}
			}
		}
	}
	return ix.applyOnTheFly(ctx)
}

func (ix *ixATE) applyOnTheFly(ctx context.Context) error {
	const (
		applyOnTheFlyArg = "globals/topology"
		applyOnTheFlyOp  = "globals/topology/operations/applyonthefly"
	)
	applyOnTheFlyArgs := ixweb.OpArgs{ix.c.Session().AbsPath(applyOnTheFlyArg)}
	if err := ix.runOp(ctx, applyOnTheFlyOp, applyOnTheFlyArgs, nil); err != nil {
		return fmt.Errorf("could not apply topology changes: %w", err)
	}
	return nil
}

func (ix *ixATE) UpdateNetworkGroups(ctx context.Context, ifs []*opb.InterfaceConfig) error {
	if err := validateInterfaces(ifs); err != nil {
		return err
	}
	if err := ix.configureTopology(ifs); err != nil {
		return err
	}
	for _, intf := range ix.intfs {
		for _, ng := range intf.netToNetworkGroup {
			if ng == nil {
				continue
			}
			if err := ix.importConfig(ctx, ng, false, peersImportTimeout, "network groups update"); err != nil {
				return fmt.Errorf("could not update network groups: %w", err)
			}
		}
	}
	return ix.applyOnTheFly(ctx)
}

func (ix *ixATE) SendBGPGracefulRestart(ctx context.Context, peerIDs []uint32, delay time.Duration) error {
	if len(peerIDs) == 0 {
		return fmt.Errorf("no bgp peers provided")
	}
	var peerNodes []ixconfig.IxiaCfgNode
	for _, intf := range ix.intfs {
		for _, id := range peerIDs {
			if peer, ok := intf.idToBGPv4Peer[id]; ok {
				peerNodes = append(peerNodes, peer)
			}
			if peer, ok := intf.idToBGPv6Peer[id]; ok {
				peerNodes = append(peerNodes, peer)
			}
		}
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, peerNodes...); err != nil {
		return fmt.Errorf("error updating IDs for bgp peers: %w", err)
	}
	for _, node := range peerNodes {
		nodeID, err := ix.c.NodeID(node)
		if err != nil {
			return fmt.Errorf("error getting ID for bgp peer: %w", err)
		}
		var op string
		switch n := node.(type) {
		case *ixconfig.TopologyBgpIpv4Peer:
			op = bgpPeerV4GracefulRestartOp
		case *ixconfig.TopologyBgpIpv6Peer:
			op = bgpPeerV6GracefulRestartOp
		default:
			return fmt.Errorf("invalid BGP peer node type: %v(%T)", n, node)
		}
		args := ixweb.OpArgs{nodeID, []int{1}, uint32(delay.Seconds())}
		if err := ix.c.Session().Post(ctx, op, args, nil); err != nil {
			return fmt.Errorf("error sending BGP graceful restart: %w", err)
		}
	}
	return nil
}

func (ix *ixATE) SendBGPPeerNotification(ctx context.Context, peerIDs []uint32, code int, subCode int) error {
	if len(peerIDs) == 0 {
		return fmt.Errorf("no bgp peers provided")
	}
	var peerNodes []ixconfig.IxiaCfgNode
	for _, intf := range ix.intfs {
		for _, id := range peerIDs {
			if peer, ok := intf.idToBGPv4Peer[id]; ok {
				peerNodes = append(peerNodes, peer)
			}
			if peer, ok := intf.idToBGPv6Peer[id]; ok {
				peerNodes = append(peerNodes, peer)
			}
		}
	}
	if err := ix.c.UpdateIDs(ctx, ix.cfg, peerNodes...); err != nil {
		return fmt.Errorf("error updating IDs for bgp peers: %w", err)
	}
	for _, node := range peerNodes {
		nodeID, err := ix.c.NodeID(node)
		if err != nil {
			return fmt.Errorf("error getting ID for bgp peer: %w", err)
		}
		var op string
		switch n := node.(type) {
		case *ixconfig.TopologyBgpIpv4Peer:
			op = bgpPeerV4NotifyOp
		case *ixconfig.TopologyBgpIpv6Peer:
			op = bgpPeerV6NotifyOp
		default:
			return fmt.Errorf("invalid BGP peer node type: %v(%T)", n, node)
		}
		args := ixweb.OpArgs{nodeID, []int{1}, code, subCode}
		if err := ix.c.Session().Post(ctx, op, args, nil); err != nil {
			return fmt.Errorf("error sending BGP notification: %w", err)
		}
	}
	return nil
}

func validateFlows(fs []*opb.Flow) error {
	for _, f := range fs {
		if len(f.GetSrcEndpoints()) == 0 {
			return fmt.Errorf("flow has no src endpointd")
		}
		if len(f.GetDstEndpoints()) == 0 {
			return fmt.Errorf("flow has no dst endpoints")
		}
	}
	return nil
}

func validateInterfaces(ifs []*opb.InterfaceConfig) error {
	if len(ifs) == 0 {
		return fmt.Errorf("zero interfaces to configure, need at least one")
	}
	if len(ifs) > maxIntfs {
		return fmt.Errorf("%v interfaces to configure, must be at most %v", len(ifs), maxIntfs)
	}
	intfs := make(map[string]bool)

	for _, i := range ifs {
		if i.GetPort() == "" && i.GetLag() == "" {
			return fmt.Errorf("interface has no port or lag specified: %v", i)
		}
		if i.GetLag() != "" && i.GetEnableLacp() {
			return fmt.Errorf("interface should not specify both a LAG and that LACP is enabled: %v", i)
		}
		if intfs[i.GetName()] {
			return fmt.Errorf("duplicate interface name: %s", i.GetName())
		}
		intfs[i.GetName()] = true
		nets := make(map[string]bool)
		for _, n := range i.GetNetworks() {
			if nets[n.GetName()] {
				return fmt.Errorf("duplicate network name: %s", n.GetName())
			}
			nets[n.GetName()] = true
		}
		if err := validateIP(i.GetIpv4(), "ipv4 on "+i.GetName()); err != nil {
			return err
		}
		if err := validateIP(i.GetIpv6(), "ipv6 on "+i.GetName()); err != nil {
			return err
		}
	}
	return nil
}

func validateIP(ipc *opb.IpConfig, desc string) error {
	if ipc == nil {
		return nil
	}
	addr := ipc.GetAddressCidr()
	gway := ipc.GetDefaultGateway()
	_, an, err := net.ParseCIDR(addr)
	if err != nil {
		return fmt.Errorf("%s address is not valid CIDR notation: %s", desc, addr)
	}
	gi := net.ParseIP(gway)
	if gi == nil {
		return fmt.Errorf("%s default gateway is not valid IP notation: %s", desc, gway)
	}
	if !gi.IsUnspecified() && !an.Contains(gi) {
		return fmt.Errorf("%s default gateway is not in CIDR range %s: %s", desc, addr, gway)
	}
	return nil
}
