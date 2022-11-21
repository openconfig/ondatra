# Binding Interface

The [Binding interface](binding.go) abstracts Ondatra from the particulars of
how to access the devices under test. Those particulars include the inventory of
devices and their properties, how to reserve them, and how to reach their
various protocol endpoints.

An implementation of the `Reserve` method of `Binding` returns a `Reservation`
struct, which in turn contains implementations of the `DUT` and `ATE`
interfaces. Substantial care was taken in writing the godoc comments on
`Binding`, `DUT`, `ATE`, and all other interfaces in the binding package, so
please read those carefully to ensure each method of an implementation adheres
to the stated contract.

## Abstract Structs

Ondatra provides a set of "abstract" structs that *must* be embedded in any
implementation of the binding interfaces. For example, any implementation of the
`DUT` interface must embed `AbstractDUT` and any implementation of `ATE` must
embed `AbstractATE`. This is enforced by the presence of a
`mustEmbedAbstractDUT` method in `DUT` and a `mustEmbedAbstactATE` method in
`ATE`, un-exported methods that are only implemented in `AbstractDUT` and
`Abstract ATE`, respectively.

The primary purpose of the abstract structs is to enable binding implementations
to be *partial*. It is perfectly acceptable for a binding leave some interface
methods unimplemented, and embedding the abstract structs gives them an easy way
to do that. It also allows Ondatra developers to add new binding methods in the
future without breaking existing implementations. A similar effect could have
been obtained by embedding the interface itself, but then calls to unimplemented
methods would produce cryptic panics rather than the prescriptive error messages
produced by the abstract implementations.

## Dialing gRPC

The binding includes many methods that dial gRPC endpoints, like `DialGNMI`,
`DialGNOI`, etc. Each of these methods accepts a context and a variadic number
of `DialOption` arguments. The implementation of these methods *must* include a
call to `grpc.DialContext` with the specified context (or perhaps a context
derived for the specified context) is the first argument.

Ondatra calls these `Dial*` functions with some options that it believes all
gRPC dial calls should include, but it is then up to the implementation to
append any additional necessary or desirable options for connecting to the gRPC
endpoint. Every such `Dial*` implementation *must* append a dial option that
specifies the transport credentials necessary to reach the endpoint; otherwise
the call to `DialContext` will fail. For example:

```
opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
grpc.DialContext(ctx, addr, opts...)
```

Here are some other options you might consider adding, depending on your lab and
network environment:

*   `grpc.WithDefaultCallOptions(grpc.WaitForReady(true))`<br /> Try this option
    if you are experiencing transient failures in the RPC channel. Such failures
    normally manifest as "connection refused" RPC errors. One downside of this
    option is that it can cause the RPC to hang indefinitely if the channel
    never returns to a ready state. For that reason, we recommend only using
    this option in combination with the "default timeout" options documented
    below. To read more about the semantics of `WaitForReady`
    [https://github.com/grpc/grpc/blob/master/doc/wait-for-ready.md](here).

*   `grpcutil.WithUnaryDefaultTimeout(defaultTimeout)` and
    `grpcutil.WithStreamDefaultTimeout(defaultTimeout)`<br /> Try these options
    if you want to ensure all RPCs have timeouts configured. If the context
    passed to any RPC does not already have a timeout or deadline, these options
    cause the RPC to be issued with a new context that has the specified
    `defaultTimeout`. These options are provided by the Ondatra
    [`grpcutil`](https://github.com/openconfig/ondatra/blob/main/binding/grpcutil/grpcutil.go)
    package.
