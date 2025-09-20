package biz

import "context"

type ArticleRepo interface {
	FindByID(ctx context.Context, id uint64) (*Article, error)
	FindByName(ctx context.Context, name string) (*Article, error)
	ListAll(ctx context.Context) ([]*Article, error)
	Save(ctx context.Context, u *Article) (*Article, error)
	Update(ctx context.Context, u *Article) (*Article, error)
	Delete(ctx context.Context, u *Article) error
	FindByUserId(ctx context.Context, userId uint64) ([]*Article, error)

	ListByTitle(ctx context.Context, title string) ([]*Article, error)
	ListByPage(ctx context.Context, offset, limit int) ([]*Article, error)
	FindAuthorById(ctx context.Context, id uint64) (*Author, error)
}

type VideoPostRepo interface {
	FindByID(ctx context.Context, id uint64) (*VideoPost, error)
	FindByName(ctx context.Context, name string) (*VideoPost, error)
	ListAll(ctx context.Context) ([]*VideoPost, error)
	Save(ctx context.Context, u *VideoPost) (*VideoPost, error)
	Update(ctx context.Context, u *VideoPost) (*VideoPost, error)
	Delete(ctx context.Context, u *VideoPost) error
	FindByUserId(ctx context.Context, userId uint64)
}

type CommentRepo interface {
	FindByID(ctx context.Context, id uint64) (*Comment, error)
	FindByName(ctx context.Context, name string) (*Comment, error)
	ListAll(ctx context.Context) ([]*Comment, error)
	Save(ctx context.Context, u *Comment) (*Comment, error)
	Update(ctx context.Context, u *Comment) (*Comment, error)
	Delete(ctx context.Context, u *Comment) error
	FindByUserId(ctx context.Context, userId uint64) ([]*Comment, error)
	List(ctx context.Context, p *BatchSelectParam) ([]*Comment, error)
}
