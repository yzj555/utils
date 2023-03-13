package tcp_chat

import "strconv"

type AddressConfig struct {
	addrIp   string
	addrPort int
}

func (t *AddressConfig) SetIp(ip string) {
	t.addrIp = ip
}

func (t *AddressConfig) SetPort(port int) {
	t.addrPort = port
}

func (t *AddressConfig) GetAddr() string {
	return t.addrIp + ":" + strconv.Itoa(t.addrPort)
}

func NewAddrConfig(ip string, port int) *AddressConfig {
	return &AddressConfig{addrIp: ip, addrPort: port}
}
