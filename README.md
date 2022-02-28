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
additional flags related to the reservation of the testbed:

*   `-testbed` (*required*): Path to the testbed text proto file.
*   `-wait_time` (*optional*): Maximum amount of time the test should wait until
    the testbed is ready. If not specified, the binding chooses the amount of
    time to wait.
*   `-run_time` (*optional*): Timeout of the test run, excluding the wait time
    for the testbed to be ready. If not specified, no limit is imposed.

In addition, the binding implementation is free to define its own set of
optional or required flags.

## Testing on KNE

You don't have to code your own binding implementation before getting started
with Ondatra, because Ondatra comes packaged with a binding for
[KNE](https://github.com/google/kne), and an
[example Ondatra test](knebind/integration/integration_test.go) that uses that
binding. See the [knebind README](knebind/README.md) for more on how to use the
KNE binding and run the example test.
