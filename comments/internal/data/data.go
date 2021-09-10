package data

import (
	"comments/internal/biz"
	"comments/internal/conf"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCommentRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB

}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
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
	if err:=gormDB.AutoMigrate(biz.Comment{});err!=nil{
		log.Fatalf("fail to connect mysql :%v",err)
	}


	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}
	d:=&Data{
		db: gormDB,
	}
	cleanup := func (){
		//log.Info("closing the data resources")
		log.Info("message", "closing the data resources")


	}
	return d, cleanup, nil
}

