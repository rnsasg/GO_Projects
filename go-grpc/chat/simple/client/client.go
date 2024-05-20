package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/rnsasg/GO_Projects/go-grpc/chat/simple/proto"
	"google.golang.org/grpc"
)

func Start(port int) pb.ChatServiceClient {
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error in startin client: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	cli := pb.NewChatServiceClient(conn)
	return cli
}

func Greet(client pb.ChatServiceClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.ChatMessage{Msg: message})

	if err != nil {
		log.Fatalf("Error in talking with server : %s", err.Error())
		os.Exit(1)
	}

	log.Printf("Response from Server %s", resp.Msg)
}
