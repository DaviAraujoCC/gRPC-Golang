package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received: %v", message)
	return &Message{Body: "Hello, you just typed: " + message.Body}, nil
}
