# grpc-to-http
The gRPC-Gateway is a plugin that generates a reverse-proxy server which translates a RESTful HTTP API into gRPC and vice versa.

In other words, gRPC-Gateway will create a layer over your gRPC services that will act as a Restful/JSON service to a client. This server is generated according to the [`google.api.http`](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46) annotations in your service definitions.

## Why gRPC-Gateway?
gRPC Gateway is a tool that allows clients to access gRPC servers using HTTP/1.1 or JSON. This is useful because some clients may not have the ability to directly call gRPC services, or the gRPC service may need to be exposed over HTTP/1.1 for legacy or compatibility reasons.

The gRPC Gateway acts as a proxy between the client and the gRPC service. When a client sends an HTTP/1.1 request to the gRPC Gateway, the gateway translates the request to a gRPC call and forwards it to the gRPC server. The gRPC server processes the request and returns the response to the gRPC Gateway, which translates the response to an HTTP/1.1 or JSON response and returns it to the client.

To use gRPC Gateway, you typically define a separate set of protobuf messages and gRPC service definitions that represent the HTTP/1.1 or JSON API that you want to expose to clients. These definitions are then used to generate a gRPC Gateway reverse proxy that implements the translation between HTTP/1.1 or JSON and gRPC.

Overall, gRPC Gateway makes it easy to provide HTTP/1.1 or JSON access to gRPC services, allowing clients to use whichever protocol is most appropriate for their needs.

>**Note**:  
>* Legacy clients might not support gRPC and require a Restful/JSON interface
>* Browsers may not support gRPC out of the box; so for the web client that wants to interact with gRPC services, gRPC-Gateway is the go-to option.

## Setup
### commands used
generate stubs and proxy from proto buffer file:
```sh
make all
```

```sh
go run main.go -endpoint='localhost:9090' -addr=':9090' -network='tcp'
```

## Links
* https://github.com/grpc-ecosystem/grpc-gateway
* https://youtu.be/Pq1paKC-fXk
