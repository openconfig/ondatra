# Ondatra gNMI API

The Ondatra gNMI API provides test helpers that wrap the API calls to the
[ygnmi library](https://github.com/openconfig/ygnmi/). Ondatra uses ygnmi (and
by extension [ygot](https://github.com/openconfig/ygot)) to autogenerate gNMI
paths and value structs from an input set of YANG modules. Please read the
[ygnmi README](https://github.com/openconfig/ygnmi#readme) to learn more about
how ygnmi translates gNMI paths to Go API calls and how it handles noncompliance
errors in telemetry it receives.

## Available gNMI Paths

Ondatra provide a combination of OpenConfig and Open Traffic Generator YANG
modules as input to the ygnmi auto-generation.

The OpenConfig paths are documented here:

*   https://openconfig.net/projects/models/paths/
*   https://openconfig.net/projects/models/schemadocs/

Information on the Open Traffic Generator YANG modules is available here:

*   https://github.com/open-traffic-generator/models-yang

If you are using the IxNetwork (aka "ATE") API, see the
[IxNetwork Telemetry](#ixnetwork-telemetry) section below to learn which
OpenConfig paths are supported by our IxNetwork gNMI implementation.

## Best Practice: Avoid `time.Sleep`

A test often defines a goal state based on telemetry and waits for the device to
reach that state. It may be tempting to add a `time.Sleep` call to insert such a
wait into the test, but sleeping is both flaky (the sleep may be too short) and
wasteful (the sleep may be unnecessarily long). Tests can more reliably and
efficiently wait for a particular telemetry state by using `gnmi.Await` or
`gnmi.Watch`.

*   **Example 1: Wait to an interface to be up**

    Use `Await` to wait until a single interface is up:

    ```go
    gnmi.Await(t, dut, gnmi.OC().Interface(dp.Name()).OperStatus().State(), time.Minute, oc.Interface_OperStatus_UP)
    ```

*   **Example 2: Wait for at least 100 packets to be sent**

    Use `Watch` to wait until a DUT sends at least 100 packets. Using `Await`
    wouldn't be wise here, because the counter may never register exactly 100.

    ```go
    watch := gnmi.Watch(t, dut, gnmi.OC().Interface(dp.Name()).Counters().OutPkts().State(), time.Minute, func(val *ygnmi.Value[uint64]) bool {
        count, ok := val.Val()
        return ok && count > 100
    })
    if val, ok := watch.Await(t); !ok {
        t.Fatalf("DUT did not reach target state: got %v", val)
    }
    ```

*   **Example 3: Wait for several interfaces to be up**

    Use a batch query with `Watch` to wait for several interfaces to be up in
    parallel. Using multiple `Await` calls wouldn't be as efficient, because
    that would cause the test to wait for each interface in serial.

    ```go
    batch := gnmi.OCBatch()
    for _, port := range dut.Ports() {
        batch.AddPaths(gnmi.OC().Interface(port.Name()))
    }
    watch := gnmi.Watch(t, dut, batch.State(), time.Minute, func(val *ygnmi.Value[*oc.Root]) bool {
        root, _ := val.Val()
        for _, port := range dut.Ports() {
            if root.GetInterface(port.Name()).GetOperStatus() != oc.Interface_OperStatus_UP {
                return false
            }
        }
        return true
    })
    if val, ok := watch.Await(t); !ok {
        t.Fatalf("DUT did not reach target state: got %v", val)
    }
    ```

## IxNetwork Telemetry

When using Ondatra's IxNetwork (aka "ATE") API, a test can gather statistics and
protocol data about the IxNetwork session via gNMI. This section describes the
available IxNetwork telemetry and the corresponding OpenConfig paths at which it
can be queried. As the IxNetwork telemetry requires a conversion from a REST API
to gNMI, you may encounter some necessary limitations in its behavior.

### Statistics Views

IxNetwork statistics views are exported to gNMI paths according to the following
table:

Ixia Stats         | gNMI API Path
------------------ | ----------------------------------------------
Port Stats         | gnmi.OC().Interface($PortName)
Port CPU Stats     | gnmi.OC().Component($PortName).Cpu()
Flow Stats         | gnmi.OC().Flow($FlowName)
Ingress Flow Stats | gnmi.OC().Flow($FlowName).IngressTrackingAny()
Egress Flow Stats  | gnmi.OC().Flow($FlowName).EgressTrackingAny()

### Routing Data

Learned routing data is exported under an OpenConfig "network instance" with the
same name as the interface over which the data was learned. Specifically,
routing data is available under gNMI paths that start
`gnmi.OC().NetworkInstance($InterfaceName)`, where `$InterfaceName` is the
argument to an `ate.AddInterface($InterfaceName)` call in the test.

The following routing data is available via gNMI:

*   BGP IPv4 RIB

    ```
    gnmi.OC().NetworkInstance($InterfaceName).
    Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
    AfiSafi(oc.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST).Ipv4Unicast().
    Neighbor($DutIp).AdjRibInPre().Route($Prefix, $PathId)
    ```

*   BGP IPv6 RIB

    ```
    gnmi.OC().NetworkInstance($InterfaceName).
    Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
    AfiSafi(oc.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST).Ipv6Unicast().
    Neighbor($DutIp).AdjRibInPre().Route($Prefix, $PathId)
    ```

*   BGP Attr Sets

    ```
    gnmi.OC().NetworkInstance($InterfaceName).
    Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
    AttrSet($Index)
    ```

*   BGP Communities

    ```
    gnmi.OC().NetworkInstance($InterfaceName).
    Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
    Community($Index)
    ```

*   ISIS LSPs

    ```
    gnmi.OC().NetworkInstance($InterfaceName).
    Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS, "0").Isis().
    Level($LevelNum).Lsp($LspId)
    ```

*   RSVP-TE Sessions

    ```
    gnmi.OC().NetworkInstance($InterfaceName).Mpls().SignalingProtocols().
    RsvpTe().SessionId($LocalIndex)
    ```
