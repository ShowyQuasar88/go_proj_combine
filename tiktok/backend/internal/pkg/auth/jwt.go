package auth

import (
	"backend/internal/conf"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	Secret string
	Expire int32
	Issuer string
}

type JWTHelper struct {
	secret []byte
	expire time.Duration
	issuer string
}

// NewJWTConfig 从 Security 配置创建 JWT 配置
func NewJWTConfig(c *conf.Security) *JWTConfig {
	return &JWTConfig{
		Secret: c.Jwt.Secret,
		Expire: c.Jwt.Expire,
		Issuer: "tiktok.api", // 可以从配置文件读取
	}
}

// NewJWTHelper 从 JWT 配置创建 Helper
func NewJWTHelper(c *JWTConfig) *JWTHelper {
	return &JWTHelper{
		secret: []byte(c.Secret),
		expire: time.Duration(c.Expire) * time.Second,
		issuer: c.Issuer,
	}
}

func (j *JWTHelper) GenerateToken(userID, username string) (string, error) {
	now := time.Now()
	claims := Claims{
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.expire)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    j.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}
