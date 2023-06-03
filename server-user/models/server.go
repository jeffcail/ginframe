package models

import (
	"github.com/jeffcail/ginframe/server-user/core"
	"time"
)

type ServerFLag struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Flag       string    `gorm:"column:flag;NOT NULL;comment:'服务'"`
	ServerName string    `gorm:"column:server_name;NOT NULL;comment:'服务名称'"`
	ServerIp   string    `gorm:"column:server_ip;NOT NULL;comment:'服务ip'"`
	Sum        int       `gorm:"column:sum;NOT NULL;comment:'服务累计数量'"`
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL;comment:'创建时间'"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL;comment:'更新时间'"`
}

func (se *ServerFLag) TableName() string {
	return "server_flag"
}

const (
	NotFound = "record not found"
)

// FindUserByServerNameAndIP 根据 server 查找server服务信息
func (se *ServerFLag) FindUserByServerNameAndIP(serverName, ip string) (*ServerFLag, error) {
	s := &ServerFLag{}
	err := core.Db.Where("server_name = ?", serverName).Where("server_ip = ?", ip).First(s).Error
	if err != nil && err.Error() != NotFound {
		return nil, err
	}
	return s, nil
}
