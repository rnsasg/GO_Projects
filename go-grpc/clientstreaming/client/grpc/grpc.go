package grpc

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	protoport "github.com/rnsasg/GO_Projects/go-grpc/clientstreaming"
)

type Port struct {
	client protoport.PortServiceClient
}

func NewPort(conn grpc.ClientConnInterface) Port {
	return Port{
		client: protoport.NewPortServiceClient(conn),
	}
}

func (p Port) Create(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// I am hardcoding this here but you should not!
	requests := []*protoport.CreateRequest{
		{Id: "id-1", Name: "name-1"},
		{Id: "id-2", Name: "name-2"},
		{Id: "id-3", Name: "name-3"},
		{Id: "id-4", Name: "name-4"},
	}

	stream, err := p.client.Create(ctx)
	if err != nil {
		return fmt.Errorf("create stream: %w", err)
	}

	for _, request := range requests {
		if err := stream.Send(request); err != nil {
			return fmt.Errorf("send stream: %w", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("close and receive: %w", err)
	}

	log.Printf("%+v\n", response)

	return nil
}
