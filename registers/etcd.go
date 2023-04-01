package registers

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Etcd struct {
	client *clientv3.Client
}

func (e *Etcd) Register(info RegisterInfo) {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{info.Addr()},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	e.client = cli
}

func (e *Etcd) RegisterClient(info RegisterInfo) {
	add, err := e.client.Cluster.MemberAdd(context.Background(), []string{"http://" + info.Addr()})
	if err != nil {
		panic(err)
	}
	log.Println(add.Member)
}

func (e *Etcd) GetIp(serverName string) (info RegisterInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (e *Etcd) Close() {
	//TODO implement me
	panic("implement me")
}

func (e *Etcd) RegisterInfo() RegisterInfo {
	//TODO implement me
	panic("implement me")
}
