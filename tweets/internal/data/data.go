package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tweets/internal/biz"
	"tweets/internal/conf"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTweetRepo,NewKafkaProducer,NewDB,NewCache)

type Data struct {
	// TODO wrapped database client
	db *gorm.DB
	rdb *redis.Client
	kp  sarama.AsyncProducer

	log *log.Helper

}

// NewData .
func NewData(producer sarama.AsyncProducer,db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	log:=log.NewHelper(logger)


	d:=&Data{
		db: db,
		rdb:rdb,
		kp:producer,
		log: log,
	}


	cleanup := func (){
		//log.Info("closing the data resources")
		// 关闭kafka redis db待做
		log.Info("message", "closing the data resources")

	}
	return d, cleanup, nil
}

func NewKafkaProducer(conf *conf.Data,logger log.Logger) sarama.AsyncProducer {
	c := sarama.NewConfig()
	// 待设置
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}

func NewDB(conf* conf.Data,logger log.Logger) *gorm.DB{
	log:=log.NewHelper(logger)

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})

	if err!=nil{
		log.Fatal("failed opening connection to mysql :%v",err)
	}

	// 初始化数据表时，只用一次
	if err:=db.AutoMigrate(biz.Tweet{});err!=nil{
		log.Fatalf("fail to connect mysql :%v",err)
	}
	return db
}

func NewCache(conf* conf.Data,logger log.Logger) *redis.Client{
	log:=log.NewHelper(logger)

	rdb:=redis.NewClient(&redis.Options{
		Addr:               conf.Redis.Addr,

	})

	log.Info("connect to redis")
	return rdb
}