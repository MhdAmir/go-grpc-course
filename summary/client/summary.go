package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MhdAmir/grpc-go-course/summary/proto"
)

func doSummary(c pb.SummaryServiceClient) {
	log.Println("do Summary invoked")
	var firstNumber int32
	var secondNumber int32

	log.Printf("Input your first number : ")
	fmt.Scanf("%d", &firstNumber)

	log.Printf("Input your first number : ")
	fmt.Scanf("%d", &secondNumber)

	res, err := c.Summary(context.Background(), &pb.SummaryRequest{
		FirstNumber: firstNumber,
		LastNumber:  secondNumber,
	})

	if err != nil {
		log.Fatalf("Couldn't Summarize: %v\n", err)
	}

	log.Printf("Summary of %d + %d = %d", firstNumber, secondNumber, res.Result)

}
