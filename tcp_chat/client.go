package tcp_chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var isLogin = false

func Client(listenConfig *AddressConfig) {
	fmt.Println("准备开始建立连接")
	conn, err := net.Dial("tcp", listenConfig.GetAddr())
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Println("连接成功")
	defer conn.Close()

	go receiveMsgFromServer(conn)

	login(conn)

	inputReader := bufio.NewReader(os.Stdin)
	for {
		if !isLogin {
			continue
		}
		fmt.Print("local：")
		readString, _ := inputReader.ReadString('\n')
		trim := strings.Trim(readString, "\r\n")
		if trim == "" {
			continue
		}
		if strings.ToUpper(trim) == "Q" {
			return
		}
		_, err = conn.Write([]byte(readString))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func login(conn net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入名称：")
	name, _ := inputReader.ReadString('\n')

	loginMsg := NewLoginMsg(LoginMsg)
	loginMsg.SetContent(name)
	err := loginMsg.SendMsg(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 接收消息
func receiveMsgFromServer(conn net.Conn) {
	for {
		receiveMsg := ReceiveMsg(conn)
		if receiveMsg == nil {
			continue
		}

		if receiveMsg.GetMsgType() == LoginMsg && receiveMsg.GetContent() == "success" {
			isLogin = true
		}
		if !isLogin {
			if receiveMsg.GetMsgType() == LoginMsg {
				fmt.Println(receiveMsg.GetContent())
			}
			login(conn)
			continue
		}

		fmt.Println("\nserver：", receiveMsg.GetContent())
		fmt.Print("local：")
	}
}
