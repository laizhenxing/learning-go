package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"

	test1 "learning-micro/test1/proto"
)

func main() {
	// 创建服务
	srv := micro.NewService(micro.Name("T1.hello"))

	// 初始化
	srv.Init()

	// 创建客户端
	// 这里的name,是一个服务的名称
	cli := test1.NewT1Service("test1", srv.Client())

	// 请求并处理结果
	rsp, err := cli.Hello(context.TODO(), &test1.HelloRequest{
		Name: "cangsf",
	})
	if err != nil {
		fmt.Println(err)
	}
	// 打印结果
	fmt.Println(rsp.Rsp)
}