# Ondatra: Open Network Device Automated Test Runner and API

Ondatra is a framework for writing and running tests against both real and
containerized network devices.

## Building

To build and execute Ondatra's unit tests, run the following:

```
$ go generate ./...
$ go build ./...
$ go test $(go list ./... | grep -v /integration)
```

## Binding API

The Ondatra *binding* is the API layer through which Ondatra connects to and
controls the devices in your test environment. For an Ondatra test to run in
your environment, you must provide an implementation of the
[Binding](internal/binding/binding.go) interface. The interface defines the
primitive operations that Ondatra needs to be able to test against devices in
your lab. When authoring a new binding implementation, many methods can be left
unimplemented to start, with their implementation postponed until you have a
test that requires that piece of functionality.
