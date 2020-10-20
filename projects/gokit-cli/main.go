package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
	csLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httpTransport "github.com/go-kit/kit/transport/http"
	csapi "github.com/hashicorp/consul/api"

	"gokit-cli/constance"
	"gokit-cli/services"
	"gokit-cli/util"
)

func main_1() {
	tgt, _ := url.Parse("http://localhost:8080")
	cli := httpTransport.NewClient("GET", tgt, services.GetUserInfoRequest, services.GetUserInfoResponse)
	getUserInfo := cli.Endpoint()

	ctx := context.Background()
	res, err := getUserInfo(ctx, services.UserRequest{Uid: 101})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	userInfo := res.(services.UserResponse)
	fmt.Printf("%#v", userInfo)
}

func main_2() {
	//第一步，创建client
	config := csapi.DefaultConfig()
	config.Address = constance.RegisterCenter // 注册中心地址
	apiCli, _ := csapi.NewClient(config)
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
	endpoints, _ := endpointer.Endpoints()
	fmt.Println("服务有", len(endpoints), "条")
	if len(endpoints) == 0 {
		fmt.Println("没有发现可用服务")
		os.Exit(1)
	}

	// 第四步，执行

	// 手动轮询
	//for {
	//	i := rand.Intn(len(endpoints))
	//	getUserInfo := endpoints[i]
	//	ctx := context.Background()
	//	res, err := getUserInfo(ctx, services.UserRequest{Uid: 102})
	//	if err != nil {
	//		fmt.Println("执行服务时出错：", err)
	//		os.Exit(2)
	//	}
	//	userInfo := res.(services.UserResponse)
	//	fmt.Printf("%#v\n", userInfo)
	//	time.Sleep(3 * time.Second)
	//}

	// 使用 go-kit 轮询
	// 创建一个负载均衡器
	//mylb := lb.NewRoundRobin(endpointer)	// 轮询器
	mylb := lb.NewRandom(endpointer, time.Now().UnixNano()) // 随机器
	for {
		getUserInfo, err := mylb.Endpoint()
		if err != nil {
			fmt.Println("获取服务出错:", err)
			os.Exit(3)
		}
		ctx := context.Background()
		res, err := getUserInfo(ctx, services.UserRequest{Uid: 102})
		if err != nil {
			fmt.Println("执行服务错误：", err)
			os.Exit(2)
		}
		userInfo := res.(services.UserResponse)
		fmt.Printf("%v\n", userInfo)
		time.Sleep(3 * time.Second)

	}
}

func main() {
	// 使用熔断器
	configA := hystrix.CommandConfig{
		Timeout:                3000,
		MaxConcurrentRequests:  5,  // 最大并发数
		RequestVolumeThreshold: 3,  // 熔断器阀值
		ErrorPercentThreshold:  20, // 错误百分比
		SleepWindow:            10, // 检查服务时间间隔
	}
	hystrix.ConfigureCommand("getUser", configA)
	err := hystrix.Do("getUser", func() error {
		// 业务代码
		userInfo, err := util.GetUser()
		if err != nil {
			return err
		}
		fmt.Println(userInfo)
		return nil
	}, func(err error) error {
		fmt.Println("降级用户")
		// 服务降级
		return err
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
