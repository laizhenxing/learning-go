package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GetShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if reqId, ok := v.(string); ok {
		return reqId
	}
	return ""
}
