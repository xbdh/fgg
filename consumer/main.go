package main

import (
	"context"
	"encoding/json"
	"fmt"
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	"sync"
	fv1 "consumer/friendship/v1"
	nv1 "consumer/newsfeed/v1"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

type NewsFeed struct {
	UserId uint64    `json:"user_id"`
	TweetId uint64   `json:"tweet_id"`
}

func main() {

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
	friendshipCliet:=fv1.NewFriendshipClient(conn1)

	conn2,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///newsfeeds"),
		grpc.WithDiscovery(r),
	)

	if err!=nil{
		fmt.Println("newsfeed" ,err)
		return
	}
	newsfeedsCliet:=nv1.NewNewsFeedClient(conn2)

	newsfeed:=NewsFeed{}

	//创建新的消费者
	consumer, err := sarama.NewConsumer([]string{"172.20.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("fail to start consumer", err)
	}

	// 指定分区
	//cp,err:=consumer.ConsumePartition("tweet",0,sarama.OffsetNewest)
	//
	//defer consumer.Close()
	//
	//for {
	//	select {
	//		case msg:=<-cp.Messages():
	//			fmt.Println(string(msg.Value))
	//
	//	case <-cp.Errors():
	//		fmt.Println(err)
	//		break
	//
	//	}
	//}

	// 根据topic获取所有的分区列表
	partitionList, err := consumer.Partitions("tweet")
	if err != nil {
		fmt.Println("fail to get list of partition,err:", err)
	}
	fmt.Println(partitionList)
	//遍历所有的分区
	for p := range partitionList {
		//针对每一个分区创建一个对应分区的消费者
		pc, err := consumer.ConsumePartition("tweet", int32(p), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", p, err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		//异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("partition:%d Offse:%d Key:%v Value:%s \n",
					msg.Partition, msg.Offset, msg.Key, msg.Value)
				json.Unmarshal(msg.Value,&newsfeed)

				followers,err:=friendshipCliet.Followers(context.Background(),&fv1.FollowerRequest{ToUserId: newsfeed.UserId})
				if err!=nil{
					fmt.Println(err)
					return
				}
				fmt.Println(followers.FriendshipInfo)

				for i:=0;i<len(followers.FriendshipInfo);i++{
					resl,err:=newsfeedsCliet.CreateNewsFeed(context.Background(),&nv1.CreateRequest{
						UserId:  followers.FriendshipInfo[i].FromUserId,
						TweetId: newsfeed.TweetId,
					})
					if err!=nil{
						fmt.Println("newsfeed send wrong")
					}
					fmt.Println(resl)

				}

			}
		}(pc)
	}

	wg.Wait()
}
// newsfeed 待更新