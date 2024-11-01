package main

import (
	"github.com/showyquasar88/proj-combine/omsv2/common"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	GrpcAddr = common.EnvString("GRPC_ADDR", "localhost:9001")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	srv := NewService(store)
	NewGrpcHandler(grpcServer, srv)

	// srv.CreateOrder(context.Background())

	log.Println("GRPC Server Started at ", GrpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
