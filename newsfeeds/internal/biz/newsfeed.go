package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type NewsFeed struct {
	ID        uint64           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time		`json:"created_at"`
	UserId uint64                `json:"user_id"`
	TweetId uint64               `json:"tweet_id"`
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


