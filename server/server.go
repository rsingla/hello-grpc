package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/rsingla/hello-grpc/helloworld"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := &pb.HelloResponse{Message: "Hello " + in.Name}
	log.Printf("response: %s", resp.Message)
	return resp, nil
}

func (s *server) LetsParty(ctx context.Context, in *pb.PartyRequest) (*pb.PartyResponse, error) {
	resp := &pb.PartyResponse{}
	resp.Name = in.Name
	resp.Description = "Let's party " + in.Name
	resp.InviteSent = true
	resp.Cost = 1100
	resp.RemainingPasses = 115

	log.Printf("response: %s", resp)
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
