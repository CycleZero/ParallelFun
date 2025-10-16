package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"parallelfun-api/app/server/internal/biz"
)

type serverRepo struct {
	data *Data
	log  *log.Helper
}

func (s *serverRepo) Save(ctx context.Context, server *biz.Server) (*biz.Server, error) {
	err := s.data.db.WithContext(ctx).Create(server).Error
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (s *serverRepo) Update(ctx context.Context, server *biz.Server) (*biz.Server, error) {
	err := s.data.db.WithContext(ctx).Save(server).Error
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (s *serverRepo) FindByID(ctx context.Context, id uint) (*biz.Server, error) {
	var server biz.Server
	err := s.data.db.WithContext(ctx).Where("id = ?", id).First(&server).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &server, nil
}

func (s *serverRepo) Delete(ctx context.Context, id uint) error {
	return s.data.db.WithContext(ctx).Delete(&biz.Server{}, id).Error
}

func (s *serverRepo) ListAll(ctx context.Context) ([]*biz.Server, error) {
	var servers []*biz.Server
	err := s.data.db.WithContext(ctx).Find(&servers).Error
	return servers, err
}

func (s *serverRepo) FindByOwnerId(ctx context.Context, ownerId uint) ([]*biz.Server, error) {
	var servers []*biz.Server
	err := s.data.db.WithContext(ctx).Where("owner_id = ?", ownerId).Find(&servers).Error
	return servers, err
}

func (s *serverRepo) FindByAddress(ctx context.Context, address string) (*biz.Server, error) {
	var server biz.Server
	err := s.data.db.WithContext(ctx).Where("address = ?", address).First(&server).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &server, nil
}

func NewServerRepo(data *Data, logger log.Logger) biz.ServerRepo {
	//return &serverRepo{
	//	data: data,
	//	log:  log.NewHelper(log.With(logger, "module", "server/data")),
	//}
	s := &serverRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "server/data")),
	}

	return s
}
