package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	gkLog "github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"

	"gokit/services"
	"gokit/util"
)

func main() {
	// 使用命令行参数指定serviceName,port
	name := flag.String("name", "", "服务名称")
	port := flag.Int("p", 0, "服务端口")
	flag.Parse()
	if *name == "" {
		log.Fatal("请指定服务名称")
	}
	if *port == 0 {
		log.Fatal("请指定服务端口")
	}
	util.SetServiceNameAndPort(*name, *port)

	user := &services.UserService{} // 服务
	limit := rate.NewLimiter(1, 5)  // 支持最大请求5,每秒1个
	var logger gkLog.Logger
	endp := services.RateLimit(limit)(
		services.UserLogMiddleware(logger)(
			services.CheckTokenMiddleware()(
				services.GenUserEndpoint(user),
			),
		),
	)

	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorEncoder(services.MyErrorEncoder),
	}
	serviceHandler := httpTransport.NewServer(endp, services.DecodeUserRequest, services.EncodeUserResponse, options...)

	// 增加 handler 用于获取用户 token
	accSrv := &services.AccessService{}
	accSrvEndpoint := services.AccessEndpoint(accSrv)
	accSrvHandler := httpTransport.NewServer(accSrvEndpoint, services.DecodeAccessRequest, services.EncodeAccessResponse, options...)

	// 使用第三方mux路由
	r := mux.NewRouter()
	{
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serviceHandler)
		r.Methods("GET").Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status": "ok"}`))
		})
		// 获取token
		r.Methods("POST").Path(`/access_token`).Handler(accSrvHandler)
	}

	errChan := make(chan error)
	go func() {
		// 注册consul服务
		util.RegisterService()
		err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()
	// 监听退出信号
	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sigChan)
	}()

	err := <-errChan
	util.UnRegisterService()
	fmt.Println("exit: ", err)
}
