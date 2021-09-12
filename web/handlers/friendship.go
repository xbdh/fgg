package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	friendshipV1 "web/api/friendship/v1"
	"web/global"
	"web/models"
)

func PostFriend(c *gin.Context)  {
	friendship:=models.Friendship{}

	if err:=c.ShouldBindJSON(&friendship);err!=nil{
		log.Println("解析 friendship 数据错误")

		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}

	resp,err:=global.FriendshipCli.CreateFriendship(context.Background(),&friendshipV1.CreateRequest{
		FromUserId: friendship.FromId,
		ToUserId:   friendship.ToId,
	})

	if err!=nil{
		log.Println("create friendship wrong")
		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}
	c.JSON(200,gin.H{

		"data":resp.FriendshipInfo,
	})


}

func Follower(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("userid"))

	resp,err:=global.FriendshipCli.Followers(context.Background(),&friendshipV1.FollowerRequest{
		ToUserId: uint64(id),
	})



	if err!=nil{
		log.Println("get follower wrong")
		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}
	c.JSON(200,gin.H{

		"data":resp.FriendshipInfo,
	})

}

func Following(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("userid"))
	fmt.Println(id)
	resp,err:=global.FriendshipCli.Followings(context.Background(),&friendshipV1.FollowingRequest{
		FromUserId: uint64(id),
	})



	if err!=nil{
		log.Println("get following wrong")
		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}
	c.JSON(200,gin.H{

		"data":resp.FriendshipInfo,
	})


}
