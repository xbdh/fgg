package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	UserId uint64
	Content string
	LikesCount uint64
	CommentsCount uint64
}

type NewsFeed struct {
	 UserId uint64    `json:"user_id"`
	 TweetId uint64   `json:"tweet_id"`
}

type TweetRepo interface {

	GetTweetByUserId(ctx context.Context, userId uint64) ([]*Tweet, error)

	CreateTweet(ctx context.Context, tweet *Tweet) error

	GetTweetByTweetId(ctx context.Context, tweetId uint64) (*Tweet, error)
}

type TweetUsecase struct {
	repo TweetRepo
	log  *log.Helper
}

func (tuc *TweetUsecase) GetTweetByUserId(ctx context.Context, userId uint64) ([]*Tweet, error) {
	return tuc.repo.GetTweetByUserId(ctx,userId)
}

func (tuc *TweetUsecase) GetTweetByTweetId(ctx context.Context, tweetId uint64) (*Tweet, error) {
	return tuc.repo.GetTweetByTweetId(ctx,tweetId)
}

func (tuc *TweetUsecase) CreateTweet(ctx context.Context, tweet *Tweet) error {
	return tuc.repo.CreateTweet(ctx,tweet)
}

func NewTweetUsecase(repo TweetRepo, logger log.Logger) *TweetUsecase {
	return &TweetUsecase{repo: repo, log: log.NewHelper(logger)}
}


