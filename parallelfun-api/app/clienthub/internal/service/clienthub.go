package service

import (
	"context"

	pb "parallelfun-api/api/clienthub/v1"
)

type ClientHubService struct {
	pb.UnimplementedClientHubServer
}

func NewClientHubService() *ClientHubService {
	return &ClientHubService{}
}

func (s *ClientHubService) SendRpcCommand(ctx context.Context, req *pb.SendRpcCommandRequest) (*pb.SendRpcCommandReply, error) {
	return &pb.SendRpcCommandReply{}, nil
}
