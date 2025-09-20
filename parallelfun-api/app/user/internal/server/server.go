package server

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"parallelfun-api/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer,
	NewHTTPServer,
	NewRegistrar,
	NewDiscovery,
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
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

func NewDiscovery(conf *conf.Registry) registry.Discovery {
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
