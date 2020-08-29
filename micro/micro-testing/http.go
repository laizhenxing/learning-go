package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	HelloMicro "micro-testing/proto/hellomicro"
	"micro-testing/router"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2380"),
	)

	srv := micro.NewService(
		micro.Name("httpsrv"),
	)

	rpcSrv := HelloMicro.NewHelloMicroService("hellomicro", srv.Client())

	r := router.NewRouter(rpcSrv)

	httpSrv := web.NewService(
		web.Name("httpClient"),
		web.Address(":8081"),
		web.Handler(r),
		web.Registry(reg),
	)

	httpSrv.Init()
	httpSrv.Run()
}
