package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	userv1 "parallelfun-api/api/user/v1"
	"parallelfun-api/app/server/internal/biz"
)

type userRepo struct {
	client userv1.UserClient
	logger *log.Helper
}

func (repo *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	reply, err := repo.client.UpdateUser(ctx, &userv1.UpdateUserRequest{User: &userv1.UserInfo{
		Id:    uint64(user.ID),
		Name:  user.Name,
		Gid:   user.GameId,
		Email: user.Email,
		Role:  int32(user.Role),
	}})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Model:  gorm.Model{ID: uint(reply.User.Id)},
		Name:   reply.User.Name,
		GameId: reply.User.Gid,
		Email:  reply.User.Email,
		Role:   biz.Role(reply.User.Role),
	}, nil
}

func (repo *userRepo) FindByID(ctx context.Context, u uint) (*biz.User, error) {
	reply, err := repo.client.GetUserById(ctx, &userv1.GetUserByIdRequest{Id: uint64(u)})
	if err != nil || len(reply.User) <= 0 {
		return nil, err
	}
	return &biz.User{
		Model: gorm.Model{
			ID: uint(reply.User[0].Id),
		},
		Name:   reply.User[0].Name,
		Email:  reply.User[0].Email,
		Role:   biz.Role(reply.User[0].Role),
		GameId: reply.User[0].Gid,
	}, nil
}

func (repo *userRepo) FindByGameId(ctx context.Context, s string) (*biz.User, error) {
	reply, err := repo.client.GetUserByGameId(ctx, &userv1.GetUserByGameIdRequest{GameId: s})
	if err != nil || len(reply.User) <= 0 {
		return nil, err
	}
	return &biz.User{
		Model: gorm.Model{
			ID: uint(reply.User[0].Id),
		},
		Name:   reply.User[0].Name,
		Email:  reply.User[0].Email,
		Role:   biz.Role(reply.User[0].Role),
		GameId: reply.User[0].Gid,
	}, nil
}

func (repo *userRepo) BatchFindByGameId(ctx context.Context, strings []string) ([]*biz.User, error) {
	reply, err := repo.client.BatchGetUserByGameId(ctx, &userv1.BatchGetUserByGameIdRequest{Gids: strings})
	if err != nil {
		return nil, err
	}

	users := make([]*biz.User, 0, len(reply.Users))
	for _, userInfo := range reply.Users {
		users = append(users, &biz.User{
			Model: gorm.Model{
				ID: uint(userInfo.Id),
			},
			Name:   userInfo.Name,
			Email:  userInfo.Email,
			Role:   biz.Role(userInfo.Role),
			GameId: userInfo.Gid,
		})
	}
	return users, nil
}

func (repo *userRepo) BatchFindById(ctx context.Context, uints []uint) ([]*biz.User, error) {
	ids := make([]uint64, 0, len(uints))
	for _, id := range uints {
		ids = append(ids, uint64(id))
	}

	reply, err := repo.client.BatchGetUserById(ctx, &userv1.BatchGetUserByIdRequest{Ids: ids})
	if err != nil {
		return nil, err
	}

	users := make([]*biz.User, 0, len(reply.Users))
	for _, userInfo := range reply.Users {
		users = append(users, &biz.User{
			Model: gorm.Model{
				ID: uint(userInfo.Id),
			},
			Name:   userInfo.Name,
			Email:  userInfo.Email,
			Role:   biz.Role(userInfo.Role),
			GameId: userInfo.Gid,
		})
	}
	return users, nil
}

func NewUserRepo(client userv1.UserClient, logger log.Logger) biz.UserRepo {
	return &userRepo{
		client: client,
		logger: log.NewHelper(log.With(logger, "module", "data.user")),
	}
}
