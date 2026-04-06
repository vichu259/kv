# kv

A distributed key-value store built on the Raft consensus protocol, with gRPC transport.

## Overview

This project implements the foundation for a Raft-based distributed KV store in Go. It exposes a gRPC interface for inter-node communication and includes a test client for manual verification.

## Project Structure

```
.
├── api/raftpb/          # Protobuf definitions and generated gRPC code
│   └── raft.proto       # RaftService: RequestVote, AppendEntries, Ping
├── cmd/
│   ├── app/main.go      # Server entrypoint
│   └── client/client.go # Test client (Ping)
├── internal/
│   ├── raft/node.go     # Raft node (RaftServiceServer implementation)
│   └── transport/       # gRPC server with graceful shutdown
└── go.mod
```

## Prerequisites

- Go 1.21+
- `protoc` + `protoc-gen-go` + `protoc-gen-go-grpc` (only needed to regenerate protos)

## Running

**Start the server:**

```bash
go run ./cmd/app
```

The server listens on `:50051`.

**Run the test client:**

```bash
go run ./cmd/client
```

Expected output:

```
Received response: Hello from Raft Node!
```

## Regenerating Protobuf

```bash
cd api/raftpb
bash build.sh
```

## RPC Surface

| Method           | Description                          |
|------------------|--------------------------------------|
| `Ping`           | Health check — returns a hello message |
| `RequestVote`    | Raft leader election (stub)          |
| `AppendEntries`  | Raft log replication (stub)          |

## Graceful Shutdown

The server handles `SIGINT` / `SIGTERM` and attempts a graceful stop with a 5-second timeout before forcing shutdown.
