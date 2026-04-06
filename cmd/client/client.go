package main

import (
	"context"
	"log"

	"github.com/vichu259/kv/api/raftpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := raftpb.NewRaftServiceClient(conn)

	resp, err := client.Ping(context.Background(), &raftpb.PingRequest{})
	if err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	log.Printf("Received response: %s", resp.Message)
}
