package main

import (
	"os"

	"todolist/config"
	"todolist/model"
	"todolist/router"
)

func main() {
	// 初始化配置文件
	if err := config.InitConfig(); err != nil {
		os.Exit(1)
	}

	// 初始化数据库连接
	if err := model.InitDB(); err != nil {
		os.Exit(1)
	}
	defer model.DB.Close()

	r := router.InitRouter()

	r.Run(":8080")
}
