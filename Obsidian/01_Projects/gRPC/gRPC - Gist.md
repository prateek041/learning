gRPC makes developing distributed applications simpler by enabling client to call methods directly on a server application on a different machine as if it were a local object.

The idea is to define [[service]] and specifying what methods that can be called remotely with their parameters and return types.

The server implements the interface hence implementing the business logic. The client has `stub` that provides the same method as the service on the server

## [[Protocol Buffers]]

[[protocol buffers]], google's open source mechanism for serialising structured data (converting structured data into bytes or any form that can be saved).

We start with defining the structure of the data into a proto file, for example

```proto
message Person {
	string name = 1;
	int32 id = 2;
	bool has_ponycopter = 3;
}
```

then we use a compiler to generate data access classes in our preferred language. this provides the implementations of methods needed for serialisation/deserialisation, accessing fields etc.

Connecting all the concepts, we have a service and messages, which define how method parameters will look, which in code would look something like this.

```proto
// The greeter service
service Greeter {
// supported method with HelloRequest as argument and HelloReply as 
// return message
rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing name of type string.
message HelloRequest {
string name = 1;
}

// The response message containing greeting of type string.
message HelloReply {
string message = 1;
}
```
