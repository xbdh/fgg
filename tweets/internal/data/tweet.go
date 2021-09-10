package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"tweets/internal/biz"
)

type tweetRepo struct {
	data *Data
	log  *log.Helper
}

func (tr *tweetRepo) GetTweetByUserId(ctx context.Context, userId uint64) ([]*biz.Tweet, error) {
	tweets:=[]*biz.Tweet{}

	result:=tr.data.db.WithContext(ctx).Where("user_id = ?",userId).Find(&tweets)
	return tweets,result.Error
}

func (tr *tweetRepo) GetTweetByTweetId(ctx context.Context, tweetId uint64) (*biz.Tweet, error) {
	tweet:=biz.Tweet{}

	result:=tr.data.db.WithContext(ctx).Where("id = ?",tweetId).Find(&tweet)
	return &tweet,result.Error
}

func (tr *tweetRepo) CreateTweet(ctx context.Context, tweet *biz.Tweet) error {
	result:=tr.data.db.WithContext(ctx).Create(tweet)

	// 待处理
	if result.Error==nil{
		newsfeed:=&biz.NewsFeed{
			UserId: tweet.UserId,
			TweetId: uint64(tweet.ID),
		}
		b,err :=json.Marshal(newsfeed)
		if err!=nil{
			return err
		}

		tr.data.kp.Input()<-&sarama.ProducerMessage{
			Topic:     "tweet",
			Value:     sarama.ByteEncoder(b),
		}

		select {
			case err:=<-tr.data.kp.Errors():
				tr.log.Error("kafka send data wrong",err)
			default:
				tr.log.Info("kafka send data success")

		}
		fmt.Println("send ok")
	}
	return result.Error
}

// NewGreeterRepo .
func NewTweetRepo(data *Data, logger log.Logger) biz.TweetRepo {
	return &tweetRepo{
		data: data,
		log:  log.NewHelper(logger),
	}

}


