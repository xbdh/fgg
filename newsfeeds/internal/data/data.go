package data

import (
	"context"
	"database/sql"
	consul "github.com/go-kratos/consul/registry"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	"newsfeeds/internal/biz"

	"newsfeeds/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	tweetCli "newsfeeds/api/tweet/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewNewsfeedRepo,NewDiscovery, NewTweetsServiceClient)


// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
	tc  tweetCli.TweetsClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger,tc  tweetCli.TweetsClient) (*Data, func(), error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: drv,
	}), &gorm.Config{})

	// 初始化数据表时，只用一次
	if err:=gormDB.AutoMigrate(biz.NewsFeed{});err!=nil{
		log.Fatalf("fail to connect mysql :%v",err)
	}


	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}
	d:=&Data{
		db: gormDB,
		tc: tc,
	}
	cleanup := func (){
		//log.Info("closing the data resources")
		log.Info("message", "closing the data resources")


	}
	return d, cleanup, nil
}

// NewData .




func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewTweetsServiceClient(r registry.Discovery) tweetCli.TweetsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tweets"),
		grpc.WithDiscovery(r),

	)
	if err != nil {
		panic(err)
	}
	return tweetCli.NewTweetsClient(conn)
}