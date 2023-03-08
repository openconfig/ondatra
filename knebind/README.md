# KNE Binding

The KNE Binding is an implementation of the Ondatra binding interface that runs
on a network topology created with
[openconfig/kne](https://github.com/openconfig/kne).

## Flags

The KNE Binding requires a `--topology` flag be passed that specifies the path
to a KNE topology text proto. It also accepts a number of optional flags, as
shown in the following table.

Key             | Description
--------------- | -------------------------------------------
`topology`      | Path to a KNE topology text proto
`kubeconfig`    | Optional path to your kubeconfig file; defaults to ~/.kube/config
`skip_reset`    | If true, skip initial device reset during reservation; defaults to false
`node_creds`    | Repeated per-node credentials in the form "nodeName/username/password"
`vendor_creds`  | Repeated per-vendor credentials in the form "VENDOR/username/password"
`default_creds` | Optional default credentials in the form "username/password"

### Device Credentials

An example combination of credentials flags:

```
--node_creds=n1/user/pass --node_creds=root/root123 \
--vendor_creds=CISCO/admin/admin --default_creds=tester/hunter2
```

The per-node credentials take precedence over the per-vendor credentials, and
they both take precedence over the default credentials. If the node does not
match a provided per-node or per-vendor credential and no default credentials
are provided, then no credentials will be used.

## Running the Integration Test

To execute the test, you must create a local KNE topology with at least two
linked nodes and pass both the testbed and topology as flags to the test:

```
go test -testbed=testbed.textproto -topology=topology.textproto
```

This repo includes an
[example integration test](integration/integration_test.go) that uses the KNE
binding, a [testbed file](integration/testbed.textproto) for that
test, and a [KNE topology file](integration/topology.textproto) that is matched
by the testbed.
