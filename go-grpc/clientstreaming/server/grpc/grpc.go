package grpc

import (
	"io"
	"log"

	protoport "github.com/rnsasg/GO_Projects/go-grpc/clientstreaming"
)

type Port struct{}

func (p Port) Create(stream protoport.PortService_CreateServer) error {
	var total int32

	for {
		port, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&protoport.CreateResponse{
				Total: total,
			})
		}
		if err != nil {
			return err
		}

		total++
		log.Printf("%+v\n", port)
	}
}
