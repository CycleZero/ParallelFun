package data

import (
	"context"
	"parallelfun-api/app/clienthub/internal/biz"
)

type clientHubRepo struct {
	data *Data
}

func NewClientHubRepo(data *Data) biz.ClientHubRepo {
	return &clientHubRepo{
		data: data,
	}
}

func (r *clientHubRepo) GetClientInfoByClientId(ctx context.Context, clientId string) (*biz.ClientInfo, error) {
	return nil, nil
}
