package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{rdb: rdb}
}

// 用户token相关操作
const (
	UserTokenKeyPrefix = "user:token:"  // 用户token前缀
	DefaultExpiration  = 24 * time.Hour // 默认过期时间
)

// SetUserToken 存储用户token
func (c *Cache) SetUserToken(ctx context.Context, userID, token string, expire time.Duration) error {
	key := UserTokenKeyPrefix + userID
	return c.rdb.Set(ctx, key, token, expire).Err()
}

// GetUserToken 获取用户token
func (c *Cache) GetUserToken(ctx context.Context, userID string) (string, error) {
	key := UserTokenKeyPrefix + userID
	return c.rdb.Get(ctx, key).Result()
}

// DelUserToken 删除用户token(登出时使用)
func (c *Cache) DelUserToken(ctx context.Context, userID string) error {
	key := UserTokenKeyPrefix + userID
	return c.rdb.Del(ctx, key).Err()
}

// IsUserTokenExist 检查用户token是否存在
func (c *Cache) IsUserTokenExist(ctx context.Context, userID string) bool {
	key := UserTokenKeyPrefix + userID
	exists, _ := c.rdb.Exists(ctx, key).Result()
	return exists > 0
}
