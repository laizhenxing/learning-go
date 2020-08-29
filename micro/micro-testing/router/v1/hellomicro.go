package v1

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	HelloMicro "micro-testing/proto/hellomicro"
)

func Hello(c *gin.Context)  {
	var req HelloMicro.HelloRequest
	c.Bind(&req)

	srv := c.Keys["srv"].(HelloMicro.HelloMicroService)

	rsp, err := srv.HelloMicro(context.Background(), &req)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": rsp.Data})
}
