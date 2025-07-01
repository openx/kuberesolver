package main

import (
	"context"
	"github.com/openx/kuberesolver"
	pb "github.com/openx/kuberesolver/test/proto/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	kuberesolver.RegisterInCluster()
	addr := "server.default.svc.cluster.local:50051"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	conn, err = grpc.NewClient(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	for i := 0; i < 500; i++ {
		resp, err := client.GetServerIP(context.Background(), &pb.Request{})
		if err != nil {
			log.Printf("error on request %d: %v", i, err)
			continue
		}
		log.Printf("Response #%d: %s", i, resp.Ip)
	}
}
