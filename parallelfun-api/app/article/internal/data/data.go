package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"parallelfun-api/app/article/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	minio "github.com/minio/minio-go/v7"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	userv1 "parallelfun-api/api/user/v1"
	"parallelfun-api/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewArticleRepo,
	NewUserClient,
	NewMediaRepo,
	NewCommentRepo,
	NewUserRepo,
)

// Data .
type Data struct {
	db          *gorm.DB
	ucli        userv1.UserClient
	minioClient *minio.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, ucli userv1.UserClient, minioClient *minio.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	dsn := c.Database.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&biz.Article{}, &biz.VideoPost{}, &biz.Comment{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &Data{db: db, ucli: ucli, minioClient: minioClient}, cleanup, nil
}

func NewUserClient(conf *conf.Registry, dis registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///parallelfun.service.user.grpc"),
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
		//panic(err)
		log.Error("new user client error", err)
		return nil
	}

	c := userv1.NewUserClient(conn)
	return c
}
