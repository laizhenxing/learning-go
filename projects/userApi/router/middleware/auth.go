package middleware

import (
	"github.com/gin-gonic/gin"

	"userApi/handler"
	"userApi/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, err, handler.EmptyResponse)
			c.Abort()
			return
		}
		c.Next()
	}
}
