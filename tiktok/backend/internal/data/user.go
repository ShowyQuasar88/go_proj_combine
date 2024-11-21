package data

import (
	"backend/internal/biz"
	"backend/internal/pkg/crypto"
	"backend/internal/pkg/utils"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       string `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
}

func (User) TableName() string {
	return "tb_user"
}

type userRepo struct {
	data   *Data
	crypto *crypto.Crypto
	log    *log.Helper
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	hashedPassword, err := r.crypto.HashPassword(u.Password)
	if err != nil {
		return err
	}

	encryptedPhone, err := r.crypto.Encrypt(u.Phone)
	if err != nil {
		return err
	}

	encryptedEmail, err := r.crypto.Encrypt(u.Email)
	if err != nil {
		return err
	}

	inserted := &User{
		ID:       utils.NewUUID(),
		Username: u.Username,
		Password: hashedPassword,
		Phone:    encryptedPhone,
		Email:    encryptedEmail,
	}

	return r.data.db.WithContext(ctx).Create(inserted).Error
}

func NewUserRepo(data *Data, crypto *crypto.Crypto, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:   data,
		crypto: crypto,
		log:    log.NewHelper(logger),
	}
}
