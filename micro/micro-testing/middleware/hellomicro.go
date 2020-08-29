package middleware

import (
	"github.com/gin-gonic/gin"
	HelloMicro "micro-testing/proto/hellomicro"
)

func InitMiddleware(srv HelloMicro.HelloMicroService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys =  make(map[string]interface{})
		c.Keys["srv"] = srv
		c.Next()
	}
}
