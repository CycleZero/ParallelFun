package biz

import "context"

type ServerRepo interface {
	Save(context.Context, *Server) (*Server, error)
	Update(context.Context, *Server) (*Server, error)
	FindByID(context.Context, uint) (*Server, error)
}
