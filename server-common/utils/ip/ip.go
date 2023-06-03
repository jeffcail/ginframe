package _ip

import "net"

var ip = ""

func GetIp() string {
	if ip != "" {
		return ip
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("获取本机地址失败:" + err.Error())
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.IsPrivate() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return ip
			}
		}
	}
	panic("获取本机地址失败")
}
