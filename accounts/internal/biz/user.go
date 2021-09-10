package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Password string
	Email string

}

type UserRepo interface {

	ListUser(ctx context.Context,ids []int64) ([]*User, error)
	GetById(ctx context.Context, id int64) (*User, error)
	GetByName(ctx context.Context, name string) (*User, error)

	CreateUser(ctx context.Context, User *User) error

	UpdateUser(ctx context.Context, id int64, User *User) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}


func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}


func (u *UserUsecase) ListUser(ctx context.Context, ids []int64) ([]*User, error) {

	return u.repo.ListUser(ctx,ids)
}

func (u *UserUsecase) GetById(ctx context.Context, id int64) (*User, error) {
	panic("implement me")
}

func (u *UserUsecase) GetByName(ctx context.Context, name string) (*User, error) {

	return u.repo.GetByName(ctx,name)
}

func (u *UserUsecase) CreateUser(ctx context.Context, User *User) error {
	return u.repo.CreateUser(ctx,User)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, id int64, User *User) error {
	panic("implement me")
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}



