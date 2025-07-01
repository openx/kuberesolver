package main

import (
	"context"
	"github.com/openx/kuberesolver"
	pb "github.com/openx/kuberesolver/test/proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	kuberesolver.RegisterInCluster()
	addr := "kubernetes:///server:50051"
	defaultServiceConfig := "{\"loadBalancingConfig\": [{\"round_robin\": {}}]}"

	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(defaultServiceConfig),
	)

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	ipRequest := make(map[string]int16)

	client := pb.NewEchoClient(conn)
	log.Printf("start testing")
	for i := 0; i < 10; i++ {
		resp, err := client.GetServerIP(context.Background(), &pb.Request{})
		if err != nil {
			log.Printf("error on request %d: %v", i, err)
			continue
		}
		if k := ipRequest[resp.Ip]; k != 0 {
			ipRequest[resp.Ip]++
		} else {
			ipRequest[resp.Ip] = 1
		}
	}
	for k, v := range ipRequest {
		log.Printf("key: %s  value: %v\n", k, v)
	}

	time.Sleep(120 * time.Second)
}
