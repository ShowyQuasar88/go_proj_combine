package data

import (
	"backend/internal/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

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
		if err := sqlDB.Close(); err != nil {
			log.Errorf("failed to close database resources: %v", err)
		}
	}

	return &Data{
		db:  db,
		log: log,
	}, cleanup, nil
}
