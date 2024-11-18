package main

import (
	"log"
	"net"

	pb "github.com/MhdAmir/grpc-go-course/summary/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SummaryServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	server := grpc.NewServer()
	pb.RegisterSummaryServiceServer(server, &Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
