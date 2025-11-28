package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	userv1 "parallelfun-api/api/user/v1"
	"parallelfun-api/app/article/internal/biz"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (r *articleRepo) FindAuthorById(ctx context.Context, id uint64) (*biz.Author, error) {
	res, err := r.data.ucli.GetUserById(ctx, &userv1.GetUserByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &biz.Author{
		ID:   res.GetUser().GetId(),
		Name: res.GetUser().GetName(),
	}, nil
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "article/data")),
	}
}

func (r *articleRepo) FindByID(ctx context.Context, id uint64) (*biz.Article, error) {
	var article biz.Article
	err := r.data.db.Where("id = ?", id).First(&article).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find article by id: %d", id)
	}
	return &article, nil
}

func (r *articleRepo) FindByName(ctx context.Context, name string) (*biz.Article, error) {
	var article biz.Article
	err := r.data.db.Where("title = ?", name).First(&article).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find article by name: %s", name)
	}
	return &article, nil
}

func (r *articleRepo) ListAll(ctx context.Context) ([]*biz.Article, error) {
	var articles []*biz.Article
	err := r.data.db.Find(&articles).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to list all articles")
	}
	return articles, nil
}

func (r *articleRepo) Save(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	err := r.data.db.Create(a).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to save article")
	}
	return a, nil
}

func (r *articleRepo) Update(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	err := r.data.db.Save(a).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to update article")
	}
	return a, nil
}

func (r *articleRepo) Delete(ctx context.Context, a *biz.Article) error {
	err := r.data.db.Delete(a).Error
	if err != nil {
		return errors.Wrap(err, "failed to delete article")
	}
	return nil
}

func (r *articleRepo) FindByUserId(ctx context.Context, userId uint64) ([]*biz.Article, error) {
	var articles []*biz.Article
	err := r.data.db.Where("user_id = ?", userId).Find(&articles).Error
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find articles by user id: %d", userId)
	}
	return articles, nil
}

func (r *articleRepo) ListByTitle(ctx context.Context, title string) ([]*biz.Article, error) {
	var articles []*biz.Article
	err := r.data.db.Where("title LIKE ?", "%"+title+"%").Find(&articles).Error
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find articles by title: %s", title)
	}
	return articles, nil
}

func (r *articleRepo) ListByPage(ctx context.Context, offset, limit int) ([]*biz.Article, error) {
	var articles []*biz.Article
	err := r.data.db.Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list articles by page, offset: %d, limit: %d", offset, limit)
	}
	return articles, nil
}
