package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ServerUsecase struct {
	log        *log.Helper
	serverRepo ServerRepo
	userRepo   UserRepo
}

func NewServerUsecase(repo ServerRepo, logger log.Logger, userRepo UserRepo) *ServerUsecase {
	return &ServerUsecase{
		serverRepo: repo,
		userRepo:   userRepo,
		log:        log.NewHelper(log.With(logger, "module", "server/biz")),
	}
}

func (uc *ServerUsecase) CreateServer(ctx context.Context, server *Server) (*Server, error) {
	uc.log.WithContext(ctx).Infof("CreateServer: %v", server.Name)
	return uc.serverRepo.Save(ctx, server)
}

func (uc *ServerUsecase) UpdateServer(ctx context.Context, server *Server) (*Server, error) {
	uc.log.WithContext(ctx).Infof("UpdateServer: %v", server.ID)
	return uc.serverRepo.Update(ctx, server)
}

func (uc *ServerUsecase) DeleteServer(ctx context.Context, id uint) error {
	uc.log.WithContext(ctx).Infof("DeleteServer: %d", id)
	return uc.serverRepo.Delete(ctx, id)
}

func (uc *ServerUsecase) GetServerById(ctx context.Context, id uint) (*Server, error) {
	uc.log.WithContext(ctx).Infof("GetServerById: %d", id)
	return uc.serverRepo.FindByID(ctx, id)
}

func (uc *ServerUsecase) ListAllServers(ctx context.Context) ([]*Server, error) {
	uc.log.WithContext(ctx).Infof("ListAllServers")
	return uc.serverRepo.ListAll(ctx)
}

func (uc *ServerUsecase) ListServersByOwnerId(ctx context.Context, ownerId uint) ([]*Server, *User, error) {
	uc.log.WithContext(ctx).Infof("ListServersByOwnerId: %d", ownerId)
	res, err := uc.serverRepo.FindByOwnerId(ctx, ownerId)
	if err != nil {
		return nil, nil, err
	}
	user, err := uc.userRepo.FindByID(ctx, ownerId)
	if err != nil {
		return nil, nil, err
	}
	return res, user, nil

}

func (uc *ServerUsecase) GetServerByAddress(ctx context.Context, address string) (*Server, error) {
	uc.log.WithContext(ctx).Infof("GetServerByAddress: %s", address)
	return uc.serverRepo.FindByAddress(ctx, address)
}
