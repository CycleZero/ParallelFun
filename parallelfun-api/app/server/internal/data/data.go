package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"parallelfun-api/app/server/internal/biz"
	"parallelfun-api/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewServerRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	dsn := c.Database.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&biz.Server{})
	if err != nil {
		return nil, nil, err
	}
	return &Data{db: db}, cleanup, nil
}
