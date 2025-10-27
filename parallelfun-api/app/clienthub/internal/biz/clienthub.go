package biz

import (
	"context"
	"time"
)

type ClientHubUseCase struct {
	connManager *ConnManager
	cliRepo     ClientHubRepo
	ctx         context.Context
}

func NewClientHubUseCase(clientHubRepo ClientHubRepo) *ClientHubUseCase {
	ctx := context.Background()
	return &ClientHubUseCase{
		connManager: NewConnManager(ctx, 256),
		cliRepo:     clientHubRepo,
		ctx:         ctx,
	}
}

func (uc *ClientHubUseCase) SendRpcMsg(reqctx context.Context, clientId string, req *RpcRequest) ([]byte, error) {

	ctx, cancel := context.WithTimeout(reqctx, 10*time.Second)
	defer cancel()
	rc, err := uc.connManager.SendRpcMsg(ctx, clientId, req)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		cancel()
		return nil, ctx.Err()
	case res := <-rc:
		return res, nil
	}
}
