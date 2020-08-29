package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
)

type Ping struct {

}

func (p *Ping) Pong(ctx context.Context, req *string, rsp *string) error {
	*rsp = "pong"
	return nil
}

func main() {
	// 创建一个新的服务
	srv := micro.NewService(
		micro.Name("go.micro.srvEg"),
	)

	//初始化
	srv.Init()

	// set handler
	micro.RegisterHandler(srv.Server(), new(Ping))

	//运行服务
	if err := srv.Run(); err != nil {
		log.Fatalln(err)
	}
}
