package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/rnsasg/GO_Projects/go-grpc/chat"
	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50001))
	if err != nil {
		log.Fatalf("Error in listening on port %s", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &ChatServer{})

	log.Println("Starting gRPC Server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error in starting gRPC server")
	}
}

func (c *ChatServer) SayHello(ctx context.Context, req *pb.ChatMessage) (*pb.ChatMessage, error) {
	log.Printf("Message from client: %s", req.Msg)
	return &pb.ChatMessage{Msg: "Hare Rama"}, nil
}
