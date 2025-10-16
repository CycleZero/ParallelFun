package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"parallelfun-api/app/article/internal/biz"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func (c *commentRepo) FindByID(ctx context.Context, id uint64) (*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepo) FindByName(ctx context.Context, name string) (*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepo) ListAll(ctx context.Context) ([]*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepo) Save(ctx context.Context, u *biz.Comment) (*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepo) Update(ctx context.Context, u *biz.Comment) (*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepo) Delete(ctx context.Context, u *biz.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepo) FindByUserId(ctx context.Context, userId uint64) ([]*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *commentRepo) List(ctx context.Context, p *biz.BatchSelectParam) ([]*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
