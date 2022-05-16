package discovery

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdDiscovery struct {
	cli *clientv3.Client // etcd连接句柄
	ctx context.Context  //上下文
}

func NewEtcdDiscovery() (*EtcdDiscovery, error) {

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
func RetRandAddr(key []string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(key))
	return key[num], nil
}
func (srv *EtcdDiscovery) GetAddrFromServiceId(key []string) (string, error) {
	var addr = []string{}
	for _, val := range key {
		resp, err := srv.cli.Get(srv.ctx, val)
		if err != nil {
			return "", err
		}
		for _, val := range resp.Kvs {
			//fmt.Printf("GET response：%s=%s\n", val.Key, val.Value)
			if string(val.Value) != "" {
				addr = append(addr, string(val.Value))
			}
		}
	}
	if len(addr) == 0 {
		return "", errors.New("发现服务失败，key对应服务不存在")
	}
	for _, srvaddr := range addr {
		fmt.Printf("可使用的addr%s\n", srvaddr)
	}
	raddr, _ := RetRandAddr(addr)
	return raddr, nil
}
