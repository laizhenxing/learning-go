package module

import (
	"fmt"
	"net"
	"strconv"

	"webcrawler/errors"
)

// 组件网络地址的类型
// 实现了 net.Addr 接口
type mAddr struct {
	// 网络协议
	network string
	// 网络地址
	address string
}

// 用于获取访问组件需遵循的网络协议
func (ma *mAddr) Network() string {
	return ma.network
}

// 用于获取组件的网络地址
func (ma *mAddr) String() string {
	return ma.address
}

// 创建并返回一个网络地址值
func NewAddr(network, ip string, port uint64) (net.Addr, error) {
	if network != "http" && network != "https" {
		errMsg := fmt.Sprintf("illegal network for module address: %s", network)
		return nil, errors.NewIllegalParameterError(errMsg)
	}
	if ip := net.ParseIP(ip); ip == nil {
		return nil, errors.NewIllegalParameterError(fmt.Sprintf("illegal ip for module adress: %s", ip))
	}
	return &mAddr{
		network: network,
		address: ip + ":" + strconv.Itoa(int(port)),
	}, nil
}
