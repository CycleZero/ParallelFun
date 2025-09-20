package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"parallelfun-api/app/server/internal/biz"
)

type serverRepo struct {
	data *Data
	log  *log.Helper
}

func (s *serverRepo) Save(ctx context.Context, server *biz.Server) (*biz.Server, error) {
	err := s.data.db.Create(server).Error
	return server, err
}

func (s *serverRepo) Update(ctx context.Context, server *biz.Server) (*biz.Server, error) {
	err := s.data.db.Save(server).Error
	return server, err
}

func (s *serverRepo) FindByID(ctx context.Context, id uint) (*biz.Server, error) {
	var server biz.Server
	err := s.data.db.Where("id = ?", id).First(&server).Error
	if err != nil {
		return nil, err
	}
	return &server, nil
}

func (s *serverRepo) ListAll(ctx context.Context) ([]*biz.Server, error) {
	var servers []*biz.Server
	err := s.data.db.Find(&servers).Error
	return servers, err
}

func (s *serverRepo) Delete(ctx context.Context, id uint) error {
	return s.data.db.Delete(&biz.Server{}, id).Error
}

func NewServerRepo(data *Data, logger log.Logger) biz.ServerRepo {
	return &serverRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "server/data")),
	}
}
