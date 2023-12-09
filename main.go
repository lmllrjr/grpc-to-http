package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/lmllrjr/grpc-to-http/gateway"
	gRPCServer "github.com/lmllrjr/grpc-to-http/server"
	"golang.org/x/net/context"
)

var (
	endpoint = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	addr     = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network  = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	// https://pkg.go.dev/github.com/golang/glog#section-readme:~:text=%2Dlogtostderr%3Dfalse%0A%09Logs,default%20temporary%20directory.

	// set log to only log to stderr
	flag.Set("logtostderr", "true")

	// set log dir to `./logs` and log to stderr, too
	// flag.Set("log_dir", "logs")
	// flag.Set("alsologtostderr", "true")

	flag.Parse()
	defer glog.Flush()

	errors := make(chan error)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := gateway.Options{
		Addr:       ":8080",
		GRPCServer: gateway.Endpoint{Network: *network, Addr: *endpoint},
	}

	go func() {
		errors <- gRPCServer.Run(ctx, *network, *addr)
	}()

	go func() {
		errors <- gateway.Run(ctx, opts)
	}()

	for err := range errors {
		glog.Fatal(err)
		return
	}
}
