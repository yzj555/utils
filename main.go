package main

import (
	"fmt"
	"utils/base"
)

func main() {
	//localIp := ynet.GetLocalIp()
	//if localIp == "" {
	//	fmt.Println("未获取到有效IP")
	//	return
	//}
	//listenAddr := tcp_chat.NewAddrConfig(localIp, 20000)
	//tcp_chat.NewChatServer().Server(listenAddr)

	//remoteAddr := tcp_chat.NewAddrConfig("cn-cd-dx-7.natfrp.cloud", 57628)
	//tcp_chat.Client(remoteAddr)

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(base.GetMapKeys(myMap))
	fmt.Println(base.GetMapValues(myMap))

}
