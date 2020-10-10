package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todolist/pkg/errno"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{})  {
	code, msg := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func EmptyResult() interface{} {
	return struct {}{}
}