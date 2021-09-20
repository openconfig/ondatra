The knebind implementation requires a config file, specified in YAML, the path
to which must be specified in a `--config` flag to the Ondatra test. The file
supports the following keys:

Key        | Required? | Description
---------- | :-------: | ----------------------------------
`username` | yes       | username to log into the KNE nodes
`password` | yes       | password to log into the KNE nodes
`topology` | yes       | path to a KNE topology text proto
`cli`      | no        | path to the kne_cli binary
`kubecfg`  | no        | path to your kubeconfig file

If `cli` and `kubecfg` are not specified, they will be inferred from the `PATH`
environment.

An example YAML config file:

```
username: tester
password: hunter2
topology: /home/tester/topo.textproto
cli: /home/tester/go/bin/kne_cli
kubecfg: /home/tester/go/bin/.kube/config
```
