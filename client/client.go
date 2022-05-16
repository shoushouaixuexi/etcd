package main

import (
	pb "ServiceRegistrationDiscovery/ServiceRegistrationDiscovery"
	"ServiceRegistrationDiscovery/discovery"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var srvId = []string{"10003", "10004"}

// addr  = flag.String("addr", "106.53.97.81:10001", "the address to connect to")
//srvId = flag.String("srvId", "10001", "get srvinfo from etcd")

func main() {
	//log.Printf("nihao")
	flag.Parse()
	log.Printf("开始服务发现：")
	// 开始发现服务
	disc, err := discovery.NewEtcdDiscovery()
	if err != nil {
		log.Fatalf("did not connect etcd: %v", err)
	}
	srvaddr, err := disc.GetAddrFromServiceId(srvId)
	if err != nil {
		log.Fatalf("Service can not find: %v", err)
		return
	}
	log.Printf("需要访问的服务地址信息: %s", srvaddr)
	log.Printf("服务发现结束！")
	// 输出服务信息

	// Set up a connection to the server.
	conn, err := grpc.Dial(srvaddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var request pb.ServiceRequest
	request.Operation = append(request.Operation, &pb.Operation{
		FirstNumber:   3,
		SecondNumber:  2,
		ServiceNumber: 0,
	})
	log.Printf("调用服务的参数为%v,%v", 3, 2)
	log.Printf("发起rpc服务调用：")
	r, err := c.ServiceMethod(ctx, &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Code: %v", r.GetCode())
	log.Printf("Greeting: %s", r.GetMsg())
	for _, value := range r.Result {
		log.Printf("result: %v", value)
	}
	log.Printf("rpc服务调用结束！")
}
