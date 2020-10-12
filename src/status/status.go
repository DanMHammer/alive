package status

import (
	"log"

	"github.com/DanMHammer/statusmonitor/cache"
	"golang.org/x/net/context"
)

// Server - server struct
type Server struct {
	Cache *cache.Cache
}

// CheckIn Handler for server
func (s *Server) CheckIn(ctx context.Context, message *Status) (*Message, error) {
	log.Printf("Recieved status from client %d: %s", message.Id, message.Timestamp)

	s.Cache.SaveStatus(message)

	return &Message{Body: "Hello From the Server!"}, nil
}
