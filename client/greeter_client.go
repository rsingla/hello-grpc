package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/rsingla/hello-grpc/helloworld"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Rajeev"})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Printf("response: %s", res.Message)
}
