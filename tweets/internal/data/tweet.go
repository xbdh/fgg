package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"tweets/internal/biz"
)

type tweetRepo struct {
	data *Data
	log  *log.Helper
}

func (tr *tweetRepo) GetTweetByUserId(ctx context.Context, userId uint64) ([]*biz.Tweet, error) {
	tweets:=[]*biz.Tweet{}

	// 待做
	//pattern:=fmt.Sprintf("%s:%d","user_id:%dtweet_id:*",userId)
	//
	//stweets,err:=tr.data.rdb.Keys(ctx,pattern).Result()
	//
	//if err==nil {
	//
	//	erri :=json.Unmarshal([]byte(stweet),&tweet)
	//	if erri!=nil{
	//		tr.data.log.Info("解析失败")
	//
	//	}
	//	tr.data.log.Info("get data from redis")
	//	return &tweet,nil
	//}
	result:=tr.data.db.WithContext(ctx).Where("user_id = ?",userId).Find(&tweets)
	return tweets,result.Error
}

func (tr *tweetRepo) GetTweetByTweetId(ctx context.Context, tweetId uint64) (*biz.Tweet, error) {
	tweet:=biz.Tweet{}

	key:=fmt.Sprintf("%s:%d","tweet_id:%d",tweetId)

	stweet,err:=tr.data.rdb.Get(ctx,key).Result()
	if err==nil {
		erri :=json.Unmarshal([]byte(stweet),&tweet)
		if erri!=nil{
			tr.data.log.Info("解析失败")

		}
		tr.data.log.Info("get data from redis")
		return &tweet,nil
	}

	result:=tr.data.db.WithContext(ctx).Where("id = ?",tweetId).Find(&tweet)
	return &tweet,result.Error
}

func (tr *tweetRepo) CreateTweet(ctx context.Context, tweet *biz.Tweet) error {
	result:=tr.data.db.WithContext(ctx).Create(tweet)


	// 待处理
	if result.Error==nil{
		// 缓存
		tweetkey:=fmt.Sprintf("%s:%d","tweet_id:%d",tweet.ID)
		userTweetkey:=fmt.Sprintf("%s:%d","user_id:%dtweet_it%d",tweet.UserId,tweet.ID)
		btweet,err:=json.Marshal(tweet)
		if err!=nil{
			return err
		}


		err=tr.data.rdb.Set(ctx,tweetkey,string(btweet),time.Hour).Err()
		if err!=nil{
			tr.log.Error("写入数据库成功，写入缓存失败: wrong:",err)
			// 不return？？
		}
		err=tr.data.rdb.Set(ctx,userTweetkey,string(btweet),time.Hour).Err()
		if err!=nil{
			tr.log.Error("写入数据库成功，写入缓存失败: wrong:",err)
			// 不return？？
		}
		tr.log.Info("store data to redis success")

		// 写入kafka
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
		//fmt.Println("send ok")
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


