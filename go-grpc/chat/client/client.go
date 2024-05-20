package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/rnsasg/GO_Projects/go-grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	address := fmt.Sprintf("localhost:%d", 50001)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error in startin client: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.SayHello(ctx, &pb.ChatMessage{Msg: "Hare Krishna"})
	if err != nil {
		log.Fatalf("Error in talking with server : %s", err.Error())
		os.Exit(1)
	}

	log.Printf("Response from Server %s", resp.Msg)
}
