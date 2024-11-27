package service

import (
	v1 "backend/api/v1"
	"backend/internal/biz"
	"backend/internal/pkg/response"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/anypb"
	"gorm.io/gorm"
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
		Phone:    req.Phone,
		Email:    req.Email,
	}
	if err := s.uc.Register(ctx, user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return response.SuccessWithMsg("用户已存在"), nil
		}
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, err.Error()), err
	}
	return response.SuccessWithMsg("注册成功"), nil
}

func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.Response, error) {
	user := &biz.User{
		Username: req.Username,
		Password: req.Password,
	}
	userResp, err := s.uc.Login(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(v1.ErrorCode_USER_NOT_FOUND, "用户不存在或密码错误"), nil
		}
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, err.Error()), nil
	}

	data, err := anypb.New(userResp)
	if err != nil {
		return response.Error(v1.ErrorCode_SYSTEM_ERROR, err.Error()), nil
	}

	return response.SuccessWithData(data), nil
}
