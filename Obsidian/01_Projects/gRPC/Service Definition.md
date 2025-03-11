RPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types.

[[protocol buffers]] are used the interface definition language (IDL) for defining the service interface and the payload structure.

There are four types of Service methods:
- Unary RPC (single request, single response)
```protobuf
rpc SayHello(HelloRequest)return (HelloResponse);
```
- Server streaming (single request, stream of data back)
```protobuf
rpc LotsOfReplies(HelloRequest)returns(stream HelloResponse)
```
- Client streaming (sequence of messages, single response)
```protobuf
rpc LotsofGreetings(stream HelloRequest)returns(HelloResponse)
```
- Bidirectional streaming (sequence of messages, sequences of response)
```protobuf
rpc Bidirectional(stream HelloRequest)returns(stream HelloResponse)
```

The flow is straightforward:
- define your service in a `proto` file.
- use a compiler to generate code for your client and server
- implement the service interface on the server side.

## Synchronous vs Asynchronous

gRPC supports both blocking and non-blocking code implementations

## RPC lifecycle

### Unary RPC

- The client calls a stub method, which notifies the server that a Remote Procedure Call has been invoked with necessary information like client's metadata, method name and deadlines (if applicable).
- The server might send back it's own metadata or wait for client's request message.
- Once the server gets the message, it executes the business logic on it, generates a response and returns it to the client with status details and optional trailing metadata.
- if response status is OK, client gets the response and completing the lifecycle.

### Server Streaming RPC

Similar to Unary RPC, just instead of sending a single message as response, the server returns a stream of messages. After sending all the messages, the server's status details and optional trailing metadata is send to the client. This completes the lifecycle on the server side, the client waits till it gets all the messages.

### Client Streaming RPC

Similar to Server Streaming RPC, just instead of server sending a stream of messages, client sends them.

### Bidirectional RPC

The call is initiated by the client invoking a server method including the client metadata, method name and deadlines. The server can choose to send back the metadata or wait for the client to start streaming messages.

Client and server streaming process is application specific, since both the streams are independent and client/server can decide how to process the messages. The server can wait for all the client messages and then start responding, or it can respond to every client request as it arrives.

