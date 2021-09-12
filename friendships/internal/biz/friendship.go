package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Friendship struct {
	ID        uint64           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time		`json:"created_at"`
	FromId uint64          `json:"from_id"`
	ToId uint64            `json:"to_id"`
}

type FriendshipRepo interface {

	GetFollowings(ctx context.Context,fromUserId uint64) ([]*Friendship,error)

	GetFollowers(ctx context.Context,toUserId uint64) ([]*Friendship,error)

	CreateFriendship(ctx context.Context, friendship *Friendship) error
}

type FriendshipUsecase struct {
	repo FriendshipRepo
	log  *log.Helper
}

func (fuc *FriendshipUsecase) GetFollowings(ctx context.Context, fromUserId uint64) ([]*Friendship, error) {
	return fuc.repo.GetFollowings(ctx,fromUserId)
}

func (fuc *FriendshipUsecase) GetFollowers(ctx context.Context, toUserId uint64) ([]*Friendship, error) {
	return fuc.repo.GetFollowers(ctx,toUserId)
}

func (fuc *FriendshipUsecase) CreateFriendship(ctx context.Context, friendship *Friendship) error {
	return fuc.repo.CreateFriendship(ctx,friendship)
}

func NewFriendshipUsecase(repo FriendshipRepo, logger log.Logger) *FriendshipUsecase {
	return &FriendshipUsecase{repo: repo, log: log.NewHelper(logger)}
}


