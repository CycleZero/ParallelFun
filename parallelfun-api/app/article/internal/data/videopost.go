package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"parallelfun-api/app/article/internal/biz"
)

type videoPostRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoPostRepo(data *Data, logger log.Logger) biz.VideoPostRepo {
	return &videoPostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (v *videoPostRepo) FindByID(ctx context.Context, id uint64) (*biz.VideoPost, error) {

	//TODO implement me
	panic("implement me")
}

func (v *videoPostRepo) FindByName(ctx context.Context, name string) (*biz.VideoPost, error) {
	//TODO implement me
	panic("implement me")
}

func (v *videoPostRepo) ListAll(ctx context.Context) ([]*biz.VideoPost, error) {
	//TODO implement me
	panic("implement me")
}

func (v *videoPostRepo) Save(ctx context.Context, u *biz.VideoPost) (*biz.VideoPost, error) {
	//TODO implement me
	panic("implement me")
}

func (v *videoPostRepo) Update(ctx context.Context, u *biz.VideoPost) (*biz.VideoPost, error) {
	//TODO implement me
	panic("implement me")
}

func (v *videoPostRepo) Delete(ctx context.Context, u *biz.VideoPost) error {
	//TODO implement me
	panic("implement me")
}

func (v *videoPostRepo) FindByUserId(ctx context.Context, userId uint64) {
	//TODO implement me
	panic("implement me")
}
