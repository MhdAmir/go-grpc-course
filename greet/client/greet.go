package main

import (
	"context"
	"log"

	pb "github.com/MhdAmir/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do Greet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Amir",
	})

	if err != nil {
		log.Fatalf("Couldn't Greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
