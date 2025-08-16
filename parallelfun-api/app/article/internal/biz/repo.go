package biz

import "context"

type ArticleRepo interface {
	FindByID(ctx context.Context, id int64) (*Article, error)
	FindByName(ctx context.Context, name string) (*Article, error)
}
