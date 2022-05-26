package discovery

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type etcdResolverBuilder struct {
	etcdClient *clientv3.Client
}

func NewEtcdResolverBuilder() *etcdResolverBuilder {

	// 创建etcd客户端连接
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://106.53.97.81:12379", "http://106.53.97.81:22379", "http://106.53.97.81:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Println("client get etcd failed,error", err)
		panic(err)
	}

	return &etcdResolverBuilder{
		etcdClient: etcdClient,
	}
}

// 客户端通过 Dial 方法对指定服务进行拨号时，grpc resolver 查找注册的 Builder 实例调用其 Build() 方法构建自定义 Resolver
func (erb *etcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {

	// 获取指定前缀的etcd节点值
	// etcd://sub/-> /etcd-sub
	// /etcd-sub->/etcd/sub-service-1  /etcd->/etcd/sub-service-2
	prefix := "/" + target.URL.Scheme + "-" + target.URL.Host

	log.Println(prefix)
	//log.Println(target.URL.Scheme)
	// 获取 etcd 中服务保存的ip列表
	res, err := erb.etcdClient.Get(context.Background(), prefix, clientv3.WithPrefix())

	if err != nil {
		log.Println("Build etcd get addr failed; err:", err)
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	// 创建etcd解析节点
	es := &etcdResolver{
		cc:         cc,
		etcdClient: erb.etcdClient,
		ctx:        ctx,
		cancel:     cancelFunc,
		scheme:     target.URL.Scheme,
	}

	// 将获取到的ip和port保存到本地的map中
	log.Printf("etcd res:%+v\n", res)
	for _, kv := range res.Kvs {
		es.store(kv.Key, kv.Value)
	}
	// fmt.Println(res.Kvs)
	// 更新拨号里的ip列表
	es.updateState()

	// 监听etcd中的服务是否变化
	go es.watcher()
	return es, nil
}

// Scheme 实现了 resolver.Builder.Scheme 方法
// Scheme 方法定义了 resolver 的协议名

func (erb *etcdResolverBuilder) Scheme() string {
	return "etcd"
}
