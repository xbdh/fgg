package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"

	newsfeedV1 "web/api/newsfeed/v1"
	"web/global"

	"context"
)



func List(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	resp,err:=global.NewsfeedCLi.ListFollowTweet(context.Background(),&newsfeedV1.ListRequest{
		UserId: uint64(id),
	})



	if err!=nil{
		log.Println("get newsfeed wrong")
		c.JSON(400,gin.H{
			"data":"wrong",
		})
		return
	}
	c.JSON(200,gin.H{

		"data":resp.NewFeedInfo,
	})

}

