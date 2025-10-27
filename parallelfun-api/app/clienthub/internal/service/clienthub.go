package service

import (
	"context"
	"parallelfun-api/app/clienthub/internal/biz"

	pb "parallelfun-api/api/clienthub/v1"
)

type ClientHubService struct {
	pb.UnimplementedClientHubServer
	uc *biz.ClientHubUseCase
}

func NewClientHubService(uc *biz.ClientHubUseCase) *ClientHubService {
	return &ClientHubService{
		uc: uc,
	}
}

func (s *ClientHubService) SendRpcCommand(ctx context.Context, req *pb.SendRpcCommandRequest) (*pb.SendRpcCommandReply, error) {
	rpcreq := &biz.RpcRequest{
		Method: req.Command,
		Params: req.Params,
	}
	res, err := s.uc.SendRpcMsg(ctx, req.ClientId, rpcreq)
	if err != nil {
		return nil, err
	}
	return &pb.SendRpcCommandReply{
		Message: string(res),
	}, nil
}
