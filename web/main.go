package main

import (
	"github.com/gin-gonic/gin"
	"web/handlers"
)





func main() {
	router := gin.Default()
	// 使用kratos中间件
	//router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	userGruop:=router.Group("/v1/user")
	{
		userGruop.POST("/signup",handlers.Register)
		userGruop.POST("/signin",handlers.PassWordLogin)
		userGruop.GET("/name/:name")
		userGruop.GET("/id/:id")
	}

	tweetGroup:=router.Group("/v1/tweet")
	{
		tweetGroup.POST("/",handlers.PostTweet)
		tweetGroup.GET("/:id",handlers.GetTweetById)
		tweetGroup.GET("/user/:id",handlers.GetTweetByUserId)

	}


	friendshipGroup:=router.Group("/v1/friendship")
	{
		friendshipGroup.POST("/",handlers.PostFriend)
		friendshipGroup.GET("/follower/:userid",handlers.Follower)
		friendshipGroup.GET("/following/:userid",handlers.Following)
	}

	newsfeedGroup:=router.Group("/v1/newsfeed")
	{
		newsfeedGroup.POST("/")
		newsfeedGroup.GET("/list/:id",handlers.List)
	}

router.Run(":8888")

	//httpSrv := http.NewServer(
	//	http.Address(":8000"),
	//
	//	)
	//
	//httpSrv.HandlePrefix("/", router)
	//
	//
	//app := kratos.New(
	//	kratos.Name("gin"),
	//	kratos.Server(
	//		httpSrv,
	//	),
	//)
	//if err := app.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
