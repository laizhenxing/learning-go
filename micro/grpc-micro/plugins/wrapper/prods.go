package wrapper

import (
	"context"
	"fmt"
	"grpc-micro/handler"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type prodsWrapper struct {
	client.Client
}

// 超时熔断处理
func (p *prodsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("prodsWrapper开始调用...")

	cmdName := req.Service() + "." + req.Endpoint()

	// 1、配置config
	cfg := hystrix.CommandConfig{
		Timeout: 1000, // 1s
		RequestVolumeThreshold: 2,
		ErrorPercentThreshold: 50,
		SleepWindow: 5000,
	}
	// 2、配置command
	hystrix.ConfigureCommand(cmdName, cfg)

	return hystrix.Do(cmdName, func() error {
		return p.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		handler.DefaultCommonProds(rsp)
		return nil
	})
}

func NewProdsWrapper(c client.Client) client.Client {
	return &prodsWrapper{c}
}