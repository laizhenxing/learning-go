package v1

import (
	"context"
	"github.com/gin-gonic/gin"

	"grpc-micro/proto/prod"
)

func PanicError(err error)  {
	if err != nil {
		panic(err)
	}
}

func GetProdList(c *gin.Context) {
	var (
		req prod.ProdRequest
		rsp *prod.ProdResponse
		//err error
	)

	PanicError(c.Bind(&req))


	prodSrv := c.Keys["prodSrv"].(prod.ProdService)

	rsp, _ = prodSrv.GetProdList(context.Background(), &req)

	// 熔断代码(服务器通信出现超时时)
	// 1. 配置config
	//cfg := hystrix.CommandConfig{
	//	Timeout: 1000, // 1s
	//}
	//// 2. 配置command
	//hystrix.ConfigureCommand("getprods", cfg)
	//// 3. 执行，使用Do方法
	//
	//err = hystrix.Do("getprods", func() error {
	//	rsp, err = prodSrv.GetProdList(context.Background(), &req)
	//	return err
	//}, func(e error) error {
	//	rsp = handler.DefaultProds()
	//	return nil
	//}) // 第二个回调是降级方法

	//if err != nil {
	//	c.JSON(500, gin.H{"msg": err.Error()})
	//	return
	//}
	//
	//fmt.Println("rsp", rsp)
	c.JSON(200, gin.H{"data": rsp.Data})
}

func GetProdDetail(c *gin.Context)  {
	var (
		req prod.ProdRequest
		rsp *prod.ProdDetailResponse
		//err error
	)

	PanicError(c.BindUri(&req))

	prodSrv := c.Keys["prodSrv"].(prod.ProdService)

	rsp, _ = prodSrv.GetProdDetail(context.Background(), &req)

	c.JSON(200, gin.H{"data": rsp.Data})
}
