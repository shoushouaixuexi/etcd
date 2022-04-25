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

var (
	// addr  = flag.String("addr", "106.53.97.81:10001", "the address to connect to")
	srvId = flag.String("srvId", "10001", "get srvinfo from etcd")
)

func main() {
	//log.Printf("nihao")
	flag.Parse()

	// 开始发现服务
	disc, err := discovery.NewEtcdRegister()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	srvaddr, err := disc.GetAddrFromServiceId(*srvId)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("srvaddr: %s", srvaddr)
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
		FirstNumber:   2,
		SecondNumber:  3,
		ServiceNumber: 0,
	})
	r, err := c.ServiceMethod(ctx, &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Code: %v", r.GetCode())
	log.Printf("Greeting: %s", r.GetMsg())
	for _, value := range r.Result {
		log.Printf("result: %v", value)
	}

}
