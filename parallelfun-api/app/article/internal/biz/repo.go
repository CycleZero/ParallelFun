package biz

import "context"

type BaseRepo[T any] interface {
	FindByID(ctx context.Context, id uint64) (T, error)
	ListAll(ctx context.Context) ([]T, error)
	Save(ctx context.Context, u T) (T, error)
	Update(ctx context.Context, u T) (T, error)
	Delete(ctx context.Context, u T) error
}

type ArticleRepo interface {
	BaseRepo[*Article]
	FindByName(ctx context.Context, name string) (*Article, error)
	FindByUserId(ctx context.Context, userId uint64) ([]*Article, error)
	ListByTitle(ctx context.Context, title string) ([]*Article, error)
	ListByPage(ctx context.Context, offset, limit int) ([]*Article, error)
	FindAuthorById(ctx context.Context, id uint64) (*Author, error)
}

type VideoPostRepo interface {
	BaseRepo[*VideoPost]
	FindByName(ctx context.Context, name string) (*VideoPost, error)
	FindByUserId(ctx context.Context, userId uint64)
}

type CommentRepo interface {
	BaseRepo[*Comment]
	FindByName(ctx context.Context, name string) (*Comment, error)
	FindByUserId(ctx context.Context, userId uint64) ([]*Comment, error)
	List(ctx context.Context, p *BatchSelectParam) ([]*Comment, error)
}
