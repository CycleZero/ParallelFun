package biz

import "context"

type ClientHubRepo interface {
	GetClientInfoByClientId(ctx context.Context, clientId string) (*ClientInfo, error)
}
