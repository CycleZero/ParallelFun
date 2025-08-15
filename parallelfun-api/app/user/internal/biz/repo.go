package biz

import "context"

type UserRepo interface {
	FindByID(ctx context.Context, id int64) (*User, error)

	ListAll(ctx context.Context) ([]*User, error)

	Save(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) (*User, error)
	Delete(ctx context.Context, u *User) error
	ListByName(ctx context.Context, name string) ([]*User, error)
}
