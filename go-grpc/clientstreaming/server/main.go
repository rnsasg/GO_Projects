package main

import (
	"log"
	"net"

	protoport "github.com/rnsasg/GO_Projects/go-grpc/clientstreaming"
	"github.com/rnsasg/GO_Projects/go-grpc/clientstreaming/server/grpc"
	gogrpc "google.golang.org/grpc"
)

func main() {
	log.Println("server")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := gogrpc.NewServer()
	protoServer := grpc.Port{}

	protoport.RegisterPortServiceServer(server, protoServer)

	log.Fatalln(server.Serve(listener))
}
