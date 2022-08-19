[![Actions Status](https://github.com/openconfig/ondatra/workflows/Go/badge.svg)](https://github.com/openconfig/ondatra/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/openconfig/ondatra)](https://goreportcard.com/report/github.com/openconfig/ondatra)
[![GoDoc](https://godoc.org/istio.io/istio?status.svg)](https://pkg.go.dev/github.com/openconfig/ondatra)
[![License: BSD](https://img.shields.io/badge/license-Apache%202-blue)](https://opensource.org/licenses/Apache-2.0)

# Ondatra: Open Network Device Automated Test Runner and API

Ondatra is a framework for writing and running tests against both real and
containerized network devices.

## Building Ondatra

To build and execute Ondatra's unit tests, run the following:

```
go generate ./...
go build ./...
go test $(go list ./... | grep -v /integration)
```

## Binding API

The Ondatra *binding* is the API layer through which Ondatra connects to and
controls the devices in your test environment. For an Ondatra test to run in
your environment, you must provide an implementation of the
[Binding](binding/binding.go) interface. The interface defines the primitive
operations that Ondatra needs to be able to test against devices in your lab.
When authoring a new binding implementation, many methods can be left
unimplemented to start, with their implementation postponed until you have a
test that requires that piece of functionality.

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

## Running an Ondatra Test

An Ondatra test is a Go test, and so is run with `go test`, albeit with some
additional flags to control the execution of the test:

*   `-testbed` (*required*): Path to the testbed text proto file.
*   `-wait_time` (*optional*): Maximum amount of time the test should wait until
    the testbed is ready. If not specified, the binding chooses the amount of
    time to wait.
*   `-run_time` (*optional*): Timeout of the test run, excluding the wait time
    for the testbed to be ready. If not specified, no limit is imposed.
*   `-xml (*optional*): File path to write JUnit XML test results; disables
    normal Go test logging.
*   `-debug (*optional*): Whether the test is run in debug mode.

In addition, the binding implementation is free to define its own set of
optional or required flags.

## Debugging an Ondatra Test

To run an Ondatra test in "debug mode," pass the `-debug` flag to `go test.`
Debug mode gives you the option allows you to insert breakpoints in your code
with one simple line:

```
ondatra.Debug().Breakpoint(t)
```

It also offers a menu option to pause the test immediately after the testbed is
reserved. This is useful if you want to just reserve the same testbed for manual
testing, or to manually inspect the testbed before the test cases run.

## Testing on KNE

You don't have to code your own binding implementation before getting started
with Ondatra, because Ondatra comes packaged with a binding for
[KNE](https://github.com/openconfig/kne), and an
[example Ondatra test](knebind/integration/integration_test.go) that uses that
binding. See the [knebind README](knebind/README.md) for more on how to use the
KNE binding and run the example test.
