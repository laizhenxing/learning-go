package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grpc-micro/proto/prod"
)

func InitMiddleware(prodSrv prod.ProdService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodSrv"] = prodSrv
		c.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(500, gin.H{"msg": fmt.Sprintf("%s", err)})
				c.Abort()
			}
		}()

		c.Next()
	}
}