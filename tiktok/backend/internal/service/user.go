package service

import (
	v1 "backend/api/v1"
	"backend/internal/biz"
	"context"
)

// UserService 是用户服务的实现
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUseCase
}

// NewUserService 创建一个用户服务
func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.Response, error) {
	user := &biz.User{
		Username: req.Username,
		Password: req.Password,
	}
	if err := s.uc.Register(ctx, user); err != nil {
		return nil, err
	}
	return &v1.Response{}, nil
}
