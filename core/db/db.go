package db

import (
	"fmt"
	"github.com/jeffcail/ginframe/common/global"
	_gorm "github.com/jeffcail/ginframe/pkg/gorm"
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func InitDb() {
	arg := []int{global.Config.Mysql.MaxOpenConns, global.Config.Mysql.MaxIdleConns}
	Db, err = _gorm.InitGormMysql(global.Config.Mysql.DbDsn, arg)
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql connection success...")
}

// InitRedisClient 初始化redis 连接
func InitRedisClient() (*redis.Client, error) {
	addr := global.Config.Redis.Addr
	password := global.Config.Redis.Password
	db := global.Config.Redis.Db
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("redis connection success...")
	return client, nil
}
