package server

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	userv1 "parallelfun-api/api/user/v1"
	"parallelfun-api/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
	NewRegistrar,
	NewDiscovery,
	NewUserClient,
)

func NewUserClient(conf *conf.Registry, dis registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint("discovery:///parallelfun.service.user.grpc"),
		grpc.WithDiscovery(dis),
	)
	if err != nil {
		log.Println("failed to dial:", err)
		return nil
	}

	cli := userv1.NewUserClient(conn)
	return cli

}
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	return NewRegistryClient(conf)
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	return NewRegistryClient(conf)
}

func NewRegistryClient(conf *conf.Registry) *nacos.Registry {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.GetAddress(), conf.Nacos.GetPort()),
	}
	cc := &constant.ClientConfig{
		Username: conf.Nacos.GetUsername(), // 获取用户名
		Password: conf.Nacos.GetPassword(), // 获取密码
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  cc,
		},
	)
	if err != nil {
		panic(err)
	}
	r := nacos.New(client)
	return r
}
