package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/rnsasg/GO_Projects/go-grpc/bidirstreaming"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMathServer
}

func main() {
	lis, err := net.Listen("tcp", ":50003")
	if err != nil {
		fmt.Printf("Error in Listening on port %d, err: %v", 50003, err)
	}

	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s server) Max(srv pb.Math_MaxServer) error {
	log.Println("start new server")
	var max int32
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Receive the data
		req, err := srv.Recv()

		if err != io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}

		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// continue if number reveived from stream
		// less than max
		if req.Num <= max {
			continue
		}

		max = req.Num
		resp := pb.Response{Result: max}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new max=%d", max)
	}
}
