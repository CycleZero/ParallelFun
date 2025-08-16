package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Name)
	return uc.repo.Save(ctx, u)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id uint) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %d", id)
	return uc.repo.FindByID(ctx, id)
}
