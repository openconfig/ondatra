# IxNetwork JSON Config Client Usage

The IxNetwork JSON config client provides a way to control IxNetwork sessions
using a JSON representation of the complete test session configuration, instead
of making multiple REST API calls to modify particular test objects/features.
This document describes basic usage of the ixconfig. It does not cover how to
implement specific protocols or the design/development of the client itself.

## IxNetwork Config Model

IxNetwork JSON configs are a representation of the IxNetwork config model, a
hierarchical structure of configuration objects that define any particular test
scenario. Each object in the JSON config has an XPath, a unique identifier
pointing to its location in the hierarchy.

The next sections describe the config objects that define the networking setup
of a test session: vports, topologies, and traffic items. A more thorough
introduction to these (and how to interact with them through the IxNetwork Web
UI) can be found in Ixia's own
[tutorial](https://www.youtube.com/watch?v=A0mbZuP94jo).

### VPorts

VPorts represent interfaces that are connected to devices under test. They are
associated with exactly one physical port or one LAG of physical ports on the
Ixia chassis. Other config objects reference the vport abstractions instead of
directly specifying physical ports.

### Topologies

A topology is a block of vports (at least one) configured with a particular set
of protocols. A vport can be associated with at most one topology.

Within a topology are device groups (at least one) that define virtual devices
with a similar role (e.g. a host or a router.) Each device group configuration
is applied to every port in the topology.

Every device group has a specific protocol stack configured with a device
multiplier that determines the number of devices that appear as part of the
group. Device groups may also contain network groups representing routes that
are advertised through the protocols configured for the group.

### Traffic Items

Traffic items are configured on one or more vports. They represent traffic flows
used to test network devices. Traffic is configured to be generated from source
ports on the Ixia and measurements are configured on destination ports. A port
may be both a source and destination for traffic flows.

## Creating a Client Connection

Using the JSON config client requires that you have an `ixweb.Session` instance
connected to an active IxNetwork session. You can then create/test the config
client connection as follows:

```go
import (
  ...
  "github.com/openconfig/ondatra/internal/ixconfig"
  ...
)
...
  var sess *ixweb.Session
  ...
  client = ixconfig.New(sess)
  cfg, err := client.GetConfig()
  if err != nil {
    return err
  }
  // Check that the config is non-empty to validate your connection. Not necessary in general.
  if cfg.XPath().String() != "/" {
    return fmt.Errorf("Invalid config returned from IxNetwork session %v.", sess)
  }
...
```

## Basic Configuration

The following sections describe the steps to set a minimal configuration for an
IxNetwork session, assuming that a client has been created and a config has been
fetched from a fresh session as above.

Notes on the following example:

*   Assigning names to configuration objects is not strictly necessary, but
    helps visibility when debugging through the Web UI.
*   Any fields that are not set in a config object will be set (by the IxNetwork
    server) to their default value as defined by the IxNetwork REST API.
*   The order of construction does not matter, only the final config that is
    pushed to the IxNetwork session. (In the example, only the xpath update and
    vport attachment to the topology depend on prior steps.)

### Port Assignment Config

Assign physical ports to vport objects to configure protocols/traffic flows to
run on/through them. For each port you wish to use, reference it as
"\<chassis\_ip\>;\<card\_number\>;\<port_number\>" (eg. "10.1.1.1;3;4"). Then
create an assignment like the following:

```go
vport := &ixconfig.Vport{
    Name: "Example Port",
    Location: "10.1.1.1;3;4",
})
cfg.Vport = append(cfg.Vport, vport)
```

### Minimal Topology Config

First create a topology and add it to the config:

```go
topology := &ixconfig.Topology{
  Name: "Example Topology",
}
cfg.Topology := append(cfg.Topology, topology)
```

To reference the vport in the topology config, you need to know its XPath. Apply
the initial config to update XPaths for all config objects. (The 'true' argument
to 'overwrite' in the examples also ensures a full reset the IxNetwork session's
config.) Then attach the vport to the topology:

```go
client.SetConfig(cfg, cfg, true)
topology.Vports = append(topology.Ports, vport.XPath().String())
```

Finally configure a device group with a basic protocol configuration for the
attached ports, and attach it to the topology. Note the use of the Multivalue\*
helper functions, which simplify construction of Multivalue config objects:

```go
deviceGroup := &ixconfig.TopologyDeviceGroup{
  Name:       "Example DeviceGroup",
  Multiplier: ixconfig.NumberUint32(1),
  Enabled:    ixconfig.MultivalueTrue(),
  Ethernet:   []*ixconfig.TopologyEthernet{
    Name: ixconfig.String("Example ethernet protocol configuration"),
    Mtu: ixconfig.MultivalueUint32(2000),
  },
}
topology.DeviceGroup = append(topology.DeviceGroup, deviceGroup)
```

### Config Push

Call SetConfig referencing just the Topology config with the 'overwrite' arg set
to 'false'. Setting 'overwrite' to 'false' updates only the values provided in
the specified config object. It preserves the rest of the existing configuration
for the IxNetwork session. Any values not explicitly set in the config and not
already present in the existing config will be set to default values by the
IxNetwork server.

```go
err := client.SetConfig(cfg, cfg.Topology, false)
```

### Config Caveats

*   The `GetConfig` method constructs an entirely new config object, which will
    not include augmented information like object IDs (discussed below), so
    you'll typically only want to call `GetConfig` once at the beginning of your
    test session.
*   Don't reuse config objects in multiple places when constructing your config.
    They are reference objects and must have a unique XPath for their particular
    location in the config.

## Operations

Operations are endpoints in the IxNetwork REST API that trigger a variety of
behaviors, such as starting traffic flows. They need to be executed using REST
API paths, rather than JSON config-based operations, but the client provides
access to the underlying Session to do this. Global operations (in this example,
starting protocols) can be executed directly using their known REST API path
(relative to the root of the IxNetwork session).

```go
rsp, err := client.Session().Post(ctx, "operations/startallprotocols", ixweb.OpArgs{"sync"}))
```

Other operations may apply to specific config objects and require a lookup to
obtain the REST API path for the target object. The path can be retrieved using
the `UpdateIDs` method. This path may be needed to specify a particular
operation path or to specify a target for a global operation. The following
example shows an example of the latter, resetting a port's CPU:

```go
configClient.UpdateIDs(cfg, vport)
rsp, err := client.Session().Post(ctx, "vport/operations/resetportcpu", ixweb.OpArgs{[]string{vport.GetRestID()}})
```

## Development Advice

### Understanding Config Object Fields/Structure

The set of available configuration objects/options within IxNetwork is
humongous. The following options are recommended for navigating the config
hierarchy:

*   Use the REST API explorer from a session in the IxNetwork Web UI. The object
    hierarchy is exactly the same, except that all object names in the config
    library are capitalized as required for public access in golang.
*   Explores the OpenAPI spec https://openixia.github.io/ixnetwork_openapi.

Looking at godocs for the generated library is *not* recommended, because
neither can perform reasonably given the size of the generated code.

### Building Complex Configurations

The IxNetwork REST API does not always produce specific pointers to
configuration errors, and it cannot understand the particular intent of a
configuration anyways, so developing/debugging large/complex test scenarios can
be difficult. The following practices are recommended to make that easier.

*   Start with as large as possible of a known working config that approximates
    the scenario you'd like to construct. Apply that config using the JSON
    config client.
*   Make small changes to the known working config to work towards your desired
    config and continually validate those in the Web UI after applying changes
    using the JSON config client.
*   For particularly complex scenarios, use the Web UI initially for
    configuration and then use the client to view what the JSON configuration
    looks like:

```go
import (
  ...
  "github.com/kylelemons/godebug/pretty"
  "github.com/openconfig/ondatra/internal/ixconfig"
  ...
)
...
  cfg, _ := client.GetConfig()
  // Printing the entire config would be noisy.
  pretty.Print(cfg.Topology)
  pretty.Print(cfg.Traffic)
...
```
