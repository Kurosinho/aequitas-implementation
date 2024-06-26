package main

import (
	"context"
	"log"
	"magisterium/stayalive"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	stayalive.UnimplementedStayAliveServiceServer
}

func (m *gRPCServer) StayAlive(ctx context.Context, request *stayalive.StayAliveRequest) (*stayalive.StayAliveResponse, error) {
	log.Println("It's alive!")
	return &stayalive.StayAliveResponse{AliveResp: bool(true)}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:2222")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	gRPCServer := &gRPCServer{}
	stayalive.RegisterStayAliveServiceServer(s, gRPCServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
