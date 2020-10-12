package status

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) CheckIn(ctx context.Context, message *Status) (*Message, error) {
	log.Printf("Recieved status from client %d: %s", message.Id, message.Timestamp)
	return &Message{Body: "Hello From the Server!"}, nil
}
