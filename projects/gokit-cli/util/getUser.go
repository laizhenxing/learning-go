package util

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	csLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httpTransport "github.com/go-kit/kit/transport/http"
	csapi "github.com/hashicorp/consul/api"

	"gokit-cli/constance"
	"gokit-cli/services"
)

func GetUser() (string, error) {
	//第一步，创建client
	config := csapi.DefaultConfig()
	config.Address = constance.RegisterCenter // 注册中心地址
	apiCli, err := csapi.NewClient(config)
	if err != nil {
		return "", err
	}
	cli := consul.NewClient(apiCli)

	// 第二步，创建一个 instancer
	var logger csLog.Logger
	logger = csLog.NewLogfmtLogger(os.Stdout)
	tags := []string{"primary"}
	// 可实时查询服务实例的状态信息
	instancer := consul.NewInstancer(cli, logger, "userService", tags, true)

	// 第三步，创建 Endpointer
	factory := func(serviceUrl string) (endpoint.Endpoint, io.Closer, error) {
		fmt.Println("serviceUrl: ", serviceUrl)
		tgt, _ := url.Parse("http://" + serviceUrl)
		return httpTransport.NewClient("GET", tgt, services.GetUserInfoRequest, services.GetUserInfoResponse).Endpoint(), nil, nil
	}
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	// 判断是否有可用的服务
	//endpoints, err := endpointer.Endpoints()
	//if err != nil {
	//	return "", nil
	//}
	//fmt.Println("服务有", len(endpoints), "条")
	//if len(endpoints) == 0 {
	//	fmt.Println("没有发现可用服务")
	//	//os.Exit(1)
	//	return "", err
	//}

	// 第四步，执行
	// 使用 go-kit 轮询
	// 创建一个负载均衡器
	//mylb := lb.NewRoundRobin(endpointer)	// 轮询器
	mylb := lb.NewRandom(endpointer, time.Now().UnixNano()) // 随机器
	getUserInfo, err := mylb.Endpoint()
	if err != nil {
		//fmt.Println("获取服务出错:", err)
		return "", err
	}
	ctx := context.Background()
	res, err := getUserInfo(ctx, services.UserRequest{Uid: 102})
	if err != nil {
		//fmt.Println("执行服务错误：", err)
		return "", err
	}
	userInfo := res.(services.UserResponse)
	fmt.Printf("%v\n", userInfo)
	return userInfo.Result, nil
}
