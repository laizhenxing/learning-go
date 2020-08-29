package handler

import (
	"context"
	HelloMicro "micro-testing/proto/hellomicro"
	"strconv"
)

type HelloMicroService struct {

}

func (h HelloMicroService) HelloMicro(ctx context.Context, req *HelloMicro.HelloRequest, rsp *HelloMicro.HelloResponse) error {
	rsp.Data = "Call HelloMicro func, Id is: " + strconv.Itoa(int(req.Id))
	return nil
}

