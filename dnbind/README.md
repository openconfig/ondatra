# Drivenets Binding

Drivenets Binding is an implementation of the Ondatra binding interface with support
limited to [Config](https://pkg.go.dev/github.com/openconfig/ondatra/config) interface.


## Flags

Drivenets Binding integration tests requires a `--config` flag be passed that specifies device information.  
Running Drivenets Binding tests requires passing device credential by using `--node_cres` flag.


### Device Credentials

An example of credentials flags:

```
--node_creds=hostname/user/pass
```

An example of yaml configuration file where `id` needs to match testbed `id`:

```
nodes:
  - id: testbed_id
    hostname:  foo
    credentials:
      username: name
      password: pass
```


## Running the Integration Test

To execute the test, you must update config.yaml with your Drivenets device details
and pass both the testbed and config files as flags to the test:

```
go test github.com/openconfig/ondatra/dnbind/integration --testbed=testbed.textproto --config=config.yaml
```

This repo includes an
[example integration test](integration/integration_test.go) that uses the Drivenets
binding, a [testbed file](integration/testbed.textproto) for that test, and a
[mock configuration file](integration/config.yaml) that is matched by the
testbed.

## Session idle timeout

CLI might disconnect if idle for prolonged period of time.  
One can disable session timeout be configuring: ```system login session-timeout 0```.  
This does not affect active CLI connection, so for intendeed behaviour one must reconnect.

Sample usage:  
```golang
  dut := ondatra.DUT(t, "dut")
  // disable session timeout
	dut.Config().New().
		WithDrivenetsText(
			`system login session-timeout 0
            `).
		Append(t)
  dut := ondatra.DUT(t, "dut")
  /* your code goes here with disabled session timeout */
```


## Interactive Commands

Operational commands requiring confirmation or user input of any kind are not supported by this API.

### Disable Show commands pagination (Recommended)

By running ```set cli-terminal-length 0``` as an initial step or apending ``` | no-more ``` to all show commands.  
Paginated output leads to command getting stuck waiting for prompt.

### List of unsupported commands

- ‘run monitor ...’  
- Any command followed by ```‘| monitor interval’```  
  
*Commands that exit from CLI:*
- ‘exit’ 
- ‘quit’ 
 
*Interactive commands that initiate outbound connection and require user input:*
- ‘run ssh ...’ 
- ‘run ipmi ...’ 
- ‘run start shell ...’  
 
*Commands that require confirmation:*
- GI: 
    - ‘request system delete’
    - ‘request system deploy’
    - ‘request system install’
    - ‘request system revert-stack’
    - ‘request system target-stack ...’
    - ‘request system tech-support ...’ ***can bypass using ‘force’ keyword***
 
- DNOS:
    - ‘load override golden-config'
    - ‘request system tech-support ...’ ***can bypass using ‘force’ keyword***
    - ‘request system restart factory-default' 
    - ‘request file copy’ 
    - ‘request file delete’ 
    - ‘request interface management <xxx> access-list' 
    - ‘request system delete’ 
    - ‘request system generate golden-config' 
    - ‘request system ncc switchover’ 
    - ‘request system container restart’ 
    - ‘request system process restart’ 
    - ‘request system process stop’ 
    - ‘request system restart <xxx>’ 
    - ‘request system revert-stack' 
    - ‘request system target-stack' 
 
- Interactive configuration commands:  
    - ‘system profile’
 
- Configuration commands that take plain-text password value.  
  Using plain-text option is interactive.  
  If you want to configure password value, pass the encrypted password value instead
    - ‘system aaa-server radius server password’ 
    - ‘system login ipmi user password’ 
    - ‘system login ncm user password’ 
    - ‘system login user password’ 
    - ‘system ntp authentication key-id' 
    - ‘system snmp user authentication password’ 
    - ‘protocols mpls traffic-engineering pcep authentication enabled password’ 
    - ‘protocols ldp authentication md5 password’ 
    - ‘protocols ldp neighbor <> authentication md5 password’ 
    - ‘protocols ospf area <> interface <> authentication-key md5 key-id <> password’ 
    - ‘protocols ospfv3 area <> authentication ipsec spi <> md5 password’ 
    - ‘protocols ospfv3 area <> interface <> authentication ‘ 
    - ‘system aaa-server tacacs server priority <> address <> password’ 
    - ‘protocols bgp <> neighbor <> authentication md5 password’ 
