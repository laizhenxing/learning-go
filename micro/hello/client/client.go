package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	greeter "learning-micro/hello/proto"
)

func main() {
	// 创建一个新的服务
	service := micro.NewService(micro.Name("Greeter.client"))
	service.Init()

	// 创建新的客户端
	cli := greeter.NewGreeterService("greeter", service.Client())

	// 调用greeter
	rsp, err := cli.Hello(context.TODO(), &greeter.HelloRequest{Name: "xiaoli"})
	if err != nil {
		fmt.Println(err)
	}

	// 打印响应
	fmt.Println(rsp.Greeting)
}
