package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/DanMHammer/statusmonitor/proto"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"

	"github.com/DanMHammer/statusmonitor/cache"

	"github.com/gorilla/mux"

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

	s := server{}

	grpcServer := grpc.NewServer()

	proto.RegisterStatusServiceServer(grpcServer, &s)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/all", getAll)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)

	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	httpListener := m.Match(cmux.Any())

	g := errgroup.Group{}
	g.Go(func() error {
		return grpcServer.Serve(grpcListener)
	})
	g.Go(func() error {
		return http.Serve(httpListener, router)
	})
	g.Go(func() error {
		return m.Serve()
	})

	// Wait for them and check for errors
	err = g.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

// CheckIn Handler for server
func (s *server) CheckIn(ctx context.Context, stat *proto.Status) (*proto.Message, error) {
	log.Printf("Recieved status from client %d: %s", stat.Id, stat.Latest)

	Cache.Save(fmt.Sprint(stat.Id), stat.Started, stat.Latest)

	return &proto.Message{Body: "Hello From the Server!"}, nil
}

func getAll(w http.ResponseWriter, r *http.Request) {
	results := Cache.GetAll()

	json.NewEncoder(w).Encode(results)
}
