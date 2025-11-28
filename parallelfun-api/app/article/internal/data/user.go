package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	userv1 "parallelfun-api/api/user/v1"
	"parallelfun-api/app/article/internal/biz"
)

type userRepo struct {
	ucli userv1.UserClient
	log  *log.Helper
}

func (u *userRepo) FindByGid(ctx context.Context, name string) (*biz.User, error) {
	res, err := u.ucli.BatchGetUserByGameId(ctx, &userv1.BatchGetUserByGameIdRequest{Gids: []string{name}})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Model: gorm.Model{
			ID: uint(res.Users[0].Id),
		},
		Email:  res.Users[0].Email,
		GameId: res.Users[0].Gid,
		Name:   res.Users[0].Name,
		Role:   biz.Role(res.Users[0].Role),
	}, nil
}

func (u *userRepo) FindById(ctx context.Context, id uint64) (*biz.User, error) {
	res, err := u.ucli.GetUserById(ctx, &userv1.GetUserByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Model: gorm.Model{
			ID: uint(res.User.Id),
		},
		Email:  res.User.Email,
		GameId: res.User.Gid,
		Name:   res.User.Name,
		Role:   biz.Role(res.User.Role),
	}, nil
}

func NewUserRepo(ucli userv1.UserClient, logger log.Logger) biz.UserRepo {
	return &userRepo{
		ucli: ucli,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}

}
