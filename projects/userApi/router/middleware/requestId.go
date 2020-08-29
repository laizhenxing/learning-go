package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		c.Request.Header.Set("X-Request-Id", requestId)

		log.Info("request info", lager.Data{"X-Request-Id": requestId, "URI": c.Request.RequestURI, "Host": c.Request.Host})
		c.Next()
	}
}
