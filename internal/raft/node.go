package raft

import (
	"context"
	"sync"
	"time"

	"github.com/vichu259/kv/api/raftpb"
)

type Node struct {
	raftpb.UnimplementedRaftServiceServer

	mu sync.Mutex

	id    string
	peers []string
	state State

	// Persistent state on all servers
	currentTerm int
	votedFor    string
	log         []raftpb.LogEntry

	// Volatile state on all servers

	// Highest log entry known to be committed
	commitIndex int
	// Highest log entry applied to state machine
	lastApplied int

	// Volatile state on leaders

	// For each server, index of the next log entry to send to that server
	nextIndex map[string]int
	// For each server, index of highest log entry known to be replicated on server
	matchIndex map[string]int

	electionTimeout  time.Duration
	heartbeatTimeout time.Duration
}

func (*Node) Ping(context.Context, *raftpb.PingRequest) (*raftpb.PingResponse, error) {
	return &raftpb.PingResponse{
		Message: "Hello from Raft Node!",
	}, nil
}

func (*Node) RequestVote(context.Context, *raftpb.RequestVoteRequest) (*raftpb.RequestVoteResponse, error) {
	return &raftpb.RequestVoteResponse{
		Term:        0,
		VoteGranted: false,
	}, nil
}
