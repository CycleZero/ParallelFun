package service

import (
	"context"
	"parallelfun-api/app/user/internal/biz"

	pb "parallelfun-api/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, _ := s.uc.GetUser(ctx, uint(req.Id))
	return &pb.GetUserReply{Id: uint64(u.ID), Name: u.Name}, nil
}
func (s *UserService) NewUser(ctx context.Context, req *pb.NewUserRequest) (*pb.NewUserReply, error) {
	return &pb.NewUserReply{}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
