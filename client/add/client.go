package main

import (
	pb "ServiceRegistrationDiscovery/ServiceRegistrationDiscovery"
	"ServiceRegistrationDiscovery/discovery"
	"ServiceRegistrationDiscovery/lb"
	"context"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var (
	srvName = flag.String("id", "etcd://sub/", "The server id")
)

func main() {

	// 注册自定义ETCD解析器
	etcdResolverBuilder := discovery.NewEtcdResolverBuilder()
	// resolver定义了 gRPC 中名称解析的 API
	// Register 将解析器构建器注册到解析器映射。
	resolver.Register(etcdResolverBuilder)
	// Set up a connection to the server.
	//"etcd:///"-> etcd//authority/endpoint"
	conn, err := grpc.Dial(*srvName, grpc.WithBalancerName(lb.WEIGHT_LOAD_BALANCE), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	firstNumber := 1
	secondNumber := 10
	for i := 0; i < 10; i++ {

		var request pb.ServiceRequest
		request.Operation = append(request.Operation, &pb.Operation{
			FirstNumber:   int32(firstNumber),
			SecondNumber:  int32(secondNumber),
			ServiceNumber: 0,
		})
		//  打印请求参数
		fmt.Println(request.Operation)
		r, err := c.ServiceMethod(ctx, &request)
		if err != nil {
			fmt.Println(err)
		}
		//  打印响应结果
		fmt.Println(r.Result)
		//firstNumber++
		//secondNumber++
		time.Sleep(2 * time.Second)
	}
	// fmt.Println("rpc服务调用结束！")
}
