package service

import (
	"context"
	"parallelfun-api/app/server/internal/biz"

	pb "parallelfun-api/api/server/v1"
)

type ServerServiceService struct {
	pb.UnimplementedServerServiceServer

	uc *biz.ServerUsecase
}

func NewServerServiceService(uc *biz.ServerUsecase) *ServerServiceService {
	return &ServerServiceService{uc: uc}
}

func (s *ServerServiceService) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	return &pb.CreateServerReply{}, nil
}
func (s *ServerServiceService) UpdateServer(ctx context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
	return &pb.UpdateServerReply{}, nil
}
func (s *ServerServiceService) DeleteServer(ctx context.Context, req *pb.DeleteServerRequest) (*pb.DeleteServerReply, error) {
	return &pb.DeleteServerReply{}, nil
}
func (s *ServerServiceService) GetServer(ctx context.Context, req *pb.GetServerRequest) (*pb.GetServerReply, error) {
	return &pb.GetServerReply{}, nil
}
func (s *ServerServiceService) ListServer(ctx context.Context, req *pb.ListServerRequest) (*pb.ListServerReply, error) {
	return &pb.ListServerReply{}, nil
}
