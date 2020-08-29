package main

import (
	"fmt"
	"gin-micro/helper"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	"gin-micro/model"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2380"),
	)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.Handle("GET", "/user", func(c *gin.Context) {
			c.String(200, "user api")
		})
		v1.Handle("POST", "/prods", func(c *gin.Context) {
			var ph helper.ProductHelper
			err := c.Bind(&ph)
			fmt.Println(ph)
			if err != nil || ph.Size <= 0 {
				ph = helper.ProductHelper{
					Size: 2,
				}
			}
			data := model.GenerateProduct(ph.Size)
			fmt.Println("main data: ", data)
			c.JSON(200, gin.H{
				"data": data,
			})
		})
	}

	srv := web.NewService(
		web.Name("go.micro.gin-micro"),
		//web.Address(":8081"),
		web.Handler(r),
		web.Metadata(map[string]string{"protocol": "http"}), // 添加这行代码
		web.Registry(reg),
	)

	//srv.Init()
	srv.Run()
}
