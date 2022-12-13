[![Actions Status](https://github.com/openconfig/ondatra/workflows/Go/badge.svg)](https://github.com/openconfig/ondatra/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/openconfig/ondatra)](https://goreportcard.com/report/github.com/openconfig/ondatra)
[![GoDoc](https://godoc.org/istio.io/istio?status.svg)](https://pkg.go.dev/github.com/openconfig/ondatra)
[![License: BSD](https://img.shields.io/badge/license-Apache%202-blue)](https://opensource.org/licenses/Apache-2.0)

# Ondatra: Open Network Device Automated Test Runner and API

Ondatra is a framework for writing and running tests against both real and
containerized network devices.

For an introduction to Ondatra, take the
[Ondatra Tour](https://docs.google.com/viewer?url=https://raw.githubusercontent.com/openconfig/ondatra/main/internal/tour/tour.pdf).

## Building Ondatra

To build and execute Ondatra's unit tests, run the following:

```
go generate ./...
go build ./...
go test $(go list ./... | grep -v /integration)
```

## Providing a Binding

The Ondatra *binding* is the API layer through which Ondatra connects to and
controls the devices in your test environment. For an Ondatra test to run in
your environment, you must provide an implementation of the `Binding` interface.
For [testing on KNE](#testing-on-kne), Ondatra comes bundled with a binding
implementation for accessing KNE on your local machine. For testing on your
physical devices,
[read how to author your own Binding implementation](binding/README.md).

## Testbed File

To run an Ondatra test, the user must specify the *testbed* of resources that
the Ondatra test runner should reserve in advance. The testbed is specified in
an external text file in protobuf text format. The protobuf `Testbed` message is
defined in [proto/testbed.proto](proto/testbed.proto), and an example testbed
can be found in
[knebind/integration/testbed.textproto](knebind/integration/testbed.textproto).
As the proto definition and example show, testbed consists of the DUTs (devices
under test), ATEs (automated test equipment), the links between them, as well as
properties of the DUTs, ATEs, and links. It is the job of the `Reserve` method
in the binding implementation to locate available resources that match the
abstract topology and criteria specified in the testbed file.

## Writing an Ondatra Test

Ondatra provides a set of
[fluent interfaces](https://en.wikipedia.org/wiki/Fluent_interface) for
configuring and interacting with network devices. See the
[full API reference documentation](https://pkg.go.dev/github.com/openconfig/ondatra#section-documentation).

The Ondatra gNMI API plays a particularly central role in Ondatra tests. It
provide a set of test helpers for querying telemetry and setting the state of
the device via gNMI. [Read more about the Ondatra gNMI API](gnmi/README.md).

## Running an Ondatra Test

An Ondatra test is a Go test, and so is run with `go test`, albeit with some
additional flags to control the execution of the test:

*   `-testbed` (*required*): Path to the testbed text proto file.
*   `-wait_time` (*optional*): Maximum amount of time the test should wait until
    the testbed is ready. If not specified, the binding chooses the amount of
    time to wait.
*   `-run_time` (*optional*): Timeout of the test run, excluding the wait time
    for the testbed to be ready. If not specified, no limit is imposed.
*   `-xml` (*optional*): File path to write JUnit XML test results; disables
    normal Go test logging.
*   `-debug` (*optional*): Whether the test is run in debug mode.
*   `-reserve` (*optional*): Reservation id or a mapping of device and port IDs
    to names; allowed only in [debug mode](#debugging-an-ondatra-test)

To run a subset of the test cases of an Ondatra test, use the Go test `-run`
flag. While the `-run` flag accepts an arbitrary
[Go regexp](https://golang.org/s/re2syntax), to match a single test case it is
usually sufficient to just pass the name of the test function:

```shell
$ go test -testbed=testbed.textproto -config=config.yaml -run=$RUN_THIS_TEST_CASE
```

Read the
[Go doc on subtests](https://pkg.go.dev/testing#hdr-Subtests_and_Sub_benchmarks)
for more details on matching subtests with the `-run` flag.

## Debugging an Ondatra Test

To run an Ondatra test in debug mode, pass the `-debug` flag to `go test`. Debug
mode allows you to insert breakpoints in your code with one simple line:

```go
ondatra.Debug().Breakpoint(t)
```

The `Breakpoint` and `Breakpointf` functions allow you to include custom text in
the breakpoint message, too. For example:

```go
ondatra.Debug().Breakpoint(t, "this should be unreachable")
ondatra.Debug().Breakpointf(t, "myVar has value %v", myVar)
```

Debug mode also offers a menu option to pause the test immediately after the
testbed is reserved. This is useful if you want to manually inspect the testbed
before the test cases run, or to run a separate test execution against with the
same reservation.

To debug a test against a pre-allocated reservation, set the `-reserve` flag to
the reservation ID:

```shell
$ go test -testbed=testbed.textproto -config=config.yaml -debug -reserve=123abc
```

To debug the test against specifically-named devices, set the `-reserve` flag to
a comma-separated list of *id=name* strings at the command line, where ports are
named with the syntax *deviceID:portID=portName*:

```shell
$ go test -testbed=testbed.textproto -config=config.yaml -debug \
    -reserve=dut=mydevice,dut:port1=Ethernet1/1,ate=myixia,ate:port2=2/3
```

## Logging Verbosity

Ondatra always sets the Go test `-v` flag to true for verbose test output, so
there is no need to set this flag explicitly in your test invocations.

Ondatra uses [glog](https://pkg.go.dev/github.com/golang/glog) for its own
logging. By default, glog logs to a temporary dir, but setting the
`-alsologtostderr` flag will output those logs to stderr. You can increase the
verbosity glog by setting its `-v` flag to a positive integer. Increasing the
value of `v` produces increasingly verbose and granular details in the logs, as
outlined in this table:

Log Level | Example Information
--------- | --------------------------------------------------------
1         | `SetRequest`/`SubscribeRequest` dumps
2         | Each `Update`/`Delete` received in a `SubscribeResponse`
3         | Ixia gNMI translation details

Because go test has its own `-v` flag, setting glog's `-v` value must be
preceded by `-args` to avoid it being interpreted by go test itself. For
example:

```
go test -alsologtostderr -args -v=1
```

See the
[glog package documentation](https://pkg.go.dev/github.com/golang/glog#pkg-overview)
for more information on all the glog flags.

## XML Test Report

Ondatra has the ability to output test results in JUnit XML format. If you pass
`-xml=[path]` to your `go test` invocation, Ondatra will use
[go-junit-report](https://github.com/jstemmer/go-junit-report) to translate the
Go test log to an XML file at the provided path.

To attach properties to the test or individual test cases in the XML report, use
the `ondatra.Report()` API. For example, `ondatra.Report().AddTestProperty()`
attaches a key-value property to the currently running test case.

The help with the development of external tooling that makes use of the XML
output, the `report` package also provides the utility functions `ReadXML` and
`ExtractProperties` for decoding the XML and extracting a properties map,
respectively.

## Testing on KNE

You don't have to code your own binding implementation before getting started
with Ondatra, because Ondatra comes packaged with a binding for
[KNE](https://github.com/openconfig/kne), and an
[example Ondatra test](knebind/integration/integration_test.go) that uses that
binding. See the [knebind README](knebind/README.md) for more on how to use the
KNE binding and run the example test.
