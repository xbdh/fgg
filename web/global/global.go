package global

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	grpcx "google.golang.org/grpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"time"

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


// Set global trace provider
func setTracerProvider(url string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("web"),
			attribute.String("env", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
func init(){
	if err := setTracerProvider("http://127.0.0.1:14268/api/traces"); err != nil {
		panic(err)
	}


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

		grpc.WithTimeout(2*time.Second),
		// for tracing remote ip recording
		grpc.WithMiddleware(
			tracing.Client(),
		),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)


	if err!=nil{
		fmt.Println("friendship" ,err)
		fmt.Println(err)
		return
	}
	FriendshipCli=friendshipV1.NewFriendshipClient(conn1)

	conn2,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///newsfeeds"),
		grpc.WithDiscovery(r),

		grpc.WithTimeout(2*time.Second),
		// for tracing remote ip recording
		grpc.WithMiddleware(
			tracing.Client(),
		),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
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

		grpc.WithTimeout(2*time.Second),
		// for tracing remote ip recording
		grpc.WithMiddleware(
			tracing.Client(),
		),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)

	if err!=nil{
		fmt.Println("account" ,err)
		return
	}
	AccountCli=accountV1.NewUserClient(conn3)


	conn4,err:=grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tweets"),
		grpc.WithDiscovery(r),

		grpc.WithTimeout(2*time.Second),
		// for tracing remote ip recording
		grpc.WithMiddleware(
			tracing.Client(),
			),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)

	if err!=nil{
		fmt.Println("tweets" ,err)
		return
	}
	TweetCli =tweetV1.NewTweetsClient(conn4)



}

