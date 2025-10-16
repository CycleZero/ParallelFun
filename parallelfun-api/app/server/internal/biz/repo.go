package biz

import "context"

type ServerRepo interface {
	Save(context.Context, *Server) (*Server, error)
	Update(context.Context, *Server) (*Server, error)
	FindByID(context.Context, uint) (*Server, error)
	Delete(context.Context, uint) error
	ListAll(context.Context) ([]*Server, error)
	FindByOwnerId(context.Context, uint) ([]*Server, error)
	FindByAddress(context.Context, string) (*Server, error)
}
