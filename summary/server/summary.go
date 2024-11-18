package main

import (
	"context"
	"log"

	pb "github.com/MhdAmir/grpc-go-course/summary/proto"
)

func (s *Server) Summary(ctx context.Context, in *pb.SummaryRequest) (*pb.SummaryResponse, error) {
	log.Printf("Summary was invoked with %v", in)

	return &pb.SummaryResponse{
		Result: in.FirstNumber + in.LastNumber,
	}, nil
}
