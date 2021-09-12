package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	tweetV1 "web/api/tweet/v1"
	"web/global"
	"web/models"
)

func PostTweet(c *gin.Context)  {
	tweet:=models.Tweet{}

	if err:=c.ShouldBindJSON(&tweet);err!=nil{
		log.Println("解析tweet数据错误")

		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}

	fmt.Println(tweet)
	resp,err:=global.TweetCli.CreateTweets(context.Background(),&tweetV1.CreateRequest{
		UserId:        tweet.UserId,
		Content:       tweet.Content,

	})

	if err!=nil{
		log.Println("create tweet wrong")
		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}
	c.JSON(200,gin.H{

		"data":resp.TweetInfo,
	})


}

func GetTweetById(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))

	resp,err:=global.TweetCli.GetTweetByTweetId(context.Background(),&tweetV1.TweetIdRequest{
		TweetId: uint64(id),
	})
	if err!=nil{
		log.Println("get tweet by id wrong")
		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}
	c.JSON(200,gin.H{

		"data":resp.TweetInfo,
	})

}

func GetTweetByUserId(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))

	resp,err:=global.TweetCli.GetTweetByUserId(context.Background(),&tweetV1.UserIdRequest{
		UserId: uint64(id),
	})
	if err!=nil{
		log.Println("get tweet by user id wrong")
		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}
	c.JSON(200,gin.H{

		"data":resp.TweetInfo,
	})

}
