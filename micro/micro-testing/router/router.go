package router

import (
	"github.com/gin-gonic/gin"
	"micro-testing/middleware"

	HelloMicro "micro-testing/proto/hellomicro"
	v1 "micro-testing/router/v1"
)

func NewRouter(srv HelloMicro.HelloMicroService) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.InitMiddleware(srv))

	v1G := r.Group("/v1")
	{
		v1G.POST("/hello", v1.Hello)
	}

	return r
}
