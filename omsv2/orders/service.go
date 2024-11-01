package main

import (
	"context"
	"github.com/showyquasar88/proj-combine/omsv2/common"
	pb "github.com/showyquasar88/proj-combine/omsv2/common/api"
)

type Service struct {
	store OrderStock
}

func NewService(store OrderStock) *Service {
	return &Service{store}
}

func (s *Service) CreateOrder(ctx context.Context) error {
	return nil
}

func (s *Service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}
	mergeItemsQuantities(p.Items)

	return nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make([]*pb.ItemsWithQuantity, 0)
	for _, item := range items {
		found := false
		for _, q := range merged {
			if q.ID == item.ID {
				q.Quantity += item.Quantity
				found = true
			}
		}
		if !found {
			merged = append(merged, item)
		}
	}
	return merged
}
