package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/DanMHammer/statusmonitor/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

var idPtr = flag.Int("id", 1, "machine id")

func main() {
	flag.Parse()

	fmt.Println(*idPtr)

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithTimeout(10*time.Minute))

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	c := proto.NewStatusServiceClient(conn)

	started := ptypes.TimestampNow()

	for tick := range time.Tick(10 * time.Second) {

		fmt.Println(tick)
		ts := ptypes.TimestampNow()

		message := proto.Status{
			Latest:  ts,
			Started: started,
			Id:      int32(*idPtr),
		}

		response, err := c.CheckIn(context.Background(), &message)

		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("Response from Server: %s", response.Body)
	}

	defer conn.Close()
}
