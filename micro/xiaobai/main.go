package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"

	"xiaobai/handler"
	"xiaobai/subscriber"

	xiaobai "xiaobai/proto/xiaobai"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.xiaobai"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	xiaobai.RegisterXiaobaiHandler(service.Server(), new(handler.Xiaobai))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.xiaobai", service.Server(), new(subscriber.Xiaobai))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
