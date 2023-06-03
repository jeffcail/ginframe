package _redis

import (
	"gopkg.in/redis.v5"
)

// InitRedis 初始化redis
func InitRedis(addr, pass string, rdb int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       rdb,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
