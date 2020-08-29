package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	"grpc-micro-client/proto/prod"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2380"),
	)
	r := gin.Default()

	httpSrv := web.NewService(
		web.Name("go.micro.httpSrv"),
		web.Address(":8080"),
		web.Handler(r),
		web.Registry(reg),
	)

	rpcCliSrv := micro.NewService(micro.Name("go.micro.rpcCliSrv"))

	prodRpcSrv := prod.NewProdService("go.micro.grpc", rpcCliSrv.Client())

	v1 := r.Group("/v1")
	{
		v1.Handle("POST", "/prods", func(c *gin.Context) {
			var req prod.ProdoRequest
			err := c.Bind(&req)
			if err != nil {
				c.JSON(400, gin.H{ "msg": err.Error() })
			}

			rsp, _ := prodRpcSrv.GetProdList(context.Background(), &req)

			c.JSON(200, gin.H{ "data": rsp.Data })
		})
	}

	httpSrv.Init()

	httpSrv.Run()
}
