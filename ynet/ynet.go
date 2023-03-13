package ynet

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

// GetLocalIp 获取本地ip
func GetLocalIp() string {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range adders {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetIPV4 获得外网IP
func GetIPV4() string {
	resp, err := http.Get("https://ipv4.netarm.com")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	return string(content)
}
