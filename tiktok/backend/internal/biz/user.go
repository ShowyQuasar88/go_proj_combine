package biz

import (
	v1 "backend/api/v1"
	"backend/internal/data/cache"
	"backend/internal/pkg/auth"
	"backend/internal/pkg/crypto"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
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
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	SaveUserToken(ctx context.Context, userID, token string, expire time.Duration) error
	GetUserToken(ctx context.Context, userID string) (string, error)
}

// UserUseCase 相当于用户业务逻辑和底层数据交互的中间代理
type UserUseCase struct {
	repo      UserRepo
	crypto    *crypto.Crypto
	jwtHelper *auth.JWTHelper
	cache     *cache.Cache
	log       *log.Helper
}

// NewUserUseCase 返回一个中间代理，所有需要进行数据处理的操作经过这个中间代理
func NewUserUseCase(repo UserRepo, crypto *crypto.Crypto, jwtHelper *auth.JWTHelper, cache *cache.Cache, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:      repo,
		crypto:    crypto,
		cache:     cache,
		jwtHelper: jwtHelper,
		log:       log.NewHelper(logger),
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

func (uc *UserUseCase) Login(ctx context.Context, u *User) (*v1.LoginResponse, error) {
	uc.log.WithContext(ctx).Infof("[login start]: %v", u.Username)
	user, err := uc.repo.GetUserByUsername(ctx, u.Username)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("[get user info error]: %v, error: %v", u.Username, err)
		return nil, err
	}

	encryptPassword, err := uc.crypto.HashPassword(u.Password)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("[hash password error]: %v", err)
		return nil, err
	}
	if uc.crypto.CheckPasswordHash(encryptPassword, user.Password) {
		return nil, gorm.ErrRecordNotFound
	}

	token, err := uc.jwtHelper.GenerateToken(user.ID, user.Username)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("[generate token failed]: %v, error: %v", user.Username, err)
		return nil, err
	}

	expire := 24 * time.Hour
	if err := uc.cache.SetUserToken(ctx, user.ID, token, expire); err != nil {
		uc.log.WithContext(ctx).Errorf("[set cache token failed]: %v, error: %v", user.Username, err)
		return nil, err
	}
	uc.log.WithContext(ctx).Infof("[login finish]: %v", u.Username)

	return &v1.LoginResponse{
		Token:     token,
		ExpiresIn: int64(expire),
		UserInfo: &v1.UserInfo{
			Id:       user.ID,
			Username: user.Username,
		},
	}, nil
}
