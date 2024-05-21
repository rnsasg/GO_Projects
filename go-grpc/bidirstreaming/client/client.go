package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/rnsasg/GO_Projects/go-grpc/bidirstreaming"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())
	conn, err := grpc.Dial(":50003", grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("Error in connect to server %v", err)
	}

	cli := pb.NewMathClient(conn)
	stream, err := cli.Max(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}
	var max int32
	ctx := stream.Context()
	done := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			// generates random number and sends it to stream
			rnd := int32(rand.Intn(i))
			req := pb.Request{Num: rnd}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("%d sent", req.Num)
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			max = resp.Result
			log.Printf("new max %d received", max)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Printf("finished with max=%d", max)
}
