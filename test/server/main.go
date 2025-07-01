package main

import (
	"context"
	"fmt"
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
	hostname, _ := os.Hostname()
	ip := getOutboundIP()
	return &pb.Response{Ip: fmt.Sprintf("%s (%s)", ip, hostname)}, nil
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "unknown"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func main() {
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
