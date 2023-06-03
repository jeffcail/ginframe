package cachedb

import (
	"github.com/jeffcail/ginframe/server-user/core"
)

// SetUserServerFlag 初始化user服务唯一标识
func SetUserServerFlag(ip, serverName string) error {
	return core.Rd.Set(ip, serverName, -1).Err()
}

// GetUserServerFlag 获取user服务唯一标识
func GetUserServerFlag(k string) string {
	result, _ := core.Rd.Get(k).Result()
	return result
}
