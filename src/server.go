package main

import (
	"flag"
	"log"
	"net"

	"github.com/DanMHammer/statusmonitor/cache"
	"github.com/DanMHammer/statusmonitor/status"

	"google.golang.org/grpc"
)

var cacheEngineFlag = flag.String("engine", "gocache", "Storage engine to use for hashes and messages.  Supported: redis, gocache. Default: gocache")

// Cache engine for results
var Cache cache.CacheEngine

func main() {
	Cache, _ = SetupCache()

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := status.Server{Cache}

	grpcServer := grpc.NewServer()

	status.RegisterStatusServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
