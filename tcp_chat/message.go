package tcp_chat

import (
	"encoding/json"
	"fmt"
	"net"
)

type MsgType int

var LoginMsg MsgType = 0  // 登录消息
var ServerMsg MsgType = 1 // 服务端消息
var ClientMsg MsgType = 2 // 客户端消息

type Message struct {
	MsgType  MsgType
	FromAddr string
	ToAddr   string
	Content  interface{}
}

func NewMsg(from, to string, msgType MsgType) *Message {
	return &Message{FromAddr: from, ToAddr: to, MsgType: msgType}
}

func NewLoginMsg(msgType MsgType) *Message {
	return &Message{MsgType: msgType}
}

func (m *Message) GetMsgType() MsgType {
	return m.MsgType
}

func (m *Message) GetFromAddr() string {
	return m.FromAddr
}

func (m Message) GetToAddr() string {
	return m.ToAddr
}

func (m *Message) SetContent(content interface{}) {
	m.Content = content
}

func (m *Message) GetContent() interface{} {
	return m.Content
}

func (m *Message) SendMsg(conn net.Conn) error {
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = conn.Write(res)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ReceiveMsg(conn net.Conn) *Message {
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil || n == 0 {
		return nil
	}
	data := buf[:n]
	receiveMsg := &Message{}
	err = json.Unmarshal(data, receiveMsg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return receiveMsg
}
