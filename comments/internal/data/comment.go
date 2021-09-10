package data

import (
	"comments/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func (c *commentRepo) ListTweetComment(ctx context.Context, tweetId uint64) ([]*biz.Comment, error) {
	comments:=[]*biz.Comment{}

	result:=c.data.db.WithContext(ctx).Where("tweet_id = ?",tweetId).Find(&comments)

	return comments, result.Error
}

func (c *commentRepo) CreateComment(ctx context.Context, comment *biz.Comment) error {
	result:=c.data.db.WithContext(ctx).Create(comment)

	return result.Error
}

// NewCommentRepo .
func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

