package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	r := gin.Default()

	r.Handle("GET", "/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "index page",
		})
	})

	srv := web.NewService(
		web.Name("go.micro.gin.index"),
		web.Address(":8080"),
		web.Handler(r),
	)
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
