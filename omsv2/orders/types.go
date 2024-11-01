package main

import (
	"context"
	pb "github.com/showyquasar88/proj-combine/omsv2/common/api"
)

type OrderService interface {
	CreateOrder(ctx context.Context) error
	ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error
}

type OrderStock interface {
	Create(ctx context.Context) error
}
