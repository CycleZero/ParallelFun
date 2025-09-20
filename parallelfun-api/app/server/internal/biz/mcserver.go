package biz

import "github.com/go-kratos/kratos/v2/log"

type ServerUsecase struct {
	log  *log.Helper
	repo ServerRepo
}

func NewServerUsecase(repo ServerRepo, logger log.Logger) *ServerUsecase {
	return &ServerUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "server/biz")),
	}
}
