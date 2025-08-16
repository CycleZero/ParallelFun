package service

import (
	"context"

	pb "parallelfun-api/api/article/v1"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) GetArticleById(ctx context.Context, req *pb.GetArticleByIdRequest) (*pb.GetArticleByIdReply, error) {
	return &pb.GetArticleByIdReply{}, nil
}
func (s *ArticleService) GetArticleList(ctx context.Context, req *pb.GetArticleListRequest) (*pb.GetArticleListReply, error) {
	return &pb.GetArticleListReply{}, nil
}
func (s *ArticleService) NewArticle(ctx context.Context, req *pb.NewArticleRequest) (*pb.NewArticleReply, error) {
	return &pb.NewArticleReply{}, nil
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
func (s *ArticleService) GetArticleListByUserId(ctx context.Context, req *pb.GetArticleListByUserIdRequest) (*pb.GetArticleListByUserIdReply, error) {
	return &pb.GetArticleListByUserIdReply{}, nil
}
func (s *ArticleService) GetArticleListByTitle(ctx context.Context, req *pb.GetArticleListByTitleRequest) (*pb.GetArticleListByTitleReply, error) {
	return &pb.GetArticleListByTitleReply{}, nil
}
