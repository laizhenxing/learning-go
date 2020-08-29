package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"grpc-micro/handler"
	"grpc-micro/proto/prod"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2380"),
	)

	srv := micro.NewService(
		micro.Name("go.micro.grpc"),
		micro.Address(":8085"),
		micro.Registry(reg),
	)

	srv.Init()

	err := prod.RegisterProdServiceHandler(srv.Server(), new(handler.ProdService))
	if err != nil {
		log.Fatal(err)
	}

	srv.Run()
}
