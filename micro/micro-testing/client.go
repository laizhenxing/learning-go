package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	HelloMicro "micro-testing/proto/hellomicro"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2380"),
	)

	// 创建一个新服务
	srv := micro.NewService(
		micro.Name("hellomicro.client"),
		micro.Registry(reg),
	)
	// 初始化
	srv.Init()

	// 创建客户端
	cli := HelloMicro.NewHelloMicroService("hellomicro", srv.Client())

	// 远程调用hellomicro服务的Hello方法
	rsp, err := cli.HelloMicro(context.TODO(), &HelloMicro.HelloRequest{
		Id: 369,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Data)
}
