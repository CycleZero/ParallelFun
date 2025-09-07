package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"parallelfun-api/app/user/internal/biz"
)

type userRepo struct {
	data *Data

	log *log.Helper
}

func (r *userRepo) GetRole(ctx context.Context, id uint) (biz.Role, error) {
	var role biz.Role
	err := r.data.db.Where("id=?", id).Select("role").First(&role).Error
	if err != nil {
		return biz.Unknown, err
	}
	return role, err
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user/data")),
	}
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*biz.User, error) {
	var u *biz.User
	r.data.db.Where("id=?", id).First(u)
	return u, nil
}

func (r *userRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	var u []*biz.User
	r.data.db.Find(&u)

	return u, nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	result := r.data.db.Create(u)
	return u, result.Error

}

func (r *userRepo) Update(ctx context.Context, u *biz.User) (*biz.User, error) {
	result := r.data.db.Save(u)
	return u, result.Error
}

func (r *userRepo) Delete(ctx context.Context, u *biz.User) error {
	return r.data.db.Delete(u).Error
}

func (r *userRepo) ListByName(ctx context.Context, name string) ([]*biz.User, error) {
	var u []*biz.User
	result := r.data.db.Where("name=?", name).Find(&u)
	return u, result.Error
}

func (r *userRepo) FindByName(ctx context.Context, name string) (*biz.User, error) {
	var u *biz.User
	result := r.data.db.Where("name=?", name).First(&u)
	return u, result.Error
}
