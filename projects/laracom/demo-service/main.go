package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"

	proto "github.com/laizhenxing/laracom/demo-service/proto/demo"
)

const (
	address  = "localhost:9091"
	grpcPort = ":9999"
	httpPort = "8002"
	appName  = "Demo Service"
)

// 服务处理器
type DemoService struct {

}
// 实现DemoServiceHandler服务处理器
func (d *DemoService) SayHello(c context.Context, req *proto.DemoRequest, rsp *proto.DemoResponse) error {
	rsp.Text = "hello, " + req.Name
	return nil
}

func main() {
	// 生成一个服务实例
	srv := micro.NewService(micro.Name("laracom.demo.service"))
	// 初始化
	srv.Init()
	//注册服务处理器
	proto.RegisterDemoServiceHandler(srv.Server(), &DemoService{})
	// 运行
	if err := srv.Run(); err != nil {
		log.Fatalf("服务启动失败：%v", err)
	}
}
