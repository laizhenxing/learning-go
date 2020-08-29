package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	"grpc-micro/plugins/wrapper"
	"grpc-micro/proto/prod"
	"grpc-micro/router"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2380"),
	)

	// rpc 服务客户端
	rpcCliSrv := micro.NewService(
		micro.Name("go.micro.rpcCliSrv"),
		// 使用 wrapper对客户端进行修饰
		micro.WrapClient(wrapper.NewLogWrapper),
		// 使用 prodsWrapper进行超时熔断处理修饰
		micro.WrapClient(wrapper.NewProdsWrapper),
	)
	// rpc 服务端
	prodSrv := prod.NewProdService("go.micro.grpc", rpcCliSrv.Client())

	r := router.NewRouter(prodSrv)

	httpSrv := web.NewService(
		web.Name("go.micro.httpSrv"),
		web.Handler(r),
		web.Address(":8080"),
		web.Registry(reg),
	)

	httpSrv.Init()
	httpSrv.Run()
}
