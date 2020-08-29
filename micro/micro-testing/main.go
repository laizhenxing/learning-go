package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"micro-testing/handler"
	_ "micro-testing/model"
	"micro-testing/proto/user"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2379"),
	)

	srv := micro.NewService(
		//micro.Name("api.com.hellomicro"),
		// 规则：在Namespace的基础上需要加(.api); 如 namespace=user,这里的服务名称就是 user.api.随意
		micro.Name("user.api.myapp"),
		micro.Address(":8083"),
		micro.Registry(reg),
	)

	// 初始化
	srv.Init()

	// 注册handler
	//HelloMicro.RegisterHelloMicroServiceHandler(srv.Server(), new(handler.HelloMicroService))
	user.RegisterUserServiceHandler(srv.Server(), new(handler.UserHandler))

	srv.Run()
}
