package registers

import (
	"testing"
)

func TestEtcd(t *testing.T) {
	// 创建一个etcd客户端
	etcd := Etcd{}
	etcd.Register(RegisterInfo{
		Ip:   "localhost",
		Port: 2379,
	})
	etcd.RegisterClient(RegisterInfo{
		Ip:   "localhost",
		Port: 22379,
	})
}
