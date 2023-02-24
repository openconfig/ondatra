# KNE Binding

The KNE Binding is an implementation of the Ondatra binding interface that runs
on a network topology created with
[openconfig/kne](https://github.com/openconfig/kne).

## Config File

The KNE Binding requires a YAML config file that lets the binding know how to
connect to your KNE topology. The YAML must be specified in a `-config` flag
passed to the Ondatra test. The file supports the following keys:

Key           | Required? | Description
------------- | --------- | ----------------------------------
`username`    | no        | default username to log into the KNE nodes
`password`    | no        | default password to log into the KNE nodes
`credentials` | no        | map of credentials to use for specific nodes / vendors
`topology`    | yes       | path to a KNE topology text proto
`kubecfg`     | no        | path to your kubeconfig file
`skip_reset`  | no        | if true, skip initial device reset that happens during reservation

If `kubecfg` is not specified, it will be inferred from the `PATH` environment.

A basic example YAML config file:

```
username: tester
password: hunter2
topology: /home/tester/topo.textproto
```

### Device Credentials

An example YAML config file with optional extra credential fields:

```
username: tester
password: hunter2
credentials:
  node:
    r1:
      username: user
      password: pass
    r2:
      username: root
      password: root123
  vendor:
    ARISTA:
      username: admin
      password: admin
topology: /home/tester/topo.textproto
```

The per-node credentials are used over the per-vendor credentials. The default
username/password are used if there are no matches. A precedence order is
defined below:

1. credential provided for a specific node by name
1. credential provided for a vendor of the node
1. credential from the default username and password fields
1. no credentials

## Running the Integration Test

This repo includes an
[example integration test](integration/integration_test.go) that uses the KNE
binding, a [testbed file](integration/testbed.textproto) for that
test, and a [KNE topology file](integration/topology.textproto) that is matched
by the testbed. To execute the test, you must:

1.  create a local KNE topology with at least two linked nodes, as the testbed
    requires
1.  create a config file for your topology, as specified above
1.  run the test passing both the testbed and config file flags:

```
go test -testbed=testbed.textproto -config=path/to/config.yaml
```

A YAML config file that works with that
topology is:

```
username: admin
password: admin
topology: /[YOUR GIT CLONE PATH]/ondatra/knebind/integration/topology.textproto
```
