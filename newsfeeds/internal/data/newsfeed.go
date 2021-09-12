package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "newsfeeds/api/tweet/v1"
	"newsfeeds/internal/biz"
)

type newsfeedRepo struct {
	data *Data
	log  *log.Helper
}

func (n *newsfeedRepo) ListFollowTweet(ctx context.Context, userId uint64) ([]*biz.AllFeed, error) {
	newsfeeds:=[]*biz.NewsFeed{}

	//倒序
	result:=n.data.db.WithContext(ctx).Where("user_id = ?",userId).Order("created_at DESC").Find(&newsfeeds)

	allfeeds:=[]*biz.AllFeed{}

	if result.Error==nil{
		for i:=0;i<len(newsfeeds);i++{
			allfeed:=biz.AllFeed{}
			tweet:=biz.Tweet{}
			tweetIdReply,err:=n.data.tc.GetTweetByTweetId(ctx,&v1.TweetIdRequest{
				TweetId: newsfeeds[i].TweetId,
			})
			if err!=nil{
				return nil, err
			}
			tweet.Id=tweetIdReply.TweetInfo.Id
			tweet.UserId=tweetIdReply.TweetInfo.UserId
			tweet.Content=tweetIdReply.TweetInfo.Content
			//tweet.LikesCount=tweetIdReply.TweetInfo.LikesCount
			//tweet.CommentsCount=tweetIdReply.TweetInfo.CommentsCount
			tweet.CreateAt=tweetIdReply.TweetInfo.CreatedAt.AsTime()

			allfeed.Nf=newsfeeds[i];
			allfeed.T=&tweet
			allfeeds=append(allfeeds,&allfeed )
		}


	}
	return allfeeds,result.Error
}

func (n *newsfeedRepo) CreateNewsfeed(ctx context.Context, newsfeed *biz.NewsFeed) error {

	result := n.data.db.WithContext(ctx).Create(newsfeed)
	return result.Error
}

// NewGreeterRepo .
func NewNewsfeedRepo(data *Data, logger log.Logger) biz.NewsFeedRepo {
	return &newsfeedRepo{

		data: data,
		log:  log.NewHelper(logger),
	}
}


