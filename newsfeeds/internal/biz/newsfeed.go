package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"

	"gorm.io/gorm"
)

type NewsFeed struct {
	gorm.Model
	UserId uint64
	TweetId uint64
}

type Tweet struct {
	Id uint64
	UserId uint64
	Content string
	LikesCount uint64
	CommentsCount uint64
	CreateAt time.Time
}

type AllFeed struct {
	Nf *NewsFeed
	T *Tweet
}

type NewsFeedRepo interface {
	ListFollowTweet(ctx context.Context,userId uint64) ([]*AllFeed,error)

	CreateNewsfeed(ctx context.Context, newsfeed *NewsFeed) error
}

type NewsfeedUsecase struct {
	repo NewsFeedRepo
	log  *log.Helper
}

func (nuc *NewsfeedUsecase) ListFollowTweet(ctx context.Context, userId uint64) ([]*AllFeed, error) {
	return nuc.repo.ListFollowTweet(ctx,userId)
}

func (nuc *NewsfeedUsecase) CreateNewsfeed(ctx context.Context, newsfeed *NewsFeed) error {
	return nuc.repo.CreateNewsfeed(ctx,newsfeed)
}

func NewGreeterUsecase(repo NewsFeedRepo, logger log.Logger) * NewsfeedUsecase {
	return & NewsfeedUsecase{repo: repo, log: log.NewHelper(logger)}
}


