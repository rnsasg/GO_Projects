package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rnsasg/GO_Projects/go-grpc/serverstreaming"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	client := pb.NewStreamServiceClient(conn)
	in := &pb.Request{Id: 1}

	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Result)
		}
	}()
	<-done
	log.Printf("finished")
}
