package data

import (
	"github.com/go-kratos/kratos/v2/log"
	clienthubv1 "parallelfun-api/api/clienthub/v1"
	"parallelfun-api/app/server/internal/biz"
)

type serverApi struct {
	cli    clienthubv1.ClientHubClient
	logger log.Logger
}

func NewServerApi(cli clienthubv1.ClientHubClient, logger log.Logger) biz.ServerApi {
	return &serverApi{
		cli:    cli,
		logger: logger,
	}
}
