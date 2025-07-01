package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/openx/kuberesolver/test/proto/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) GetServerIP(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	ip := os.Getenv("IP_ADDR")
	return &pb.Response{Ip: ip}, nil
}

func main() {
	log.Println(os.Getenv("IP_ADDR"))
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
