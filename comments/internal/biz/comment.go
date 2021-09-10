package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId uint64
	TweetId uint64
	Content string
}

type CommentRepo interface {
	ListTweetComment(ctx context.Context,tweetId uint64) ([]*Comment,error)

	CreateComment(ctx context.Context, comment *Comment) error
}


type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func (c *CommentUsecase) ListTweetComment(ctx context.Context, tweetId uint64) ([]*Comment, error) {
	return c.repo.ListTweetComment(ctx,tweetId)
}

func (c *CommentUsecase) CreateComment(ctx context.Context, comment *Comment) error {
	return c.repo.CreateComment(ctx,comment)
}

func NewGreeterUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{repo: repo, log: log.NewHelper(logger)}
}

