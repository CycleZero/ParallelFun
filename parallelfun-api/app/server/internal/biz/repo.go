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

type UserRepo interface {
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, uint) (*User, error)
	FindByGameId(context.Context, string) (*User, error)
	BatchFindByGameId(context.Context, []string) ([]*User, error)
	BatchFindById(context.Context, []uint) ([]*User, error)
}
