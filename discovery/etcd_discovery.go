package discovery

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdDiscovery struct {
	cli *clientv3.Client // etcd连接句柄
	ctx context.Context  //上下文
}

func NewEtcdRegister() (*EtcdDiscovery, error) {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://106.53.97.81:12379", "http://106.53.97.81:22379", "http://106.53.97.81:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Printf("new etcd client failed,error %v \n", err)
		return nil, err
	}

	ctx, _ := context.WithCancel(context.Background())

	srv := &EtcdDiscovery{
		cli: client,
		ctx: ctx,
	}
	return srv, nil
}

func (srv *EtcdDiscovery) GetAddrFromServiceId(key string) (string, error) {
	resp, err := srv.cli.Get(srv.ctx, key)
	if err != nil {
		return "", err
	}
	addr := ""
	for _, val := range resp.Kvs {
		fmt.Printf("GET response：%s=%s\n", val.Key, val.Value)
		addr = string(val.Value)
	}
	return addr, nil
}
