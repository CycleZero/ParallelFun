package service

import (
	"context"
	"gorm.io/gorm"
	"parallelfun-api/app/article/internal/biz"

	pb "parallelfun-api/api/article/v1"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
	uc *biz.ArticleUsecase
}

func NewArticleService(uc *biz.ArticleUsecase) *ArticleService {
	return &ArticleService{uc: uc}
}

func (s *ArticleService) GetArticleById(ctx context.Context, req *pb.GetArticleByIdRequest) (*pb.GetArticleByIdReply, error) {
	res, _ := s.uc.GetArticleById(ctx, req.Id)
	author, err := s.uc.GetAuthorById(ctx, res.AuthorID)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleByIdReply{
		Article: &pb.ArticleInfo{
			Id:      uint64(res.ID),
			Title:   res.Title,
			Content: res.Content,
			Author: &pb.Author{
				Id:   author.ID,
				Name: author.Name,
			},
		},
	}, nil
}
func (s *ArticleService) GetArticleList(ctx context.Context, req *pb.GetArticleListRequest) (*pb.GetArticleListReply, error) {
	_, err := s.uc.GetArticleListByPage(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	// TODO: 批量获取作者信息
	return &pb.GetArticleListReply{}, nil
}
func (s *ArticleService) NewArticle(ctx context.Context, req *pb.NewArticleRequest) (*pb.NewArticleReply, error) {
	res, err := s.uc.NewArticle(ctx, &biz.Article{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: req.AuthorId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.NewArticleReply{
		Id:       uint64(res.ID),
		AuthorId: res.AuthorID,
		Success:  true,
	}, nil
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	res, err := s.uc.UpdateArticle(ctx, &biz.Article{
		Model:   gorm.Model{ID: uint(req.Id)},
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateArticleReply{
		Article: s.toArticleInfo(ctx, res),
	}, nil
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	err := s.uc.DeleteArticle(ctx, &biz.Article{Model: gorm.Model{ID: uint(req.Id)}})
	if err != nil {
		return &pb.DeleteArticleReply{Success: false}, err
	}
	return &pb.DeleteArticleReply{Success: true}, nil
}
func (s *ArticleService) GetArticleListByUserId(ctx context.Context, req *pb.GetArticleListByUserIdRequest) (*pb.GetArticleListByUserIdReply, error) {
	_, err := s.uc.GetArticleListByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleListByUserIdReply{
		//todo 批量获取作者信息
	}, nil
}
func (s *ArticleService) GetArticleListByTitle(ctx context.Context, req *pb.GetArticleListByTitleRequest) (*pb.GetArticleListByTitleReply, error) {
	return &pb.GetArticleListByTitleReply{}, nil
}

func (s *ArticleService) UploadMedia(ctx context.Context, req *pb.UploadMediaRequest) (*pb.UploadMediaReply, error) {
	r, err := s.uc.GenerateUploadUrl(ctx, req.UserId, req.FileName, req.Type)
	if err != nil {
		return nil, err
	}
	return &pb.UploadMediaReply{
		Url: r,
	}, nil
}

func (s *ArticleService) toArticleInfo(ctx context.Context, article *biz.Article) *pb.ArticleInfo {
	author, err := s.uc.GetAuthorById(ctx, article.AuthorID)
	if err != nil {
		return nil
	}

	return &pb.ArticleInfo{
		Id:      uint64(article.ID),
		Title:   article.Title,
		Content: article.Content,
		Author: &pb.Author{
			Id:   author.ID,
			Name: author.Name,
		},
	}
}
