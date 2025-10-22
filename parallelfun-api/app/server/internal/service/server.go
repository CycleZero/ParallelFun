package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"parallelfun-api/api"
	pb "parallelfun-api/api/server/v1"
	userpb "parallelfun-api/api/user/v1"
	"parallelfun-api/app/server/internal/biz"
)

type ServerService struct {
	pb.UnimplementedServerServer
	uc  *biz.ServerUsecase
	log *log.Helper
}

func NewServerService(uc *biz.ServerUsecase, logger log.Logger) *ServerService {
	return &ServerService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "server/service")),
	}
}

func (s *ServerService) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	server := &biz.Server{
		Name:    req.Name,
		Address: req.Address,
		Port:    uint(req.Port),
		OwnerId: uint(req.OwnerId),
	}

	createdServer, err := s.uc.CreateServer(ctx, server)
	if err != nil {
		return nil, err
	}

	return &pb.CreateServerReply{
		ServerInfo: &pb.ServerInfo{
			Id:          uint64(createdServer.ID),
			Name:        createdServer.Name,
			Address:     createdServer.Address,
			Port:        uint32(createdServer.Port),
			Status:      int32(createdServer.Status),
			OwnerId:     uint64(createdServer.OwnerId),
			Avatar:      createdServer.Avatar,
			Cover:       createdServer.Cover,
			Description: createdServer.Description,
			Tags:        createdServer.Tags,
		},
	}, nil
}

func (s *ServerService) UpdateServer(ctx context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
	server := &biz.Server{
		Name:        req.ServerInfo.Name,
		Address:     req.ServerInfo.Address,
		Port:        uint(req.ServerInfo.Port),
		Status:      int(req.ServerInfo.Status),
		OwnerId:     uint(req.ServerInfo.OwnerId),
		Avatar:      req.ServerInfo.Avatar,
		Cover:       req.ServerInfo.Cover,
		Description: req.ServerInfo.Description,
		Tags:        req.ServerInfo.Tags,
	}

	updatedServer, err := s.uc.UpdateServer(ctx, server)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateServerReply{
		Id: uint64(updatedServer.ID),
		ServerInfo: &pb.ServerInfo{
			Id:          uint64(updatedServer.ID),
			Name:        updatedServer.Name,
			Address:     updatedServer.Address,
			Port:        uint32(updatedServer.Port),
			Status:      int32(updatedServer.Status),
			OwnerId:     uint64(updatedServer.OwnerId),
			Avatar:      updatedServer.Avatar,
			Cover:       updatedServer.Cover,
			Description: updatedServer.Description,
			Tags:        updatedServer.Tags,
		},
	}, nil
}

func (s *ServerService) DeleteServer(ctx context.Context, req *pb.DeleteServerRequest) (*pb.DeleteServerReply, error) {
	err := s.uc.DeleteServer(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteServerReply{
		Success: true,
	}, nil
}

func (s *ServerService) GetServerById(ctx context.Context, req *pb.GetServerByIdRequest) (*pb.GetServerByIdReply, error) {
	server, err := s.uc.GetServerById(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	if server == nil {
		return &pb.GetServerByIdReply{}, nil
	}

	return &pb.GetServerByIdReply{
		Server: &pb.ServerInfo{
			Id:          uint64(server.ID),
			Name:        server.Name,
			Address:     server.Address,
			Port:        uint32(server.Port),
			Status:      int32(server.Status),
			OwnerId:     uint64(server.OwnerId),
			Avatar:      server.Avatar,
			Cover:       server.Cover,
			Description: server.Description,
			Tags:        server.Tags,
		},
	}, nil
}

func (s *ServerService) ListServer(ctx context.Context, req *pb.ListServerByOwnerIdRequest) (*pb.ListServerByOwnerIdReply, error) {
	servers, user, err := s.uc.ListServersByOwnerId(ctx, uint(req.OwnerId))
	if err != nil {
		return nil, err
	}

	serverInfos := make([]*pb.ServerInfo, 0, len(servers))
	for _, server := range servers {
		serverInfos = append(serverInfos, s.convertToServerInfo(server))
	}

	return &pb.ListServerByOwnerIdReply{
		Servers: serverInfos,
		PageInfo: &api.PageInfoReply{
			Total:      int64(len(serverInfos)),
			PageNum:    req.PageInfo.PageNum,
			PageSize:   req.PageInfo.PageSize,
			TotalPages: 0,
		},
		Owner: &userpb.UserInfo{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Gid:   user.GameId,
			Email: user.Email,
			Role:  int32(user.Role),
		},
	}, nil
}

func (s *ServerService) convertToServerInfo(server *biz.Server) *pb.ServerInfo {
	serverInfo := &pb.ServerInfo{
		Id:          uint64(server.ID),
		Name:        server.Name,
		Address:     server.Address,
		Port:        uint32(server.Port),
		Status:      int32(server.Status),
		OwnerId:     uint64(server.OwnerId),
		CreatedAt:   timestamppb.New(server.CreatedAt),
		UpdatedAt:   timestamppb.New(server.UpdatedAt),
		Avatar:      server.Avatar,
		Cover:       server.Cover,
		Description: server.Description,
		Tags:        server.Tags,
	}

	// 处理 DeletedAt
	if server.DeletedAt.Valid == true {
		serverInfo.DeletedAt = timestamppb.New(server.DeletedAt.Time)
	}

	return serverInfo
}

func (s *ServerService) GetServerOnlinePlayer(ctx context.Context, req *pb.GetServerOnlinePlayerRequest) (*pb.GetServerOnlinePlayerReply, error) {
	// TODO: 实现获取服务器在线玩家逻辑
	s.log.Info("GetServerOnlinePlayer", req)
	return &pb.GetServerOnlinePlayerReply{Players: make([]*pb.ServerPlayerInfo, 10)}, nil
}

func (s *ServerService) AddServerOnlinePlayer(ctx context.Context, req *pb.AddServerOnlinePlayerRequest) (*pb.AddServerOnlinePlayerReply, error) {
	// TODO: 实现添加服务器在线玩家逻辑
	return &pb.AddServerOnlinePlayerReply{}, nil
}

func (s *ServerService) RemoveServerOnlinePlayer(ctx context.Context, req *pb.RemoveServerOnlinePlayerRequest) (*pb.RemoveServerOnlinePlayerReply, error) {
	// TODO: 实现移除服务器在线玩家逻辑
	return &pb.RemoveServerOnlinePlayerReply{}, nil
}

func (s *ServerService) GetServerOnlinePlayerCount(ctx context.Context, req *pb.GetServerOnlinePlayerCountRequest) (*pb.GetServerOnlinePlayerCountReply, error) {
	// TODO: 实现获取服务器在线玩家数量逻辑
	return &pb.GetServerOnlinePlayerCountReply{}, nil
}
