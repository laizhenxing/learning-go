package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	xiaobai "xiaobai/proto/xiaobai"
)

type Xiaobai struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Xiaobai) Call(ctx context.Context, req *xiaobai.Request, rsp *xiaobai.Response) error {
	log.Info("Received Xiaobai.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Xiaobai) Stream(ctx context.Context, req *xiaobai.StreamingRequest, stream xiaobai.Xiaobai_StreamStream) error {
	log.Infof("Received Xiaobai.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&xiaobai.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Xiaobai) PingPong(ctx context.Context, stream xiaobai.Xiaobai_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&xiaobai.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
