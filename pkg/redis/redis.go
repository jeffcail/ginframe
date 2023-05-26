package _redis

import (
	"github.com/jeffcail/ginframe/common/global"
	"gopkg.in/redis.v5"
)

// InitRedis 初始化redis
func InitRedis() (*redis.Client, error) {
	addr := global.Config.Redis.Addr
	password := global.Config.Redis.Password
	db := global.Config.Redis.Db
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
