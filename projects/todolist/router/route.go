package router

import (
	"net/http"
	v1 "todolist/router/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	g := gin.Default()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	// 告诉gin前端静态资源的路径
	g.Static("/static", "static")
	// 告诉gin模版文件的路径
	g.LoadHTMLGlob("templates/*")

	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := g.Group("/v1")
	{
		v1Group.POST("/todo", v1.AddTodo)
		v1Group.GET("/todo/:id", nil)
		v1Group.GET("/todo/list", nil)
		v1Group.PUT("/todo/:id", nil)
		v1Group.DELETE("/todo/:id", nil)
	}

	return g
}