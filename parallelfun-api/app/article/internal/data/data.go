package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	userv1 "parallelfun-api/api/user/v1"
	"parallelfun-api/app/article/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewArticleRepo,
	NewUserClient,
)

// Data .
type Data struct {
	db   *gorm.DB
	ucli userv1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, ucli userv1.UserClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	dsn := "user=root password=poyuan666 dbname=parallelfun port=35432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Data{db: db, ucli: ucli}, cleanup, nil
}

func NewUserClient(conf *conf.Registry, dis registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///beer.user.service"),
		grpc.WithDiscovery(dis),
		//grpc.WithMiddleware(
		//	tracing.Client(tracing.WithTracerProvider(tp)),
		//	recovery.Recovery(),
		//	jwt.Client(func(token *jwt2.Token) (interface{}, error) {
		//		return []byte(ac.ServiceKey), nil
		//	}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		//),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}
