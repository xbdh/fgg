package data

import (
	"context"
	"accounts/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) ListUser(ctx context.Context, ids []int64) ([]*biz.User, error) {
	users:=[]*biz.User{}
	result:=u.data.db.WithContext(ctx).Where("id IN ?", ids).Find(&users)
	return users,result.Error
}

func (u *userRepo) GetById(ctx context.Context, id int64) (*biz.User, error) {
	user:=biz.User{}
	result:=u.data.db.WithContext(ctx).First(&user,id)
	return &user, result.Error
}

func (u *userRepo) GetByName(ctx context.Context, name string) (*biz.User, error) {
	user:=biz.User{}
	result:=u.data.db.WithContext(ctx).First(&user,name)
	return &user, result.Error
}


func (u *userRepo) CreateUser(ctx context.Context, User *biz.User) error {

	result:=u.data.db.Create(User)
	return result.Error
}

func (u *userRepo) UpdateUser(ctx context.Context, id int64, User *biz.User) error {
	panic("implement me")
}

func (u *userRepo) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}


