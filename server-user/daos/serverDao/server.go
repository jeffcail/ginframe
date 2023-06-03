package serverDao

import "github.com/jeffcail/ginframe/server-user/models"

// CheckIpIsNewServer 检查当前启动服务是否是新服务
func CheckIpIsNewServer(ip, serverName string) (*models.ServerFLag, error) {
	m := new(models.ServerFLag)
	s, err := m.FindUserByServerNameAndIP(serverName, ip)
	if err != nil {
		return nil, err
	}
	return s, nil
}
