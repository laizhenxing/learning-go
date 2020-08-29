package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"

	pb "github.com/laizhenxing/laracom/demo-service/proto/demo"
)

func main() {
	srv := micro.NewService(micro.Name("laracom.demo.cli"))
	srv.Init()

	cli := pb.NewDemoServiceClient("laracom.demo.service", srv.Client())
	rsp, err := cli.SayHello(context.TODO(), &pb.DemoRequest{Name: "兴小狸"})
	if err != nil {
		log.Fatalf("调用 laracom.demo.service 服务失败: %v", err)
		return
	}
	log.Println("go-micro response: ", rsp.Text)
}
