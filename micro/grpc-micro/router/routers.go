package router

import (
	"github.com/gin-gonic/gin"

	"grpc-micro/middleware"
	"grpc-micro/proto/prod"
	v1 "grpc-micro/router/v1"
)

func NewRouter(prodSrv prod.ProdService) *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.InitMiddleware(prodSrv), middleware.ErrorMiddleware())

	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", v1.GetProdList)
		v1Group.Handle("GET","/prod/:pid", v1.GetProdDetail)
	}

	return r

}