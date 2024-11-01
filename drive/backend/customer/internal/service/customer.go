package service

import (
	"context"
	"customer/api/customer/v1"
)

type CustomerService struct {
	v1.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *v1.GetVerifyCodeReq) (*v1.GetVerifyCodeResp, error) {
	return &v1.GetVerifyCodeResp{}, nil
}
