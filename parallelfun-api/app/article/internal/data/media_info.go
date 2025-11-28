package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"parallelfun-api/app/article/internal/biz"
)

type mediaInfoRepo struct {
	data *Data
	log  *log.Helper
}

func NewMediaRepo(data *Data, logger log.Logger) biz.MediaInfoRepo {
	return &mediaInfoRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "article/data")),
	}
}

// FindByID 根据ID查找媒体信息
func (m *mediaInfoRepo) FindByID(ctx context.Context, id uint64) (*biz.MediaInfo, error) {
	var mediaInfo biz.MediaInfo
	if err := m.data.db.WithContext(ctx).Where("id = ?", id).First(&mediaInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &mediaInfo, nil
}

// ListAll 获取所有媒体信息
func (m *mediaInfoRepo) ListAll(ctx context.Context) ([]*biz.MediaInfo, error) {
	var mediaInfos []*biz.MediaInfo
	if err := m.data.db.WithContext(ctx).Find(&mediaInfos).Error; err != nil {
		return nil, err
	}
	return mediaInfos, nil
}

// Save 保存媒体信息
func (m *mediaInfoRepo) Save(ctx context.Context, u *biz.MediaInfo) (*biz.MediaInfo, error) {
	err := m.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(u).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Update 更新媒体信息
func (m *mediaInfoRepo) Update(ctx context.Context, u *biz.MediaInfo) (*biz.MediaInfo, error) {
	err := m.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", u.ID).Updates(u).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete 删除媒体信息
func (m *mediaInfoRepo) Delete(ctx context.Context, u *biz.MediaInfo) error {
	return m.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Delete(u).Error
	})
}

// FindByArticleId 根据文章ID查找媒体信息
func (m *mediaInfoRepo) FindByArticleId(ctx context.Context, articleId uint64) ([]*biz.MediaInfo, error) {
	var mediaInfos []*biz.MediaInfo
	if err := m.data.db.WithContext(ctx).Where("article_id = ?", articleId).Find(&mediaInfos).Error; err != nil {
		return nil, err
	}
	return mediaInfos, nil
}

// BatchSave 批量保存媒体信息
func (m *mediaInfoRepo) BatchSave(ctx context.Context, infos []*biz.MediaInfo) ([]*biz.MediaInfo, error) {
	if len(infos) == 0 {
		return infos, nil
	}

	err := m.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 使用批量插入提高效率
		if err := tx.CreateInBatches(infos, 100).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return infos, nil
}

// DeleteByArticleId 根据文章ID删除媒体信息
func (m *mediaInfoRepo) DeleteByArticleId(ctx context.Context, articleId uint64) error {
	return m.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("article_id = ?", articleId).Delete(&biz.MediaInfo{}).Error
	})
}
