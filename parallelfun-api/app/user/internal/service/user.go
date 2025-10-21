package service

import (
	"context"
	"errors"
	"gorm.io/gorm"

	pb "parallelfun-api/api/user/v1"
	"parallelfun-api/app/user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdReply, error) {
	user, err := s.uc.GetUser(ctx, uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.GetUserByIdReply{}, nil
		}
		return nil, err
	}

	return &pb.GetUserByIdReply{
		User: []*pb.UserInfo{
			&pb.UserInfo{
				Id:    uint64(user.ID),
				Name:  user.Name,
				Gid:   user.GameId,
				Email: user.Email,
				Role:  int32(user.Role),
			},
		},
	}, nil
}

func (s *UserService) NewUser(ctx context.Context, req *pb.NewUserRequest) (*pb.NewUserReply, error) {
	user := &biz.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := s.uc.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.NewUserReply{
		NewUser: &pb.UserInfo{
			Id:    uint64(createdUser.ID),
			Name:  createdUser.Name,
			Gid:   createdUser.GameId,
			Email: createdUser.Email,
			Role:  int32(createdUser.Role),
		},
	}, nil
}

func (s *UserService) GetUserByGameId(ctx context.Context, req *pb.GetUserByGameIdRequest) (*pb.GetUserByGameIdReply, error) {
	user, err := s.uc.FindByGameId(ctx, req.GameId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.GetUserByGameIdReply{}, nil
		}
		return nil, err
	}

	return &pb.GetUserByGameIdReply{
		User: []*pb.UserInfo{
			&pb.UserInfo{
				Id:    uint64(user.ID),
				Name:  user.Name,
				Gid:   user.GameId,
				Email: user.Email,
				Role:  int32(user.Role),
			},
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	user := &biz.User{
		Model:  gorm.Model{ID: uint(req.User.Id)},
		Name:   req.User.Name,
		Email:  req.User.Email,
		GameId: req.User.Gid,
		Role:   biz.Role(req.User.Role),
	}

	updatedUser, err := s.uc.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserReply{
		User: &pb.UserInfo{
			Id:    uint64(updatedUser.ID),
			Name:  updatedUser.Name,
			Gid:   updatedUser.GameId,
			Email: updatedUser.Email,
			Role:  int32(updatedUser.Role),
		},
		Success: "User updated successfully",
	}, nil
}

func (s *UserService) BatchGetUserByGameId(ctx context.Context, req *pb.BatchGetUserByGameIdRequest) (*pb.BatchGetUserByGameIdReply, error) {
	users, err := s.uc.BatchFindByGameId(ctx, req.Gids)
	if err != nil {
		return nil, err
	}

	userInfos := make([]*pb.UserInfo, 0, len(users))
	for _, user := range users {
		userInfos = append(userInfos, &pb.UserInfo{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Gid:   user.GameId,
			Email: user.Email,
			Role:  int32(user.Role),
		})
	}

	return &pb.BatchGetUserByGameIdReply{
		Users: userInfos,
	}, nil
}

func (s *UserService) BatchGetUserById(ctx context.Context, req *pb.BatchGetUserByIdRequest) (*pb.BatchGetUserByIdReply, error) {
	ids := make([]uint, 0, len(req.Ids))
	for _, id := range req.Ids {
		ids = append(ids, uint(id))
	}

	users, err := s.uc.BatchFindById(ctx, ids)
	if err != nil {
		return nil, err
	}

	userInfos := make([]*pb.UserInfo, 0, len(users))
	for _, user := range users {
		userInfos = append(userInfos, &pb.UserInfo{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Gid:   user.GameId,
			Email: user.Email,
			Role:  int32(user.Role),
		})
	}

	return &pb.BatchGetUserByIdReply{
		Users: userInfos,
	}, nil
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	user := &biz.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	registeredUser, err := s.uc.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		NewUser: &pb.UserInfo{
			Id:    uint64(registeredUser.ID),
			Name:  registeredUser.Name,
			Gid:   registeredUser.GameId,
			Email: registeredUser.Email,
			Role:  int32(registeredUser.Role),
		},
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, token, err := s.uc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		User: &pb.UserInfo{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Gid:   user.GameId,
			Email: user.Email,
			Role:  int32(user.Role),
		},
		Token: token,
	}, nil
}

func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	err := s.uc.Logout(ctx, req.Id, req.Token)
	if err != nil {
		return nil, err
	}

	return &pb.LogoutReply{
		Message: "Logout successful",
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	user := &biz.User{
		Model: gorm.Model{ID: uint(req.Id)},
	}
	err := s.uc.DeleteUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserReply{
		Message: "User deleted successfully",
	}, nil
}
