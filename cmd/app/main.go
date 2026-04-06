package main

import (
	"fmt"

	"github.com/vichu259/kv/internal/raft"
	"github.com/vichu259/kv/internal/transport"
)

func main() {
	server := transport.NewGRPCServer("50051", &raft.Node{})
	if err := server.Start(); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
