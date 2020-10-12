package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/DanMHammer/statusmonitor/proto"

	"github.com/DanMHammer/statusmonitor/cache"

	"google.golang.org/grpc"
)

type server struct{}

var cacheEngineFlag = flag.String("engine", "gocache", "Storage engine to use for hashes and messages.  Supported: redis, gocache. Default: gocache")

var minutesToExpire = 5
var minutesToDelete = 10

// Cache engine for results
var Cache cache.CacheEngine

func main() {
	Cache, _ = cache.SetupCache(*cacheEngineFlag, minutesToExpire, minutesToDelete)

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := server{}

	grpcServer := grpc.NewServer()

	proto.RegisterStatusServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

// CheckIn Handler for server
func (s *server) CheckIn(ctx context.Context, message *proto.Status) (*proto.Message, error) {
	log.Printf("Recieved status from client %d: %s", message.Id, message.Timestamp)

	Cache.Save(fmt.Sprint(message.Id), fmt.Sprint(message.Timestamp))

	return &proto.Message{Body: "Hello From the Server!"}, nil
}
