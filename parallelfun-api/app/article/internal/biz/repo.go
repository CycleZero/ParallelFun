package biz

import "context"

type ArticleRepo interface {
	FindByID(ctx context.Context, id uint64) (*Article, error)
	FindByName(ctx context.Context, name string) (*Article, error)
	ListAll(ctx context.Context) ([]*Article, error)
	Save(ctx context.Context, u *Article) (*Article, error)
	Update(ctx context.Context, u *Article) (*Article, error)
	Delete(ctx context.Context, u *Article) error
	FindByUserId(ctx context.Context, userId uint64)
}
