package biz

import "context"

type UserRepo interface {
	FindByID(ctx context.Context, id uint) (*User, error)
	FindByName(ctx context.Context, name string) (*User, error)
	ListAll(ctx context.Context) ([]*User, error)

	FindByEmail(ctx context.Context, email string) (*User, error)

	Save(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) (*User, error)
	Delete(ctx context.Context, u *User) error
	ListByName(ctx context.Context, name string) ([]*User, error)
	GetRole(ctx context.Context, id uint) (Role, error)
	GetUserByGameId(ctx context.Context, gameId string) (*User, error)
	BatchGetUserByGameId(ctx context.Context, gameIds []string) ([]*User, error)
	BatchFindById(ctx context.Context, ids []uint) ([]*User, error)
}
