package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Name)
	return uc.repo.Save(ctx, u)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id uint) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %d", id)
	return uc.repo.FindByID(ctx, id)
}

func (uc *UserUseCase) FindByEmail(ctx context.Context, email string) (*User, error) {
	uc.log.WithContext(ctx).Infof("FindByEmail: %s", email)
	return uc.repo.FindByEmail(ctx, email)
}

func (uc *UserUseCase) ListByName(ctx context.Context, name string) ([]*User, error) {
	uc.log.WithContext(ctx).Infof("ListByName: %s", name)
	return uc.repo.ListByName(ctx, name)
}

func (uc *UserUseCase) GetRole(ctx context.Context, id uint) (Role, error) {
	uc.log.WithContext(ctx).Infof("GetRole: %d", id)
	return uc.repo.GetRole(ctx, id)
}

func (uc *UserUseCase) BatchFindByGameId(ctx context.Context, gameIds []string) ([]*User, error) {
	uc.log.WithContext(ctx).Infof("BatchFindByGameId: %v", gameIds)
	return uc.repo.BatchGetUserByGameId(ctx, gameIds)
}

func (uc *UserUseCase) FindByGameId(ctx context.Context, gameId string) (*User, error) {
	uc.log.WithContext(ctx).Infof("FindByGameId: %s", gameId)
	return uc.repo.GetUserByGameId(ctx, gameId)
}

func (uc *UserUseCase) BatchFindById(ctx context.Context, ids []uint) ([]*User, error) {
	uc.log.WithContext(ctx).Infof("BatchFindById: %v", ids)
	// 注意：UserRepo 接口中没有定义 BatchFindById 方法，需要补充
	// 或者在 service 层直接调用其他方法实现
	return nil, nil
}
func (uc *UserUseCase) UpdateUser(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", u.ID)
	return uc.repo.Update(ctx, u)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, u *User) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", u.ID)
	return uc.repo.Delete(ctx, u)
}

func (uc *UserUseCase) Register(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("Register: %v", u.Name)
	// 这里应该包含密码加密等注册逻辑
	return uc.repo.Save(ctx, u)
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (*User, string, error) {
	uc.log.WithContext(ctx).Infof("Login: %s", email)
	// 这里应该包含密码验证和token生成逻辑
	user, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}
	// 验证密码逻辑...
	token := "generated_token" // 实际应该生成真实的token
	return user, token, nil
}

func (uc *UserUseCase) Logout(ctx context.Context, id string, token string) error {
	uc.log.WithContext(ctx).Infof("Logout: %s", id)
	// 这里应该包含token失效等登出逻辑
	return nil
}
