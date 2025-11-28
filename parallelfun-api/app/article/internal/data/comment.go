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

// FindByID 根据ID查找评论
func (c *commentRepo) FindByID(ctx context.Context, id uint64) (*biz.Comment, error) {
	var comment biz.Comment
	if err := c.data.db.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// FindByName 根据名称查找评论（这里假设是根据内容查找）
func (c *commentRepo) FindByName(ctx context.Context, name string) (*biz.Comment, error) {
	var comment biz.Comment
	if err := c.data.db.Where("content LIKE ?", "%"+name+"%").First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// ListAll 获取所有评论
func (c *commentRepo) ListAll(ctx context.Context) ([]*biz.Comment, error) {
	var comments []*biz.Comment
	if err := c.data.db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// Save 创建新评论
func (c *commentRepo) Save(ctx context.Context, comment *biz.Comment) (*biz.Comment, error) {
	if err := c.data.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

// Update 更新评论
func (c *commentRepo) Update(ctx context.Context, comment *biz.Comment) (*biz.Comment, error) {
	if err := c.data.db.Save(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

// Delete 删除评论
func (c *commentRepo) Delete(ctx context.Context, comment *biz.Comment) error {
	return c.data.db.Delete(comment).Error
}

// FindByUserId 根据用户ID查找评论
func (c *commentRepo) FindByUserId(ctx context.Context, userId uint64) ([]*biz.Comment, error) {
	var comments []*biz.Comment
	if err := c.data.db.Where("author_id = ?", userId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// List 分页获取评论列表
// TODO 更换分页查找逻辑
func (c *commentRepo) List(ctx context.Context, p *biz.BatchSelectParam) ([]*biz.Comment, error) {
	var comments []*biz.Comment
	db := c.data.db

	// 如果提供了IDs，则按ID筛选
	if len(p.IDs) > 0 {
		db = db.Where("id IN ?", p.IDs)
	}

	// 按作者ID筛选
	if p.AuthorID != 0 {
		db = db.Where("author_id = ?", p.AuthorID)
	}

	// 按文章ID筛选（假设BatchSelectParam中应该有ArticleID字段）
	// 这里需要您确认是否应在BatchSelectParam中添加ArticleID字段

	// 分页处理
	offset := (p.PageNum - 1) * p.PageSize
	if err := db.Offset(offset).Limit(p.PageSize).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

// NewCommentRepo 创建新的评论仓库实例
func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
