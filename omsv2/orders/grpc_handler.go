package main

import (
	"context"
	pb "github.com/showyquasar88/proj-combine/omsv2/common/api"
	"google.golang.org/grpc"
	"log"
)

type GrpcHandler struct {
	pb.UnimplementedOrderServiceServer

	service OrderService
}

func NewGrpcHandler(grpcServer *grpc.Server, srv OrderService) {
	h := &GrpcHandler{
		service: srv,
	}
	pb.RegisterOrderServiceServer(grpcServer, h)
}

func (h *GrpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("New order received! Order %v", p)
	resp := &pb.CreateOrderResponse{
		Order: &pb.Order{
			ID: "42",
		},
	}
	return resp, nil
}
