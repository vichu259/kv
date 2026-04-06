package raft

import (
	"context"
	"sync"

	"github.com/vichu259/kv/api/raftpb"
)

type Node struct {
	raftpb.UnimplementedRaftServiceServer

	mu sync.Mutex
}

func (*Node) Ping(context.Context, *raftpb.PingRequest) (*raftpb.PingResponse, error) {
	return &raftpb.PingResponse{
		Message: "Hello from Raft Node!",
	}, nil
}
