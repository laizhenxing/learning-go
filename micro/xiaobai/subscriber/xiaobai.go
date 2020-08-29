package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	xiaobai "xiaobai/proto/xiaobai"
)

type Xiaobai struct{}

func (e *Xiaobai) Handle(ctx context.Context, msg *xiaobai.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *xiaobai.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
