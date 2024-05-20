package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/rnsasg/GO_Projects/go-grpc/chat/simple/proto"
	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func Start(port int) {
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &ChatServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("Error in listening on port %s", err.Error())
	}
	log.Println("Starting gRPC Server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error in starting gRPC server")
	} else {
		log.Println("Succesfully Started gRPC Server")
	}
}

func (c *ChatServer) SayHello(ctx context.Context, req *pb.ChatMessage) (*pb.ChatMessage, error) {
	log.Printf("Message from client: %s", req.Msg)
	return &pb.ChatMessage{Msg: "Welcome Back"}, nil
}
