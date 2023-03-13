package tcp_chat

import "net"

type Connect struct {
	Name    string
	IsLogin bool
	Conn    net.Conn
}
