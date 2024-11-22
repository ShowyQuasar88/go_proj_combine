package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       string
	Username string
	Password string
	Phone    string
	Email    string
}

// UserRepo 数据访问层的抽象接口
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

// UserUseCase 相当于用户业务逻辑和底层数据交互的中间代理
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUseCase 返回一个中间代理，所有需要进行数据处理的操作经过这个中间代理
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Register(ctx context.Context, u *User) error {
	uc.log.WithContext(ctx).Infof("[register start]: %v", u.Username)
	err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("[register failed]: %v, error: %v", u.Username, err)
		return err
	}
	uc.log.WithContext(ctx).Infof("[register success]: %v", u.Username)
	return nil
}
