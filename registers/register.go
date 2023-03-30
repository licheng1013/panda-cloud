package registers

import "fmt"

// Register 注册中心必须实现的接口！
type Register interface {
	// Register 注册远程中心信息
	Register(info RegisterInfo)
	// RegisterClient 注册本地客户端信息
	RegisterClient(info RegisterInfo)
	// GetIp 获取ip
	GetIp(serverName string) (info RegisterInfo, err error)
	// Close 用于关机等操作
	Close()
	// RegisterInfo 注册信息
	RegisterInfo() RegisterInfo
}

type RegisterInfo struct {
	// ip
	Ip string
	// 端口
	Port uint16
	// 当前注册服务名
	ServiceName string
}

func (v RegisterInfo) Addr() string {
	return v.Ip + ":" + fmt.Sprint(v.Port)
}
