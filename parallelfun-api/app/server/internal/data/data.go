package data

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	userv1 "parallelfun-api/api/user/v1"
	"parallelfun-api/app/server/internal/biz"
	"parallelfun-api/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewServerRepo,
	NewUserRepo,
)

// Data .
type Data struct {
	db         *gorm.DB
	rdb        *redis.Client
	userClient userv1.UserClient
	log        *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, userClient userv1.UserClient) (*Data, func(), error) {
	dsn := c.Database.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&biz.Server{})
	if err != nil {
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr, // Redis服务器地址和端口
		Password: "",           // Redis访问密码，如果没有可以为空字符串
		DB:       0,            // 使Redis数据库编号，默认为0用的
	})
	cleanup := func() {
		d, _ := db.DB()
		err := d.Close()
		if err != nil {
			log.Info("failed to close the data resources", err)
		}
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb, userClient: userClient}, cleanup, nil
}
