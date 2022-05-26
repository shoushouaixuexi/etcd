package discovery

import (
	"ServiceRegistrationDiscovery/lb"
	"context"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"

	//"sgrpc/grpc/lb"
	"sync"
)

type etcdResolver struct {
	ctx        context.Context     // 进程上下文
	cancel     context.CancelFunc  // 取消操作
	cc         resolver.ClientConn // 根据服务地址信息建立的链接
	etcdClient *clientv3.Client    // etcd链接句柄
	scheme     string              // 解析器简介
	ipPool     sync.Map            // 存储前缀对应的服务信息地址
}

// ResolveNow 实现了 resolver.Resolver.ResolveNow 方法
// 再次调用解析器解析目标的提示信息
func (e *etcdResolver) ResolveNow(resolver.ResolveNowOptions) {
	log.Println("etcd resolver resolve now")
}

// Close 实现了 resolver.Resolver.Close 方法
// 关闭解析器
func (e *etcdResolver) Close() {
	log.Println("etcd resolver close")
	e.cancel()
}

// watcher 轮询并更新指定前缀服务的实例变化
func (e *etcdResolver) watcher() {

	watchChan := e.etcdClient.Watch(context.Background(), "/"+e.scheme, clientv3.WithPrefix())

	for {
		select {
		case val := <-watchChan:
			for _, event := range val.Events {
				switch event.Type {
				case 0: // 0 是有数据增加
					e.store(event.Kv.Key, event.Kv.Value)
					log.Println("put:", string(event.Kv.Key))
					e.updateState()
				case 1: // 1是有数据减少
					log.Println("del:", string(event.Kv.Key))
					e.del(event.Kv.Key)
					e.updateState()
				}
			}
		case <-e.ctx.Done():
			return
		}

	}
}

// 新增服务实例地址信息到map
func (e *etcdResolver) store(k, v []byte) {
	e.ipPool.Store(string(k), string(v))
}

// 删除服务实例地址信息
func (e *etcdResolver) del(key []byte) {
	e.ipPool.Delete(string(key))
}

//更新 resolver.Resolver.ClientConn 配置（链接权重）
func (e *etcdResolver) updateState() {
	var addrList resolver.State

	// 模拟权重设置
	log.Println("etcdResolver updateSate called...")
	var i = 1
	// 为ipPool的每个值调用函数func() 返回flase截至
	e.ipPool.Range(func(k, v interface{}) bool {
		tA, ok := v.(string)
		if !ok {
			return false
		}
		log.Printf("conn.UpdateState key[%v];val[%v]\n", k, v)

		// 模拟设置权重
		addr := resolver.Address{
			BalancerAttributes: attributes.New(lb.WeightAttributeKey{}, lb.WeightAddrInfo{
				Weight: i,
			}),
			Addr: tA,
		}

		addrList.Addresses = append(addrList.Addresses, addr)
		// 地址的权重逐渐增加
		i++
		return true
	})
	// UpdateState 适当地更新 ClientConn 的状态,既更新链接对应的权重
	e.cc.UpdateState(addrList)
}
