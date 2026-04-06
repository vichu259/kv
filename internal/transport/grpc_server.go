package transport

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vichu259/kv/api/raftpb"
	"github.com/vichu259/kv/internal/raft"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	Node   *raft.Node
	server *grpc.Server
	port   string
}

func NewGRPCServer(port string, node *raft.Node) *GRPCServer {
	server := grpc.NewServer()
	raftpb.RegisterRaftServiceServer(server, node)
	return &GRPCServer{
		Node:   node,
		server: server,
		port:   port,
	}
}

func (g *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		return err
	}

	log.Println("🚀 Raft node started on port", g.port)

	go func() {
		if err := g.server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	g.waitForShutdown()
	return nil
}

func (g *GRPCServer) waitForShutdown() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		g.server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("✅ Server stopped gracefully")
	case <-ctx.Done():
		log.Println("⚠️ Force stopping server")
		g.server.Stop()
	}
}
