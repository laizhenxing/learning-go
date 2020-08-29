package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"userApi/config"
	"userApi/model"
	v "userApi/pkg/version"
	"userApi/router"
	"userApi/router/middleware"
)

var (
	cfg = pflag.StringP("config", "c", "", "userApi config file path.")
	version = pflag.BoolP("version", "v", false, "show version info")
)

func main() {
	pflag.Parse()
	// 查看 version 信息
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	g := gin.New()
	// 设置运行模式，从配置文件中读取配置
	gin.SetMode(viper.GetString("runmode"))

	// 全局中间件
	middlewares := []gin.HandlerFunc{
		middleware.RequestId(),
		middleware.Logging(),
	}
	// 加载中间件
	router.Load(g, middlewares...)

	// 检查路由的健康状态
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response: ", err)
		}
		log.Info("the router has been deployed successfully.")
	}()

	// start to listen the incoming requests
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	// 支持 https
	if cert != "" && key != "" {
		go func() {
			addr := viper.GetString("tls.addr")
			log.Infof("Start to listening incoming requests on http address: %s", addr)
			log.Infof(http.ListenAndServeTLS(addr, cert, key, g).Error())
		}()
	}

	log.Infof("Start to listening incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < 10; i++ {
		rsp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && rsp.StatusCode == 200 {
			return nil
		}

		time.Sleep(time.Second)
		// 睡眠1s继续请求
		log.Infof("Waiting for the router retry 1 second. %s", viper.GetString("url")+"/sd/health")
	}
	return errors.New("Cannot connect to the router.")
}
