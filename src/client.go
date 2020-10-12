package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/DanMHammer/status/status"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

var idPtr = flag.Int("id", 1, "machine id")

func main() {
	flag.Parse()

	fmt.Println(*idPtr)

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := status.NewStatusServiceClient(conn)

	for tick := range time.Tick(10 * time.Second) {

		fmt.Println(tick)
		ts := ptypes.TimestampNow()

		message := status.Status{
			Timestamp: ts,
			Id:        int32(*idPtr),
		}

		response, err := c.CheckIn(context.Background(), &message)

		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("Response from Server: %s", response.Body)
	}
}
