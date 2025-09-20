//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"parallelfun-api/app/social/internal/biz"
	"parallelfun-api/app/social/internal/conf"
	"parallelfun-api/app/social/internal/data"
	"parallelfun-api/app/social/internal/server"
	"parallelfun-api/app/social/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
