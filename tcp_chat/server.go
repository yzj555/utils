package tcp_chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type ChatServer struct {
	connList map[int]net.Conn
	connMap  map[string]*Connect
	currConn int
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		connList: make(map[int]net.Conn),
		connMap:  make(map[string]*Connect),
		//addrList: make(map[int]string, 0),
	}
}

func (cs *ChatServer) Server(listenConfig *AddressConfig) bool {
	fmt.Println("准备开始建立连接")
	listen, err := net.Listen("tcp", listenConfig.GetAddr())
	fmt.Printf("监听端口成功: %v\n", listenConfig.GetAddr())
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return false
	}
	go cs.serverMsgToClient()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go cs.process(conn)
	}
}

// 获取连接地址
func getRemoteAddr(conn net.Conn) string {
	addr := strings.Split(conn.RemoteAddr().String(), ":")
	return addr[0] + ":" + addr[1]
}

// 处理接受到的客户端消息
func (cs *ChatServer) process(conn net.Conn) {
	remoteAddr := getRemoteAddr(conn)
	index := len(cs.connList) + 1
	cs.connList[index] = conn
	fmt.Println("新连接：", remoteAddr)
	cs.printConnList()
	defer conn.Close()
	for {
		receiveMsg := ReceiveMsg(conn)
		if receiveMsg == nil {
			continue
		}
		if receiveMsg.GetMsgType() == LoginMsg {
			loginSuccess := cs.HandleLogin(conn, receiveMsg)
			if !loginSuccess {
				break
			}
			continue
		} else {
			connect, err := cs.connMap[receiveMsg.GetToAddr()]
			if !err || connect.IsLogin {
				toMsg := NewMsg("server", receiveMsg.GetFromAddr(), ServerMsg)
				toMsg.SetContent("目标用户不存在或不在线！")
			} else {
				toMsg := NewMsg(receiveMsg.GetFromAddr(), receiveMsg.GetToAddr(), ClientMsg)
				toMsg.SetContent(receiveMsg.GetContent())
				err := toMsg.SendMsg(connect.Conn)
				if err != nil {
					toMsg.FromAddr = "server"
					toMsg.ToAddr = receiveMsg.GetFromAddr()
					toMsg.MsgType = ServerMsg
					toMsg.SetContent("消息发送失败！")
					toMsg.SendMsg(conn)
					continue
				}
			}
		}
		//fmt.Println("\nclient->", remoteAddr, "发来的数据：", strings.Trim(recvStr, "\r\n"))
	}
	cs.closeConn(index)
	fmt.Println("断开连接：", getRemoteAddr(conn))
	cs.printConnList()
}

func (cs *ChatServer) HandleLogin(conn net.Conn, msg *Message) bool {
	if msg.GetMsgType() != LoginMsg {
		return false
	}
	loginMsgResult := NewLoginMsg(LoginMsg)
	defer loginMsgResult.SendMsg(conn)
	name := msg.GetContent().(string)
	connect, err := cs.connMap[name]
	if err && connect.IsLogin {
		loginMsgResult.SetContent("当前账号已被占用！")
		return false
	}
	connect = &Connect{IsLogin: true, Conn: conn, Name: name}
	cs.connMap[name] = connect
	loginMsgResult.SetContent("success")
	return true
}

// 关闭连接
func (cs *ChatServer) closeConn(index int) {
	delete(cs.connList, index)
	if index == cs.currConn {
		cs.choiceConn()
	}
}

// 服务端发送消息给客户端
func (cs *ChatServer) serverMsgToClient() {
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	trim := strings.Trim(str, "\r\n")
	for {
		if len(cs.connList) == 0 {
			fmt.Println("当前无连接")
			continue
		}
		if 0 >= cs.currConn {
			//fmt.Println("未选择连接！请输入“choice”")
			cs.currConn = cs.choiceConn()
			continue
		}
		nowConn := cs.connList[cs.currConn]
		fmt.Print("server->", getRemoteAddr(nowConn), "：")
		trim, _ = inputReader.ReadString('\n')
		trim = strings.Trim(trim, "\r\n")
		if trim == "choice" {
			cs.currConn = cs.choiceConn()
			continue
		}
		_, err := nowConn.Write([]byte(trim))
		if err != nil {
			continue
		}
		fmt.Println("消息发送成功")
	}
}

// 选择连接
func (cs *ChatServer) choiceConn() int {
	cs.printConnList()
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("请选择连接(回车确认)：")
	str, _ := inputReader.ReadString('\n')
	trim := strings.Trim(str, "\r\n")
	index, _ := strconv.Atoi(trim)
	_, ok := cs.connList[index]
	if !ok {
		fmt.Println("输入值无效！")
		return -1
	}
	return index
}

// 输出打印连接列表
func (cs *ChatServer) printConnList() {
	fmt.Println("已连接列表：")
	for i, coon := range cs.connList {
		fmt.Println(i, "：", getRemoteAddr(coon))
	}
}
