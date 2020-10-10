package router

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"userApi/handler/sd"
	v1 "userApi/handler/v1"
	"userApi/router/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// pprof router
	pprof.Register(g)

	g.POST("/v1/login", v1.Login) // 用户登录
	g.POST("/v1/user", v1.Create) // 创建用户

	v1Group := g.Group("/v1")
	// 使用身份验证中间价
	v1Group.Use(middleware.AuthMiddleware())
	{

		v1Group.GET("/users", v1.GetList)      // 获取用户列表
		v1Group.GET("/user/:username", v1.Get) // 获取指定用户
		v1Group.PUT("/user/:id", v1.Update)    // 更新用户
		v1Group.DELETE("/user/:id", v1.Delete) // 删除用户
	}

	// the health check handles
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
