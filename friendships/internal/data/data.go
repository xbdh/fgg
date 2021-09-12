package data

import (
	"context"

	"friendships/internal/biz"
	"friendships/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/go-redis/redis/v8"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewFriendshipRepo,NewCache,NewDB)

type Data struct {
	// TODO wrapped database client
	db *gorm.DB
	rdb *redis.Client


	log *log.Helper

}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	log:=log.NewHelper(logger)


	d:=&Data{
		db: db,
		rdb:rdb,

		log: log,
	}


	cleanup := func (){
		//log.Info("closing the data resources")
		// 关闭kafka redis db待做
		log.Info("message", "closing the data resources")

	}
	return d, cleanup, nil
}



func NewDB(conf* conf.Data,logger log.Logger) *gorm.DB{
	log:=log.NewHelper(logger)

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})

	if err!=nil{
		log.Fatal("failed opening connection to mysql :%v",err)
		return nil
	}

	// 初始化数据表时，只用一次
	if err:=db.AutoMigrate(biz.Friendship{});err!=nil{
		log.Fatalf("fail to connect mysql :%v",err)
	}
	return db
}

func NewCache(conf* conf.Data,logger log.Logger) *redis.Client{
	log:=log.NewHelper(logger)

	rdb:=redis.NewClient(&redis.Options{
		Addr:               conf.Redis.Addr,

	})

	_,err:=rdb.Ping(context.Background()).Result()
	if err!=nil{
		log.Fatal("redis cant ping")
	}
	log.Info("connect to redis")
	return rdb
}