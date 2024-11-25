package data

import (
	"backend/internal/conf"
	"backend/internal/data/cache"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewUserRepo,
	NewRedisClient,
	cache.NewCache,
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	log *log.Helper
}

func NewRedisClient(c *conf.Data, logger log.Logger) (*redis.Client, error) {
	// 初始化 redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		ReadTimeout:  time.Duration(c.Redis.ReadTimeout.AsDuration()),
		WriteTimeout: time.Duration(c.Redis.WriteTimeout.AsDuration()),
	})

	// 创建一个带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// 测试连接
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.NewHelper(logger).Errorf("redis connect error: %v", err)
		return nil, err
	}

	log.NewHelper(logger).Info("redis connect success")
	return rdb, nil
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	// 初始化数据库连接
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Errorf("failed opening connection to postgres: %v", err)
		return nil, nil, err
	}

	// 获取通用数据库对象 sql.DB,然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("failed to get db instance: %v", err)
		return nil, nil, err
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	cleanup := func() {
		log.Info("closing database resources")
		if err = sqlDB.Close(); err != nil {
			log.Errorf("failed to close database resources: %v", err)
		}
	}

	return &Data{
		db:  db,
		log: log,
	}, cleanup, nil
}
