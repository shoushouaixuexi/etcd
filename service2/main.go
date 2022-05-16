package main

import (
	pb "ServiceRegistrationDiscovery/ServiceRegistrationDiscovery"
	"ServiceRegistrationDiscovery/register"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port    = flag.Int("port", 10002, "The server port")
	ip      = flag.String("ip", "106.53.97.81", "The server ip")
	srvName = flag.String("id", "10002", "The server id")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) ServiceMethod(ctx context.Context, in *pb.ServiceRequest) (*pb.ServiceReply, error) {

	var result pb.ServiceReply
	for _, x := range in.GetOperation() {
		result.Result = append(result.Result, x.GetFirstNumber()+x.GetSecondNumber())
	}
	result.Msg = ""
	result.Code = 0
	return &result, nil
}

func main() {
	// 启动服务并且监听地址
	flag.Parse()
	// ------ 开始将服务信息注册到etcd
	log.Printf("server 注册开始  ")
	srvRegister, err := register.NewEtcdRegister()
	if err != nil {
		log.Println(err)
		return
	}
	// 程序退出自动关闭注册器链接
	defer srvRegister.Close()
	// 服务监听地址
	srvAddr := *ip + fmt.Sprintf(":%d", *port)
	//fmt.Printf(srvAddr)
	fmt.Printf("service id:%s,service ip-port:%s\n", fmt.Sprintf(":%d", *port), srvAddr)
	srvRegister.RegisterServer(*srvName, srvAddr, 60)
	if err != nil {
		log.Printf("register error %v \n", err)
		return
	}
	log.Printf("server 注册结束  ")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}

}
