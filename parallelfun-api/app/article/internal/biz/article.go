package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleUsecase) GetArticleById(ctx context.Context, id uint64) (*Article, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *ArticleUsecase) GetArticleList(ctx context.Context) ([]*Article, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *ArticleUsecase) NewArticle(ctx context.Context, a *Article) (*Article, error) {
	return uc.repo.Save(ctx, a)
}

func (uc *ArticleUsecase) UpdateArticle(ctx context.Context, a *Article) (*Article, error) {
	return uc.repo.Update(ctx, a)
}

func (uc *ArticleUsecase) DeleteArticle(ctx context.Context, a *Article) error {
	return uc.repo.Delete(ctx, a)
}

func (uc *ArticleUsecase) GetArticleListByUserId(ctx context.Context, userId uint64) ([]*Article, error) {
	return uc.repo.FindByUserId(ctx, userId)
}

func (uc *ArticleUsecase) GetArticleListByTitle(ctx context.Context, title string) ([]*Article, error) {
	return uc.repo.ListByTitle(ctx, title)
}

func (uc *ArticleUsecase) GetArticleListByPage(ctx context.Context, offset, limit int) ([]*Article, error) {
	return uc.repo.ListByPage(ctx, offset, limit)
}

func (uc *ArticleUsecase) GetAuthorById(ctx context.Context, id uint64) (*Author, error) {
	return uc.repo.FindAuthorById(ctx, id)
}
