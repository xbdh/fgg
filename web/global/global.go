package global

import (
	"context"
	"fmt"

	friendshipV1 "web/api/friendship/v1"
	newsfeedV1 "web/api/newsfeed/v1"
	tweetV1 "web/api/tweet/v1"
	accountV1 "web/api/user/v1"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
)
var (
	FriendshipCli friendshipV1.FriendshipClient
	TweetCli     tweetV1.TweetsClient
	NewsfeedCLi  newsfeedV1.NewsFeedClient
	AccountCli  accountV1.UserClient
)
func init(){

	//服务注册与发现
	c:=consulAPI.DefaultConfig()
	c.Address="127.0.0.1:8500"
	c.Scheme="http"
	cli,err:=consulAPI.NewClient(c)
	if err!=nil{
		fmt.Println(err)
		return
	}
	r:=consul.New(cli,consul.WithHealthCheck(false))

	conn1,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///friendships"),
		grpc.WithDiscovery(r),
	)


	if err!=nil{
		fmt.Println(err)
		return
	}
	FriendshipCli=friendshipV1.NewFriendshipClient(conn1)

	conn2,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///newsfeeds"),
		grpc.WithDiscovery(r),
	)

	if err!=nil{
		fmt.Println("newsfeed" ,err)
		return
	}
	NewsfeedCLi=newsfeedV1.NewNewsFeedClient(conn2)

	conn3,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///accounts"),
		grpc.WithDiscovery(r),
	)

	if err!=nil{
		fmt.Println("newsfeed" ,err)
		return
	}
	AccountCli=accountV1.NewUserClient(conn3)


	conn4,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tweets"),
		grpc.WithDiscovery(r),
	)

	if err!=nil{
		fmt.Println("newsfeed" ,err)
		return
	}
	TweetCli =tweetV1.NewTweetsClient(conn4)



}

