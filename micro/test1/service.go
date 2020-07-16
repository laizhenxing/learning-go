package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"

	test1 "learning-micro/test1/proto"
)

type T1 struct {

}

func (t *T1) Hello(ctx context.Context, req *test1.HelloRequest, rsp *test1.HelloResponse) error {
	rsp.Rsp = "Test micro~" + req.Name
	return nil
}

func main() {
	// 创建一个服务
	srv := micro.NewService(micro.Name("test1"))

	// 初始化
	srv.Init()

	// 注册处理器
	test1.RegisterT1ServiceHandler(srv.Server(), new(T1))

	// 运行
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}